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

// checks if the LinkedTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkedTransaction{}

// LinkedTransaction struct for LinkedTransaction
type LinkedTransaction struct {
	// Filter by the SourceTransactionID. Get all the linked transactions created from a particular ACCPAY invoice
	SourceTransactionID *string `json:"SourceTransactionID,omitempty"`
	// The line item identifier from the source transaction.
	SourceLineItemID *string `json:"SourceLineItemID,omitempty"`
	// Filter by the combination of ContactID and Status. Get all the linked transactions that have been assigned to a particular customer and have a particular status e.g. GET /LinkedTransactions?ContactID=4bb34b03-3378-4bb2-a0ed-6345abf3224e&Status=APPROVED.
	ContactID *string `json:"ContactID,omitempty"`
	// Filter by the TargetTransactionID. Get all the linked transactions  allocated to a particular ACCREC invoice
	TargetTransactionID *string `json:"TargetTransactionID,omitempty"`
	// The line item identifier from the target transaction. It is possible  to link multiple billable expenses to the same TargetLineItemID.
	TargetLineItemID *string `json:"TargetLineItemID,omitempty"`
	// The Xero identifier for an Linked Transaction e.g./LinkedTransactions/297c2dc5-cc47-4afd-8ec8-74990b8761e9
	LinkedTransactionID *string `json:"LinkedTransactionID,omitempty"`
	// Filter by the combination of ContactID and Status. Get all the linked transactions that have been assigned to a particular customer and have a particular status e.g. GET /LinkedTransactions?ContactID=4bb34b03-3378-4bb2-a0ed-6345abf3224e&Status=APPROVED.
	Status *string `json:"Status,omitempty"`
	// This will always be BILLABLEEXPENSE. More types may be added in future.
	Type *string `json:"Type,omitempty"`
	// The last modified date in UTC format
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	// The Type of the source tranasction. This will be ACCPAY if the linked transaction was created from an invoice and SPEND if it was created from a bank transaction.
	SourceTransactionTypeCode *string `json:"SourceTransactionTypeCode,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
}

// NewLinkedTransaction instantiates a new LinkedTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedTransaction() *LinkedTransaction {
	this := LinkedTransaction{}
	return &this
}

// NewLinkedTransactionWithDefaults instantiates a new LinkedTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedTransactionWithDefaults() *LinkedTransaction {
	this := LinkedTransaction{}
	return &this
}

// GetSourceTransactionID returns the SourceTransactionID field value if set, zero value otherwise.
func (o *LinkedTransaction) GetSourceTransactionID() string {
	if o == nil || IsNil(o.SourceTransactionID) {
		var ret string
		return ret
	}
	return *o.SourceTransactionID
}

// GetSourceTransactionIDOk returns a tuple with the SourceTransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetSourceTransactionIDOk() (*string, bool) {
	if o == nil || IsNil(o.SourceTransactionID) {
		return nil, false
	}
	return o.SourceTransactionID, true
}

// HasSourceTransactionIDField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasSourceTransactionIDField() bool {
	if o != nil && !IsNil(o.SourceTransactionID) {
		return true
	}

	return false
}

// SetSourceTransactionID gets a reference to the given string and assigns it to the SourceTransactionID field.
func (o *LinkedTransaction) SetSourceTransactionID(v string) {
	o.SourceTransactionID = &v
}

// GetSourceLineItemID returns the SourceLineItemID field value if set, zero value otherwise.
func (o *LinkedTransaction) GetSourceLineItemID() string {
	if o == nil || IsNil(o.SourceLineItemID) {
		var ret string
		return ret
	}
	return *o.SourceLineItemID
}

// GetSourceLineItemIDOk returns a tuple with the SourceLineItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetSourceLineItemIDOk() (*string, bool) {
	if o == nil || IsNil(o.SourceLineItemID) {
		return nil, false
	}
	return o.SourceLineItemID, true
}

// HasSourceLineItemIDField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasSourceLineItemIDField() bool {
	if o != nil && !IsNil(o.SourceLineItemID) {
		return true
	}

	return false
}

// SetSourceLineItemID gets a reference to the given string and assigns it to the SourceLineItemID field.
func (o *LinkedTransaction) SetSourceLineItemID(v string) {
	o.SourceLineItemID = &v
}

// GetContactID returns the ContactID field value if set, zero value otherwise.
func (o *LinkedTransaction) GetContactID() string {
	if o == nil || IsNil(o.ContactID) {
		var ret string
		return ret
	}
	return *o.ContactID
}

// GetContactIDOk returns a tuple with the ContactID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetContactIDOk() (*string, bool) {
	if o == nil || IsNil(o.ContactID) {
		return nil, false
	}
	return o.ContactID, true
}

// HasContactIDField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasContactIDField() bool {
	if o != nil && !IsNil(o.ContactID) {
		return true
	}

	return false
}

// SetContactID gets a reference to the given string and assigns it to the ContactID field.
func (o *LinkedTransaction) SetContactID(v string) {
	o.ContactID = &v
}

// GetTargetTransactionID returns the TargetTransactionID field value if set, zero value otherwise.
func (o *LinkedTransaction) GetTargetTransactionID() string {
	if o == nil || IsNil(o.TargetTransactionID) {
		var ret string
		return ret
	}
	return *o.TargetTransactionID
}

// GetTargetTransactionIDOk returns a tuple with the TargetTransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetTargetTransactionIDOk() (*string, bool) {
	if o == nil || IsNil(o.TargetTransactionID) {
		return nil, false
	}
	return o.TargetTransactionID, true
}

// HasTargetTransactionIDField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasTargetTransactionIDField() bool {
	if o != nil && !IsNil(o.TargetTransactionID) {
		return true
	}

	return false
}

// SetTargetTransactionID gets a reference to the given string and assigns it to the TargetTransactionID field.
func (o *LinkedTransaction) SetTargetTransactionID(v string) {
	o.TargetTransactionID = &v
}

// GetTargetLineItemID returns the TargetLineItemID field value if set, zero value otherwise.
func (o *LinkedTransaction) GetTargetLineItemID() string {
	if o == nil || IsNil(o.TargetLineItemID) {
		var ret string
		return ret
	}
	return *o.TargetLineItemID
}

// GetTargetLineItemIDOk returns a tuple with the TargetLineItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetTargetLineItemIDOk() (*string, bool) {
	if o == nil || IsNil(o.TargetLineItemID) {
		return nil, false
	}
	return o.TargetLineItemID, true
}

// HasTargetLineItemIDField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasTargetLineItemIDField() bool {
	if o != nil && !IsNil(o.TargetLineItemID) {
		return true
	}

	return false
}

// SetTargetLineItemID gets a reference to the given string and assigns it to the TargetLineItemID field.
func (o *LinkedTransaction) SetTargetLineItemID(v string) {
	o.TargetLineItemID = &v
}

// GetLinkedTransactionID returns the LinkedTransactionID field value if set, zero value otherwise.
func (o *LinkedTransaction) GetLinkedTransactionID() string {
	if o == nil || IsNil(o.LinkedTransactionID) {
		var ret string
		return ret
	}
	return *o.LinkedTransactionID
}

// GetLinkedTransactionIDOk returns a tuple with the LinkedTransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetLinkedTransactionIDOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedTransactionID) {
		return nil, false
	}
	return o.LinkedTransactionID, true
}

// HasLinkedTransactionIDField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasLinkedTransactionIDField() bool {
	if o != nil && !IsNil(o.LinkedTransactionID) {
		return true
	}

	return false
}

// SetLinkedTransactionID gets a reference to the given string and assigns it to the LinkedTransactionID field.
func (o *LinkedTransaction) SetLinkedTransactionID(v string) {
	o.LinkedTransactionID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LinkedTransaction) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatusField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasStatusField() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LinkedTransaction) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LinkedTransaction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasTypeField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasTypeField() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LinkedTransaction) SetType(v string) {
	o.Type = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *LinkedTransaction) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTCField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasUpdatedDateUTCField() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *LinkedTransaction) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetSourceTransactionTypeCode returns the SourceTransactionTypeCode field value if set, zero value otherwise.
func (o *LinkedTransaction) GetSourceTransactionTypeCode() string {
	if o == nil || IsNil(o.SourceTransactionTypeCode) {
		var ret string
		return ret
	}
	return *o.SourceTransactionTypeCode
}

// GetSourceTransactionTypeCodeOk returns a tuple with the SourceTransactionTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetSourceTransactionTypeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceTransactionTypeCode) {
		return nil, false
	}
	return o.SourceTransactionTypeCode, true
}

// HasSourceTransactionTypeCodeField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasSourceTransactionTypeCodeField() bool {
	if o != nil && !IsNil(o.SourceTransactionTypeCode) {
		return true
	}

	return false
}

// SetSourceTransactionTypeCode gets a reference to the given string and assigns it to the SourceTransactionTypeCode field.
func (o *LinkedTransaction) SetSourceTransactionTypeCode(v string) {
	o.SourceTransactionTypeCode = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *LinkedTransaction) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTransaction) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrorsField returns a boolean if a field has been set.
func (o *LinkedTransaction) HasValidationErrorsField() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *LinkedTransaction) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

func (o LinkedTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkedTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceTransactionID) {
		toSerialize["SourceTransactionID"] = o.SourceTransactionID
	}
	if !IsNil(o.SourceLineItemID) {
		toSerialize["SourceLineItemID"] = o.SourceLineItemID
	}
	if !IsNil(o.ContactID) {
		toSerialize["ContactID"] = o.ContactID
	}
	if !IsNil(o.TargetTransactionID) {
		toSerialize["TargetTransactionID"] = o.TargetTransactionID
	}
	if !IsNil(o.TargetLineItemID) {
		toSerialize["TargetLineItemID"] = o.TargetLineItemID
	}
	if !IsNil(o.LinkedTransactionID) {
		toSerialize["LinkedTransactionID"] = o.LinkedTransactionID
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.SourceTransactionTypeCode) {
		toSerialize["SourceTransactionTypeCode"] = o.SourceTransactionTypeCode
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableLinkedTransaction struct {
	value *LinkedTransaction
	isSet bool
}

func (v NullableLinkedTransaction) Get() *LinkedTransaction {
	return v.value
}

func (v *NullableLinkedTransaction) Set(val *LinkedTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedTransaction(val *LinkedTransaction) *NullableLinkedTransaction {
	return &NullableLinkedTransaction{value: val, isSet: true}
}

func (v NullableLinkedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


