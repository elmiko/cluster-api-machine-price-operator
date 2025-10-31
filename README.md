# cluster-api-machine-price-operator

This is an experiment to demonstrate how pricing information could be exposed
through Cluster API resources.

It is not meant for production use cases.

It is meant to demonstrate how the Cluster Autoscaler and Karpenter projects
could use pricing information during their cluster scaling operations.

adds this annotations to MachineDeployments
```
cluster.x-k8s.io/machine-current-price: "0.0"
```

the value will depend on the cloud provider implementation.
