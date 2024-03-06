# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | See Invoice Types | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | See LineItems | [optional] 
**Date** | Pointer to **string** | Date invoice was issued – YYYY-MM-DD. If the Date element is not specified it will default to the current date based on the timezone setting of the organisation | [optional] 
**DueDate** | Pointer to **string** | Date invoice is due – YYYY-MM-DD | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**InvoiceNumber** | Pointer to **string** | ACCREC – Unique alpha numeric code identifying invoice (when missing will auto-generate from your Organisation Invoice Settings) (max length &#x3D; 255) | [optional] 
**Reference** | Pointer to **string** | ACCREC only – additional reference number | [optional] 
**BrandingThemeID** | Pointer to **string** | See BrandingThemes | [optional] 
**Url** | Pointer to **string** | URL link to a source document – shown as “Go to [appName]” in the Xero app | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**CurrencyRate** | Pointer to **float64** | The currency rate for a multicurrency invoice. If no rate is specified, the XE.com day rate is used. (max length &#x3D; [18].[6]) | [optional] 
**Status** | Pointer to **string** | See Invoice Status Codes | [optional] 
**SentToContact** | Pointer to **bool** | Boolean to set whether the invoice in the Xero app should be marked as “sent”. This can be set only on invoices that have been approved | [optional] 
**ExpectedPaymentDate** | Pointer to **string** | Shown on sales invoices (Accounts Receivable) when this has been set | [optional] 
**PlannedPaymentDate** | Pointer to **string** | Shown on bills (Accounts Payable) when this has been set | [optional] 
**CISDeduction** | Pointer to **float64** | CIS deduction for UK contractors | [optional] [readonly] 
**CISRate** | Pointer to **float64** | CIS Deduction rate for the organisation | [optional] [readonly] 
**SubTotal** | Pointer to **float64** | Total of invoice excluding taxes | [optional] [readonly] 
**TotalTax** | Pointer to **float64** | Total tax on invoice | [optional] [readonly] 
**Total** | Pointer to **float64** | Total of Invoice tax inclusive (i.e. SubTotal + TotalTax). This will be ignored if it doesn’t equal the sum of the LineAmounts | [optional] [readonly] 
**TotalDiscount** | Pointer to **float64** | Total of discounts applied on the invoice line items | [optional] [readonly] 
**InvoiceID** | Pointer to **string** | Xero generated unique identifier for invoice | [optional] 
**RepeatingInvoiceID** | Pointer to **string** | Xero generated unique identifier for repeating invoices | [optional] 
**HasAttachments** | Pointer to **bool** | boolean to indicate if an invoice has an attachment | [optional] [readonly] [default to false]
**IsDiscounted** | Pointer to **bool** | boolean to indicate if an invoice has a discount | [optional] [readonly] 
**Payments** | Pointer to [**[]Payment**](Payment.md) | See Payments | [optional] [readonly] 
**Prepayments** | Pointer to [**[]Prepayment**](Prepayment.md) | See Prepayments | [optional] [readonly] 
**Overpayments** | Pointer to [**[]Overpayment**](Overpayment.md) | See Overpayments | [optional] [readonly] 
**AmountDue** | Pointer to **float64** | Amount remaining to be paid on invoice | [optional] [readonly] 
**AmountPaid** | Pointer to **float64** | Sum of payments received for invoice | [optional] [readonly] 
**FullyPaidOnDate** | Pointer to **string** | The date the invoice was fully paid. Only returned on fully paid invoices | [optional] [readonly] 
**AmountCredited** | Pointer to **float64** | Sum of all credit notes, over-payments and pre-payments applied to invoice | [optional] [readonly] 
**UpdatedDateUTC** | Pointer to **string** | Last modified date UTC format | [optional] [readonly] 
**CreditNotes** | Pointer to [**[]CreditNote**](CreditNote.md) | Details of credit notes that have been applied to an invoice | [optional] [readonly] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 
**HasErrors** | Pointer to **bool** | A boolean to indicate if a invoice has an validation errors | [optional] [default to false]
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 
**Warnings** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of warning messages from the API | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Invoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Invoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Invoice) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *Invoice) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetContact

`func (o *Invoice) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Invoice) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Invoice) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *Invoice) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetLineItems

`func (o *Invoice) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Invoice) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Invoice) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItemsField

