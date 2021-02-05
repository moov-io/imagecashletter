---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Docker

We publish a [public Docker image `moov/imagecashletter`](https://hub.docker.com/r/moov/imagecashletter/) from Docker Hub or use this repository. No configuration is required to serve on `:8083` and metrics at `:9093/metrics` in Prometheus format. We also have Docker images for [OpenShift](https://quay.io/repository/moov/imagecashletter?tab=tags) published as `quay.io/moov/imagecashletter`.

Moov ImageCashLetter is dependent on Docker being properly installed and running on your machine. Ensure that Docker is running. If your Docker client has issues connecting to the service, review the [Docker getting started guide](https://docs.docker.com/get-started/).

```
docker ps
```
```
CONTAINER ID        IMAGE        COMMAND        CREATED        STATUS        PORTS        NAMES
```

Pull & start the Docker image:
```
docker pull moov/imagecashletter:latest
docker run -p 8083:8083 -p 9093:9093 moov/imagecashletter:latest
```

List files stored in-memory:
```
curl localhost:8083/files
```
```
null
```

Upload an x9 file (binary):
```
curl -X POST --data-binary "@./test/testdata/valid-ascii.x937" http://localhost:8083/files/create
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"","standardLevel":"03","testIndicator":"T","immediateDestination":"061000146","immediateOrigin":"026073150", ...
```

Retrieve an existing x9 file (JSON):
```
curl http://localhost:8083/files/<YOUR-UNIQUE-FILE-ID>
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"","standardLevel":"03","testIndicator":"T","immediateDestination":"061000146","immediateOrigin":"026073150", ...
```

Create an x9 file from JSON:
```
curl -X POST -H "content-type: application/json" localhost:8083/files/create --data @./test/testdata/icl-valid.json
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"","standardLevel":"35","testIndicator":"T","immediateDestination":"231380104","immediateOrigin":"121042882", ...
```

Get the formatted file:
```
curl localhost:8083/files/<YOUR-UNIQUE-FILE-ID>/contents
```
```
P0135T231380104121042882201810032219NCitadel      Wells Fargo    US   P100123138010412104288220181003201810032219IGA1   Contact Name 5558675552  P200123138010412104288220181003201810039999   1  01             P25   123456789 031300012       555888100001000001       GD1Y030BP261121042882201810031       938383      01  Test Payee   Y10
...
```