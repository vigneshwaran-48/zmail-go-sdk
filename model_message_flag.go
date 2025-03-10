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
	"fmt"
)

// MessageFlag the model 'MessageFlag'
type MessageFlag string

// List of MessageFlag
const (
	MESSAGEFLAG_INFO MessageFlag = "info"
	MESSAGEFLAG_IMPORTANT MessageFlag = "important"
	MESSAGEFLAG_FOLLOWUP MessageFlag = "followup"
	MESSAGEFLAG_FLAG_NOT_SET MessageFlag = "flag_not_set"
)

// All allowed values of MessageFlag enum
var AllowedMessageFlagEnumValues = []MessageFlag{
	"info",
	"important",
	"followup",
	"flag_not_set",
}

func (v *MessageFlag) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageFlag(value)
	for _, existing := range AllowedMessageFlagEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageFlag", value)
}

// NewMessageFlagFromValue returns a pointer to a valid MessageFlag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageFlagFromValue(v string) (*MessageFlag, error) {
	ev := MessageFlag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageFlag: valid values are %v", v, AllowedMessageFlagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageFlag) IsValid() bool {
	for _, existing := range AllowedMessageFlagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageFlag value
func (v MessageFlag) Ptr() *MessageFlag {
	return &v
}

type NullableMessageFlag struct {
	value *MessageFlag
	isSet bool
}

func (v NullableMessageFlag) Get() *MessageFlag {
	return v.value
}

func (v *NullableMessageFlag) Set(val *MessageFlag) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageFlag(val *MessageFlag) *NullableMessageFlag {
	return &NullableMessageFlag{value: val, isSet: true}
}

func (v NullableMessageFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

