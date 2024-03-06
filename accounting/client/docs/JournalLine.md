# JournalLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JournalLineID** | Pointer to **string** | Xero identifier for Journal | [optional] 
**AccountID** | Pointer to **string** | See Accounts | [optional] 
**AccountCode** | Pointer to **string** | See Accounts | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**AccountName** | Pointer to **string** | See AccountCodes | [optional] 
**Description** | Pointer to **string** | The description from the source transaction line item. Only returned if populated. | [optional] 
**NetAmount** | Pointer to **float64** | Net amount of journal line. This will be a positive value for a debit and negative for a credit | [optional] 
**GrossAmount** | Pointer to **float64** | Gross amount of journal line (NetAmount + TaxAmount). | [optional] 
**TaxAmount** | Pointer to **float64** | Total tax on a journal line | [optional] [readonly] 
**TaxType** | Pointer to **string** | The tax type from taxRates | [optional] 
**TaxName** | Pointer to **string** | see TaxRates | [optional] 
**TrackingCategories** | Pointer to [**[]TrackingCategory**](TrackingCategory.md) | Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 &lt;TrackingCategory&gt; elements. | [optional] 

## Methods

### NewJournalLine

`func NewJournalLine() *JournalLine`

NewJournalLine instantiates a new JournalLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalLineWithDefaults

`func NewJournalLineWithDefaults() *JournalLine`

NewJournalLineWithDefaults instantiates a new JournalLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJournalLineID

`func (o *JournalLine) GetJournalLineID() string`

GetJournalLineID returns the JournalLineID field if non-nil, zero value otherwise.

### GetJournalLineIDOk

`func (o *JournalLine) GetJournalLineIDOk() (*string, bool)`

GetJournalLineIDOk returns a tuple with the JournalLineID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalLineID

`func (o *JournalLine) SetJournalLineID(v string)`

SetJournalLineID sets JournalLineID field to given value.

### HasJournalLineIDField

`func (o *JournalLine) HasJournalLineIDField() bool`

HasJournalLineIDField returns a boolean if a field has been set.

### GetAccountID

`func (o *JournalLine) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *JournalLine) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *JournalLine) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountIDField

`func (o *JournalLine) HasAccountIDField() bool`

HasAccountIDField returns a boolean if a field has been set.

### GetAccountCode

`func (o *JournalLine) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *JournalLine) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *JournalLine) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.

### HasAccountCodeField

`func (o *JournalLine) HasAccountCodeField() bool`

HasAccountCodeField returns a boolean if a field has been set.

### GetAccountType

`func (o *JournalLine) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *JournalLine) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *JournalLine) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountTypeField

`func (o *JournalLine) HasAccountTypeField() bool`

HasAccountTypeField returns a boolean if a field has been set.

### GetAccountName

`func (o *JournalLine) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *JournalLine) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *JournalLine) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountNameField

`func (o *JournalLine) HasAccountNameField() bool`

HasAccountNameField returns a boolean if a field has been set.

### GetDescription

`func (o *JournalLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescriptionField

`func (o *JournalLine) HasDescriptionField() bool`

HasDescriptionField returns a boolean if a field has been set.

### GetNetAmount

`func (o *JournalLine) GetNetAmount() float64`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *JournalLine) GetNetAmountOk() (*float64, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *JournalLine) SetNetAmount(v float64)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmountField

`func (o *JournalLine) HasNetAmountField() bool`

HasNetAmountField returns a boolean if a field has been set.

### GetGrossAmount

`func (o *JournalLine) GetGrossAmount() float64`

GetGrossAmount returns the GrossAmount field if non-nil, zero value otherwise.

### GetGrossAmountOk

`func (o *JournalLine) GetGrossAmountOk() (*float64, bool)`

GetGrossAmountOk returns a tuple with the GrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmount

`func (o *JournalLine) SetGrossAmount(v float64)`

SetGrossAmount sets GrossAmount field to given value.

### HasGrossAmountField

`func (o *JournalLine) HasGrossAmountField() bool`

HasGrossAmountField returns a boolean if a field has been set.

### GetTaxAmount

`func (o *JournalLine) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *JournalLine) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *JournalLine) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmountField

`func (o *JournalLine) HasTaxAmountField() bool`

HasTaxAmountField returns a boolean if a field has been set.

### GetTaxType

`func (o *JournalLine) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *JournalLine) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *JournalLine) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxTypeField

`func (o *JournalLine) HasTaxTypeField() bool`

HasTaxTypeField returns a boolean if a field has been set.

### GetTaxName

`func (o *JournalLine) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *JournalLine) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *JournalLine) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.

### HasTaxNameField

`func (o *JournalLine) HasTaxNameField() bool`

HasTaxNameField returns a boolean if a field has been set.

### GetTrackingCategories

`func (o *JournalLine) GetTrackingCategories() []TrackingCategory`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *JournalLine) GetTrackingCategoriesOk() (*[]TrackingCategory, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *JournalLine) SetTrackingCategories(v []TrackingCategory)`

SetTrackingCategories sets TrackingCategories field to given value.

### HasTrackingCategoriesField

`func (o *JournalLine) HasTrackingCategoriesField() bool`

HasTrackingCategoriesField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


