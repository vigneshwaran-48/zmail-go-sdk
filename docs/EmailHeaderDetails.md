# EmailHeaderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderContent** | Pointer to [**EmailHeaderDetailsHeaderContent**](EmailHeaderDetailsHeaderContent.md) |  | [optional] 
**MessageId** | Pointer to **int64** |  | [optional] 

## Methods

### NewEmailHeaderDetails

`func NewEmailHeaderDetails() *EmailHeaderDetails`

NewEmailHeaderDetails instantiates a new EmailHeaderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailHeaderDetailsWithDefaults

`func NewEmailHeaderDetailsWithDefaults() *EmailHeaderDetails`

NewEmailHeaderDetailsWithDefaults instantiates a new EmailHeaderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderContent

`func (o *EmailHeaderDetails) GetHeaderContent() EmailHeaderDetailsHeaderContent`

GetHeaderContent returns the HeaderContent field if non-nil, zero value otherwise.

### GetHeaderContentOk

`func (o *EmailHeaderDetails) GetHeaderContentOk() (*EmailHeaderDetailsHeaderContent, bool)`

GetHeaderContentOk returns a tuple with the HeaderContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderContent

`func (o *EmailHeaderDetails) SetHeaderContent(v EmailHeaderDetailsHeaderContent)`

SetHeaderContent sets HeaderContent field to given value.

### HasHeaderContent

`func (o *EmailHeaderDetails) HasHeaderContent() bool`

HasHeaderContent returns a boolean if a field has been set.

### GetMessageId

`func (o *EmailHeaderDetails) GetMessageId() int64`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EmailHeaderDetails) GetMessageIdOk() (*int64, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EmailHeaderDetails) SetMessageId(v int64)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *EmailHeaderDetails) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


