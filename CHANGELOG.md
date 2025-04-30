## v0.12.0 (Released 2025-04-30)

IMPROVEMENTS

- add support for return code `N` for `CashLetterHeader.ReturnsIndicator` ([#403](https://github.com/moov-io/imagecashletter/pull/403) by [@obscure76](https://github.com/obscure76))
- validation logic updates for FRB compatibility mode ([#398](https://github.com/moov-io/imagecashletter/pull/398) by [@mtnat](https://github.com/mtnat))
- skip `ReturnLocationRoutingNumber` zero validation when FRB compatibility is enabled ([#388](https://github.com/moov-io/imagecashletter/pull/388) by [@Hillrunner2008](https://github.com/Hillrunner2008))

BUILD

- build(deps): bump nokogiri from 1.16.7 to 1.18.8
- build(deps): bump uri from 0.13.0 to 0.13.2
- fix(deps): update module github.com/moov-io/base to v0.54.4
- fix(deps): update module golang.org/x/oauth2 to v0.26.0
- fix(deps): update module github.com/stretchr/testify to v1.10.0
- build(webui): check both `lib/` and `misc/` for `wasm_exec.js`
- chore(deps): update golang docker tag to v1.24

## v0.11.4 (Released 2024-12-05)

IMPROVEMENTS

- extend ASCII compatibility in FRB compatibility mode (See [PR #386](https://github.com/moov-io/imagecashletter/pull/386) by [@Hillrunner2008](https://github.com/Hillrunner2008))

BUILD

- build(deps): bump rexml from 3.3.6 to 3.3.9 in /docs
- fix(deps): update module golang.org/x/oauth2 to v0.24.0
- build: increase fuzz job timeout
- fix(deps): update module github.com/stretchr/testify to v1.10.0

## v0.11.3 (Released 2024-10-28)

IMPROVEMENTS

- build bundle control records when creating files (#381)
- fix: Allow empty EndorsingBankItemSequenceNumber in FRB mode (#376)
- fix: nil checks from fuzzers

BUILD

- build(deps): bump rexml from 3.2.8 to 3.3.6 in /docs
- build(deps): bump webrick from 1.8.1 to 1.8.2 in /docs
- chore(deps): update dependency github-pages to v232
- fix(deps): update module github.com/moov-io/base to v0.51.1
- fix(deps): update module github.com/prometheus/client_golang to v1.20.5

## v0.11.0 (Released 2024-07-08)

ADDITIONS

- config: add an `FRB_COMPATIBILITY_MODE` environment variable to enable Federal Reserve Bank (FRB) compatibility mode (See [PR #307](https://github.com/moov-io/imagecashletter/pull/364) by [@smithbk](https://github.com/smithbk))

BUG FIXES

- writer: correct ordering of records in return images (See [PR #307](https://github.com/moov-io/imagecashletter/pull/361) by [@smithbk](https://github.com/smithbk))
- client: add missing quotes to `payorBankCheckDigit` example value in OpenAPI specification

BUILD

- fix(deps): update minimum Go version to 1.21
- fix(deps): update module golang.org/x/oauth2 to v0.21.0
- fix(deps): update module github.com/moov-io/base to v0.50.0

## v0.10.1 (Released 2024-04-17)

IMPROVEMENTS

- fix: cap buffer growth

BUILD

- fix(deps): update module golang.org/x/oauth2 to v0.19.0
- build(deps): bump google.golang.org/protobuf from 1.32.0 to 1.33.0
- fix(deps): update module github.com/gdamore/encoding to v1.0.1

## v0.10.0 (Released 2024-02-19)

ADDITIONS

- api: introduce a `v2` files API with an improved Create File endpoint

## v0.9.5 (Released 2023-08-15)

IMPROVEMENTS

- CashLetterControl: make `SettlementDate` field optional to support additional specification variations

## v0.9.4 (Released 2023-08-14)

IMPROVEMENTS

- client: update `eCEInstitutionRoutingNumber` to `eceInstitutionRoutingNumber` to follow existing naming conventions ((See [PR #312](https://github.com/moov-io/imagecashletter/pull/312) by [@DennyWeinberg](https://github.com/DennyWeinberg)))
- all: use `testify` for assertions in all tests, improving readability
- reader: add `Unwrap()` method to `ParseError` so `errors.Is()` and `errors.As()` behave as expected

BUG FIXES

- ImageViewAnalysis: fix copy/paste bug in `AmountInWordsUsabilityField()`
- ReturnDetailAddendumB: implement validation for the `PayorBankBusinessDate` field
- ReturnDetailAddendumD: implement validation for the `EndorsingBankItemSequenceNumber` field
- UserPayeeEndorsement: implement validation for the `FormatTypeVersionLevel` field

## v0.9.3 (Released 2023-07-31)

BUILD

- build: update minimum Go version to 1.20
- build: expose ports for HTTP API and health check
- deps: update openapi codegen to v4.3.1
- deps: update module github.com/moov-io/base to v0.45.1
- deps: update module github.com/prometheus/client_golang to v1.16.0
- deps: update module github.com/stretchr/testify to v1.8.4
- deps: update module golang.org/x/oauth2 to v0.10.0
- deps: bump activesupport to 7.0.4.3 in /docs

BUG FIXES

- client: rename File Header field `testFileIndicator` to `testIndicator` (See [PR #307](https://github.com/moov-io/imagecashletter/pull/307) by [@DennyWeinberg](https://github.com/DennyWeinberg))
- credit: fix validation of `ECEInstitutionItemSequenceNumber`

## v0.9.2 (Released 2023-03-07)

IMPROVEMENTS

- cmd/server: use READER_BUFFER_SIZE to override default bufio buffer size

BUILD

- build(deps): bump activesupport from 6.0.3.4 to 6.0.6.1 in /docs
- build(deps): bump golang.org/x/net from 0.6.0 to 0.7.0
- build(deps): bump nokogiri from 1.13.6 to 1.13.10 in /docs
- fix(deps): update module github.com/moov-io/base to v0.39.0
- fix(deps): update module github.com/prometheus/client_golang to v1.13.1
- fix(deps): update module github.com/stretchr/testify to v1.8.2
- fix(deps): update module golang.org/x/oauth2 to v0.6.0

## v0.9.1 (Released 2022-08-11)

ADDITIONS

- file: include new `Credit` (record type 61) in file, reader, and writer

IMPROVEMENTS

- all: replace deprecated `ioutil` function calls with equivalents from `io` and `os`

BUG FIXES

- cashLetter: fix out of range error when iterating over `ReturnDetailAddendumD` records

BUILD

- deps: upgrade docker build to Go 1.19
- deps: update model github.com/prometheus/client_golang to v1.13.0
- deps: update model golang.org/x/oauth2 to v0.0.0-20220808172628-8227340efae7

## v0.9.0 (Released 2022-08-02)

ADDITIONS

- cashLetter: add credit item (record type 61)

BUILD

 - deps: update module github.com/moov-io/base to v0.33.0
 - deps: update module github.com/prometheus/client_golang to v1.12.2
 - deps: update module github.com/stretchr/testify to v1.8.0

## v0.8.0 (Released 2022-03-29)

IMPROVEMENTS

- cashLetter: return error if CashLetterControl is missing or invalid

## v0.7.4 (Released 2022-02-10)

IMPROVEMENTS

- checkDetail: make `MICRValidIndicator` validation more flexible
- reader: handle check image data larger than `bufio`'s default buffer size
- examples: fix writer to build a valid file, including a check image

BUILD

- deps: update github.com/prometheus/client_golang to v1.12.1

## v0.7.3 (Released 2022-01-18)

BUG FIXES

- cmd/webui: pass a data URI through instead of raw file contents

BUILD

- fix(deps): update module github.com/moov-io/base to v0.27.5

## v0.7.2 (Released 2021-12-09)

IMPROVEMENTS

- returnDetailAddendumD: make `EndorsingBankItemSequenceNumber` conditional to support older specifications

## v0.7.1 (Released 2021-10-25)

IMPROVEMENTS

- returnDetailAddendumA: make `BOFDItemSequenceNumber` conditional to support more spec versions

BUILD

- deps: update moov-io/base to v0.26.0

## v0.7.0 (Released 2021-10-11)

IMPROVEMENTS

- api: make ebcdic the default encoding for readers and writers

BUILD

- deps: update moov-io/base to v0.25.0

## v0.6.6 (Released 2021-08-13)

BUG FIXES

- reader: nil check before parsing CollectionTypeIndicator

BUILD

- fix: Dockerfile to reduce vulnerabilities
- fix(deps): update module github.com/moov-io/base to v0.21.1

## v0.6.5 (Released 2021-07-16)

BUILD

- build(deps): bump addressable from 2.7.0 to 2.8.0 in /docs
- fix: Dockerfile to reduce vulnerabilities
- fix: Dockerfile.webui to reduce vulnerabilities

## v0.6.4 (Released 2021-07-07)

BREAKING CHANGES

- reader: pass CollectionTypeIndicator to CashLetterControl's `Validate()` function (Issue: [#185](https://github.com/moov-io/imagecashletter/pull/185))

BUILD

- fix(deps): update module github.com/go-kit/kit to v0.11.0
- fix(deps): update module github.com/moov-io/base to v0.20.1

## v0.6.3 (Released 2021-06-28)

BUG FIXES

- imageViewDetail: replace the conditional Security Key Size field with whitespace if it isn't a valid value

IMPROVEMENTS

- api: return HTTP 404 instead of empty response when a specified resource is not found

## v0.6.2 (Released 2021-01-29)

BUG FIXES

- returnDetailAddendumB: make `PayorBankBusinessDate` conditional; don't marshal empty `time.Time{}`

## v0.6.1 (Released 2021-01-20)

IMPROVEMENTS

- readme: add table of contents
- readme: document reader/writer options
- readme: update cURL examples (and fix broken example files)
- readme: add a section about running on Google Cloud Platform

BUG FIXES

- icl: don't overwrite existing file values with defaults

BUILD

- deps: update github.com/moov-io/base to v0.15.4
- deps: update github.com/prometheus/client_golang to v1.8.0
- deps: update github.com/stretchr/testify to v1.7.0

## v0.6.0 (Released 2020-12-24)

ADDITIONS

- writer: always write collated image views (e.g. 50, 52, 50, 52)

BUILD:

- deps: update github.com/moov-io/base to v0.15.2
- deps: update github.com/moov-io/paygate to v0.9.2

## v0.5.2 (Released 2020-12-10)

BUG FIXES

- returnDetail: change ReturnNotificationIndicator from `int` to `string` to prevent storing `0` when the field should have been empty

## v0.5.1

BUG FIXES

- writer: fix indexing error when writing collated images

## v0.5.0

ADDITIONS

- all: update project to use [Moov's logger](https://github.com/moov-io/base/log) instead of [Go kit](https://github.com/go-kit/kit/tree/master/log)
- cmd/webui: initial setup with pretty JSON formatting example
- file: Add support for writing EBCDIC
- file: Allow File to be written as DTSU with control bytes dictating line lengths rather than line breaks
- imageViewData: add decode method for ImageData
- imageViewData: attempt base64 decode when generating a file
- server: pass `ReaderOption`/`WriterOption` for variable line lengths

BUG FIXES

- file: Do not overwrite institution sequence number if one already exists
- file: populate recordType inside each record's JSON unmarshal
- file: setup additional nil checks

IMPROVEMENTS

- api: add returnLocationRoutingNumber to BundleHeader
- api: include missing imageViewDataSize
- api: use a correct example timestamp
- examples: add the Inserted Length Field to each record in the imagecacheletter-read example file
- examples: use `ReadVariableLineLengthOption` in examples/imagecacheletter-read

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
