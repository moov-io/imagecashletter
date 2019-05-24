# CheckDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | CheckDetail ID | [optional] 
**AuxiliaryOnUs** | **string** | AuxiliaryOnUs identifies a code used on commercial checks at the discretion of the payor bank. | [optional] 
**ExternalProcessingCode** | **string** | ExternalProcessingCode identifies a code used for special purposes as authorized by the Accredited Standards Committee X9. Also known as Position 44. | [optional] 
**PayorBankRoutingNumber** | **string** | PayorBankRoutingNumber identifies a number that identifies the institution by or through which the item is payable. Must be a valid routing and transit number issued by the ABA’s Routing Number Registrar. Shall represent the first 8 digits of a 9-digit routing number or 8 numeric digits of a 4 dash 4 routing number. A valid routing number consists of 2 fields: the eight- digit Payor Bank Routing Number  and the one-digit Payor Bank Routing Number Check Digit.  | [optional] 
**PayorBankCheckDigit** | **string** | PayorBankCheckDigit identifies a digit representing the routing number check digit.  The combination of Payor Bank Routing Number and payor Bank Routing Number Check Digit must be a mod-checked routing number with a valid check digit.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


