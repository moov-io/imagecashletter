// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Errors specific to a ImageViewAnalysis Record

// ImageViewAnalysis Record
type ImageViewAnalysis struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// recordType defines the type of record.
	recordType string
	// GlobalImageQuality is a code that indicates whether the image view was tested for any of the conditions related
	// to image quality defined in the Image Quality Information.
	// Values:
	// 0: The image was not tested for any of the image quality conditions
	// 1: The image was tested and one or more image quality conditions were reported
	// 2: The image was tested and no image quality conditions were reported
	GlobalImageQuality int `json:"globalImageQuality"`
	// GlobalImageUsability is a code that indicates whether the image view was tested for any of the conditions
	// related to image usability defined in the Image Usability Information.
	// Values:
	// 0: The image was not tested for any of the image usability conditions
	// 1: The image was tested and one or more image usability conditions were reported
	// 2: The image was tested and no image usability conditions were reported
	GlobalImageUsability int `json:"globalImageUsability"`
	// ImagingBankSpecificTest designates the capture institution may be able to perform specific tests that can
	// indicate a potentially problematic image view caused by conditions other than those listed in the Image Quality
	// and Image Usability Information fields. By mutual agreement, clearing partners can use the UserField to report
	// the presence or absence of additional image conditions found through tests that are particular to the specific
	// imaging institution. The meaning and interpretation of the User Field data must be understood and agreed upon
	// between participants.
	// Values:
	// 0: No user-defined tests were made for other image quality/usability conditions
	// 1: Other user-defined image conditions were tested and one or more are reported in the UserField
	// 2: Other user-defined image conditions were tested and none are reported in the UserField
	ImagingBankSpecificTest int `json:"imagingBankSpecificTest"`
	// PartialImage is a code that indicates if only a portion of the image view is represented digitally while the
	// other portion is suspected to be missing or corrupt.
	// Values:
	// 0: Test not done
	// 1: Condition present
	// 2: Condition not present
	PartialImage int `json:"partialImage"`
	// ExcessiveImageSkew is a code that indicates if the image view skew exceeds an acceptable value. This value is
	// specific to the imaging institution’s own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: Condition present
	// 2: Condition not present
	ExcessiveImageSkew int `json:"excessiveImageSkew"`
	// PiggybackImage is A code that indicates if a “piggyback” condition has been detected. With a “piggyback”
	// condition, the intended image view may be extended, obscured, or replaced by image(s) of additional document(s).
	// A piggyback occurs when two or more documents are fed together and captured as one document when only a single
	// document should have been fed and captured.
	// Values:
	// 0: Test not done
	// 1: Condition present
	// 2: Condition not present
	PiggybackImage int `json:"piggybackImage"`
	// TooLightOrTooDark is a code that indicates if the image view is too light or too dark. The value is specific to
	// the imaging institution’s own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: Condition present
	// 2: Condition not present
	TooLightOrTooDark int `json:"tooLightOrTooDark"`
	// StreaksAndOrBands is a A code that indicates if the image view is likely corrupted due to streaks and/or bands.
	// Streaks and bands can be caused by such problems as dirt, dust, ink, or debris on a lens or in the optical path,
	// and failures in the imaging equipment scanner.
	// Values:
	// 0: Test not done
	// 1: Condition present
	// 2: Condition not present
	StreaksAndOrBands int `json:"streaksAndOrBands"`
	// BelowMinimumImageSize is a code that indicates if the size of the compressed image view is below an acceptable
	// value. The value is specific to the imaging institution’s own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: Condition present
	// 2: Condition not present
	BelowMinimumImageSize int `json:"belowMinimumImageSize"`
	// ExceedsMaximumImageSize is a code that indicates if the size of the compressed image view is above an
	// acceptable value. The value is specific to the imaging institution’s own defined requirements and/or
	// constraints.
	// Values:
	// 0: Test not done
	// 1: Condition present
	// 2: Condition not present
	ExceedsMaximumImageSize int `json:"exceedsMaximumImageSize"`
	// reserved is a field reserved for future use.  Reserved should be blank.
	reserved string
	// ImageEnabledPOD is a code that indicates if the image view was used within an image-enabled POD
	// (Proof of Deposit) application.
	// Values:
	// 0: It is unknown whether the image was used within an image-enabled POD application.
	// 1: Image was not used within an image-enabled POD application.
	// 2: Image was used within an image-enabled POD application.
	ImageEnabledPOD int `json:"imageEnabledPOD"`
	// SourceDocumentBad is a code that indicates if it is possible to obtain a better image from the source document
	// when it is known that the current image of the document is unusable.
	// Values:
	// 0: Test not done
	// 1: Image is unusable. It is not possible to obtain a better image since the source document is bad.
	// 2: Image is unusable. It is likely possible to obtain a better image since the source document is good.
	SourceDocumentBad int `json:"sourceDocumentBad"`
	// DateUsability is a code that indicates if the date Area of Interest is usable and readable from the image. The
	// definition of the Area of Interest for image usability testing purposes is specific to the imaging institution's
	// own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the date is unusable and unreadable
	// 2: From the image the date is usable and readable
	DateUsability int `json:"dateUsability"`
	// PayeeUsability is a code that indicates if the payee name Area of Interest is usable and readable from the
	// image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging
	// institution's own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the payee is unusable and unreadable
	// 2: From the image the payee is usable and readable
	PayeeUsability int `json:"payeeUsability"`
	// ConvenienceAmountUsability is a code that indicates if the convenience amount Area of Interest is usable and
	// readable from the image. The definition of the Area of Interest for image usability testing purposes is
	// specific to the imaging institution's own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the convenience amount is unusable and unreadable
	// 2: From the image the convenience amount is usable and readable
	ConvenienceAmountUsability int `json:"convenienceAmountUsability"`
	// AmountInWordsUsability is a code that indicates if the amount in words (legal amount) Area of Interest is usable
	// and readable from the image. The definition of the Area of Interest for image usability testing purposes is
	// specific to the imaging institution's own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the amount in words is unusable and unreadable
	// 2: From the image the amount in words is usable and readable
	AmountInWordsUsability int `json:"amountInWordsUsability"`
	// SignatureUsability is a code that indicates if the signature Area of Interest is usable and readable from the
	// image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging
	// institution's own defined requirements and/or constraints.
	// Values:
	// 0 Test not done
	// 1 From the image the signature(s) is/are unusable and unreadable
	// 2 From the image the signature(s) is/are usable and readable
	SignatureUsability int `json:"signatureUsability"`
	// PayorNameAddressUsability is a code that indicates if the payor name and address Area of Interest is usable and
	// readable from the image. The definition of the Area of Interest for image usability testing purposes is specific
	// to the imaging institution's own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the payor name and address is unusable and unreadable
	// 2: From the image the payor name and address is usable and readable
	PayorNameAddressUsability int `json:"payorNameAddressUsability"`
	// MICRLineUsability is a code that indicates if the MICR line Area of Interest is usable and readable from the
	// image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging
	// institution's own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the MICR line is unusable and unreadable
	// 2: From the image the MICR line is usable and readable
	MICRLineUsability int `json:"micrLineUsability"`
	// MemoLineUsability is code that indicates if the memo line Area of Interest is usable and readable from the
	// image. The definition of the Area of Interest for image usability testing purposes is specific to the imaging
	// institution's own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the memo line is unusable and unreadable
	// 2: From the image the memo line is usable and readable
	MemoLineUsability int `json:"memoLineUsability"`
	// PayorBankNameAddressUsability is a code that indicates if the payor bank name and address Area of Interest is
	// usable and readable from the image. The definition of the Area of Interest for image usability testing purposes
	// is specific to the imaging institution's own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the payor bank name and address is unusable and unreadable
	// 2: From the image the payor bank name and address is usable and readable
	PayorBankNameAddressUsability int `json:"payorBankNameAddressUsability"`
	// PayeeEndorsementUsability is a code that indicates if the payee endorsement Area of Interest is usable and
	// readable from the image. The definition of the Area of Interest for image usability testing purposes is specific
	// to the imaging institution's own defined requirements and/or constraints.
	// Values:
	// 0: Test not done
	// 1: From the image the payor bank name and address is unusable and unreadable
	// 2: From the image the payor bank name and address is usable and readable
	PayeeEndorsementUsability int `json:"payeeEndorsementUsability"`
	// BOFDEndorsementUsability is a code that indicates if the Bank of First Deposit (BOFD) endorsement Area of
	// Interest is usable and readable from the image. The definition of the Area of Interest for image usability
	// testing purposes is specific to the imaging institution's own defined requirements and/or constraints.
	// 0: Test not done
	// 1: From the image the BOFD endorsement is unusable and unreadable
	// 2: From the image the BOFD endorsement is usable and readable
	BOFDEndorsementUsability int `json:"bofdEndorsementUsability"`
	// TransitEndorsementUsability is a code that indicates if the transit endorsement Area of Interest is usable and
	// readable from the image. The definition of the Area of Interest for image usability testing purposes is specific
	// to the imaging institution's own defined requirements and/or constraints.
	// Values:Values:
	// 0: Test not done
	// 1: From the image the transit endorsement(s) is/are unusable and unreadable
	// 2: From the image the transit endorsement(s) is/are usable and readable
	TransitEndorsementUsability int `json:"transitEndorsementUsability"`
	// reservedTwo is a field reserved for future use.  Reserved should be blank.
	reservedTwo string
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField"`
	// reservedThree is a field reserved for future use.  Reserved should be blank.
	reservedThree string
	// validator is composed for ImageCashLetter data validation
	validator
	// converters is composed for ImageCashLetter to golang Converters
	converters
}

