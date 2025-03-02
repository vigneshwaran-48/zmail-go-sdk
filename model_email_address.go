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

// checks if the EmailAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailAddress{}

// EmailAddress struct for EmailAddress
type EmailAddress struct {
	IsAlias *bool `json:"isAlias,omitempty"`
	IsPrimary *bool `json:"isPrimary,omitempty"`
	MailId *string `json:"mailId,omitempty"`
	IsConfirmed *bool `json:"isConfirmed,omitempty"`
}

// NewEmailAddress instantiates a new EmailAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailAddress() *EmailAddress {
	this := EmailAddress{}
	return &this
}

// NewEmailAddressWithDefaults instantiates a new EmailAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailAddressWithDefaults() *EmailAddress {
	this := EmailAddress{}
	return &this
}

// GetIsAlias returns the IsAlias field value if set, zero value otherwise.
func (o *EmailAddress) GetIsAlias() bool {
	if o == nil || IsNil(o.IsAlias) {
		var ret bool
		return ret
	}
	return *o.IsAlias
}

// GetIsAliasOk returns a tuple with the IsAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAddress) GetIsAliasOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAlias) {
		return nil, false
	}
	return o.IsAlias, true
}

// HasIsAlias returns a boolean if a field has been set.
func (o *EmailAddress) HasIsAlias() bool {
	if o != nil && !IsNil(o.IsAlias) {
		return true
	}

	return false
}

// SetIsAlias gets a reference to the given bool and assigns it to the IsAlias field.
func (o *EmailAddress) SetIsAlias(v bool) {
	o.IsAlias = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *EmailAddress) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAddress) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *EmailAddress) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *EmailAddress) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetMailId returns the MailId field value if set, zero value otherwise.
func (o *EmailAddress) GetMailId() string {
	if o == nil || IsNil(o.MailId) {
		var ret string
		return ret
	}
	return *o.MailId
}

// GetMailIdOk returns a tuple with the MailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAddress) GetMailIdOk() (*string, bool) {
	if o == nil || IsNil(o.MailId) {
		return nil, false
	}
	return o.MailId, true
}

// HasMailId returns a boolean if a field has been set.
func (o *EmailAddress) HasMailId() bool {
	if o != nil && !IsNil(o.MailId) {
		return true
	}

	return false
}

// SetMailId gets a reference to the given string and assigns it to the MailId field.
func (o *EmailAddress) SetMailId(v string) {
	o.MailId = &v
}

// GetIsConfirmed returns the IsConfirmed field value if set, zero value otherwise.
func (o *EmailAddress) GetIsConfirmed() bool {
	if o == nil || IsNil(o.IsConfirmed) {
		var ret bool
		return ret
	}
	return *o.IsConfirmed
}

// GetIsConfirmedOk returns a tuple with the IsConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAddress) GetIsConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConfirmed) {
		return nil, false
	}
	return o.IsConfirmed, true
}

// HasIsConfirmed returns a boolean if a field has been set.
func (o *EmailAddress) HasIsConfirmed() bool {
	if o != nil && !IsNil(o.IsConfirmed) {
		return true
	}

	return false
}

// SetIsConfirmed gets a reference to the given bool and assigns it to the IsConfirmed field.
func (o *EmailAddress) SetIsConfirmed(v bool) {
	o.IsConfirmed = &v
}

func (o EmailAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAlias) {
		toSerialize["isAlias"] = o.IsAlias
	}
	if !IsNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if !IsNil(o.MailId) {
		toSerialize["mailId"] = o.MailId
	}
	if !IsNil(o.IsConfirmed) {
		toSerialize["isConfirmed"] = o.IsConfirmed
	}
	return toSerialize, nil
}

type NullableEmailAddress struct {
	value *EmailAddress
	isSet bool
}

func (v NullableEmailAddress) Get() *EmailAddress {
	return v.value
}

func (v *NullableEmailAddress) Set(val *EmailAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailAddress(val *EmailAddress) *NullableEmailAddress {
	return &NullableEmailAddress{value: val, isSet: true}
}

func (v NullableEmailAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


