# Overview
<!-- Place this tag where you want the button to render. -->
<a class="github-button" href="https://github.com/moov-io/imagecashletter" data-size="large" data-show-count="true" aria-label="Star moov-io/imagecashletter on GitHub">moov-io/imagecashletter</a>
<a href="https://godoc.org/github.com/moov-io/imagecashletter"><img src="https://godoc.org/github.com/moov-io/imagecashletter?status.svg" /></a>
<a href="https://raw.githubusercontent.com/moov-io/imagecashletter/master/LICENSE"><img src="https://img.shields.io/badge/license-Apache2-blue.svg" /></a>

Moov ImageCashLetter implements a low level Image Cash Letter (ICL) interface for parsing, creating, and validating, ICL files. Moov ImageCashLetter exposes an HTTP API for REST based interaction. Any language which can use HTTP and JSON can leverage the ImageCashLetter Server. The API's endpoints expose both text and JSON to easily ingest or export either format.

## Running Moov ImageCashLetter Server

Moov ImageCashLetter can be deployed in multiple scenarios.

- <a href="#binary-distribution">Binary Distributions</a> are released with every versioned release. Frequently added to the VM/AMI build script for the application needing Moov ICL.
- A <a href="#docker-container">Docker container</a> is built and added to Docker Hub with every versioned released.
- Our hosted [api.moov.io](https://api.moov.io) is updated with every versioned release. Our Kubernetes example is what Moov utilizes in our production environment.

### Binary Distribution

Download the [latest Moov ImageCashLetter server release](https://github.com/moov-io/imagecashletter/releases) for your operating system and run it from a terminal.

```sh
$ ./imagecashletter-darwin-amd64
ts=2019-06-20T23:23:44.870717Z caller=main.go:75 startup="Starting imagecashletter server version v0.2.0"
ts=2019-06-20T23:23:44.871623Z caller=main.go:135 transport=HTTP addr=:8083
ts=2019-06-20T23:23:44.871692Z caller=main.go:125 admin="listening on :9093"
```

Next [Connect to Moov imagecashletter](#connecting-to-moov-imagecashletter)

### Docker Container

Moov ImageCashLetter is dependent on Docker being properly installed and running on your machine. Ensure that Docker is running. If your Docker client has issues connecting to the service review the [Docker getting started guide](https://docs.docker.com/get-started/) if you have any issues.

```sh
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```

Execute the Docker run command

```sh
$ docker run -p 8083:8083 -p 9093:9093 moov/imagecashletter:latest
ts=2019-06-21T17:03:23.782592Z caller=main.go:69 startup="Starting imagecashletter server version v0.2.0"
ts=2019-06-21T17:03:23.78314Z caller=main.go:129 transport=HTTP addr=:8083
ts=2019-06-21T17:03:23.783252Z caller=main.go:119 admin="listening on :9093"
```

Next [Connect to Moov ImageCashLetter](#connecting-to-moov-imagecashletter)

### Kubernetes

The following snippet runs the ImageCashLetter Server on [Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/) in the `apps` namespace. You could reach the ImageCashLetter instance at the following URL from inside the cluster.

```
# Needs to be ran from inside the cluster
$ curl http://imagecashletter.apps.svc.cluster.local:8083/ping
PONG

$ curl http://localhost:8083/files
{"files":[],"error":null}
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
Next [Connect to Moov ImageCashLetter](#connecting-to-moov-imagecashletter)

## Connecting to Moov ImageCashLetter

The Moov ImageCashLetter service will be running on port `8083` (with an admin port on `9093`).

Confirm that the service is running by issuing the following command or simply visiting the url in your browser [localhost:8083/ping](http://localhost:8083/ping)

```bash
$ curl http://localhost:8083/ping
PONG

$ curl http://localhost:8083/files
null
```

### API documentation

See our [API documentation](https://moov-io.github.io/imagecashletter/api/) for Moov ImageCashLetter endpoints.

### ImageCashLetter Admin Port

The port `9093` is bound by ImageCashLetter for our admin service. This HTTP server has endpoints for Prometheus metrics (`GET /metrics`), readiness (`GET /ready`) and liveness checks (`GET /live`).

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](https://moov-io.github.io/imagecashletter/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#imagecashletter`) to have an interactive discussion about the development of the project.
