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

// checks if the Receipts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Receipts{}

// Receipts struct for Receipts
type Receipts struct {
	Receipts []Receipt `json:"Receipts,omitempty"`
}

// NewReceipts instantiates a new Receipts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceipts() *Receipts {
	this := Receipts{}
	return &this
}

// NewReceiptsWithDefaults instantiates a new Receipts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptsWithDefaults() *Receipts {
	this := Receipts{}
	return &this
}

// GetReceipts returns the Receipts field value if set, zero value otherwise.
func (o *Receipts) GetReceipts() []Receipt {
	if o == nil || IsNil(o.Receipts) {
		var ret []Receipt
		return ret
	}
	return o.Receipts
}

// GetReceiptsOk returns a tuple with the Receipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipts) GetReceiptsOk() ([]Receipt, bool) {
	if o == nil || IsNil(o.Receipts) {
		return nil, false
	}
	return o.Receipts, true
}

// HasReceiptsField returns a boolean if a field has been set.
func (o *Receipts) HasReceiptsField() bool {
	if o != nil && !IsNil(o.Receipts) {
		return true
	}

	return false
}

// SetReceipts gets a reference to the given []Receipt and assigns it to the Receipts field.
func (o *Receipts) SetReceipts(v []Receipt) {
	o.Receipts = v
}

func (o Receipts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Receipts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Receipts) {
		toSerialize["Receipts"] = o.Receipts
	}
	return toSerialize, nil
}

type NullableReceipts struct {
	value *Receipts
	isSet bool
}

func (v NullableReceipts) Get() *Receipts {
	return v.value
}

func (v *NullableReceipts) Set(val *Receipts) {
	v.value = val
	v.isSet = true
}

func (v NullableReceipts) IsSet() bool {
	return v.isSet
}

func (v *NullableReceipts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceipts(val *Receipts) *NullableReceipts {
	return &NullableReceipts{value: val, isSet: true}
}

func (v NullableReceipts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceipts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


