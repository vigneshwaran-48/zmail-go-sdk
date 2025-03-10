# SendEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | **string** | Provide the sender&#39;s email address (associated to the authenticated account). Allowed values:- Valid email address corresponding to the authenticated account for the From field. | 
**ToAddress** | **string** | Provide the recipient&#39;s email address. Allowed values:- Valid recipient email address for the To field. | 
**CcAddress** | Pointer to **string** | Provide the recipient&#39;s email address for the Cc field. Allowed values:- Valid recipient email address for the Cc field. | [optional] 
**BccAddress** | Pointer to **string** | Provide the recipient&#39;s email address for the Bcc field. Allowed values:- Valid recipient email address for the Bcc field. | [optional] 
**Subject** | Pointer to **string** | Provide the subject of the email. | [optional] 
**Content** | Pointer to **string** | Provide the content of the email. | [optional] 
**MailFormat** | Pointer to **string** | Specify the format in which the mail needs to be sent. The value can be html or plaintext. The default value is html. | [optional] [default to "html"]
**AskReceipt** | Pointer to **string** | Specifies whether Read receipt from the recipient is requested or not. Allowed values:- yes - Requesting a read receipt. no - Not requesting a read receipt | [optional] 
**Encoding** | Pointer to **string** | Specifies the encoding that is to be used in the email content. Allowed values:- Big5, EUC-JP, EUC-KR, GB2312, ISO-2022-JP, ISO-8859-1, KOI8-R, Shift_JIS, US-ASCII, UTF-8, WINDOWS-1251, X-WINDOWS-ISO2022JP. The default value is UTF-8. | [optional] [default to "UTF-8"]
**IsSchedule** | Pointer to **bool** | Depending on whether the mail has to be scheduled or not, the value can be true - if the email should be scheduled. false - if the email should be sent immediately. | [optional] 
**ScheduleType** | Pointer to **int32** | Specifies the type of scheduling. Allowed values:- 1 - Schedules email to be sent after one hour from the time of the request. 2 - Schedules email to be sent after two hours from the time of the request. 3 - Schedules email to be sent after four hours from the time of the request. 4 - Schedules email to be sent by the morning of the next day from the time of the request. 5 - Schedules email to be sent by the afternoon of the next day from the time of the request. 6 - Schedules email to be sent on the custom date and time of your choice. | [optional] 
**TimeZone** | Pointer to **string** | Specify the timezone to schedule your email. This parameter is mandatory if scheduleType is set to value 6. For example:- GMT 5.30 (India Standard Time - Asia/Calcutta). | [optional] 
**ScheduleTime** | Pointer to **string** | Specify the date and time you want to schedule your email. This parameter is mandatory if scheduleType is set to value 6. Format:- MM/DD/YYYY HH:MM:SS. For example:- 09/15/2023 14:30:28 | [optional] 
**Attachments** | Pointer to [**[]SendEmailRequestAttachmentsInner**](SendEmailRequestAttachmentsInner.md) |  | [optional] 
**Mode** | Pointer to [**SendEmailMode**](SendEmailMode.md) |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 

## Methods

### NewSendEmailRequest

`func NewSendEmailRequest(fromAddress string, toAddress string, ) *SendEmailRequest`

NewSendEmailRequest instantiates a new SendEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailRequestWithDefaults

`func NewSendEmailRequestWithDefaults() *SendEmailRequest`

NewSendEmailRequestWithDefaults instantiates a new SendEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *SendEmailRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *SendEmailRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *SendEmailRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *SendEmailRequest) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *SendEmailRequest) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *SendEmailRequest) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetCcAddress

`func (o *SendEmailRequest) GetCcAddress() string`

GetCcAddress returns the CcAddress field if non-nil, zero value otherwise.

### GetCcAddressOk

`func (o *SendEmailRequest) GetCcAddressOk() (*string, bool)`

GetCcAddressOk returns a tuple with the CcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcAddress

`func (o *SendEmailRequest) SetCcAddress(v string)`

SetCcAddress sets CcAddress field to given value.

### HasCcAddress

`func (o *SendEmailRequest) HasCcAddress() bool`

HasCcAddress returns a boolean if a field has been set.

### GetBccAddress

