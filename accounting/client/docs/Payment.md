# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | Pointer to [**Invoice**](Invoice.md) |  | [optional] 
**CreditNote** | Pointer to [**CreditNote**](CreditNote.md) |  | [optional] 
**Prepayment** | Pointer to [**Prepayment**](Prepayment.md) |  | [optional] 
**Overpayment** | Pointer to [**Overpayment**](Overpayment.md) |  | [optional] 
**InvoiceNumber** | Pointer to **string** | Number of invoice or credit note you are applying payment to e.g.INV-4003 | [optional] 
**CreditNoteNumber** | Pointer to **string** | Number of invoice or credit note you are applying payment to e.g. INV-4003 | [optional] 
**BatchPayment** | Pointer to [**BatchPayment**](BatchPayment.md) |  | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**Code** | Pointer to **string** | Code of account you are using to make the payment e.g. 001 (note- not all accounts have a code value) | [optional] 
**Date** | Pointer to **string** | Date the payment is being made (YYYY-MM-DD) e.g. 2009-09-06 | [optional] 
**CurrencyRate** | Pointer to **float64** | Exchange rate when payment is received. Only used for non base currency invoices and credit notes e.g. 0.7500 | [optional] 
**Amount** | Pointer to **float64** | The amount of the payment. Must be less than or equal to the outstanding amount owing on the invoice e.g. 200.00 | [optional] 
**BankAmount** | Pointer to **float64** | The amount of the payment in the currency of the bank account. | [optional] 
**Reference** | Pointer to **string** | An optional description for the payment e.g. Direct Debit | [optional] 
**IsReconciled** | Pointer to **bool** | An optional parameter for the payment. A boolean indicating whether you would like the payment to be created as reconciled when using PUT, or whether a payment has been reconciled when using GET | [optional] 
**Status** | Pointer to **string** | The status of the payment. | [optional] 
**PaymentType** | Pointer to **string** | See Payment Types. | [optional] [readonly] 
**UpdatedDateUTC** | Pointer to **string** | UTC timestamp of last update to the payment | [optional] [readonly] 
**PaymentID** | Pointer to **string** | The Xero identifier for an Payment e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9 | [optional] 
**BatchPaymentID** | Pointer to **string** | Present if the payment was created as part of a batch. | [optional] 
**BankAccountNumber** | Pointer to **string** | The suppliers bank account number the payment is being made to | [optional] 
**Particulars** | Pointer to **string** | The suppliers bank account number the payment is being made to | [optional] 
**Details** | Pointer to **string** | The information to appear on the supplier&#39;s bank account | [optional] 
**HasAccount** | Pointer to **bool** | A boolean to indicate if a contact has an validation errors | [optional] [default to false]
**HasValidationErrors** | Pointer to **bool** | A boolean to indicate if a contact has an validation errors | [optional] [default to false]
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 
**Warnings** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of warning messages from the API | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *Payment) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *Payment) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *Payment) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.

### HasInvoiceField

`func (o *Payment) HasInvoiceField() bool`

HasInvoiceField returns a boolean if a field has been set.

### GetCreditNote

`func (o *Payment) GetCreditNote() CreditNote`

GetCreditNote returns the CreditNote field if non-nil, zero value otherwise.

### GetCreditNoteOk

`func (o *Payment) GetCreditNoteOk() (*CreditNote, bool)`

GetCreditNoteOk returns a tuple with the CreditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNote

`func (o *Payment) SetCreditNote(v CreditNote)`

SetCreditNote sets CreditNote field to given value.

### HasCreditNoteField

`func (o *Payment) HasCreditNoteField() bool`

HasCreditNoteField returns a boolean if a field has been set.

### GetPrepayment

`func (o *Payment) GetPrepayment() Prepayment`

GetPrepayment returns the Prepayment field if non-nil, zero value otherwise.

### GetPrepaymentOk

`func (o *Payment) GetPrepaymentOk() (*Prepayment, bool)`

GetPrepaymentOk returns a tuple with the Prepayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepayment

