# TaxRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of tax rate | [optional] 
**TaxType** | Pointer to **string** | The tax type | [optional] 
**TaxComponents** | Pointer to [**[]TaxComponent**](TaxComponent.md) | See TaxComponents | [optional] 
**Status** | Pointer to **string** | See Status Codes | [optional] 
**ReportTaxType** | Pointer to **string** | See ReportTaxTypes | [optional] 
**CanApplyToAssets** | Pointer to **bool** | Boolean to describe if tax rate can be used for asset accounts i.e.  true,false | [optional] [readonly] 
**CanApplyToEquity** | Pointer to **bool** | Boolean to describe if tax rate can be used for equity accounts i.e true,false | [optional] [readonly] 
**CanApplyToExpenses** | Pointer to **bool** | Boolean to describe if tax rate can be used for expense accounts  i.e. true,false | [optional] [readonly] 
**CanApplyToLiabilities** | Pointer to **bool** | Boolean to describe if tax rate can be used for liability accounts  i.e. true,false | [optional] [readonly] 
**CanApplyToRevenue** | Pointer to **bool** | Boolean to describe if tax rate can be used for revenue accounts i.e. true,false | [optional] [readonly] 
**DisplayTaxRate** | Pointer to **float64** | Tax Rate (decimal to 4dp) e.g 12.5000 | [optional] [readonly] 
**EffectiveRate** | Pointer to **float64** | Effective Tax Rate (decimal to 4dp) e.g 12.5000 | [optional] [readonly] 

## Methods

### NewTaxRate

`func NewTaxRate() *TaxRate`

NewTaxRate instantiates a new TaxRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRateWithDefaults

`func NewTaxRateWithDefaults() *TaxRate`

NewTaxRateWithDefaults instantiates a new TaxRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaxRate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxRate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxRate) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *TaxRate) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetTaxType

`func (o *TaxRate) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *TaxRate) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *TaxRate) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxTypeField

`func (o *TaxRate) HasTaxTypeField() bool`

HasTaxTypeField returns a boolean if a field has been set.

### GetTaxComponents

`func (o *TaxRate) GetTaxComponents() []TaxComponent`

GetTaxComponents returns the TaxComponents field if non-nil, zero value otherwise.

### GetTaxComponentsOk

`func (o *TaxRate) GetTaxComponentsOk() (*[]TaxComponent, bool)`

GetTaxComponentsOk returns a tuple with the TaxComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxComponents

`func (o *TaxRate) SetTaxComponents(v []TaxComponent)`

SetTaxComponents sets TaxComponents field to given value.

### HasTaxComponentsField

`func (o *TaxRate) HasTaxComponentsField() bool`

HasTaxComponentsField returns a boolean if a field has been set.

### GetStatus

`func (o *TaxRate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaxRate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaxRate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *TaxRate) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetReportTaxType

`func (o *TaxRate) GetReportTaxType() string`

GetReportTaxType returns the ReportTaxType field if non-nil, zero value otherwise.

### GetReportTaxTypeOk

`func (o *TaxRate) GetReportTaxTypeOk() (*string, bool)`

GetReportTaxTypeOk returns a tuple with the ReportTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTaxType

`func (o *TaxRate) SetReportTaxType(v string)`

SetReportTaxType sets ReportTaxType field to given value.

### HasReportTaxTypeField

`func (o *TaxRate) HasReportTaxTypeField() bool`

HasReportTaxTypeField returns a boolean if a field has been set.

### GetCanApplyToAssets

`func (o *TaxRate) GetCanApplyToAssets() bool`

GetCanApplyToAssets returns the CanApplyToAssets field if non-nil, zero value otherwise.

### GetCanApplyToAssetsOk

`func (o *TaxRate) GetCanApplyToAssetsOk() (*bool, bool)`

GetCanApplyToAssetsOk returns a tuple with the CanApplyToAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyToAssets

`func (o *TaxRate) SetCanApplyToAssets(v bool)`

SetCanApplyToAssets sets CanApplyToAssets field to given value.

### HasCanApplyToAssetsField

`func (o *TaxRate) HasCanApplyToAssetsField() bool`

HasCanApplyToAssetsField returns a boolean if a field has been set.

### GetCanApplyToEquity

`func (o *TaxRate) GetCanApplyToEquity() bool`

