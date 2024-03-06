# ConversionDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | Pointer to **int32** | The month the organisation starts using Xero. Value is an integer between 1 and 12 | [optional] 
**Year** | Pointer to **int32** | The year the organisation starts using Xero. Value is an integer greater than 2006 | [optional] 

## Methods

### NewConversionDate

`func NewConversionDate() *ConversionDate`

NewConversionDate instantiates a new ConversionDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionDateWithDefaults

`func NewConversionDateWithDefaults() *ConversionDate`

NewConversionDateWithDefaults instantiates a new ConversionDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *ConversionDate) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ConversionDate) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ConversionDate) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonthField

`func (o *ConversionDate) HasMonthField() bool`

HasMonthField returns a boolean if a field has been set.

### GetYear

`func (o *ConversionDate) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ConversionDate) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ConversionDate) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYearField

`func (o *ConversionDate) HasYearField() bool`

HasYearField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


