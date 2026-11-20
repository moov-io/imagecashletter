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

The same HTTP API is served by `go run ./cmd/server` (listens on `:8083`, admin on `:9093`).

List files stored in-memory:
```
curl localhost:8083/files
```
```
null
```

Upload an x9 file (v1, EBCDIC):

`POST /files/create` accepts JSON (`Content-Type: application/json`) or a raw ICL body. Raw bodies are always read as variable-length EBCDIC (`ReadVariableLineLengthOption` and `ReadEbcdicEncodingOption` are always set). Use `valid-ebcdic.x937`; an ASCII file such as `valid-ascii.x937` will fail with a `FileCreationDate` mandatory-field error.

```
curl -X POST --data-binary "@./test/testdata/valid-ebcdic.x937" http://localhost:8083/files/create
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"","standardLevel":"03","testIndicator":"T","immediateDestination":"061000146","immediateOrigin":"026073150", ...
```

Upload an x9 file (v2, ASCII or EBCDIC):

`POST /v2/files` accepts JSON or a multipart form upload. For multipart, send the file in a `file` part. Set that part's `Content-Type` to `text/plain` for ASCII, or `application/octet-stream` (the default) for EBCDIC. A successful create returns HTTP 201 and a `Location` header of `/v2/files/{id}`. Files created through v2 are stored in the same in-memory repository as v1, so they can also be fetched with the v1 routes below.

```
curl -F "file=@./test/testdata/valid-ascii.x937;type=text/plain" http://localhost:8083/v2/files
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"","standardLevel":"03","testIndicator":"T","immediateDestination":"061000146","immediateOrigin":"026073150", ...
```

```
curl -F "file=@./test/testdata/valid-ebcdic.x937;type=application/octet-stream" http://localhost:8083/v2/files
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

`GET /files/{id}/contents` always writes variable-length EBCDIC (`WriteVariableLineLengthOption` and `WriteEbcdicEncodingOption`). The `Content-Type` is `text/plain`, but the body is not ASCII: each record starts with a 4-byte big-endian length prefix (for example `00 00 00 50` for an 80-byte record) followed by EBCDIC data.

```
curl --output letter.x937 http://localhost:8083/files/<YOUR-UNIQUE-FILE-ID>/contents
od -An -tx1 -N 32 letter.x937
```
```
 00 00 00 50 f0 f1 f0 f3 e3 f0 f6 f1 f0 f0 f0 f1
 f4 f6 f0 f2 f6 f0 f7 f3 f1 f5 f0 f2 f0 f2 f0 f1
```

To get a readable ASCII file, create through v2 and set `Accept: text/plain`. The response is still prefixed with 4-byte record lengths (the leading NUL bytes may not display in a terminal). `Accept: application/octet-stream` returns the same file as variable-length EBCDIC.

```
curl -F "file=@./test/testdata/valid-ascii.x937;type=text/plain" \
  -H "Accept: text/plain" \
  --output letter-ascii.x937 \
  http://localhost:8083/v2/files
```
```
\x00\x00\x00P0103T061000146026073150202010230911NWave Money        Wave Money        1   ...
```
