---
layout: page
title: Binary distribution
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Binary distribution

Download the [latest Moov ImageCashLetter server release](https://github.com/moov-io/imagecashletter/releases) for your operating system and run it from a terminal.

```sh
$ ./imagecashletter-linux-amd64
ts=2026-11-14T16:00:06Z msg="Starting moov-io/imagecashletter server version v0.17.1" level=info package=main
ts=2026-11-14T16:00:06Z msg="admin server listening on [::]:9093" level=info package=main
ts=2026-11-14T16:00:06Z msg="binding to :8083 for HTTP server" level=info package=main
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