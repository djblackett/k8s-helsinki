const mustache = require('mustache')
const fs = require('fs').promises


const getDeploymentYAML = async (fields) => {
    const deploymentTemplate = await fs.readFile("deployment.mustache", "utf-8")
    const renderedYAML = mustache.render(deploymentTemplate, fields)
    console.log('Rendered Deployment YAML:', renderedYAML);
    // return renderedYAML;

    await fs.writeFile("generated-deployment.yaml", renderedYAML)
}

const fields = {
    container_name: "dep-container",
    dummysite_name: "dummysite",
    deployment_name: `dummysite-deployment`,
    namespace: "default",
    website_url: "www.example.com"
}

getDeploymentYAML(fields)