`func (o *Invoice) HasLineItemsField() bool`

HasLineItemsField returns a boolean if a field has been set.

### GetDate

`func (o *Invoice) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Invoice) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Invoice) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *Invoice) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetDueDate

`func (o *Invoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDateField

`func (o *Invoice) HasDueDateField() bool`

HasDueDateField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *Invoice) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *Invoice) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *Invoice) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *Invoice) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *Invoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Invoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Invoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumberField

`func (o *Invoice) HasInvoiceNumberField() bool`

HasInvoiceNumberField returns a boolean if a field has been set.

### GetReference

`func (o *Invoice) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Invoice) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Invoice) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *Invoice) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetBrandingThemeID

`func (o *Invoice) GetBrandingThemeID() string`

GetBrandingThemeID returns the BrandingThemeID field if non-nil, zero value otherwise.

### GetBrandingThemeIDOk

`func (o *Invoice) GetBrandingThemeIDOk() (*string, bool)`

GetBrandingThemeIDOk returns a tuple with the BrandingThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingThemeID

`func (o *Invoice) SetBrandingThemeID(v string)`

SetBrandingThemeID sets BrandingThemeID field to given value.

### HasBrandingThemeIDField

`func (o *Invoice) HasBrandingThemeIDField() bool`

HasBrandingThemeIDField returns a boolean if a field has been set.

### GetUrl

`func (o *Invoice) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Invoice) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Invoice) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrlField

`func (o *Invoice) HasUrlField() bool`

HasUrlField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Invoice) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Invoice) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Invoice) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *Invoice) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *Invoice) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *Invoice) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *Invoice) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *Invoice) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *Invoice) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetSentToContact

`func (o *Invoice) GetSentToContact() bool`

GetSentToContact returns the SentToContact field if non-nil, zero value otherwise.

### GetSentToContactOk

`func (o *Invoice) GetSentToContactOk() (*bool, bool)`

GetSentToContactOk returns a tuple with the SentToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToContact

`func (o *Invoice) SetSentToContact(v bool)`

SetSentToContact sets SentToContact field to given value.

### HasSentToContactField

`func (o *Invoice) HasSentToContactField() bool`

HasSentToContactField returns a boolean if a field has been set.

### GetExpectedPaymentDate

`func (o *Invoice) GetExpectedPaymentDate() string`

GetExpectedPaymentDate returns the ExpectedPaymentDate field if non-nil, zero value otherwise.

### GetExpectedPaymentDateOk

`func (o *Invoice) GetExpectedPaymentDateOk() (*string, bool)`

GetExpectedPaymentDateOk returns a tuple with the ExpectedPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedPaymentDate

`func (o *Invoice) SetExpectedPaymentDate(v string)`

SetExpectedPaymentDate sets ExpectedPaymentDate field to given value.

### HasExpectedPaymentDateField

`func (o *Invoice) HasExpectedPaymentDateField() bool`

HasExpectedPaymentDateField returns a boolean if a field has been set.

### GetPlannedPaymentDate

`func (o *Invoice) GetPlannedPaymentDate() string`

GetPlannedPaymentDate returns the PlannedPaymentDate field if non-nil, zero value otherwise.

### GetPlannedPaymentDateOk

`func (o *Invoice) GetPlannedPaymentDateOk() (*string, bool)`

GetPlannedPaymentDateOk returns a tuple with the PlannedPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedPaymentDate

`func (o *Invoice) SetPlannedPaymentDate(v string)`

SetPlannedPaymentDate sets PlannedPaymentDate field to given value.

### HasPlannedPaymentDateField

`func (o *Invoice) HasPlannedPaymentDateField() bool`

HasPlannedPaymentDateField returns a boolean if a field has been set.

### GetCISDeduction

`func (o *Invoice) GetCISDeduction() float64`

GetCISDeduction returns the CISDeduction field if non-nil, zero value otherwise.

### GetCISDeductionOk

`func (o *Invoice) GetCISDeductionOk() (*float64, bool)`

GetCISDeductionOk returns a tuple with the CISDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCISDeduction

`func (o *Invoice) SetCISDeduction(v float64)`

