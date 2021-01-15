[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/imagecashletter/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/imagecashletter/api/#overview">API Endpoints</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/imagecashletter?status.svg)](https://godoc.org/github.com/moov-io/imagecashletter)
[![Build Status](https://github.com/moov-io/imagecashletter/workflows/Go/badge.svg)](https://github.com/moov-io/imagecashletter/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/imagecashletter/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/imagecashletter)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/imagecashletter)](https://goreportcard.com/report/github.com/moov-io/imagecashletter)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/imagecashletter?label=project%20size)](https://github.com/moov-io/imagecashletter)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/imagecashletter)](https://hub.docker.com/r/moov/imagecashletter)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/imagecashletter)](https://github.com/moov-io/imagecashletter)
[![Twitter](https://img.shields.io/twitter/follow/moov_io?style=social)](https://twitter.com/moov_io?lang=en)

# moov-io/imagecashletter

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease-of-use.

ImageCashLetter implements a reader, writer, and validator for X9’s Specifications for [Image Cash Letter](https://en.wikipedia.org/wiki/Check_21_Act) (ICL) to provide Check 21 services in an HTTP server and Go library. The HTTP server is available in a [Docker image](#docker) and the Go package `github.com/moov-io/imagecashletter` is available.

## Project Status

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

Create a file on the HTTP server:
```
curl -XPOST --data-binary "@./test/testdata/valid-ascii.x937" http://localhost:8083/files/create
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"","standardLevel":"03","testIndicator":"T","immediateDestination":"061000146","immediateOrigin":"026073150", ...
```

Read the X9 file (in JSON form):
```
curl http://localhost:8083/files/<YOUR-UNIQUE-FILE-ID>
```
```
{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"","standardLevel":"03","testIndicator":"T","immediateDestination":"061000146","immediateOrigin":"026073150", ...
```

Create a file with JSON format on the HTTP server:
```
curl -XPOST -H "content-type: application/json" localhost:8083/files/create --data @./test/testdata/icl-valid.json
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

### Configuration Settings

The following environmental variables can be set to configure behavior in ImageCashLetter.

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE`  | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |

### Data Persistence
By design, ImageCashLetter  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart ImageCashLetter will have no files or data saved. Also, no in-memory encryption of the data is performed.

### Go Library

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/imagecashletter.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/imagecashletter

$ go doc github.com/moov-io/imagecashletter CashLetter
```

The package [`github.com/moov-io/imagecashletter`](https://pkg.go.dev/github.com/moov-io/imagecashletter) offers a Go-based Image Cash Letter file reader and writer. To get started, check out a specific example:

<details>
<summary>ICL File</summary>

 Example | Read | Write |
|---------|------|-------|
| [Link](examples/imagecashletter-read/iclFile.x937) | [Link](examples/imagecashletter-read/main.go) | [Link](examples/imagecashletter-write/main.go) |
</details>

### In-Browser ICL File Parser
Using our [in-browser utility](http://oss.moov.io/x9/), you can instantly convert X9 files into JSON. Either paste in ICL file content directly or choose a file from your local machine. This tool is particulary useful if you're handling sensitive PII or want perform some quick tests, as operations are fully client-side with nothing stored in memory. We plan to support bidirectional conversion in the near future.

## Getting Help

 channel | info
 ------- | -------
[Project Documentation](https://moov-io.github.io/imagecashletter/) | Our project documentation available online.
Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Rasberry Pi

Note: 32-bit platforms have known issues and is not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

### Fuzzing

We currently run fuzzing over ACH in the form of a [`moov/imagecashletterfuzz`](https://hub.docker.com/r/moov/imagecashletterfuzz) Docker image. You can [read more](./test/fuzz-reader/README.md) or run the image and report crasher examples to [`security@moov.io`](mailto:security@moov.io). Thanks!

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
