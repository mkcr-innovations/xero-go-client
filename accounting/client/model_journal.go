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

// checks if the Journal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Journal{}

// Journal struct for Journal
type Journal struct {
	// Xero identifier
	JournalID *string `json:"JournalID,omitempty"`
	// Date the journal was posted
	JournalDate *string `json:"JournalDate,omitempty"`
	// Xero generated journal number
	JournalNumber *int32 `json:"JournalNumber,omitempty"`
	// Created date UTC format
	CreatedDateUTC *string `json:"CreatedDateUTC,omitempty"`
	// reference field for additional indetifying information
	Reference *string `json:"Reference,omitempty"`
	// The identifier for the source transaction (e.g. InvoiceID)
	SourceID *string `json:"SourceID,omitempty"`
	// The journal source type. The type of transaction that created the journal
	SourceType *string `json:"SourceType,omitempty"`
	// See JournalLines
	JournalLines []JournalLine `json:"JournalLines,omitempty"`
}

// NewJournal instantiates a new Journal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournal() *Journal {
	this := Journal{}
	return &this
}

// NewJournalWithDefaults instantiates a new Journal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalWithDefaults() *Journal {
	this := Journal{}
	return &this
}

// GetJournalID returns the JournalID field value if set, zero value otherwise.
func (o *Journal) GetJournalID() string {
	if o == nil || IsNil(o.JournalID) {
		var ret string
		return ret
	}
	return *o.JournalID
}

// GetJournalIDOk returns a tuple with the JournalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetJournalIDOk() (*string, bool) {
	if o == nil || IsNil(o.JournalID) {
		return nil, false
	}
	return o.JournalID, true
}

// HasJournalIDField returns a boolean if a field has been set.
func (o *Journal) HasJournalIDField() bool {
	if o != nil && !IsNil(o.JournalID) {
		return true
	}

	return false
}

// SetJournalID gets a reference to the given string and assigns it to the JournalID field.
func (o *Journal) SetJournalID(v string) {
	o.JournalID = &v
}

// GetJournalDate returns the JournalDate field value if set, zero value otherwise.
func (o *Journal) GetJournalDate() string {
	if o == nil || IsNil(o.JournalDate) {
		var ret string
		return ret
	}
	return *o.JournalDate
}

// GetJournalDateOk returns a tuple with the JournalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetJournalDateOk() (*string, bool) {
	if o == nil || IsNil(o.JournalDate) {
		return nil, false
	}
	return o.JournalDate, true
}

// HasJournalDateField returns a boolean if a field has been set.
func (o *Journal) HasJournalDateField() bool {
	if o != nil && !IsNil(o.JournalDate) {
		return true
	}

	return false
}

// SetJournalDate gets a reference to the given string and assigns it to the JournalDate field.
func (o *Journal) SetJournalDate(v string) {
	o.JournalDate = &v
}

// GetJournalNumber returns the JournalNumber field value if set, zero value otherwise.
func (o *Journal) GetJournalNumber() int32 {
	if o == nil || IsNil(o.JournalNumber) {
		var ret int32
		return ret
	}
	return *o.JournalNumber
}

// GetJournalNumberOk returns a tuple with the JournalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetJournalNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.JournalNumber) {
		return nil, false
	}
	return o.JournalNumber, true
}

// HasJournalNumberField returns a boolean if a field has been set.
func (o *Journal) HasJournalNumberField() bool {
	if o != nil && !IsNil(o.JournalNumber) {
		return true
	}

	return false
}

// SetJournalNumber gets a reference to the given int32 and assigns it to the JournalNumber field.
func (o *Journal) SetJournalNumber(v int32) {
	o.JournalNumber = &v
}

// GetCreatedDateUTC returns the CreatedDateUTC field value if set, zero value otherwise.
func (o *Journal) GetCreatedDateUTC() string {
	if o == nil || IsNil(o.CreatedDateUTC) {
		var ret string
		return ret
	}
	return *o.CreatedDateUTC
}

