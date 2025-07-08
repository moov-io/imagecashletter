# CashLetterControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | CashLetterControl ID | [optional] 
**CashLetterBundleCount** | **int32** | CashLetterBundleCount identifies the total number of bundles within the cash letter. | [optional] 
**CashLetterItemsCount** | **int32** | CashLetterItemsCount identifies the total number of items within the cash letter. | 
**CashLetterTotalAmount** | **int32** | CashLetterTotalAmount identifies the total dollar value of all item amounts within the cash letter. | 
**CashLetterImagesCount** | **int32** | CashLetterImagesCount identifies the total number of ImageViewDetail(s) within the CashLetter. | [optional] 
**EceInstitutionName** | **string** | ECEInstitutionName identifies the short name of the institution that creates the CashLetterControl. | [optional] 
**SettlementDate** | [**time.Time**](time.Time.md) | SettlementDate identifies the date that the institution that creates the cash letter expects settlement. | 
**CreditTotalIndicator** | **int32** | CreditTotalIndicator is a code that indicates whether Credit Items are included in this record’s totals. If so, they will be included in TotalItemCount and FileTotalAmount. TotalRecordCount includes all records of all types regardless of the value of this field. * &#x60; &#x60; - No Credit Items * &#x60;0&#x60; - Credit Items are not included in totals * &#x60;1&#x60; - Credit Items are included in totals  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


