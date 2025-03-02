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

// checks if the MailForward type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailForward{}

// MailForward struct for MailForward
type MailForward struct {
	MailForwardTo *string `json:"mailForwardTo,omitempty"`
	Type *string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewMailForward instantiates a new MailForward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailForward() *MailForward {
	this := MailForward{}
	return &this
}

// NewMailForwardWithDefaults instantiates a new MailForward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailForwardWithDefaults() *MailForward {
	this := MailForward{}
	return &this
}

// GetMailForwardTo returns the MailForwardTo field value if set, zero value otherwise.
func (o *MailForward) GetMailForwardTo() string {
	if o == nil || IsNil(o.MailForwardTo) {
		var ret string
		return ret
	}
	return *o.MailForwardTo
}

// GetMailForwardToOk returns a tuple with the MailForwardTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailForward) GetMailForwardToOk() (*string, bool) {
	if o == nil || IsNil(o.MailForwardTo) {
		return nil, false
	}
	return o.MailForwardTo, true
}

// HasMailForwardTo returns a boolean if a field has been set.
func (o *MailForward) HasMailForwardTo() bool {
	if o != nil && !IsNil(o.MailForwardTo) {
		return true
	}

	return false
}

// SetMailForwardTo gets a reference to the given string and assigns it to the MailForwardTo field.
func (o *MailForward) SetMailForwardTo(v string) {
	o.MailForwardTo = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MailForward) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailForward) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MailForward) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MailForward) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MailForward) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailForward) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MailForward) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MailForward) SetStatus(v string) {
	o.Status = &v
}

func (o MailForward) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailForward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MailForwardTo) {
		toSerialize["mailForwardTo"] = o.MailForwardTo
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableMailForward struct {
	value *MailForward
	isSet bool
}

func (v NullableMailForward) Get() *MailForward {
	return v.value
}

func (v *NullableMailForward) Set(val *MailForward) {
	v.value = val
	v.isSet = true
}

func (v NullableMailForward) IsSet() bool {
	return v.isSet
}

func (v *NullableMailForward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailForward(val *MailForward) *NullableMailForward {
	return &NullableMailForward{value: val, isSet: true}
}

func (v NullableMailForward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailForward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


