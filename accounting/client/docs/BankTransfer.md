# BankTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromBankAccount** | [**Account**](Account.md) |  | 
**ToBankAccount** | [**Account**](Account.md) |  | 
**Amount** | **float64** | amount of the transaction | 
**Date** | Pointer to **string** | The date of the Transfer YYYY-MM-DD | [optional] 
**BankTransferID** | Pointer to **string** | The identifier of the Bank Transfer | [optional] [readonly] 
**CurrencyRate** | Pointer to **float64** | The currency rate | [optional] [readonly] 
**FromBankTransactionID** | Pointer to **string** | The Bank Transaction ID for the source account | [optional] [readonly] 
**ToBankTransactionID** | Pointer to **string** | The Bank Transaction ID for the destination account | [optional] [readonly] 
**FromIsReconciled** | Pointer to **bool** | The Bank Transaction boolean to show if it is reconciled for the source account | [optional] [default to false]
**ToIsReconciled** | Pointer to **bool** | The Bank Transaction boolean to show if it is reconciled for the destination account | [optional] [default to false]
**Reference** | Pointer to **string** | Reference for the transactions. | [optional] 
**HasAttachments** | Pointer to **bool** | Boolean to indicate if a Bank Transfer has an attachment | [optional] [readonly] [default to false]
**CreatedDateUTC** | Pointer to **string** | UTC timestamp of creation date of bank transfer | [optional] [readonly] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewBankTransfer

`func NewBankTransfer(fromBankAccount Account, toBankAccount Account, amount float64, ) *BankTransfer`

NewBankTransfer instantiates a new BankTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferWithDefaults

`func NewBankTransferWithDefaults() *BankTransfer`

NewBankTransferWithDefaults instantiates a new BankTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromBankAccount

`func (o *BankTransfer) GetFromBankAccount() Account`

GetFromBankAccount returns the FromBankAccount field if non-nil, zero value otherwise.

### GetFromBankAccountOk

`func (o *BankTransfer) GetFromBankAccountOk() (*Account, bool)`

GetFromBankAccountOk returns a tuple with the FromBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBankAccount

`func (o *BankTransfer) SetFromBankAccount(v Account)`

SetFromBankAccount sets FromBankAccount field to given value.


### GetToBankAccount

`func (o *BankTransfer) GetToBankAccount() Account`

GetToBankAccount returns the ToBankAccount field if non-nil, zero value otherwise.

### GetToBankAccountOk

`func (o *BankTransfer) GetToBankAccountOk() (*Account, bool)`

GetToBankAccountOk returns a tuple with the ToBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBankAccount

`func (o *BankTransfer) SetToBankAccount(v Account)`

SetToBankAccount sets ToBankAccount field to given value.


### GetAmount

`func (o *BankTransfer) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransfer) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransfer) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *BankTransfer) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BankTransfer) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BankTransfer) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *BankTransfer) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetBankTransferID

`func (o *BankTransfer) GetBankTransferID() string`

GetBankTransferID returns the BankTransferID field if non-nil, zero value otherwise.

### GetBankTransferIDOk

`func (o *BankTransfer) GetBankTransferIDOk() (*string, bool)`

GetBankTransferIDOk returns a tuple with the BankTransferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferID

`func (o *BankTransfer) SetBankTransferID(v string)`

SetBankTransferID sets BankTransferID field to given value.

### HasBankTransferIDField

`func (o *BankTransfer) HasBankTransferIDField() bool`

HasBankTransferIDField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *BankTransfer) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *BankTransfer) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *BankTransfer) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *BankTransfer) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetFromBankTransactionID

`func (o *BankTransfer) GetFromBankTransactionID() string`

GetFromBankTransactionID returns the FromBankTransactionID field if non-nil, zero value otherwise.

### GetFromBankTransactionIDOk

`func (o *BankTransfer) GetFromBankTransactionIDOk() (*string, bool)`

GetFromBankTransactionIDOk returns a tuple with the FromBankTransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBankTransactionID

`func (o *BankTransfer) SetFromBankTransactionID(v string)`

SetFromBankTransactionID sets FromBankTransactionID field to given value.

### HasFromBankTransactionIDField

`func (o *BankTransfer) HasFromBankTransactionIDField() bool`

HasFromBankTransactionIDField returns a boolean if a field has been set.

### GetToBankTransactionID

`func (o *BankTransfer) GetToBankTransactionID() string`

GetToBankTransactionID returns the ToBankTransactionID field if non-nil, zero value otherwise.

### GetToBankTransactionIDOk

`func (o *BankTransfer) GetToBankTransactionIDOk() (*string, bool)`

GetToBankTransactionIDOk returns a tuple with the ToBankTransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBankTransactionID

`func (o *BankTransfer) SetToBankTransactionID(v string)`

SetToBankTransactionID sets ToBankTransactionID field to given value.

### HasToBankTransactionIDField

`func (o *BankTransfer) HasToBankTransactionIDField() bool`

HasToBankTransactionIDField returns a boolean if a field has been set.

### GetFromIsReconciled

`func (o *BankTransfer) GetFromIsReconciled() bool`

GetFromIsReconciled returns the FromIsReconciled field if non-nil, zero value otherwise.

### GetFromIsReconciledOk

`func (o *BankTransfer) GetFromIsReconciledOk() (*bool, bool)`

GetFromIsReconciledOk returns a tuple with the FromIsReconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIsReconciled

`func (o *BankTransfer) SetFromIsReconciled(v bool)`

SetFromIsReconciled sets FromIsReconciled field to given value.

### HasFromIsReconciledField

`func (o *BankTransfer) HasFromIsReconciledField() bool`

HasFromIsReconciledField returns a boolean if a field has been set.

### GetToIsReconciled

`func (o *BankTransfer) GetToIsReconciled() bool`

GetToIsReconciled returns the ToIsReconciled field if non-nil, zero value otherwise.

### GetToIsReconciledOk

`func (o *BankTransfer) GetToIsReconciledOk() (*bool, bool)`

GetToIsReconciledOk returns a tuple with the ToIsReconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToIsReconciled

`func (o *BankTransfer) SetToIsReconciled(v bool)`

SetToIsReconciled sets ToIsReconciled field to given value.

### HasToIsReconciledField

`func (o *BankTransfer) HasToIsReconciledField() bool`

HasToIsReconciledField returns a boolean if a field has been set.

### GetReference

`func (o *BankTransfer) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BankTransfer) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BankTransfer) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *BankTransfer) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *BankTransfer) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *BankTransfer) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *BankTransfer) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *BankTransfer) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetCreatedDateUTC

`func (o *BankTransfer) GetCreatedDateUTC() string`

GetCreatedDateUTC returns the CreatedDateUTC field if non-nil, zero value otherwise.

### GetCreatedDateUTCOk

`func (o *BankTransfer) GetCreatedDateUTCOk() (*string, bool)`

GetCreatedDateUTCOk returns a tuple with the CreatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateUTC

`func (o *BankTransfer) SetCreatedDateUTC(v string)`

SetCreatedDateUTC sets CreatedDateUTC field to given value.

### HasCreatedDateUTCField

`func (o *BankTransfer) HasCreatedDateUTCField() bool`

HasCreatedDateUTCField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *BankTransfer) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *BankTransfer) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *BankTransfer) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *BankTransfer) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


