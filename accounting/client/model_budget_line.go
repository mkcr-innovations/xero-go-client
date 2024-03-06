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

// checks if the BudgetLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetLine{}

// BudgetLine struct for BudgetLine
type BudgetLine struct {
	// See Accounts
	AccountID *string `json:"AccountID,omitempty"`
	// See Accounts
	AccountCode *string `json:"AccountCode,omitempty"`
	BudgetBalances []BudgetBalance `json:"BudgetBalances,omitempty"`
}

// NewBudgetLine instantiates a new BudgetLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetLine() *BudgetLine {
	this := BudgetLine{}
	return &this
}

// NewBudgetLineWithDefaults instantiates a new BudgetLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetLineWithDefaults() *BudgetLine {
	this := BudgetLine{}
	return &this
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *BudgetLine) GetAccountID() string {
	if o == nil || IsNil(o.AccountID) {
		var ret string
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLine) GetAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountIDField returns a boolean if a field has been set.
func (o *BudgetLine) HasAccountIDField() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given string and assigns it to the AccountID field.
func (o *BudgetLine) SetAccountID(v string) {
	o.AccountID = &v
}

// GetAccountCode returns the AccountCode field value if set, zero value otherwise.
func (o *BudgetLine) GetAccountCode() string {
	if o == nil || IsNil(o.AccountCode) {
		var ret string
		return ret
	}
	return *o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLine) GetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCode) {
		return nil, false
	}
	return o.AccountCode, true
}

// HasAccountCodeField returns a boolean if a field has been set.
func (o *BudgetLine) HasAccountCodeField() bool {
	if o != nil && !IsNil(o.AccountCode) {
		return true
	}

	return false
}

// SetAccountCode gets a reference to the given string and assigns it to the AccountCode field.
func (o *BudgetLine) SetAccountCode(v string) {
	o.AccountCode = &v
}

// GetBudgetBalances returns the BudgetBalances field value if set, zero value otherwise.
func (o *BudgetLine) GetBudgetBalances() []BudgetBalance {
	if o == nil || IsNil(o.BudgetBalances) {
		var ret []BudgetBalance
		return ret
	}
	return o.BudgetBalances
}

// GetBudgetBalancesOk returns a tuple with the BudgetBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLine) GetBudgetBalancesOk() ([]BudgetBalance, bool) {
	if o == nil || IsNil(o.BudgetBalances) {
		return nil, false
	}
	return o.BudgetBalances, true
}

// HasBudgetBalancesField returns a boolean if a field has been set.
func (o *BudgetLine) HasBudgetBalancesField() bool {
	if o != nil && !IsNil(o.BudgetBalances) {
		return true
	}

	return false
}

// SetBudgetBalances gets a reference to the given []BudgetBalance and assigns it to the BudgetBalances field.
func (o *BudgetLine) SetBudgetBalances(v []BudgetBalance) {
	o.BudgetBalances = v
}

func (o BudgetLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BudgetLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountID) {
		toSerialize["AccountID"] = o.AccountID
	}
	if !IsNil(o.AccountCode) {
		toSerialize["AccountCode"] = o.AccountCode
	}
	if !IsNil(o.BudgetBalances) {
		toSerialize["BudgetBalances"] = o.BudgetBalances
	}
	return toSerialize, nil
}

type NullableBudgetLine struct {
	value *BudgetLine
	isSet bool
}

func (v NullableBudgetLine) Get() *BudgetLine {
	return v.value
}

func (v *NullableBudgetLine) Set(val *BudgetLine) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetLine) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetLine(val *BudgetLine) *NullableBudgetLine {
	return &NullableBudgetLine{value: val, isSet: true}
}

func (v NullableBudgetLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


