# Knative 

## Ping Pong App

### Installation

To install to cluster, just apply the manifests with kubectl.



### How to run

#### Command line

The values may vary depending on the IP Kubernetes assigns to the service and also the DNS setup configured for Knative. Assuming the cluster is exposed on localhost:8081 and the host value is as shown below (`kubectl get ksvc` will retrieve the correct value),
you can use `curl` like so:  

```bash
curl -H "Host: ping-pong-knative.default.172.18.0.3.sslip.io" http://localhost:8081
```

#### Browser

Using a browser extension that allows setting custom headers, set the host header to the address shown in `kubectl get ksvc` and navigate to the cluster's exposed address -- in my case, `localhost:8081`  
