# LineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItemID** | Pointer to **string** | LineItem unique ID | [optional] 
**Description** | Pointer to **string** | Description needs to be at least 1 char long. A line item with just a description (i.e no unit amount or quantity) can be created by specifying just a &lt;Description&gt; element that contains at least 1 character | [optional] 
**Quantity** | Pointer to **float64** | LineItem Quantity | [optional] 
**UnitAmount** | Pointer to **float64** | LineItem Unit Amount | [optional] 
**ItemCode** | Pointer to **string** | See Items | [optional] 
**AccountCode** | Pointer to **string** | See Accounts | [optional] 
**AccountID** | Pointer to **string** | The associated account ID related to this line item | [optional] 
**TaxType** | Pointer to **string** | The tax type from TaxRates | [optional] 
**TaxAmount** | Pointer to **float64** | The tax amount is auto calculated as a percentage of the line amount (see below) based on the tax rate. This value can be overriden if the calculated &lt;TaxAmount&gt; is not correct. | [optional] 
**Item** | Pointer to [**LineItemItem**](LineItemItem.md) |  | [optional] 
**LineAmount** | Pointer to **float64** | If you wish to omit either the Quantity or UnitAmount you can provide a LineAmount and Xero will calculate the missing amount for you. The line amount reflects the discounted price if either a DiscountRate or DiscountAmount has been used i.e. LineAmount &#x3D; Quantity * Unit Amount * ((100 - DiscountRate)/100) or LineAmount &#x3D; (Quantity * UnitAmount) - DiscountAmount | [optional] 
**Tracking** | Pointer to [**[]LineItemTracking**](LineItemTracking.md) | Optional Tracking Category – see Tracking.  Any LineItem can have a  maximum of 2 &lt;TrackingCategory&gt; elements. | [optional] 
**DiscountRate** | Pointer to **float64** | Percentage discount being applied to a line item (only supported on  ACCREC invoices – ACC PAY invoices and credit notes in Xero do not support discounts | [optional] 
**DiscountAmount** | Pointer to **float64** | Discount amount being applied to a line item. Only supported on ACCREC invoices and quotes. ACCPAY invoices and credit notes in Xero do not support discounts. | [optional] 
**RepeatingInvoiceID** | Pointer to **string** | The Xero identifier for a Repeating Invoice | [optional] 

## Methods

### NewLineItem

`func NewLineItem() *LineItem`

NewLineItem instantiates a new LineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineItemID

`func (o *LineItem) GetLineItemID() string`

GetLineItemID returns the LineItemID field if non-nil, zero value otherwise.

### GetLineItemIDOk

`func (o *LineItem) GetLineItemIDOk() (*string, bool)`

GetLineItemIDOk returns a tuple with the LineItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemID

`func (o *LineItem) SetLineItemID(v string)`

SetLineItemID sets LineItemID field to given value.

### HasLineItemIDField

`func (o *LineItem) HasLineItemIDField() bool`

HasLineItemIDField returns a boolean if a field has been set.

### GetDescription

`func (o *LineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescriptionField

`func (o *LineItem) HasDescriptionField() bool`

HasDescriptionField returns a boolean if a field has been set.

### GetQuantity

`func (o *LineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantityField

`func (o *LineItem) HasQuantityField() bool`

HasQuantityField returns a boolean if a field has been set.

### GetUnitAmount

`func (o *LineItem) GetUnitAmount() float64`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *LineItem) GetUnitAmountOk() (*float64, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *LineItem) SetUnitAmount(v float64)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmountField

`func (o *LineItem) HasUnitAmountField() bool`

HasUnitAmountField returns a boolean if a field has been set.

### GetItemCode

`func (o *LineItem) GetItemCode() string`

GetItemCode returns the ItemCode field if non-nil, zero value otherwise.

### GetItemCodeOk

`func (o *LineItem) GetItemCodeOk() (*string, bool)`

GetItemCodeOk returns a tuple with the ItemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCode

`func (o *LineItem) SetItemCode(v string)`

SetItemCode sets ItemCode field to given value.

### HasItemCodeField

`func (o *LineItem) HasItemCodeField() bool`

HasItemCodeField returns a boolean if a field has been set.

### GetAccountCode

`func (o *LineItem) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *LineItem) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *LineItem) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.

### HasAccountCodeField

`func (o *LineItem) HasAccountCodeField() bool`

HasAccountCodeField returns a boolean if a field has been set.

### GetAccountID

`func (o *LineItem) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *LineItem) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *LineItem) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountIDField

`func (o *LineItem) HasAccountIDField() bool`

HasAccountIDField returns a boolean if a field has been set.

### GetTaxType

`func (o *LineItem) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *LineItem) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *LineItem) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxTypeField

`func (o *LineItem) HasTaxTypeField() bool`

HasTaxTypeField returns a boolean if a field has been set.

### GetTaxAmount

`func (o *LineItem) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *LineItem) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *LineItem) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmountField

`func (o *LineItem) HasTaxAmountField() bool`

HasTaxAmountField returns a boolean if a field has been set.

### GetItem

`func (o *LineItem) GetItem() LineItemItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *LineItem) GetItemOk() (*LineItemItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *LineItem) SetItem(v LineItemItem)`

SetItem sets Item field to given value.

### HasItemField

`func (o *LineItem) HasItemField() bool`

HasItemField returns a boolean if a field has been set.

### GetLineAmount

`func (o *LineItem) GetLineAmount() float64`

GetLineAmount returns the LineAmount field if non-nil, zero value otherwise.

### GetLineAmountOk

`func (o *LineItem) GetLineAmountOk() (*float64, bool)`

GetLineAmountOk returns a tuple with the LineAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmount

`func (o *LineItem) SetLineAmount(v float64)`

SetLineAmount sets LineAmount field to given value.

### HasLineAmountField

`func (o *LineItem) HasLineAmountField() bool`

HasLineAmountField returns a boolean if a field has been set.

### GetTracking

`func (o *LineItem) GetTracking() []LineItemTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *LineItem) GetTrackingOk() (*[]LineItemTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *LineItem) SetTracking(v []LineItemTracking)`

SetTracking sets Tracking field to given value.

### HasTrackingField

`func (o *LineItem) HasTrackingField() bool`

HasTrackingField returns a boolean if a field has been set.

### GetDiscountRate

`func (o *LineItem) GetDiscountRate() float64`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *LineItem) GetDiscountRateOk() (*float64, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *LineItem) SetDiscountRate(v float64)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRateField

`func (o *LineItem) HasDiscountRateField() bool`

HasDiscountRateField returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *LineItem) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *LineItem) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *LineItem) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmountField

`func (o *LineItem) HasDiscountAmountField() bool`

HasDiscountAmountField returns a boolean if a field has been set.

### GetRepeatingInvoiceID

`func (o *LineItem) GetRepeatingInvoiceID() string`

GetRepeatingInvoiceID returns the RepeatingInvoiceID field if non-nil, zero value otherwise.

### GetRepeatingInvoiceIDOk

`func (o *LineItem) GetRepeatingInvoiceIDOk() (*string, bool)`

GetRepeatingInvoiceIDOk returns a tuple with the RepeatingInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatingInvoiceID

`func (o *LineItem) SetRepeatingInvoiceID(v string)`

SetRepeatingInvoiceID sets RepeatingInvoiceID field to given value.

### HasRepeatingInvoiceIDField

`func (o *LineItem) HasRepeatingInvoiceIDField() bool`

HasRepeatingInvoiceIDField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


