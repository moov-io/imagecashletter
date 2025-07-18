/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing, and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ImageViewData struct for ImageViewData
type ImageViewData struct {
	// ImageViewData ID
	Id string `json:"id,omitempty"`
	// ECEInstitutionRoutingNumber contains the routing and transit number of the institution that creates the bundle header.  This number is imported from the Bundle Header Record (Clause 9.4) associated with the image view conveyed in this Image View Data Property.
	EceInstitutionRoutingNumber string `json:"eceInstitutionRoutingNumber,omitempty"`
	// BundleBusinessDate is the business date of the bundle.
	BundleBusinessDate time.Time `json:"bundleBusinessDate,omitempty"`
	// CycleNumber is a code assigned by the institution that creates the bundle.  Denotes the cycle under which the bundle is created.
	CycleNumber string `json:"cycleNumber,omitempty"`
	// ECEInstitutionItemSequenceNumber is a number assigned by the institution that creates the Check or Return.  This number is imported from the Check.ECEInstitutionItemSequenceNumber or Return.ECEInstitutionItemSequenceNumber associated with the image view conveyed in this Image View Data Record. The ECE institution must construct the sequence number to guarantee uniqueness for a given routing number, business day, and cycle number. Must contain a numeric value.
	EceInstitutionItemSequenceNumber string `json:"eceInstitutionItemSequenceNumber,omitempty"`
	// SecurityOriginatorName is a unique name that creates the Digital Signature for data to be exchanged. Shall be present only under clearing arrangements and when ImageViewDetail.DigitalSignatureIndicator is 1 Shall not be present when ImageViewDetail.ImageIndicator is 0.
	SecurityOriginatorName string `json:"securityOriginatorName,omitempty"`
	// SecurityAuthenticatorName is the unique name that performs authentication on received data. Shall be present only under clearing arrangements and when ImageViewDetail.DigitalSignatureIndicator is 1 Shall not be present when ImageViewDetail.ImageIndicator is 0.
	SecurityAuthenticatorName string `json:"securityAuthenticatorName,omitempty"`
	// SecurityKeyName is a name or character sequence used by the signer (originator) to communicate a key identifierto the recipient (authenticator) so the recipient can obtain the key needed to validate the signature. The name is typically used as an identifier related to the key pair used to sign the image. The name is mutually known to the security originator and the security authenticator and is unique to this relationship. Shall be present only under clearing arrangements and when ImageViewDetail.DigitalSignatureIndicator is 1 Shall not be present when ImageViewDetail.ImageIndicator is 0.
	SecurityKeyName string `json:"securityKeyName,omitempty"`
	// ClippingOrigin is a code that defines the corner of the conveyed image view that is taken as the reference point for the clipping coordinates. Top, bottom, left, and right references apply to a view that presents a visually correct orientation. When clipping information is present, the nature of the Area of Interest defined by the clipping rectangle is determined by the value of the ImageViewDetail.ViewDescriptor. Primary front and rear views shall only have a Defined Value of 0.  Can be blank.  * `0` - Clipping information is not present–full view present * `1` - Clipping origin is top left corner of image view * `2` - Clipping origin is top right corner of image view * `3` - Clipping origin is bottom right corner of image view * `4` - Clipping origin is bottom left corner of image view
	ClippingOrigin int32 `json:"clippingOrigin,omitempty"`
	// ClippingCoordinateH1 is a number that represents the horizontal offset in pixels from the clipping origin to the nearest vertical side of the clipping rectangle. The clipping coordinates (h1, h2, v1, v2) convey the clipping rectangle’s offsets in both horizontal (h) and vertical (v) directions. The offset values collectively establish the boundary sides of the clipping rectangle. Pixels on the boundary of the clipping rectangle are included in the selected array of pixels. That is, the first pixel of the selected array is at offset (h1, v1) and the last pixel of the selected array is at offset (h2, v2). The corner pixel at the origin of the image view is assumed to have the offset value (0, 0). Shall be present if Image View Data.ClippingOrigin is present and non-zero. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000–9999
	ClippingCoordinateH1 string `json:"clippingCoordinateH1,omitempty"`
	// ClippingCoordinateH2 is a number that represents the horizontal offset in pixels from the clipping origin to the furthermost vertical side of the clipping rectangle. Shall be present if Image View Data.ClippingOrigin is present and non-zero. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000–9999
	ClippingCoordinateH2 string `json:"clippingCoordinateH2,omitempty"`
	// ClippingCoordinateV1 is a number that represents the vertical offset in pixels from the clipping origin to the nearest horizontal side of the clipping rectangle. Shall be present if Image View Data.ClippingOrigin is present and non-zero. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000–9999
	ClippingCoordinateV1 string `json:"clippingCoordinateV1,omitempty"`
	// ClippingCoordinateV2 is a number that represents the vertical offset in pixels from the clipping origin to the furthermost horizontal side of the clipping rectangle. Shall be present if Image View Data.ClippingOrigin is present and non-zero. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000–9999
	ClippingCoordinateV2 string `json:"clippingCoordinateV2,omitempty"`
	// LengthImageReferenceKey is the number of characters in the ImageViewData.ImageReferenceKey. Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid values - 0000 ImageReferenceKey is not present 0001–9999  Valid when ImageReferenceKey is present
	LengthImageReferenceKey string `json:"lengthImageReferenceKey,omitempty"`
	// ImageReferenceKey is assigned by the ECE institution that creates the CheckDetail or Return, and the related Image View Records. This designator, when used, shall uniquely identify the item image to the ECE institution. This designator is a special key with significance to the creating institution. It is intended to be used to locate within an archive the unique image associated with the item. The designator could be a full access path and name that would allow direct look up and access to the image, for example a URL. This shall match CheckDetailAddendumB.ImageReferenceKey, or ReturnAddendumCImageReferenceKey Record, if used. Valid size - 0 – 9999
	ImageReferenceKey string `json:"imageReferenceKey,omitempty"`
	// LengthDigitalSignature is the number of bytes in the Image View Data.DigitalSignature. Shall not be present when ImageViewDetail.ImageIndicator is 0.
	LengthDigitalSignature string `json:"lengthDigitalSignature,omitempty"`
	// DigitalSignature is created by applying the cryptographic algorithm and private/secret key against the data to be protected. The Digital Signature provides user authentication and data integrity. Shall be present only under clearing arrangements and when ImageViewDetail.DigitalSignatureIndicator is 1 Shall not be present when ImageViewDetail.ImageIndicator is 0. Valid size - 0-99999
	DigitalSignature string `json:"digitalSignature,omitempty"`
	// LengthImageData is the number of bytes in the ImageViewData.ImageData. Shall be present when ImageViewDetail.ImageIndicator is NOT 0 Valid values - 0000001–99999999
	LengthImageData string `json:"lengthImageData,omitempty"`
	// ImageData contains the image view. The Image Data generally consists of an image header and the image raster data. The image header provides information that is required to interpret the image raster data. The image raster data contains the scanned image of the physical item in raster (line by line) format. Each scan line comprises a set of concatenated pixels. The image comprises a set of scan lines. The image raster data is typically compressed to reduce the number of bytes needed to transmit and store the image. The header/image format type is defined by the ImageViewDetail.ImageViewFormatIndicator. The syntax and semantics of the image header/image format are understood by referring to the appropriate image format specification. The compression scheme used to compress the image raster data is specified in the ImageViewCompressionAlgorithmIdentifier and in the image header portion of the Image Data or by association with the selected image format. The data may be provided in standard Base64 encoding and will be decoded on file generation. Shall be present when ImageViewDetail.ImageIndicator Record is NOT 0. Valid size - 0-9999999
	ImageData string `json:"imageData,omitempty"`
}
