// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// Errors specific to an ImageView

// ImageView contains ImageViewDetail, ImageViewData, and ImageViewAnalysis Records
type ImageView struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// ImageViewDetail is an ICL Image View Detail Record
	ImageViewDetail *ImageViewDetail `json:"imageViewDetail,omitempty"`
	// ImageViewData is an ICL Image View Data Record
	ImageViewData *ImageViewData `json:"imageViewData,omitempty"`
	// ImageViewAnalysis is an ICL Image View Analysis Record
	ImageViewAnalysis *ImageViewAnalysis `json:"imageViewAnalysis,omitempty"`
}

// NewImageView returns an ImageView
func NewImageView() *ImageView {
	imageView := &ImageView{}
	return imageView
}