GetCanApplyToEquity returns the CanApplyToEquity field if non-nil, zero value otherwise.

### GetCanApplyToEquityOk

`func (o *TaxRate) GetCanApplyToEquityOk() (*bool, bool)`

GetCanApplyToEquityOk returns a tuple with the CanApplyToEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyToEquity

`func (o *TaxRate) SetCanApplyToEquity(v bool)`

SetCanApplyToEquity sets CanApplyToEquity field to given value.

### HasCanApplyToEquityField

`func (o *TaxRate) HasCanApplyToEquityField() bool`

HasCanApplyToEquityField returns a boolean if a field has been set.

### GetCanApplyToExpenses

`func (o *TaxRate) GetCanApplyToExpenses() bool`

GetCanApplyToExpenses returns the CanApplyToExpenses field if non-nil, zero value otherwise.

### GetCanApplyToExpensesOk

`func (o *TaxRate) GetCanApplyToExpensesOk() (*bool, bool)`

GetCanApplyToExpensesOk returns a tuple with the CanApplyToExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyToExpenses

`func (o *TaxRate) SetCanApplyToExpenses(v bool)`

SetCanApplyToExpenses sets CanApplyToExpenses field to given value.

### HasCanApplyToExpensesField

`func (o *TaxRate) HasCanApplyToExpensesField() bool`

HasCanApplyToExpensesField returns a boolean if a field has been set.

### GetCanApplyToLiabilities

`func (o *TaxRate) GetCanApplyToLiabilities() bool`

GetCanApplyToLiabilities returns the CanApplyToLiabilities field if non-nil, zero value otherwise.

### GetCanApplyToLiabilitiesOk

`func (o *TaxRate) GetCanApplyToLiabilitiesOk() (*bool, bool)`

GetCanApplyToLiabilitiesOk returns a tuple with the CanApplyToLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyToLiabilities

`func (o *TaxRate) SetCanApplyToLiabilities(v bool)`

SetCanApplyToLiabilities sets CanApplyToLiabilities field to given value.

### HasCanApplyToLiabilitiesField

`func (o *TaxRate) HasCanApplyToLiabilitiesField() bool`

HasCanApplyToLiabilitiesField returns a boolean if a field has been set.

### GetCanApplyToRevenue

`func (o *TaxRate) GetCanApplyToRevenue() bool`

GetCanApplyToRevenue returns the CanApplyToRevenue field if non-nil, zero value otherwise.

### GetCanApplyToRevenueOk

`func (o *TaxRate) GetCanApplyToRevenueOk() (*bool, bool)`

GetCanApplyToRevenueOk returns a tuple with the CanApplyToRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyToRevenue

`func (o *TaxRate) SetCanApplyToRevenue(v bool)`

SetCanApplyToRevenue sets CanApplyToRevenue field to given value.

### HasCanApplyToRevenueField

`func (o *TaxRate) HasCanApplyToRevenueField() bool`

HasCanApplyToRevenueField returns a boolean if a field has been set.

### GetDisplayTaxRate

`func (o *TaxRate) GetDisplayTaxRate() float64`

GetDisplayTaxRate returns the DisplayTaxRate field if non-nil, zero value otherwise.

### GetDisplayTaxRateOk

`func (o *TaxRate) GetDisplayTaxRateOk() (*float64, bool)`

GetDisplayTaxRateOk returns a tuple with the DisplayTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTaxRate

`func (o *TaxRate) SetDisplayTaxRate(v float64)`

SetDisplayTaxRate sets DisplayTaxRate field to given value.

### HasDisplayTaxRateField

`func (o *TaxRate) HasDisplayTaxRateField() bool`

HasDisplayTaxRateField returns a boolean if a field has been set.

### GetEffectiveRate

`func (o *TaxRate) GetEffectiveRate() float64`

GetEffectiveRate returns the EffectiveRate field if non-nil, zero value otherwise.

### GetEffectiveRateOk

`func (o *TaxRate) GetEffectiveRateOk() (*float64, bool)`

GetEffectiveRateOk returns a tuple with the EffectiveRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRate

`func (o *TaxRate) SetEffectiveRate(v float64)`

SetEffectiveRate sets EffectiveRate field to given value.

### HasEffectiveRateField

`func (o *TaxRate) HasEffectiveRateField() bool`

HasEffectiveRateField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


