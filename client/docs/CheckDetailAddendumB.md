# CheckDetailAddendumB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | CheckDetailAddendumB ID | [optional] 
**ImageReferenceKeyIndicator** | **int32** | ImageReferenceKeyIndicator identifies whether ImageReferenceKeyLength contains a variable value within the allowable range, or contains a defined value and the content is ItemReferenceKey.  * &#x60;0&#x60; - ImageReferenceKeyIndicator has a Defined Value of 0034 and ImageReferenceKey contains the Image Reference Key. * &#x60;1&#x60;- ImageReferenceKeyIndicator contains a value other than 0034; or ImageReferenceKeyIndicator contains Value 0034, which is not a Defined Value, and the content of ImageReferenceKey has no special significance with regards to an Image Reference Key; or ImageReferenceKeyIndicator is 0000, meaning the ImageReferenceKey is not present.  | [optional] 
**MicrofilmArchiveSequenceNumber** | **string** | microfilmArchiveSequenceNumber is a number that identifies the item in the microfilm archive system; it may be different than the Check.ECEInstitutionItemSequenceNumber and from the ImageReferenceKey. | 
**LengthImageReferenceKey** | **string** | ImageReferenceKeyLength is the number of characters in the ImageReferenceKey.  * &#x60;0034&#x60; - ImageReferenceKey contains the ImageReferenceKey (ImageReferenceKeyIndicator is 0). * &#x60;0000&#x60; - ImageReferenceKey not present (ImageReferenceKeyIndicator is 1). * &#x60;0001&#x60; - 9999: May include Value 0034, and ImageReferenceKey has no special significance to Image Reference Key (ImageReferenceKey is 1).  | [optional] 
**ImageReferenceKey** | **string** | ImageReferenceKey is used to find the image of the item in the image data system.  Size is variable based on lengthImageReferenceKey. The position within the file is variable based on the lengthImageReferenceKey.  | [optional] 
**Description** | **string** | Description describes the transaction.  The position within the file is variable based on the lengthImageReferenceKey. | [optional] 
**UserField** | **string** | UserField identifies a field used at the discretion of users of the standard. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