// NewImageViewAnalysis returns a new ImageViewAnalysis with default values for non exported fields
func NewImageViewAnalysis() ImageViewAnalysis {
	ivAnalysis := ImageViewAnalysis{}
	ivAnalysis.setRecordType()
	return ivAnalysis
}

func (ivAnalysis *ImageViewAnalysis) setRecordType() {
	if ivAnalysis == nil {
		return
	}
	ivAnalysis.recordType = "54"
}

// Parse takes the input record string and parses the ImageViewAnalysis values
func (ivAnalysis *ImageViewAnalysis) Parse(record string) {
	if utf8.RuneCountInString(record) < 65 {
		return // line too short
	}

	// Character position 1-2, Always "54"
	ivAnalysis.setRecordType()
	// 03-03
	ivAnalysis.GlobalImageQuality = ivAnalysis.parseNumField(record[2:3])
	// 04-04
	ivAnalysis.GlobalImageUsability = ivAnalysis.parseNumField(record[3:4])
	// 05-05
	ivAnalysis.ImagingBankSpecificTest = ivAnalysis.parseNumField(record[4:5])
	// 06-06
	ivAnalysis.PartialImage = ivAnalysis.parseNumField(record[5:6])
	// 07-07
	ivAnalysis.ExcessiveImageSkew = ivAnalysis.parseNumField(record[6:7])
	// 08-8
	ivAnalysis.PiggybackImage = ivAnalysis.parseNumField(record[7:8])
	// 09-09
	ivAnalysis.TooLightOrTooDark = ivAnalysis.parseNumField(record[8:9])
	// 10-10
	ivAnalysis.StreaksAndOrBands = ivAnalysis.parseNumField(record[9:10])
	// 11-11
	ivAnalysis.BelowMinimumImageSize = ivAnalysis.parseNumField(record[10:11])
	// 12-12
	ivAnalysis.ExceedsMaximumImageSize = ivAnalysis.parseNumField(record[11:12])
	// 13-25
	ivAnalysis.reserved = "             "
	// 26-26
	ivAnalysis.ImageEnabledPOD = ivAnalysis.parseNumField(record[25:26])
	// 27-27
	ivAnalysis.SourceDocumentBad = ivAnalysis.parseNumField(record[26:27])
	// 28-28
	ivAnalysis.DateUsability = ivAnalysis.parseNumField(record[27:28])
	// 29-29
	ivAnalysis.PayeeUsability = ivAnalysis.parseNumField(record[28:29])
	// 30-30
	ivAnalysis.ConvenienceAmountUsability = ivAnalysis.parseNumField(record[29:30])
	// 31-31
	ivAnalysis.AmountInWordsUsability = ivAnalysis.parseNumField(record[30:31])
	// 32-32
	ivAnalysis.SignatureUsability = ivAnalysis.parseNumField(record[31:32])
	// 33-33
	ivAnalysis.PayorNameAddressUsability = ivAnalysis.parseNumField(record[32:33])
	// 34-34
	ivAnalysis.MICRLineUsability = ivAnalysis.parseNumField(record[33:34])
	// 35-35
	ivAnalysis.MemoLineUsability = ivAnalysis.parseNumField(record[34:35])
	// 36-36
	ivAnalysis.PayorBankNameAddressUsability = ivAnalysis.parseNumField(record[35:36])
	// 37-37
	ivAnalysis.PayeeEndorsementUsability = ivAnalysis.parseNumField(record[36:37])
	// 38-38
	ivAnalysis.BOFDEndorsementUsability = ivAnalysis.parseNumField(record[37:38])
	// 39-39
	ivAnalysis.TransitEndorsementUsability = ivAnalysis.parseNumField(record[38:39])
	// 40-45
	ivAnalysis.reservedTwo = "      "
	// 46-65
	ivAnalysis.UserField = ivAnalysis.parseStringField(record[45:65])
	// 66-80
	ivAnalysis.reservedThree = "               "
}

