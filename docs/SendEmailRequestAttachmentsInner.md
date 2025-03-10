# SendEmailRequestAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentName** | Pointer to **string** | Specifies the name of the attachment. This parameter can be fetched from Upload Attachments API. | [optional] 
**AttachmentPath** | Pointer to **string** | Specifies the path in which the attachment is stored. This parameter can be fetched from Upload Attachments API. | [optional] 
**StoreName** | Pointer to **string** | Specifies the name of the store where the attachment is saved. This parameter can be fetched from Upload Attachments API. | [optional] 

## Methods

### NewSendEmailRequestAttachmentsInner

`func NewSendEmailRequestAttachmentsInner() *SendEmailRequestAttachmentsInner`

NewSendEmailRequestAttachmentsInner instantiates a new SendEmailRequestAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailRequestAttachmentsInnerWithDefaults

`func NewSendEmailRequestAttachmentsInnerWithDefaults() *SendEmailRequestAttachmentsInner`

NewSendEmailRequestAttachmentsInnerWithDefaults instantiates a new SendEmailRequestAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentName

`func (o *SendEmailRequestAttachmentsInner) GetAttachmentName() string`

GetAttachmentName returns the AttachmentName field if non-nil, zero value otherwise.

### GetAttachmentNameOk

`func (o *SendEmailRequestAttachmentsInner) GetAttachmentNameOk() (*string, bool)`

GetAttachmentNameOk returns a tuple with the AttachmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentName

`func (o *SendEmailRequestAttachmentsInner) SetAttachmentName(v string)`

SetAttachmentName sets AttachmentName field to given value.

### HasAttachmentName

`func (o *SendEmailRequestAttachmentsInner) HasAttachmentName() bool`

HasAttachmentName returns a boolean if a field has been set.

### GetAttachmentPath

`func (o *SendEmailRequestAttachmentsInner) GetAttachmentPath() string`

GetAttachmentPath returns the AttachmentPath field if non-nil, zero value otherwise.

### GetAttachmentPathOk

`func (o *SendEmailRequestAttachmentsInner) GetAttachmentPathOk() (*string, bool)`

GetAttachmentPathOk returns a tuple with the AttachmentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentPath

`func (o *SendEmailRequestAttachmentsInner) SetAttachmentPath(v string)`

SetAttachmentPath sets AttachmentPath field to given value.

### HasAttachmentPath

`func (o *SendEmailRequestAttachmentsInner) HasAttachmentPath() bool`

HasAttachmentPath returns a boolean if a field has been set.

### GetStoreName

`func (o *SendEmailRequestAttachmentsInner) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *SendEmailRequestAttachmentsInner) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *SendEmailRequestAttachmentsInner) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *SendEmailRequestAttachmentsInner) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


