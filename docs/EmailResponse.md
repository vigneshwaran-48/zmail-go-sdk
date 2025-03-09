# EmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**Email**](Email.md) |  | [optional] 

## Methods

### NewEmailResponse

`func NewEmailResponse() *EmailResponse`

NewEmailResponse instantiates a new EmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailResponseWithDefaults

`func NewEmailResponseWithDefaults() *EmailResponse`

NewEmailResponseWithDefaults instantiates a new EmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EmailResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *EmailResponse) GetData() Email`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailResponse) GetDataOk() (*Email, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailResponse) SetData(v Email)`

SetData sets Data field to given value.

### HasData

`func (o *EmailResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


