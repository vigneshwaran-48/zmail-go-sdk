# \EmailAPI

All URIs are relative to *https://mail.zoho.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEmails**](EmailAPI.md#ListEmails) | **Get** /api/accounts/{accountId}/messages/view | Retrieves emails



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

