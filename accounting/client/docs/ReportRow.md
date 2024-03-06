# ReportRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowType** | Pointer to [**RowType**](RowType.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Cells** | Pointer to [**[]ReportCell**](ReportCell.md) |  | [optional] 

## Methods

### NewReportRow

`func NewReportRow() *ReportRow`

NewReportRow instantiates a new ReportRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportRowWithDefaults

`func NewReportRowWithDefaults() *ReportRow`

NewReportRowWithDefaults instantiates a new ReportRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowType

`func (o *ReportRow) GetRowType() RowType`

GetRowType returns the RowType field if non-nil, zero value otherwise.

### GetRowTypeOk

`func (o *ReportRow) GetRowTypeOk() (*RowType, bool)`

GetRowTypeOk returns a tuple with the RowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowType

`func (o *ReportRow) SetRowType(v RowType)`

SetRowType sets RowType field to given value.

### HasRowTypeField

`func (o *ReportRow) HasRowTypeField() bool`

HasRowTypeField returns a boolean if a field has been set.

### GetTitle

`func (o *ReportRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitleField

`func (o *ReportRow) HasTitleField() bool`

HasTitleField returns a boolean if a field has been set.

### GetCells

`func (o *ReportRow) GetCells() []ReportCell`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *ReportRow) GetCellsOk() (*[]ReportCell, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *ReportRow) SetCells(v []ReportCell)`

SetCells sets Cells field to given value.

### HasCellsField

`func (o *ReportRow) HasCellsField() bool`

HasCellsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


