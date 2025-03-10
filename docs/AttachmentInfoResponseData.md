# AttachmentInfoResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]EmailAttachment**](EmailAttachment.md) |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 

## Methods

### NewAttachmentInfoResponseData

`func NewAttachmentInfoResponseData() *AttachmentInfoResponseData`

NewAttachmentInfoResponseData instantiates a new AttachmentInfoResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentInfoResponseDataWithDefaults

`func NewAttachmentInfoResponseDataWithDefaults() *AttachmentInfoResponseData`

NewAttachmentInfoResponseDataWithDefaults instantiates a new AttachmentInfoResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *AttachmentInfoResponseData) GetAttachments() []EmailAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AttachmentInfoResponseData) GetAttachmentsOk() (*[]EmailAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AttachmentInfoResponseData) SetAttachments(v []EmailAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AttachmentInfoResponseData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetMessageId

`func (o *AttachmentInfoResponseData) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AttachmentInfoResponseData) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AttachmentInfoResponseData) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *AttachmentInfoResponseData) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


