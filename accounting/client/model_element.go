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

// checks if the Element type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Element{}

// Element struct for Element
type Element struct {
	// Array of Validation Error message
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
	// Unique ID for batch payment object with validation error
	BatchPaymentID *string `json:"BatchPaymentID,omitempty"`
	BankTransactionID *string `json:"BankTransactionID,omitempty"`
	CreditNoteID *string `json:"CreditNoteID,omitempty"`
	ContactID *string `json:"ContactID,omitempty"`
	InvoiceID *string `json:"InvoiceID,omitempty"`
	ItemID *string `json:"ItemID,omitempty"`
	PurchaseOrderID *string `json:"PurchaseOrderID,omitempty"`
}

// NewElement instantiates a new Element object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElement() *Element {
	this := Element{}
	return &this
}

// NewElementWithDefaults instantiates a new Element object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElementWithDefaults() *Element {
	this := Element{}
	return &this
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *Element) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Element) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrorsField returns a boolean if a field has been set.
func (o *Element) HasValidationErrorsField() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *Element) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

// GetBatchPaymentID returns the BatchPaymentID field value if set, zero value otherwise.
func (o *Element) GetBatchPaymentID() string {
	if o == nil || IsNil(o.BatchPaymentID) {
		var ret string
		return ret
	}
	return *o.BatchPaymentID
}

// GetBatchPaymentIDOk returns a tuple with the BatchPaymentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Element) GetBatchPaymentIDOk() (*string, bool) {
	if o == nil || IsNil(o.BatchPaymentID) {
		return nil, false
	}
	return o.BatchPaymentID, true
}

// HasBatchPaymentIDField returns a boolean if a field has been set.
func (o *Element) HasBatchPaymentIDField() bool {
	if o != nil && !IsNil(o.BatchPaymentID) {
		return true
	}

	return false
}

// SetBatchPaymentID gets a reference to the given string and assigns it to the BatchPaymentID field.
func (o *Element) SetBatchPaymentID(v string) {
	o.BatchPaymentID = &v
}

// GetBankTransactionID returns the BankTransactionID field value if set, zero value otherwise.
func (o *Element) GetBankTransactionID() string {
	if o == nil || IsNil(o.BankTransactionID) {
		var ret string
		return ret
	}
	return *o.BankTransactionID
}

// GetBankTransactionIDOk returns a tuple with the BankTransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Element) GetBankTransactionIDOk() (*string, bool) {
	if o == nil || IsNil(o.BankTransactionID) {
		return nil, false
	}
	return o.BankTransactionID, true
}

// HasBankTransactionIDField returns a boolean if a field has been set.
func (o *Element) HasBankTransactionIDField() bool {
	if o != nil && !IsNil(o.BankTransactionID) {
		return true
	}

	return false
}

// SetBankTransactionID gets a reference to the given string and assigns it to the BankTransactionID field.
func (o *Element) SetBankTransactionID(v string) {
	o.BankTransactionID = &v
}

// GetCreditNoteID returns the CreditNoteID field value if set, zero value otherwise.
func (o *Element) GetCreditNoteID() string {
	if o == nil || IsNil(o.CreditNoteID) {
		var ret string
		return ret
	}
	return *o.CreditNoteID
}

// GetCreditNoteIDOk returns a tuple with the CreditNoteID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Element) GetCreditNoteIDOk() (*string, bool) {
	if o == nil || IsNil(o.CreditNoteID) {
		return nil, false
	}
	return o.CreditNoteID, true
}

// HasCreditNoteIDField returns a boolean if a field has been set.
func (o *Element) HasCreditNoteIDField() bool {
	if o != nil && !IsNil(o.CreditNoteID) {
		return true
	}

	return false
}

// SetCreditNoteID gets a reference to the given string and assigns it to the CreditNoteID field.
func (o *Element) SetCreditNoteID(v string) {
	o.CreditNoteID = &v
}

// GetContactID returns the ContactID field value if set, zero value otherwise.
func (o *Element) GetContactID() string {
	if o == nil || IsNil(o.ContactID) {
		var ret string
		return ret
	}
	return *o.ContactID
}

// GetContactIDOk returns a tuple with the ContactID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Element) GetContactIDOk() (*string, bool) {
	if o == nil || IsNil(o.ContactID) {
		return nil, false
	}
	return o.ContactID, true
}

