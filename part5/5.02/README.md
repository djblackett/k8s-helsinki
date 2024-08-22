# Exercise 5.02

Injecting Linkerd into ArgoRollout files was causing problems so I changed them back to standard Kubernetes deployment objects. NATS doesn't behave with linkerd so I also added the `skip-outbound-ports: 4222` rule to the `nats-config.yaml`.
