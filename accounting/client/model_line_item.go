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

// checks if the LineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItem{}

// LineItem struct for LineItem
type LineItem struct {
	// LineItem unique ID
	LineItemID *string `json:"LineItemID,omitempty"`
	// Description needs to be at least 1 char long. A line item with just a description (i.e no unit amount or quantity) can be created by specifying just a <Description> element that contains at least 1 character
	Description *string `json:"Description,omitempty"`
	// LineItem Quantity
	Quantity *float64 `json:"Quantity,omitempty"`
	// LineItem Unit Amount
	UnitAmount *float64 `json:"UnitAmount,omitempty"`
	// See Items
	ItemCode *string `json:"ItemCode,omitempty"`
	// See Accounts
	AccountCode *string `json:"AccountCode,omitempty"`
	// The associated account ID related to this line item
	AccountID *string `json:"AccountID,omitempty"`
	// The tax type from TaxRates
	TaxType *string `json:"TaxType,omitempty"`
	// The tax amount is auto calculated as a percentage of the line amount (see below) based on the tax rate. This value can be overriden if the calculated <TaxAmount> is not correct.
	TaxAmount *float64 `json:"TaxAmount,omitempty"`
	Item *LineItemItem `json:"Item,omitempty"`
	// If you wish to omit either the Quantity or UnitAmount you can provide a LineAmount and Xero will calculate the missing amount for you. The line amount reflects the discounted price if either a DiscountRate or DiscountAmount has been used i.e. LineAmount = Quantity * Unit Amount * ((100 - DiscountRate)/100) or LineAmount = (Quantity * UnitAmount) - DiscountAmount
	LineAmount *float64 `json:"LineAmount,omitempty"`
	// Optional Tracking Category – see Tracking.  Any LineItem can have a  maximum of 2 <TrackingCategory> elements.
	Tracking []LineItemTracking `json:"Tracking,omitempty"`
	// Percentage discount being applied to a line item (only supported on  ACCREC invoices – ACC PAY invoices and credit notes in Xero do not support discounts
	DiscountRate *float64 `json:"DiscountRate,omitempty"`
	// Discount amount being applied to a line item. Only supported on ACCREC invoices and quotes. ACCPAY invoices and credit notes in Xero do not support discounts.
	DiscountAmount *float64 `json:"DiscountAmount,omitempty"`
	// The Xero identifier for a Repeating Invoice
	RepeatingInvoiceID *string `json:"RepeatingInvoiceID,omitempty"`
}

// NewLineItem instantiates a new LineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItem() *LineItem {
	this := LineItem{}
	return &this
}

// NewLineItemWithDefaults instantiates a new LineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemWithDefaults() *LineItem {
	this := LineItem{}
	return &this
}

// GetLineItemID returns the LineItemID field value if set, zero value otherwise.
func (o *LineItem) GetLineItemID() string {
	if o == nil || IsNil(o.LineItemID) {
		var ret string
		return ret
	}
	return *o.LineItemID
}

// GetLineItemIDOk returns a tuple with the LineItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetLineItemIDOk() (*string, bool) {
	if o == nil || IsNil(o.LineItemID) {
		return nil, false
	}
	return o.LineItemID, true
}

// HasLineItemIDField returns a boolean if a field has been set.
func (o *LineItem) HasLineItemIDField() bool {
	if o != nil && !IsNil(o.LineItemID) {
		return true
	}

	return false
}

// SetLineItemID gets a reference to the given string and assigns it to the LineItemID field.
func (o *LineItem) SetLineItemID(v string) {
	o.LineItemID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LineItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescriptionField returns a boolean if a field has been set.
func (o *LineItem) HasDescriptionField() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LineItem) SetDescription(v string) {
	o.Description = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *LineItem) GetQuantity() float64 {
	if o == nil || IsNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantityField returns a boolean if a field has been set.
func (o *LineItem) HasQuantityField() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *LineItem) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetUnitAmount returns the UnitAmount field value if set, zero value otherwise.
func (o *LineItem) GetUnitAmount() float64 {
	if o == nil || IsNil(o.UnitAmount) {
		var ret float64
		return ret
	}
	return *o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetUnitAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.UnitAmount) {
		return nil, false
	}
	return o.UnitAmount, true
}

