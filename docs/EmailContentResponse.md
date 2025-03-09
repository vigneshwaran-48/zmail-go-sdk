# EmailContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**EmailContentResponseData**](EmailContentResponseData.md) |  | [optional] 

## Methods

### NewEmailContentResponse

`func NewEmailContentResponse() *EmailContentResponse`

NewEmailContentResponse instantiates a new EmailContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailContentResponseWithDefaults

`func NewEmailContentResponseWithDefaults() *EmailContentResponse`

NewEmailContentResponseWithDefaults instantiates a new EmailContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EmailContentResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailContentResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailContentResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailContentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *EmailContentResponse) GetData() EmailContentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailContentResponse) GetDataOk() (*EmailContentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailContentResponse) SetData(v EmailContentResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *EmailContentResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


