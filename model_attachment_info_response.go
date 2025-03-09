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

// checks if the AttachmentInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentInfoResponse{}

// AttachmentInfoResponse struct for AttachmentInfoResponse
type AttachmentInfoResponse struct {
	Status *Status `json:"status,omitempty"`
	Data *AttachmentInfoResponseData `json:"data,omitempty"`
}

// NewAttachmentInfoResponse instantiates a new AttachmentInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentInfoResponse() *AttachmentInfoResponse {
	this := AttachmentInfoResponse{}
	return &this
}

// NewAttachmentInfoResponseWithDefaults instantiates a new AttachmentInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentInfoResponseWithDefaults() *AttachmentInfoResponse {
	this := AttachmentInfoResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AttachmentInfoResponse) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentInfoResponse) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AttachmentInfoResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *AttachmentInfoResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AttachmentInfoResponse) GetData() AttachmentInfoResponseData {
	if o == nil || IsNil(o.Data) {
		var ret AttachmentInfoResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentInfoResponse) GetDataOk() (*AttachmentInfoResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AttachmentInfoResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AttachmentInfoResponseData and assigns it to the Data field.
func (o *AttachmentInfoResponse) SetData(v AttachmentInfoResponseData) {
	o.Data = &v
}

func (o AttachmentInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAttachmentInfoResponse struct {
	value *AttachmentInfoResponse
	isSet bool
}

func (v NullableAttachmentInfoResponse) Get() *AttachmentInfoResponse {
	return v.value
}

func (v *NullableAttachmentInfoResponse) Set(val *AttachmentInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentInfoResponse(val *AttachmentInfoResponse) *NullableAttachmentInfoResponse {
	return &NullableAttachmentInfoResponse{value: val, isSet: true}
}

func (v NullableAttachmentInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


