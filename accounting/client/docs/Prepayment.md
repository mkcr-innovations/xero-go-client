# Prepayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | See Prepayment Types | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Date** | Pointer to **string** | The date the prepayment is created YYYY-MM-DD | [optional] 
**Status** | Pointer to **string** | See Prepayment Status Codes | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | See Prepayment Line Items | [optional] 
**SubTotal** | Pointer to **float64** | The subtotal of the prepayment excluding taxes | [optional] 
**TotalTax** | Pointer to **float64** | The total tax on the prepayment | [optional] 
**Total** | Pointer to **float64** | The total of the prepayment(subtotal + total tax) | [optional] 
**Reference** | Pointer to **string** | Returns Invoice number field. Reference field isn&#39;t available. | [optional] [readonly] 
**UpdatedDateUTC** | Pointer to **string** | UTC timestamp of last update to the prepayment | [optional] [readonly] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**PrepaymentID** | Pointer to **string** | Xero generated unique identifier | [optional] 
**CurrencyRate** | Pointer to **float64** | The currency rate for a multicurrency prepayment. If no rate is specified, the XE.com day rate is used | [optional] 
**RemainingCredit** | Pointer to **float64** | The remaining credit balance on the prepayment | [optional] 
**Allocations** | Pointer to [**[]Allocation**](Allocation.md) | See Allocations | [optional] 
**Payments** | Pointer to [**[]Payment**](Payment.md) | See Payments | [optional] 
**AppliedAmount** | Pointer to **float64** | The amount of applied to an invoice | [optional] 
**HasAttachments** | Pointer to **bool** | boolean to indicate if a prepayment has an attachment | [optional] [readonly] [default to false]
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | See Attachments | [optional] 

## Methods

### NewPrepayment

`func NewPrepayment() *Prepayment`

NewPrepayment instantiates a new Prepayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepaymentWithDefaults

`func NewPrepaymentWithDefaults() *Prepayment`

NewPrepaymentWithDefaults instantiates a new Prepayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Prepayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Prepayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Prepayment) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *Prepayment) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetContact

`func (o *Prepayment) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Prepayment) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Prepayment) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *Prepayment) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetDate

`func (o *Prepayment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Prepayment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Prepayment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *Prepayment) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetStatus

`func (o *Prepayment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Prepayment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Prepayment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *Prepayment) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *Prepayment) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *Prepayment) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *Prepayment) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *Prepayment) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetLineItems

`func (o *Prepayment) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Prepayment) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Prepayment) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItemsField

`func (o *Prepayment) HasLineItemsField() bool`

HasLineItemsField returns a boolean if a field has been set.

### GetSubTotal

`func (o *Prepayment) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *Prepayment) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *Prepayment) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *Prepayment) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *Prepayment) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *Prepayment) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *Prepayment) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *Prepayment) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *Prepayment) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Prepayment) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Prepayment) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *Prepayment) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetReference

`func (o *Prepayment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Prepayment) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Prepayment) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *Prepayment) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Prepayment) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Prepayment) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Prepayment) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Prepayment) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Prepayment) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Prepayment) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Prepayment) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *Prepayment) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetPrepaymentID

`func (o *Prepayment) GetPrepaymentID() string`

GetPrepaymentID returns the PrepaymentID field if non-nil, zero value otherwise.

### GetPrepaymentIDOk

`func (o *Prepayment) GetPrepaymentIDOk() (*string, bool)`

GetPrepaymentIDOk returns a tuple with the PrepaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaymentID

`func (o *Prepayment) SetPrepaymentID(v string)`

SetPrepaymentID sets PrepaymentID field to given value.

### HasPrepaymentIDField

`func (o *Prepayment) HasPrepaymentIDField() bool`

HasPrepaymentIDField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *Prepayment) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *Prepayment) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *Prepayment) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *Prepayment) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetRemainingCredit

`func (o *Prepayment) GetRemainingCredit() float64`

GetRemainingCredit returns the RemainingCredit field if non-nil, zero value otherwise.

### GetRemainingCreditOk

`func (o *Prepayment) GetRemainingCreditOk() (*float64, bool)`

GetRemainingCreditOk returns a tuple with the RemainingCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCredit

`func (o *Prepayment) SetRemainingCredit(v float64)`

SetRemainingCredit sets RemainingCredit field to given value.

### HasRemainingCreditField

`func (o *Prepayment) HasRemainingCreditField() bool`

HasRemainingCreditField returns a boolean if a field has been set.

### GetAllocations

`func (o *Prepayment) GetAllocations() []Allocation`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *Prepayment) GetAllocationsOk() (*[]Allocation, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *Prepayment) SetAllocations(v []Allocation)`

SetAllocations sets Allocations field to given value.

### HasAllocationsField

`func (o *Prepayment) HasAllocationsField() bool`

HasAllocationsField returns a boolean if a field has been set.

### GetPayments

`func (o *Prepayment) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Prepayment) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Prepayment) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPaymentsField

`func (o *Prepayment) HasPaymentsField() bool`

HasPaymentsField returns a boolean if a field has been set.

### GetAppliedAmount

`func (o *Prepayment) GetAppliedAmount() float64`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *Prepayment) GetAppliedAmountOk() (*float64, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *Prepayment) SetAppliedAmount(v float64)`

SetAppliedAmount sets AppliedAmount field to given value.

### HasAppliedAmountField

`func (o *Prepayment) HasAppliedAmountField() bool`

HasAppliedAmountField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *Prepayment) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Prepayment) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Prepayment) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *Prepayment) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetAttachments

`func (o *Prepayment) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Prepayment) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Prepayment) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachmentsField

`func (o *Prepayment) HasAttachmentsField() bool`

HasAttachmentsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


