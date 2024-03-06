# ManualJournalLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineAmount** | Pointer to **float64** | total for line. Debits are positive, credits are negative value | [optional] 
**AccountCode** | Pointer to **string** | See Accounts | [optional] 
**AccountID** | Pointer to **string** | See Accounts | [optional] 
**Description** | Pointer to **string** | Description for journal line | [optional] 
**TaxType** | Pointer to **string** | The tax type from TaxRates | [optional] 
**Tracking** | Pointer to [**[]TrackingCategory**](TrackingCategory.md) | Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 &lt;TrackingCategory&gt; elements. | [optional] 
**TaxAmount** | Pointer to **float64** | The calculated tax amount based on the TaxType and LineAmount | [optional] 
**IsBlank** | Pointer to **bool** | is the line blank | [optional] 

## Methods

### NewManualJournalLine

`func NewManualJournalLine() *ManualJournalLine`

NewManualJournalLine instantiates a new ManualJournalLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalLineWithDefaults

`func NewManualJournalLineWithDefaults() *ManualJournalLine`

NewManualJournalLineWithDefaults instantiates a new ManualJournalLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineAmount

`func (o *ManualJournalLine) GetLineAmount() float64`

GetLineAmount returns the LineAmount field if non-nil, zero value otherwise.

### GetLineAmountOk

`func (o *ManualJournalLine) GetLineAmountOk() (*float64, bool)`

GetLineAmountOk returns a tuple with the LineAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmount

`func (o *ManualJournalLine) SetLineAmount(v float64)`

SetLineAmount sets LineAmount field to given value.

### HasLineAmountField

`func (o *ManualJournalLine) HasLineAmountField() bool`

HasLineAmountField returns a boolean if a field has been set.

### GetAccountCode

`func (o *ManualJournalLine) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *ManualJournalLine) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *ManualJournalLine) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.

### HasAccountCodeField

`func (o *ManualJournalLine) HasAccountCodeField() bool`

HasAccountCodeField returns a boolean if a field has been set.

### GetAccountID

`func (o *ManualJournalLine) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *ManualJournalLine) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *ManualJournalLine) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountIDField

`func (o *ManualJournalLine) HasAccountIDField() bool`

HasAccountIDField returns a boolean if a field has been set.

### GetDescription

`func (o *ManualJournalLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManualJournalLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManualJournalLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescriptionField

`func (o *ManualJournalLine) HasDescriptionField() bool`

HasDescriptionField returns a boolean if a field has been set.

### GetTaxType

`func (o *ManualJournalLine) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ManualJournalLine) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ManualJournalLine) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxTypeField

`func (o *ManualJournalLine) HasTaxTypeField() bool`

HasTaxTypeField returns a boolean if a field has been set.

### GetTracking

`func (o *ManualJournalLine) GetTracking() []TrackingCategory`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *ManualJournalLine) GetTrackingOk() (*[]TrackingCategory, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *ManualJournalLine) SetTracking(v []TrackingCategory)`

SetTracking sets Tracking field to given value.

### HasTrackingField

`func (o *ManualJournalLine) HasTrackingField() bool`

HasTrackingField returns a boolean if a field has been set.

### GetTaxAmount

`func (o *ManualJournalLine) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *ManualJournalLine) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *ManualJournalLine) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmountField

`func (o *ManualJournalLine) HasTaxAmountField() bool`

HasTaxAmountField returns a boolean if a field has been set.

### GetIsBlank

`func (o *ManualJournalLine) GetIsBlank() bool`

GetIsBlank returns the IsBlank field if non-nil, zero value otherwise.

### GetIsBlankOk

`func (o *ManualJournalLine) GetIsBlankOk() (*bool, bool)`

GetIsBlankOk returns a tuple with the IsBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlank

`func (o *ManualJournalLine) SetIsBlank(v bool)`

SetIsBlank sets IsBlank field to given value.

### HasIsBlankField

`func (o *ManualJournalLine) HasIsBlankField() bool`

HasIsBlankField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


