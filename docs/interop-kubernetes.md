---
title: Kubernetes
sidebar_label: Kubernetes
---

## Examples

Deploy opctl in kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: opctl-in-kubernetes
spec:
  replicas: 1
  template:
    spec:
      containers:
        - name: opctl
          image: opctl:0.1.24
          ports:
            # expose to other containers
            - name: http
              containerPort: 42224
              protocol: TCP
          securityContext:
            privileged: true
```