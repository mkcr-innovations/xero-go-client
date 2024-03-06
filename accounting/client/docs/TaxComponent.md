# TaxComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of Tax Component | [optional] 
**Rate** | Pointer to **float64** | Tax Rate (up to 4dp) | [optional] 
**IsCompound** | Pointer to **bool** | Boolean to describe if Tax rate is compounded. | [optional] 
**IsNonRecoverable** | Pointer to **bool** | Boolean to describe if tax rate is non-recoverable. Non-recoverable rates are only applicable to Canadian organisations | [optional] 

## Methods

### NewTaxComponent

`func NewTaxComponent() *TaxComponent`

NewTaxComponent instantiates a new TaxComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxComponentWithDefaults

`func NewTaxComponentWithDefaults() *TaxComponent`

NewTaxComponentWithDefaults instantiates a new TaxComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaxComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxComponent) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *TaxComponent) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetRate

`func (o *TaxComponent) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxComponent) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxComponent) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRateField

`func (o *TaxComponent) HasRateField() bool`

HasRateField returns a boolean if a field has been set.

### GetIsCompound

`func (o *TaxComponent) GetIsCompound() bool`

GetIsCompound returns the IsCompound field if non-nil, zero value otherwise.

### GetIsCompoundOk

`func (o *TaxComponent) GetIsCompoundOk() (*bool, bool)`

GetIsCompoundOk returns a tuple with the IsCompound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompound

`func (o *TaxComponent) SetIsCompound(v bool)`

SetIsCompound sets IsCompound field to given value.

### HasIsCompoundField

`func (o *TaxComponent) HasIsCompoundField() bool`

HasIsCompoundField returns a boolean if a field has been set.

### GetIsNonRecoverable

`func (o *TaxComponent) GetIsNonRecoverable() bool`

GetIsNonRecoverable returns the IsNonRecoverable field if non-nil, zero value otherwise.

### GetIsNonRecoverableOk

`func (o *TaxComponent) GetIsNonRecoverableOk() (*bool, bool)`

GetIsNonRecoverableOk returns a tuple with the IsNonRecoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNonRecoverable

`func (o *TaxComponent) SetIsNonRecoverable(v bool)`

SetIsNonRecoverable sets IsNonRecoverable field to given value.

### HasIsNonRecoverableField

`func (o *TaxComponent) HasIsNonRecoverableField() bool`

HasIsNonRecoverableField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


