# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Customer defined alpha numeric account code e.g 200 or SALES (max length &#x3D; 10) | [optional] 
**Name** | Pointer to **string** | Name of account (max length &#x3D; 150) | [optional] 
**AccountID** | Pointer to **string** | The Xero identifier for an account – specified as a string following  the endpoint name   e.g. /297c2dc5-cc47-4afd-8ec8-74990b8761e9 | [optional] 
**Type** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**BankAccountNumber** | Pointer to **string** | For bank accounts only (Account Type BANK) | [optional] 
**Status** | Pointer to **string** | Accounts with a status of ACTIVE can be updated to ARCHIVED. See Account Status Codes | [optional] 
**Description** | Pointer to **string** | Description of the Account. Valid for all types of accounts except bank accounts (max length &#x3D; 4000) | [optional] 
**BankAccountType** | Pointer to **string** | For bank accounts only. See Bank Account types | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**TaxType** | Pointer to **string** | The tax type from taxRates | [optional] 
**EnablePaymentsToAccount** | Pointer to **bool** | Boolean – describes whether account can have payments applied to it | [optional] 
**ShowInExpenseClaims** | Pointer to **bool** | Boolean – describes whether account code is available for use with expense claims | [optional] 
**Class** | Pointer to **string** | See Account Class Types | [optional] [readonly] 
**SystemAccount** | Pointer to **string** | If this is a system account then this element is returned. See System Account types. Note that non-system accounts may have this element set as either “” or null. | [optional] [readonly] 
**ReportingCode** | Pointer to **string** | Shown if set | [optional] 
**ReportingCodeName** | Pointer to **string** | Shown if set | [optional] [readonly] 
**HasAttachments** | Pointer to **bool** | boolean to indicate if an account has an attachment (read only) | [optional] [readonly] [default to false]
**UpdatedDateUTC** | Pointer to **string** | Last modified date UTC format | [optional] [readonly] 
**AddToWatchlist** | Pointer to **bool** | Boolean – describes whether the account is shown in the watchlist widget on the dashboard | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Account) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Account) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Account) SetCode(v string)`

SetCode sets Code field to given value.

### HasCodeField

`func (o *Account) HasCodeField() bool`

HasCodeField returns a boolean if a field has been set.

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *Account) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetAccountID

`func (o *Account) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Account) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *Account) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountIDField

`func (o *Account) HasAccountIDField() bool`

HasAccountIDField returns a boolean if a field has been set.

### GetType

`func (o *Account) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v AccountType)`

SetType sets Type field to given value.

### HasTypeField

`func (o *Account) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetBankAccountNumber

`func (o *Account) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *Account) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *Account) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumberField

`func (o *Account) HasBankAccountNumberField() bool`

HasBankAccountNumberField returns a boolean if a field has been set.

### GetStatus

`func (o *Account) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *Account) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetDescription

`func (o *Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescriptionField

`func (o *Account) HasDescriptionField() bool`

HasDescriptionField returns a boolean if a field has been set.

### GetBankAccountType

`func (o *Account) GetBankAccountType() string`

GetBankAccountType returns the BankAccountType field if non-nil, zero value otherwise.

### GetBankAccountTypeOk

`func (o *Account) GetBankAccountTypeOk() (*string, bool)`

GetBankAccountTypeOk returns a tuple with the BankAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountType

`func (o *Account) SetBankAccountType(v string)`

SetBankAccountType sets BankAccountType field to given value.

### HasBankAccountTypeField

`func (o *Account) HasBankAccountTypeField() bool`

HasBankAccountTypeField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Account) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Account) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Account) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *Account) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetTaxType

