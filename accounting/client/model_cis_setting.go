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

// checks if the CISSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CISSetting{}

// CISSetting struct for CISSetting
type CISSetting struct {
	// Boolean that describes if the contact is a CIS Subcontractor
	CISEnabled *bool `json:"CISEnabled,omitempty"`
	// CIS Deduction rate for the contact if he is a subcontractor. If the contact is not CISEnabled, then the rate is not returned
	Rate *float64 `json:"Rate,omitempty"`
}

// NewCISSetting instantiates a new CISSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCISSetting() *CISSetting {
	this := CISSetting{}
	return &this
}

// NewCISSettingWithDefaults instantiates a new CISSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCISSettingWithDefaults() *CISSetting {
	this := CISSetting{}
	return &this
}

// GetCISEnabled returns the CISEnabled field value if set, zero value otherwise.
func (o *CISSetting) GetCISEnabled() bool {
	if o == nil || IsNil(o.CISEnabled) {
		var ret bool
		return ret
	}
	return *o.CISEnabled
}

// GetCISEnabledOk returns a tuple with the CISEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CISSetting) GetCISEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CISEnabled) {
		return nil, false
	}
	return o.CISEnabled, true
}

// HasCISEnabledField returns a boolean if a field has been set.
func (o *CISSetting) HasCISEnabledField() bool {
	if o != nil && !IsNil(o.CISEnabled) {
		return true
	}

	return false
}

// SetCISEnabled gets a reference to the given bool and assigns it to the CISEnabled field.
func (o *CISSetting) SetCISEnabled(v bool) {
	o.CISEnabled = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *CISSetting) GetRate() float64 {
	if o == nil || IsNil(o.Rate) {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CISSetting) GetRateOk() (*float64, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRateField returns a boolean if a field has been set.
func (o *CISSetting) HasRateField() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *CISSetting) SetRate(v float64) {
	o.Rate = &v
}

func (o CISSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CISSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CISEnabled) {
		toSerialize["CISEnabled"] = o.CISEnabled
	}
	if !IsNil(o.Rate) {
		toSerialize["Rate"] = o.Rate
	}
	return toSerialize, nil
}

type NullableCISSetting struct {
	value *CISSetting
	isSet bool
}

func (v NullableCISSetting) Get() *CISSetting {
	return v.value
}

func (v *NullableCISSetting) Set(val *CISSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableCISSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableCISSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCISSetting(val *CISSetting) *NullableCISSetting {
	return &NullableCISSetting{value: val, isSet: true}
}

func (v NullableCISSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCISSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


