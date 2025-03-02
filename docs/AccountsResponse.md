# AccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**[]Account**](Account.md) |  | [optional] 

## Methods

### NewAccountsResponse

`func NewAccountsResponse() *AccountsResponse`

NewAccountsResponse instantiates a new AccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsResponseWithDefaults

`func NewAccountsResponseWithDefaults() *AccountsResponse`

NewAccountsResponseWithDefaults instantiates a new AccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccountsResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountsResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountsResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *AccountsResponse) GetData() []Account`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountsResponse) GetDataOk() (*[]Account, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountsResponse) SetData(v []Account)`

SetData sets Data field to given value.

### HasData

`func (o *AccountsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


