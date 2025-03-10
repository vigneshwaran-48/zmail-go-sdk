# AttachmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 

## Methods

### NewAttachmentsResponse

`func NewAttachmentsResponse() *AttachmentsResponse`

NewAttachmentsResponse instantiates a new AttachmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentsResponseWithDefaults

`func NewAttachmentsResponseWithDefaults() *AttachmentsResponse`

NewAttachmentsResponseWithDefaults instantiates a new AttachmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AttachmentsResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachmentsResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachmentsResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachmentsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *AttachmentsResponse) GetData() []Attachment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AttachmentsResponse) GetDataOk() (*[]Attachment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AttachmentsResponse) SetData(v []Attachment)`

SetData sets Data field to given value.

### HasData

`func (o *AttachmentsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


