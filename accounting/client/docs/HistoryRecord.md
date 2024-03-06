# HistoryRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **string** | details | [optional] 
**Changes** | Pointer to **string** | Name of branding theme | [optional] 
**User** | Pointer to **string** | has a value of 0 | [optional] 
**DateUTC** | Pointer to **string** | UTC timestamp of creation date of branding theme | [optional] [readonly] 

## Methods

### NewHistoryRecord

`func NewHistoryRecord() *HistoryRecord`

NewHistoryRecord instantiates a new HistoryRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryRecordWithDefaults

`func NewHistoryRecordWithDefaults() *HistoryRecord`

NewHistoryRecordWithDefaults instantiates a new HistoryRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *HistoryRecord) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HistoryRecord) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HistoryRecord) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetailsField

`func (o *HistoryRecord) HasDetailsField() bool`

HasDetailsField returns a boolean if a field has been set.

### GetChanges

`func (o *HistoryRecord) GetChanges() string`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *HistoryRecord) GetChangesOk() (*string, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *HistoryRecord) SetChanges(v string)`

SetChanges sets Changes field to given value.

### HasChangesField

`func (o *HistoryRecord) HasChangesField() bool`

HasChangesField returns a boolean if a field has been set.

### GetUser

`func (o *HistoryRecord) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HistoryRecord) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HistoryRecord) SetUser(v string)`

SetUser sets User field to given value.

### HasUserField

`func (o *HistoryRecord) HasUserField() bool`

HasUserField returns a boolean if a field has been set.

### GetDateUTC

`func (o *HistoryRecord) GetDateUTC() string`

GetDateUTC returns the DateUTC field if non-nil, zero value otherwise.

### GetDateUTCOk

`func (o *HistoryRecord) GetDateUTCOk() (*string, bool)`

GetDateUTCOk returns a tuple with the DateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUTC

`func (o *HistoryRecord) SetDateUTC(v string)`

SetDateUTC sets DateUTC field to given value.

### HasDateUTCField

`func (o *HistoryRecord) HasDateUTCField() bool`

HasDateUTCField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


