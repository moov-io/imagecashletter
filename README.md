moov-io/imagecashletter
===
[![GoDoc](https://godoc.org/github.com/moov-io/imagecashletter?status.svg)](https://godoc.org/github.com/moov-io/imagecashletter)
[![Build Status](https://github.com/moov-io/imagecashletter/workflows/Go/badge.svg)](https://github.com/moov-io/imagecashletter/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/imagecashletter/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/imagecashletter)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/imagecashletter)](https://goreportcard.com/report/github.com/moov-io/imagecashletter)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/imagecashletter/master/LICENSE)

ImageCashLetter implements a reader, writer, and validator for X9’s Specifications for [Image Cash Letter](https://en.wikipedia.org/wiki/Check_21_Act) (ICL) to provide Check 21 services in an HTTP server and Go library. The HTTP server is available in a [Docker image](#docker) and the Go package `github.com/moov-io/imagecashletter` is available.

Docs: [Project](https://moov-io.github.io/imagecashletter/) | [API Endpoints](https://moov-io.github.io/imagecashletter/api/) | [In-Browser Parser](https://demo.moov.io/x9/)

## Usage

### Docker

We publish a [public Docker image `moov/imagecashletter`](https://hub.docker.com/r/moov/imagecashletter/) from Docker Hub or use this repository. No configuration is required to serve on `:8083` and metrics at `:9093/metrics` in Prometheus format. We also have docker images for [OpenShift](https://quay.io/repository/moov/imagecashletter?tab=tags).

Start the Docker image:
```
docker run -p 8083:8083 -p 9093:9093 moov/imagecashletter:latest
```

List files stored in-memory
```
curl localhost:8083/files
```
```
{"files":[],"error":null}
```

Create a file on the HTTP server
```
curl -XPOST --data-binary "@./test/testdata/BNK20180905121042882-A.icl" http://localhost:8083/files/create
```
```
{"id":"71ae3f5bc5527cdb1efc88e1814333fd9d6d2edb","fileHeader":{"id":"","standardLevel":"35","testIndicator":"T","immediateDestination":"231380104","immediateOrigin":"121042882", ...
```

Create a file with the JSON format on the HTTP server
```
curl -XPOST -H "content-type: application/json" localhost:8083/files/create --data @./test/testdata/icl-valid.json
```
```
{"id":"8afcde4fc2cf4023e92a7a96be9dbbe44e1e9508","fileHeader":{"id":"","standardLevel":"35","testIndicator":"T","immediateDestination":"231380104","immediateOrigin":"121042882", ...
```

Get the formatted file
```
curl localhost:8083/files/8afcde4fc2cf4023e92a7a96be9dbbe44e1e9508/contents
```
```
0135T231380104121042882201810032219NCitadel           Wells Fargo        US
0123138010412104288220181003201810032219IGA1      Contact Name  5558675552
0123138010412104288220181003201810039999      1   01
      123456789 031300012             555888100001000001              GD1Y030B
1121042882201810031              938383            01   Test Payee     Y10
...
```

### Go library

`github.com/moov-io/imagecashletter` offers a Go based Image Cash Letter file reader and writer. To get started checkout a specific example:

<details>
<summary>ICL File</summary>

 Example | Read | Write |
|---------|------|-------|
| [Link](examples/imagecashletter-read/iclFile.txt) | [Link](examples/imagecashletter-read/main.go) | [Link](examples/imagecashletter-write/main.go) |
</details>

### From Source

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/imagecashletter.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/imagecashletter

$ go doc github.com/moov-io/imagecashletter CashLetter
```

### Configuration

The following environmental variables can be set to configure behavior in paygate.

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE`  | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |

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