`func (o *Payment) SetPrepayment(v Prepayment)`

SetPrepayment sets Prepayment field to given value.

### HasPrepaymentField

`func (o *Payment) HasPrepaymentField() bool`

HasPrepaymentField returns a boolean if a field has been set.

### GetOverpayment

`func (o *Payment) GetOverpayment() Overpayment`

GetOverpayment returns the Overpayment field if non-nil, zero value otherwise.

### GetOverpaymentOk

`func (o *Payment) GetOverpaymentOk() (*Overpayment, bool)`

GetOverpaymentOk returns a tuple with the Overpayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverpayment

`func (o *Payment) SetOverpayment(v Overpayment)`

SetOverpayment sets Overpayment field to given value.

### HasOverpaymentField

`func (o *Payment) HasOverpaymentField() bool`

HasOverpaymentField returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *Payment) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Payment) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Payment) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumberField

`func (o *Payment) HasInvoiceNumberField() bool`

HasInvoiceNumberField returns a boolean if a field has been set.

### GetCreditNoteNumber

`func (o *Payment) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *Payment) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *Payment) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.

### HasCreditNoteNumberField

`func (o *Payment) HasCreditNoteNumberField() bool`

HasCreditNoteNumberField returns a boolean if a field has been set.

### GetBatchPayment

`func (o *Payment) GetBatchPayment() BatchPayment`

GetBatchPayment returns the BatchPayment field if non-nil, zero value otherwise.

### GetBatchPaymentOk

`func (o *Payment) GetBatchPaymentOk() (*BatchPayment, bool)`

GetBatchPaymentOk returns a tuple with the BatchPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchPayment

`func (o *Payment) SetBatchPayment(v BatchPayment)`

SetBatchPayment sets BatchPayment field to given value.

### HasBatchPaymentField

`func (o *Payment) HasBatchPaymentField() bool`

HasBatchPaymentField returns a boolean if a field has been set.

### GetAccount

`func (o *Payment) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Payment) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Payment) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccountField

`func (o *Payment) HasAccountField() bool`

HasAccountField returns a boolean if a field has been set.

### GetCode

`func (o *Payment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Payment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Payment) SetCode(v string)`

SetCode sets Code field to given value.

### HasCodeField

`func (o *Payment) HasCodeField() bool`

HasCodeField returns a boolean if a field has been set.

### GetDate

`func (o *Payment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Payment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Payment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *Payment) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *Payment) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *Payment) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *Payment) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *Payment) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetAmount

`func (o *Payment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmountField

`func (o *Payment) HasAmountField() bool`

HasAmountField returns a boolean if a field has been set.

### GetBankAmount

`func (o *Payment) GetBankAmount() float64`

GetBankAmount returns the BankAmount field if non-nil, zero value otherwise.

### GetBankAmountOk

`func (o *Payment) GetBankAmountOk() (*float64, bool)`

GetBankAmountOk returns a tuple with the BankAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAmount

`func (o *Payment) SetBankAmount(v float64)`

SetBankAmount sets BankAmount field to given value.

### HasBankAmountField

`func (o *Payment) HasBankAmountField() bool`

HasBankAmountField returns a boolean if a field has been set.

### GetReference

`func (o *Payment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Payment) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Payment) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *Payment) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetIsReconciled

`func (o *Payment) GetIsReconciled() bool`

GetIsReconciled returns the IsReconciled field if non-nil, zero value otherwise.

### GetIsReconciledOk

`func (o *Payment) GetIsReconciledOk() (*bool, bool)`

GetIsReconciledOk returns a tuple with the IsReconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReconciled

`func (o *Payment) SetIsReconciled(v bool)`

SetIsReconciled sets IsReconciled field to given value.

### HasIsReconciledField

`func (o *Payment) HasIsReconciledField() bool`

HasIsReconciledField returns a boolean if a field has been set.

### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *Payment) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetPaymentType

`func (o *Payment) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *Payment) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *Payment) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentTypeField

