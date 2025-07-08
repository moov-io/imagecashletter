# CheckDetailAddendumA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | CheckDetailAddendumA ID | [optional] 
**RecordNumber** | **int32** | RecordNumber is a number representing the order in which each CheckDetailAddendumA was created. CheckDetailAddendumA shall be in sequential order starting with 1. | 
**ReturnLocationRoutingNumber** | **string** | ReturnLocationRoutingNumber is a valid routing and transit number indicating where returns, final return notifications, and preliminary return notifications are sent, usually the BOFD. | 
**BofdEndorsementDate** | [**time.Time**](time.Time.md) | BOFDEndorsementDate is the date of endorsement. | 
**BofdItemSequenceNumber** | **string** | BOFDItemSequenceNumber is a number that identifies the item in the CheckDetailAddendumA. | 
**BofdAccountNumber** | **string** | BOFDAccountNumber is a number that identifies the depository account at the Bank of First Deposit. | [optional] 
**BofdBranchCode** | **string** | BOFDBranchCode is a code that identifies the branch at the Bank of First Deposit. | [optional] 
**PayeeName** | **string** | PayeeName is the name of the payee from the check. | [optional] 
**TruncationIndicator** | **string** | TruncationIndicator identifies if the institution truncated the original check item. | 
**BofdConversionIndicator** | **string** | BOFDConversionIndicator is a code that indicates the conversion within the processing institution between original paper check, image, and IRD. The indicator is specific to the action of the institution that created this record.  * &#x60;0&#x60; - Did not convert physical document * &#x60;1&#x60; - Original paper converted to IRD * &#x60;2&#x60; - Original paper converted to image * &#x60;3&#x60; - IRD converted to another IRD * &#x60;4&#x60; - IRD converted to image of IRD * &#x60;5&#x60; - Image converted to an IRD * &#x60;6&#x60; - Image converted to another image (e.g., transcoded) * &#x60;7&#x60; - Did not convert image (e.g., same as source) * &#x60;8&#x60; - Undetermined  | [optional] 
**BofdCorrectionIndicator** | **int32** | BOFDCorrectionIndicator identifies whether and how the MICR line of this item was repaired by the creator of this CheckDetailAddendumA Record for fields other than Payor Bank Routing Number and Amount. * &#x60;0&#x60; - No Repair * &#x60;1&#x60; - Repaired (form of repair unknown) * &#x60;2&#x60; - Repaired without Operator intervention * &#x60;3&#x60; - Repaired with Operator intervention * &#x60;4&#x60; - Undetermined if repair has been done or not  | [optional] 
**UserField** | **string** | UserField identifies a field used at the discretion of users of the standard. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