func (ivAnalysis *ImageViewAnalysis) UnmarshalJSON(data []byte) error {
	type Alias ImageViewAnalysis
	aux := struct {
		*Alias
	}{
		(*Alias)(ivAnalysis),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ivAnalysis.setRecordType()
	return nil
}

// String writes the ImageViewAnalysis struct to a string.
func (ivAnalysis *ImageViewAnalysis) String() string {
	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(ivAnalysis.recordType)
	buf.WriteString(ivAnalysis.GlobalImageQualityField())
	buf.WriteString(ivAnalysis.GlobalImageUsabilityField())
	buf.WriteString(ivAnalysis.ImagingBankSpecificTestField())
	buf.WriteString(ivAnalysis.PartialImageField())
	buf.WriteString(ivAnalysis.ExcessiveImageSkewField())
	buf.WriteString(ivAnalysis.PiggybackImageField())
	buf.WriteString(ivAnalysis.TooLightOrTooDarkField())
	buf.WriteString(ivAnalysis.StreaksAndOrBandsField())
	buf.WriteString(ivAnalysis.BelowMinimumImageSizeField())
	buf.WriteString(ivAnalysis.ExceedsMaximumImageSizeField())
	buf.WriteString(ivAnalysis.reservedField())
	buf.WriteString(ivAnalysis.ImageEnabledPODField())
	buf.WriteString(ivAnalysis.SourceDocumentBadField())
	buf.WriteString(ivAnalysis.DateUsabilityField())
	buf.WriteString(ivAnalysis.PayeeUsabilityField())
	buf.WriteString(ivAnalysis.ConvenienceAmountUsabilityField())
	buf.WriteString(ivAnalysis.AmountInWordsUsabilityField())
	buf.WriteString(ivAnalysis.SignatureUsabilityField())
	buf.WriteString(ivAnalysis.PayorNameAddressUsabilityField())
	buf.WriteString(ivAnalysis.MICRLineUsabilityField())
	buf.WriteString(ivAnalysis.MemoLineUsabilityField())
	buf.WriteString(ivAnalysis.PayorBankNameAddressUsabilityField())
	buf.WriteString(ivAnalysis.PayeeEndorsementUsabilityField())
	buf.WriteString(ivAnalysis.BOFDEndorsementUsabilityField())
	buf.WriteString(ivAnalysis.TransitEndorsementUsabilityField())
	buf.WriteString(ivAnalysis.reservedTwoField())
	buf.WriteString(ivAnalysis.UserFieldField())
	buf.WriteString(ivAnalysis.reservedThreeField())
	return buf.String()
}

// Validate performs ImageCashLetterformat rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (ivAnalysis *ImageViewAnalysis) Validate() error {
	if err := ivAnalysis.fieldInclusion(); err != nil {
		return err
	}
	if ivAnalysis.recordType != "54" {
		msg := fmt.Sprintf(msgRecordType, 54)
		return &FieldError{FieldName: "recordType", Value: ivAnalysis.recordType, Msg: msg}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.GlobalImageQualityField()); err != nil {
		return &FieldError{FieldName: "GlobalImageQuality",
			Value: ivAnalysis.GlobalImageQualityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.GlobalImageUsabilityField()); err != nil {
		return &FieldError{FieldName: "GlobalImageUsability",
			Value: ivAnalysis.GlobalImageUsabilityField(), Msg: err.Error()}
	}

	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.ImagingBankSpecificTestField()); err != nil {
		return &FieldError{FieldName: "ImagingBankSpecificTest",
			Value: ivAnalysis.ImagingBankSpecificTestField(), Msg: err.Error()}
	}
	if err := ivAnalysis.validateConditionalFields(); err != nil {
		return err
	}
	if err := ivAnalysis.isAlphanumericSpecial(ivAnalysis.UserField); err != nil {
		return &FieldError{FieldName: "UserField", Value: ivAnalysis.UserField, Msg: err.Error()}
	}
	return nil
}

