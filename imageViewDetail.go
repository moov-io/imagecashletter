// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// Errors specific to a ImageViewDetail Record

// ImageViewDetail Record
type ImageViewDetail struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// recordType defines the type of record.
	recordType string
	// ImageIndicator is a code that indicates the presence and disposition of an image view conveyed in the related
	// ImageViewData.  When an image view is not present (0) then certain conditional fields in this ImageViewDetail and
	// the related ImageViewData shall not be present and will be filled with blank space.
	// Values:
	// 0: Image view not present
	// 1: Image view present, actual check
	// 2: Image view present, not actual check
	// 3: Image view present, unable to determine if value is 1 or 2
	ImageIndicator int `json:"imageIndicator"`
	// ImageCreatorRoutingNumber identifies the financial institution that created the image view
	// in ImageViewData.ImageData.  Format: TTTTAAAAC, where:
	//	TTTT Federal Reserve Prefix
	//	AAAA ABA Institution Identifier
	//	C Check Digit
	//	For a number that identifies a non-financial institution: NNNNNNNNN
	ImageCreatorRoutingNumber string `json:"imageCreatorRoutingNumber"`
	// ImageCreatorDate is the date assigned by the image creator for the image view conveyed in the related
	// ImageViewData.ImageData.
	ImageCreatorDate time.Time `json:"imageCreatorDate"`
	// ImageViewFormatIndicator is a code that identifies the type of image format used in the related
	// ImageViewData.ImageData. The image format type is also commonly specified by reference to the file extension
	// used when image data is saved as an image file.
	// Values:
	// Agreement not required:
	// 00: TIFF 6; Extension: TIF
	// Agreement required:
	// 01: IOCA FS 11; Extension: ICA
	// 20: PNG (Portable Network Graphics); Extension: PNG ‘21’	JFIF (JPEG File Interchange Format); Extension: JPG
	// 22: SPIFF (Still Picture Interchange File Format) (ITU-T Rec. T.84 Annex F); Extension: SPF
	// 23: JBIG data stream (ITU-T Rec. T.82/ISO/IEC 11544:1993);
	// Extension: JBG ‘24’	JPEG 2000 (ISO/IEC 15444-1:2000);
	// Extension: JP2
	ImageViewFormatIndicator string `json:"imageViewFormatIndicator"`
	// ImageViewCompressionAlgorithm is a code that identifies the algorithm or method used to compress the Image Data
	// in the related ImageViewData.ImageData.
	// Values:
	// Agreement not required:
	// 00: Group 4 facsimile compression (ITU-T Rec. T.563/CCITT Rec. T.6)
	// Agreement required:
	// 01: JPEG Baseline (JPEG Interchange Format) (ITU-T Rec. T.81/ISO/IEC 10918)
	// 02: ABIC
	// 21: PNG (Portable Network Graphics)
	// 22: JBIG (ITU-T Rec. T.82/ISO/IEC 11544:1993)
	// 23: JPEG 2000 (ISO/IEC 15444–1:2000)
	ImageViewCompressionAlgorithm string `json:"imageViewCompressionAlgorithm"`
	// ImageViewDataSize is the total number of bytes in ImageViewData.ImageData.  Use of this field is NOT recommended.
	// If data is present it shall be ignored, and ImageViewData.ImageDataLength shall take precedence.
	ImageViewDataSize string `json:"imageViewDataSize"`
	// ViewSideIndicator is a code that indicates the image view conveyed in the related ImageViewData
	// Record.ImageData An image view may be a full view of the item (i.e., the entire full face of the document)
	// or may be a partial view (snippet) as determined by ImageViewDetail.ViewDescriptor.
	// Values:
	// 0: Front image view
	// 1: Rear image view
	ViewSideIndicator int `json:"viewSideIndicator"`
	// ViewDescriptor is a code that indicates the nature of the image view based on ImageViewData.ImageData.
	// Values:
	// 00: Full view
	// 01: Partial view–unspecified Area of Interest
	// 02: Partial view–date Area of Interest
	// 03: Partial view–payee Area of Interest
	// 04: Partial view–convenience amount Area of Interest
	// 05: Partial view–amount in words (legal amount) Area of Interest
	// 06: Partial view–signature Area(s) of Interest
	// 07: Partial view–payor name and address Area of Interest
	// 08: Partial view–MICR line Area of Interest
	// 09: Partial view–memo line Area of Interest
	// 10: Partial view–payor bank name and address Area of Interest
	// 11: Partial view–payee endorsement Area of Interest
	// 12: Partial view–Bank Of First Deposit (BOFD) endorsement Area of Interest
	// 13: Partial view–transit endorsement Area of Interest
	// 14 - 99: Reserved for ImageCashLetter
	ViewDescriptor string `json:"viewDescriptor"`
	// DigitalSignatureIndicator is a code that indicates the presence or absence of a digital signature for the image
	// view contained in ImageViewData.ImageData. If present, the Digital Signature is conveyed in the related
	// ImageViewData.DigitalSignature.
	// Values:
	// 0: Digital Signature is not present
	// 1: Digital Signature is present
	DigitalSignatureIndicator int `json:"digitalSignatureIndicator"`
	// DigitalSignatureMethod is a code that identifies the cryptographic algorithm used to generate and validate the
	// Digital Signature in ImageViewData.DigitalSignature.
	// Values:
	// 00: Digital Signature Algorithm (DSA) with SHA1 (ANSI X9.30)
	// 01: RSA with MD5 (ANSI X9.31)
	// 02: RSA with MDC2 (ANSI X9.31)
	// 03: RSA with SHA1 (ANSI X9.31)
	// 04: Elliptic Curve DSA (ECDSA) with SHA1 (ANSI X9.62)
	// 05 - 99: Reserved for emerging cryptographic algorithms.
	DigitalSignatureMethod string `json:"digitalSignatureMethod"`
	// SecurityKeySize is the length in bits of the cryptographic algorithm key used to create the Digital Signature
	// in ImageViewData.DigitalSignature.
	// Values: 00001–99999
	SecurityKeySize int `json:"securityKeySize"`
	// ProtectedDataStart is a number that represents the offset in bytes from the first byte (counted as byte 1)
	// of the image data in ImageViewData.ImageData to the first byte of the image data protected by the
	// digital signature.
	// Values:
	// 0000000: Digital Signature is applied to the entire image data
	// 0000001–9999999: Valid offset values
	ProtectedDataStart int `json:"protectedDataStart"`
	// ProtectedDataLength is number of contiguous bytes of image data in the related ImageViewData.ImageData
	// protected by the digital signature starting with the byte indicated by the value of the ProtectedDataStart in
	// this ImageViewDetail. The ProtectedDataLength value shall not exceed the ImageViewData.ImageDataLength.
	// Defined Values:
	// 0000000: Digital Signature is applied to entire image data
	// 0000001–9999999: 	Valid length values
	ProtectedDataLength int `json:"protectedDataLength"`
	// ImageRecreateIndicator is a code that indicates whether the sender has the ability to recreate the image view
	// conveyed in the related ImageViewData.ImageData.
	// Values:
	// 0: Sender can recreate the image view for the duration of the agreed upon retention time frames.
	// 1: Sender cannot recreate image view.
	ImageRecreateIndicator int `json:"imageRecreateIndicator"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// OverrideIndicator is a code that indicates to a receiving exchange partner that this image view has a detected
	// image test failure that cannot be corrected and that this view shall be accepted regardless of any image test
	// failures.
	// Values:
	// "": blank/space indicates no observed image test failure present
	// 0: No override information for this view or not applicable
	// 1: Imperfect image
	// A: IQA Fail–Image view reviewed and deemed usable—no alternate format
	// B: IQA Fail–Image view reviewed and deemed usable—alternate format included in this file
	// C: IQA Fail–Image view reviewed and deemed usable–alternate format included in this file and original document
	// available
	// D: IQA Fail–Image view reviewed and deemed usable–alternate format available
	// E: IQA Fail–Image view reviewed and
	// deemed usable–original document available
	// F: IQA Fail–Image view reviewed and deemed usable–original document and alternate format available
	// G: IQA Fail–Image view reviewed and deemed unusable–no alternate format
	// H: IQA Fail–Image view reviewed and deemed unusable–alternate format included in this file
	// I: IQA Fail–Image view reviewed and deemed unusable–alternate format included in this file and original document
	// available
	// J: IQA Fail–Image view reviewed and deemed unusable–alternate format available
	// K: IQA Fail–Image view reviewed and deemed unusable–original document available
	// L: IQA Fail–Image view reviewed and deemed unusable–original document and alternate format available
	// M: IQA Fail–Image view not reviewed–no alternate format
	// N: IQA Fail–Image view not reviewed–alternate format included in this file
	// O: IQA Fail–Image view not reviewed–alternate format included in this file and original
	OverrideIndicator string `json:"overrideIndicator"`
	// reservedTwo is a field reserved for future use.  Reserved should be blank.
	reservedTwo string
	// validator is composed for ImageCashLetter data validation
	validator
	// converters is composed for ImageCashLetter to golang Converters
	converters
}

// NewImageViewDetail returns a new ImageViewDetail with default values for non exported fields
func NewImageViewDetail() ImageViewDetail {
	ivDetail := ImageViewDetail{}
	ivDetail.setRecordType()
	return ivDetail
}

func (ivDetail *ImageViewDetail) setRecordType() {
	if ivDetail == nil {
		return
	}
	ivDetail.recordType = "50"
}

// Parse takes the input record string and parses the ImageViewDetail values
func (ivDetail *ImageViewDetail) Parse(record string) {
	if utf8.RuneCountInString(record) < 67 {
		return // line too short
	}

	// Character position 1-2, Always "50"
	ivDetail.setRecordType()
	// 03-03
	ivDetail.ImageIndicator = ivDetail.parseNumField(record[2:3])
	// 04-12
	ivDetail.ImageCreatorRoutingNumber = ivDetail.parseStringField(record[3:12])
	// 13-20
	ivDetail.ImageCreatorDate = ivDetail.parseYYYYMMDDDate(record[12:20])
	// 21-22
	ivDetail.ImageViewFormatIndicator = ivDetail.parseStringField(record[20:22])
	// 23-24
	ivDetail.ImageViewCompressionAlgorithm = ivDetail.parseStringField(record[22:24])
	// 25-31
	ivDetail.ImageViewDataSize = ivDetail.parseStringField(record[24:31])
	// 32-32
	ivDetail.ViewSideIndicator = ivDetail.parseNumField(record[31:32])
	// 33-34
	ivDetail.ViewDescriptor = ivDetail.parseStringField(record[32:34])
	// 35-35
	ivDetail.DigitalSignatureIndicator = ivDetail.parseNumField(record[34:35])
	// 36-37
	ivDetail.DigitalSignatureMethod = ivDetail.parseStringField(record[35:37])
	// 38-42
	ivDetail.SecurityKeySize = ivDetail.parseNumField(record[37:42])
	// 43-49
	ivDetail.ProtectedDataStart = ivDetail.parseNumField(record[42:49])
	// 50-56
	ivDetail.ProtectedDataLength = ivDetail.parseNumField(record[49:56])
	// 57-57
	ivDetail.ImageRecreateIndicator = ivDetail.parseNumField(record[56:57])
	// 58-65
	ivDetail.UserField = ivDetail.parseStringField(record[57:65])
	// 66-66
	ivDetail.reserved = " "
	// 67-67
	ivDetail.OverrideIndicator = ivDetail.parseStringField(record[66:67])
	// 68-80
	ivDetail.reservedTwo = "             "
}

func (ivDetail *ImageViewDetail) UnmarshalJSON(data []byte) error {
	type Alias ImageViewDetail
	aux := struct {
		*Alias
	}{
		(*Alias)(ivDetail),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ivDetail.setRecordType()
	return nil
}

// String writes the ImageViewDetail struct to a string.
func (ivDetail *ImageViewDetail) String() string {
	if ivDetail == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(ivDetail.recordType)
	buf.WriteString(ivDetail.ImageIndicatorField())
	buf.WriteString(ivDetail.ImageCreatorRoutingNumberField())
	buf.WriteString(ivDetail.ImageCreatorDateField())
	buf.WriteString(ivDetail.ImageViewFormatIndicatorField())
	buf.WriteString(ivDetail.ImageViewCompressionAlgorithmField())
	buf.WriteString(ivDetail.ImageViewDataSizeField())
	buf.WriteString(ivDetail.ViewSideIndicatorField())
	buf.WriteString(ivDetail.ViewDescriptorField())
	buf.WriteString(ivDetail.DigitalSignatureIndicatorField())
	buf.WriteString(ivDetail.DigitalSignatureMethodField())
	buf.WriteString(ivDetail.SecurityKeySizeField())
	buf.WriteString(ivDetail.ProtectedDataStartField())
	buf.WriteString(ivDetail.ProtectedDataLengthField())
	buf.WriteString(ivDetail.ImageRecreateIndicatorField())
	buf.WriteString(ivDetail.UserFieldField())
	buf.WriteString(ivDetail.reservedField())
	buf.WriteString(ivDetail.OverrideIndicatorField())
	buf.WriteString(ivDetail.reservedTwoField())
	return buf.String()
}

// Validate performs ImageCashLetter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (ivDetail *ImageViewDetail) Validate() error {
	if err := ivDetail.fieldInclusion(); err != nil {
		return err
	}
	// Mandatory
	if ivDetail.recordType != "50" {
		msg := fmt.Sprintf(msgRecordType, 50)
		return &FieldError{FieldName: "recordType", Value: ivDetail.recordType, Msg: msg}
	}
	// Mandatory
	if err := ivDetail.isImageIndicator(ivDetail.ImageIndicator); err != nil {
		return &FieldError{FieldName: "ImageIndicator",
			Value: ivDetail.ImageIndicatorField(), Msg: err.Error()}
	}
	// Conditional
	if ivDetail.ImageViewFormatIndicator != "" {
		if err := ivDetail.isImageViewFormatIndicator(ivDetail.ImageViewFormatIndicator); err != nil {
			return &FieldError{FieldName: "ImageViewFormatIndicator",
				Value: ivDetail.ImageViewFormatIndicator, Msg: err.Error()}
		}
	}
	// Conditional
	if ivDetail.ImageViewCompressionAlgorithm != "" {
		if err := ivDetail.isImageViewCompressionAlgorithm(ivDetail.ImageViewCompressionAlgorithm); err != nil {
			return &FieldError{FieldName: "ImageViewCompressionAlgorithm",
				Value: ivDetail.ImageViewCompressionAlgorithm, Msg: err.Error()}
		}
	}
	// Mandatory
	if err := ivDetail.isViewSideIndicator(ivDetail.ViewSideIndicator); err != nil {
		return &FieldError{FieldName: "ViewSideIndicator",
			Value: ivDetail.ViewSideIndicatorField(), Msg: err.Error()}
	}
	// Mandatory
	if err := ivDetail.isViewDescriptor(ivDetail.ViewDescriptor); err != nil {
		return &FieldError{FieldName: "ViewDescriptor",
			Value: ivDetail.ViewDescriptor, Msg: err.Error()}
	}
	// Conditional
	if ivDetail.DigitalSignatureIndicatorField() != "" {
		if err := ivDetail.isDigitalSignatureIndicator(ivDetail.DigitalSignatureIndicator); err != nil {
			return &FieldError{FieldName: "DigitalSignatureIndicator",
				Value: ivDetail.DigitalSignatureIndicatorField(), Msg: err.Error()}
		}
	}
	// Conditional
	if ivDetail.DigitalSignatureMethod != "" {
		if ivDetail.DigitalSignatureMethod == "0" && IsFRBCompatibilityModeEnabled() {
			ivDetail.DigitalSignatureMethod = "00"
		}
		if err := ivDetail.isDigitalSignatureMethod(ivDetail.DigitalSignatureMethod); err != nil {
			return &FieldError{FieldName: "DigitalSignatureMethod",
				Value: ivDetail.DigitalSignatureMethod, Msg: err.Error()}
		}
	}
	// Conditional
	if ivDetail.ImageRecreateIndicatorField() != "" {
		if err := ivDetail.isImageRecreateIndicator(ivDetail.ImageRecreateIndicator); err != nil {
			return &FieldError{FieldName: "ImageRecreateIndicator",
				Value: ivDetail.ImageRecreateIndicatorField(), Msg: err.Error()}
		}
	}
	// Conditional
	if ivDetail.OverrideIndicator != "" {
		if err := ivDetail.isOverrideIndicator(ivDetail.OverrideIndicator); err != nil {
			return &FieldError{FieldName: "OverrideIndicator",
				Value: ivDetail.OverrideIndicatorField(), Msg: err.Error()}
		}
	}
	if err := ivDetail.isAlphanumericSpecial(ivDetail.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: ivDetail.UserField, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (ivDetail *ImageViewDetail) fieldInclusion() error {
	if ivDetail.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: ivDetail.recordType,
			Msg:   msgFieldInclusion + ", did you use ImageViewDetail()?"}
	}
	if ivDetail.ImageCreatorRoutingNumber == "" {
		return &FieldError{FieldName: "ImageCreatorRoutingNumber",
			Value: ivDetail.ImageCreatorRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use ImageViewDetail()?"}
	}
	if ivDetail.ImageCreatorRoutingNumberField() == "000000000" && !IsFRBCompatibilityModeEnabled() {
		return &FieldError{FieldName: "ImageCreatorRoutingNumber",
			Value: ivDetail.ImageCreatorRoutingNumber,
			Msg:   msgFieldInclusion + ", did you use ImageViewDetail()?"}
	}
	if ivDetail.ImageCreatorDate.IsZero() && !IsFRBCompatibilityModeEnabled() {
		return &FieldError{FieldName: "ImageCreatorDate",
			Value: ivDetail.ImageCreatorDate.String(),
			Msg:   msgFieldInclusion + ", did you use ImageViewDetail()?"}
	}
	if ivDetail.ViewDescriptor == "" {
		return &FieldError{FieldName: "ViewDescriptor",
			Value: ivDetail.ViewDescriptor,
			Msg:   msgFieldInclusion + ", did you use ImageViewDetail()?"}
	}
	return nil
}

// ImageIndicatorField gets a string of the ImageIndicator field
func (ivDetail *ImageViewDetail) ImageIndicatorField() string {
	return ivDetail.numericField(ivDetail.ImageIndicator, 1)
}

// ImageCreatorRoutingNumberField gets the ImageCreatorRoutingNumber field
func (ivDetail *ImageViewDetail) ImageCreatorRoutingNumberField() string {
	return ivDetail.stringField(ivDetail.ImageCreatorRoutingNumber, 9)
}

// ImageCreatorDateField gets the ImageCreatorDate field, format YYYYMMDD
func (ivDetail *ImageViewDetail) ImageCreatorDateField() string {
	return ivDetail.formatYYYYMMDDDate(ivDetail.ImageCreatorDate)
}

// ImageViewFormatIndicatorField gets the ImageViewFormatIndicator field
func (ivDetail *ImageViewDetail) ImageViewFormatIndicatorField() string {
	return ivDetail.alphaField(ivDetail.ImageViewFormatIndicator, 2)
}

// ImageViewCompressionAlgorithmField gets the ImageViewCompressionAlgorithm field
func (ivDetail *ImageViewDetail) ImageViewCompressionAlgorithmField() string {
	return ivDetail.alphaField(ivDetail.ImageViewCompressionAlgorithm, 2)
}

// ImageViewDataSizeField gets the ImageViewDataSize field
func (ivDetail *ImageViewDetail) ImageViewDataSizeField() string {
	return ivDetail.alphaField(ivDetail.ImageViewDataSize, 7)
}

// ViewSideIndicatorField gets a string of the ViewSideIndicator field
func (ivDetail *ImageViewDetail) ViewSideIndicatorField() string {
	return ivDetail.numericField(ivDetail.ViewSideIndicator, 1)
}

// ViewDescriptorField gets the ViewDescriptor field
func (ivDetail *ImageViewDetail) ViewDescriptorField() string {
	return ivDetail.alphaField(ivDetail.ViewDescriptor, 2)
}

// DigitalSignatureIndicatorField gets a string of the DigitalSignatureIndicator field
func (ivDetail *ImageViewDetail) DigitalSignatureIndicatorField() string {
	return ivDetail.numericField(ivDetail.DigitalSignatureIndicator, 1)
}

// DigitalSignatureMethodField gets the DigitalSignatureMethod field
func (ivDetail *ImageViewDetail) DigitalSignatureMethodField() string {
	return ivDetail.alphaField(ivDetail.DigitalSignatureMethod, 2)
}

// SecurityKeySizeField gets the SecurityKeySize field
func (ivDetail *ImageViewDetail) SecurityKeySizeField() string {
	if ivDetail.SecurityKeySize <= 0 {
		return "     "
	}
	return ivDetail.numericField(ivDetail.SecurityKeySize, 5)

}

// ProtectedDataStartField gets the ProtectedDataStart field
func (ivDetail *ImageViewDetail) ProtectedDataStartField() string {
	return ivDetail.numericField(ivDetail.ProtectedDataStart, 7)
}

// ProtectedDataLengthField gets the ProtectedDataLength field
func (ivDetail *ImageViewDetail) ProtectedDataLengthField() string {
	return ivDetail.numericField(ivDetail.ProtectedDataLength, 7)
}

// ImageRecreateIndicatorField gets a string of the ImageRecreateIndicator field
func (ivDetail *ImageViewDetail) ImageRecreateIndicatorField() string {
	return ivDetail.numericField(ivDetail.ImageRecreateIndicator, 1)
}

// UserFieldField gets the UserField field
func (ivDetail *ImageViewDetail) UserFieldField() string {
	return ivDetail.alphaField(ivDetail.UserField, 8)
}

// reservedField gets the reserved field
func (ivDetail *ImageViewDetail) reservedField() string {
	return ivDetail.alphaField(ivDetail.reserved, 1)
}

// OverrideIndicatorField gets the OverrideIndicator field
func (ivDetail *ImageViewDetail) OverrideIndicatorField() string {
	return ivDetail.alphaField(ivDetail.OverrideIndicator, 1)
}

// reservedTwoField gets the reserved field
func (ivDetail *ImageViewDetail) reservedTwoField() string {
	return ivDetail.alphaField(ivDetail.reservedTwo, 13)
}
