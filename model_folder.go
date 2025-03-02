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

// checks if the Folder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Folder{}

// Folder struct for Folder
type Folder struct {
	Path *string `json:"path,omitempty"`
	PreviousFolderId *string `json:"previousFolderId,omitempty"`
	IsArchived *int32 `json:"isArchived,omitempty"`
	FolderName *string `json:"folderName,omitempty"`
	ImapAccess *bool `json:"imapAccess,omitempty"`
	FolderType *string `json:"folderType,omitempty"`
	URI *string `json:"URI,omitempty"`
	FolderId *string `json:"folderId,omitempty"`
}

// NewFolder instantiates a new Folder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolder() *Folder {
	this := Folder{}
	return &this
}

// NewFolderWithDefaults instantiates a new Folder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderWithDefaults() *Folder {
	this := Folder{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Folder) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Folder) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Folder) SetPath(v string) {
	o.Path = &v
}

// GetPreviousFolderId returns the PreviousFolderId field value if set, zero value otherwise.
func (o *Folder) GetPreviousFolderId() string {
	if o == nil || IsNil(o.PreviousFolderId) {
		var ret string
		return ret
	}
	return *o.PreviousFolderId
}

// GetPreviousFolderIdOk returns a tuple with the PreviousFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetPreviousFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousFolderId) {
		return nil, false
	}
	return o.PreviousFolderId, true
}

// HasPreviousFolderId returns a boolean if a field has been set.
func (o *Folder) HasPreviousFolderId() bool {
	if o != nil && !IsNil(o.PreviousFolderId) {
		return true
	}

	return false
}

// SetPreviousFolderId gets a reference to the given string and assigns it to the PreviousFolderId field.
func (o *Folder) SetPreviousFolderId(v string) {
	o.PreviousFolderId = &v
}

// GetIsArchived returns the IsArchived field value if set, zero value otherwise.
func (o *Folder) GetIsArchived() int32 {
	if o == nil || IsNil(o.IsArchived) {
		var ret int32
		return ret
	}
	return *o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetIsArchivedOk() (*int32, bool) {
	if o == nil || IsNil(o.IsArchived) {
		return nil, false
	}
	return o.IsArchived, true
}

// HasIsArchived returns a boolean if a field has been set.
func (o *Folder) HasIsArchived() bool {
	if o != nil && !IsNil(o.IsArchived) {
		return true
	}

	return false
}

// SetIsArchived gets a reference to the given int32 and assigns it to the IsArchived field.
func (o *Folder) SetIsArchived(v int32) {
	o.IsArchived = &v
}

// GetFolderName returns the FolderName field value if set, zero value otherwise.
func (o *Folder) GetFolderName() string {
	if o == nil || IsNil(o.FolderName) {
		var ret string
		return ret
	}
	return *o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetFolderNameOk() (*string, bool) {
	if o == nil || IsNil(o.FolderName) {
		return nil, false
	}
	return o.FolderName, true
}

// HasFolderName returns a boolean if a field has been set.
func (o *Folder) HasFolderName() bool {
	if o != nil && !IsNil(o.FolderName) {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given string and assigns it to the FolderName field.
func (o *Folder) SetFolderName(v string) {
	o.FolderName = &v
}

// GetImapAccess returns the ImapAccess field value if set, zero value otherwise.
func (o *Folder) GetImapAccess() bool {
	if o == nil || IsNil(o.ImapAccess) {
		var ret bool
		return ret
	}
	return *o.ImapAccess
}

// GetImapAccessOk returns a tuple with the ImapAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetImapAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.ImapAccess) {
		return nil, false
	}
	return o.ImapAccess, true
}

// HasImapAccess returns a boolean if a field has been set.
func (o *Folder) HasImapAccess() bool {
	if o != nil && !IsNil(o.ImapAccess) {
		return true
	}

	return false
}

// SetImapAccess gets a reference to the given bool and assigns it to the ImapAccess field.
func (o *Folder) SetImapAccess(v bool) {
	o.ImapAccess = &v
}

// GetFolderType returns the FolderType field value if set, zero value otherwise.
func (o *Folder) GetFolderType() string {
	if o == nil || IsNil(o.FolderType) {
		var ret string
		return ret
	}
	return *o.FolderType
}

// GetFolderTypeOk returns a tuple with the FolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetFolderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FolderType) {
		return nil, false
	}
	return o.FolderType, true
}

// HasFolderType returns a boolean if a field has been set.
func (o *Folder) HasFolderType() bool {
	if o != nil && !IsNil(o.FolderType) {
		return true
	}

	return false
}

// SetFolderType gets a reference to the given string and assigns it to the FolderType field.
func (o *Folder) SetFolderType(v string) {
	o.FolderType = &v
}

// GetURI returns the URI field value if set, zero value otherwise.
func (o *Folder) GetURI() string {
	if o == nil || IsNil(o.URI) {
		var ret string
		return ret
	}
	return *o.URI
}

// GetURIOk returns a tuple with the URI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetURIOk() (*string, bool) {
	if o == nil || IsNil(o.URI) {
		return nil, false
	}
	return o.URI, true
}

// HasURI returns a boolean if a field has been set.
func (o *Folder) HasURI() bool {
	if o != nil && !IsNil(o.URI) {
		return true
	}

	return false
}

// SetURI gets a reference to the given string and assigns it to the URI field.
func (o *Folder) SetURI(v string) {
	o.URI = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *Folder) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *Folder) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *Folder) SetFolderId(v string) {
	o.FolderId = &v
}

func (o Folder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Folder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.PreviousFolderId) {
		toSerialize["previousFolderId"] = o.PreviousFolderId
	}
	if !IsNil(o.IsArchived) {
		toSerialize["isArchived"] = o.IsArchived
	}
	if !IsNil(o.FolderName) {
		toSerialize["folderName"] = o.FolderName
	}
	if !IsNil(o.ImapAccess) {
		toSerialize["imapAccess"] = o.ImapAccess
	}
	if !IsNil(o.FolderType) {
		toSerialize["folderType"] = o.FolderType
	}
	if !IsNil(o.URI) {
		toSerialize["URI"] = o.URI
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	return toSerialize, nil
}

type NullableFolder struct {
	value *Folder
	isSet bool
}

func (v NullableFolder) Get() *Folder {
	return v.value
}

func (v *NullableFolder) Set(val *Folder) {
	v.value = val
	v.isSet = true
}

func (v NullableFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolder(val *Folder) *NullableFolder {
	return &NullableFolder{value: val, isSet: true}
}

func (v NullableFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


