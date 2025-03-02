# CreateFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderName** | **string** | The name of the folder to be created. | 
**ParentFolderId** | Pointer to **string** | The ID of the parent folder under which the new folder is created. | [optional] 
**ParentFolderPath** | Pointer to **string** | The path of the parent folder under which the new folder is created. | [optional] 

## Methods

### NewCreateFolderRequest

`func NewCreateFolderRequest(folderName string, ) *CreateFolderRequest`

NewCreateFolderRequest instantiates a new CreateFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFolderRequestWithDefaults

`func NewCreateFolderRequestWithDefaults() *CreateFolderRequest`

NewCreateFolderRequestWithDefaults instantiates a new CreateFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderName

`func (o *CreateFolderRequest) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *CreateFolderRequest) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *CreateFolderRequest) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetParentFolderId

`func (o *CreateFolderRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *CreateFolderRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *CreateFolderRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *CreateFolderRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetParentFolderPath

`func (o *CreateFolderRequest) GetParentFolderPath() string`

GetParentFolderPath returns the ParentFolderPath field if non-nil, zero value otherwise.

### GetParentFolderPathOk

`func (o *CreateFolderRequest) GetParentFolderPathOk() (*string, bool)`

GetParentFolderPathOk returns a tuple with the ParentFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderPath

`func (o *CreateFolderRequest) SetParentFolderPath(v string)`

SetParentFolderPath sets ParentFolderPath field to given value.

### HasParentFolderPath

`func (o *CreateFolderRequest) HasParentFolderPath() bool`

HasParentFolderPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


