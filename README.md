[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/imagecashletter/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/imagecashletter/api/#overview">API Endpoints</a>
  ·
  <a href="https://moov.io/blog/education/imagecashletter-api-guide/">API Guide</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/imagecashletter?status.svg)](https://godoc.org/github.com/moov-io/imagecashletter)
[![Build Status](https://github.com/moov-io/imagecashletter/workflows/Go/badge.svg)](https://github.com/moov-io/imagecashletter/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/imagecashletter)](https://goreportcard.com/report/github.com/moov-io/imagecashletter)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/imagecashletter?label=project%20size)](https://github.com/moov-io/imagecashletter)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/imagecashletter/master/LICENSE)
[![Slack Channel](https://img.shields.io/badge/Slack-Moov-e01563)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/imagecashletter)](https://hub.docker.com/r/moov/imagecashletter)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/imagecashletter)](https://github.com/moov-io/imagecashletter)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/imagecashletter

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

ImageCashLetter implements a reader, writer, and validator for X9’s Specifications for [Image Cash Letter](https://en.wikipedia.org/wiki/Check_21_Act) (ICL) to provide Check 21 services in an HTTP server and Go library. The HTTP server is available in a [Docker image](#docker) and the Go package `github.com/moov-io/imagecashletter` is available.

## Table of contents

- [Project status](#project-status)
- [Usage](#usage)
  - As an API
    - [Docker](#docker) ([Config](#configuration-settings))
    - [Google Cloud](#google-cloud-run) ([Config](#configuration-settings))
    - [Data persistence](#data-persistence)
  - [As a Go module](#go-library)
  - [Command-line tools](#command-line-tools)
  - [As an in-browser parser](#in-browser-icl-file-parser)
- [Learn about Image Cash Letter](#learn-about-image-cash-letter)
- [Getting help](#getting-help)
- [Supported and tested platforms](#supported-and-tested-platforms)
- [Contributing](#contributing)
- [Related projects](#related-projects)

## Project status

Moov ImageCashLetter is actively used in multiple production environments. Please star the project if you are interested in its progress. If you have layers above ImageCashLetter to simplify tasks, perform business operations, or found bugs we would appreciate an issue or pull request. Thanks!

## Usage
The Image Cash Letter project implements an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/imagecashletter) for creating and modifying ICL files. We also have some [examples](https://pkg.go.dev/github.com/moov-io/imagecashletter/examples) of the reader and writer.

### Docker

We publish a [public Docker image `moov/imagecashletter`](https://hub.docker.com/r/moov/imagecashletter/) from Docker Hub or use this repository. No configuration is required to serve on `:8083` and metrics at `:9093/metrics` in Prometheus format. We also have Docker images for [OpenShift](https://quay.io/repository/moov/imagecashletter?tab=tags) published as `quay.io/moov/imagecashletter`.

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

### Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-icl-demo
```

Enable the [Artifact Registry](https://cloud.google.com/artifact-registry/docs/docker/quickstart) API for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following commands, substituting your unique project ID:

```
gcloud services enable artifactregistry.googleapis.com
gcloud artifacts repositories create imagecashletter \
  --repository-format=docker --location=us
gcloud auth configure-docker us-docker.pkg.dev

docker pull moov/imagecashletter
docker tag moov/imagecashletter us-docker.pkg.dev/<PROJECT-ID>/imagecashletter/imagecashletter
docker push us-docker.pkg.dev/<PROJECT-ID>/imagecashletter/imagecashletter
```

Deploy the container to Cloud Run:
```
gcloud run deploy --image us-docker.pkg.dev/<PROJECT-ID>/imagecashletter/imagecashletter --port 8083
```

Select your target platform to `1`, service name to `imagecashletter`, and region to the one closest to you (enable Google API service if a prompt appears). Upon a successful build you will be given a URL where the API has been deployed:

```
https://YOUR-ICL-APP-URL.a.run.app
```

Now you can list files stored in-memory:
```
curl https://YOUR-ICL-APP-URL.a.run.app/files
```
You should get this response:
```
null
```


### Configuration settings

The following environmental variables can be set to configure behavior in ImageCashLetter.

| Environmental Variable | Description | Default |
|------------------------|-------------|---------|
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE` | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |
| `MAX_UPLOAD_SIZE` | Maximum size (in bytes) of HTTP request bodies accepted when creating files via the v2 API. Applies to both JSON and multipart/form-data uploads. | `104857600` (100MB) |
| `READER_BUFFER_SIZE` | Size (in bytes) of the buffer used when reading raw ICL file uploads. | `bufio.MaxScanTokenSize` (64KB) |
| `SKIP_ALL_ON_FILE_CREATE` | If true, the server will use `ValidateOpts.SkipAll` as a base for all file creates (merged with any per-request opts like `?skipAll=...`). Useful for archived/non-compliant data. | false |
| `SKIP_COUNT_VALIDATION_ON_FILE_CREATE` | If true, the server will use `ValidateOpts.SkipCountValidation` as a base for all file creates (merged with per-request). | false |
| `SKIP_INVALID_CONTACT_PHONE_NUMBERS_ON_FILE_CREATE` | If true, the server will use `ValidateOpts.SkipInvalidContactPhoneNumbers` as a base for all file creates (merged with per-request). Accepts non-numeric `FileControl.ImmediateOriginContactPhoneNumber` / `CashLetterHeader.OriginatorContactPhoneNumber` values as-is. These fields are fixed-width (10 chars); a human-formatted phone number (e.g. `(831) 555-1234`) will already have been truncated by the file's writer before this library sees it, so enabling this does not recover lost digits -- it only stops rejecting the file. | false |
| `FRB_COMPATIBILITY_MODE` | If set to `true`, enables Federal Reserve Bank (FRB) compatibility mode, which relaxes some record validations to accept files produced in FRB format. | Empty (disabled) |

### Data persistence
By design, ImageCashLetter  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart ImageCashLetter will have no files or data saved. Also, no in-memory encryption of the data is performed.

### Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.26 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git clone git@github.com:moov-io/imagecashletter.git
$ cd imagecashletter

# Pull down into the Go Module cache (omit -u so Go does not bump the module's go directive)
$ go get github.com/moov-io/imagecashletter

$ go doc github.com/moov-io/imagecashletter CashLetter
```

The package [`github.com/moov-io/imagecashletter`](https://pkg.go.dev/github.com/moov-io/imagecashletter) offers a Go-based Image Cash Letter file reader and writer. To get started, check out a specific example. Run the examples from the repository root (`go run ./examples/...`); paths inside the programs are relative to the root, so `cd examples/imagecashletter-read && go run .` will fail. The write example writes to a temp file and prints the path so it does not overwrite the tracked sample.

| ICL File | Read | Write |
|---------|------|-------|
| [Link](examples/imagecashletter-read/iclFile.x937) | [Link](examples/imagecashletter-read/main.go) | [Link](examples/imagecashletter-write/main.go) |

```
$ go run ./examples/imagecashletter-read
$ go run ./examples/imagecashletter-write
```

ImageCashLetter's file handling behaviors can be modified to accommodate your specific use case. This is done by passing _options_ into ICL's `reader` and `writer` during instantiation. For example, to read EBCDIC encoded files you would instantiate a reader with `NewReader(fd, ReadVariableLineLengthOption(), ReadEbcdicEncodingOption())`.

The following options are currently supported:
| Option | Description |
|-----|-----|
| `ReadVariableLineLengthOption` | Allows Reader to split ICL files based on the Inserted Length Field. |
| `ReadEbcdicEncodingOption` | Allows Reader to decode scanned lines from EBCDIC to UTF-8. |
| `ReadValidateOpts` | Allows skipping validation checks for archived or non-compliant ICL files via ValidateOpts (e.g. SkipAll). Use `file.SetValidation(opts)` after read if needed for later Validate/Create calls. |
| `WriteVariableLineLengthOption` | Instructs the Writer to begin each record with the appropriate Inserted Length Field. |
| `WriteEbcdicEncodingOption` | Allows Writer to write file in EBCDIC. |



### Command-line tools

`cmd/readImageCashLetter` prints summary counts (or JSON with `-json`) from an ICL file. The default sample at `cmd/readImageCashLetter/BNK20181015-A.icl` is fixed-width ASCII and can be read with no extra flags. Most files under `test/testdata/` use 4-byte inserted-length prefixes and need `-variable-length`; EBCDIC files also need `-ebcdic`. Without those flags, variable-length or EBCDIC files fail with `recordType is an unknown record type`.

```
# Fixed-width ASCII (the default sample in that directory)
$ go run ./cmd/readImageCashLetter -fPath=cmd/readImageCashLetter/BNK20181015-A.icl

# Variable-length ASCII
$ go run ./cmd/readImageCashLetter -variable-length -fPath=test/testdata/valid-ascii.x937

# Variable-length EBCDIC
$ go run ./cmd/readImageCashLetter -variable-length -ebcdic -fPath=test/testdata/valid-ebcdic.x937
```

`cmd/writeImageCashLetter` generates a large ICL file for benchmarks. `cmd/server` is the HTTP API described above.

### In-browser ICL file parser
Using our [in-browser utility](https://oss.moov.io/x9/), you can instantly convert X9 files into JSON. Either paste in ICL file content directly or choose a file from your local machine. This tool is particularly useful if you're handling sensitive PII or want to perform some quick tests, as operations are fully client-side with nothing stored in memory. We plan to support bidirectional conversion in the near future.

## Learn about Image Cash Letter
- [Intro to ICL](./docs/intro.md)
- [ICL File Structure](./docs/file-structure.md)

## Getting help

 channel | info
 ------- | -------
[Project Documentation](https://moov-io.github.io/imagecashletter/) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/imagecashletter/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Raspberry Pi

Note: 32-bit platforms have known issues and are not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.26 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

### Releasing

To make a release of imagecashletter simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

### Testing

We maintain a comprehensive suite of unit tests and recommend table-driven testing when a particular function warrants several very similar test cases. To run all test files in the current directory, use `go test`.

### Fuzzing

We currently run fuzzing over Image Cash Letter files in the form of a [Github Action](https://github.com/moov-io/imagecashletter/actions/workflows/fuzz.yml). Please report crashes examples to [`oss@moov.io`](mailto:oss@moov.io). Thanks!

## Related projects
As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Watchman](https://github.com/moov-io/watchman) offers search functions over numerous trade sanction lists from the United States and European Union.

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Wire](https://github.com/moov-io/wire) implements an interface to write files for the Fedwire Funds Service, a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Metro 2](https://github.com/moov-io/metro2) provides a way to easily read, create, and validate Metro 2 format, which is used for consumer credit history reporting by the United States credit bureaus.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