`func (o *SendEmailRequest) GetBccAddress() string`

GetBccAddress returns the BccAddress field if non-nil, zero value otherwise.

### GetBccAddressOk

`func (o *SendEmailRequest) GetBccAddressOk() (*string, bool)`

GetBccAddressOk returns a tuple with the BccAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccAddress

`func (o *SendEmailRequest) SetBccAddress(v string)`

SetBccAddress sets BccAddress field to given value.

### HasBccAddress

`func (o *SendEmailRequest) HasBccAddress() bool`

HasBccAddress returns a boolean if a field has been set.

### GetSubject

`func (o *SendEmailRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendEmailRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendEmailRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SendEmailRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContent

`func (o *SendEmailRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SendEmailRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SendEmailRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SendEmailRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMailFormat

`func (o *SendEmailRequest) GetMailFormat() string`

GetMailFormat returns the MailFormat field if non-nil, zero value otherwise.

### GetMailFormatOk

`func (o *SendEmailRequest) GetMailFormatOk() (*string, bool)`

GetMailFormatOk returns a tuple with the MailFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailFormat

`func (o *SendEmailRequest) SetMailFormat(v string)`

SetMailFormat sets MailFormat field to given value.

### HasMailFormat

`func (o *SendEmailRequest) HasMailFormat() bool`

HasMailFormat returns a boolean if a field has been set.

### GetAskReceipt

`func (o *SendEmailRequest) GetAskReceipt() string`

GetAskReceipt returns the AskReceipt field if non-nil, zero value otherwise.

### GetAskReceiptOk

`func (o *SendEmailRequest) GetAskReceiptOk() (*string, bool)`

GetAskReceiptOk returns a tuple with the AskReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskReceipt

`func (o *SendEmailRequest) SetAskReceipt(v string)`

SetAskReceipt sets AskReceipt field to given value.

### HasAskReceipt

`func (o *SendEmailRequest) HasAskReceipt() bool`

HasAskReceipt returns a boolean if a field has been set.

### GetEncoding

`func (o *SendEmailRequest) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *SendEmailRequest) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *SendEmailRequest) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *SendEmailRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetIsSchedule

`func (o *SendEmailRequest) GetIsSchedule() bool`

GetIsSchedule returns the IsSchedule field if non-nil, zero value otherwise.

### GetIsScheduleOk

`func (o *SendEmailRequest) GetIsScheduleOk() (*bool, bool)`

GetIsScheduleOk returns a tuple with the IsSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSchedule

`func (o *SendEmailRequest) SetIsSchedule(v bool)`

SetIsSchedule sets IsSchedule field to given value.

### HasIsSchedule

`func (o *SendEmailRequest) HasIsSchedule() bool`

HasIsSchedule returns a boolean if a field has been set.

### GetScheduleType

`func (o *SendEmailRequest) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *SendEmailRequest) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *SendEmailRequest) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *SendEmailRequest) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetTimeZone

`func (o *SendEmailRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SendEmailRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SendEmailRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SendEmailRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetScheduleTime

`func (o *SendEmailRequest) GetScheduleTime() string`

GetScheduleTime returns the ScheduleTime field if non-nil, zero value otherwise.

### GetScheduleTimeOk

`func (o *SendEmailRequest) GetScheduleTimeOk() (*string, bool)`

GetScheduleTimeOk returns a tuple with the ScheduleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTime

`func (o *SendEmailRequest) SetScheduleTime(v string)`

SetScheduleTime sets ScheduleTime field to given value.

### HasScheduleTime

`func (o *SendEmailRequest) HasScheduleTime() bool`

HasScheduleTime returns a boolean if a field has been set.

### GetAttachments

`func (o *SendEmailRequest) GetAttachments() []SendEmailRequestAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SendEmailRequest) GetAttachmentsOk() (*[]SendEmailRequestAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SendEmailRequest) SetAttachments(v []SendEmailRequestAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SendEmailRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetMode

`func (o *SendEmailRequest) GetMode() SendEmailMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SendEmailRequest) GetModeOk() (*SendEmailMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SendEmailRequest) SetMode(v SendEmailMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SendEmailRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAction

`func (o *SendEmailRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SendEmailRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SendEmailRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SendEmailRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


