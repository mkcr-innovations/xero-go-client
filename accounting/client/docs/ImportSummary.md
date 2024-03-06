# ImportSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**ImportSummaryAccounts**](ImportSummaryAccounts.md) |  | [optional] 
**Organisation** | Pointer to [**ImportSummaryOrganisation**](ImportSummaryOrganisation.md) |  | [optional] 

## Methods

### NewImportSummary

`func NewImportSummary() *ImportSummary`

NewImportSummary instantiates a new ImportSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportSummaryWithDefaults

`func NewImportSummaryWithDefaults() *ImportSummary`

NewImportSummaryWithDefaults instantiates a new ImportSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ImportSummary) GetAccounts() ImportSummaryAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ImportSummary) GetAccountsOk() (*ImportSummaryAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ImportSummary) SetAccounts(v ImportSummaryAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccountsField

`func (o *ImportSummary) HasAccountsField() bool`

HasAccountsField returns a boolean if a field has been set.

### GetOrganisation

`func (o *ImportSummary) GetOrganisation() ImportSummaryOrganisation`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *ImportSummary) GetOrganisationOk() (*ImportSummaryOrganisation, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *ImportSummary) SetOrganisation(v ImportSummaryOrganisation)`

SetOrganisation sets Organisation field to given value.

### HasOrganisationField

`func (o *ImportSummary) HasOrganisationField() bool`

HasOrganisationField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


