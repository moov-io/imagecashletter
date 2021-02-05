---
layout: page
title: Binary Distribution
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Binary Distribution

Download the [latest Moov ImageCashLetter server release](https://github.com/moov-io/imagecashletter/releases) for your operating system and run it from a terminal.

```sh
$ ./imagecashletter-darwin-amd64
ts=2019-06-20T23:23:44.870717Z caller=main.go:75 startup="Starting imagecashletter server version v0.2.0"
ts=2019-06-20T23:23:44.871623Z caller=main.go:135 transport=HTTP addr=:8083
ts=2019-06-20T23:23:44.871692Z caller=main.go:125 admin="listening on :9093"
```

## Connecting to Moov ImageCashLetter

The Moov ImageCashLetter service will be running on port `8083` (with an admin port on `9093`).

Confirm that the service is running by issuing the following command or simply visiting [localhost:8083/ping](http://localhost:8083/ping) in your browser.

```bash
$ curl http://localhost:8083/ping
PONG

$ curl http://localhost:8083/files
null
```