// HasContactIDField returns a boolean if a field has been set.
func (o *Element) HasContactIDField() bool {
	if o != nil && !IsNil(o.ContactID) {
		return true
	}

	return false
}

// SetContactID gets a reference to the given string and assigns it to the ContactID field.
func (o *Element) SetContactID(v string) {
	o.ContactID = &v
}

// GetInvoiceID returns the InvoiceID field value if set, zero value otherwise.
func (o *Element) GetInvoiceID() string {
	if o == nil || IsNil(o.InvoiceID) {
		var ret string
		return ret
	}
	return *o.InvoiceID
}

// GetInvoiceIDOk returns a tuple with the InvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Element) GetInvoiceIDOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceID) {
		return nil, false
	}
	return o.InvoiceID, true
}

// HasInvoiceIDField returns a boolean if a field has been set.
func (o *Element) HasInvoiceIDField() bool {
	if o != nil && !IsNil(o.InvoiceID) {
		return true
	}

	return false
}

// SetInvoiceID gets a reference to the given string and assigns it to the InvoiceID field.
func (o *Element) SetInvoiceID(v string) {
	o.InvoiceID = &v
}

// GetItemID returns the ItemID field value if set, zero value otherwise.
func (o *Element) GetItemID() string {
	if o == nil || IsNil(o.ItemID) {
		var ret string
		return ret
	}
	return *o.ItemID
}

// GetItemIDOk returns a tuple with the ItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Element) GetItemIDOk() (*string, bool) {
	if o == nil || IsNil(o.ItemID) {
		return nil, false
	}
	return o.ItemID, true
}

// HasItemIDField returns a boolean if a field has been set.
func (o *Element) HasItemIDField() bool {
	if o != nil && !IsNil(o.ItemID) {
		return true
	}

	return false
}

// SetItemID gets a reference to the given string and assigns it to the ItemID field.
func (o *Element) SetItemID(v string) {
	o.ItemID = &v
}

// GetPurchaseOrderID returns the PurchaseOrderID field value if set, zero value otherwise.
func (o *Element) GetPurchaseOrderID() string {
	if o == nil || IsNil(o.PurchaseOrderID) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderID
}

// GetPurchaseOrderIDOk returns a tuple with the PurchaseOrderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Element) GetPurchaseOrderIDOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderID) {
		return nil, false
	}
	return o.PurchaseOrderID, true
}

// HasPurchaseOrderIDField returns a boolean if a field has been set.
func (o *Element) HasPurchaseOrderIDField() bool {
	if o != nil && !IsNil(o.PurchaseOrderID) {
		return true
	}

	return false
}

// SetPurchaseOrderID gets a reference to the given string and assigns it to the PurchaseOrderID field.
func (o *Element) SetPurchaseOrderID(v string) {
	o.PurchaseOrderID = &v
}

func (o Element) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Element) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	if !IsNil(o.BatchPaymentID) {
		toSerialize["BatchPaymentID"] = o.BatchPaymentID
	}
	if !IsNil(o.BankTransactionID) {
		toSerialize["BankTransactionID"] = o.BankTransactionID
	}
	if !IsNil(o.CreditNoteID) {
		toSerialize["CreditNoteID"] = o.CreditNoteID
	}
	if !IsNil(o.ContactID) {
		toSerialize["ContactID"] = o.ContactID
	}
	if !IsNil(o.InvoiceID) {
		toSerialize["InvoiceID"] = o.InvoiceID
	}
	if !IsNil(o.ItemID) {
		toSerialize["ItemID"] = o.ItemID
	}
	if !IsNil(o.PurchaseOrderID) {
		toSerialize["PurchaseOrderID"] = o.PurchaseOrderID
	}
	return toSerialize, nil
}

type NullableElement struct {
	value *Element
	isSet bool
}

func (v NullableElement) Get() *Element {
	return v.value
}

func (v *NullableElement) Set(val *Element) {
	v.value = val
	v.isSet = true
}

func (v NullableElement) IsSet() bool {
	return v.isSet
}

func (v *NullableElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElement(val *Element) *NullableElement {
	return &NullableElement{value: val, isSet: true}
}

func (v NullableElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

