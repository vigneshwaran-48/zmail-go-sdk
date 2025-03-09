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

// checks if the EmailHeaderRawResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailHeaderRawResponseData{}

// EmailHeaderRawResponseData struct for EmailHeaderRawResponseData
type EmailHeaderRawResponseData struct {
	HeaderContent *string `json:"headerContent,omitempty"`
	MessageId *string `json:"messageId,omitempty"`
}

// NewEmailHeaderRawResponseData instantiates a new EmailHeaderRawResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailHeaderRawResponseData() *EmailHeaderRawResponseData {
	this := EmailHeaderRawResponseData{}
	return &this
}

// NewEmailHeaderRawResponseDataWithDefaults instantiates a new EmailHeaderRawResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailHeaderRawResponseDataWithDefaults() *EmailHeaderRawResponseData {
	this := EmailHeaderRawResponseData{}
	return &this
}

// GetHeaderContent returns the HeaderContent field value if set, zero value otherwise.
func (o *EmailHeaderRawResponseData) GetHeaderContent() string {
	if o == nil || IsNil(o.HeaderContent) {
		var ret string
		return ret
	}
	return *o.HeaderContent
}

// GetHeaderContentOk returns a tuple with the HeaderContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailHeaderRawResponseData) GetHeaderContentOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderContent) {
		return nil, false
	}
	return o.HeaderContent, true
}

// HasHeaderContent returns a boolean if a field has been set.
func (o *EmailHeaderRawResponseData) HasHeaderContent() bool {
	if o != nil && !IsNil(o.HeaderContent) {
		return true
	}

	return false
}

// SetHeaderContent gets a reference to the given string and assigns it to the HeaderContent field.
func (o *EmailHeaderRawResponseData) SetHeaderContent(v string) {
	o.HeaderContent = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *EmailHeaderRawResponseData) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailHeaderRawResponseData) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *EmailHeaderRawResponseData) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *EmailHeaderRawResponseData) SetMessageId(v string) {
	o.MessageId = &v
}

func (o EmailHeaderRawResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailHeaderRawResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeaderContent) {
		toSerialize["headerContent"] = o.HeaderContent
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	return toSerialize, nil
}

type NullableEmailHeaderRawResponseData struct {
	value *EmailHeaderRawResponseData
	isSet bool
}

func (v NullableEmailHeaderRawResponseData) Get() *EmailHeaderRawResponseData {
	return v.value
}

func (v *NullableEmailHeaderRawResponseData) Set(val *EmailHeaderRawResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailHeaderRawResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailHeaderRawResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailHeaderRawResponseData(val *EmailHeaderRawResponseData) *NullableEmailHeaderRawResponseData {
	return &NullableEmailHeaderRawResponseData{value: val, isSet: true}
}

func (v NullableEmailHeaderRawResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailHeaderRawResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


