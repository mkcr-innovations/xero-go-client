# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportName** | Pointer to **string** | See Prepayment Types | [optional] 
**ReportType** | Pointer to **string** | See Prepayment Types | [optional] 
**ReportTitle** | Pointer to **string** | See Prepayment Types | [optional] 
**ReportDate** | Pointer to **string** | Date of report | [optional] 
**UpdatedDateUTC** | Pointer to **string** | Updated Date | [optional] [readonly] 
**Contacts** | Pointer to [**[]TenNinetyNineContact**](TenNinetyNineContact.md) |  | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportName

`func (o *Report) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *Report) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *Report) SetReportName(v string)`

SetReportName sets ReportName field to given value.

### HasReportNameField

`func (o *Report) HasReportNameField() bool`

HasReportNameField returns a boolean if a field has been set.

### GetReportType

`func (o *Report) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *Report) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *Report) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportTypeField

`func (o *Report) HasReportTypeField() bool`

HasReportTypeField returns a boolean if a field has been set.

### GetReportTitle

`func (o *Report) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *Report) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *Report) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitleField

`func (o *Report) HasReportTitleField() bool`

HasReportTitleField returns a boolean if a field has been set.

### GetReportDate

`func (o *Report) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *Report) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *Report) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDateField

`func (o *Report) HasReportDateField() bool`

HasReportDateField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Report) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Report) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Report) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Report) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetContacts

`func (o *Report) GetContacts() []TenNinetyNineContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Report) GetContactsOk() (*[]TenNinetyNineContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Report) SetContacts(v []TenNinetyNineContact)`

SetContacts sets Contacts field to given value.

### HasContactsField

`func (o *Report) HasContactsField() bool`

HasContactsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


