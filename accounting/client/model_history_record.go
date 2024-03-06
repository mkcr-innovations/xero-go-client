/*
Xero Accounting API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.1
Contact: api@xero.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the HistoryRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryRecord{}

// HistoryRecord struct for HistoryRecord
type HistoryRecord struct {
	// details
	Details *string `json:"Details,omitempty"`
	// Name of branding theme
	Changes *string `json:"Changes,omitempty"`
	// has a value of 0
	User *string `json:"User,omitempty"`
	// UTC timestamp of creation date of branding theme
	DateUTC *string `json:"DateUTC,omitempty"`
}

// NewHistoryRecord instantiates a new HistoryRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryRecord() *HistoryRecord {
	this := HistoryRecord{}
	return &this
}

// NewHistoryRecordWithDefaults instantiates a new HistoryRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryRecordWithDefaults() *HistoryRecord {
	this := HistoryRecord{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *HistoryRecord) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryRecord) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetailsField returns a boolean if a field has been set.
func (o *HistoryRecord) HasDetailsField() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *HistoryRecord) SetDetails(v string) {
	o.Details = &v
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *HistoryRecord) GetChanges() string {
	if o == nil || IsNil(o.Changes) {
		var ret string
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryRecord) GetChangesOk() (*string, bool) {
	if o == nil || IsNil(o.Changes) {
		return nil, false
	}
	return o.Changes, true
}

// HasChangesField returns a boolean if a field has been set.
func (o *HistoryRecord) HasChangesField() bool {
	if o != nil && !IsNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given string and assigns it to the Changes field.
func (o *HistoryRecord) SetChanges(v string) {
	o.Changes = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *HistoryRecord) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryRecord) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUserField returns a boolean if a field has been set.
func (o *HistoryRecord) HasUserField() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *HistoryRecord) SetUser(v string) {
	o.User = &v
}

// GetDateUTC returns the DateUTC field value if set, zero value otherwise.
func (o *HistoryRecord) GetDateUTC() string {
	if o == nil || IsNil(o.DateUTC) {
		var ret string
		return ret
	}
	return *o.DateUTC
}

// GetDateUTCOk returns a tuple with the DateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryRecord) GetDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.DateUTC) {
		return nil, false
	}
	return o.DateUTC, true
}

// HasDateUTCField returns a boolean if a field has been set.
func (o *HistoryRecord) HasDateUTCField() bool {
	if o != nil && !IsNil(o.DateUTC) {
		return true
	}

	return false
}

// SetDateUTC gets a reference to the given string and assigns it to the DateUTC field.
func (o *HistoryRecord) SetDateUTC(v string) {
	o.DateUTC = &v
}

func (o HistoryRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["Details"] = o.Details
	}
	if !IsNil(o.Changes) {
		toSerialize["Changes"] = o.Changes
	}
	if !IsNil(o.User) {
		toSerialize["User"] = o.User
	}
	if !IsNil(o.DateUTC) {
		toSerialize["DateUTC"] = o.DateUTC
	}
	return toSerialize, nil
}

type NullableHistoryRecord struct {
	value *HistoryRecord
	isSet bool
}

func (v NullableHistoryRecord) Get() *HistoryRecord {
	return v.value
}

func (v *NullableHistoryRecord) Set(val *HistoryRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryRecord(val *HistoryRecord) *NullableHistoryRecord {
	return &NullableHistoryRecord{value: val, isSet: true}
}

func (v NullableHistoryRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

