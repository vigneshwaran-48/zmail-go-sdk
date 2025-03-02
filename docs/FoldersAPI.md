# \FoldersAPI

All URIs are relative to *https://mail.zoho.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FoldersAPI.md#CreateFolder) | **Post** /accounts/{accountId}/folders | Create a New Folder



## CreateFolder

> FolderResponse CreateFolder(ctx, accountId).CreateFolderRequest(createFolderRequest).Execute()

Create a New Folder



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
	accountId := "accountId_example" // string | The ID of the account where the folder will be created.
	createFolderRequest := *openapiclient.NewCreateFolderRequest("newFolder") // CreateFolderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.CreateFolder(context.Background(), accountId).CreateFolderRequest(createFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.CreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder`: FolderResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.CreateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account where the folder will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFolderRequest** | [**CreateFolderRequest**](CreateFolderRequest.md) |  | 

### Return type

[**FolderResponse**](FolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

