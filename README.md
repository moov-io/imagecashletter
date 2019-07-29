moov-io/imagecashletter
===
[![GoDoc](https://godoc.org/github.com/moov-io/imagecashletter?status.svg)](https://godoc.org/github.com/moov-io/imagecashletter)
[![Build Status](https://travis-ci.com/moov-io/imagecashletter.svg?branch=master)](https://travis-ci.com/moov-io/imagecashletter)
[![Coverage Status](https://codecov.io/gh/moov-io/imagecashletter/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/imagecashletter)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/imagecashletter)](https://goreportcard.com/report/github.com/moov-io/imagecashletter)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/imagecashletter/master/LICENSE)

X9’s Specifications for ICL (Image Cash Letter) to provide Check 21 services

Package `github.com/moov-io/imagecashletter` implements a file reader and writer for parsing [Image Cash Letter](https://en.wikipedia.org/wiki/Check_21_Act) (ICL) files.

Docs: [docs.moov.io](https://docs.moov.io/icl/) | [api docs](https://api.moov.io/apps/imagecashletter/)

## Project Status

Moov IMAGECASHLETTER is under active development and in production for multiple companies.  Please star the project if you are interested in its progress.

## Usage

### Go library

`github.com/moov-io/imagecashletter` offers a Go based Image Cash Letter file reader and writer. To get started checkout a specific example:

<details>
<summary>ICL File</summary>

 Example | Read | Write |
|---------|------|-------|
| [Link](examples/imagecashletter-read/iclFile.txt) | [Link](examples/imagecashletter-read/main.go) | [Link](examples/imagecashletter-write/main.go) |
</details>

### From Source

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and thus requires Go 1.11+. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases) as well. We highly recommend you use a tagged release for production.

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
[Project Documentation](http://docs.moov.io/) | Our project documentation available online.
Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request an invite to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Rasberry Pi

Note: 32-bit platforms have known issues and is not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

Note: This project uses Go Modules, which requires Go 1.11 or higher, but we ship the vendor directory in our repository.

### Fuzzing

We currently run fuzzing over ACH in the form of a [`moov/imagecashletterfuzz`](https://hub.docker.com/r/moov/imagecashletterfuzz) Docker image. You can [read more](./test/fuzz-reader/README.md) or run the image and report crasher examples to [`security@moov.io`](mailto:security@moov.io). Thanks!

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
