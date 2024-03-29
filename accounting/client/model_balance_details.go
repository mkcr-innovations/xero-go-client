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

// checks if the BalanceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceDetails{}

// BalanceDetails An array to specify multiple currency balances of an account
type BalanceDetails struct {
	// The opening balances of the account. Debits are positive, credits are negative values
	Balance *float64 `json:"Balance,omitempty"`
	// The currency of the balance (Not required for base currency)
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// (Optional) Exchange rate to base currency when money is spent or received. If not specified, XE rate for the day is applied
	CurrencyRate *float64 `json:"CurrencyRate,omitempty"`
}

// NewBalanceDetails instantiates a new BalanceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceDetails() *BalanceDetails {
	this := BalanceDetails{}
	return &this
}

// NewBalanceDetailsWithDefaults instantiates a new BalanceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceDetailsWithDefaults() *BalanceDetails {
	this := BalanceDetails{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *BalanceDetails) GetBalance() float64 {
	if o == nil || IsNil(o.Balance) {
		var ret float64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceDetails) GetBalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalanceField returns a boolean if a field has been set.
func (o *BalanceDetails) HasBalanceField() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float64 and assigns it to the Balance field.
func (o *BalanceDetails) SetBalance(v float64) {
	o.Balance = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BalanceDetails) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceDetails) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCodeField returns a boolean if a field has been set.
func (o *BalanceDetails) HasCurrencyCodeField() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BalanceDetails) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencyRate returns the CurrencyRate field value if set, zero value otherwise.
func (o *BalanceDetails) GetCurrencyRate() float64 {
	if o == nil || IsNil(o.CurrencyRate) {
		var ret float64
		return ret
	}
	return *o.CurrencyRate
}

// GetCurrencyRateOk returns a tuple with the CurrencyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceDetails) GetCurrencyRateOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrencyRate) {
		return nil, false
	}
	return o.CurrencyRate, true
}

// HasCurrencyRateField returns a boolean if a field has been set.
func (o *BalanceDetails) HasCurrencyRateField() bool {
	if o != nil && !IsNil(o.CurrencyRate) {
		return true
	}

	return false
}

// SetCurrencyRate gets a reference to the given float64 and assigns it to the CurrencyRate field.
func (o *BalanceDetails) SetCurrencyRate(v float64) {
	o.CurrencyRate = &v
}

func (o BalanceDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balance) {
		toSerialize["Balance"] = o.Balance
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencyRate) {
		toSerialize["CurrencyRate"] = o.CurrencyRate
	}
	return toSerialize, nil
}

type NullableBalanceDetails struct {
	value *BalanceDetails
	isSet bool
}

func (v NullableBalanceDetails) Get() *BalanceDetails {
	return v.value
}

func (v *NullableBalanceDetails) Set(val *BalanceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceDetails(val *BalanceDetails) *NullableBalanceDetails {
	return &NullableBalanceDetails{value: val, isSet: true}
}

func (v NullableBalanceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


