# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**Description** | Pointer to **string** | Name of Currency | [optional] 

## Methods

### NewCurrency

`func NewCurrency() *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Currency) GetCode() CurrencyCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Currency) GetCodeOk() (*CurrencyCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Currency) SetCode(v CurrencyCode)`

SetCode sets Code field to given value.

### HasCodeField

`func (o *Currency) HasCodeField() bool`

HasCodeField returns a boolean if a field has been set.

### GetDescription

`func (o *Currency) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Currency) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Currency) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescriptionField

`func (o *Currency) HasDescriptionField() bool`

HasDescriptionField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


