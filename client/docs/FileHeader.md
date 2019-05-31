# FileHeader

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | FileHeader ID | [optional] 
**StandardLevel** | **string** | StandardLevel identifies the standard level of the file.  * &#x60;03&#x60; - DSTU X9.37 - 2003 * &#x60;30&#x60; - X9.100-187-2008 * &#x60;35&#x60; - X9.100-187-2013 and 2016  | 
**TestFileIndicator** | **string** | TestFileIndicator dentifies whether the file is a test or production file.  * &#x60;T&#x60; - Test File * &#x60;P&#x60; - Production File  | 
**ImmediateDestination** | **string** | ImmediateDestination is the routing and transit number of the Federal Reserve Bank (FRB) or receiver to which the file is being sent.  | 
**ImmediateOrigin** | **string** | ImmediateOrigin is the routing and transit number of the Federal Reserve Bank (FRB) or originator from which the file is being sent.  | 
**FileCreationDate** | [**time.Time**](time.Time.md) | FileCreationDate is the date that the immediate origin institution creates the file. | 
**FileCreationTime** | **string** | FileCreationTime is the time the immediate origin institution creates the file. (Format - hhmm, where - hh hour, mm minute)  | 
**ResendIndicator** | **string** | ResendIndicator Indicates whether the file has been previously transmitted. (Y - Yes, N - No) | 
**ImmediateDestinationName** | **string** | ImmediateDestinationName Identifies the short name of the institution that receives the file. | [optional] 
**ImmediateOriginName** | **string** | immediateOriginName identifies the short name of the institution that sends the file. | [optional] 
**FileIDModifier** | **string** | FileIDModifier is a code that permits multiple files, created on the same date, same time and between the same institutions, to be distinguished one from another. If all of the following fields in a previous file are equal to the same fields in this file: FileHeader ImmediateDestination, ImmediateOrigin, FileCreationDate, and FileCreationTime, it must be defined.  | [optional] 
**CountryCode** | **string** | CountryCode is a 2-character code as approved by the International Organization for Standardization (ISO) used to identify the country in which the payer bank is located.  | [optional] 
**UserField** | **string** | UserField identifies a field used at the discretion of users of the standard. | [optional] 
**CompanionDocumentIndicator** | **string** | CompanionDocumentIndicator identifies a field used to indicate the Companion Document being used. Shall be present only under clearing arrangements. Companion Document usage and values defined by clearing arrangements. Values: 0–9 Reserved for United States use A–J Reserved for Canadian use Other - as defined by clearing arrangements.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


