# Returns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Return ID | [optional] 
**PayorBankRoutingNumber** | **string** | PayorBankRoutingNumber identifies the institution by or through which the item is payable. Must be a valid routing and transit number issued by the ABA’s Routing Number Registrar. Shall represent the first 8 digits of a 9-digit routing number or 8 numeric digits of a 4 dash 4 routing number. A valid routing number consists of 2 fields: the eight-digit Payor Bank Routing Number and the one-digit Payor Bank Routing Number Check Digit.  | [optional] 
**PayorBankCheckDigit** | **string** | PayorBankCheckDigit identifies the routing number check digit.  The combination of Payor Bank Routing Number and Payor Bank Routing Number Check Digit must be a mod-checked routing number with a valid check digit.  | [optional] 
**OnUs** | **string** | OnUs identifies data specified by the payor bank. On-Us data usually consists of the payor’s account number, a serial number or transaction code, or both. | [optional] 
**ItemAmount** | **int32** | Amount identifies the amount of the check.  All amounts fields have two implied decimal points. e.g., 100000 is $1,000.00. | [optional] 
**ReturnReason** | **string** | ReturnReason is a code that indicates the reason for non-payment. | [optional] 
**AddendumCount** | **int32** | AddendumCount is a number of Check Detail Record Addenda to follow. This represents the number of CheckDetailAddendumA, CheckDetailAddendumB, and CheckDetailAddendumC types. It matches the total number of addendum records associated with this item. The standard supports up to 99 addendum records. | [optional] 
**DocumentationTypeIndicator** | **string** | DocumentationTypeIndicator identifies a code that indicates the type of documentation that supports the check record.  This field is superseded by the Cash Letter Documentation Type Indicator in the Cash Letter Header Record for all Defined Values except ‘Z’ Not Same Type. In the case of Defined Value of ‘Z’, the Documentation Type Indicator in this record takes precedent.  Shall be present when Cash Letter Documentation Type Indicator in the Cash Letter Header Record is Defined Value of ‘Z’.  * &#x60;A&#x60; - No image provided, paper provided separately * &#x60;B&#x60; - No image provided, paper provided separately, image upon request * &#x60;C&#x60; - Image provided separately, no paper provided * &#x60;D&#x60; - Image provided separately, no paper provided, image upon request * &#x60;E&#x60; - Image and paper provided separately * &#x60;F&#x60; - Image and paper provided separately, image upon request * &#x60;G&#x60; - Image included, no paper provided * &#x60;H&#x60; - Image included, no paper provided, image upon request * &#x60;I&#x60; - Image included, paper provided separately * &#x60;J&#x60; - Image included, paper provided separately, image upon request * &#x60;K&#x60; - No image provided, no paper provided * &#x60;L&#x60; - No image provided, no paper provided, image upon request * &#x60;M&#x60; - No image provided, Electronic Check provided separately  | [optional] 
**ForwardBundleDate** | [**time.Time**](time.Time.md) | ForwardBundleDate represents for electronic check exchange items, the year, month, and day that designate the business date of the original forward bundle. This data is transferred from the BundleHeader BundleBusinessDate.  For items presented in paper cash letters, the year, month, and day that the cash letter was created. | [optional] 
**EceInstitutionItemSequenceNumber** | **string** | ECEInstitutionItemSequenceNumber identifies a number assigned by the institution that creates the CheckDetail. Field must contain a numeric value. It cannot be all blanks. | [optional] 
**ExternalProcessingCode** | **string** | ExternalProcessingCode identifies a code used for special purposes as authorized by the Accredited Standards Committee X9. Also known as Position 44. | [optional] 
**ReturnNotificationIndicator** | **string** | ReturnNotificationIndicator is a code that identifies the type of notification. The CashLetterHeader.CollectionTypeIndicator and BundleHeader.CollectionTypeIndicator equalling &#x60;05&#x60; or &#x60;06&#x60; takes precedence over this field.  * &#x60;1&#x60; - Preliminary notification * &#x60;2&#x60; - Final notification  | [optional] 
**ArchiveTypeIndicator** | **string** | ArchiveTypeIndicator is a code that indicates the type of archive that supports this Check. Access method, availability, and time frames shall be defined by clearing arrangements. * &#x60;A&#x60; - Microfilm * &#x60;B&#x60; - Image * &#x60;C&#x60; - Paper * &#x60;D&#x60; - Microfilm and image * &#x60;E&#x60; - Microfilm and paper * &#x60;F&#x60; - Image and paper * &#x60;G&#x60; - Microfilm, image, and paper * &#x60;H&#x60; - Electronic Check Instrument * &#x60;I&#x60; - None  | [optional] 
**TimesReturned** | **int32** | TimesReturned is a code used to indicate the number of times the paying bank has returned this item.  * &#x60;0&#x60; - The item has been returned an unknown number of times * &#x60;1&#x60; - The item has been returned once * &#x60;2&#x60; - The item has been returned twice * &#x60;3&#x60; - The item has been returned three times  | [optional] 
**ReturnDetailAddendumA** | [**[]ReturnDetailAddendumA**](ReturnDetailAddendumA.md) |  | [optional] 
**ReturnDetailAddendumB** | [**[]ReturnDetailAddendumB**](ReturnDetailAddendumB.md) |  | [optional] 
**ReturnDetailAddendumC** | [**[]ReturnDetailAddendumC**](ReturnDetailAddendumC.md) |  | [optional] 
**ReturnDetailAddendumD** | [**[]ReturnDetailAddendumD**](ReturnDetailAddendumD.md) |  | [optional] 
**ImageViewDetail** | [**[]ImageViewDetail**](ImageViewDetail.md) |  | [optional] 
**ImageViewData** | [**[]ImageViewData**](ImageViewData.md) |  | [optional] 
**ImageViewAnalysis** | [**[]ImageViewAnalysis**](ImageViewAnalysis.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


