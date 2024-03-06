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

// checks if the ManualJournals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualJournals{}

// ManualJournals struct for ManualJournals
type ManualJournals struct {
	ManualJournals []ManualJournal `json:"ManualJournals,omitempty"`
}

// NewManualJournals instantiates a new ManualJournals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualJournals() *ManualJournals {
	this := ManualJournals{}
	return &this
}

// NewManualJournalsWithDefaults instantiates a new ManualJournals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualJournalsWithDefaults() *ManualJournals {
	this := ManualJournals{}
	return &this
}

// GetManualJournals returns the ManualJournals field value if set, zero value otherwise.
func (o *ManualJournals) GetManualJournals() []ManualJournal {
	if o == nil || IsNil(o.ManualJournals) {
		var ret []ManualJournal
		return ret
	}
	return o.ManualJournals
}

// GetManualJournalsOk returns a tuple with the ManualJournals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournals) GetManualJournalsOk() ([]ManualJournal, bool) {
	if o == nil || IsNil(o.ManualJournals) {
		return nil, false
	}
	return o.ManualJournals, true
}

// HasManualJournalsField returns a boolean if a field has been set.
func (o *ManualJournals) HasManualJournalsField() bool {
	if o != nil && !IsNil(o.ManualJournals) {
		return true
	}

	return false
}

// SetManualJournals gets a reference to the given []ManualJournal and assigns it to the ManualJournals field.
func (o *ManualJournals) SetManualJournals(v []ManualJournal) {
	o.ManualJournals = v
}

func (o ManualJournals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualJournals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManualJournals) {
		toSerialize["ManualJournals"] = o.ManualJournals
	}
	return toSerialize, nil
}

type NullableManualJournals struct {
	value *ManualJournals
	isSet bool
}

func (v NullableManualJournals) Get() *ManualJournals {
	return v.value
}

func (v *NullableManualJournals) Set(val *ManualJournals) {
	v.value = val
	v.isSet = true
}

func (v NullableManualJournals) IsSet() bool {
	return v.isSet
}

func (v *NullableManualJournals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualJournals(val *ManualJournals) *NullableManualJournals {
	return &NullableManualJournals{value: val, isSet: true}
}

func (v NullableManualJournals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualJournals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


