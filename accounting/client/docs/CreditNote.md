# CreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | See Credit Note Types | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Date** | Pointer to **string** | The date the credit note is issued YYYY-MM-DD. If the Date element is not specified then it will default to the current date based on the timezone setting of the organisation | [optional] 
**DueDate** | Pointer to **string** | Date invoice is due – YYYY-MM-DD | [optional] 
**Status** | Pointer to **string** | See Credit Note Status Codes | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | See Invoice Line Items | [optional] 
**SubTotal** | Pointer to **float64** | The subtotal of the credit note excluding taxes | [optional] 
**TotalTax** | Pointer to **float64** | The total tax on the credit note | [optional] 
**Total** | Pointer to **float64** | The total of the Credit Note(subtotal + total tax) | [optional] 
**CISDeduction** | Pointer to **float64** | CIS deduction for UK contractors | [optional] [readonly] 
**CISRate** | Pointer to **float64** | CIS Deduction rate for the organisation | [optional] [readonly] 
**UpdatedDateUTC** | Pointer to **string** | UTC timestamp of last update to the credit note | [optional] [readonly] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**FullyPaidOnDate** | Pointer to **string** | Date when credit note was fully paid(UTC format) | [optional] 
**CreditNoteID** | Pointer to **string** | Xero generated unique identifier | [optional] 
**CreditNoteNumber** | Pointer to **string** | ACCRECCREDIT – Unique alpha numeric code identifying credit note (when missing will auto-generate from your Organisation Invoice Settings) | [optional] 
**Reference** | Pointer to **string** | ACCRECCREDIT only – additional reference number | [optional] 
**SentToContact** | Pointer to **bool** | boolean to indicate if a credit note has been sent to a contact via  the Xero app (currently read only) | [optional] [readonly] 
**CurrencyRate** | Pointer to **float64** | The currency rate for a multicurrency invoice. If no rate is specified, the XE.com day rate is used | [optional] 
**RemainingCredit** | Pointer to **float64** | The remaining credit balance on the Credit Note | [optional] 
**Allocations** | Pointer to [**[]Allocation**](Allocation.md) | See Allocations | [optional] 
**AppliedAmount** | Pointer to **float64** | The amount of applied to an invoice | [optional] 
**Payments** | Pointer to [**[]Payment**](Payment.md) | See Payments | [optional] 
**BrandingThemeID** | Pointer to **string** | See BrandingThemes | [optional] 
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**HasAttachments** | Pointer to **bool** | boolean to indicate if a credit note has an attachment | [optional] [default to false]
**HasErrors** | Pointer to **bool** | A boolean to indicate if a credit note has an validation errors | [optional] [default to false]
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 
**Warnings** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of warning messages from the API | [optional] 

## Methods

### NewCreditNote

`func NewCreditNote() *CreditNote`

NewCreditNote instantiates a new CreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteWithDefaults

`func NewCreditNoteWithDefaults() *CreditNote`

NewCreditNoteWithDefaults instantiates a new CreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreditNote) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditNote) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditNote) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *CreditNote) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetContact

`func (o *CreditNote) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CreditNote) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CreditNote) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *CreditNote) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetDate

`func (o *CreditNote) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreditNote) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreditNote) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *CreditNote) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetDueDate

`func (o *CreditNote) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CreditNote) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CreditNote) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDateField

`func (o *CreditNote) HasDueDateField() bool`

HasDueDateField returns a boolean if a field has been set.

### GetStatus

`func (o *CreditNote) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditNote) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditNote) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *CreditNote) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *CreditNote) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *CreditNote) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *CreditNote) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *CreditNote) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetLineItems

`func (o *CreditNote) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreditNote) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreditNote) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItemsField

`func (o *CreditNote) HasLineItemsField() bool`

HasLineItemsField returns a boolean if a field has been set.

### GetSubTotal

`func (o *CreditNote) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *CreditNote) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *CreditNote) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *CreditNote) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *CreditNote) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *CreditNote) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *CreditNote) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *CreditNote) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *CreditNote) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CreditNote) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CreditNote) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *CreditNote) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetCISDeduction

`func (o *CreditNote) GetCISDeduction() float64`

GetCISDeduction returns the CISDeduction field if non-nil, zero value otherwise.

### GetCISDeductionOk

`func (o *CreditNote) GetCISDeductionOk() (*float64, bool)`

GetCISDeductionOk returns a tuple with the CISDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCISDeduction

`func (o *CreditNote) SetCISDeduction(v float64)`

SetCISDeduction sets CISDeduction field to given value.

### HasCISDeductionField

`func (o *CreditNote) HasCISDeductionField() bool`

HasCISDeductionField returns a boolean if a field has been set.

### GetCISRate

`func (o *CreditNote) GetCISRate() float64`

GetCISRate returns the CISRate field if non-nil, zero value otherwise.

### GetCISRateOk

`func (o *CreditNote) GetCISRateOk() (*float64, bool)`

GetCISRateOk returns a tuple with the CISRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCISRate

`func (o *CreditNote) SetCISRate(v float64)`

SetCISRate sets CISRate field to given value.

### HasCISRateField

`func (o *CreditNote) HasCISRateField() bool`

HasCISRateField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *CreditNote) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *CreditNote) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *CreditNote) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *CreditNote) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CreditNote) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CreditNote) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CreditNote) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *CreditNote) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetFullyPaidOnDate

`func (o *CreditNote) GetFullyPaidOnDate() string`

GetFullyPaidOnDate returns the FullyPaidOnDate field if non-nil, zero value otherwise.

### GetFullyPaidOnDateOk

