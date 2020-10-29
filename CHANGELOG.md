## v0.5.0 (Unreleased)

ADDITIONS

- cmd/webui: initial setup with pretty JSON formatting example
- file: Add support for writing EBCDIC
- file: Allow File to be written as DTSU with control bytes dictating line lengths rather than line breaks.
- imageViewData: add decode method for ImageData
- imageViewData: attempt base64 decode when generating a file

BUG FIXES

- file: Do not overwrite institution sequence number if one already exists
- file: populate recordType inside each record's JSON unmarshal
- file: setup additional nil checks

IMPROVEMENTS

- api: add returnLocationRoutingNumber to BundleHeader
- api: include missing imageViewDataSize
- api: use a correct example timestamp

BUILD

- chore(deps): update golang docker tag to v1.15
- chore(deps): update module gorilla/mux to v1.8.0
- chore(deps): update module prometheus/client_golang to v1.7.1

## v0.4.3 (Released 2020-07-07)

BUILD

- build: add OpenShift [`quay.io/moov/imagecashletter`](https://quay.io/repository/moov/imagecashletter) Docker image
- build: convert to Actions from TravisCI
- chore(deps): update module prometheus/client_golang to v1.6.0
- chore(deps): upgrade github.com/gorilla/websocket to v1.4.2

## v0.4.2 (Released 2020-04-14)

BUILD

- build: fix windows install of 'make'

## v0.4.1 (Released 2020-04-14)

IMPROVEMENTS

- api: use shorter summaries

BUILD

- build: upgrade to Go 1.14.x
- build: upgrade staticcheck to 2020.1.3

## v0.4.0 (Released 2020-04-14)

ADDITIONS

- server: add version handler to admin HTTP server

IMPROVEMENTS

- icl: log crasher file after it's parsed
- api,client: rename models whose names are duplicated across projects
- api,client: use shared Error model

BUILD

- Update module prometheus/client_golang to v1.2.1
- build: run sonatype-nexus-community/nancy in CI

## v0.3.0 (Released 2019-10-18)

ADDITIONS

- file: add FileFromJSON to decode ICL files
- cmd/server: decode a file as JSON or plain text

BUG FIXES

- reader: setup a File internally before reading

BUILD

- build: upgrade to Go 1.13 and Debian 10
- build: update openapi-generator to v4.1.3

## v0.2.0 (Released 2019-08-20)

BREAKING CHANGE

In our OpenAPI we've renamed fields generated as `Id` to `ID`, which is more in-line with Go's style conventions.

ADDITIONS

- Add RuneCountInString check to Parse(record string) functions
- cmd/server: bind HTTP server with TLS if HTTPS_* variables are defined

BUG FIXES

- all: check record lengths before parsing them
- all: fix range checks w.r.t added crasher files

BUILD

- build: upgrade openapi-generator to 4.1.0
- cmd/server: update github.com/moov-io/base to v0.10.0

## v0.1.1 (Released 2019-06-25)

BUG FIXES

- all: fixup panics found from first hour of fuzzing

IMPROVEMENTS

- build: push moov/imagecashletter:latest on 'make release-push'

## v0.1.0 (Released 2019-06-19)

- Initial release
