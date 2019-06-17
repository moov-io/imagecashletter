# \ImageCashLetterFilesApi

All URIs are relative to *http://localhost:8083*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddICLToFile**](ImageCashLetterFilesApi.md#AddICLToFile) | **Post** /files/{file_id}/cashLetters | Add CashLetter to File
[**CreateICLFile**](ImageCashLetterFilesApi.md#CreateICLFile) | **Post** /files/create | Create a new File object
[**DeleteICLFile**](ImageCashLetterFilesApi.md#DeleteICLFile) | **Delete** /files/{file_id} | Permanently deletes a File and associated CashLetters and Bundles. It cannot be undone.
[**DeleteICLFromFile**](ImageCashLetterFilesApi.md#DeleteICLFromFile) | **Delete** /files/{file_id}/cashLetters/{cashLetter_id} | Delete a CashLetter from a File
[**GetICLFileByID**](ImageCashLetterFilesApi.md#GetICLFileByID) | **Get** /files/{file_id} | Retrieves the details of an existing File. You need only supply the unique File identifier that was returned upon creation.
[**GetICLFileContents**](ImageCashLetterFilesApi.md#GetICLFileContents) | **Get** /files/{file_id}/contents | Assembles the existing file (Cash Letters, Bundles and Controls) records, computes sequence numbers and totals. Returns plaintext file.
[**GetICLFiles**](ImageCashLetterFilesApi.md#GetICLFiles) | **Get** /files | Gets a list of Files
[**UpdateICLFile**](ImageCashLetterFilesApi.md#UpdateICLFile) | **Post** /files/{file_id} | Updates the specified File Header by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
[**ValidateICLFile**](ImageCashLetterFilesApi.md#ValidateICLFile) | **Get** /files/{file_id}/validate | Validates the existing file. You need only supply the unique File identifier that was returned upon creation.



## AddICLToFile

> AddICLToFile(ctx, fileId, cashLetter, optional)
Add CashLetter to File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
**cashLetter** | [**CashLetter**](CashLetter.md)|  | 
 **optional** | ***AddICLToFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddICLToFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateICLFile

> File CreateICLFile(ctx, createFile, optional)
Create a new File object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createFile** | [**CreateFile**](CreateFile.md)| Content of the ImageCashLetter file (in json or raw text) | 
 **optional** | ***CreateICLFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateICLFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteICLFile

> DeleteICLFile(ctx, fileId, optional)
Permanently deletes a File and associated CashLetters and Bundles. It cannot be undone.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
 **optional** | ***DeleteICLFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteICLFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteICLFromFile

> DeleteICLFromFile(ctx, fileId, cashLetterId, optional)
Delete a CashLetter from a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
**cashLetterId** | **string**| CashLetter ID | 
 **optional** | ***DeleteICLFromFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteICLFromFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetICLFileByID

> File GetICLFileByID(ctx, fileId, optional)
Retrieves the details of an existing File. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
 **optional** | ***GetICLFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetICLFileByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetICLFileContents

> string GetICLFileContents(ctx, fileId, optional)
Assembles the existing file (Cash Letters, Bundles and Controls) records, computes sequence numbers and totals. Returns plaintext file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
 **optional** | ***GetICLFileContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetICLFileContentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetICLFiles

> []File GetICLFiles(ctx, optional)
Gets a list of Files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetICLFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetICLFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**[]File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateICLFile

> File UpdateICLFile(ctx, fileId, createFile, optional)
Updates the specified File Header by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
**createFile** | [**CreateFile**](CreateFile.md)|  | 
 **optional** | ***UpdateICLFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateICLFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateICLFile

> File ValidateICLFile(ctx, fileId, optional)
Validates the existing file. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string**| File ID | 
 **optional** | ***ValidateICLFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateICLFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

