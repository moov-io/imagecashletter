# BundleControl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | BundleControl ID | [optional] 
**BundleItemsCount** | **int32** | BundleItemsCount identifies the total number of items within the bundle. | 
**BundleTotalAmount** | **int32** | BundleTotalAmount identifies the total amount of item amounts within the bundle. | 
**MicrValidTotalAmount** | **int32** | MICRValidTotalAmount identifies the total amount of all CheckDetail Records within the bundle which contains 1 in the MICRValidIndicator. | [optional] 
**BundleImagesCount** | **int32** | BundleImagesCount identifies the total number of Image ViewDetail Records  within the bundle. | [optional] 
**UserField** | **string** | UserField identifies a field used at the discretion of users of the standard. | [optional] 
**CreditTotalIndicator** | **string** | CreditTotalIndicator is a code that indicates whether Credits Items are included in the totals. If so they  they will be included in this record’s: Items Within Bundle Count, Bundle Total Amount and, Images Within Bundle Count * &#x60; &#x60; - No Credit Items * &#x60;0&#x60; - Credit Items are not included in totals * &#x60;1&#x60; - Credit Items are included in totals  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


