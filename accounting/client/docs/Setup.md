# Setup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionDate** | Pointer to [**ConversionDate**](ConversionDate.md) |  | [optional] 
**ConversionBalances** | Pointer to [**[]ConversionBalances**](ConversionBalances.md) | Balance supplied for each account that has a value as at the conversion date. | [optional] 
**Accounts** | Pointer to [**[]Account**](Account.md) |  | [optional] 

## Methods

### NewSetup

`func NewSetup() *Setup`

NewSetup instantiates a new Setup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupWithDefaults

`func NewSetupWithDefaults() *Setup`

NewSetupWithDefaults instantiates a new Setup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionDate

`func (o *Setup) GetConversionDate() ConversionDate`

GetConversionDate returns the ConversionDate field if non-nil, zero value otherwise.

### GetConversionDateOk

`func (o *Setup) GetConversionDateOk() (*ConversionDate, bool)`

GetConversionDateOk returns a tuple with the ConversionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionDate

`func (o *Setup) SetConversionDate(v ConversionDate)`

SetConversionDate sets ConversionDate field to given value.

### HasConversionDateField

`func (o *Setup) HasConversionDateField() bool`

HasConversionDateField returns a boolean if a field has been set.

### GetConversionBalances

`func (o *Setup) GetConversionBalances() []ConversionBalances`

GetConversionBalances returns the ConversionBalances field if non-nil, zero value otherwise.

### GetConversionBalancesOk

`func (o *Setup) GetConversionBalancesOk() (*[]ConversionBalances, bool)`

GetConversionBalancesOk returns a tuple with the ConversionBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionBalances

`func (o *Setup) SetConversionBalances(v []ConversionBalances)`

SetConversionBalances sets ConversionBalances field to given value.

### HasConversionBalancesField

`func (o *Setup) HasConversionBalancesField() bool`

HasConversionBalancesField returns a boolean if a field has been set.

### GetAccounts

`func (o *Setup) GetAccounts() []Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Setup) GetAccountsOk() (*[]Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Setup) SetAccounts(v []Account)`

SetAccounts sets Accounts field to given value.

### HasAccountsField

`func (o *Setup) HasAccountsField() bool`

HasAccountsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


