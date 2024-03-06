# BalanceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **float64** | The opening balances of the account. Debits are positive, credits are negative values | [optional] 
**CurrencyCode** | Pointer to **string** | The currency of the balance (Not required for base currency) | [optional] 
**CurrencyRate** | Pointer to **float64** | (Optional) Exchange rate to base currency when money is spent or received. If not specified, XE rate for the day is applied | [optional] 

## Methods

### NewBalanceDetails

`func NewBalanceDetails() *BalanceDetails`

NewBalanceDetails instantiates a new BalanceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceDetailsWithDefaults

`func NewBalanceDetailsWithDefaults() *BalanceDetails`

NewBalanceDetailsWithDefaults instantiates a new BalanceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BalanceDetails) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceDetails) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceDetails) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalanceField

`func (o *BalanceDetails) HasBalanceField() bool`

HasBalanceField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BalanceDetails) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BalanceDetails) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BalanceDetails) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *BalanceDetails) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *BalanceDetails) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *BalanceDetails) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *BalanceDetails) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *BalanceDetails) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


