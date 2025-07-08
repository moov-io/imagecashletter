# CheckDetailAddendumC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | CheckDetailAddendumC ID | [optional] 
**RecordNumber** | **int32** | RecordNumber is a number representing the order in which each CheckDetailAddendumC was created. CheckDetailAddendumC shall be in sequential order starting with 1. | 
**EndorsingBankRoutingNumber** | **string** | EndorsingBankRoutingNumber is a valid routing and transit number indicating the bank that endorsed the check. | 
**BofdEndorsementBusinessDate** | [**time.Time**](time.Time.md) | BOFDEndorsementBusinessDate is the date of endorsement. | 
**EndorsingBankSequenceNumber** | **string** | EndorsingItemSequenceNumber is a number that identifies the item at the endorsing bank. | [optional] 
**TruncationIndicator** | **string** | TruncationIndicator identifies if the institution truncated the original check item. | 
**EndorsingBankConversionIndicator** | **string** | EndorsingBankConversionIndicator is a code that indicates the conversion within the processing institution between original paper check, image, and IRD. The indicator is specific to the action of the institution identified in the EndorsingBankRoutingNumber.  * &#x60;0&#x60; - Did not convert physical document * &#x60;1&#x60; - Original paper converted to IRD * &#x60;2&#x60; - Original paper converted to image * &#x60;3&#x60; - IRD converted to another IRD * &#x60;4&#x60; - IRD converted to image of IRD * &#x60;5&#x60; - Image converted to an IRD * &#x60;6&#x60; - Image converted to another image (e.g., transcoded) * &#x60;7&#x60; - Did not convert image (e.g., same as source) * &#x60;8&#x60; - Undetermined  | [optional] 
**EndorsingBankCorrectionIndicator** | **int32** | EndorsingBankCorrectionIndicator identifies whether and how the MICR line of this item was repaired by the creator of this CheckDetailAddendumC Record for fields other than Payor Bank Routing Number and Amount.  * &#x60;0&#x60; - No Repair * &#x60;1&#x60; - Repaired (form of repair unknown) * &#x60;2&#x60; - Repaired without Operator intervention * &#x60;3&#x60; - Repaired with Operator intervention * &#x60;4&#x60; - Undetermined if repair has been done or not  | [optional] 
**ReturnReason** | **string** | ReturnReason is a code that indicates the reason for non-payment. | [optional] 
**UserField** | **string** | UserField identifies a field used at the discretion of users of the standard. | [optional] 
**EndorsingBankIdentifier** | **int32** | * &#x60;0&#x60; - Depository Bank (BOFD) - this value is used when the CheckDetailAddendumC Record reflects the Return * &#x60;Processing Bank in lieu of BOFD. * &#x60;1&#x60; - Other Collecting Bank * &#x60;2&#x60; - Other Returning Bank * &#x60;3&#x60; - Payor Bank  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


