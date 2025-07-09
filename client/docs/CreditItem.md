# CreditItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | CreditItem ID | [optional] 
**AuxiliaryOnUs** | **string** | AuxiliaryOnUs identifies a code used at the discretion of the creating bank. The handling of dashes and spaces shall be determined between the exchange partners. | [optional] 
**ExternalProcessingCode** | **string** | ExternalProcessingCode identifies a code used for special purposes as authorized by the Accredited Standards Committee X9. Also known as Position 44. | [optional] 
**PostingBankRoutingNumber** | **string** | PostingBankRoutingNumber is a routing number assigned by the posting bank to identify this credit. | 
**OnUs** | **string** | OnUs identifies data specified by the payor bank. On-Us data usually consists of the payor’s account number, a serial number or transaction code, or both. | [optional] 
**ItemAmount** | **int32** | Amount identifies the amount of the check.  All amounts fields have two implied decimal points. e.g., 100000 is $1,000.00. | [optional] 
**CreditItemSequenceNumber** | **string** | CreditItemSequenceNumber identifies a number assigned by the institution that creates the CreditItem. | 
**DocumentationTypeIndicator** | **string** | DocumentationTypeIndicator is a code used to indicate the type of documentation that supports this record. Shall be present when Cash Letter Documentation Type Indicator in the Cash Letter Header Record is Defined Value of ‘Z’.  * &#x60;A&#x60; - No image provided, paper provided separately * &#x60;B&#x60; - No image provided, paper provided separately, image upon request * &#x60;C&#x60; - Image provided separately, no paper provided * &#x60;D&#x60; - Image provided separately, no paper provided, image upon request * &#x60;E&#x60; - Image and paper provided separately * &#x60;F&#x60; - Image and paper provided separately, image upon request * &#x60;G&#x60; - Image included, no paper provided * &#x60;H&#x60; - Image included, no paper provided, image upon request * &#x60;I&#x60; - Image included, paper provided separately * &#x60;J&#x60; - Image included, paper provided separately, image upon request * &#x60;K&#x60; - No image provided, no paper provided * &#x60;L&#x60; - No image provided, no paper provided, image upon request  | [optional] 
**AccountTypeCode** | **string** | AccountTypeCode is a code that indicates the type of account to which this CreditItem is associated.  * &#x60;0&#x60; - Unknown * &#x60;1&#x60; - DDA account * &#x60;2&#x60; - General Ledger account * &#x60;3&#x60; - Savings account * &#x60;4&#x60; - Money Market account * &#x60;5&#x60; - Other Account  | [optional] 
**SourceWorkCode** | **string** | SourceWorkCode is a code used to identify the source of the work associated with this CreditItem.  * &#x60;00&#x60; - Unknown * &#x60;01&#x60; - Internal–ATM * &#x60;02&#x60; - Internal–Branch * &#x60;03&#x60; - Internal–Other * &#x60;04&#x60; - External–Bank to Bank (Correspondent) * &#x60;05&#x60; - External–Business to Bank (Customer) * &#x60;06&#x60; - External–Business to Bank Remote Capture * &#x60;07&#x60; - External–Processor to Bank * &#x60;08&#x60; - External–Bank to Processor * &#x60;09&#x60; - Lockbox * &#x60;10&#x60; - International–Internal * &#x60;11&#x60; - International–External * &#x60;21–50&#x60; - User Defined  | [optional] 
**UserField** | **string** | UserField identifies a field used at the discretion of users of the standard. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


