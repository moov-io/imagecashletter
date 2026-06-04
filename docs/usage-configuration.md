---
layout: page
title: API configuration
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Configuration settings

The following environmental variables can be set to configure behavior in ImageCashLetter.

| Environmental Variable | Description                                                                                                                                       | Default             |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---------------------|
| `HTTPS_CERT_FILE`      | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP.              | Empty               |
| `HTTPS_KEY_FILE`       | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`.                                                                   | Empty               |
| `MAX_UPLOAD_SIZE`      | Maximum size (in bytes) of HTTP request bodies accepted when creating files via the v2 API. Applies to both JSON and multipart/form-data uploads. | `104857600` (100MB) |
| `READER_BUFFER_SIZE`   | Size (in bytes) of the buffer used when reading ICL files (JSON or raw uploads). | `bufio.MaxScanTokenSize` (64KB) |
| `SKIP_ALL_ON_FILE_CREATE` | If true, the server will use `ValidateOpts.SkipAll` as a base for all file creates (merged with any per-request opts like `?skipAll=...`). Useful for archived/non-compliant data. | false |
| `SKIP_COUNT_VALIDATION_ON_FILE_CREATE` | If true, the server will use `ValidateOpts.SkipCountValidation` as a base for all file creates (merged with per-request). | false |

## Data persistence
By design, ImageCashLetter  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart ImageCashLetter will have no files or data saved. Also, no in-memory encryption of the data is performed.
