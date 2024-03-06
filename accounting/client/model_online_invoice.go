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

// checks if the OnlineInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnlineInvoice{}

// OnlineInvoice struct for OnlineInvoice
type OnlineInvoice struct {
	// the URL to an online invoice
	OnlineInvoiceUrl *string `json:"OnlineInvoiceUrl,omitempty"`
}

// NewOnlineInvoice instantiates a new OnlineInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnlineInvoice() *OnlineInvoice {
	this := OnlineInvoice{}
	return &this
}

// NewOnlineInvoiceWithDefaults instantiates a new OnlineInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnlineInvoiceWithDefaults() *OnlineInvoice {
	this := OnlineInvoice{}
	return &this
}

// GetOnlineInvoiceUrl returns the OnlineInvoiceUrl field value if set, zero value otherwise.
func (o *OnlineInvoice) GetOnlineInvoiceUrl() string {
	if o == nil || IsNil(o.OnlineInvoiceUrl) {
		var ret string
		return ret
	}
	return *o.OnlineInvoiceUrl
}

// GetOnlineInvoiceUrlOk returns a tuple with the OnlineInvoiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnlineInvoice) GetOnlineInvoiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OnlineInvoiceUrl) {
		return nil, false
	}
	return o.OnlineInvoiceUrl, true
}

// HasOnlineInvoiceUrlField returns a boolean if a field has been set.
func (o *OnlineInvoice) HasOnlineInvoiceUrlField() bool {
	if o != nil && !IsNil(o.OnlineInvoiceUrl) {
		return true
	}

	return false
}

// SetOnlineInvoiceUrl gets a reference to the given string and assigns it to the OnlineInvoiceUrl field.
func (o *OnlineInvoice) SetOnlineInvoiceUrl(v string) {
	o.OnlineInvoiceUrl = &v
}

func (o OnlineInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnlineInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnlineInvoiceUrl) {
		toSerialize["OnlineInvoiceUrl"] = o.OnlineInvoiceUrl
	}
	return toSerialize, nil
}

type NullableOnlineInvoice struct {
	value *OnlineInvoice
	isSet bool
}

func (v NullableOnlineInvoice) Get() *OnlineInvoice {
	return v.value
}

func (v *NullableOnlineInvoice) Set(val *OnlineInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableOnlineInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableOnlineInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnlineInvoice(val *OnlineInvoice) *NullableOnlineInvoice {
	return &NullableOnlineInvoice{value: val, isSet: true}
}

func (v NullableOnlineInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnlineInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


