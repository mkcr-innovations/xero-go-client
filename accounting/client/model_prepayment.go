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

// checks if the Prepayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Prepayment{}

// Prepayment struct for Prepayment
type Prepayment struct {
	// See Prepayment Types
	Type *string `json:"Type,omitempty"`
	Contact *Contact `json:"Contact,omitempty"`
	// The date the prepayment is created YYYY-MM-DD
	Date *string `json:"Date,omitempty"`
	// See Prepayment Status Codes
	Status *string `json:"Status,omitempty"`
	LineAmountTypes *LineAmountTypes `json:"LineAmountTypes,omitempty"`
	// See Prepayment Line Items
	LineItems []LineItem `json:"LineItems,omitempty"`
	// The subtotal of the prepayment excluding taxes
	SubTotal *float64 `json:"SubTotal,omitempty"`
	// The total tax on the prepayment
	TotalTax *float64 `json:"TotalTax,omitempty"`
	// The total of the prepayment(subtotal + total tax)
	Total *float64 `json:"Total,omitempty"`
	// Returns Invoice number field. Reference field isn't available.
	Reference *string `json:"Reference,omitempty"`
	// UTC timestamp of last update to the prepayment
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"`
	// Xero generated unique identifier
	PrepaymentID *string `json:"PrepaymentID,omitempty"`
	// The currency rate for a multicurrency prepayment. If no rate is specified, the XE.com day rate is used
	CurrencyRate *float64 `json:"CurrencyRate,omitempty"`
	// The remaining credit balance on the prepayment
	RemainingCredit *float64 `json:"RemainingCredit,omitempty"`
	// See Allocations
	Allocations []Allocation `json:"Allocations,omitempty"`
	// See Payments
	Payments []Payment `json:"Payments,omitempty"`
	// The amount of applied to an invoice
	AppliedAmount *float64 `json:"AppliedAmount,omitempty"`
	// boolean to indicate if a prepayment has an attachment
	HasAttachments *bool `json:"HasAttachments,omitempty"`
	// See Attachments
	Attachments []Attachment `json:"Attachments,omitempty"`
}

// NewPrepayment instantiates a new Prepayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrepayment() *Prepayment {
	this := Prepayment{}
	return &this
}

// NewPrepaymentWithDefaults instantiates a new Prepayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrepaymentWithDefaults() *Prepayment {
	this := Prepayment{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Prepayment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasTypeField returns a boolean if a field has been set.
func (o *Prepayment) HasTypeField() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Prepayment) SetType(v string) {
	o.Type = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *Prepayment) GetContact() Contact {
	if o == nil || IsNil(o.Contact) {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetContactOk() (*Contact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContactField returns a boolean if a field has been set.
func (o *Prepayment) HasContactField() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *Prepayment) SetContact(v Contact) {
	o.Contact = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Prepayment) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDateField returns a boolean if a field has been set.
func (o *Prepayment) HasDateField() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Prepayment) SetDate(v string) {
	o.Date = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Prepayment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatusField returns a boolean if a field has been set.
func (o *Prepayment) HasStatusField() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Prepayment) SetStatus(v string) {
	o.Status = &v
}

// GetLineAmountTypes returns the LineAmountTypes field value if set, zero value otherwise.
func (o *Prepayment) GetLineAmountTypes() LineAmountTypes {
	if o == nil || IsNil(o.LineAmountTypes) {
		var ret LineAmountTypes
		return ret
	}
	return *o.LineAmountTypes
}

// GetLineAmountTypesOk returns a tuple with the LineAmountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetLineAmountTypesOk() (*LineAmountTypes, bool) {
	if o == nil || IsNil(o.LineAmountTypes) {
		return nil, false
	}
	return o.LineAmountTypes, true
}

// HasLineAmountTypesField returns a boolean if a field has been set.
func (o *Prepayment) HasLineAmountTypesField() bool {
	if o != nil && !IsNil(o.LineAmountTypes) {
		return true
	}

	return false
}

// SetLineAmountTypes gets a reference to the given LineAmountTypes and assigns it to the LineAmountTypes field.
func (o *Prepayment) SetLineAmountTypes(v LineAmountTypes) {
	o.LineAmountTypes = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *Prepayment) GetLineItems() []LineItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItemsField returns a boolean if a field has been set.
func (o *Prepayment) HasLineItemsField() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *Prepayment) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetSubTotal returns the SubTotal field value if set, zero value otherwise.
func (o *Prepayment) GetSubTotal() float64 {
	if o == nil || IsNil(o.SubTotal) {
		var ret float64
		return ret
	}
	return *o.SubTotal
}

// GetSubTotalOk returns a tuple with the SubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetSubTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.SubTotal) {
		return nil, false
	}
	return o.SubTotal, true
}

