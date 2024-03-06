# Overpayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | See Overpayment Types | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Date** | Pointer to **string** | The date the overpayment is created YYYY-MM-DD | [optional] 
**Status** | Pointer to **string** | See Overpayment Status Codes | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | See Overpayment Line Items | [optional] 
**SubTotal** | Pointer to **float64** | The subtotal of the overpayment excluding taxes | [optional] 
**TotalTax** | Pointer to **float64** | The total tax on the overpayment | [optional] 
**Total** | Pointer to **float64** | The total of the overpayment (subtotal + total tax) | [optional] 
**UpdatedDateUTC** | Pointer to **string** | UTC timestamp of last update to the overpayment | [optional] [readonly] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**OverpaymentID** | Pointer to **string** | Xero generated unique identifier | [optional] 
**CurrencyRate** | Pointer to **float64** | The currency rate for a multicurrency overpayment. If no rate is specified, the XE.com day rate is used | [optional] 
**RemainingCredit** | Pointer to **float64** | The remaining credit balance on the overpayment | [optional] 
**Allocations** | Pointer to [**[]Allocation**](Allocation.md) | See Allocations | [optional] 
**AppliedAmount** | Pointer to **float64** | The amount of applied to an invoice | [optional] 
**Payments** | Pointer to [**[]Payment**](Payment.md) | See Payments | [optional] 
**HasAttachments** | Pointer to **bool** | boolean to indicate if a overpayment has an attachment | [optional] [readonly] [default to false]
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | See Attachments | [optional] 

## Methods

### NewOverpayment

`func NewOverpayment() *Overpayment`

NewOverpayment instantiates a new Overpayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverpaymentWithDefaults

`func NewOverpaymentWithDefaults() *Overpayment`

NewOverpaymentWithDefaults instantiates a new Overpayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Overpayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Overpayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Overpayment) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *Overpayment) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetContact

`func (o *Overpayment) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Overpayment) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Overpayment) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *Overpayment) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetDate

`func (o *Overpayment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Overpayment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Overpayment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *Overpayment) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetStatus

`func (o *Overpayment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Overpayment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Overpayment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *Overpayment) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *Overpayment) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *Overpayment) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *Overpayment) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *Overpayment) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetLineItems

`func (o *Overpayment) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Overpayment) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Overpayment) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItemsField

`func (o *Overpayment) HasLineItemsField() bool`

HasLineItemsField returns a boolean if a field has been set.

### GetSubTotal

`func (o *Overpayment) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *Overpayment) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *Overpayment) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *Overpayment) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *Overpayment) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *Overpayment) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *Overpayment) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *Overpayment) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *Overpayment) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Overpayment) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Overpayment) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *Overpayment) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Overpayment) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Overpayment) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Overpayment) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Overpayment) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Overpayment) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Overpayment) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Overpayment) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *Overpayment) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetOverpaymentID

`func (o *Overpayment) GetOverpaymentID() string`

GetOverpaymentID returns the OverpaymentID field if non-nil, zero value otherwise.

### GetOverpaymentIDOk

`func (o *Overpayment) GetOverpaymentIDOk() (*string, bool)`

GetOverpaymentIDOk returns a tuple with the OverpaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverpaymentID

`func (o *Overpayment) SetOverpaymentID(v string)`

SetOverpaymentID sets OverpaymentID field to given value.

### HasOverpaymentIDField

`func (o *Overpayment) HasOverpaymentIDField() bool`

HasOverpaymentIDField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *Overpayment) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *Overpayment) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *Overpayment) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *Overpayment) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetRemainingCredit

`func (o *Overpayment) GetRemainingCredit() float64`

GetRemainingCredit returns the RemainingCredit field if non-nil, zero value otherwise.

### GetRemainingCreditOk

`func (o *Overpayment) GetRemainingCreditOk() (*float64, bool)`

GetRemainingCreditOk returns a tuple with the RemainingCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCredit

`func (o *Overpayment) SetRemainingCredit(v float64)`

SetRemainingCredit sets RemainingCredit field to given value.

### HasRemainingCreditField

`func (o *Overpayment) HasRemainingCreditField() bool`

HasRemainingCreditField returns a boolean if a field has been set.

### GetAllocations

`func (o *Overpayment) GetAllocations() []Allocation`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *Overpayment) GetAllocationsOk() (*[]Allocation, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *Overpayment) SetAllocations(v []Allocation)`

SetAllocations sets Allocations field to given value.

### HasAllocationsField

`func (o *Overpayment) HasAllocationsField() bool`

HasAllocationsField returns a boolean if a field has been set.

### GetAppliedAmount

`func (o *Overpayment) GetAppliedAmount() float64`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *Overpayment) GetAppliedAmountOk() (*float64, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *Overpayment) SetAppliedAmount(v float64)`

SetAppliedAmount sets AppliedAmount field to given value.

### HasAppliedAmountField

`func (o *Overpayment) HasAppliedAmountField() bool`

HasAppliedAmountField returns a boolean if a field has been set.

### GetPayments

`func (o *Overpayment) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Overpayment) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Overpayment) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPaymentsField

`func (o *Overpayment) HasPaymentsField() bool`

HasPaymentsField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *Overpayment) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Overpayment) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Overpayment) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *Overpayment) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetAttachments

`func (o *Overpayment) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Overpayment) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Overpayment) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachmentsField

`func (o *Overpayment) HasAttachmentsField() bool`

HasAttachmentsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


