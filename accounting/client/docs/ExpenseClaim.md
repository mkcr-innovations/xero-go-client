# ExpenseClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpenseClaimID** | Pointer to **string** | Xero generated unique identifier for an expense claim | [optional] 
**Status** | Pointer to **string** | Current status of an expense claim – see status types | [optional] 
**Payments** | Pointer to [**[]Payment**](Payment.md) | See Payments | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Receipts** | Pointer to [**[]Receipt**](Receipt.md) |  | [optional] 
**UpdatedDateUTC** | Pointer to **string** | Last modified date UTC format | [optional] [readonly] 
**Total** | Pointer to **float64** | The total of an expense claim being paid | [optional] [readonly] 
**AmountDue** | Pointer to **float64** | The amount due to be paid for an expense claim | [optional] [readonly] 
**AmountPaid** | Pointer to **float64** | The amount still to pay for an expense claim | [optional] [readonly] 
**PaymentDueDate** | Pointer to **string** | The date when the expense claim is due to be paid YYYY-MM-DD | [optional] [readonly] 
**ReportingDate** | Pointer to **string** | The date the expense claim will be reported in Xero YYYY-MM-DD | [optional] [readonly] 
**ReceiptID** | Pointer to **string** | The Xero identifier for the Receipt e.g. e59a2c7f-1306-4078-a0f3-73537afcbba9 | [optional] 

## Methods

### NewExpenseClaim

`func NewExpenseClaim() *ExpenseClaim`

NewExpenseClaim instantiates a new ExpenseClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseClaimWithDefaults

`func NewExpenseClaimWithDefaults() *ExpenseClaim`

NewExpenseClaimWithDefaults instantiates a new ExpenseClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpenseClaimID

`func (o *ExpenseClaim) GetExpenseClaimID() string`

GetExpenseClaimID returns the ExpenseClaimID field if non-nil, zero value otherwise.

### GetExpenseClaimIDOk

`func (o *ExpenseClaim) GetExpenseClaimIDOk() (*string, bool)`

GetExpenseClaimIDOk returns a tuple with the ExpenseClaimID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseClaimID

`func (o *ExpenseClaim) SetExpenseClaimID(v string)`

SetExpenseClaimID sets ExpenseClaimID field to given value.

### HasExpenseClaimIDField

`func (o *ExpenseClaim) HasExpenseClaimIDField() bool`

HasExpenseClaimIDField returns a boolean if a field has been set.

### GetStatus

`func (o *ExpenseClaim) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpenseClaim) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpenseClaim) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *ExpenseClaim) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetPayments

`func (o *ExpenseClaim) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *ExpenseClaim) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *ExpenseClaim) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPaymentsField

`func (o *ExpenseClaim) HasPaymentsField() bool`

HasPaymentsField returns a boolean if a field has been set.

### GetUser

`func (o *ExpenseClaim) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExpenseClaim) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExpenseClaim) SetUser(v User)`

SetUser sets User field to given value.

### HasUserField

`func (o *ExpenseClaim) HasUserField() bool`

HasUserField returns a boolean if a field has been set.

### GetReceipts

`func (o *ExpenseClaim) GetReceipts() []Receipt`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *ExpenseClaim) GetReceiptsOk() (*[]Receipt, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *ExpenseClaim) SetReceipts(v []Receipt)`

SetReceipts sets Receipts field to given value.

### HasReceiptsField

`func (o *ExpenseClaim) HasReceiptsField() bool`

HasReceiptsField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *ExpenseClaim) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *ExpenseClaim) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *ExpenseClaim) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *ExpenseClaim) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetTotal

`func (o *ExpenseClaim) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExpenseClaim) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExpenseClaim) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *ExpenseClaim) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetAmountDue

`func (o *ExpenseClaim) GetAmountDue() float64`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *ExpenseClaim) GetAmountDueOk() (*float64, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *ExpenseClaim) SetAmountDue(v float64)`

SetAmountDue sets AmountDue field to given value.

### HasAmountDueField

`func (o *ExpenseClaim) HasAmountDueField() bool`

HasAmountDueField returns a boolean if a field has been set.

### GetAmountPaid

`func (o *ExpenseClaim) GetAmountPaid() float64`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *ExpenseClaim) GetAmountPaidOk() (*float64, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *ExpenseClaim) SetAmountPaid(v float64)`

SetAmountPaid sets AmountPaid field to given value.

### HasAmountPaidField

`func (o *ExpenseClaim) HasAmountPaidField() bool`

HasAmountPaidField returns a boolean if a field has been set.

### GetPaymentDueDate

`func (o *ExpenseClaim) GetPaymentDueDate() string`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *ExpenseClaim) GetPaymentDueDateOk() (*string, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *ExpenseClaim) SetPaymentDueDate(v string)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDateField

`func (o *ExpenseClaim) HasPaymentDueDateField() bool`

HasPaymentDueDateField returns a boolean if a field has been set.

### GetReportingDate

`func (o *ExpenseClaim) GetReportingDate() string`

GetReportingDate returns the ReportingDate field if non-nil, zero value otherwise.

### GetReportingDateOk

`func (o *ExpenseClaim) GetReportingDateOk() (*string, bool)`

GetReportingDateOk returns a tuple with the ReportingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingDate

`func (o *ExpenseClaim) SetReportingDate(v string)`

SetReportingDate sets ReportingDate field to given value.

### HasReportingDateField

`func (o *ExpenseClaim) HasReportingDateField() bool`

HasReportingDateField returns a boolean if a field has been set.

### GetReceiptID

`func (o *ExpenseClaim) GetReceiptID() string`

GetReceiptID returns the ReceiptID field if non-nil, zero value otherwise.

### GetReceiptIDOk

`func (o *ExpenseClaim) GetReceiptIDOk() (*string, bool)`

GetReceiptIDOk returns a tuple with the ReceiptID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptID

`func (o *ExpenseClaim) SetReceiptID(v string)`

SetReceiptID sets ReceiptID field to given value.

### HasReceiptIDField

`func (o *ExpenseClaim) HasReceiptIDField() bool`

HasReceiptIDField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


