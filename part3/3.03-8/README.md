# Comparison of DBaaS and Persistent Volumes on Kubernetes

## DBaas Pros:
- easier to manage (update, scale, disaster management)
- easier to scale
- often better for SQL/relational DBs (difficult to shard)
- Less expertise required


## Kubernetes PV and Statefulset Pros:
- cheaper
- more fine grained control
- everything on one platform (K8s)
- works better with data compliance, data locality restrictions, or multi-region architecture needs




### 1 year cost comparison  
CloudSQL - $257, 593.68  
GKE -      $126, 712.80

As an aside, Koyeb estimates that a fully in-house implemented Kubernetes cluster may cost nearly $570 000 if you pay 4 devops engineers the average yearly salary in the USA. Only $5000 of that estimate is for the hardware to host the control plane and worker nodes. 




## Analysis:
At the end of the day, I think the choice between DBaas and a self-managed Kubernetes solution comes down to the size of the company and the expertise of its devops team. A smaller company may not have the technical expertise required to use a Kubernetes solution, and the cost of hiring somone with this expertise may nullify any monetary savings. However, larger companies will naturally have more staff and may find it easier to justify the expense of in-house cloud experts to manage Kubernetes databases. 

For a startup that is just getting going and trying to gain momentum, the simplicity of DBaas is a boon. There is always the option to transition to a Kubernetes solution later on when the company and its product is more stable.

For a well estabished company that already has a large staff of cloud experts, managing their own Kubernetes database implementations will be significantly easier. Larger companies will often require more database resources, so they also stand to save the most money by managing databases themselves. 


Sources: 

https://www.postgresql.eu/events/pgdayparis2022/sessions/session/3612/slides/294/PostgreSQL%20in%20the%20Cloud_%20DBaaS%20vs%20Kubernetes.pdf

https://www.percona.com/blog/should-you-deploy-your-databases-on-kubernetes-and-what-makes-statefulset-worthwhile/

https://www.koyeb.com/blog/the-true-cost-of-kubernetes-people-time-and-productivity
