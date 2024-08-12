const k8s = require('@kubernetes/client-node')
const mustache = require('mustache')
const request = require('request')
const JSONStream = require('json-stream')
const fs = require('fs').promises

// Use Kubernetes client to interact with Kubernetes

const timeouts = {}

const kc = new k8s.KubeConfig();

process.env.NODE_ENV === 'development' ? kc.loadFromDefault() : kc.loadFromCluster()

const opts = {}
kc.applyToRequest(opts)

const client = kc.makeApiClient(k8s.CoreV1Api);

const sendRequestToApi = async (api, method = 'get', options = {}) => {
    return new Promise((resolve, reject) => {
        const url = `${kc.getCurrentCluster().server}${api}`;
        const requestOptions = {
            ...opts,
            ...options,
            headers: {
                ...options.headers,
                ...opts.headers
            }
        };

        console.log(`Sending ${method.toUpperCase()} request to ${url} with options:`, requestOptions);

        request[method](url, requestOptions, (err, res) => {
            if (err) {
                console.error(`Error in ${method.toUpperCase()} request to ${url}:`, err);
                return reject(err);
            }

            try {
                const responseBody = JSON.parse(res.body);
                console.log(`Received response from ${url}:`, responseBody);
                resolve(responseBody);
            } catch (parseError) {
                console.error(`Error parsing response from ${url}:`, parseError);
                reject(parseError);
            }
        });
    });
};



// fields from the dummysite resource given by k8s API
const fieldsFromDummysite = (object) => {
    if (!object || !object.metadata) {
        console.error('Invalid object structure:', object)
        throw new Error('Invalid object structure')
    }
    const fields = {
        dummysite_name: object.metadata.name,
        container_name: object.metadata.name,
        deployment_name: `${object.metadata.name}-deployment`,
        namespace: object.metadata.namespace,
        website_url: object.spec.website_url,
    }
    console.log('Extracted fields:', fields)
    return fields;
}


// fields to inject into deployment template
const fieldsFromDeployment = (object) => {
    if (!object || !object.metadata) {
        console.error('Invalid object structure:', object)
        throw new Error('Invalid object structure')
    }

    const fields = {
        container_name: object.metadata.labels.name,
        dummysite_name: object.metadata.labels.name,
        deployment_name: `${object.metadata.labels.name}-deployment`,
        namespace: object.metadata.namespace,
        website_url: object.spec.website_url,
    }
    console.log('Extracted fields:', fields)
    return fields;
}


//
const getDeploymentYAML = async (fields) => {
    const deploymentTemplate = await fs.readFile("deployment.mustache", "utf-8")
    const renderedYAML = mustache.render(deploymentTemplate, fields)
    console.log('Rendered Deployment YAML:', renderedYAML);
    return renderedYAML;
}

// check if deployment already exists
const jobForDummysiteAlreadyExists = async (fields) => {
    const { dummysite_name, namespace } = fields
    const { items } = await sendRequestToApi(`/apis/apps/v1/namespaces/${namespace}/deployments`)
    if (!items) {
        console.error('Failed to retrieve deployments for namespace:', namespace);
        return false;
    }

    return items.find(item => item?.metadata?.labels?.dummysite === dummysite_name)
}

const createDeployment = async (fields) => {
    console.log('Scheduling new job number', fields.length, 'for dummysite', fields.dummysite_name, 'to namespace', fields.namespace)

    const yaml = await getDeploymentYAML(fields)

    return sendRequestToApi(`/apis/apps/v1/namespaces/${fields.namespace}/deployments`, 'post', {
        headers: {
            'Content-Type': 'application/yaml'
        },
        body: yaml
    })
}


const removeDeployment = async ({ namespace, deployment_name }) => {
    const pods = await sendRequestToApi(`/api/v1/namespaces/${namespace}/pods/`)

    if (!pods || !pods.items) {
        console.error('Failed to retrieve pods for namespace:', namespace);
        return;
    }

    pods.items.filter(pod => pod.metadata.labels && pod.metadata.labels['job-name'] === deployment_name).forEach(pod => removePod({ namespace, pod_name: pod.metadata.name }))

    return sendRequestToApi(`/apis/apps/v1/namespaces/${namespace}/deployments/${deployment_name}`, 'delete')
}

const removeDummysite = ({ namespace, dummysite_name }) => sendRequestToApi(`/apis/stable.dwk/v1/namespaces/${namespace}/dummysites/${dummysite_name}`, 'delete')

const removePod = ({ namespace, pod_name }) => sendRequestToApi(`/api/v1/namespaces/${namespace}/pods/${pod_name}`, 'delete')

const cleanupForDummysite = async ({ namespace, dummysite_name }) => {
    console.log('Doing cleanup')
    clearTimeout(timeouts[dummysite_name])

    const deployments = await sendRequestToApi(`/apis/apps/v1/namespaces/${namespace}/deployments`)

    if (!deployments || !deployments.items) {
        console.error('Failed to retrieve deployments for namespace:', namespace);
        return;
    }

    deployments.items.forEach(dep => {
        // ignores all other deployment objects
        if (!dep.metadata.labels || !dep.metadata.labels.dummysite === dummysite_name) return

        removeDeployment({ namespace, deployment_name: dep.metadata.name })
    })
}

const rescheduleDeployment = (deploymentObject) => {
    try {
        const fields = fieldsFromDeployment(deploymentObject)
        if (Number(fields.length) <= 1) {
            console.log('dummysite ended. Removing dummysite.')
            return removeDummysite(fields)
        }
    } catch (error) {
        console.error('Error rescheduling deployment:', error);
    }
}


const maintainStatus = async () => {
    (await client.listPodForAllNamespaces()).body // A bug in the client(?) was fixed by sending a request and not caring about response

    /**
     * Watch dummysites
     */

    const dummysite_stream = new JSONStream()

    dummysite_stream.on('data', async (data) => {
        console.log('Received dummysite event:', data);

        const { type, object } = data;
        if (!type || !object) {
            console.error('Invalid event structure:', data);
            return;
        }

        try {
            const fields = fieldsFromDummysite(object)

            if (type === 'ADDED') {
                if (await jobForDummysiteAlreadyExists(fields)) return // Restarting application would create new 0th deployments without this check
                await createDeployment(fields)
            }
            if (type === 'DELETED') cleanupForDummysite(fields)
        } catch (error) {
            console.error('Error processing dummysite event:', error);
        }
    })

    request.get(`${kc.getCurrentCluster().server}/apis/stable.dwk/v1/dummysites?watch=true`, opts).pipe(dummysite_stream)

    /**
     * Watch Deployments
     */

    const deployment_stream = new JSONStream()

    deployment_stream.on('data', async (data) => {
        console.log('Received deployment event:', data);

        const { type, object } = data;
        if (!type || !object) {
            console.error('Invalid event structure:', data);
            return;
        }

        try {
            if (!object.metadata || !object.metadata.labels || !object.metadata.labels.dummysite) return // If it's not dummysite job don't handle
            if (type === 'DELETED' || object.metadata.deletionTimestamp) return // Do not handle deleted deployments
            if (!object?.status?.succeeded) return

            rescheduleDeployment(object)
        } catch (error) {
            console.error('Error processing deployment event:', error);
        }
    })

    request.get(`${kc.getCurrentCluster().server}/apis/apps/v1/deployments?watch=true`, opts).pipe(deployment_stream)
}

maintainStatus()