`func (o *CreditNote) GetFullyPaidOnDateOk() (*string, bool)`

GetFullyPaidOnDateOk returns a tuple with the FullyPaidOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyPaidOnDate

`func (o *CreditNote) SetFullyPaidOnDate(v string)`

SetFullyPaidOnDate sets FullyPaidOnDate field to given value.

### HasFullyPaidOnDateField

`func (o *CreditNote) HasFullyPaidOnDateField() bool`

HasFullyPaidOnDateField returns a boolean if a field has been set.

### GetCreditNoteID

`func (o *CreditNote) GetCreditNoteID() string`

GetCreditNoteID returns the CreditNoteID field if non-nil, zero value otherwise.

### GetCreditNoteIDOk

`func (o *CreditNote) GetCreditNoteIDOk() (*string, bool)`

GetCreditNoteIDOk returns a tuple with the CreditNoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteID

`func (o *CreditNote) SetCreditNoteID(v string)`

SetCreditNoteID sets CreditNoteID field to given value.

### HasCreditNoteIDField

`func (o *CreditNote) HasCreditNoteIDField() bool`

HasCreditNoteIDField returns a boolean if a field has been set.

### GetCreditNoteNumber

`func (o *CreditNote) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *CreditNote) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *CreditNote) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.

### HasCreditNoteNumberField

`func (o *CreditNote) HasCreditNoteNumberField() bool`

HasCreditNoteNumberField returns a boolean if a field has been set.

### GetReference

`func (o *CreditNote) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreditNote) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreditNote) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *CreditNote) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetSentToContact

`func (o *CreditNote) GetSentToContact() bool`

GetSentToContact returns the SentToContact field if non-nil, zero value otherwise.

### GetSentToContactOk

`func (o *CreditNote) GetSentToContactOk() (*bool, bool)`

GetSentToContactOk returns a tuple with the SentToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToContact

`func (o *CreditNote) SetSentToContact(v bool)`

SetSentToContact sets SentToContact field to given value.

### HasSentToContactField

`func (o *CreditNote) HasSentToContactField() bool`

HasSentToContactField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *CreditNote) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *CreditNote) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *CreditNote) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *CreditNote) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetRemainingCredit

`func (o *CreditNote) GetRemainingCredit() float64`

GetRemainingCredit returns the RemainingCredit field if non-nil, zero value otherwise.

### GetRemainingCreditOk

`func (o *CreditNote) GetRemainingCreditOk() (*float64, bool)`

GetRemainingCreditOk returns a tuple with the RemainingCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCredit

`func (o *CreditNote) SetRemainingCredit(v float64)`

SetRemainingCredit sets RemainingCredit field to given value.

### HasRemainingCreditField

`func (o *CreditNote) HasRemainingCreditField() bool`

HasRemainingCreditField returns a boolean if a field has been set.

### GetAllocations

`func (o *CreditNote) GetAllocations() []Allocation`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *CreditNote) GetAllocationsOk() (*[]Allocation, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *CreditNote) SetAllocations(v []Allocation)`

SetAllocations sets Allocations field to given value.

### HasAllocationsField

`func (o *CreditNote) HasAllocationsField() bool`

HasAllocationsField returns a boolean if a field has been set.

### GetAppliedAmount

`func (o *CreditNote) GetAppliedAmount() float64`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *CreditNote) GetAppliedAmountOk() (*float64, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *CreditNote) SetAppliedAmount(v float64)`

SetAppliedAmount sets AppliedAmount field to given value.

### HasAppliedAmountField

`func (o *CreditNote) HasAppliedAmountField() bool`

HasAppliedAmountField returns a boolean if a field has been set.

### GetPayments

`func (o *CreditNote) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *CreditNote) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *CreditNote) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPaymentsField

`func (o *CreditNote) HasPaymentsField() bool`

HasPaymentsField returns a boolean if a field has been set.

### GetBrandingThemeID

`func (o *CreditNote) GetBrandingThemeID() string`

GetBrandingThemeID returns the BrandingThemeID field if non-nil, zero value otherwise.

### GetBrandingThemeIDOk

`func (o *CreditNote) GetBrandingThemeIDOk() (*string, bool)`

GetBrandingThemeIDOk returns a tuple with the BrandingThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingThemeID

`func (o *CreditNote) SetBrandingThemeID(v string)`

SetBrandingThemeID sets BrandingThemeID field to given value.

### HasBrandingThemeIDField

`func (o *CreditNote) HasBrandingThemeIDField() bool`

HasBrandingThemeIDField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *CreditNote) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *CreditNote) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *CreditNote) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *CreditNote) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *CreditNote) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *CreditNote) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *CreditNote) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *CreditNote) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetHasErrors

`func (o *CreditNote) GetHasErrors() bool`

GetHasErrors returns the HasErrors field if non-nil, zero value otherwise.

### GetHasErrorsOk

`func (o *CreditNote) GetHasErrorsOk() (*bool, bool)`

GetHasErrorsOk returns a tuple with the HasErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasErrors

`func (o *CreditNote) SetHasErrors(v bool)`

SetHasErrors sets HasErrors field to given value.

### HasHasErrorsField

`func (o *CreditNote) HasHasErrorsField() bool`

HasHasErrorsField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *CreditNote) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CreditNote) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CreditNote) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *CreditNote) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.

### GetWarnings

`func (o *CreditNote) GetWarnings() []ValidationError`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CreditNote) GetWarningsOk() (*[]ValidationError, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CreditNote) SetWarnings(v []ValidationError)`

SetWarnings sets Warnings field to given value.

### HasWarningsField

`func (o *CreditNote) HasWarningsField() bool`

HasWarningsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


