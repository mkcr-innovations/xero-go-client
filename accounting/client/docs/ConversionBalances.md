# ConversionBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCode** | Pointer to **string** | The account code for a account | [optional] 
**Balance** | Pointer to **float64** | The opening balances of the account. Debits are positive, credits are negative values | [optional] 
**BalanceDetails** | Pointer to [**[]BalanceDetails**](BalanceDetails.md) |  | [optional] 

## Methods

### NewConversionBalances

`func NewConversionBalances() *ConversionBalances`

NewConversionBalances instantiates a new ConversionBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionBalancesWithDefaults

`func NewConversionBalancesWithDefaults() *ConversionBalances`

NewConversionBalancesWithDefaults instantiates a new ConversionBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCode

`func (o *ConversionBalances) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *ConversionBalances) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *ConversionBalances) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.

### HasAccountCodeField

`func (o *ConversionBalances) HasAccountCodeField() bool`

HasAccountCodeField returns a boolean if a field has been set.

### GetBalance

`func (o *ConversionBalances) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ConversionBalances) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ConversionBalances) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalanceField

`func (o *ConversionBalances) HasBalanceField() bool`

HasBalanceField returns a boolean if a field has been set.

### GetBalanceDetails

`func (o *ConversionBalances) GetBalanceDetails() []BalanceDetails`

GetBalanceDetails returns the BalanceDetails field if non-nil, zero value otherwise.

### GetBalanceDetailsOk

`func (o *ConversionBalances) GetBalanceDetailsOk() (*[]BalanceDetails, bool)`

GetBalanceDetailsOk returns a tuple with the BalanceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDetails

`func (o *ConversionBalances) SetBalanceDetails(v []BalanceDetails)`

SetBalanceDetails sets BalanceDetails field to given value.

### HasBalanceDetailsField

`func (o *ConversionBalances) HasBalanceDetailsField() bool`

HasBalanceDetailsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


