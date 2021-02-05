---
layout: page
title: Kubernetes
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Kubernetes

The following snippet runs the ImageCashLetter Server on [Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/) in the `apps` namespace. You can reach the ImageCashLetter instance at the following URL from inside the cluster.

```
# Needs to be ran from inside the cluster
$ curl http://imagecashletter.apps.svc.cluster.local:8083/ping
PONG
```

Kubernetes manifest - save in a file (`imagecashletter.yaml`) and apply with `kubectl apply -f imagecashletter.yaml`.

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: apps
---
apiVersion: v1
kind: Service
metadata:
  name: imagecashletter
  namespace: apps
spec:
  type: ClusterIP
  selector:
    app: imagecashletter
  ports:
    - name: http
      protocol: TCP
      port: 8083
      targetPort: 8083
    - name: metrics
      protocol: TCP
      port: 9093
      targetPort: 9093
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: imagecashletter
  namespace: apps
  labels:
    app: imagecashletter
spec:
  replicas: 1
  selector:
    matchLabels:
      app: imagecashletter
  template:
    metadata:
      labels:
        app: imagecashletter
    spec:
      containers:
      - image: moov/imagecashletter:latest
        imagePullPolicy: Always
        name: imagecashletter
        args:
          - -http.addr=:8083
          - -admin.addr=:9093
        ports:
          - containerPort: 8083
            name: http
            protocol: TCP
          - containerPort: 9093
            name: metrics
            protocol: TCP
        resources:
          limits:
            cpu: 0.1
            memory: 50Mi
          requests:
            cpu: 25m
            memory: 10Mi
        readinessProbe:
          httpGet:
            path: /ping
            port: 8083
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          httpGet:
            path: /ping
            port: 8083
          initialDelaySeconds: 5
          periodSeconds: 10
      restartPolicy: Always
```