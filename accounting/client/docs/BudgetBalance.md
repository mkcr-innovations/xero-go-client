# BudgetBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **string** | Period the amount applies to (e.g. “2019-08”) | [optional] 
**Amount** | Pointer to **float64** | LineItem Quantity | [optional] 
**UnitAmount** | Pointer to **float64** | Budgeted amount | [optional] 
**Notes** | Pointer to **string** | Any footnotes associated with this balance | [optional] 

## Methods

### NewBudgetBalance

`func NewBudgetBalance() *BudgetBalance`

NewBudgetBalance instantiates a new BudgetBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetBalanceWithDefaults

`func NewBudgetBalanceWithDefaults() *BudgetBalance`

NewBudgetBalanceWithDefaults instantiates a new BudgetBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *BudgetBalance) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BudgetBalance) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BudgetBalance) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriodField

`func (o *BudgetBalance) HasPeriodField() bool`

HasPeriodField returns a boolean if a field has been set.

### GetAmount

`func (o *BudgetBalance) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BudgetBalance) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BudgetBalance) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmountField

`func (o *BudgetBalance) HasAmountField() bool`

HasAmountField returns a boolean if a field has been set.

### GetUnitAmount

`func (o *BudgetBalance) GetUnitAmount() float64`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *BudgetBalance) GetUnitAmountOk() (*float64, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *BudgetBalance) SetUnitAmount(v float64)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmountField

`func (o *BudgetBalance) HasUnitAmountField() bool`

HasUnitAmountField returns a boolean if a field has been set.

### GetNotes

`func (o *BudgetBalance) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BudgetBalance) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BudgetBalance) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotesField

`func (o *BudgetBalance) HasNotesField() bool`

HasNotesField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


