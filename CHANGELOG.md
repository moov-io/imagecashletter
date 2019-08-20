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