`func (o *Account) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *Account) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *Account) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxTypeField

`func (o *Account) HasTaxTypeField() bool`

HasTaxTypeField returns a boolean if a field has been set.

### GetEnablePaymentsToAccount

`func (o *Account) GetEnablePaymentsToAccount() bool`

GetEnablePaymentsToAccount returns the EnablePaymentsToAccount field if non-nil, zero value otherwise.

### GetEnablePaymentsToAccountOk

`func (o *Account) GetEnablePaymentsToAccountOk() (*bool, bool)`

GetEnablePaymentsToAccountOk returns a tuple with the EnablePaymentsToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePaymentsToAccount

`func (o *Account) SetEnablePaymentsToAccount(v bool)`

SetEnablePaymentsToAccount sets EnablePaymentsToAccount field to given value.

### HasEnablePaymentsToAccountField

`func (o *Account) HasEnablePaymentsToAccountField() bool`

HasEnablePaymentsToAccountField returns a boolean if a field has been set.

### GetShowInExpenseClaims

`func (o *Account) GetShowInExpenseClaims() bool`

GetShowInExpenseClaims returns the ShowInExpenseClaims field if non-nil, zero value otherwise.

### GetShowInExpenseClaimsOk

`func (o *Account) GetShowInExpenseClaimsOk() (*bool, bool)`

GetShowInExpenseClaimsOk returns a tuple with the ShowInExpenseClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInExpenseClaims

`func (o *Account) SetShowInExpenseClaims(v bool)`

SetShowInExpenseClaims sets ShowInExpenseClaims field to given value.

### HasShowInExpenseClaimsField

`func (o *Account) HasShowInExpenseClaimsField() bool`

HasShowInExpenseClaimsField returns a boolean if a field has been set.

### GetClass

`func (o *Account) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Account) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Account) SetClass(v string)`

SetClass sets Class field to given value.

### HasClassField

`func (o *Account) HasClassField() bool`

HasClassField returns a boolean if a field has been set.

### GetSystemAccount

`func (o *Account) GetSystemAccount() string`

GetSystemAccount returns the SystemAccount field if non-nil, zero value otherwise.

### GetSystemAccountOk

`func (o *Account) GetSystemAccountOk() (*string, bool)`

GetSystemAccountOk returns a tuple with the SystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccount

`func (o *Account) SetSystemAccount(v string)`

SetSystemAccount sets SystemAccount field to given value.

### HasSystemAccountField

`func (o *Account) HasSystemAccountField() bool`

HasSystemAccountField returns a boolean if a field has been set.

### GetReportingCode

`func (o *Account) GetReportingCode() string`

GetReportingCode returns the ReportingCode field if non-nil, zero value otherwise.

### GetReportingCodeOk

`func (o *Account) GetReportingCodeOk() (*string, bool)`

GetReportingCodeOk returns a tuple with the ReportingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCode

`func (o *Account) SetReportingCode(v string)`

SetReportingCode sets ReportingCode field to given value.

### HasReportingCodeField

`func (o *Account) HasReportingCodeField() bool`

HasReportingCodeField returns a boolean if a field has been set.

### GetReportingCodeName

`func (o *Account) GetReportingCodeName() string`

GetReportingCodeName returns the ReportingCodeName field if non-nil, zero value otherwise.

### GetReportingCodeNameOk

`func (o *Account) GetReportingCodeNameOk() (*string, bool)`

GetReportingCodeNameOk returns a tuple with the ReportingCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingCodeName

`func (o *Account) SetReportingCodeName(v string)`

SetReportingCodeName sets ReportingCodeName field to given value.

### HasReportingCodeNameField

`func (o *Account) HasReportingCodeNameField() bool`

HasReportingCodeNameField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *Account) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Account) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Account) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *Account) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Account) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Account) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Account) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Account) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetAddToWatchlist

`func (o *Account) GetAddToWatchlist() bool`

GetAddToWatchlist returns the AddToWatchlist field if non-nil, zero value otherwise.

### GetAddToWatchlistOk

`func (o *Account) GetAddToWatchlistOk() (*bool, bool)`

GetAddToWatchlistOk returns a tuple with the AddToWatchlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToWatchlist

`func (o *Account) SetAddToWatchlist(v bool)`

SetAddToWatchlist sets AddToWatchlist field to given value.

### HasAddToWatchlistField

`func (o *Account) HasAddToWatchlistField() bool`

HasAddToWatchlistField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Account) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Account) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Account) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Account) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


