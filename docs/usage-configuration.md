---
layout: page
title: API configuration
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Configuration settings

The following environmental variables can be set to configure behavior in ImageCashLetter.

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE`  | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |

## Data persistence
By design, ImageCashLetter  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart ImageCashLetter will have no files or data saved. Also, no in-memory encryption of the data is performed.