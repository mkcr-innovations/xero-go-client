# Purchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitPrice** | Pointer to **float64** | Unit Price of the item. By default UnitPrice is rounded to two decimal places. You can use 4 decimal places by adding the unitdp&#x3D;4 querystring parameter to your request. | [optional] 
**AccountCode** | Pointer to **string** | Default account code to be used for purchased/sale. Not applicable to the purchase details of tracked items | [optional] 
**COGSAccountCode** | Pointer to **string** | Cost of goods sold account. Only applicable to the purchase details of tracked items. | [optional] 
**TaxType** | Pointer to **string** | The tax type from TaxRates | [optional] 

## Methods

### NewPurchase

`func NewPurchase() *Purchase`

NewPurchase instantiates a new Purchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseWithDefaults

`func NewPurchaseWithDefaults() *Purchase`

NewPurchaseWithDefaults instantiates a new Purchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitPrice

`func (o *Purchase) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *Purchase) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *Purchase) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPriceField

`func (o *Purchase) HasUnitPriceField() bool`

HasUnitPriceField returns a boolean if a field has been set.

### GetAccountCode

`func (o *Purchase) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *Purchase) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *Purchase) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.

### HasAccountCodeField

`func (o *Purchase) HasAccountCodeField() bool`

HasAccountCodeField returns a boolean if a field has been set.

### GetCOGSAccountCode

`func (o *Purchase) GetCOGSAccountCode() string`

GetCOGSAccountCode returns the COGSAccountCode field if non-nil, zero value otherwise.

### GetCOGSAccountCodeOk

`func (o *Purchase) GetCOGSAccountCodeOk() (*string, bool)`

GetCOGSAccountCodeOk returns a tuple with the COGSAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCOGSAccountCode

`func (o *Purchase) SetCOGSAccountCode(v string)`

SetCOGSAccountCode sets COGSAccountCode field to given value.

### HasCOGSAccountCodeField

`func (o *Purchase) HasCOGSAccountCodeField() bool`

HasCOGSAccountCodeField returns a boolean if a field has been set.

### GetTaxType

`func (o *Purchase) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *Purchase) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *Purchase) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxTypeField

`func (o *Purchase) HasTaxTypeField() bool`

HasTaxTypeField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


