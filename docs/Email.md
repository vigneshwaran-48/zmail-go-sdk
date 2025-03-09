# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to **string** | A brief summary of the email. | [optional] 
**SentDateInGMT** | Pointer to **int64** | The sent date of the email in GMT, represented as a Unix timestamp in milliseconds. | [optional] 
**CalendarType** | Pointer to **int32** | The type of calendar associated with the email. | [optional] 
**Subject** | Pointer to **string** | The subject of the email. | [optional] 
**MessageId** | Pointer to **string** | The unique identifier of the email message. | [optional] 
**ThreadCount** | Pointer to **string** | The number of emails in the thread. | [optional] 
**Flagid** | Pointer to **string** | The flag identifier associated with the email (e.g., flag_not_set). | [optional] 
**Status2** | Pointer to **string** | An alternative status code for the email. | [optional] 
**Priority** | Pointer to **string** | The priority of the email. | [optional] 
**HasInline** | Pointer to **string** | Indicates whether the email has inline content (true/false, represented as string). | [optional] 
**ToAddress** | Pointer to **string** | The recipient&#39;s email address, including display name. | [optional] 
**FolderId** | Pointer to **string** | The unique identifier of the folder containing the email. | [optional] 
**CcAddress** | Pointer to **string** | The CC recipient&#39;s email address, or \&quot;Not Provided\&quot;. | [optional] 
**ThreadId** | Pointer to **string** | The unique identifier of the email thread. | [optional] 
**HasAttachment** | Pointer to **string** | Indicates whether the email has attachments (0 or 1, represented as string). | [optional] 
**Size** | Pointer to **string** | The size of the email. | [optional] 
**Sender** | Pointer to **string** | The sender&#39;s display name. | [optional] 
**ReceivedTime** | Pointer to **int64** | The received time of the email, represented as a Unix timestamp in milliseconds. | [optional] 
**FromAddress** | Pointer to **string** | The sender&#39;s email address. | [optional] 
**Status** | Pointer to **string** | The status of the email (e.g., 1 for read, 0 for unread). | [optional] 

## Methods

### NewEmail

`func NewEmail() *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *Email) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Email) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Email) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Email) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSentDateInGMT

`func (o *Email) GetSentDateInGMT() int64`

GetSentDateInGMT returns the SentDateInGMT field if non-nil, zero value otherwise.

### GetSentDateInGMTOk

`func (o *Email) GetSentDateInGMTOk() (*int64, bool)`

GetSentDateInGMTOk returns a tuple with the SentDateInGMT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentDateInGMT

`func (o *Email) SetSentDateInGMT(v int64)`

SetSentDateInGMT sets SentDateInGMT field to given value.

### HasSentDateInGMT

`func (o *Email) HasSentDateInGMT() bool`

HasSentDateInGMT returns a boolean if a field has been set.

### GetCalendarType

`func (o *Email) GetCalendarType() int32`

GetCalendarType returns the CalendarType field if non-nil, zero value otherwise.

### GetCalendarTypeOk

`func (o *Email) GetCalendarTypeOk() (*int32, bool)`

GetCalendarTypeOk returns a tuple with the CalendarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarType

`func (o *Email) SetCalendarType(v int32)`

SetCalendarType sets CalendarType field to given value.

### HasCalendarType

`func (o *Email) HasCalendarType() bool`

HasCalendarType returns a boolean if a field has been set.

### GetSubject

`func (o *Email) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Email) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Email) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Email) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessageId

