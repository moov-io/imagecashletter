---
layout: page
title: Go library
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.18 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/imagecashletter/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/imagecashletter.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/imagecashletter

$ go doc github.com/moov-io/imagecashletter CashLetter
```

The package [`github.com/moov-io/imagecashletter`](https://pkg.go.dev/github.com/moov-io/imagecashletter) offers a Go-based Image Cash Letter file reader and writer. To get started, check out a specific example:

| ICL File | Read | Write |
|---------|------|-------|
| [Link](https://github.com/moov-io/imagecashletter/blob/master/examples/imagecashletter-read/iclFile.x937) | [Link](https://github.com/moov-io/imagecashletter/blob/master/examples/imagecashletter-read/main.go) | [Link](https://github.com/moov-io/imagecashletter/blob/master/examples/imagecashletter-write/main.go) |

ImageCashLetter's file handling behaviors can be modified to accommodate your specific use case. This is done by passing _options_ into ICL's `reader` and `writer` during instantiation. For example, to read EBCDID encoded files you would instantiate a reader with `NewReader(fd, ReadVariableLineLengthOption(), ReadEbcdicEncodingOption())`.

The following options are currently supported:

| Option | Description |
|-----|-----|
| `ReadVariableLineLengthOption` | Allows Reader to split ICL files based on the Inserted Length Field. |
| `ReadEbcdicEncodingOption` | Allows Reader to decode scanned lines from EBCDIC to UTF-8. |
| `WriteVariableLineLengthOption` | Instructs the Writer to begin each record with the appropriate Inserted Length Field. |
| `WriteEbcdicEncodingOption` | Allows Writer to write file in EBCDIC. |
