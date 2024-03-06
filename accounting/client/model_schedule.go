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

// checks if the Schedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule{}

// Schedule struct for Schedule
type Schedule struct {
	// Integer used with the unit e.g. 1 (every 1 week), 2 (every 2 months)
	Period *int32 `json:"Period,omitempty"`
	// One of the following - WEEKLY or MONTHLY
	Unit *string `json:"Unit,omitempty"`
	// Integer used with due date type e.g 20 (of following month), 31 (of current month)
	DueDate *int32 `json:"DueDate,omitempty"`
	// the payment terms
	DueDateType *string `json:"DueDateType,omitempty"`
	// Date the first invoice of the current version of the repeating schedule was generated (changes when repeating invoice is edited)
	StartDate *string `json:"StartDate,omitempty"`
	// The calendar date of the next invoice in the schedule to be generated
	NextScheduledDate *string `json:"NextScheduledDate,omitempty"`
	// Invoice end date – only returned if the template has an end date set
	EndDate *string `json:"EndDate,omitempty"`
}

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule() *Schedule {
	this := Schedule{}
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *Schedule) GetPeriod() int32 {
	if o == nil || IsNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriodField returns a boolean if a field has been set.
func (o *Schedule) HasPeriodField() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *Schedule) SetPeriod(v int32) {
	o.Period = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Schedule) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnitField returns a boolean if a field has been set.
func (o *Schedule) HasUnitField() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Schedule) SetUnit(v string) {
	o.Unit = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *Schedule) GetDueDate() int32 {
	if o == nil || IsNil(o.DueDate) {
		var ret int32
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetDueDateOk() (*int32, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDateField returns a boolean if a field has been set.
func (o *Schedule) HasDueDateField() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given int32 and assigns it to the DueDate field.
func (o *Schedule) SetDueDate(v int32) {
	o.DueDate = &v
}

// GetDueDateType returns the DueDateType field value if set, zero value otherwise.
func (o *Schedule) GetDueDateType() string {
	if o == nil || IsNil(o.DueDateType) {
		var ret string
		return ret
	}
	return *o.DueDateType
}

// GetDueDateTypeOk returns a tuple with the DueDateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetDueDateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DueDateType) {
		return nil, false
	}
	return o.DueDateType, true
}

// HasDueDateTypeField returns a boolean if a field has been set.
func (o *Schedule) HasDueDateTypeField() bool {
	if o != nil && !IsNil(o.DueDateType) {
		return true
	}

	return false
}

// SetDueDateType gets a reference to the given string and assigns it to the DueDateType field.
func (o *Schedule) SetDueDateType(v string) {
	o.DueDateType = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Schedule) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDateField returns a boolean if a field has been set.
func (o *Schedule) HasStartDateField() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *Schedule) SetStartDate(v string) {
	o.StartDate = &v
}

// GetNextScheduledDate returns the NextScheduledDate field value if set, zero value otherwise.
func (o *Schedule) GetNextScheduledDate() string {
	if o == nil || IsNil(o.NextScheduledDate) {
		var ret string
		return ret
	}
	return *o.NextScheduledDate
}

// GetNextScheduledDateOk returns a tuple with the NextScheduledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetNextScheduledDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextScheduledDate) {
		return nil, false
	}
	return o.NextScheduledDate, true
}

// HasNextScheduledDateField returns a boolean if a field has been set.
func (o *Schedule) HasNextScheduledDateField() bool {
	if o != nil && !IsNil(o.NextScheduledDate) {
		return true
	}

	return false
}

// SetNextScheduledDate gets a reference to the given string and assigns it to the NextScheduledDate field.
func (o *Schedule) SetNextScheduledDate(v string) {
	o.NextScheduledDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Schedule) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDateField returns a boolean if a field has been set.
func (o *Schedule) HasEndDateField() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *Schedule) SetEndDate(v string) {
	o.EndDate = &v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["Period"] = o.Period
	}
	if !IsNil(o.Unit) {
		toSerialize["Unit"] = o.Unit
	}
	if !IsNil(o.DueDate) {
		toSerialize["DueDate"] = o.DueDate
	}
	if !IsNil(o.DueDateType) {
		toSerialize["DueDateType"] = o.DueDateType
	}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.NextScheduledDate) {
		toSerialize["NextScheduledDate"] = o.NextScheduledDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["EndDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

