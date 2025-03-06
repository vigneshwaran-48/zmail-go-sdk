# FoldersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**[]Folder**](Folder.md) |  | [optional] 

## Methods

### NewFoldersResponse

`func NewFoldersResponse() *FoldersResponse`

NewFoldersResponse instantiates a new FoldersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersResponseWithDefaults

`func NewFoldersResponseWithDefaults() *FoldersResponse`

NewFoldersResponseWithDefaults instantiates a new FoldersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FoldersResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FoldersResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FoldersResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FoldersResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *FoldersResponse) GetData() []Folder`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FoldersResponse) GetDataOk() (*[]Folder, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FoldersResponse) SetData(v []Folder)`

SetData sets Data field to given value.

### HasData

`func (o *FoldersResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


