# BudgetLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | Pointer to **string** | See Accounts | [optional] 
**AccountCode** | Pointer to **string** | See Accounts | [optional] 
**BudgetBalances** | Pointer to [**[]BudgetBalance**](BudgetBalance.md) |  | [optional] 

## Methods

### NewBudgetLine

`func NewBudgetLine() *BudgetLine`

NewBudgetLine instantiates a new BudgetLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetLineWithDefaults

`func NewBudgetLineWithDefaults() *BudgetLine`

NewBudgetLineWithDefaults instantiates a new BudgetLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountID

`func (o *BudgetLine) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *BudgetLine) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *BudgetLine) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountIDField

`func (o *BudgetLine) HasAccountIDField() bool`

HasAccountIDField returns a boolean if a field has been set.

### GetAccountCode

`func (o *BudgetLine) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *BudgetLine) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *BudgetLine) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.

### HasAccountCodeField

`func (o *BudgetLine) HasAccountCodeField() bool`

HasAccountCodeField returns a boolean if a field has been set.

### GetBudgetBalances

`func (o *BudgetLine) GetBudgetBalances() []BudgetBalance`

GetBudgetBalances returns the BudgetBalances field if non-nil, zero value otherwise.

### GetBudgetBalancesOk

`func (o *BudgetLine) GetBudgetBalancesOk() (*[]BudgetBalance, bool)`

GetBudgetBalancesOk returns a tuple with the BudgetBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetBalances

`func (o *BudgetLine) SetBudgetBalances(v []BudgetBalance)`

SetBudgetBalances sets BudgetBalances field to given value.

### HasBudgetBalancesField

`func (o *BudgetLine) HasBudgetBalancesField() bool`

HasBudgetBalancesField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


