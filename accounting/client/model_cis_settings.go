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

// checks if the CISSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CISSettings{}

// CISSettings struct for CISSettings
type CISSettings struct {
	CISSettings []CISSetting `json:"CISSettings,omitempty"`
}

// NewCISSettings instantiates a new CISSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCISSettings() *CISSettings {
	this := CISSettings{}
	return &this
}

// NewCISSettingsWithDefaults instantiates a new CISSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCISSettingsWithDefaults() *CISSettings {
	this := CISSettings{}
	return &this
}

// GetCISSettings returns the CISSettings field value if set, zero value otherwise.
func (o *CISSettings) GetCISSettings() []CISSetting {
	if o == nil || IsNil(o.CISSettings) {
		var ret []CISSetting
		return ret
	}
	return o.CISSettings
}

// GetCISSettingsOk returns a tuple with the CISSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CISSettings) GetCISSettingsOk() ([]CISSetting, bool) {
	if o == nil || IsNil(o.CISSettings) {
		return nil, false
	}
	return o.CISSettings, true
}

// HasCISSettingsField returns a boolean if a field has been set.
func (o *CISSettings) HasCISSettingsField() bool {
	if o != nil && !IsNil(o.CISSettings) {
		return true
	}

	return false
}

// SetCISSettings gets a reference to the given []CISSetting and assigns it to the CISSettings field.
func (o *CISSettings) SetCISSettings(v []CISSetting) {
	o.CISSettings = v
}

func (o CISSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CISSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CISSettings) {
		toSerialize["CISSettings"] = o.CISSettings
	}
	return toSerialize, nil
}

type NullableCISSettings struct {
	value *CISSettings
	isSet bool
}

func (v NullableCISSettings) Get() *CISSettings {
	return v.value
}

func (v *NullableCISSettings) Set(val *CISSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCISSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCISSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCISSettings(val *CISSettings) *NullableCISSettings {
	return &NullableCISSettings{value: val, isSet: true}
}

func (v NullableCISSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCISSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


