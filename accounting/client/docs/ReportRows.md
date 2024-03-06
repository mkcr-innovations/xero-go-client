# ReportRows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowType** | Pointer to [**RowType**](RowType.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Cells** | Pointer to [**[]ReportCell**](ReportCell.md) |  | [optional] 
**Rows** | Pointer to [**[]ReportRow**](ReportRow.md) |  | [optional] 

## Methods

### NewReportRows

`func NewReportRows() *ReportRows`

NewReportRows instantiates a new ReportRows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportRowsWithDefaults

`func NewReportRowsWithDefaults() *ReportRows`

NewReportRowsWithDefaults instantiates a new ReportRows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowType

`func (o *ReportRows) GetRowType() RowType`

GetRowType returns the RowType field if non-nil, zero value otherwise.

### GetRowTypeOk

`func (o *ReportRows) GetRowTypeOk() (*RowType, bool)`

GetRowTypeOk returns a tuple with the RowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowType

`func (o *ReportRows) SetRowType(v RowType)`

SetRowType sets RowType field to given value.

### HasRowTypeField

`func (o *ReportRows) HasRowTypeField() bool`

HasRowTypeField returns a boolean if a field has been set.

### GetTitle

`func (o *ReportRows) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportRows) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportRows) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitleField

`func (o *ReportRows) HasTitleField() bool`

HasTitleField returns a boolean if a field has been set.

### GetCells

`func (o *ReportRows) GetCells() []ReportCell`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *ReportRows) GetCellsOk() (*[]ReportCell, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *ReportRows) SetCells(v []ReportCell)`

SetCells sets Cells field to given value.

### HasCellsField

`func (o *ReportRows) HasCellsField() bool`

HasCellsField returns a boolean if a field has been set.

### GetRows

`func (o *ReportRows) GetRows() []ReportRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ReportRows) GetRowsOk() (*[]ReportRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ReportRows) SetRows(v []ReportRow)`

SetRows sets Rows field to given value.

### HasRowsField

`func (o *ReportRows) HasRowsField() bool`

HasRowsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


