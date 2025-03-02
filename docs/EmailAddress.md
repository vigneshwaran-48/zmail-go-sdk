# EmailAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAlias** | Pointer to **bool** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**MailId** | Pointer to **string** |  | [optional] 
**IsConfirmed** | Pointer to **bool** |  | [optional] 

## Methods

### NewEmailAddress

`func NewEmailAddress() *EmailAddress`

NewEmailAddress instantiates a new EmailAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAddressWithDefaults

`func NewEmailAddressWithDefaults() *EmailAddress`

NewEmailAddressWithDefaults instantiates a new EmailAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAlias

`func (o *EmailAddress) GetIsAlias() bool`

GetIsAlias returns the IsAlias field if non-nil, zero value otherwise.

### GetIsAliasOk

`func (o *EmailAddress) GetIsAliasOk() (*bool, bool)`

GetIsAliasOk returns a tuple with the IsAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlias

`func (o *EmailAddress) SetIsAlias(v bool)`

SetIsAlias sets IsAlias field to given value.

### HasIsAlias

`func (o *EmailAddress) HasIsAlias() bool`

HasIsAlias returns a boolean if a field has been set.

### GetIsPrimary

`func (o *EmailAddress) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *EmailAddress) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *EmailAddress) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *EmailAddress) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetMailId

`func (o *EmailAddress) GetMailId() string`

GetMailId returns the MailId field if non-nil, zero value otherwise.

### GetMailIdOk

`func (o *EmailAddress) GetMailIdOk() (*string, bool)`

GetMailIdOk returns a tuple with the MailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailId

`func (o *EmailAddress) SetMailId(v string)`

SetMailId sets MailId field to given value.

### HasMailId

`func (o *EmailAddress) HasMailId() bool`

HasMailId returns a boolean if a field has been set.

### GetIsConfirmed

`func (o *EmailAddress) GetIsConfirmed() bool`

GetIsConfirmed returns the IsConfirmed field if non-nil, zero value otherwise.

### GetIsConfirmedOk

`func (o *EmailAddress) GetIsConfirmedOk() (*bool, bool)`

GetIsConfirmedOk returns a tuple with the IsConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfirmed

`func (o *EmailAddress) SetIsConfirmed(v bool)`

SetIsConfirmed sets IsConfirmed field to given value.

### HasIsConfirmed

`func (o *EmailAddress) HasIsConfirmed() bool`

HasIsConfirmed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


