# \EmailAPI

All URIs are relative to *https://mail.zoho.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessageContent**](EmailAPI.md#GetMessageContent) | **Get** /api/accounts/{accountId}/folders/{folderId}/messages/{messageId}/content | Retrives the content of email
[**GetMessageDetails**](EmailAPI.md#GetMessageDetails) | **Get** /api/accounts/{accountId}/folders/{folderId}/messages/{messageId}/details | Retrives the message details
[**GetMessageHeader**](EmailAPI.md#GetMessageHeader) | **Get** /api/accounts/{accountId}/folders/{folderId}/messages/{messageId}/header | Retrives message headers
[**GetOriginalMessage**](EmailAPI.md#GetOriginalMessage) | **Get** /api/accounts/{accountId}/messages/{messageId}/originalmessage | Retrives the original content of email
[**ListEmails**](EmailAPI.md#ListEmails) | **Get** /api/accounts/{accountId}/messages/view | Retrieves emails
[**SearchEmails**](EmailAPI.md#SearchEmails) | **Get** /api/accounts/{accountId}/messages/search | Searches emails



## GetMessageContent

> EmailContentResponse GetMessageContent(ctx, accountId, folderId, messageId).IncludeBlockContent(includeBlockContent).Execute()

Retrives the content of email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vigneshwaran-48/zmail-go-sdk"
)

func main() {
	accountId := "accountId_example" // string | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition.
	folderId := "folderId_example" // string | This key is used to identify the folder to be used.
	messageId := "messageId_example" // string | This key is used to identify the message to be used.
	includeBlockContent := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetMessageContent(context.Background(), accountId, folderId, messageId).IncludeBlockContent(includeBlockContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetMessageContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageContent`: EmailContentResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetMessageContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition. | 
**folderId** | **string** | This key is used to identify the folder to be used. | 
**messageId** | **string** | This key is used to identify the message to be used. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeBlockContent** | **bool** |  | [default to false]

### Return type

[**EmailContentResponse**](EmailContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageDetails

> EmailResponse GetMessageDetails(ctx, accountId, folderId, messageId).Execute()

Retrives the message details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vigneshwaran-48/zmail-go-sdk"
)

func main() {
	accountId := "accountId_example" // string | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition.
	folderId := "folderId_example" // string | This key is used to identify the folder to be used.
	messageId := "messageId_example" // string | This key is used to identify the message to be used.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetMessageDetails(context.Background(), accountId, folderId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetMessageDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageDetails`: EmailResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetMessageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition. | 
**folderId** | **string** | This key is used to identify the folder to be used. | 
**messageId** | **string** | This key is used to identify the message to be used. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmailResponse**](EmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageHeader

> GetMessageHeader200Response GetMessageHeader(ctx, accountId, folderId, messageId).Raw(raw).Execute()

Retrives message headers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vigneshwaran-48/zmail-go-sdk"
)

func main() {
	accountId := "accountId_example" // string | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition.
	folderId := "folderId_example" // string | This key is used to identify the folder to be used.
	messageId := "messageId_example" // string | This key is used to identify the message to be used.
	raw := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetMessageHeader(context.Background(), accountId, folderId, messageId).Raw(raw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetMessageHeader``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageHeader`: GetMessageHeader200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetMessageHeader`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition. | 
**folderId** | **string** | This key is used to identify the folder to be used. | 
**messageId** | **string** | This key is used to identify the message to be used. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **raw** | **bool** |  | [default to true]

### Return type

[**GetMessageHeader200Response**](GetMessageHeader200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOriginalMessage

> EmailContentResponse GetOriginalMessage(ctx, accountId, messageId).Execute()

Retrives the original content of email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vigneshwaran-48/zmail-go-sdk"
)

func main() {
	accountId := "accountId_example" // string | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition.
	messageId := "messageId_example" // string | This key is used to identify the message to be used.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetOriginalMessage(context.Background(), accountId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetOriginalMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOriginalMessage`: EmailContentResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetOriginalMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition. | 
**messageId** | **string** | This key is used to identify the message to be used. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOriginalMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailContentResponse**](EmailContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmails

> EmailsResponse ListEmails(ctx, accountId).FolderId(folderId).Start(start).Limit(limit).Status(status).Flagid(flagid).Labelid(labelid).ThreadId(threadId).SortBy(sortBy).Sortorder(sortorder).Includeto(includeto).Includesent(includesent).Includearchive(includearchive).AttachedMails(attachedMails).InlinedMails(inlinedMails).FlaggedMails(flaggedMails).RespondedMails(respondedMails).ThreadedMails(threadedMails).Execute()

Retrieves emails



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vigneshwaran-48/zmail-go-sdk"
)

func main() {
	accountId := "accountId_example" // string | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition.
	folderId := int64(789) // int64 | The unique identifier for the folder from which the emails need to be retrieved. This can be fetched using the Get all folders API. (optional)
	start := int32(56) // int32 | The starting sequence number of the emails to be retrieved. (optional) (default to 1)
	limit := int32(56) // int32 | The number of emails to be retrieved from the start value mentioned. Allowed values:- Min. value:- 1 and max. value:- 200. (optional) (default to 10)
	status := "status_example" // string | Retrieve emails by read or unread status. (optional) (default to "all")
	flagid := int32(56) // int32 | The unique identifier for the flag used to retrieve emails based on a specific flag type. (optional)
	labelid := int64(789) // int64 | The unique identifier for the label used to retrieve emails based on a specific label. (optional)
	threadId := int64(789) // int64 | The unique identifier for the flag used to retrieve emails of the given threadId. (optional)
	sortBy := "sortBy_example" // string | The basis on which the sorting of the list of emails should be done. (optional) (default to "date")
	sortorder := true // bool | The order in which the sorting of the list of emails should be done. (optional) (default to false)
	includeto := true // bool | Whether to details need to be included or not on the list of emails retrieved. (optional) (default to false)
	includesent := true // bool | Whether sent emails need to be included or not on the list of emails retrieved. (optional) (default to false)
	includearchive := true // bool | Whether archived emails need to be included or not on the list of emails retrieved. (optional) (default to false)
	attachedMails := true // bool | Retrieve only the emails with attachments. (optional) (default to false)
	inlinedMails := true // bool | Retrieve only the emails with inline images. (optional) (default to false)
	flaggedMails := true // bool | Retrieve only flagged emails. (optional) (default to false)
	respondedMails := true // bool | Retrieve only emails with replies. (optional) (default to false)
	threadedMails := true // bool | Retrieve emails that are a part of conversations. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.ListEmails(context.Background(), accountId).FolderId(folderId).Start(start).Limit(limit).Status(status).Flagid(flagid).Labelid(labelid).ThreadId(threadId).SortBy(sortBy).Sortorder(sortorder).Includeto(includeto).Includesent(includesent).Includearchive(includearchive).AttachedMails(attachedMails).InlinedMails(inlinedMails).FlaggedMails(flaggedMails).RespondedMails(respondedMails).ThreadedMails(threadedMails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.ListEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEmails`: EmailsResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.ListEmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderId** | **int64** | The unique identifier for the folder from which the emails need to be retrieved. This can be fetched using the Get all folders API. | 
 **start** | **int32** | The starting sequence number of the emails to be retrieved. | [default to 1]
 **limit** | **int32** | The number of emails to be retrieved from the start value mentioned. Allowed values:- Min. value:- 1 and max. value:- 200. | [default to 10]
 **status** | **string** | Retrieve emails by read or unread status. | [default to &quot;all&quot;]
 **flagid** | **int32** | The unique identifier for the flag used to retrieve emails based on a specific flag type. | 
 **labelid** | **int64** | The unique identifier for the label used to retrieve emails based on a specific label. | 
 **threadId** | **int64** | The unique identifier for the flag used to retrieve emails of the given threadId. | 
 **sortBy** | **string** | The basis on which the sorting of the list of emails should be done. | [default to &quot;date&quot;]
 **sortorder** | **bool** | The order in which the sorting of the list of emails should be done. | [default to false]
 **includeto** | **bool** | Whether to details need to be included or not on the list of emails retrieved. | [default to false]
 **includesent** | **bool** | Whether sent emails need to be included or not on the list of emails retrieved. | [default to false]
 **includearchive** | **bool** | Whether archived emails need to be included or not on the list of emails retrieved. | [default to false]
 **attachedMails** | **bool** | Retrieve only the emails with attachments. | [default to false]
 **inlinedMails** | **bool** | Retrieve only the emails with inline images. | [default to false]
 **flaggedMails** | **bool** | Retrieve only flagged emails. | [default to false]
 **respondedMails** | **bool** | Retrieve only emails with replies. | [default to false]
 **threadedMails** | **bool** | Retrieve emails that are a part of conversations. | [default to false]

### Return type

[**EmailsResponse**](EmailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEmails

> EmailsResponse SearchEmails(ctx, accountId).SearchKey(searchKey).ReceivedTime(receivedTime).Start(start).Limit(limit).Includeto(includeto).Execute()

Searches emails



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vigneshwaran-48/zmail-go-sdk"
)

func main() {
	accountId := "accountId_example" // string | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition.
	searchKey := "searchKey_example" // string | Specifies the search criteria. Ensure that the searchKey is constructed using the search syntax defined in this help page. For example:- To search for new emails, provide the searchKey as newMails.
	receivedTime := int64(789) // int64 | Specifies the time before which emails were received. It allows users to filter emails based on their received timestamp. Format:- Unix timestamp in milliseconds. Example:- 1609459200000. By default, the API retrieves emails received before 2 minutes from the current time unless a specific timestamp is provided. (optional)
	start := int32(56) // int32 | Specifies the starting sequence number of the emails to be retrieved. (optional) (default to 1)
	limit := int32(56) // int32 | Specifies the number of emails to be retrieved from the start value mentioned. Allowed values:- Min. value:- 1 and max. value:- 200. (optional) (default to 10)
	includeto := true // bool | Specifies whether To details need to be included or not on the list of emails retrieved. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.SearchEmails(context.Background(), accountId).SearchKey(searchKey).ReceivedTime(receivedTime).Start(start).Limit(limit).Includeto(includeto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.SearchEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEmails`: EmailsResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.SearchEmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchKey** | **string** | Specifies the search criteria. Ensure that the searchKey is constructed using the search syntax defined in this help page. For example:- To search for new emails, provide the searchKey as newMails. | 
 **receivedTime** | **int64** | Specifies the time before which emails were received. It allows users to filter emails based on their received timestamp. Format:- Unix timestamp in milliseconds. Example:- 1609459200000. By default, the API retrieves emails received before 2 minutes from the current time unless a specific timestamp is provided. | 
 **start** | **int32** | Specifies the starting sequence number of the emails to be retrieved. | [default to 1]
 **limit** | **int32** | Specifies the number of emails to be retrieved from the start value mentioned. Allowed values:- Min. value:- 1 and max. value:- 200. | [default to 10]
 **includeto** | **bool** | Specifies whether To details need to be included or not on the list of emails retrieved. | [default to false]

### Return type

[**EmailsResponse**](EmailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

