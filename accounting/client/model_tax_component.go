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

// checks if the TaxComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxComponent{}

// TaxComponent struct for TaxComponent
type TaxComponent struct {
	// Name of Tax Component
	Name *string `json:"Name,omitempty"`
	// Tax Rate (up to 4dp)
	Rate *float64 `json:"Rate,omitempty"`
	// Boolean to describe if Tax rate is compounded.
	IsCompound *bool `json:"IsCompound,omitempty"`
	// Boolean to describe if tax rate is non-recoverable. Non-recoverable rates are only applicable to Canadian organisations
	IsNonRecoverable *bool `json:"IsNonRecoverable,omitempty"`
}

// NewTaxComponent instantiates a new TaxComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxComponent() *TaxComponent {
	this := TaxComponent{}
	return &this
}

// NewTaxComponentWithDefaults instantiates a new TaxComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxComponentWithDefaults() *TaxComponent {
	this := TaxComponent{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaxComponent) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxComponent) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasNameField returns a boolean if a field has been set.
func (o *TaxComponent) HasNameField() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaxComponent) SetName(v string) {
	o.Name = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *TaxComponent) GetRate() float64 {
	if o == nil || IsNil(o.Rate) {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxComponent) GetRateOk() (*float64, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRateField returns a boolean if a field has been set.
func (o *TaxComponent) HasRateField() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *TaxComponent) SetRate(v float64) {
	o.Rate = &v
}

// GetIsCompound returns the IsCompound field value if set, zero value otherwise.
func (o *TaxComponent) GetIsCompound() bool {
	if o == nil || IsNil(o.IsCompound) {
		var ret bool
		return ret
	}
	return *o.IsCompound
}

// GetIsCompoundOk returns a tuple with the IsCompound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxComponent) GetIsCompoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompound) {
		return nil, false
	}
	return o.IsCompound, true
}

// HasIsCompoundField returns a boolean if a field has been set.
func (o *TaxComponent) HasIsCompoundField() bool {
	if o != nil && !IsNil(o.IsCompound) {
		return true
	}

	return false
}

// SetIsCompound gets a reference to the given bool and assigns it to the IsCompound field.
func (o *TaxComponent) SetIsCompound(v bool) {
	o.IsCompound = &v
}

// GetIsNonRecoverable returns the IsNonRecoverable field value if set, zero value otherwise.
func (o *TaxComponent) GetIsNonRecoverable() bool {
	if o == nil || IsNil(o.IsNonRecoverable) {
		var ret bool
		return ret
	}
	return *o.IsNonRecoverable
}

// GetIsNonRecoverableOk returns a tuple with the IsNonRecoverable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxComponent) GetIsNonRecoverableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNonRecoverable) {
		return nil, false
	}
	return o.IsNonRecoverable, true
}

// HasIsNonRecoverableField returns a boolean if a field has been set.
func (o *TaxComponent) HasIsNonRecoverableField() bool {
	if o != nil && !IsNil(o.IsNonRecoverable) {
		return true
	}

	return false
}

// SetIsNonRecoverable gets a reference to the given bool and assigns it to the IsNonRecoverable field.
func (o *TaxComponent) SetIsNonRecoverable(v bool) {
	o.IsNonRecoverable = &v
}

func (o TaxComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Rate) {
		toSerialize["Rate"] = o.Rate
	}
	if !IsNil(o.IsCompound) {
		toSerialize["IsCompound"] = o.IsCompound
	}
	if !IsNil(o.IsNonRecoverable) {
		toSerialize["IsNonRecoverable"] = o.IsNonRecoverable
	}
	return toSerialize, nil
}

type NullableTaxComponent struct {
	value *TaxComponent
	isSet bool
}

func (v NullableTaxComponent) Get() *TaxComponent {
	return v.value
}

func (v *NullableTaxComponent) Set(val *TaxComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxComponent(val *TaxComponent) *NullableTaxComponent {
	return &NullableTaxComponent{value: val, isSet: true}
}

func (v NullableTaxComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


