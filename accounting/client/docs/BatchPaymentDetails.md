# BatchPaymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | Pointer to **string** | Bank account number for use with Batch Payments | [optional] 
**BankAccountName** | Pointer to **string** | Name of bank for use with Batch Payments | [optional] 
**Details** | Pointer to **string** | (Non-NZ Only) These details are sent to the org’s bank as a reference for the batch payment transaction. They will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement imported into Xero. Maximum field length &#x3D; 18 | [optional] 
**Code** | Pointer to **string** | (NZ Only) Optional references for the batch payment transaction. It will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement you import into Xero. | [optional] 
**Reference** | Pointer to **string** | (NZ Only) Optional references for the batch payment transaction. It will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement you import into Xero. | [optional] 

## Methods

### NewBatchPaymentDetails

`func NewBatchPaymentDetails() *BatchPaymentDetails`

NewBatchPaymentDetails instantiates a new BatchPaymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentDetailsWithDefaults

`func NewBatchPaymentDetailsWithDefaults() *BatchPaymentDetails`

NewBatchPaymentDetailsWithDefaults instantiates a new BatchPaymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountNumber

`func (o *BatchPaymentDetails) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *BatchPaymentDetails) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *BatchPaymentDetails) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumberField

`func (o *BatchPaymentDetails) HasBankAccountNumberField() bool`

HasBankAccountNumberField returns a boolean if a field has been set.

### GetBankAccountName

`func (o *BatchPaymentDetails) GetBankAccountName() string`

GetBankAccountName returns the BankAccountName field if non-nil, zero value otherwise.

### GetBankAccountNameOk

`func (o *BatchPaymentDetails) GetBankAccountNameOk() (*string, bool)`

GetBankAccountNameOk returns a tuple with the BankAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountName

`func (o *BatchPaymentDetails) SetBankAccountName(v string)`

SetBankAccountName sets BankAccountName field to given value.

### HasBankAccountNameField

`func (o *BatchPaymentDetails) HasBankAccountNameField() bool`

HasBankAccountNameField returns a boolean if a field has been set.

### GetDetails

`func (o *BatchPaymentDetails) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BatchPaymentDetails) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BatchPaymentDetails) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetailsField

`func (o *BatchPaymentDetails) HasDetailsField() bool`

HasDetailsField returns a boolean if a field has been set.

### GetCode

`func (o *BatchPaymentDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BatchPaymentDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BatchPaymentDetails) SetCode(v string)`

SetCode sets Code field to given value.

### HasCodeField

`func (o *BatchPaymentDetails) HasCodeField() bool`

HasCodeField returns a boolean if a field has been set.

### GetReference

`func (o *BatchPaymentDetails) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BatchPaymentDetails) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BatchPaymentDetails) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *BatchPaymentDetails) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


