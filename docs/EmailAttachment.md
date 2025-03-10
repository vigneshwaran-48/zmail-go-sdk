# EmailAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentSize** | Pointer to **int32** | Size of the attachment in bytes. | [optional] 
**AttachmentName** | Pointer to **string** | Name of the attachment file. | [optional] 
**AttachmentId** | Pointer to **string** | Unique identifier for the attachment. | [optional] 

## Methods

### NewEmailAttachment

`func NewEmailAttachment() *EmailAttachment`

NewEmailAttachment instantiates a new EmailAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAttachmentWithDefaults

`func NewEmailAttachmentWithDefaults() *EmailAttachment`

NewEmailAttachmentWithDefaults instantiates a new EmailAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentSize

`func (o *EmailAttachment) GetAttachmentSize() int32`

GetAttachmentSize returns the AttachmentSize field if non-nil, zero value otherwise.

### GetAttachmentSizeOk

`func (o *EmailAttachment) GetAttachmentSizeOk() (*int32, bool)`

GetAttachmentSizeOk returns a tuple with the AttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentSize

`func (o *EmailAttachment) SetAttachmentSize(v int32)`

SetAttachmentSize sets AttachmentSize field to given value.

### HasAttachmentSize

`func (o *EmailAttachment) HasAttachmentSize() bool`

HasAttachmentSize returns a boolean if a field has been set.

### GetAttachmentName

`func (o *EmailAttachment) GetAttachmentName() string`

GetAttachmentName returns the AttachmentName field if non-nil, zero value otherwise.

### GetAttachmentNameOk

`func (o *EmailAttachment) GetAttachmentNameOk() (*string, bool)`

GetAttachmentNameOk returns a tuple with the AttachmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentName

`func (o *EmailAttachment) SetAttachmentName(v string)`

SetAttachmentName sets AttachmentName field to given value.

### HasAttachmentName

`func (o *EmailAttachment) HasAttachmentName() bool`

HasAttachmentName returns a boolean if a field has been set.

### GetAttachmentId

`func (o *EmailAttachment) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *EmailAttachment) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *EmailAttachment) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *EmailAttachment) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


