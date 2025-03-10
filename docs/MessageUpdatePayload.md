# MessageUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**MessageUpdateMode**](MessageUpdateMode.md) |  | 
**MessageId** | **[]string** |  | 
**ThreadId** | Pointer to **[]string** |  | [optional] 
**IsFolderSpecific** | Pointer to **bool** |  | [optional] [default to false]
**FolderId** | Pointer to **string** |  | [optional] 
**LabelId** | Pointer to **[]string** |  | [optional] 
**IsArchive** | Pointer to **bool** |  | [optional] [default to false]
**Flagid** | Pointer to [**MessageFlag**](MessageFlag.md) |  | [optional] 
**DestfolderId** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageUpdatePayload

`func NewMessageUpdatePayload(mode MessageUpdateMode, messageId []string, ) *MessageUpdatePayload`

NewMessageUpdatePayload instantiates a new MessageUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageUpdatePayloadWithDefaults

`func NewMessageUpdatePayloadWithDefaults() *MessageUpdatePayload`

NewMessageUpdatePayloadWithDefaults instantiates a new MessageUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *MessageUpdatePayload) GetMode() MessageUpdateMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MessageUpdatePayload) GetModeOk() (*MessageUpdateMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MessageUpdatePayload) SetMode(v MessageUpdateMode)`

SetMode sets Mode field to given value.


### GetMessageId

`func (o *MessageUpdatePayload) GetMessageId() []string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageUpdatePayload) GetMessageIdOk() (*[]string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageUpdatePayload) SetMessageId(v []string)`

SetMessageId sets MessageId field to given value.


### GetThreadId

`func (o *MessageUpdatePayload) GetThreadId() []string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *MessageUpdatePayload) GetThreadIdOk() (*[]string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *MessageUpdatePayload) SetThreadId(v []string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *MessageUpdatePayload) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetIsFolderSpecific

`func (o *MessageUpdatePayload) GetIsFolderSpecific() bool`

GetIsFolderSpecific returns the IsFolderSpecific field if non-nil, zero value otherwise.

### GetIsFolderSpecificOk

`func (o *MessageUpdatePayload) GetIsFolderSpecificOk() (*bool, bool)`

GetIsFolderSpecificOk returns a tuple with the IsFolderSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolderSpecific

`func (o *MessageUpdatePayload) SetIsFolderSpecific(v bool)`

SetIsFolderSpecific sets IsFolderSpecific field to given value.

### HasIsFolderSpecific

`func (o *MessageUpdatePayload) HasIsFolderSpecific() bool`

HasIsFolderSpecific returns a boolean if a field has been set.

### GetFolderId

`func (o *MessageUpdatePayload) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MessageUpdatePayload) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MessageUpdatePayload) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MessageUpdatePayload) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetLabelId

`func (o *MessageUpdatePayload) GetLabelId() []string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *MessageUpdatePayload) GetLabelIdOk() (*[]string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *MessageUpdatePayload) SetLabelId(v []string)`

SetLabelId sets LabelId field to given value.

### HasLabelId

`func (o *MessageUpdatePayload) HasLabelId() bool`

HasLabelId returns a boolean if a field has been set.

### GetIsArchive

`func (o *MessageUpdatePayload) GetIsArchive() bool`

GetIsArchive returns the IsArchive field if non-nil, zero value otherwise.

### GetIsArchiveOk

`func (o *MessageUpdatePayload) GetIsArchiveOk() (*bool, bool)`

GetIsArchiveOk returns a tuple with the IsArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchive

`func (o *MessageUpdatePayload) SetIsArchive(v bool)`

SetIsArchive sets IsArchive field to given value.

### HasIsArchive

`func (o *MessageUpdatePayload) HasIsArchive() bool`

HasIsArchive returns a boolean if a field has been set.

### GetFlagid

`func (o *MessageUpdatePayload) GetFlagid() MessageFlag`

GetFlagid returns the Flagid field if non-nil, zero value otherwise.

### GetFlagidOk

`func (o *MessageUpdatePayload) GetFlagidOk() (*MessageFlag, bool)`

GetFlagidOk returns a tuple with the Flagid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagid

`func (o *MessageUpdatePayload) SetFlagid(v MessageFlag)`

SetFlagid sets Flagid field to given value.

### HasFlagid

`func (o *MessageUpdatePayload) HasFlagid() bool`

HasFlagid returns a boolean if a field has been set.

### GetDestfolderId

`func (o *MessageUpdatePayload) GetDestfolderId() string`

GetDestfolderId returns the DestfolderId field if non-nil, zero value otherwise.

### GetDestfolderIdOk

`func (o *MessageUpdatePayload) GetDestfolderIdOk() (*string, bool)`

GetDestfolderIdOk returns a tuple with the DestfolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestfolderId

`func (o *MessageUpdatePayload) SetDestfolderId(v string)`

SetDestfolderId sets DestfolderId field to given value.

### HasDestfolderId

`func (o *MessageUpdatePayload) HasDestfolderId() bool`

HasDestfolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


