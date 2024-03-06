# PaymentTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bills** | Pointer to [**Bill**](Bill.md) |  | [optional] 
**Sales** | Pointer to [**Bill**](Bill.md) |  | [optional] 

## Methods

### NewPaymentTerm

`func NewPaymentTerm() *PaymentTerm`

NewPaymentTerm instantiates a new PaymentTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTermWithDefaults

`func NewPaymentTermWithDefaults() *PaymentTerm`

NewPaymentTermWithDefaults instantiates a new PaymentTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBills

`func (o *PaymentTerm) GetBills() Bill`

GetBills returns the Bills field if non-nil, zero value otherwise.

### GetBillsOk

`func (o *PaymentTerm) GetBillsOk() (*Bill, bool)`

GetBillsOk returns a tuple with the Bills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBills

`func (o *PaymentTerm) SetBills(v Bill)`

SetBills sets Bills field to given value.

### HasBillsField

`func (o *PaymentTerm) HasBillsField() bool`

HasBillsField returns a boolean if a field has been set.

### GetSales

`func (o *PaymentTerm) GetSales() Bill`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *PaymentTerm) GetSalesOk() (*Bill, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *PaymentTerm) SetSales(v Bill)`

SetSales sets Sales field to given value.

### HasSalesField

`func (o *PaymentTerm) HasSalesField() bool`

HasSalesField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


