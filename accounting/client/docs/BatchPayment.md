# BatchPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**Reference** | Pointer to **string** | (NZ Only) Optional references for the batch payment transaction. It will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement you import into Xero. | [optional] 
**Particulars** | Pointer to **string** | (NZ Only) Optional references for the batch payment transaction. It will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement you import into Xero. | [optional] 
**Code** | Pointer to **string** | (NZ Only) Optional references for the batch payment transaction. It will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement you import into Xero. | [optional] 
**Details** | Pointer to **string** | (Non-NZ Only) These details are sent to the org’s bank as a reference for the batch payment transaction. They will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement imported into Xero. Maximum field length &#x3D; 18 | [optional] 
**Narrative** | Pointer to **string** | (UK Only) Only shows on the statement line in Xero. Max length &#x3D;18 | [optional] 
**BatchPaymentID** | Pointer to **string** | The Xero generated unique identifier for the bank transaction (read-only) | [optional] [readonly] 
**DateString** | Pointer to **string** | Date the payment is being made (YYYY-MM-DD) e.g. 2009-09-06 | [optional] 
**Date** | Pointer to **string** | Date the payment is being made (YYYY-MM-DD) e.g. 2009-09-06 | [optional] 
**Amount** | Pointer to **float64** | The amount of the payment. Must be less than or equal to the outstanding amount owing on the invoice e.g. 200.00 | [optional] 
**Payments** | Pointer to [**[]Payment**](Payment.md) | An array of payments | [optional] 
**Type** | Pointer to **string** | PAYBATCH for bill payments or RECBATCH for sales invoice payments (read-only) | [optional] [readonly] 
**Status** | Pointer to **string** | AUTHORISED or DELETED (read-only). New batch payments will have a status of AUTHORISED. It is not possible to delete batch payments via the API. | [optional] [readonly] 
**TotalAmount** | Pointer to **float64** | The total of the payments that make up the batch (read-only) | [optional] [readonly] 
**UpdatedDateUTC** | Pointer to **string** | UTC timestamp of last update to the payment | [optional] [readonly] 
**IsReconciled** | Pointer to **bool** | Booelan that tells you if the batch payment has been reconciled (read-only) | [optional] [readonly] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewBatchPayment

`func NewBatchPayment() *BatchPayment`

NewBatchPayment instantiates a new BatchPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentWithDefaults

`func NewBatchPaymentWithDefaults() *BatchPayment`

NewBatchPaymentWithDefaults instantiates a new BatchPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BatchPayment) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BatchPayment) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BatchPayment) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccountField

`func (o *BatchPayment) HasAccountField() bool`

HasAccountField returns a boolean if a field has been set.

### GetReference

`func (o *BatchPayment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BatchPayment) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BatchPayment) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *BatchPayment) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetParticulars

`func (o *BatchPayment) GetParticulars() string`

GetParticulars returns the Particulars field if non-nil, zero value otherwise.

### GetParticularsOk

`func (o *BatchPayment) GetParticularsOk() (*string, bool)`

GetParticularsOk returns a tuple with the Particulars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticulars

`func (o *BatchPayment) SetParticulars(v string)`

SetParticulars sets Particulars field to given value.

### HasParticularsField

`func (o *BatchPayment) HasParticularsField() bool`

HasParticularsField returns a boolean if a field has been set.

### GetCode

`func (o *BatchPayment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BatchPayment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BatchPayment) SetCode(v string)`

SetCode sets Code field to given value.

### HasCodeField

`func (o *BatchPayment) HasCodeField() bool`

HasCodeField returns a boolean if a field has been set.

### GetDetails

`func (o *BatchPayment) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BatchPayment) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BatchPayment) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetailsField

`func (o *BatchPayment) HasDetailsField() bool`

HasDetailsField returns a boolean if a field has been set.

### GetNarrative

`func (o *BatchPayment) GetNarrative() string`

GetNarrative returns the Narrative field if non-nil, zero value otherwise.

### GetNarrativeOk

`func (o *BatchPayment) GetNarrativeOk() (*string, bool)`

GetNarrativeOk returns a tuple with the Narrative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrative

`func (o *BatchPayment) SetNarrative(v string)`

SetNarrative sets Narrative field to given value.

### HasNarrativeField

`func (o *BatchPayment) HasNarrativeField() bool`

HasNarrativeField returns a boolean if a field has been set.

### GetBatchPaymentID

`func (o *BatchPayment) GetBatchPaymentID() string`

GetBatchPaymentID returns the BatchPaymentID field if non-nil, zero value otherwise.

### GetBatchPaymentIDOk

`func (o *BatchPayment) GetBatchPaymentIDOk() (*string, bool)`

GetBatchPaymentIDOk returns a tuple with the BatchPaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchPaymentID

`func (o *BatchPayment) SetBatchPaymentID(v string)`

SetBatchPaymentID sets BatchPaymentID field to given value.

### HasBatchPaymentIDField

`func (o *BatchPayment) HasBatchPaymentIDField() bool`

HasBatchPaymentIDField returns a boolean if a field has been set.

### GetDateString

`func (o *BatchPayment) GetDateString() string`

GetDateString returns the DateString field if non-nil, zero value otherwise.

### GetDateStringOk

`func (o *BatchPayment) GetDateStringOk() (*string, bool)`

GetDateStringOk returns a tuple with the DateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateString

`func (o *BatchPayment) SetDateString(v string)`

SetDateString sets DateString field to given value.

### HasDateStringField

`func (o *BatchPayment) HasDateStringField() bool`

HasDateStringField returns a boolean if a field has been set.

### GetDate

`func (o *BatchPayment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BatchPayment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BatchPayment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *BatchPayment) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetAmount

`func (o *BatchPayment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BatchPayment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BatchPayment) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmountField

`func (o *BatchPayment) HasAmountField() bool`

HasAmountField returns a boolean if a field has been set.

### GetPayments

`func (o *BatchPayment) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *BatchPayment) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *BatchPayment) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPaymentsField

`func (o *BatchPayment) HasPaymentsField() bool`

HasPaymentsField returns a boolean if a field has been set.

### GetType

`func (o *BatchPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BatchPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BatchPayment) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *BatchPayment) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetStatus

`func (o *BatchPayment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchPayment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchPayment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *BatchPayment) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetTotalAmount

`func (o *BatchPayment) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *BatchPayment) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *BatchPayment) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmountField

`func (o *BatchPayment) HasTotalAmountField() bool`

HasTotalAmountField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *BatchPayment) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *BatchPayment) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *BatchPayment) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *BatchPayment) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetIsReconciled

`func (o *BatchPayment) GetIsReconciled() bool`

GetIsReconciled returns the IsReconciled field if non-nil, zero value otherwise.

### GetIsReconciledOk

`func (o *BatchPayment) GetIsReconciledOk() (*bool, bool)`

GetIsReconciledOk returns a tuple with the IsReconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReconciled

`func (o *BatchPayment) SetIsReconciled(v bool)`

SetIsReconciled sets IsReconciled field to given value.

### HasIsReconciledField

`func (o *BatchPayment) HasIsReconciledField() bool`

HasIsReconciledField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *BatchPayment) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *BatchPayment) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *BatchPayment) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *BatchPayment) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