SetCISDeduction sets CISDeduction field to given value.

### HasCISDeductionField

`func (o *Invoice) HasCISDeductionField() bool`

HasCISDeductionField returns a boolean if a field has been set.

### GetCISRate

`func (o *Invoice) GetCISRate() float64`

GetCISRate returns the CISRate field if non-nil, zero value otherwise.

### GetCISRateOk

`func (o *Invoice) GetCISRateOk() (*float64, bool)`

GetCISRateOk returns a tuple with the CISRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCISRate

`func (o *Invoice) SetCISRate(v float64)`

SetCISRate sets CISRate field to given value.

### HasCISRateField

`func (o *Invoice) HasCISRateField() bool`

HasCISRateField returns a boolean if a field has been set.

### GetSubTotal

`func (o *Invoice) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *Invoice) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *Invoice) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *Invoice) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *Invoice) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *Invoice) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *Invoice) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *Invoice) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *Invoice) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *Invoice) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *Invoice) GetTotalDiscount() float64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *Invoice) GetTotalDiscountOk() (*float64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *Invoice) SetTotalDiscount(v float64)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscountField

`func (o *Invoice) HasTotalDiscountField() bool`

HasTotalDiscountField returns a boolean if a field has been set.

### GetInvoiceID

`func (o *Invoice) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *Invoice) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *Invoice) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceIDField

`func (o *Invoice) HasInvoiceIDField() bool`

HasInvoiceIDField returns a boolean if a field has been set.

### GetRepeatingInvoiceID

`func (o *Invoice) GetRepeatingInvoiceID() string`

GetRepeatingInvoiceID returns the RepeatingInvoiceID field if non-nil, zero value otherwise.

### GetRepeatingInvoiceIDOk

`func (o *Invoice) GetRepeatingInvoiceIDOk() (*string, bool)`

GetRepeatingInvoiceIDOk returns a tuple with the RepeatingInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatingInvoiceID

`func (o *Invoice) SetRepeatingInvoiceID(v string)`

SetRepeatingInvoiceID sets RepeatingInvoiceID field to given value.

### HasRepeatingInvoiceIDField

`func (o *Invoice) HasRepeatingInvoiceIDField() bool`

HasRepeatingInvoiceIDField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *Invoice) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Invoice) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Invoice) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *Invoice) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetIsDiscounted

`func (o *Invoice) GetIsDiscounted() bool`

GetIsDiscounted returns the IsDiscounted field if non-nil, zero value otherwise.

### GetIsDiscountedOk

`func (o *Invoice) GetIsDiscountedOk() (*bool, bool)`

GetIsDiscountedOk returns a tuple with the IsDiscounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiscounted

`func (o *Invoice) SetIsDiscounted(v bool)`

SetIsDiscounted sets IsDiscounted field to given value.

### HasIsDiscountedField

`func (o *Invoice) HasIsDiscountedField() bool`

HasIsDiscountedField returns a boolean if a field has been set.

### GetPayments

`func (o *Invoice) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Invoice) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Invoice) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPaymentsField

`func (o *Invoice) HasPaymentsField() bool`

HasPaymentsField returns a boolean if a field has been set.

### GetPrepayments

`func (o *Invoice) GetPrepayments() []Prepayment`

GetPrepayments returns the Prepayments field if non-nil, zero value otherwise.

### GetPrepaymentsOk

`func (o *Invoice) GetPrepaymentsOk() (*[]Prepayment, bool)`

GetPrepaymentsOk returns a tuple with the Prepayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepayments

`func (o *Invoice) SetPrepayments(v []Prepayment)`

SetPrepayments sets Prepayments field to given value.

### HasPrepaymentsField

`func (o *Invoice) HasPrepaymentsField() bool`

HasPrepaymentsField returns a boolean if a field has been set.

### GetOverpayments

`func (o *Invoice) GetOverpayments() []Overpayment`

GetOverpayments returns the Overpayments field if non-nil, zero value otherwise.

### GetOverpaymentsOk

`func (o *Invoice) GetOverpaymentsOk() (*[]Overpayment, bool)`

GetOverpaymentsOk returns a tuple with the Overpayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverpayments

`func (o *Invoice) SetOverpayments(v []Overpayment)`

SetOverpayments sets Overpayments field to given value.

### HasOverpaymentsField

`func (o *Invoice) HasOverpaymentsField() bool`

HasOverpaymentsField returns a boolean if a field has been set.

### GetAmountDue

`func (o *Invoice) GetAmountDue() float64`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *Invoice) GetAmountDueOk() (*float64, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *Invoice) SetAmountDue(v float64)`