`func (o *Email) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Email) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Email) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Email) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetThreadCount

`func (o *Email) GetThreadCount() string`

GetThreadCount returns the ThreadCount field if non-nil, zero value otherwise.

### GetThreadCountOk

`func (o *Email) GetThreadCountOk() (*string, bool)`

GetThreadCountOk returns a tuple with the ThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadCount

`func (o *Email) SetThreadCount(v string)`

SetThreadCount sets ThreadCount field to given value.

### HasThreadCount

`func (o *Email) HasThreadCount() bool`

HasThreadCount returns a boolean if a field has been set.

### GetFlagid

`func (o *Email) GetFlagid() string`

GetFlagid returns the Flagid field if non-nil, zero value otherwise.

### GetFlagidOk

`func (o *Email) GetFlagidOk() (*string, bool)`

GetFlagidOk returns a tuple with the Flagid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagid

`func (o *Email) SetFlagid(v string)`

SetFlagid sets Flagid field to given value.

### HasFlagid

`func (o *Email) HasFlagid() bool`

HasFlagid returns a boolean if a field has been set.

### GetStatus2

`func (o *Email) GetStatus2() string`

GetStatus2 returns the Status2 field if non-nil, zero value otherwise.

### GetStatus2Ok

`func (o *Email) GetStatus2Ok() (*string, bool)`

GetStatus2Ok returns a tuple with the Status2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2

`func (o *Email) SetStatus2(v string)`

SetStatus2 sets Status2 field to given value.

### HasStatus2

`func (o *Email) HasStatus2() bool`

HasStatus2 returns a boolean if a field has been set.

### GetPriority

`func (o *Email) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Email) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Email) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Email) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetHasInline

`func (o *Email) GetHasInline() string`

GetHasInline returns the HasInline field if non-nil, zero value otherwise.

### GetHasInlineOk

`func (o *Email) GetHasInlineOk() (*string, bool)`

GetHasInlineOk returns a tuple with the HasInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInline

`func (o *Email) SetHasInline(v string)`

SetHasInline sets HasInline field to given value.

### HasHasInline

`func (o *Email) HasHasInline() bool`

HasHasInline returns a boolean if a field has been set.

### GetToAddress

`func (o *Email) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *Email) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *Email) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *Email) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetFolderId

`func (o *Email) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Email) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Email) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Email) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetCcAddress

`func (o *Email) GetCcAddress() string`

GetCcAddress returns the CcAddress field if non-nil, zero value otherwise.

### GetCcAddressOk

`func (o *Email) GetCcAddressOk() (*string, bool)`

GetCcAddressOk returns a tuple with the CcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcAddress

`func (o *Email) SetCcAddress(v string)`

SetCcAddress sets CcAddress field to given value.

### HasCcAddress

`func (o *Email) HasCcAddress() bool`

HasCcAddress returns a boolean if a field has been set.

### GetThreadId

`func (o *Email) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Email) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Email) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *Email) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetHasAttachment

`func (o *Email) GetHasAttachment() string`

GetHasAttachment returns the HasAttachment field if non-nil, zero value otherwise.

### GetHasAttachmentOk

`func (o *Email) GetHasAttachmentOk() (*string, bool)`

GetHasAttachmentOk returns a tuple with the HasAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachment

`func (o *Email) SetHasAttachment(v string)`

SetHasAttachment sets HasAttachment field to given value.

### HasHasAttachment

`func (o *Email) HasHasAttachment() bool`

HasHasAttachment returns a boolean if a field has been set.

### GetSize

`func (o *Email) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Email) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Email) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Email) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSender

`func (o *Email) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Email) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Email) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *Email) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetReceivedTime

`func (o *Email) GetReceivedTime() int64`

GetReceivedTime returns the ReceivedTime field if non-nil, zero value otherwise.

### GetReceivedTimeOk

`func (o *Email) GetReceivedTimeOk() (*int64, bool)`

GetReceivedTimeOk returns a tuple with the ReceivedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTime

`func (o *Email) SetReceivedTime(v int64)`

SetReceivedTime sets ReceivedTime field to given value.

### HasReceivedTime

`func (o *Email) HasReceivedTime() bool`

HasReceivedTime returns a boolean if a field has been set.

### GetFromAddress

`func (o *Email) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *Email) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *Email) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *Email) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetStatus

`func (o *Email) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Email) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Email) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Email) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


