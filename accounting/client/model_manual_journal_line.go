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

// checks if the ManualJournalLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualJournalLine{}

// ManualJournalLine struct for ManualJournalLine
type ManualJournalLine struct {
	// total for line. Debits are positive, credits are negative value
	LineAmount *float64 `json:"LineAmount,omitempty"`
	// See Accounts
	AccountCode *string `json:"AccountCode,omitempty"`
	// See Accounts
	AccountID *string `json:"AccountID,omitempty"`
	// Description for journal line
	Description *string `json:"Description,omitempty"`
	// The tax type from TaxRates
	TaxType *string `json:"TaxType,omitempty"`
	// Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 <TrackingCategory> elements.
	Tracking []TrackingCategory `json:"Tracking,omitempty"`
	// The calculated tax amount based on the TaxType and LineAmount
	TaxAmount *float64 `json:"TaxAmount,omitempty"`
	// is the line blank
	IsBlank *bool `json:"IsBlank,omitempty"`
}

// NewManualJournalLine instantiates a new ManualJournalLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualJournalLine() *ManualJournalLine {
	this := ManualJournalLine{}
	return &this
}

// NewManualJournalLineWithDefaults instantiates a new ManualJournalLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualJournalLineWithDefaults() *ManualJournalLine {
	this := ManualJournalLine{}
	return &this
}

// GetLineAmount returns the LineAmount field value if set, zero value otherwise.
func (o *ManualJournalLine) GetLineAmount() float64 {
	if o == nil || IsNil(o.LineAmount) {
		var ret float64
		return ret
	}
	return *o.LineAmount
}

// GetLineAmountOk returns a tuple with the LineAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournalLine) GetLineAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.LineAmount) {
		return nil, false
	}
	return o.LineAmount, true
}

// HasLineAmountField returns a boolean if a field has been set.
func (o *ManualJournalLine) HasLineAmountField() bool {
	if o != nil && !IsNil(o.LineAmount) {
		return true
	}

	return false
}

// SetLineAmount gets a reference to the given float64 and assigns it to the LineAmount field.
func (o *ManualJournalLine) SetLineAmount(v float64) {
	o.LineAmount = &v
}

// GetAccountCode returns the AccountCode field value if set, zero value otherwise.
func (o *ManualJournalLine) GetAccountCode() string {
	if o == nil || IsNil(o.AccountCode) {
		var ret string
		return ret
	}
	return *o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournalLine) GetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCode) {
		return nil, false
	}
	return o.AccountCode, true
}

// HasAccountCodeField returns a boolean if a field has been set.
func (o *ManualJournalLine) HasAccountCodeField() bool {
	if o != nil && !IsNil(o.AccountCode) {
		return true
	}

	return false
}

// SetAccountCode gets a reference to the given string and assigns it to the AccountCode field.
func (o *ManualJournalLine) SetAccountCode(v string) {
	o.AccountCode = &v
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *ManualJournalLine) GetAccountID() string {
	if o == nil || IsNil(o.AccountID) {
		var ret string
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournalLine) GetAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountIDField returns a boolean if a field has been set.
func (o *ManualJournalLine) HasAccountIDField() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given string and assigns it to the AccountID field.
func (o *ManualJournalLine) SetAccountID(v string) {
	o.AccountID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManualJournalLine) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournalLine) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescriptionField returns a boolean if a field has been set.
func (o *ManualJournalLine) HasDescriptionField() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManualJournalLine) SetDescription(v string) {
	o.Description = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *ManualJournalLine) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournalLine) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxTypeField returns a boolean if a field has been set.
func (o *ManualJournalLine) HasTaxTypeField() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *ManualJournalLine) SetTaxType(v string) {
	o.TaxType = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *ManualJournalLine) GetTracking() []TrackingCategory {
	if o == nil || IsNil(o.Tracking) {
		var ret []TrackingCategory
		return ret
	}
	return o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournalLine) GetTrackingOk() ([]TrackingCategory, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTrackingField returns a boolean if a field has been set.
func (o *ManualJournalLine) HasTrackingField() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given []TrackingCategory and assigns it to the Tracking field.
func (o *ManualJournalLine) SetTracking(v []TrackingCategory) {
	o.Tracking = v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *ManualJournalLine) GetTaxAmount() float64 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float64
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournalLine) GetTaxAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmountField returns a boolean if a field has been set.
func (o *ManualJournalLine) HasTaxAmountField() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float64 and assigns it to the TaxAmount field.
func (o *ManualJournalLine) SetTaxAmount(v float64) {
	o.TaxAmount = &v
}

// GetIsBlank returns the IsBlank field value if set, zero value otherwise.
func (o *ManualJournalLine) GetIsBlank() bool {
	if o == nil || IsNil(o.IsBlank) {
		var ret bool
		return ret
	}
	return *o.IsBlank
}

// GetIsBlankOk returns a tuple with the IsBlank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournalLine) GetIsBlankOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBlank) {
		return nil, false
	}
	return o.IsBlank, true
}

// HasIsBlankField returns a boolean if a field has been set.
func (o *ManualJournalLine) HasIsBlankField() bool {
	if o != nil && !IsNil(o.IsBlank) {
		return true
	}

	return false
}

// SetIsBlank gets a reference to the given bool and assigns it to the IsBlank field.
func (o *ManualJournalLine) SetIsBlank(v bool) {
	o.IsBlank = &v
}

func (o ManualJournalLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualJournalLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineAmount) {
		toSerialize["LineAmount"] = o.LineAmount
	}
	if !IsNil(o.AccountCode) {
		toSerialize["AccountCode"] = o.AccountCode
	}
	if !IsNil(o.AccountID) {
		toSerialize["AccountID"] = o.AccountID
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.TaxType) {
		toSerialize["TaxType"] = o.TaxType
	}
	if !IsNil(o.Tracking) {
		toSerialize["Tracking"] = o.Tracking
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["TaxAmount"] = o.TaxAmount
	}
	if !IsNil(o.IsBlank) {
		toSerialize["IsBlank"] = o.IsBlank
	}
	return toSerialize, nil
}

type NullableManualJournalLine struct {
	value *ManualJournalLine
	isSet bool
}

func (v NullableManualJournalLine) Get() *ManualJournalLine {
	return v.value
}

func (v *NullableManualJournalLine) Set(val *ManualJournalLine) {
	v.value = val
	v.isSet = true
}

func (v NullableManualJournalLine) IsSet() bool {
	return v.isSet
}

func (v *NullableManualJournalLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualJournalLine(val *ManualJournalLine) *NullableManualJournalLine {
	return &NullableManualJournalLine{value: val, isSet: true}
}

func (v NullableManualJournalLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualJournalLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


