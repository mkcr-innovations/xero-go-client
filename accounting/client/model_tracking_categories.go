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

// checks if the TrackingCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingCategories{}

// TrackingCategories struct for TrackingCategories
type TrackingCategories struct {
	TrackingCategories []TrackingCategory `json:"TrackingCategories,omitempty"`
}

// NewTrackingCategories instantiates a new TrackingCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingCategories() *TrackingCategories {
	this := TrackingCategories{}
	return &this
}

// NewTrackingCategoriesWithDefaults instantiates a new TrackingCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingCategoriesWithDefaults() *TrackingCategories {
	this := TrackingCategories{}
	return &this
}

// GetTrackingCategories returns the TrackingCategories field value if set, zero value otherwise.
func (o *TrackingCategories) GetTrackingCategories() []TrackingCategory {
	if o == nil || IsNil(o.TrackingCategories) {
		var ret []TrackingCategory
		return ret
	}
	return o.TrackingCategories
}

// GetTrackingCategoriesOk returns a tuple with the TrackingCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingCategories) GetTrackingCategoriesOk() ([]TrackingCategory, bool) {
	if o == nil || IsNil(o.TrackingCategories) {
		return nil, false
	}
	return o.TrackingCategories, true
}

// HasTrackingCategoriesField returns a boolean if a field has been set.
func (o *TrackingCategories) HasTrackingCategoriesField() bool {
	if o != nil && !IsNil(o.TrackingCategories) {
		return true
	}

	return false
}

// SetTrackingCategories gets a reference to the given []TrackingCategory and assigns it to the TrackingCategories field.
func (o *TrackingCategories) SetTrackingCategories(v []TrackingCategory) {
	o.TrackingCategories = v
}

func (o TrackingCategories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackingCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingCategories) {
		toSerialize["TrackingCategories"] = o.TrackingCategories
	}
	return toSerialize, nil
}

type NullableTrackingCategories struct {
	value *TrackingCategories
	isSet bool
}

func (v NullableTrackingCategories) Get() *TrackingCategories {
	return v.value
}

func (v *NullableTrackingCategories) Set(val *TrackingCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingCategories(val *TrackingCategories) *NullableTrackingCategories {
	return &NullableTrackingCategories{value: val, isSet: true}
}

func (v NullableTrackingCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


