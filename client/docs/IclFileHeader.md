# IclFileHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | FileHeader ID | [optional] 
**StandardLevel** | **string** | StandardLevel identifies the standard level of the file.  * &#x60;03&#x60; - DSTU X9.37-2003 * &#x60;30&#x60; - X9.100-187-2008 * &#x60;35&#x60; - X9.100-187-2013 and X9.100-187-2016  | 
**TestIndicator** | **string** | TestIndicator identifies whether the file is a test or production file.  * &#x60;T&#x60; - Test File * &#x60;P&#x60; - Production File  | 
**ImmediateDestination** | **string** | ImmediateDestination is the routing and transit number of the Federal Reserve Bank (FRB) or receiver to which the file is being sent.  | 
**ImmediateOrigin** | **string** | ImmediateOrigin is the routing and transit number of the Federal Reserve Bank (FRB) or originator from which the file is being sent.  | 
**FileCreationDate** | [**time.Time**](time.Time.md) | FileCreationDate is the date the immediate origin institution creates the file. | 
**FileCreationTime** | [**time.Time**](time.Time.md) | FileCreationTime is the time the immediate origin institution creates the file. | 
**ResendIndicator** | **string** | ResendIndicator indicates whether the file has been previously transmitted. (Y - Yes, N - No) | 
**ImmediateDestinationName** | **string** | ImmediateDestinationName identifies the short name of the institution that receives the file. | [optional] 
**ImmediateOriginName** | **string** | immediateOriginName identifies the short name of the institution that sends the file. | [optional] 
**FileIDModifier** | **string** | FileIDModifier is a code that permits multiple files, created on the same date, at the same time, and sent between the same institutions, to be distinguished from one another. If FileHeader ImmediateDestination, ImmediateOrigin, FileCreationDate, and FileCreationTime in a previous file are equal to the same fields in this file, FileIDModifier must be defined.  | [optional] 
**CountryCode** | **string** | CountryCode is a 2-character code as approved by the International Organization for Standardization (ISO) used to identify the country in which the payer bank is located.  | [optional] 
**UserField** | **string** | UserField identifies a field used at the discretion of users of the standard. | [optional] 
**CompanionDocumentIndicator** | **string** | CompanionDocumentIndicator indicates the Companion Document being used. It shall be present only under clearing arrangements, where Companion Document usage and values are defined. Values: * 0–9 - Reserved for United States use * A–J - Reserved for Canadian use * Other - Defined by clearing arrangements  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


