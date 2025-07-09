# BundleHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | BundleHeader ID | [optional] 
**CollectionTypeIndicator** | **string** | A code that identifies the type of bundle. It is the same value as the CollectionTypeIndicator in the CashLetterHeader within which the bundle is contained, unless the CollectionTypeIndicator in the CashLetterHeader is 99.  * &#x60;00&#x60; - Preliminary Forward Information * &#x60;01&#x60; - Forward Presentment * &#x60;02&#x60; - Forward Presentment - Same-Day Settlement * &#x60;03&#x60; - Return * &#x60;04&#x60; - Return Notification * &#x60;05&#x60; - Preliminary Return Notification * &#x60;06&#x60; - Final Return Notification  | [optional] 
**DestinationRoutingNumber** | **string** | DestinationRoutingNumber contains the routing and transit number of the institution that receives and processes the cash letter or the bundle. | [optional] 
**EceInstitutionRoutingNumber** | **string** | ECEInstitutionRoutingNumber contains the routing and transit number of the institution that that creates the bundle header. | [optional] 
**BundleBusinessDate** | [**time.Time**](time.Time.md) | BundleBusinessDate is the business date of the bundle. | [optional] 
**BundleCreationDate** | [**time.Time**](time.Time.md) | BundleCreationDate is the date that the bundle is created. | [optional] 
**BundleID** | **string** | BundleID is a number that identifies the bundle, assigned by the institution that creates the bundle. | [optional] 
**BundleSequenceNumber** | **string** | BundleSequenceNumber is a number assigned by the institution that creates the bundle. Usually denotes the relative position of the bundle within the cash letter. | [optional] 
**CycleNumber** | **string** | CycleNumber is a code assigned by the institution that creates the bundle.  Denotes the cycle under which the bundle is created. | [optional] 
**ReturnLocationRoutingNumber** | **string** | ReturnLocationRoutingNumber is a bank routing number used by some processors. This will be blank in the resulting file if it is empty. | [optional] 
**UserField** | **string** | UserField identifies a field used at the discretion of users of the standard. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


