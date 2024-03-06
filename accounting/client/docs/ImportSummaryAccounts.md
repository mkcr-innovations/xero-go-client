# ImportSummaryAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total number of accounts in the org | [optional] 
**New** | Pointer to **int32** | The number of new accounts created | [optional] 
**Updated** | Pointer to **int32** | The number of accounts updated | [optional] 
**Deleted** | Pointer to **int32** | The number of accounts deleted | [optional] 
**Locked** | Pointer to **int32** | The number of locked accounts | [optional] 
**System** | Pointer to **int32** | The number of system accounts | [optional] 
**Errored** | Pointer to **int32** | The number of accounts that had an error | [optional] 
**Present** | Pointer to **bool** |  | [optional] 
**NewOrUpdated** | Pointer to **int32** | The number of new or updated accounts | [optional] 

## Methods

### NewImportSummaryAccounts

`func NewImportSummaryAccounts() *ImportSummaryAccounts`

NewImportSummaryAccounts instantiates a new ImportSummaryAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportSummaryAccountsWithDefaults

`func NewImportSummaryAccountsWithDefaults() *ImportSummaryAccounts`

NewImportSummaryAccountsWithDefaults instantiates a new ImportSummaryAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ImportSummaryAccounts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ImportSummaryAccounts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ImportSummaryAccounts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *ImportSummaryAccounts) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetNew

`func (o *ImportSummaryAccounts) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *ImportSummaryAccounts) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *ImportSummaryAccounts) SetNew(v int32)`

SetNew sets New field to given value.

### HasNewField

`func (o *ImportSummaryAccounts) HasNewField() bool`

HasNewField returns a boolean if a field has been set.

### GetUpdated

`func (o *ImportSummaryAccounts) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ImportSummaryAccounts) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ImportSummaryAccounts) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdatedField

`func (o *ImportSummaryAccounts) HasUpdatedField() bool`

HasUpdatedField returns a boolean if a field has been set.

### GetDeleted

`func (o *ImportSummaryAccounts) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ImportSummaryAccounts) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ImportSummaryAccounts) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeletedField

`func (o *ImportSummaryAccounts) HasDeletedField() bool`

HasDeletedField returns a boolean if a field has been set.

### GetLocked

`func (o *ImportSummaryAccounts) GetLocked() int32`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ImportSummaryAccounts) GetLockedOk() (*int32, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ImportSummaryAccounts) SetLocked(v int32)`

SetLocked sets Locked field to given value.

### HasLockedField

`func (o *ImportSummaryAccounts) HasLockedField() bool`

HasLockedField returns a boolean if a field has been set.

### GetSystem

`func (o *ImportSummaryAccounts) GetSystem() int32`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ImportSummaryAccounts) GetSystemOk() (*int32, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ImportSummaryAccounts) SetSystem(v int32)`

SetSystem sets System field to given value.

### HasSystemField

`func (o *ImportSummaryAccounts) HasSystemField() bool`

HasSystemField returns a boolean if a field has been set.

### GetErrored

`func (o *ImportSummaryAccounts) GetErrored() int32`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *ImportSummaryAccounts) GetErroredOk() (*int32, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *ImportSummaryAccounts) SetErrored(v int32)`

SetErrored sets Errored field to given value.

### HasErroredField

`func (o *ImportSummaryAccounts) HasErroredField() bool`

HasErroredField returns a boolean if a field has been set.

### GetPresent

`func (o *ImportSummaryAccounts) GetPresent() bool`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *ImportSummaryAccounts) GetPresentOk() (*bool, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *ImportSummaryAccounts) SetPresent(v bool)`

SetPresent sets Present field to given value.

### HasPresentField

`func (o *ImportSummaryAccounts) HasPresentField() bool`

HasPresentField returns a boolean if a field has been set.

### GetNewOrUpdated

`func (o *ImportSummaryAccounts) GetNewOrUpdated() int32`

GetNewOrUpdated returns the NewOrUpdated field if non-nil, zero value otherwise.

### GetNewOrUpdatedOk

`func (o *ImportSummaryAccounts) GetNewOrUpdatedOk() (*int32, bool)`

GetNewOrUpdatedOk returns a tuple with the NewOrUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrUpdated

`func (o *ImportSummaryAccounts) SetNewOrUpdated(v int32)`

SetNewOrUpdated sets NewOrUpdated field to given value.

### HasNewOrUpdatedField

`func (o *ImportSummaryAccounts) HasNewOrUpdatedField() bool`

HasNewOrUpdatedField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


