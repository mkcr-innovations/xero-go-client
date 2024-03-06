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
	"bytes"
	"fmt"
)

// checks if the ManualJournal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualJournal{}

// ManualJournal struct for ManualJournal
type ManualJournal struct {
	// Description of journal being posted
	Narration string `json:"Narration"`
	// See JournalLines
	JournalLines []ManualJournalLine `json:"JournalLines,omitempty"`
	// Date journal was posted – YYYY-MM-DD
	Date *string `json:"Date,omitempty"`
	LineAmountTypes *LineAmountTypes `json:"LineAmountTypes,omitempty"`
	// See Manual Journal Status Codes
	Status *string `json:"Status,omitempty"`
	// Url link to a source document – shown as “Go to [appName]” in the Xero app
	Url *string `json:"Url,omitempty"`
	// Boolean – default is true if not specified
	ShowOnCashBasisReports *bool `json:"ShowOnCashBasisReports,omitempty"`
	// Boolean to indicate if a manual journal has an attachment
	HasAttachments *bool `json:"HasAttachments,omitempty"`
	// Last modified date UTC format
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	// The Xero identifier for a Manual Journal
	ManualJournalID *string `json:"ManualJournalID,omitempty"`
	// A string to indicate if a invoice status
	StatusAttributeString *string `json:"StatusAttributeString,omitempty"`
	// Displays array of warning messages from the API
	Warnings []ValidationError `json:"Warnings,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
	// Displays array of attachments from the API
	Attachments []Attachment `json:"Attachments,omitempty"`
}

type _ManualJournal ManualJournal

// NewManualJournal instantiates a new ManualJournal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualJournal(narration string) *ManualJournal {
	this := ManualJournal{}
	this.Narration = narration
	return &this
}

// NewManualJournalWithDefaults instantiates a new ManualJournal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualJournalWithDefaults() *ManualJournal {
	this := ManualJournal{}
	return &this
}

// GetNarration returns the Narration field value
func (o *ManualJournal) GetNarration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Narration
}

// GetNarrationOk returns a tuple with the Narration field value
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetNarrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Narration, true
}

// SetNarration sets field value
func (o *ManualJournal) SetNarration(v string) {
	o.Narration = v
}

// GetJournalLines returns the JournalLines field value if set, zero value otherwise.
func (o *ManualJournal) GetJournalLines() []ManualJournalLine {
	if o == nil || IsNil(o.JournalLines) {
		var ret []ManualJournalLine
		return ret
	}
	return o.JournalLines
}

// GetJournalLinesOk returns a tuple with the JournalLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetJournalLinesOk() ([]ManualJournalLine, bool) {
	if o == nil || IsNil(o.JournalLines) {
		return nil, false
	}
	return o.JournalLines, true
}

// HasJournalLinesField returns a boolean if a field has been set.
func (o *ManualJournal) HasJournalLinesField() bool {
	if o != nil && !IsNil(o.JournalLines) {
		return true
	}

	return false
}

// SetJournalLines gets a reference to the given []ManualJournalLine and assigns it to the JournalLines field.
func (o *ManualJournal) SetJournalLines(v []ManualJournalLine) {
	o.JournalLines = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ManualJournal) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDateField returns a boolean if a field has been set.
func (o *ManualJournal) HasDateField() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ManualJournal) SetDate(v string) {
	o.Date = &v
}

// GetLineAmountTypes returns the LineAmountTypes field value if set, zero value otherwise.
func (o *ManualJournal) GetLineAmountTypes() LineAmountTypes {
	if o == nil || IsNil(o.LineAmountTypes) {
		var ret LineAmountTypes
		return ret
	}
	return *o.LineAmountTypes
}

// GetLineAmountTypesOk returns a tuple with the LineAmountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetLineAmountTypesOk() (*LineAmountTypes, bool) {
	if o == nil || IsNil(o.LineAmountTypes) {
		return nil, false
	}
	return o.LineAmountTypes, true
}

// HasLineAmountTypesField returns a boolean if a field has been set.
func (o *ManualJournal) HasLineAmountTypesField() bool {
	if o != nil && !IsNil(o.LineAmountTypes) {
		return true
	}

	return false
}

// SetLineAmountTypes gets a reference to the given LineAmountTypes and assigns it to the LineAmountTypes field.
func (o *ManualJournal) SetLineAmountTypes(v LineAmountTypes) {
	o.LineAmountTypes = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ManualJournal) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatusField returns a boolean if a field has been set.
func (o *ManualJournal) HasStatusField() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ManualJournal) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ManualJournal) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrlField returns a boolean if a field has been set.
func (o *ManualJournal) HasUrlField() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ManualJournal) SetUrl(v string) {
	o.Url = &v
}

// GetShowOnCashBasisReports returns the ShowOnCashBasisReports field value if set, zero value otherwise.
func (o *ManualJournal) GetShowOnCashBasisReports() bool {
	if o == nil || IsNil(o.ShowOnCashBasisReports) {
		var ret bool
		return ret
	}
	return *o.ShowOnCashBasisReports
}

// GetShowOnCashBasisReportsOk returns a tuple with the ShowOnCashBasisReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetShowOnCashBasisReportsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowOnCashBasisReports) {
		return nil, false
	}
	return o.ShowOnCashBasisReports, true
}

// HasShowOnCashBasisReportsField returns a boolean if a field has been set.
func (o *ManualJournal) HasShowOnCashBasisReportsField() bool {
	if o != nil && !IsNil(o.ShowOnCashBasisReports) {
		return true
	}

	return false
}

// SetShowOnCashBasisReports gets a reference to the given bool and assigns it to the ShowOnCashBasisReports field.
func (o *ManualJournal) SetShowOnCashBasisReports(v bool) {
	o.ShowOnCashBasisReports = &v
}

// GetHasAttachments returns the HasAttachments field value if set, zero value otherwise.
func (o *ManualJournal) GetHasAttachments() bool {
	if o == nil || IsNil(o.HasAttachments) {
		var ret bool
		return ret
	}
	return *o.HasAttachments
}

// GetHasAttachmentsOk returns a tuple with the HasAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetHasAttachmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAttachments) {
		return nil, false
	}
	return o.HasAttachments, true
}

// HasHasAttachmentsField returns a boolean if a field has been set.
func (o *ManualJournal) HasHasAttachmentsField() bool {
	if o != nil && !IsNil(o.HasAttachments) {
		return true
	}

	return false
}

// SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.
func (o *ManualJournal) SetHasAttachments(v bool) {
	o.HasAttachments = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *ManualJournal) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTCField returns a boolean if a field has been set.
func (o *ManualJournal) HasUpdatedDateUTCField() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *ManualJournal) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetManualJournalID returns the ManualJournalID field value if set, zero value otherwise.
func (o *ManualJournal) GetManualJournalID() string {
	if o == nil || IsNil(o.ManualJournalID) {
		var ret string
		return ret
	}
	return *o.ManualJournalID
}

// GetManualJournalIDOk returns a tuple with the ManualJournalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetManualJournalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ManualJournalID) {
		return nil, false
	}
	return o.ManualJournalID, true
}

// HasManualJournalIDField returns a boolean if a field has been set.
func (o *ManualJournal) HasManualJournalIDField() bool {
	if o != nil && !IsNil(o.ManualJournalID) {
		return true
	}

	return false
}

// SetManualJournalID gets a reference to the given string and assigns it to the ManualJournalID field.
func (o *ManualJournal) SetManualJournalID(v string) {
	o.ManualJournalID = &v
}

// GetStatusAttributeString returns the StatusAttributeString field value if set, zero value otherwise.
func (o *ManualJournal) GetStatusAttributeString() string {
	if o == nil || IsNil(o.StatusAttributeString) {
		var ret string
		return ret
	}
	return *o.StatusAttributeString
}

// GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetStatusAttributeStringOk() (*string, bool) {
	if o == nil || IsNil(o.StatusAttributeString) {
		return nil, false
	}
	return o.StatusAttributeString, true
}

// HasStatusAttributeStringField returns a boolean if a field has been set.
func (o *ManualJournal) HasStatusAttributeStringField() bool {
	if o != nil && !IsNil(o.StatusAttributeString) {
		return true
	}

	return false
}

// SetStatusAttributeString gets a reference to the given string and assigns it to the StatusAttributeString field.
func (o *ManualJournal) SetStatusAttributeString(v string) {
	o.StatusAttributeString = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ManualJournal) GetWarnings() []ValidationError {
	if o == nil || IsNil(o.Warnings) {
		var ret []ValidationError
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetWarningsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarningsField returns a boolean if a field has been set.
func (o *ManualJournal) HasWarningsField() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []ValidationError and assigns it to the Warnings field.
func (o *ManualJournal) SetWarnings(v []ValidationError) {
	o.Warnings = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *ManualJournal) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrorsField returns a boolean if a field has been set.
func (o *ManualJournal) HasValidationErrorsField() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *ManualJournal) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ManualJournal) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualJournal) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachmentsField returns a boolean if a field has been set.
func (o *ManualJournal) HasAttachmentsField() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *ManualJournal) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o ManualJournal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualJournal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Narration"] = o.Narration
	if !IsNil(o.JournalLines) {
		toSerialize["JournalLines"] = o.JournalLines
	}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.LineAmountTypes) {
		toSerialize["LineAmountTypes"] = o.LineAmountTypes
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !IsNil(o.ShowOnCashBasisReports) {
		toSerialize["ShowOnCashBasisReports"] = o.ShowOnCashBasisReports
	}
	if !IsNil(o.HasAttachments) {
		toSerialize["HasAttachments"] = o.HasAttachments
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.ManualJournalID) {
		toSerialize["ManualJournalID"] = o.ManualJournalID
	}
	if !IsNil(o.StatusAttributeString) {
		toSerialize["StatusAttributeString"] = o.StatusAttributeString
	}
	if !IsNil(o.Warnings) {
		toSerialize["Warnings"] = o.Warnings
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	if !IsNil(o.Attachments) {
		toSerialize["Attachments"] = o.Attachments
	}
	return toSerialize, nil
}

func (o *ManualJournal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Narration",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varManualJournal := _ManualJournal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManualJournal)

	if err != nil {
		return err
	}

	*o = ManualJournal(varManualJournal)

	return err
}

type NullableManualJournal struct {
	value *ManualJournal
	isSet bool
}

func (v NullableManualJournal) Get() *ManualJournal {
	return v.value
}

func (v *NullableManualJournal) Set(val *ManualJournal) {
	v.value = val
	v.isSet = true
}

func (v NullableManualJournal) IsSet() bool {
	return v.isSet
}

func (v *NullableManualJournal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualJournal(val *ManualJournal) *NullableManualJournal {
	return &NullableManualJournal{value: val, isSet: true}
}

func (v NullableManualJournal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualJournal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


