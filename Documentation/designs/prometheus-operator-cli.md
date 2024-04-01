# Prometheus Operator CLI

## Summary

This document proposes a new CLI tool to manage Prometheus Operator resources, allowing users to create, update, and delete Prometheus Operator resources using a simple CLI interface as well as
troubleshoot and debug Prometheus Operator resources.

## Background

During KubeCon Paris 2024, we received feedback from users that managing Prometheus Operator resources is difficult and error-prone. Users reported that they often encounter issues when creating, updating, and deleting Prometheus Operator resources, and that troubleshooting and debugging Prometheus Operator resources is challenging, below you can find some of the issues that users reported:

### Why my Prometheus resource is not being created?

People were struggling to troubleshoot why their Prometheus resources were not being created. They were not sure if the issue was with the Prometheus Operator or with the Prometheus resource itself.
After long investigation, they found out that the Prometheus Operator were only watching the namespace where the Prometheus Operator was deployed, and the Prometheus resource was being created in a different namespace.

### Why my ServiceMonitor is not being created?

People were struggling to troubleshoot why their ServiceMonitor resources were not being created. They were not sure if the issue was with the Prometheus Operator or with the ServiceMonitor resource itself.

After long investigation, they found out that the matchingLabels field in the ServiceMonitor resource was not matching the labels of the target Service.

These are just some common issues that users reported, but there are many other issues that users face when managing Prometheus Operator resources, and we believe that a CLI tool can help users to manage Prometheus Operator resources more easily and efficiently.

## Proposal

We propose to create a new CLI tool allowing users to create, update, and delete Prometheus Operator resources using a simple CLI interface as well as troubleshoot and debug Prometheus Operator resources.

The main idea is to distribute this CLI as a [kubectl plugin](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) and the user experience should be similar to the `kubectl` CLI, as you can see in the example below:

```bash
kubectl prometheus-operator create prometheus --name my-prometheus --namespace my-namespace --replicas 2 --shards 2
```

### Where the code will live?

The prometheus-operator-cli code will be placed in a dedicated repository under the [Prometheus organization](https://github.com/prometheus-operator), this will allow us to have a separeted lifecycle for the Prometheus Operator and the Prometheus Operator CLI and the following benefits:

- We can have a dedicated team to maintain the Prometheus Operator CLI
- We can use the Prometheus Operator client as a customer and be more careful with the changes that we make in the Prometheus Operator codebase
- We can have a dedicated release cycle for the Prometheus Operator CLI

### General features

The Prometheus Operator CLI should have the following features, but not limited to:

#### Resource scaffolding

Allow users create Prometheus Operator resources using a simple CLI interface, the CLI should generate the resource manifest based on the user input.

This is especially useful for users that are not familiar with the Prometheus Operator resources, they can use the CLI to generate the resource manifest and then customize it according to their needs.

#### Linting and Validation

Allow users to validate Prometheus Operator resources before creating them, the CLI should check if the resource manifest is valid and if it follows the best practices.

This is especially useful for CI pipelines, where users can validate the Prometheus Operator resources before applying them to the cluster.

#### Troubleshooting and Debugging

Allow users to troubleshoot and debug Prometheus Operator resources, the CLI should provide useful information about the Prometheus Operator resources, such as if the resource is being watched by the Prometheus Operator, if the labels are matching, etc.