// HasUnitAmountField returns a boolean if a field has been set.
func (o *LineItem) HasUnitAmountField() bool {
	if o != nil && !IsNil(o.UnitAmount) {
		return true
	}

	return false
}

// SetUnitAmount gets a reference to the given float64 and assigns it to the UnitAmount field.
func (o *LineItem) SetUnitAmount(v float64) {
	o.UnitAmount = &v
}

// GetItemCode returns the ItemCode field value if set, zero value otherwise.
func (o *LineItem) GetItemCode() string {
	if o == nil || IsNil(o.ItemCode) {
		var ret string
		return ret
	}
	return *o.ItemCode
}

// GetItemCodeOk returns a tuple with the ItemCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetItemCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemCode) {
		return nil, false
	}
	return o.ItemCode, true
}

// HasItemCodeField returns a boolean if a field has been set.
func (o *LineItem) HasItemCodeField() bool {
	if o != nil && !IsNil(o.ItemCode) {
		return true
	}

	return false
}

// SetItemCode gets a reference to the given string and assigns it to the ItemCode field.
func (o *LineItem) SetItemCode(v string) {
	o.ItemCode = &v
}

// GetAccountCode returns the AccountCode field value if set, zero value otherwise.
func (o *LineItem) GetAccountCode() string {
	if o == nil || IsNil(o.AccountCode) {
		var ret string
		return ret
	}
	return *o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCode) {
		return nil, false
	}
	return o.AccountCode, true
}

// HasAccountCodeField returns a boolean if a field has been set.
func (o *LineItem) HasAccountCodeField() bool {
	if o != nil && !IsNil(o.AccountCode) {
		return true
	}

	return false
}

