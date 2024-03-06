# ReportWithRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportID** | Pointer to **string** | ID of the Report | [optional] 
**ReportName** | Pointer to **string** | Name of the report | [optional] 
**ReportTitle** | Pointer to **string** | Title of the report | [optional] 
**ReportType** | Pointer to **string** | The type of report (BalanceSheet,ProfitLoss, etc) | [optional] 
**ReportTitles** | Pointer to **[]string** | Report titles array (3 to 4 strings with the report name, orgnisation name and time frame of report) | [optional] 
**ReportDate** | Pointer to **string** | Date of report | [optional] 
**Rows** | Pointer to [**[]ReportRows**](ReportRows.md) |  | [optional] 
**UpdatedDateUTC** | Pointer to **string** | Updated Date | [optional] [readonly] 
**Fields** | Pointer to [**[]ReportFields**](ReportFields.md) |  | [optional] 

## Methods

### NewReportWithRow

`func NewReportWithRow() *ReportWithRow`

NewReportWithRow instantiates a new ReportWithRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithRowWithDefaults

`func NewReportWithRowWithDefaults() *ReportWithRow`

NewReportWithRowWithDefaults instantiates a new ReportWithRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportID

`func (o *ReportWithRow) GetReportID() string`

GetReportID returns the ReportID field if non-nil, zero value otherwise.

### GetReportIDOk

`func (o *ReportWithRow) GetReportIDOk() (*string, bool)`

GetReportIDOk returns a tuple with the ReportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportID

`func (o *ReportWithRow) SetReportID(v string)`

SetReportID sets ReportID field to given value.

### HasReportIDField

`func (o *ReportWithRow) HasReportIDField() bool`

HasReportIDField returns a boolean if a field has been set.

### GetReportName

`func (o *ReportWithRow) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *ReportWithRow) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *ReportWithRow) SetReportName(v string)`

SetReportName sets ReportName field to given value.

### HasReportNameField

`func (o *ReportWithRow) HasReportNameField() bool`

HasReportNameField returns a boolean if a field has been set.

### GetReportTitle

`func (o *ReportWithRow) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *ReportWithRow) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *ReportWithRow) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitleField

`func (o *ReportWithRow) HasReportTitleField() bool`

HasReportTitleField returns a boolean if a field has been set.

### GetReportType

`func (o *ReportWithRow) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportWithRow) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportWithRow) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportTypeField

`func (o *ReportWithRow) HasReportTypeField() bool`

HasReportTypeField returns a boolean if a field has been set.

### GetReportTitles

`func (o *ReportWithRow) GetReportTitles() []string`

GetReportTitles returns the ReportTitles field if non-nil, zero value otherwise.

### GetReportTitlesOk

`func (o *ReportWithRow) GetReportTitlesOk() (*[]string, bool)`

GetReportTitlesOk returns a tuple with the ReportTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitles

`func (o *ReportWithRow) SetReportTitles(v []string)`

SetReportTitles sets ReportTitles field to given value.

### HasReportTitlesField

`func (o *ReportWithRow) HasReportTitlesField() bool`

HasReportTitlesField returns a boolean if a field has been set.

### GetReportDate

`func (o *ReportWithRow) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *ReportWithRow) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *ReportWithRow) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDateField

`func (o *ReportWithRow) HasReportDateField() bool`

HasReportDateField returns a boolean if a field has been set.

### GetRows

`func (o *ReportWithRow) GetRows() []ReportRows`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ReportWithRow) GetRowsOk() (*[]ReportRows, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ReportWithRow) SetRows(v []ReportRows)`

SetRows sets Rows field to given value.

### HasRowsField

`func (o *ReportWithRow) HasRowsField() bool`

HasRowsField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *ReportWithRow) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *ReportWithRow) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *ReportWithRow) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *ReportWithRow) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetFields

`func (o *ReportWithRow) GetFields() []ReportFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ReportWithRow) GetFieldsOk() (*[]ReportFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ReportWithRow) SetFields(v []ReportFields)`

SetFields sets Fields field to given value.

### HasFieldsField

`func (o *ReportWithRow) HasFieldsField() bool`

HasFieldsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


