# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **int32** | Integer used with the unit e.g. 1 (every 1 week), 2 (every 2 months) | [optional] 
**Unit** | Pointer to **string** | One of the following - WEEKLY or MONTHLY | [optional] 
**DueDate** | Pointer to **int32** | Integer used with due date type e.g 20 (of following month), 31 (of current month) | [optional] 
**DueDateType** | Pointer to **string** | the payment terms | [optional] 
**StartDate** | Pointer to **string** | Date the first invoice of the current version of the repeating schedule was generated (changes when repeating invoice is edited) | [optional] 
**NextScheduledDate** | Pointer to **string** | The calendar date of the next invoice in the schedule to be generated | [optional] 
**EndDate** | Pointer to **string** | Invoice end date – only returned if the template has an end date set | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *Schedule) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Schedule) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Schedule) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriodField

`func (o *Schedule) HasPeriodField() bool`

HasPeriodField returns a boolean if a field has been set.

### GetUnit

`func (o *Schedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Schedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Schedule) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnitField

`func (o *Schedule) HasUnitField() bool`

HasUnitField returns a boolean if a field has been set.

### GetDueDate

`func (o *Schedule) GetDueDate() int32`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Schedule) GetDueDateOk() (*int32, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Schedule) SetDueDate(v int32)`

SetDueDate sets DueDate field to given value.

### HasDueDateField

`func (o *Schedule) HasDueDateField() bool`

HasDueDateField returns a boolean if a field has been set.

### GetDueDateType

`func (o *Schedule) GetDueDateType() string`

GetDueDateType returns the DueDateType field if non-nil, zero value otherwise.

### GetDueDateTypeOk

`func (o *Schedule) GetDueDateTypeOk() (*string, bool)`

GetDueDateTypeOk returns a tuple with the DueDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateType

`func (o *Schedule) SetDueDateType(v string)`

SetDueDateType sets DueDateType field to given value.

### HasDueDateTypeField

`func (o *Schedule) HasDueDateTypeField() bool`

HasDueDateTypeField returns a boolean if a field has been set.

### GetStartDate

`func (o *Schedule) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Schedule) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Schedule) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDateField

`func (o *Schedule) HasStartDateField() bool`

HasStartDateField returns a boolean if a field has been set.

### GetNextScheduledDate

`func (o *Schedule) GetNextScheduledDate() string`

GetNextScheduledDate returns the NextScheduledDate field if non-nil, zero value otherwise.

### GetNextScheduledDateOk

`func (o *Schedule) GetNextScheduledDateOk() (*string, bool)`

GetNextScheduledDateOk returns a tuple with the NextScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledDate

`func (o *Schedule) SetNextScheduledDate(v string)`

SetNextScheduledDate sets NextScheduledDate field to given value.

### HasNextScheduledDateField

`func (o *Schedule) HasNextScheduledDateField() bool`

HasNextScheduledDateField returns a boolean if a field has been set.

### GetEndDate

`func (o *Schedule) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Schedule) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Schedule) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDateField

`func (o *Schedule) HasEndDateField() bool`

HasEndDateField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


