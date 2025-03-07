# \FoldersAPI

All URIs are relative to *https://mail.zoho.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FoldersAPI.md#CreateFolder) | **Post** /accounts/{accountId}/folders | Create a New Folder
[**DeleteFolder**](FoldersAPI.md#DeleteFolder) | **Delete** /accounts/{accountId}/folders/{folderId} | Delete a folder
[**GetAllFolders**](FoldersAPI.md#GetAllFolders) | **Get** /accounts/{accountId}/folders | Get all folders of the account
[**GetFolder**](FoldersAPI.md#GetFolder) | **Get** /accounts/{accountId}/folders/{folderId} | Get a specific folder of the account
[**UpdateFolder**](FoldersAPI.md#UpdateFolder) | **Put** /accounts/{accountId}/folders/{folderId} | Update a folder



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


## DeleteFolder

> NoDataResponse DeleteFolder(ctx, accountId, folderId).Execute()

Delete a folder



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
	folderId := "folderId_example" // string | This key is used to identify the folder to be fetched.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.DeleteFolder(context.Background(), accountId, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.DeleteFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFolder`: NoDataResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.DeleteFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition. | 
**folderId** | **string** | This key is used to identify the folder to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NoDataResponse**](NoDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFolders

> FoldersResponse GetAllFolders(ctx, accountId).Execute()

Get all folders of the account



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
	accountId := "accountId_example" // string | The ID of the account where the folder will be listed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetAllFolders(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetAllFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFolders`: FoldersResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetAllFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account where the folder will be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FoldersResponse**](FoldersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> FolderResponse GetFolder(ctx, accountId, folderId).Execute()

Get a specific folder of the account



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
	accountId := "accountId_example" // string | The ID of the account where the folder will be listed.
	folderId := "folderId_example" // string | This key is used to identify the folder to be fetched.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFolder(context.Background(), accountId, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolder`: FolderResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The ID of the account where the folder will be listed. | 
**folderId** | **string** | This key is used to identify the folder to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FolderResponse**](FolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolder

> NoDataResponse UpdateFolder(ctx, accountId, folderId).FolderUpdatePayload(folderUpdatePayload).Execute()

Update a folder



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
	folderId := "folderId_example" // string | This key is used to identify the folder to be fetched.
	folderUpdatePayload := *openapiclient.NewFolderUpdatePayload(openapiclient.FolderUpdateMode("move")) // FolderUpdatePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolder(context.Background(), accountId, folderId).FolderUpdatePayload(folderUpdatePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolder`: NoDataResponse
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | This key is used to identify the account from which the folders have to be fetched. It is generated during account addition. | 
**folderId** | **string** | This key is used to identify the folder to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **folderUpdatePayload** | [**FolderUpdatePayload**](FolderUpdatePayload.md) |  | 

### Return type

[**NoDataResponse**](NoDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