`func (o *Payment) HasPaymentTypeField() bool`

HasPaymentTypeField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Payment) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Payment) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Payment) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Payment) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetPaymentID

`func (o *Payment) GetPaymentID() string`

GetPaymentID returns the PaymentID field if non-nil, zero value otherwise.

### GetPaymentIDOk

`func (o *Payment) GetPaymentIDOk() (*string, bool)`

GetPaymentIDOk returns a tuple with the PaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentID

`func (o *Payment) SetPaymentID(v string)`

SetPaymentID sets PaymentID field to given value.

### HasPaymentIDField

`func (o *Payment) HasPaymentIDField() bool`

HasPaymentIDField returns a boolean if a field has been set.

### GetBatchPaymentID

`func (o *Payment) GetBatchPaymentID() string`

GetBatchPaymentID returns the BatchPaymentID field if non-nil, zero value otherwise.

### GetBatchPaymentIDOk

`func (o *Payment) GetBatchPaymentIDOk() (*string, bool)`

GetBatchPaymentIDOk returns a tuple with the BatchPaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchPaymentID

`func (o *Payment) SetBatchPaymentID(v string)`

SetBatchPaymentID sets BatchPaymentID field to given value.

### HasBatchPaymentIDField

`func (o *Payment) HasBatchPaymentIDField() bool`

HasBatchPaymentIDField returns a boolean if a field has been set.

### GetBankAccountNumber

`func (o *Payment) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *Payment) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *Payment) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumberField

`func (o *Payment) HasBankAccountNumberField() bool`

HasBankAccountNumberField returns a boolean if a field has been set.

### GetParticulars

`func (o *Payment) GetParticulars() string`

GetParticulars returns the Particulars field if non-nil, zero value otherwise.

### GetParticularsOk

`func (o *Payment) GetParticularsOk() (*string, bool)`

GetParticularsOk returns a tuple with the Particulars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticulars

`func (o *Payment) SetParticulars(v string)`

SetParticulars sets Particulars field to given value.

### HasParticularsField

`func (o *Payment) HasParticularsField() bool`

HasParticularsField returns a boolean if a field has been set.

### GetDetails

`func (o *Payment) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Payment) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Payment) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetailsField

`func (o *Payment) HasDetailsField() bool`

HasDetailsField returns a boolean if a field has been set.

### GetHasAccount

`func (o *Payment) GetHasAccount() bool`

GetHasAccount returns the HasAccount field if non-nil, zero value otherwise.

### GetHasAccountOk

`func (o *Payment) GetHasAccountOk() (*bool, bool)`

GetHasAccountOk returns a tuple with the HasAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccount

`func (o *Payment) SetHasAccount(v bool)`

SetHasAccount sets HasAccount field to given value.

### HasHasAccountField

`func (o *Payment) HasHasAccountField() bool`

HasHasAccountField returns a boolean if a field has been set.

### GetHasValidationErrors

`func (o *Payment) GetHasValidationErrors() bool`

GetHasValidationErrors returns the HasValidationErrors field if non-nil, zero value otherwise.

### GetHasValidationErrorsOk

`func (o *Payment) GetHasValidationErrorsOk() (*bool, bool)`

GetHasValidationErrorsOk returns a tuple with the HasValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValidationErrors

`func (o *Payment) SetHasValidationErrors(v bool)`

SetHasValidationErrors sets HasValidationErrors field to given value.

### HasHasValidationErrorsField

`func (o *Payment) HasHasValidationErrorsField() bool`

HasHasValidationErrorsField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *Payment) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *Payment) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *Payment) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *Payment) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Payment) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Payment) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Payment) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Payment) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.

### GetWarnings

`func (o *Payment) GetWarnings() []ValidationError`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Payment) GetWarningsOk() (*[]ValidationError, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Payment) SetWarnings(v []ValidationError)`

SetWarnings sets Warnings field to given value.

### HasWarningsField

`func (o *Payment) HasWarningsField() bool`

HasWarningsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