// validateConditionalFields makes calls to validate Image View Analysis conditional fields
func (ivAnalysis *ImageViewAnalysis) validateConditionalFields() error {
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.PartialImageField()); err != nil {
		return &FieldError{FieldName: "PartialImage",
			Value: ivAnalysis.PartialImageField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.ExcessiveImageSkewField()); err != nil {
		return &FieldError{FieldName: "ExcessiveImageSkew",
			Value: ivAnalysis.ExcessiveImageSkewField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.PiggybackImageField()); err != nil {
		return &FieldError{FieldName: "PiggybackImage",
			Value: ivAnalysis.PiggybackImageField(), Msg: err.Error()}

	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.TooLightOrTooDarkField()); err != nil {
		return &FieldError{FieldName: "TooLightOrTooDark",
			Value: ivAnalysis.TooLightOrTooDarkField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.StreaksAndOrBandsField()); err != nil {
		return &FieldError{FieldName: "StreaksAndOrBands",
			Value: ivAnalysis.StreaksAndOrBandsField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.BelowMinimumImageSizeField()); err != nil {
		return &FieldError{FieldName: "BelowMinimumImageSize",
			Value: ivAnalysis.BelowMinimumImageSizeField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.ExceedsMaximumImageSizeField()); err != nil {
		return &FieldError{FieldName: "ExceedsMaximumImageSize",
			Value: ivAnalysis.ExceedsMaximumImageSizeField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.ImageEnabledPODField()); err != nil {
		return &FieldError{FieldName: "ImageEnabledPOD",
			Value: ivAnalysis.ImageEnabledPODField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.SourceDocumentBadField()); err != nil {
		return &FieldError{FieldName: "SourceDocumentBad",
			Value: ivAnalysis.SourceDocumentBadField(), Msg: err.Error()}
	}
	if err := ivAnalysis.validateUsabilityFields(); err != nil {
		return err
	}
	return nil
}

// validateUsabilityFields makes calls to validate Image View Analysis usability fields
func (ivAnalysis *ImageViewAnalysis) validateUsabilityFields() error {
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.DateUsabilityField()); err != nil {
		return &FieldError{FieldName: "DateUsability",
			Value: ivAnalysis.DateUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.PayeeUsabilityField()); err != nil {
		return &FieldError{FieldName: "PayeeUsability",
			Value: ivAnalysis.PayeeUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.ConvenienceAmountUsabilityField()); err != nil {
		return &FieldError{FieldName: "ConvenienceAmountUsability",
			Value: ivAnalysis.ConvenienceAmountUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.AmountInWordsUsabilityField()); err != nil {
		return &FieldError{FieldName: "AmountInWordsUsability",
			Value: ivAnalysis.AmountInWordsUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.SignatureUsabilityField()); err != nil {
		return &FieldError{FieldName: "SignatureUsability",
			Value: ivAnalysis.SignatureUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.PayorNameAddressUsabilityField()); err != nil {
		return &FieldError{FieldName: "PayorNameAddressUsability",
			Value: ivAnalysis.PayorNameAddressUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.MICRLineUsabilityField()); err != nil {
		return &FieldError{FieldName: "MICRLineUsability",
			Value: ivAnalysis.MICRLineUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.MemoLineUsabilityField()); err != nil {
		return &FieldError{FieldName: "MemoLineUsability",
			Value: ivAnalysis.MemoLineUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.PayorBankNameAddressUsabilityField()); err != nil {
		return &FieldError{FieldName: "PayorBankNameAddressUsability",
			Value: ivAnalysis.PayorBankNameAddressUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.PayeeEndorsementUsabilityField()); err != nil {
		return &FieldError{FieldName: "PayeeEndorsementUsability",
			Value: ivAnalysis.PayeeEndorsementUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.BOFDEndorsementUsabilityField()); err != nil {
		return &FieldError{FieldName: "BOFDEndorsementUsability",
			Value: ivAnalysis.BOFDEndorsementUsabilityField(), Msg: err.Error()}
	}
	if err := ivAnalysis.isImageViewAnalysisValid(ivAnalysis.TransitEndorsementUsabilityField()); err != nil {
		return &FieldError{FieldName: "TransitEndorsementUsability",
			Value: ivAnalysis.TransitEndorsementUsabilityField(), Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (ivAnalysis *ImageViewAnalysis) fieldInclusion() error {
	if ivAnalysis.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: ivAnalysis.recordType,
			Msg:   msgFieldInclusion + ", did you use ImageViewAnalysis()?"}
	}
	return nil
}

// GlobalImageQualityField gets a string of the GlobalImageQuality field
func (ivAnalysis *ImageViewAnalysis) GlobalImageQualityField() string {
	return ivAnalysis.numericField(ivAnalysis.GlobalImageQuality, 1)
}

// GlobalImageUsabilityField gets a string of the GlobalImageUsability field
func (ivAnalysis *ImageViewAnalysis) GlobalImageUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.GlobalImageUsability, 1)
}

// ImagingBankSpecificTestField gets a string of the ImagingBankSpecificTest field
func (ivAnalysis *ImageViewAnalysis) ImagingBankSpecificTestField() string {
	return ivAnalysis.numericField(ivAnalysis.ImagingBankSpecificTest, 1)
}

// PartialImageField gets a string of the PartialImage field
func (ivAnalysis *ImageViewAnalysis) PartialImageField() string {
	return ivAnalysis.numericField(ivAnalysis.PartialImage, 1)
}

// ExcessiveImageSkewField gets a string of the ExcessiveImageSkew field
func (ivAnalysis *ImageViewAnalysis) ExcessiveImageSkewField() string {
	return ivAnalysis.numericField(ivAnalysis.ExcessiveImageSkew, 1)
}

// PiggybackImageField gets a string of the PiggybackImage field
func (ivAnalysis *ImageViewAnalysis) PiggybackImageField() string {
	return ivAnalysis.numericField(ivAnalysis.PiggybackImage, 1)
}

// TooLightOrTooDarkField gets a string of the TooLightOrTooDark field
func (ivAnalysis *ImageViewAnalysis) TooLightOrTooDarkField() string {
	return ivAnalysis.numericField(ivAnalysis.TooLightOrTooDark, 1)
}

// StreaksAndOrBandsField gets a string of the StreaksAndOrBands field
func (ivAnalysis *ImageViewAnalysis) StreaksAndOrBandsField() string {
	return ivAnalysis.numericField(ivAnalysis.StreaksAndOrBands, 1)
}

// BelowMinimumImageSizeField gets a string of the BelowMinimumImageSize field
func (ivAnalysis *ImageViewAnalysis) BelowMinimumImageSizeField() string {
	return ivAnalysis.numericField(ivAnalysis.BelowMinimumImageSize, 1)
}

// ExceedsMaximumImageSizeField gets a string of the ExceedsMaximumImageSize field
func (ivAnalysis *ImageViewAnalysis) ExceedsMaximumImageSizeField() string {
	return ivAnalysis.numericField(ivAnalysis.ExceedsMaximumImageSize, 1)
}

// reservedField gets the reserved field
func (ivAnalysis *ImageViewAnalysis) reservedField() string {
	return ivAnalysis.alphaField(ivAnalysis.reserved, 13)
}

// ImageEnabledPODField gets a string of the ImageEnabledPOD field
func (ivAnalysis *ImageViewAnalysis) ImageEnabledPODField() string {
	return ivAnalysis.numericField(ivAnalysis.ImageEnabledPOD, 1)
}

// SourceDocumentBadField gets a string of the SourceDocumentBad field
func (ivAnalysis *ImageViewAnalysis) SourceDocumentBadField() string {
	return ivAnalysis.numericField(ivAnalysis.SourceDocumentBad, 1)
}

// DateUsabilityField gets a string of the DateUsability field
func (ivAnalysis *ImageViewAnalysis) DateUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.DateUsability, 1)
}

// PayeeUsabilityField gets a string of the PayeeUsability field
func (ivAnalysis *ImageViewAnalysis) PayeeUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.PayeeUsability, 1)
}

// ConvenienceAmountUsabilityField gets a string of the ConvenienceAmountUsability field
func (ivAnalysis *ImageViewAnalysis) ConvenienceAmountUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.ConvenienceAmountUsability, 1)
}

// AmountInWordsUsabilityField gets a string of the AmountInWordsUsability field
func (ivAnalysis *ImageViewAnalysis) AmountInWordsUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.AmountInWordsUsability, 1)
}

// SignatureUsabilityField gets a string of the SignatureUsability  field
func (ivAnalysis *ImageViewAnalysis) SignatureUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.SignatureUsability, 1)
}

// PayorNameAddressUsabilityField gets a string of the PayorNameAddressUsability field
func (ivAnalysis *ImageViewAnalysis) PayorNameAddressUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.PayorNameAddressUsability, 1)
}

// MICRLineUsabilityField gets a string of the MICRLineUsability field
func (ivAnalysis *ImageViewAnalysis) MICRLineUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.MICRLineUsability, 1)
}