// SetAccountCode gets a reference to the given string and assigns it to the AccountCode field.
func (o *LineItem) SetAccountCode(v string) {
	o.AccountCode = &v
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *LineItem) GetAccountID() string {
	if o == nil || IsNil(o.AccountID) {
		var ret string
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountIDField returns a boolean if a field has been set.
func (o *LineItem) HasAccountIDField() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given string and assigns it to the AccountID field.
func (o *LineItem) SetAccountID(v string) {
	o.AccountID = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *LineItem) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxTypeField returns a boolean if a field has been set.
func (o *LineItem) HasTaxTypeField() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *LineItem) SetTaxType(v string) {
	o.TaxType = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *LineItem) GetTaxAmount() float64 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float64
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetTaxAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmountField returns a boolean if a field has been set.
func (o *LineItem) HasTaxAmountField() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float64 and assigns it to the TaxAmount field.
func (o *LineItem) SetTaxAmount(v float64) {
	o.TaxAmount = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *LineItem) GetItem() LineItemItem {
	if o == nil || IsNil(o.Item) {
		var ret LineItemItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetItemOk() (*LineItemItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItemField returns a boolean if a field has been set.
func (o *LineItem) HasItemField() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given LineItemItem and assigns it to the Item field.
func (o *LineItem) SetItem(v LineItemItem) {
	o.Item = &v
}

// GetLineAmount returns the LineAmount field value if set, zero value otherwise.
func (o *LineItem) GetLineAmount() float64 {
	if o == nil || IsNil(o.LineAmount) {
		var ret float64
		return ret
	}
	return *o.LineAmount
}

// GetLineAmountOk returns a tuple with the LineAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetLineAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.LineAmount) {
		return nil, false
	}
	return o.LineAmount, true
}

// HasLineAmountField returns a boolean if a field has been set.
func (o *LineItem) HasLineAmountField() bool {
	if o != nil && !IsNil(o.LineAmount) {
		return true
	}

	return false
}

// SetLineAmount gets a reference to the given float64 and assigns it to the LineAmount field.
func (o *LineItem) SetLineAmount(v float64) {
	o.LineAmount = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *LineItem) GetTracking() []LineItemTracking {
	if o == nil || IsNil(o.Tracking) {
		var ret []LineItemTracking
		return ret
	}
	return o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetTrackingOk() ([]LineItemTracking, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTrackingField returns a boolean if a field has been set.
func (o *LineItem) HasTrackingField() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given []LineItemTracking and assigns it to the Tracking field.
func (o *LineItem) SetTracking(v []LineItemTracking) {
	o.Tracking = v
}

// GetDiscountRate returns the DiscountRate field value if set, zero value otherwise.
func (o *LineItem) GetDiscountRate() float64 {
	if o == nil || IsNil(o.DiscountRate) {
		var ret float64
		return ret
	}
	return *o.DiscountRate
}

// GetDiscountRateOk returns a tuple with the DiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetDiscountRateOk() (*float64, bool) {
	if o == nil || IsNil(o.DiscountRate) {
		return nil, false
	}
	return o.DiscountRate, true
}

// HasDiscountRateField returns a boolean if a field has been set.
func (o *LineItem) HasDiscountRateField() bool {
	if o != nil && !IsNil(o.DiscountRate) {
		return true
	}

	return false
}

// SetDiscountRate gets a reference to the given float64 and assigns it to the DiscountRate field.
func (o *LineItem) SetDiscountRate(v float64) {
	o.DiscountRate = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *LineItem) GetDiscountAmount() float64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret float64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetDiscountAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmountField returns a boolean if a field has been set.
func (o *LineItem) HasDiscountAmountField() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given float64 and assigns it to the DiscountAmount field.
func (o *LineItem) SetDiscountAmount(v float64) {
	o.DiscountAmount = &v
}

// GetRepeatingInvoiceID returns the RepeatingInvoiceID field value if set, zero value otherwise.
func (o *LineItem) GetRepeatingInvoiceID() string {
	if o == nil || IsNil(o.RepeatingInvoiceID) {
		var ret string
		return ret
	}
	return *o.RepeatingInvoiceID
}

// GetRepeatingInvoiceIDOk returns a tuple with the RepeatingInvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetRepeatingInvoiceIDOk() (*string, bool) {
	if o == nil || IsNil(o.RepeatingInvoiceID) {
		return nil, false
	}
	return o.RepeatingInvoiceID, true
}

// HasRepeatingInvoiceIDField returns a boolean if a field has been set.
func (o *LineItem) HasRepeatingInvoiceIDField() bool {
	if o != nil && !IsNil(o.RepeatingInvoiceID) {
		return true
	}

	return false
}

// SetRepeatingInvoiceID gets a reference to the given string and assigns it to the RepeatingInvoiceID field.
func (o *LineItem) SetRepeatingInvoiceID(v string) {
	o.RepeatingInvoiceID = &v
}

func (o LineItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineItemID) {
		toSerialize["LineItemID"] = o.LineItemID
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if !IsNil(o.UnitAmount) {
		toSerialize["UnitAmount"] = o.UnitAmount
	}
	if !IsNil(o.ItemCode) {
		toSerialize["ItemCode"] = o.ItemCode
	}
	if !IsNil(o.AccountCode) {
		toSerialize["AccountCode"] = o.AccountCode
	}
	if !IsNil(o.AccountID) {
		toSerialize["AccountID"] = o.AccountID
	}
	if !IsNil(o.TaxType) {
		toSerialize["TaxType"] = o.TaxType
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["TaxAmount"] = o.TaxAmount
	}
	if !IsNil(o.Item) {
		toSerialize["Item"] = o.Item
	}
	if !IsNil(o.LineAmount) {
		toSerialize["LineAmount"] = o.LineAmount
	}
	if !IsNil(o.Tracking) {
		toSerialize["Tracking"] = o.Tracking
	}
	if !IsNil(o.DiscountRate) {
		toSerialize["DiscountRate"] = o.DiscountRate
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["DiscountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.RepeatingInvoiceID) {
		toSerialize["RepeatingInvoiceID"] = o.RepeatingInvoiceID
	}
	return toSerialize, nil
}

type NullableLineItem struct {
	value *LineItem
	isSet bool
}

func (v NullableLineItem) Get() *LineItem {
	return v.value
}

func (v *NullableLineItem) Set(val *LineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItem(val *LineItem) *NullableLineItem {
	return &NullableLineItem{value: val, isSet: true}
}

func (v NullableLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


