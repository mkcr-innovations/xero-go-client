# ReportCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]ReportAttribute**](ReportAttribute.md) |  | [optional] 

## Methods

### NewReportCell

`func NewReportCell() *ReportCell`

NewReportCell instantiates a new ReportCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCellWithDefaults

`func NewReportCellWithDefaults() *ReportCell`

NewReportCellWithDefaults instantiates a new ReportCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ReportCell) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReportCell) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReportCell) SetValue(v string)`

SetValue sets Value field to given value.

### HasValueField

`func (o *ReportCell) HasValueField() bool`

HasValueField returns a boolean if a field has been set.

### GetAttributes

`func (o *ReportCell) GetAttributes() []ReportAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ReportCell) GetAttributesOk() (*[]ReportAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ReportCell) SetAttributes(v []ReportAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributesField

`func (o *ReportCell) HasAttributesField() bool`

HasAttributesField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