// MemoLineUsabilityField gets a string of the MemoLineUsability field
func (ivAnalysis *ImageViewAnalysis) MemoLineUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.MemoLineUsability, 1)
}

// PayorBankNameAddressUsabilityField gets a string of the PayorBankNameAddressUsability field
func (ivAnalysis *ImageViewAnalysis) PayorBankNameAddressUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.PayorBankNameAddressUsability, 1)
}

// PayeeEndorsementUsabilityField gets a string of the PayeeEndorsementUsability field
func (ivAnalysis *ImageViewAnalysis) PayeeEndorsementUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.PayeeEndorsementUsability, 1)
}

// BOFDEndorsementUsabilityField gets a string of the BOFDEndorsementUsability field
func (ivAnalysis *ImageViewAnalysis) BOFDEndorsementUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.BOFDEndorsementUsability, 1)
}

// TransitEndorsementUsabilityField gets a string of the TransitEndorsementUsability field
func (ivAnalysis *ImageViewAnalysis) TransitEndorsementUsabilityField() string {
	return ivAnalysis.numericField(ivAnalysis.TransitEndorsementUsability, 1)
}

// reservedTwoField gets the reservedTwo field
func (ivAnalysis *ImageViewAnalysis) reservedTwoField() string {
	return ivAnalysis.alphaField(ivAnalysis.reservedTwo, 6)
}

// UserFieldField gets the UserField field
func (ivAnalysis *ImageViewAnalysis) UserFieldField() string {
	return ivAnalysis.alphaField(ivAnalysis.UserField, 20)
}

// reservedThreeField gets the reservedThree field
func (ivAnalysis *ImageViewAnalysis) reservedThreeField() string {
	return ivAnalysis.alphaField(ivAnalysis.reservedThree, 15)
}
