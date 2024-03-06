# CISSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CISEnabled** | Pointer to **bool** | Boolean that describes if the contact is a CIS Subcontractor | [optional] 
**Rate** | Pointer to **float64** | CIS Deduction rate for the contact if he is a subcontractor. If the contact is not CISEnabled, then the rate is not returned | [optional] [readonly] 

## Methods

### NewCISSetting

`func NewCISSetting() *CISSetting`

NewCISSetting instantiates a new CISSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCISSettingWithDefaults

`func NewCISSettingWithDefaults() *CISSetting`

NewCISSettingWithDefaults instantiates a new CISSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCISEnabled

`func (o *CISSetting) GetCISEnabled() bool`

GetCISEnabled returns the CISEnabled field if non-nil, zero value otherwise.

### GetCISEnabledOk

`func (o *CISSetting) GetCISEnabledOk() (*bool, bool)`

GetCISEnabledOk returns a tuple with the CISEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCISEnabled

`func (o *CISSetting) SetCISEnabled(v bool)`

SetCISEnabled sets CISEnabled field to given value.

### HasCISEnabledField

`func (o *CISSetting) HasCISEnabledField() bool`

HasCISEnabledField returns a boolean if a field has been set.

### GetRate

`func (o *CISSetting) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CISSetting) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CISSetting) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRateField

`func (o *CISSetting) HasRateField() bool`

HasRateField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