// GetCreatedDateUTCOk returns a tuple with the CreatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetCreatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDateUTC) {
		return nil, false
	}
	return o.CreatedDateUTC, true
}

// HasCreatedDateUTCField returns a boolean if a field has been set.
func (o *Journal) HasCreatedDateUTCField() bool {
	if o != nil && !IsNil(o.CreatedDateUTC) {
		return true
	}

	return false
}

// SetCreatedDateUTC gets a reference to the given string and assigns it to the CreatedDateUTC field.
func (o *Journal) SetCreatedDateUTC(v string) {
	o.CreatedDateUTC = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Journal) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReferenceField returns a boolean if a field has been set.
func (o *Journal) HasReferenceField() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Journal) SetReference(v string) {
	o.Reference = &v
}

// GetSourceID returns the SourceID field value if set, zero value otherwise.
func (o *Journal) GetSourceID() string {
	if o == nil || IsNil(o.SourceID) {
		var ret string
		return ret
	}
	return *o.SourceID
}

// GetSourceIDOk returns a tuple with the SourceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetSourceIDOk() (*string, bool) {
	if o == nil || IsNil(o.SourceID) {
		return nil, false
	}
	return o.SourceID, true
}

// HasSourceIDField returns a boolean if a field has been set.
func (o *Journal) HasSourceIDField() bool {
	if o != nil && !IsNil(o.SourceID) {
		return true
	}

	return false
}

// SetSourceID gets a reference to the given string and assigns it to the SourceID field.
func (o *Journal) SetSourceID(v string) {
	o.SourceID = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *Journal) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceTypeField returns a boolean if a field has been set.
func (o *Journal) HasSourceTypeField() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *Journal) SetSourceType(v string) {
	o.SourceType = &v
}

// GetJournalLines returns the JournalLines field value if set, zero value otherwise.
func (o *Journal) GetJournalLines() []JournalLine {
	if o == nil || IsNil(o.JournalLines) {
		var ret []JournalLine
		return ret
	}
	return o.JournalLines
}

// GetJournalLinesOk returns a tuple with the JournalLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetJournalLinesOk() ([]JournalLine, bool) {
	if o == nil || IsNil(o.JournalLines) {
		return nil, false
	}
	return o.JournalLines, true
}

// HasJournalLinesField returns a boolean if a field has been set.
func (o *Journal) HasJournalLinesField() bool {
	if o != nil && !IsNil(o.JournalLines) {
		return true
	}

	return false
}

// SetJournalLines gets a reference to the given []JournalLine and assigns it to the JournalLines field.
func (o *Journal) SetJournalLines(v []JournalLine) {
	o.JournalLines = v
}

func (o Journal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Journal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JournalID) {
		toSerialize["JournalID"] = o.JournalID
	}
	if !IsNil(o.JournalDate) {
		toSerialize["JournalDate"] = o.JournalDate
	}
	if !IsNil(o.JournalNumber) {
		toSerialize["JournalNumber"] = o.JournalNumber
	}
	if !IsNil(o.CreatedDateUTC) {
		toSerialize["CreatedDateUTC"] = o.CreatedDateUTC
	}
	if !IsNil(o.Reference) {
		toSerialize["Reference"] = o.Reference
	}
	if !IsNil(o.SourceID) {
		toSerialize["SourceID"] = o.SourceID
	}
	if !IsNil(o.SourceType) {
		toSerialize["SourceType"] = o.SourceType
	}
	if !IsNil(o.JournalLines) {
		toSerialize["JournalLines"] = o.JournalLines
	}
	return toSerialize, nil
}

type NullableJournal struct {
	value *Journal
	isSet bool
}

func (v NullableJournal) Get() *Journal {
	return v.value
}

func (v *NullableJournal) Set(val *Journal) {
	v.value = val
	v.isSet = true
}

func (v NullableJournal) IsSet() bool {
	return v.isSet
}

func (v *NullableJournal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournal(val *Journal) *NullableJournal {
	return &NullableJournal{value: val, isSet: true}
}

func (v NullableJournal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

