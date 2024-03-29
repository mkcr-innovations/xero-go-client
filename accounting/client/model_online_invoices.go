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

// checks if the OnlineInvoices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnlineInvoices{}

// OnlineInvoices struct for OnlineInvoices
type OnlineInvoices struct {
	OnlineInvoices []OnlineInvoice `json:"OnlineInvoices,omitempty"`
}

// NewOnlineInvoices instantiates a new OnlineInvoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnlineInvoices() *OnlineInvoices {
	this := OnlineInvoices{}
	return &this
}

// NewOnlineInvoicesWithDefaults instantiates a new OnlineInvoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnlineInvoicesWithDefaults() *OnlineInvoices {
	this := OnlineInvoices{}
	return &this
}

// GetOnlineInvoices returns the OnlineInvoices field value if set, zero value otherwise.
func (o *OnlineInvoices) GetOnlineInvoices() []OnlineInvoice {
	if o == nil || IsNil(o.OnlineInvoices) {
		var ret []OnlineInvoice
		return ret
	}
	return o.OnlineInvoices
}

// GetOnlineInvoicesOk returns a tuple with the OnlineInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnlineInvoices) GetOnlineInvoicesOk() ([]OnlineInvoice, bool) {
	if o == nil || IsNil(o.OnlineInvoices) {
		return nil, false
	}
	return o.OnlineInvoices, true
}

// HasOnlineInvoicesField returns a boolean if a field has been set.
func (o *OnlineInvoices) HasOnlineInvoicesField() bool {
	if o != nil && !IsNil(o.OnlineInvoices) {
		return true
	}

	return false
}

// SetOnlineInvoices gets a reference to the given []OnlineInvoice and assigns it to the OnlineInvoices field.
func (o *OnlineInvoices) SetOnlineInvoices(v []OnlineInvoice) {
	o.OnlineInvoices = v
}

func (o OnlineInvoices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnlineInvoices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnlineInvoices) {
		toSerialize["OnlineInvoices"] = o.OnlineInvoices
	}
	return toSerialize, nil
}

type NullableOnlineInvoices struct {
	value *OnlineInvoices
	isSet bool
}

func (v NullableOnlineInvoices) Get() *OnlineInvoices {
	return v.value
}

func (v *NullableOnlineInvoices) Set(val *OnlineInvoices) {
	v.value = val
	v.isSet = true
}

func (v NullableOnlineInvoices) IsSet() bool {
	return v.isSet
}

func (v *NullableOnlineInvoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnlineInvoices(val *OnlineInvoices) *NullableOnlineInvoices {
	return &NullableOnlineInvoices{value: val, isSet: true}
}

func (v NullableOnlineInvoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnlineInvoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


