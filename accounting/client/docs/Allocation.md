# Allocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationID** | Pointer to **string** | Xero generated unique identifier | [optional] 
**Invoice** | [**Invoice**](Invoice.md) |  | 
**Overpayment** | Pointer to [**Overpayment**](Overpayment.md) |  | [optional] 
**Prepayment** | Pointer to [**Prepayment**](Prepayment.md) |  | [optional] 
**CreditNote** | Pointer to [**CreditNote**](CreditNote.md) |  | [optional] 
**Amount** | **float64** | the amount being applied to the invoice | 
**Date** | **string** | the date the allocation is applied YYYY-MM-DD. | 
**IsDeleted** | Pointer to **bool** | A flag that returns true when the allocation is succesfully deleted | [optional] [readonly] 
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewAllocation

`func NewAllocation(invoice Invoice, amount float64, date string, ) *Allocation`

NewAllocation instantiates a new Allocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationWithDefaults

`func NewAllocationWithDefaults() *Allocation`

NewAllocationWithDefaults instantiates a new Allocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationID

`func (o *Allocation) GetAllocationID() string`

GetAllocationID returns the AllocationID field if non-nil, zero value otherwise.

### GetAllocationIDOk

`func (o *Allocation) GetAllocationIDOk() (*string, bool)`

GetAllocationIDOk returns a tuple with the AllocationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationID

`func (o *Allocation) SetAllocationID(v string)`

SetAllocationID sets AllocationID field to given value.

### HasAllocationIDField

`func (o *Allocation) HasAllocationIDField() bool`

HasAllocationIDField returns a boolean if a field has been set.

### GetInvoice

`func (o *Allocation) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *Allocation) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *Allocation) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.


### GetOverpayment

`func (o *Allocation) GetOverpayment() Overpayment`

GetOverpayment returns the Overpayment field if non-nil, zero value otherwise.

### GetOverpaymentOk

`func (o *Allocation) GetOverpaymentOk() (*Overpayment, bool)`

GetOverpaymentOk returns a tuple with the Overpayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverpayment

`func (o *Allocation) SetOverpayment(v Overpayment)`

SetOverpayment sets Overpayment field to given value.

### HasOverpaymentField

`func (o *Allocation) HasOverpaymentField() bool`

HasOverpaymentField returns a boolean if a field has been set.

### GetPrepayment

`func (o *Allocation) GetPrepayment() Prepayment`

GetPrepayment returns the Prepayment field if non-nil, zero value otherwise.

### GetPrepaymentOk

`func (o *Allocation) GetPrepaymentOk() (*Prepayment, bool)`

GetPrepaymentOk returns a tuple with the Prepayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepayment

`func (o *Allocation) SetPrepayment(v Prepayment)`

SetPrepayment sets Prepayment field to given value.

### HasPrepaymentField

`func (o *Allocation) HasPrepaymentField() bool`

HasPrepaymentField returns a boolean if a field has been set.

### GetCreditNote

`func (o *Allocation) GetCreditNote() CreditNote`

GetCreditNote returns the CreditNote field if non-nil, zero value otherwise.

### GetCreditNoteOk

`func (o *Allocation) GetCreditNoteOk() (*CreditNote, bool)`

GetCreditNoteOk returns a tuple with the CreditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNote

`func (o *Allocation) SetCreditNote(v CreditNote)`

SetCreditNote sets CreditNote field to given value.

### HasCreditNoteField

`func (o *Allocation) HasCreditNoteField() bool`

HasCreditNoteField returns a boolean if a field has been set.

### GetAmount

`func (o *Allocation) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Allocation) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Allocation) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *Allocation) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Allocation) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Allocation) SetDate(v string)`

SetDate sets Date field to given value.


### GetIsDeleted

`func (o *Allocation) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Allocation) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Allocation) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeletedField

`func (o *Allocation) HasIsDeletedField() bool`

HasIsDeletedField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *Allocation) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *Allocation) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *Allocation) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *Allocation) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Allocation) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Allocation) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Allocation) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Allocation) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


