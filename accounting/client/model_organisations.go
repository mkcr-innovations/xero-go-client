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

// checks if the Organisations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organisations{}

// Organisations struct for Organisations
type Organisations struct {
	Organisations []Organisation `json:"Organisations,omitempty"`
}

// NewOrganisations instantiates a new Organisations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisations() *Organisations {
	this := Organisations{}
	return &this
}

// NewOrganisationsWithDefaults instantiates a new Organisations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationsWithDefaults() *Organisations {
	this := Organisations{}
	return &this
}

// GetOrganisations returns the Organisations field value if set, zero value otherwise.
func (o *Organisations) GetOrganisations() []Organisation {
	if o == nil || IsNil(o.Organisations) {
		var ret []Organisation
		return ret
	}
	return o.Organisations
}

// GetOrganisationsOk returns a tuple with the Organisations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organisations) GetOrganisationsOk() ([]Organisation, bool) {
	if o == nil || IsNil(o.Organisations) {
		return nil, false
	}
	return o.Organisations, true
}

// HasOrganisationsField returns a boolean if a field has been set.
func (o *Organisations) HasOrganisationsField() bool {
	if o != nil && !IsNil(o.Organisations) {
		return true
	}

	return false
}

// SetOrganisations gets a reference to the given []Organisation and assigns it to the Organisations field.
func (o *Organisations) SetOrganisations(v []Organisation) {
	o.Organisations = v
}

func (o Organisations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organisations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organisations) {
		toSerialize["Organisations"] = o.Organisations
	}
	return toSerialize, nil
}

type NullableOrganisations struct {
	value *Organisations
	isSet bool
}

func (v NullableOrganisations) Get() *Organisations {
	return v.value
}

func (v *NullableOrganisations) Set(val *Organisations) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisations) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisations(val *Organisations) *NullableOrganisations {
	return &NullableOrganisations{value: val, isSet: true}
}

func (v NullableOrganisations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


