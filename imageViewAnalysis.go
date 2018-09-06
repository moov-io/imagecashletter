// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// ToDo: Handle inserted length field (variable length) Big Endian and Little Endian format

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
}

// NewImageViewAnalysis returns a new ImageViewAnalysis with default values for non exported fields
func NewImageViewAnalysis() *ImageViewAnalysis {
	imageAnalysis := &ImageViewAnalysis{
		recordType: "54",
	}
	return imageAnalysis
}

// Parse takes the input record string and parses the ImageViewAnalysis values

// String writes the ImageViewAnalysis struct to a string.

// Validate performs X9 format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.

// Get properties
