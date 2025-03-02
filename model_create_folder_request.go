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
	"bytes"
	"fmt"
)

// checks if the CreateFolderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFolderRequest{}

// CreateFolderRequest struct for CreateFolderRequest
type CreateFolderRequest struct {
	// The name of the folder to be created.
	FolderName string `json:"folderName"`
	// The ID of the parent folder under which the new folder is created.
	ParentFolderId *string `json:"parentFolderId,omitempty"`
	// The path of the parent folder under which the new folder is created.
	ParentFolderPath *string `json:"parentFolderPath,omitempty"`
}

type _CreateFolderRequest CreateFolderRequest

// NewCreateFolderRequest instantiates a new CreateFolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFolderRequest(folderName string) *CreateFolderRequest {
	this := CreateFolderRequest{}
	this.FolderName = folderName
	return &this
}

// NewCreateFolderRequestWithDefaults instantiates a new CreateFolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFolderRequestWithDefaults() *CreateFolderRequest {
	this := CreateFolderRequest{}
	return &this
}

// GetFolderName returns the FolderName field value
func (o *CreateFolderRequest) GetFolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value
// and a boolean to check if the value has been set.
func (o *CreateFolderRequest) GetFolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderName, true
}

// SetFolderName sets field value
func (o *CreateFolderRequest) SetFolderName(v string) {
	o.FolderName = v
}

// GetParentFolderId returns the ParentFolderId field value if set, zero value otherwise.
func (o *CreateFolderRequest) GetParentFolderId() string {
	if o == nil || IsNil(o.ParentFolderId) {
		var ret string
		return ret
	}
	return *o.ParentFolderId
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFolderRequest) GetParentFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentFolderId) {
		return nil, false
	}
	return o.ParentFolderId, true
}

// HasParentFolderId returns a boolean if a field has been set.
func (o *CreateFolderRequest) HasParentFolderId() bool {
	if o != nil && !IsNil(o.ParentFolderId) {
		return true
	}

	return false
}

// SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.
func (o *CreateFolderRequest) SetParentFolderId(v string) {
	o.ParentFolderId = &v
}

// GetParentFolderPath returns the ParentFolderPath field value if set, zero value otherwise.
func (o *CreateFolderRequest) GetParentFolderPath() string {
	if o == nil || IsNil(o.ParentFolderPath) {
		var ret string
		return ret
	}
	return *o.ParentFolderPath
}

// GetParentFolderPathOk returns a tuple with the ParentFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFolderRequest) GetParentFolderPathOk() (*string, bool) {
	if o == nil || IsNil(o.ParentFolderPath) {
		return nil, false
	}
	return o.ParentFolderPath, true
}

// HasParentFolderPath returns a boolean if a field has been set.
func (o *CreateFolderRequest) HasParentFolderPath() bool {
	if o != nil && !IsNil(o.ParentFolderPath) {
		return true
	}

	return false
}

// SetParentFolderPath gets a reference to the given string and assigns it to the ParentFolderPath field.
func (o *CreateFolderRequest) SetParentFolderPath(v string) {
	o.ParentFolderPath = &v
}

func (o CreateFolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFolderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folderName"] = o.FolderName
	if !IsNil(o.ParentFolderId) {
		toSerialize["parentFolderId"] = o.ParentFolderId
	}
	if !IsNil(o.ParentFolderPath) {
		toSerialize["parentFolderPath"] = o.ParentFolderPath
	}
	return toSerialize, nil
}

func (o *CreateFolderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"folderName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateFolderRequest := _CreateFolderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFolderRequest)

	if err != nil {
		return err
	}

	*o = CreateFolderRequest(varCreateFolderRequest)

	return err
}

type NullableCreateFolderRequest struct {
	value *CreateFolderRequest
	isSet bool
}

func (v NullableCreateFolderRequest) Get() *CreateFolderRequest {
	return v.value
}

func (v *NullableCreateFolderRequest) Set(val *CreateFolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFolderRequest(val *CreateFolderRequest) *NullableCreateFolderRequest {
	return &NullableCreateFolderRequest{value: val, isSet: true}
}

func (v NullableCreateFolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


