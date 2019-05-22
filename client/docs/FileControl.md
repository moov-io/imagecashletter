# FileControl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | FileControl ID | [optional] 
**CashLetterCount** | **int32** | Identifies the total number of cash letters within the file. | 
**TotalRecordCount** | **int32** | Identifies the total number of records of all types sent in the file, including the FileControl. | 
**TotalItemCount** | **int32** | Identifies the total number of Items sent within the file. | 
**FileTotalAmount** | **int32** | Identifies the total Item amount of the complete file. | 
**ImmediateOriginContactName** | **string** | Identifies contact at the institution that creates the file. | [optional] 
**ImmediateOriginContactPhoneNumber** | **string** | Identifies the phone number of the contact at the institution that creates the file. | [optional] 
**CreditTotalIndicator** | **string** | Is a code that indicates whether Credits Items are included in this record’s totals. If so they will be included in TotalItemCount and FileTotalAmount. TotalRecordCount includes all records of all types regardless of the value of this field. * &#x60; &#x60; - No Credit Items * &#x60;0&#x60; - Credit Items are not included in totals * &#x60;1&#x60; - Credit Items are included in totals  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


