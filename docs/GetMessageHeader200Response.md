# GetMessageHeader200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**EmailHeaderRawResponseData**](EmailHeaderRawResponseData.md) |  | [optional] 

## Methods

### NewGetMessageHeader200Response

`func NewGetMessageHeader200Response() *GetMessageHeader200Response`

NewGetMessageHeader200Response instantiates a new GetMessageHeader200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessageHeader200ResponseWithDefaults

`func NewGetMessageHeader200ResponseWithDefaults() *GetMessageHeader200Response`

NewGetMessageHeader200ResponseWithDefaults instantiates a new GetMessageHeader200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetMessageHeader200Response) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMessageHeader200Response) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMessageHeader200Response) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMessageHeader200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetMessageHeader200Response) GetData() EmailHeaderRawResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMessageHeader200Response) GetDataOk() (*EmailHeaderRawResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMessageHeader200Response) SetData(v EmailHeaderRawResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetMessageHeader200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


