# \AccountsAPI

All URIs are relative to *https://mail.zoho.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getmailaccounts**](AccountsAPI.md#Getmailaccounts) | **Get** /accounts | Get all accounts



## Getmailaccounts

> AccountsResponse Getmailaccounts(ctx).Execute()

Get all accounts



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.Getmailaccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.Getmailaccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getmailaccounts`: AccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.Getmailaccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetmailaccountsRequest struct via the builder pattern


### Return type

[**AccountsResponse**](AccountsResponse.md)

### Authorization

[zohomail_auth](../README.md#zohomail_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

