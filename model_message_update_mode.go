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

// MessageUpdateMode the model 'MessageUpdateMode'
type MessageUpdateMode string

// List of MessageUpdateMode
const (
	MARK_AS_READ MessageUpdateMode = "markAsRead"
	MARK_AS_UNREAD MessageUpdateMode = "markAsUnread"
	MOVE_MESSAGE MessageUpdateMode = "moveMessage"
	SET_FLAG MessageUpdateMode = "setFlag"
	APPLY_LABEL MessageUpdateMode = "applyLabel"
	REMOVE_LABEL MessageUpdateMode = "removeLabel"
	REMOVE_ALL_LABELS MessageUpdateMode = "removeAllLabels"
	ARCHIVE_MAILS MessageUpdateMode = "archiveMails"
	UN_ARCHIVE_MAILS MessageUpdateMode = "unArchiveMails"
	MOVE_TO_SPAM MessageUpdateMode = "moveToSpam"
	MARK_NOT_SPAM MessageUpdateMode = "markNotSpam"
)

// All allowed values of MessageUpdateMode enum
var AllowedMessageUpdateModeEnumValues = []MessageUpdateMode{
	"markAsRead",
	"markAsUnread",
	"moveMessage",
	"setFlag",
	"applyLabel",
	"removeLabel",
	"removeAllLabels",
	"archiveMails",
	"unArchiveMails",
	"moveToSpam",
	"markNotSpam",
}

func (v *MessageUpdateMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageUpdateMode(value)
	for _, existing := range AllowedMessageUpdateModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageUpdateMode", value)
}

// NewMessageUpdateModeFromValue returns a pointer to a valid MessageUpdateMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageUpdateModeFromValue(v string) (*MessageUpdateMode, error) {
	ev := MessageUpdateMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageUpdateMode: valid values are %v", v, AllowedMessageUpdateModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageUpdateMode) IsValid() bool {
	for _, existing := range AllowedMessageUpdateModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageUpdateMode value
func (v MessageUpdateMode) Ptr() *MessageUpdateMode {
	return &v
}

type NullableMessageUpdateMode struct {
	value *MessageUpdateMode
	isSet bool
}

func (v NullableMessageUpdateMode) Get() *MessageUpdateMode {
	return v.value
}

func (v *NullableMessageUpdateMode) Set(val *MessageUpdateMode) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageUpdateMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageUpdateMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageUpdateMode(val *MessageUpdateMode) *NullableMessageUpdateMode {
	return &NullableMessageUpdateMode{value: val, isSet: true}
}

func (v NullableMessageUpdateMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageUpdateMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

