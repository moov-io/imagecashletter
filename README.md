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
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
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

### Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-icl-demo
```

Enable the [Container Registry](https://cloud.google.com/container-registry) API for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following Docker commands, substituting your unique project ID:

```
docker pull moov/imagecashletter
docker tag moov/imagecashletter gcr.io/<PROJECT-ID>/imagecashletter
docker push gcr.io/<PROJECT-ID>/imagecashletter
```

Deploy the container to Cloud Run:
```
gcloud run deploy --image gcr.io/<PROJECT-ID>/imagecashletter --port 8083
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
|-----|-----|-----|
| `READER_BUFFER_SIZE` | Size of buffer to use with `bufio.Scanner`. | Check `bufio.MaxScanTokenSize` |
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE`  | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |
| `FRB_COMPATIBILITY_MODE`  | If set, enables Federal Reserve Bank (FRB) compatibility mode. | Empty |

### Data persistence
By design, ImageCashLetter  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart ImageCashLetter will have no files or data saved. Also, no in-memory encryption of the data is performed.

### Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.18 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/imagecashletter.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/imagecashletter

$ go doc github.com/moov-io/imagecashletter CashLetter
```

The package [`github.com/moov-io/imagecashletter`](https://pkg.go.dev/github.com/moov-io/imagecashletter) offers a Go-based Image Cash Letter file reader and writer. To get started, check out a specific example:

 ICL File | Read | Write |
|---------|------|-------|
| [Link](examples/imagecashletter-read/iclFile.x937) | [Link](examples/imagecashletter-read/main.go) | [Link](examples/imagecashletter-write/main.go) |

ImageCashLetter's file handling behaviors can be modified to accommodate your specific use case. This is done by passing _options_ into ICL's `reader` and `writer` during instantiation. For example, to read EBCDID encoded files you would instantiate a reader with `NewReader(fd, ReadVariableLineLengthOption(), ReadEbcdicEncodingOption())`.

The following options are currently supported:
| Option | Description |
|-----|-----|
| `ReadVariableLineLengthOption` | Allows Reader to split ICL files based on the Inserted Length Field. |
| `ReadEbcdicEncodingOption` | Allows Reader to decode scanned lines from EBCDIC to UTF-8. |
| `WriteVariableLineLengthOption` | Instructs the Writer to begin each record with the appropriate Inserted Length Field. |
| `WriteEbcdicEncodingOption` | Allows Writer to write file in EBCDIC. |


### In-browser ICL file parser
Using our [in-browser utility](http://oss.moov.io/x9/), you can instantly convert X9 files into JSON. Either paste in ICL file content directly or choose a file from your local machine. This tool is particularly useful if you're handling sensitive PII or want to perform some quick tests, as operations are fully client-side with nothing stored in memory. We plan to support bidirectional conversion in the near future.

## Learn about Image Cash Letter
- [Intro to ICL](./docs/intro.md)
- [ICL File Structure](./docs/file-structure.md)

## Getting help

 channel | info
 ------- | -------
[Project Documentation](https://moov-io.github.io/imagecashletter/) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](github.com/moov-io/imagecashletter/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Raspberry Pi

Note: 32-bit platforms have known issues and are not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.18 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

### Releasing

To make a release of imagecashletter simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

### Testing

We maintain a comprehensive suite of unit tests and recommend table-driven testing when a particular function warrants several very similar test cases. To run all test files in the current directory, use `go test`.

### Fuzzing

We currently run fuzzing over ACH in the form of a [Github Action](https://github.com/moov-io/imagecashletter/actions/workflows/fuzz.yml). Please report crashes examples to [`oss@moov.io`](mailto:oss@moov.io). Thanks!

## Related projects
As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Watchman](https://github.com/moov-io/watchman) offers search functions over numerous trade sanction lists from the United States and European Union.

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Wire](https://github.com/moov-io/wire) implements an interface to write files for the Fedwire Funds Service, a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Metro 2](https://github.com/moov-io/metro2) provides a way to easily read, create, and validate Metro 2 format, which is used for consumer credit history reporting by the United States credit bureaus.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
