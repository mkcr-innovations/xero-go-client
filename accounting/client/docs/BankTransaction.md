# BankTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | See Bank Transaction Types | 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**LineItems** | [**[]LineItem**](LineItem.md) | See LineItems | 
**BankAccount** | [**Account**](Account.md) |  | 
**IsReconciled** | Pointer to **bool** | Boolean to show if transaction is reconciled | [optional] 
**Date** | Pointer to **string** | Date of transaction – YYYY-MM-DD | [optional] 
**Reference** | Pointer to **string** | Reference for the transaction. Only supported for SPEND and RECEIVE transactions. | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**CurrencyRate** | Pointer to **float64** | Exchange rate to base currency when money is spent or received. e.g.0.7500 Only used for bank transactions in non base currency. If this isn’t specified for non base currency accounts then either the user-defined rate (preference) or the XE.com day rate will be used. Setting currency is only supported on overpayments. | [optional] 
**Url** | Pointer to **string** | URL link to a source document – shown as “Go to App Name” | [optional] 
**Status** | Pointer to **string** | See Bank Transaction Status Codes | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**SubTotal** | Pointer to **float64** | Total of bank transaction excluding taxes | [optional] 
**TotalTax** | Pointer to **float64** | Total tax on bank transaction | [optional] 
**Total** | Pointer to **float64** | Total of bank transaction tax inclusive | [optional] 
**BankTransactionID** | Pointer to **string** | Xero generated unique identifier for bank transaction | [optional] 
**PrepaymentID** | Pointer to **string** | Xero generated unique identifier for a Prepayment. This will be returned on BankTransactions with a Type of SPEND-PREPAYMENT or RECEIVE-PREPAYMENT | [optional] [readonly] 
**OverpaymentID** | Pointer to **string** | Xero generated unique identifier for an Overpayment. This will be returned on BankTransactions with a Type of SPEND-OVERPAYMENT or RECEIVE-OVERPAYMENT | [optional] [readonly] 
**UpdatedDateUTC** | Pointer to **string** | Last modified date UTC format | [optional] [readonly] 
**HasAttachments** | Pointer to **bool** | Boolean to indicate if a bank transaction has an attachment | [optional] [readonly] [default to false]
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewBankTransaction

`func NewBankTransaction(type_ string, lineItems []LineItem, bankAccount Account, ) *BankTransaction`

NewBankTransaction instantiates a new BankTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransactionWithDefaults

`func NewBankTransactionWithDefaults() *BankTransaction`

NewBankTransactionWithDefaults instantiates a new BankTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BankTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetContact

`func (o *BankTransaction) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *BankTransaction) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *BankTransaction) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *BankTransaction) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetLineItems

`func (o *BankTransaction) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BankTransaction) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BankTransaction) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.


### GetBankAccount

`func (o *BankTransaction) GetBankAccount() Account`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *BankTransaction) GetBankAccountOk() (*Account, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *BankTransaction) SetBankAccount(v Account)`

SetBankAccount sets BankAccount field to given value.


### GetIsReconciled

`func (o *BankTransaction) GetIsReconciled() bool`

GetIsReconciled returns the IsReconciled field if non-nil, zero value otherwise.

### GetIsReconciledOk

`func (o *BankTransaction) GetIsReconciledOk() (*bool, bool)`

GetIsReconciledOk returns a tuple with the IsReconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReconciled

`func (o *BankTransaction) SetIsReconciled(v bool)`

SetIsReconciled sets IsReconciled field to given value.

### HasIsReconciledField

`func (o *BankTransaction) HasIsReconciledField() bool`

HasIsReconciledField returns a boolean if a field has been set.

### GetDate

`func (o *BankTransaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BankTransaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BankTransaction) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *BankTransaction) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetReference

`func (o *BankTransaction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BankTransaction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BankTransaction) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *BankTransaction) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BankTransaction) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BankTransaction) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BankTransaction) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *BankTransaction) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *BankTransaction) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *BankTransaction) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *BankTransaction) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *BankTransaction) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetUrl

`func (o *BankTransaction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BankTransaction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BankTransaction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrlField

`func (o *BankTransaction) HasUrlField() bool`

HasUrlField returns a boolean if a field has been set.

### GetStatus

`func (o *BankTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *BankTransaction) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *BankTransaction) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *BankTransaction) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *BankTransaction) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *BankTransaction) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetSubTotal

`func (o *BankTransaction) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *BankTransaction) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *BankTransaction) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *BankTransaction) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *BankTransaction) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *BankTransaction) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *BankTransaction) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *BankTransaction) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *BankTransaction) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BankTransaction) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BankTransaction) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *BankTransaction) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetBankTransactionID

`func (o *BankTransaction) GetBankTransactionID() string`

GetBankTransactionID returns the BankTransactionID field if non-nil, zero value otherwise.

### GetBankTransactionIDOk

`func (o *BankTransaction) GetBankTransactionIDOk() (*string, bool)`

GetBankTransactionIDOk returns a tuple with the BankTransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransactionID

`func (o *BankTransaction) SetBankTransactionID(v string)`

SetBankTransactionID sets BankTransactionID field to given value.

### HasBankTransactionIDField

`func (o *BankTransaction) HasBankTransactionIDField() bool`

HasBankTransactionIDField returns a boolean if a field has been set.

### GetPrepaymentID

`func (o *BankTransaction) GetPrepaymentID() string`

GetPrepaymentID returns the PrepaymentID field if non-nil, zero value otherwise.

### GetPrepaymentIDOk

`func (o *BankTransaction) GetPrepaymentIDOk() (*string, bool)`

GetPrepaymentIDOk returns a tuple with the PrepaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaymentID

`func (o *BankTransaction) SetPrepaymentID(v string)`

SetPrepaymentID sets PrepaymentID field to given value.

### HasPrepaymentIDField

`func (o *BankTransaction) HasPrepaymentIDField() bool`

HasPrepaymentIDField returns a boolean if a field has been set.

### GetOverpaymentID

`func (o *BankTransaction) GetOverpaymentID() string`

GetOverpaymentID returns the OverpaymentID field if non-nil, zero value otherwise.

### GetOverpaymentIDOk

`func (o *BankTransaction) GetOverpaymentIDOk() (*string, bool)`

GetOverpaymentIDOk returns a tuple with the OverpaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverpaymentID

`func (o *BankTransaction) SetOverpaymentID(v string)`

SetOverpaymentID sets OverpaymentID field to given value.

### HasOverpaymentIDField

`func (o *BankTransaction) HasOverpaymentIDField() bool`

HasOverpaymentIDField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *BankTransaction) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *BankTransaction) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *BankTransaction) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *BankTransaction) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *BankTransaction) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *BankTransaction) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *BankTransaction) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *BankTransaction) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *BankTransaction) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *BankTransaction) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *BankTransaction) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *BankTransaction) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *BankTransaction) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *BankTransaction) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *BankTransaction) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *BankTransaction) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


