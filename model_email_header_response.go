/*
Zoho Mail API

This is a Zoho Mail API OpenAPI 3.0 specification

API version: 1.0.0
Contact: vigneshwaran4817@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zmail

import (
	"encoding/json"
)

// checks if the EmailHeaderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailHeaderResponse{}

// EmailHeaderResponse struct for EmailHeaderResponse
type EmailHeaderResponse struct {
	Status *Status `json:"status,omitempty"`
	Data *EmailHeaderDetails `json:"data,omitempty"`
}

// NewEmailHeaderResponse instantiates a new EmailHeaderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailHeaderResponse() *EmailHeaderResponse {
	this := EmailHeaderResponse{}
	return &this
}

// NewEmailHeaderResponseWithDefaults instantiates a new EmailHeaderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailHeaderResponseWithDefaults() *EmailHeaderResponse {
	this := EmailHeaderResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmailHeaderResponse) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailHeaderResponse) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmailHeaderResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *EmailHeaderResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailHeaderResponse) GetData() EmailHeaderDetails {
	if o == nil || IsNil(o.Data) {
		var ret EmailHeaderDetails
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailHeaderResponse) GetDataOk() (*EmailHeaderDetails, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailHeaderResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EmailHeaderDetails and assigns it to the Data field.
func (o *EmailHeaderResponse) SetData(v EmailHeaderDetails) {
	o.Data = &v
}

func (o EmailHeaderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailHeaderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEmailHeaderResponse struct {
	value *EmailHeaderResponse
	isSet bool
}

func (v NullableEmailHeaderResponse) Get() *EmailHeaderResponse {
	return v.value
}

func (v *NullableEmailHeaderResponse) Set(val *EmailHeaderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailHeaderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailHeaderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailHeaderResponse(val *EmailHeaderResponse) *NullableEmailHeaderResponse {
	return &NullableEmailHeaderResponse{value: val, isSet: true}
}

func (v NullableEmailHeaderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailHeaderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


