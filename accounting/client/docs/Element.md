# Element

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Array of Validation Error message | [optional] 
**BatchPaymentID** | Pointer to **string** | Unique ID for batch payment object with validation error | [optional] 
**BankTransactionID** | Pointer to **string** |  | [optional] 
**CreditNoteID** | Pointer to **string** |  | [optional] 
**ContactID** | Pointer to **string** |  | [optional] 
**InvoiceID** | Pointer to **string** |  | [optional] 
**ItemID** | Pointer to **string** |  | [optional] 
**PurchaseOrderID** | Pointer to **string** |  | [optional] 

## Methods

### NewElement

`func NewElement() *Element`

NewElement instantiates a new Element object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElementWithDefaults

`func NewElementWithDefaults() *Element`

NewElementWithDefaults instantiates a new Element object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationErrors

`func (o *Element) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Element) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Element) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Element) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.

### GetBatchPaymentID

`func (o *Element) GetBatchPaymentID() string`

GetBatchPaymentID returns the BatchPaymentID field if non-nil, zero value otherwise.

### GetBatchPaymentIDOk

`func (o *Element) GetBatchPaymentIDOk() (*string, bool)`

GetBatchPaymentIDOk returns a tuple with the BatchPaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchPaymentID

`func (o *Element) SetBatchPaymentID(v string)`

SetBatchPaymentID sets BatchPaymentID field to given value.

### HasBatchPaymentIDField

`func (o *Element) HasBatchPaymentIDField() bool`

HasBatchPaymentIDField returns a boolean if a field has been set.

### GetBankTransactionID

`func (o *Element) GetBankTransactionID() string`

GetBankTransactionID returns the BankTransactionID field if non-nil, zero value otherwise.

### GetBankTransactionIDOk

`func (o *Element) GetBankTransactionIDOk() (*string, bool)`

GetBankTransactionIDOk returns a tuple with the BankTransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransactionID

`func (o *Element) SetBankTransactionID(v string)`

SetBankTransactionID sets BankTransactionID field to given value.

### HasBankTransactionIDField

`func (o *Element) HasBankTransactionIDField() bool`

HasBankTransactionIDField returns a boolean if a field has been set.

### GetCreditNoteID

`func (o *Element) GetCreditNoteID() string`

GetCreditNoteID returns the CreditNoteID field if non-nil, zero value otherwise.

### GetCreditNoteIDOk

`func (o *Element) GetCreditNoteIDOk() (*string, bool)`

GetCreditNoteIDOk returns a tuple with the CreditNoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteID

`func (o *Element) SetCreditNoteID(v string)`

SetCreditNoteID sets CreditNoteID field to given value.

### HasCreditNoteIDField

`func (o *Element) HasCreditNoteIDField() bool`

HasCreditNoteIDField returns a boolean if a field has been set.

### GetContactID

`func (o *Element) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *Element) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *Element) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactIDField

`func (o *Element) HasContactIDField() bool`

HasContactIDField returns a boolean if a field has been set.

### GetInvoiceID

`func (o *Element) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *Element) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *Element) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceIDField

`func (o *Element) HasInvoiceIDField() bool`

HasInvoiceIDField returns a boolean if a field has been set.

### GetItemID

`func (o *Element) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *Element) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *Element) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemIDField

`func (o *Element) HasItemIDField() bool`

HasItemIDField returns a boolean if a field has been set.

### GetPurchaseOrderID

`func (o *Element) GetPurchaseOrderID() string`

GetPurchaseOrderID returns the PurchaseOrderID field if non-nil, zero value otherwise.

### GetPurchaseOrderIDOk

`func (o *Element) GetPurchaseOrderIDOk() (*string, bool)`

GetPurchaseOrderIDOk returns a tuple with the PurchaseOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderID

`func (o *Element) SetPurchaseOrderID(v string)`

SetPurchaseOrderID sets PurchaseOrderID field to given value.

### HasPurchaseOrderIDField

`func (o *Element) HasPurchaseOrderIDField() bool`

HasPurchaseOrderIDField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