// HasSubTotalField returns a boolean if a field has been set.
func (o *Prepayment) HasSubTotalField() bool {
	if o != nil && !IsNil(o.SubTotal) {
		return true
	}

	return false
}

// SetSubTotal gets a reference to the given float64 and assigns it to the SubTotal field.
func (o *Prepayment) SetSubTotal(v float64) {
	o.SubTotal = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *Prepayment) GetTotalTax() float64 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float64
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetTotalTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTaxField returns a boolean if a field has been set.
func (o *Prepayment) HasTotalTaxField() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float64 and assigns it to the TotalTax field.
func (o *Prepayment) SetTotalTax(v float64) {
	o.TotalTax = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Prepayment) GetTotal() float64 {
	if o == nil || IsNil(o.Total) {
		var ret float64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotalField returns a boolean if a field has been set.
func (o *Prepayment) HasTotalField() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float64 and assigns it to the Total field.
func (o *Prepayment) SetTotal(v float64) {
	o.Total = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Prepayment) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReferenceField returns a boolean if a field has been set.
func (o *Prepayment) HasReferenceField() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Prepayment) SetReference(v string) {
	o.Reference = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *Prepayment) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTCField returns a boolean if a field has been set.
func (o *Prepayment) HasUpdatedDateUTCField() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *Prepayment) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Prepayment) GetCurrencyCode() CurrencyCode {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret CurrencyCode
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCodeField returns a boolean if a field has been set.
func (o *Prepayment) HasCurrencyCodeField() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given CurrencyCode and assigns it to the CurrencyCode field.
func (o *Prepayment) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = &v
}

// GetPrepaymentID returns the PrepaymentID field value if set, zero value otherwise.
func (o *Prepayment) GetPrepaymentID() string {
	if o == nil || IsNil(o.PrepaymentID) {
		var ret string
		return ret
	}
	return *o.PrepaymentID
}

// GetPrepaymentIDOk returns a tuple with the PrepaymentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetPrepaymentIDOk() (*string, bool) {
	if o == nil || IsNil(o.PrepaymentID) {
		return nil, false
	}
	return o.PrepaymentID, true
}

// HasPrepaymentIDField returns a boolean if a field has been set.
func (o *Prepayment) HasPrepaymentIDField() bool {
	if o != nil && !IsNil(o.PrepaymentID) {
		return true
	}

	return false
}

// SetPrepaymentID gets a reference to the given string and assigns it to the PrepaymentID field.
func (o *Prepayment) SetPrepaymentID(v string) {
	o.PrepaymentID = &v
}

// GetCurrencyRate returns the CurrencyRate field value if set, zero value otherwise.
func (o *Prepayment) GetCurrencyRate() float64 {
	if o == nil || IsNil(o.CurrencyRate) {
		var ret float64
		return ret
	}
	return *o.CurrencyRate
}

// GetCurrencyRateOk returns a tuple with the CurrencyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetCurrencyRateOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrencyRate) {
		return nil, false
	}
	return o.CurrencyRate, true
}

// HasCurrencyRateField returns a boolean if a field has been set.
func (o *Prepayment) HasCurrencyRateField() bool {
	if o != nil && !IsNil(o.CurrencyRate) {
		return true
	}

	return false
}

// SetCurrencyRate gets a reference to the given float64 and assigns it to the CurrencyRate field.
func (o *Prepayment) SetCurrencyRate(v float64) {
	o.CurrencyRate = &v
}

// GetRemainingCredit returns the RemainingCredit field value if set, zero value otherwise.
func (o *Prepayment) GetRemainingCredit() float64 {
	if o == nil || IsNil(o.RemainingCredit) {
		var ret float64
		return ret
	}
	return *o.RemainingCredit
}

// GetRemainingCreditOk returns a tuple with the RemainingCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetRemainingCreditOk() (*float64, bool) {
	if o == nil || IsNil(o.RemainingCredit) {
		return nil, false
	}
	return o.RemainingCredit, true
}

// HasRemainingCreditField returns a boolean if a field has been set.
func (o *Prepayment) HasRemainingCreditField() bool {
	if o != nil && !IsNil(o.RemainingCredit) {
		return true
	}

	return false
}

// SetRemainingCredit gets a reference to the given float64 and assigns it to the RemainingCredit field.
func (o *Prepayment) SetRemainingCredit(v float64) {
	o.RemainingCredit = &v
}

// GetAllocations returns the Allocations field value if set, zero value otherwise.
func (o *Prepayment) GetAllocations() []Allocation {
	if o == nil || IsNil(o.Allocations) {
		var ret []Allocation
		return ret
	}
	return o.Allocations
}

