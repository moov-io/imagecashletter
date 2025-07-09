# CashLetterHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | CashLetterHeader ID | [optional] 
**CollectionTypeIndicator** | **string** | CollectionTypeIndicator is a code that identifies the type of cash letter.  * &#x60;00&#x60; - Preliminary Forward Information * &#x60;01&#x60; - Forward Presentment * &#x60;02&#x60; - Forward Presentment - Same-Day Settlement * &#x60;03&#x60; - Return * &#x60;04&#x60; - Return Notification * &#x60;05&#x60; - Preliminary Return Notification * &#x60;06&#x60; - Final Return Notification * &#x60;20&#x60; - No Detail * &#x60;99&#x60; - Bundles not the same collection type. Use of the value is only allowed by clearing arrangement.  | [optional] 
**DestinationRoutingNumber** | **string** | DestinationRoutingNumber is the routing and transit number of the institution that receives and processes the cash letter or the bundle. | [optional] 
**EceInstitutionRoutingNumber** | **string** | ECEInstitutionRoutingNumber is the routing and transit number of the institution that creates the Cash Letter Header record. | [optional] 
**CashLetterBusinessDate** | [**time.Time**](time.Time.md) | cashLetterBusinessDate is the business date of the cash letter. | [optional] 
**CashLetterCreationDate** | [**time.Time**](time.Time.md) | cashLetterCreationDate is the date that the cash letter is created. | [optional] 
**CashLetterCreationTime** | [**time.Time**](time.Time.md) | CashLetterCreationTime is the time that the cash letter is created. | [optional] 
**RecordTypeIndicator** | **string** | RecordTypeIndicator is a code that indicates the presence of records or the type of records contained in the cash letter. If an image is associated with any Check or Return, the cash letter must have a RecordTypeIndicator of I or F.  * &#x60;N&#x60; - No electronic check records or image records (Type 2x’s, 3x’s, 5x’s); e.g., an empty cash letter. * &#x60;E&#x60; - Cash letter contains electronic check records with no images (Type 2x’s and 3x’s only). * &#x60;I&#x60; - Cash letter contains electronic check records (Type 2x’s, 3x’s) and image records (Type 5x’s). * &#x60;F&#x60; - Cash letter contains electronic check records (Type 2x’s and 3x’s) and image records (Type 5x’s) that correspond to a previously sent cash letter (i.e., E file).  | [optional] 
**DocumentationTypeIndicator** | **string** | DocumentationTypeIndicator is a code that indicates the type of documentation that supports all check records in the cash letter.  * &#x60;A&#x60; - No image provided, paper provided separately * &#x60;B&#x60; - No image provided, paper provided separately, image upon request * &#x60;C&#x60; - Image provided separately, no paper provided * &#x60;D&#x60; - Image provided separately, no paper provided, image upon request * &#x60;E&#x60; - Image and paper provided separately * &#x60;F&#x60; - Image and paper provided separately, image upon request * &#x60;G&#x60; - Image included, no paper provided * &#x60;H&#x60; - Image included, no paper provided, image upon request * &#x60;I&#x60; - Image included, paper provided separately * &#x60;J&#x60; - Image included, paper provided separately, image upon request * &#x60;K&#x60; - No image provided, no paper provided * &#x60;L&#x60; - No image provided, no paper provided, image upon request * &#x60;M&#x60; - No image provided, Electronic Check provided separately * &#x60;Z&#x60; - Not Same Type–Documentation associated with each item in Cash Letter will be different.  | [optional] 
**CashLetterID** | **string** | CashLetterID uniquely identifies the cash letter. It is assigned by the institution that creates the cash letter and must be unique within a Cash Letter Business Date. | [optional] 
**OriginatorContactName** | **string** | OriginatorContactName is the name of a contact at the institution that creates the cash letter. | [optional] 
**OriginatorContactPhoneNumber** | **string** | OriginatorContactPhoneNumber is the phone number of the contact at the institution that creates the cash letter. | [optional] 
**FedWorkType** | **string** | fedWorkType is any valid code specified by the Federal Reserve Bank. | [optional] 
**ReturnsIndicator** | **string** | ReturnsIndicator identifies type of returns.  * &#x60; &#x60; - Original Message * &#x60;E&#x60; - Administrative - items being returned that are handled by the bank and usually do not directly affect the customer or its account. * &#x60;R&#x60; - Customer–items being returned that directly affect a customer’s account. * &#x60;J&#x60; - Reject Return  | [optional] 
**UserField** | **string** | UserField is a field used at the discretion of users of the standard | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


