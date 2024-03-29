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

// checks if the ConversionBalances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConversionBalances{}

// ConversionBalances Balance supplied for each account that has a value as at the conversion date.
type ConversionBalances struct {
	// The account code for a account
	AccountCode *string `json:"AccountCode,omitempty"`
	// The opening balances of the account. Debits are positive, credits are negative values
	Balance *float64 `json:"Balance,omitempty"`
	BalanceDetails []BalanceDetails `json:"BalanceDetails,omitempty"`
}

// NewConversionBalances instantiates a new ConversionBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversionBalances() *ConversionBalances {
	this := ConversionBalances{}
	return &this
}

// NewConversionBalancesWithDefaults instantiates a new ConversionBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversionBalancesWithDefaults() *ConversionBalances {
	this := ConversionBalances{}
	return &this
}

// GetAccountCode returns the AccountCode field value if set, zero value otherwise.
func (o *ConversionBalances) GetAccountCode() string {
	if o == nil || IsNil(o.AccountCode) {
		var ret string
		return ret
	}
	return *o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionBalances) GetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCode) {
		return nil, false
	}
	return o.AccountCode, true
}

// HasAccountCodeField returns a boolean if a field has been set.
func (o *ConversionBalances) HasAccountCodeField() bool {
	if o != nil && !IsNil(o.AccountCode) {
		return true
	}

	return false
}

// SetAccountCode gets a reference to the given string and assigns it to the AccountCode field.
func (o *ConversionBalances) SetAccountCode(v string) {
	o.AccountCode = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *ConversionBalances) GetBalance() float64 {
	if o == nil || IsNil(o.Balance) {
		var ret float64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionBalances) GetBalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalanceField returns a boolean if a field has been set.
func (o *ConversionBalances) HasBalanceField() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float64 and assigns it to the Balance field.
func (o *ConversionBalances) SetBalance(v float64) {
	o.Balance = &v
}

// GetBalanceDetails returns the BalanceDetails field value if set, zero value otherwise.
func (o *ConversionBalances) GetBalanceDetails() []BalanceDetails {
	if o == nil || IsNil(o.BalanceDetails) {
		var ret []BalanceDetails
		return ret
	}
	return o.BalanceDetails
}

// GetBalanceDetailsOk returns a tuple with the BalanceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionBalances) GetBalanceDetailsOk() ([]BalanceDetails, bool) {
	if o == nil || IsNil(o.BalanceDetails) {
		return nil, false
	}
	return o.BalanceDetails, true
}

// HasBalanceDetailsField returns a boolean if a field has been set.
func (o *ConversionBalances) HasBalanceDetailsField() bool {
	if o != nil && !IsNil(o.BalanceDetails) {
		return true
	}

	return false
}

// SetBalanceDetails gets a reference to the given []BalanceDetails and assigns it to the BalanceDetails field.
func (o *ConversionBalances) SetBalanceDetails(v []BalanceDetails) {
	o.BalanceDetails = v
}

func (o ConversionBalances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConversionBalances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountCode) {
		toSerialize["AccountCode"] = o.AccountCode
	}
	if !IsNil(o.Balance) {
		toSerialize["Balance"] = o.Balance
	}
	if !IsNil(o.BalanceDetails) {
		toSerialize["BalanceDetails"] = o.BalanceDetails
	}
	return toSerialize, nil
}

type NullableConversionBalances struct {
	value *ConversionBalances
	isSet bool
}

func (v NullableConversionBalances) Get() *ConversionBalances {
	return v.value
}

func (v *NullableConversionBalances) Set(val *ConversionBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableConversionBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableConversionBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversionBalances(val *ConversionBalances) *NullableConversionBalances {
	return &NullableConversionBalances{value: val, isSet: true}
}

func (v NullableConversionBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversionBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


