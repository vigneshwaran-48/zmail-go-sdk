# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AccountDisplayName** | Pointer to **string** |  | [optional] 
**IncomingUserName** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**MailboxStatus** | Pointer to **string** |  | [optional] 
**ImapAccessEnabled** | Pointer to **bool** |  | [optional] 
**PopAccessEnabled** | Pointer to **bool** |  | [optional] 
**SmtpStatus** | Pointer to **bool** |  | [optional] 
**MailForward** | Pointer to [**[]MailForward**](MailForward.md) |  | [optional] 
**EmailAddress** | Pointer to [**[]EmailAddress**](EmailAddress.md) |  | [optional] 
**VacationResponse** | Pointer to [**[]VacationResponse**](VacationResponse.md) |  | [optional] 
**SendMailDetails** | Pointer to [**[]SendMailDetails**](SendMailDetails.md) |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**URI** | Pointer to **string** |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Account) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Account) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Account) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Account) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountDisplayName

`func (o *Account) GetAccountDisplayName() string`

GetAccountDisplayName returns the AccountDisplayName field if non-nil, zero value otherwise.

### GetAccountDisplayNameOk

`func (o *Account) GetAccountDisplayNameOk() (*string, bool)`

GetAccountDisplayNameOk returns a tuple with the AccountDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDisplayName

`func (o *Account) SetAccountDisplayName(v string)`

SetAccountDisplayName sets AccountDisplayName field to given value.

### HasAccountDisplayName

`func (o *Account) HasAccountDisplayName() bool`

HasAccountDisplayName returns a boolean if a field has been set.

### GetIncomingUserName

`func (o *Account) GetIncomingUserName() string`

GetIncomingUserName returns the IncomingUserName field if non-nil, zero value otherwise.

### GetIncomingUserNameOk

`func (o *Account) GetIncomingUserNameOk() (*string, bool)`

GetIncomingUserNameOk returns a tuple with the IncomingUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingUserName

`func (o *Account) SetIncomingUserName(v string)`

SetIncomingUserName sets IncomingUserName field to given value.

### HasIncomingUserName

`func (o *Account) HasIncomingUserName() bool`

HasIncomingUserName returns a boolean if a field has been set.

### GetAccountName

`func (o *Account) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Account) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Account) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Account) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Account) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMailboxStatus

`func (o *Account) GetMailboxStatus() string`

GetMailboxStatus returns the MailboxStatus field if non-nil, zero value otherwise.

### GetMailboxStatusOk

`func (o *Account) GetMailboxStatusOk() (*string, bool)`

GetMailboxStatusOk returns a tuple with the MailboxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxStatus

`func (o *Account) SetMailboxStatus(v string)`

SetMailboxStatus sets MailboxStatus field to given value.

### HasMailboxStatus

`func (o *Account) HasMailboxStatus() bool`

HasMailboxStatus returns a boolean if a field has been set.

### GetImapAccessEnabled

`func (o *Account) GetImapAccessEnabled() bool`

GetImapAccessEnabled returns the ImapAccessEnabled field if non-nil, zero value otherwise.

### GetImapAccessEnabledOk

`func (o *Account) GetImapAccessEnabledOk() (*bool, bool)`

GetImapAccessEnabledOk returns a tuple with the ImapAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapAccessEnabled

`func (o *Account) SetImapAccessEnabled(v bool)`

SetImapAccessEnabled sets ImapAccessEnabled field to given value.

### HasImapAccessEnabled

`func (o *Account) HasImapAccessEnabled() bool`

HasImapAccessEnabled returns a boolean if a field has been set.

### GetPopAccessEnabled

`func (o *Account) GetPopAccessEnabled() bool`

GetPopAccessEnabled returns the PopAccessEnabled field if non-nil, zero value otherwise.

### GetPopAccessEnabledOk

`func (o *Account) GetPopAccessEnabledOk() (*bool, bool)`

GetPopAccessEnabledOk returns a tuple with the PopAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopAccessEnabled

`func (o *Account) SetPopAccessEnabled(v bool)`

SetPopAccessEnabled sets PopAccessEnabled field to given value.

### HasPopAccessEnabled

`func (o *Account) HasPopAccessEnabled() bool`

HasPopAccessEnabled returns a boolean if a field has been set.

### GetSmtpStatus

`func (o *Account) GetSmtpStatus() bool`

GetSmtpStatus returns the SmtpStatus field if non-nil, zero value otherwise.

### GetSmtpStatusOk

`func (o *Account) GetSmtpStatusOk() (*bool, bool)`

GetSmtpStatusOk returns a tuple with the SmtpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpStatus

`func (o *Account) SetSmtpStatus(v bool)`

SetSmtpStatus sets SmtpStatus field to given value.

### HasSmtpStatus

`func (o *Account) HasSmtpStatus() bool`

HasSmtpStatus returns a boolean if a field has been set.

### GetMailForward

`func (o *Account) GetMailForward() []MailForward`

GetMailForward returns the MailForward field if non-nil, zero value otherwise.

### GetMailForwardOk

`func (o *Account) GetMailForwardOk() (*[]MailForward, bool)`

GetMailForwardOk returns a tuple with the MailForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailForward

`func (o *Account) SetMailForward(v []MailForward)`

SetMailForward sets MailForward field to given value.

### HasMailForward

`func (o *Account) HasMailForward() bool`

HasMailForward returns a boolean if a field has been set.

### GetEmailAddress

`func (o *Account) GetEmailAddress() []EmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Account) GetEmailAddressOk() (*[]EmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Account) SetEmailAddress(v []EmailAddress)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Account) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetVacationResponse

`func (o *Account) GetVacationResponse() []VacationResponse`

GetVacationResponse returns the VacationResponse field if non-nil, zero value otherwise.

### GetVacationResponseOk

`func (o *Account) GetVacationResponseOk() (*[]VacationResponse, bool)`

GetVacationResponseOk returns a tuple with the VacationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationResponse

`func (o *Account) SetVacationResponse(v []VacationResponse)`

SetVacationResponse sets VacationResponse field to given value.

### HasVacationResponse

`func (o *Account) HasVacationResponse() bool`

HasVacationResponse returns a boolean if a field has been set.

### GetSendMailDetails

`func (o *Account) GetSendMailDetails() []SendMailDetails`

GetSendMailDetails returns the SendMailDetails field if non-nil, zero value otherwise.

### GetSendMailDetailsOk

`func (o *Account) GetSendMailDetailsOk() (*[]SendMailDetails, bool)`

GetSendMailDetailsOk returns a tuple with the SendMailDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMailDetails

`func (o *Account) SetSendMailDetails(v []SendMailDetails)`

SetSendMailDetails sets SendMailDetails field to given value.

### HasSendMailDetails

`func (o *Account) HasSendMailDetails() bool`

HasSendMailDetails returns a boolean if a field has been set.

### GetAddress

`func (o *Account) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Account) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Account) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Account) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetURI

`func (o *Account) GetURI() string`

GetURI returns the URI field if non-nil, zero value otherwise.

### GetURIOk

`func (o *Account) GetURIOk() (*string, bool)`

GetURIOk returns a tuple with the URI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURI

`func (o *Account) SetURI(v string)`

SetURI sets URI field to given value.

### HasURI

`func (o *Account) HasURI() bool`

HasURI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


