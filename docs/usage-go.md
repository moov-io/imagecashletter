---
layout: page
title: Go library
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.26 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git clone git@github.com:moov-io/imagecashletter.git
$ cd imagecashletter

# Pull down into the Go Module cache (omit -u so Go does not bump the module's go directive)
$ go get github.com/moov-io/imagecashletter

$ go doc github.com/moov-io/imagecashletter CashLetter
```

The package [`github.com/moov-io/imagecashletter`](https://pkg.go.dev/github.com/moov-io/imagecashletter) offers a Go-based Image Cash Letter file reader and writer. To get started, check out a specific example:

| ICL File | Read | Write |
|---------|------|-------|
| [Link](https://github.com/moov-io/imagecashletter/blob/master/examples/imagecashletter-read/iclFile.x937) | [Link](https://github.com/moov-io/imagecashletter/blob/master/examples/imagecashletter-read/main.go) | [Link](https://github.com/moov-io/imagecashletter/blob/master/examples/imagecashletter-write/main.go) |

ImageCashLetter's file handling behaviors can be modified to accommodate your specific use case. This is done by passing _options_ into ICL's `reader` and `writer` during instantiation. For example, to read EBCDIC encoded files you would instantiate a reader with `NewReader(fd, ReadVariableLineLengthOption(), ReadEbcdicEncodingOption())`.

The following options are currently supported:

| Option | Description |
|-----|-----|
| `ReadVariableLineLengthOption` | Allows Reader to split ICL files based on the Inserted Length Field. The 4-byte length prefix is big-endian (Motorola byte order), which is the Federal Reserve default. |
| `ReadLittleEndianVariableLengthOption` | Same as `ReadVariableLineLengthOption`, but the length prefix is little-endian (Intel byte order). Use this for X9.37 files produced by JHA teller systems on Windows. Do not combine it with `ReadVariableLineLengthOption`. |
| `ReadEbcdicEncodingOption` | Allows Reader to decode scanned lines from EBCDIC to UTF-8. |
| `ReadValidateOpts` | Allows skipping validation checks for archived or non-compliant ICL files via ValidateOpts (e.g. SkipAll). Use `file.SetValidation(opts)` after read if needed for later Validate/Create calls. |
| `WriteVariableLineLengthOption` | Instructs the Writer to begin each record with a big-endian Inserted Length Field. The writer has no little-endian counterpart. |
| `WriteEbcdicEncodingOption` | Allows Writer to write file in EBCDIC. |

Run the examples from the repository root (`go run ./examples/...`); paths inside the programs are relative to the root, so `cd examples/imagecashletter-read && go run .` will fail. The write example writes to a temp file and prints the path so it does not overwrite the tracked sample.

```
$ go run ./examples/imagecashletter-read
$ go run ./examples/imagecashletter-write
```

## Command-line tools

`cmd/readImageCashLetter` prints summary counts (or JSON with `-json`) from an ICL file. The default sample at `cmd/readImageCashLetter/BNK20181015-A.icl` is fixed-width ASCII and can be read with no extra flags. Most files under `test/testdata/` use 4-byte inserted-length prefixes and need `-variable-length`; EBCDIC files also need `-ebcdic`. Without those flags, variable-length or EBCDIC files fail with `recordType is an unknown record type`.

```
# Fixed-width ASCII (the default sample in that directory)
$ go run ./cmd/readImageCashLetter -fPath=cmd/readImageCashLetter/BNK20181015-A.icl

# Variable-length ASCII
$ go run ./cmd/readImageCashLetter -variable-length -fPath=test/testdata/valid-ascii.x937

# Variable-length EBCDIC
$ go run ./cmd/readImageCashLetter -variable-length -ebcdic -fPath=test/testdata/valid-ebcdic.x937
```
