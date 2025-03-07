# FolderUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**FolderUpdateMode**](FolderUpdateMode.md) |  | 
**ParentFolderId** | Pointer to **string** | This parameter refers to the folder ID of the destination parent folder where the folder should be moved. | [optional] 
**PreviousFolderId** | Pointer to **string** | This parameter refers to the folder ID of a specific folder under the destination parent folder. This determines the exact position where the folder should be placed. | [optional] 
**FolderName** | Pointer to **string** | This parameter refers to the new name of the folder. | [optional] 

## Methods

### NewFolderUpdatePayload

`func NewFolderUpdatePayload(mode FolderUpdateMode, ) *FolderUpdatePayload`

NewFolderUpdatePayload instantiates a new FolderUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderUpdatePayloadWithDefaults

`func NewFolderUpdatePayloadWithDefaults() *FolderUpdatePayload`

NewFolderUpdatePayloadWithDefaults instantiates a new FolderUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *FolderUpdatePayload) GetMode() FolderUpdateMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FolderUpdatePayload) GetModeOk() (*FolderUpdateMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FolderUpdatePayload) SetMode(v FolderUpdateMode)`

SetMode sets Mode field to given value.


### GetParentFolderId

`func (o *FolderUpdatePayload) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *FolderUpdatePayload) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *FolderUpdatePayload) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *FolderUpdatePayload) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetPreviousFolderId

`func (o *FolderUpdatePayload) GetPreviousFolderId() string`

GetPreviousFolderId returns the PreviousFolderId field if non-nil, zero value otherwise.

### GetPreviousFolderIdOk

`func (o *FolderUpdatePayload) GetPreviousFolderIdOk() (*string, bool)`

GetPreviousFolderIdOk returns a tuple with the PreviousFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousFolderId

`func (o *FolderUpdatePayload) SetPreviousFolderId(v string)`

SetPreviousFolderId sets PreviousFolderId field to given value.

### HasPreviousFolderId

`func (o *FolderUpdatePayload) HasPreviousFolderId() bool`

HasPreviousFolderId returns a boolean if a field has been set.

### GetFolderName

`func (o *FolderUpdatePayload) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *FolderUpdatePayload) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *FolderUpdatePayload) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *FolderUpdatePayload) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


