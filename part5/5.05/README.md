# Rancher VS OpenShift

## Common features

- web based interfaces for managing clusters
- major cloud providers have integrated workflows
- implement extra RBAC security practices and simplifies managing security
- CI/CD pipelines can be integrated
- provide support (costs vary between tiers)  

## Benefits of OpenShift

- all in one solution for enterprises - built in CI/CD, monitoring, service mesh, etc
- opinionated - reduces decision making complexity
- focuses on developer experience
- stronger emphasis on security and compliance control - good for finance, government, healthcare, etc

## Negatives of OpenShift

- requires a subscription, thus locking you into one ecosystem
- less flexibility for preexisting clusters - e.g. can only run CentOS
- cost can be high, esp. for smaller organizations

## Benefits of Rancher

- more flexible - can manage multiple clusters from different providers 
- Helm can be used for package management and installation
- open source product - paid support is entirely optional  
- can be removed without losing clusters - less vendor lock-in 
- generally cheaper than OpenShift

## Negatives of Rancher

- increased freedom comes at the cost of more complexity
- may require more employees with more expertise
- security requires extra configuration to match that of OpenShift


## Verdict

Aside from large enterprise organizations, Rancher is probably the better fit for most use cases. However, each one excels in different circumstances. 

### Sources  

<https://www.densify.com/openshift-tutorial/rancher-vs-openshift/>
<https://www.rancher.com/why-rancher>
<https://www.redhat.com/en/technologies/cloud-computing/openshift>