SetAmountDue sets AmountDue field to given value.

### HasAmountDueField

`func (o *Invoice) HasAmountDueField() bool`

HasAmountDueField returns a boolean if a field has been set.

### GetAmountPaid

`func (o *Invoice) GetAmountPaid() float64`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *Invoice) GetAmountPaidOk() (*float64, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *Invoice) SetAmountPaid(v float64)`

SetAmountPaid sets AmountPaid field to given value.

### HasAmountPaidField

`func (o *Invoice) HasAmountPaidField() bool`

HasAmountPaidField returns a boolean if a field has been set.

### GetFullyPaidOnDate

`func (o *Invoice) GetFullyPaidOnDate() string`

GetFullyPaidOnDate returns the FullyPaidOnDate field if non-nil, zero value otherwise.

### GetFullyPaidOnDateOk

`func (o *Invoice) GetFullyPaidOnDateOk() (*string, bool)`

GetFullyPaidOnDateOk returns a tuple with the FullyPaidOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyPaidOnDate

`func (o *Invoice) SetFullyPaidOnDate(v string)`

SetFullyPaidOnDate sets FullyPaidOnDate field to given value.

### HasFullyPaidOnDateField

`func (o *Invoice) HasFullyPaidOnDateField() bool`

HasFullyPaidOnDateField returns a boolean if a field has been set.

### GetAmountCredited

`func (o *Invoice) GetAmountCredited() float64`

GetAmountCredited returns the AmountCredited field if non-nil, zero value otherwise.

### GetAmountCreditedOk

`func (o *Invoice) GetAmountCreditedOk() (*float64, bool)`

GetAmountCreditedOk returns a tuple with the AmountCredited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCredited

`func (o *Invoice) SetAmountCredited(v float64)`

SetAmountCredited sets AmountCredited field to given value.

### HasAmountCreditedField

`func (o *Invoice) HasAmountCreditedField() bool`

HasAmountCreditedField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Invoice) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Invoice) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Invoice) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Invoice) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetCreditNotes

`func (o *Invoice) GetCreditNotes() []CreditNote`

GetCreditNotes returns the CreditNotes field if non-nil, zero value otherwise.

### GetCreditNotesOk

`func (o *Invoice) GetCreditNotesOk() (*[]CreditNote, bool)`

GetCreditNotesOk returns a tuple with the CreditNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotes

`func (o *Invoice) SetCreditNotes(v []CreditNote)`

SetCreditNotes sets CreditNotes field to given value.

### HasCreditNotesField

`func (o *Invoice) HasCreditNotesField() bool`

HasCreditNotesField returns a boolean if a field has been set.

### GetAttachments

`func (o *Invoice) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Invoice) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Invoice) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachmentsField

`func (o *Invoice) HasAttachmentsField() bool`

HasAttachmentsField returns a boolean if a field has been set.

### GetHasErrors

`func (o *Invoice) GetHasErrors() bool`

GetHasErrors returns the HasErrors field if non-nil, zero value otherwise.

### GetHasErrorsOk

`func (o *Invoice) GetHasErrorsOk() (*bool, bool)`

GetHasErrorsOk returns a tuple with the HasErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasErrors

`func (o *Invoice) SetHasErrors(v bool)`

SetHasErrors sets HasErrors field to given value.

### HasHasErrorsField

`func (o *Invoice) HasHasErrorsField() bool`

HasHasErrorsField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *Invoice) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *Invoice) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *Invoice) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *Invoice) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Invoice) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Invoice) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Invoice) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Invoice) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.

### GetWarnings

`func (o *Invoice) GetWarnings() []ValidationError`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Invoice) GetWarningsOk() (*[]ValidationError, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Invoice) SetWarnings(v []ValidationError)`

SetWarnings sets Warnings field to given value.

### HasWarningsField

`func (o *Invoice) HasWarningsField() bool`

HasWarningsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


