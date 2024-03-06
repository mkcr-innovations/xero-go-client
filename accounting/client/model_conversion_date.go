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

// checks if the ConversionDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConversionDate{}

// ConversionDate The date when the organisation starts using Xero
type ConversionDate struct {
	// The month the organisation starts using Xero. Value is an integer between 1 and 12
	Month *int32 `json:"Month,omitempty"`
	// The year the organisation starts using Xero. Value is an integer greater than 2006
	Year *int32 `json:"Year,omitempty"`
}

// NewConversionDate instantiates a new ConversionDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversionDate() *ConversionDate {
	this := ConversionDate{}
	return &this
}

// NewConversionDateWithDefaults instantiates a new ConversionDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversionDateWithDefaults() *ConversionDate {
	this := ConversionDate{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *ConversionDate) GetMonth() int32 {
	if o == nil || IsNil(o.Month) {
		var ret int32
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionDate) GetMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonthField returns a boolean if a field has been set.
func (o *ConversionDate) HasMonthField() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int32 and assigns it to the Month field.
func (o *ConversionDate) SetMonth(v int32) {
	o.Month = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *ConversionDate) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionDate) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYearField returns a boolean if a field has been set.
func (o *ConversionDate) HasYearField() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *ConversionDate) SetYear(v int32) {
	o.Year = &v
}

func (o ConversionDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConversionDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Month) {
		toSerialize["Month"] = o.Month
	}
	if !IsNil(o.Year) {
		toSerialize["Year"] = o.Year
	}
	return toSerialize, nil
}

type NullableConversionDate struct {
	value *ConversionDate
	isSet bool
}

func (v NullableConversionDate) Get() *ConversionDate {
	return v.value
}

func (v *NullableConversionDate) Set(val *ConversionDate) {
	v.value = val
	v.isSet = true
}

func (v NullableConversionDate) IsSet() bool {
	return v.isSet
}

func (v *NullableConversionDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversionDate(val *ConversionDate) *NullableConversionDate {
	return &NullableConversionDate{value: val, isSet: true}
}

func (v NullableConversionDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversionDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


