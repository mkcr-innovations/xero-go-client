# CISOrgSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CISContractorEnabled** | Pointer to **bool** | true or false - Boolean that describes if the organisation is a CIS Contractor | [optional] 
**CISSubContractorEnabled** | Pointer to **bool** | true or false - Boolean that describes if the organisation is a CIS SubContractor | [optional] 
**Rate** | Pointer to **float64** | CIS Deduction rate for the organisation | [optional] [readonly] 

## Methods

### NewCISOrgSetting

`func NewCISOrgSetting() *CISOrgSetting`

NewCISOrgSetting instantiates a new CISOrgSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCISOrgSettingWithDefaults

`func NewCISOrgSettingWithDefaults() *CISOrgSetting`

NewCISOrgSettingWithDefaults instantiates a new CISOrgSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCISContractorEnabled

`func (o *CISOrgSetting) GetCISContractorEnabled() bool`

GetCISContractorEnabled returns the CISContractorEnabled field if non-nil, zero value otherwise.

### GetCISContractorEnabledOk

`func (o *CISOrgSetting) GetCISContractorEnabledOk() (*bool, bool)`

GetCISContractorEnabledOk returns a tuple with the CISContractorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCISContractorEnabled

`func (o *CISOrgSetting) SetCISContractorEnabled(v bool)`

SetCISContractorEnabled sets CISContractorEnabled field to given value.

### HasCISContractorEnabledField

`func (o *CISOrgSetting) HasCISContractorEnabledField() bool`

HasCISContractorEnabledField returns a boolean if a field has been set.

### GetCISSubContractorEnabled

`func (o *CISOrgSetting) GetCISSubContractorEnabled() bool`

GetCISSubContractorEnabled returns the CISSubContractorEnabled field if non-nil, zero value otherwise.

### GetCISSubContractorEnabledOk

`func (o *CISOrgSetting) GetCISSubContractorEnabledOk() (*bool, bool)`

GetCISSubContractorEnabledOk returns a tuple with the CISSubContractorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCISSubContractorEnabled

`func (o *CISOrgSetting) SetCISSubContractorEnabled(v bool)`

SetCISSubContractorEnabled sets CISSubContractorEnabled field to given value.

### HasCISSubContractorEnabledField

`func (o *CISOrgSetting) HasCISSubContractorEnabledField() bool`

HasCISSubContractorEnabledField returns a boolean if a field has been set.

### GetRate

`func (o *CISOrgSetting) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CISOrgSetting) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CISOrgSetting) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRateField

`func (o *CISOrgSetting) HasRateField() bool`

HasRateField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


