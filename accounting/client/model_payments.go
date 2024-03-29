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

// checks if the Payments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payments{}

// Payments struct for Payments
type Payments struct {
	Payments []Payment `json:"Payments,omitempty"`
}

// NewPayments instantiates a new Payments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayments() *Payments {
	this := Payments{}
	return &this
}

// NewPaymentsWithDefaults instantiates a new Payments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentsWithDefaults() *Payments {
	this := Payments{}
	return &this
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *Payments) GetPayments() []Payment {
	if o == nil || IsNil(o.Payments) {
		var ret []Payment
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payments) GetPaymentsOk() ([]Payment, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPaymentsField returns a boolean if a field has been set.
func (o *Payments) HasPaymentsField() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []Payment and assigns it to the Payments field.
func (o *Payments) SetPayments(v []Payment) {
	o.Payments = v
}

func (o Payments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payments) {
		toSerialize["Payments"] = o.Payments
	}
	return toSerialize, nil
}

type NullablePayments struct {
	value *Payments
	isSet bool
}

func (v NullablePayments) Get() *Payments {
	return v.value
}

func (v *NullablePayments) Set(val *Payments) {
	v.value = val
	v.isSet = true
}

func (v NullablePayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayments(val *Payments) *NullablePayments {
	return &NullablePayments{value: val, isSet: true}
}

func (v NullablePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


