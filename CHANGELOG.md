## v0.2.0 (Unreleased)

ADDITIONS

- Add RuneCountInString check to Parse(record string) functions
- cmd/server: bind HTTP server with TLS if HTTPS_* variables are defined

BUG FIXES

- all: check record lengths before parsing them

## v0.1.1 (Released 2019-06-25)

BUG FIXES

- all: fixup panics found from first hour of fuzzing

IMPROVEMENTS

- build: push moov/imagecashletter:latest on 'make release-push'

## v0.1.0 (Released 2019-06-19)

- Initial release