// GetAllocationsOk returns a tuple with the Allocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetAllocationsOk() ([]Allocation, bool) {
	if o == nil || IsNil(o.Allocations) {
		return nil, false
	}
	return o.Allocations, true
}

// HasAllocationsField returns a boolean if a field has been set.
func (o *Prepayment) HasAllocationsField() bool {
	if o != nil && !IsNil(o.Allocations) {
		return true
	}

	return false
}

// SetAllocations gets a reference to the given []Allocation and assigns it to the Allocations field.
func (o *Prepayment) SetAllocations(v []Allocation) {
	o.Allocations = v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *Prepayment) GetPayments() []Payment {
	if o == nil || IsNil(o.Payments) {
		var ret []Payment
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetPaymentsOk() ([]Payment, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPaymentsField returns a boolean if a field has been set.
func (o *Prepayment) HasPaymentsField() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []Payment and assigns it to the Payments field.
func (o *Prepayment) SetPayments(v []Payment) {
	o.Payments = v
}

// GetAppliedAmount returns the AppliedAmount field value if set, zero value otherwise.
func (o *Prepayment) GetAppliedAmount() float64 {
	if o == nil || IsNil(o.AppliedAmount) {
		var ret float64
		return ret
	}
	return *o.AppliedAmount
}

// GetAppliedAmountOk returns a tuple with the AppliedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetAppliedAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.AppliedAmount) {
		return nil, false
	}
	return o.AppliedAmount, true
}

// HasAppliedAmountField returns a boolean if a field has been set.
func (o *Prepayment) HasAppliedAmountField() bool {
	if o != nil && !IsNil(o.AppliedAmount) {
		return true
	}

	return false
}

// SetAppliedAmount gets a reference to the given float64 and assigns it to the AppliedAmount field.
func (o *Prepayment) SetAppliedAmount(v float64) {
	o.AppliedAmount = &v
}

// GetHasAttachments returns the HasAttachments field value if set, zero value otherwise.
func (o *Prepayment) GetHasAttachments() bool {
	if o == nil || IsNil(o.HasAttachments) {
		var ret bool
		return ret
	}
	return *o.HasAttachments
}

// GetHasAttachmentsOk returns a tuple with the HasAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetHasAttachmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAttachments) {
		return nil, false
	}
	return o.HasAttachments, true
}

// HasHasAttachmentsField returns a boolean if a field has been set.
func (o *Prepayment) HasHasAttachmentsField() bool {
	if o != nil && !IsNil(o.HasAttachments) {
		return true
	}

	return false
}

// SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.
func (o *Prepayment) SetHasAttachments(v bool) {
	o.HasAttachments = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *Prepayment) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prepayment) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachmentsField returns a boolean if a field has been set.
func (o *Prepayment) HasAttachmentsField() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *Prepayment) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o Prepayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Prepayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Contact) {
		toSerialize["Contact"] = o.Contact
	}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.LineAmountTypes) {
		toSerialize["LineAmountTypes"] = o.LineAmountTypes
	}
	if !IsNil(o.LineItems) {
		toSerialize["LineItems"] = o.LineItems
	}
	if !IsNil(o.SubTotal) {
		toSerialize["SubTotal"] = o.SubTotal
	}
	if !IsNil(o.TotalTax) {
		toSerialize["TotalTax"] = o.TotalTax
	}
	if !IsNil(o.Total) {
		toSerialize["Total"] = o.Total
	}
	if !IsNil(o.Reference) {
		toSerialize["Reference"] = o.Reference
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.PrepaymentID) {
		toSerialize["PrepaymentID"] = o.PrepaymentID
	}
	if !IsNil(o.CurrencyRate) {
		toSerialize["CurrencyRate"] = o.CurrencyRate
	}
	if !IsNil(o.RemainingCredit) {
		toSerialize["RemainingCredit"] = o.RemainingCredit
	}
	if !IsNil(o.Allocations) {
		toSerialize["Allocations"] = o.Allocations
	}
	if !IsNil(o.Payments) {
		toSerialize["Payments"] = o.Payments
	}
	if !IsNil(o.AppliedAmount) {
		toSerialize["AppliedAmount"] = o.AppliedAmount
	}
	if !IsNil(o.HasAttachments) {
		toSerialize["HasAttachments"] = o.HasAttachments
	}
	if !IsNil(o.Attachments) {
		toSerialize["Attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullablePrepayment struct {
	value *Prepayment
	isSet bool
}

func (v NullablePrepayment) Get() *Prepayment {
	return v.value
}

func (v *NullablePrepayment) Set(val *Prepayment) {
	v.value = val
	v.isSet = true
}

func (v NullablePrepayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePrepayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrepayment(val *Prepayment) *NullablePrepayment {
	return &NullablePrepayment{value: val, isSet: true}
}

func (v NullablePrepayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrepayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


