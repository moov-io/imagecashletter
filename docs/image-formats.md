---
layout: page
title: Image Formats
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Image Formats
Image data issues are some of the most common problems with ICL files. Each image file holds some general information like size, compression, and orientation, which should be specified in the Image View Detail Record (Type 50). Primary image data is located under the Image View Data Record (Type 52) and must be binary — we’ve seen cases where it’s incorrectly Base64 encoded. This includes a header section that expands on the contents of Record Type 50, followed by image raster data. Lastly, information pertaining to image quality should be supplied under the Image View Analysis Record (Type 54).

There are many possible image file formats, such as TIFF, JPEG, GIF, etc. However, TIFF 6.0 is the current standard for image cash letters and other formats will likely cause problems during processing. As there are many slight variations of TIFF 6.0, [ANSI X9.100-181](https://webstore.ansi.org/standards/ascx9/ansix91001812010) (TIFF Image Format for Image Exchange) is used to prevent incompatibility with processing systems.

If an image has multiple views, it will require separate image records for each view. A view must have associated detail and data records, but not necessarily an analysis record. Detail and data records are typically provided by an image capture institution, while the analysis record is filled out after quality assurance checks by the exchange network.
