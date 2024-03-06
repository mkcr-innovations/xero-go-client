# PurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | See LineItems | [optional] 
**Date** | Pointer to **string** | Date purchase order was issued – YYYY-MM-DD. If the Date element is not specified then it will default to the current date based on the timezone setting of the organisation | [optional] 
**DeliveryDate** | Pointer to **string** | Date the goods are to be delivered – YYYY-MM-DD | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | Unique alpha numeric code identifying purchase order (when missing will auto-generate from your Organisation Invoice Settings) | [optional] 
**Reference** | Pointer to **string** | Additional reference number | [optional] 
**BrandingThemeID** | Pointer to **string** | See BrandingThemes | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**Status** | Pointer to **string** | See Purchase Order Status Codes | [optional] 
**SentToContact** | Pointer to **bool** | Boolean to set whether the purchase order should be marked as “sent”. This can be set only on purchase orders that have been approved or billed | [optional] 
**DeliveryAddress** | Pointer to **string** | The address the goods are to be delivered to | [optional] 
**AttentionTo** | Pointer to **string** | The person that the delivery is going to | [optional] 
**Telephone** | Pointer to **string** | The phone number for the person accepting the delivery | [optional] 
**DeliveryInstructions** | Pointer to **string** | A free text feild for instructions (500 characters max) | [optional] 
**ExpectedArrivalDate** | Pointer to **string** | The date the goods are expected to arrive. | [optional] 
**PurchaseOrderID** | Pointer to **string** | Xero generated unique identifier for purchase order | [optional] 
**CurrencyRate** | Pointer to **float64** | The currency rate for a multicurrency purchase order. If no rate is specified, the XE.com day rate is used. | [optional] 
**SubTotal** | Pointer to **float64** | Total of purchase order excluding taxes | [optional] [readonly] 
**TotalTax** | Pointer to **float64** | Total tax on purchase order | [optional] [readonly] 
**Total** | Pointer to **float64** | Total of Purchase Order tax inclusive (i.e. SubTotal + TotalTax) | [optional] [readonly] 
**TotalDiscount** | Pointer to **float64** | Total of discounts applied on the purchase order line items | [optional] [readonly] 
**HasAttachments** | Pointer to **bool** | boolean to indicate if a purchase order has an attachment | [optional] [readonly] [default to false]
**UpdatedDateUTC** | Pointer to **string** | Last modified date UTC format | [optional] [readonly] 
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 
**Warnings** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of warning messages from the API | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 

## Methods

### NewPurchaseOrder

`func NewPurchaseOrder() *PurchaseOrder`

NewPurchaseOrder instantiates a new PurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderWithDefaults

`func NewPurchaseOrderWithDefaults() *PurchaseOrder`

NewPurchaseOrderWithDefaults instantiates a new PurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *PurchaseOrder) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PurchaseOrder) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PurchaseOrder) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *PurchaseOrder) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetLineItems

`func (o *PurchaseOrder) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *PurchaseOrder) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *PurchaseOrder) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItemsField

`func (o *PurchaseOrder) HasLineItemsField() bool`

HasLineItemsField returns a boolean if a field has been set.

### GetDate

`func (o *PurchaseOrder) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PurchaseOrder) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PurchaseOrder) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *PurchaseOrder) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *PurchaseOrder) GetDeliveryDate() string`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *PurchaseOrder) GetDeliveryDateOk() (*string, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *PurchaseOrder) SetDeliveryDate(v string)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDateField

`func (o *PurchaseOrder) HasDeliveryDateField() bool`

HasDeliveryDateField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *PurchaseOrder) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *PurchaseOrder) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *PurchaseOrder) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *PurchaseOrder) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *PurchaseOrder) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *PurchaseOrder) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *PurchaseOrder) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumberField

`func (o *PurchaseOrder) HasPurchaseOrderNumberField() bool`

HasPurchaseOrderNumberField returns a boolean if a field has been set.

### GetReference

`func (o *PurchaseOrder) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PurchaseOrder) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PurchaseOrder) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *PurchaseOrder) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetBrandingThemeID

`func (o *PurchaseOrder) GetBrandingThemeID() string`

GetBrandingThemeID returns the BrandingThemeID field if non-nil, zero value otherwise.

### GetBrandingThemeIDOk

`func (o *PurchaseOrder) GetBrandingThemeIDOk() (*string, bool)`

GetBrandingThemeIDOk returns a tuple with the BrandingThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingThemeID

`func (o *PurchaseOrder) SetBrandingThemeID(v string)`

SetBrandingThemeID sets BrandingThemeID field to given value.

### HasBrandingThemeIDField

`func (o *PurchaseOrder) HasBrandingThemeIDField() bool`

HasBrandingThemeIDField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PurchaseOrder) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PurchaseOrder) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PurchaseOrder) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *PurchaseOrder) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetStatus

`func (o *PurchaseOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PurchaseOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PurchaseOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *PurchaseOrder) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetSentToContact

`func (o *PurchaseOrder) GetSentToContact() bool`

GetSentToContact returns the SentToContact field if non-nil, zero value otherwise.

### GetSentToContactOk

`func (o *PurchaseOrder) GetSentToContactOk() (*bool, bool)`

GetSentToContactOk returns a tuple with the SentToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToContact

`func (o *PurchaseOrder) SetSentToContact(v bool)`

SetSentToContact sets SentToContact field to given value.

### HasSentToContactField

`func (o *PurchaseOrder) HasSentToContactField() bool`

HasSentToContactField returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *PurchaseOrder) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *PurchaseOrder) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *PurchaseOrder) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddressField

`func (o *PurchaseOrder) HasDeliveryAddressField() bool`

HasDeliveryAddressField returns a boolean if a field has been set.

### GetAttentionTo

`func (o *PurchaseOrder) GetAttentionTo() string`

GetAttentionTo returns the AttentionTo field if non-nil, zero value otherwise.

### GetAttentionToOk

`func (o *PurchaseOrder) GetAttentionToOk() (*string, bool)`

GetAttentionToOk returns a tuple with the AttentionTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttentionTo

`func (o *PurchaseOrder) SetAttentionTo(v string)`

SetAttentionTo sets AttentionTo field to given value.

### HasAttentionToField

`func (o *PurchaseOrder) HasAttentionToField() bool`

HasAttentionToField returns a boolean if a field has been set.

### GetTelephone

`func (o *PurchaseOrder) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *PurchaseOrder) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *PurchaseOrder) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephoneField

`func (o *PurchaseOrder) HasTelephoneField() bool`

HasTelephoneField returns a boolean if a field has been set.

### GetDeliveryInstructions

`func (o *PurchaseOrder) GetDeliveryInstructions() string`

GetDeliveryInstructions returns the DeliveryInstructions field if non-nil, zero value otherwise.

### GetDeliveryInstructionsOk

`func (o *PurchaseOrder) GetDeliveryInstructionsOk() (*string, bool)`

GetDeliveryInstructionsOk returns a tuple with the DeliveryInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryInstructions

`func (o *PurchaseOrder) SetDeliveryInstructions(v string)`

SetDeliveryInstructions sets DeliveryInstructions field to given value.

### HasDeliveryInstructionsField

`func (o *PurchaseOrder) HasDeliveryInstructionsField() bool`

HasDeliveryInstructionsField returns a boolean if a field has been set.

### GetExpectedArrivalDate

`func (o *PurchaseOrder) GetExpectedArrivalDate() string`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *PurchaseOrder) GetExpectedArrivalDateOk() (*string, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *PurchaseOrder) SetExpectedArrivalDate(v string)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.

### HasExpectedArrivalDateField

`func (o *PurchaseOrder) HasExpectedArrivalDateField() bool`

HasExpectedArrivalDateField returns a boolean if a field has been set.

### GetPurchaseOrderID

`func (o *PurchaseOrder) GetPurchaseOrderID() string`

GetPurchaseOrderID returns the PurchaseOrderID field if non-nil, zero value otherwise.

### GetPurchaseOrderIDOk

`func (o *PurchaseOrder) GetPurchaseOrderIDOk() (*string, bool)`

GetPurchaseOrderIDOk returns a tuple with the PurchaseOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderID

`func (o *PurchaseOrder) SetPurchaseOrderID(v string)`

SetPurchaseOrderID sets PurchaseOrderID field to given value.

### HasPurchaseOrderIDField

`func (o *PurchaseOrder) HasPurchaseOrderIDField() bool`

HasPurchaseOrderIDField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *PurchaseOrder) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *PurchaseOrder) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *PurchaseOrder) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *PurchaseOrder) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetSubTotal

`func (o *PurchaseOrder) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *PurchaseOrder) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *PurchaseOrder) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *PurchaseOrder) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *PurchaseOrder) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *PurchaseOrder) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *PurchaseOrder) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *PurchaseOrder) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *PurchaseOrder) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PurchaseOrder) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PurchaseOrder) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *PurchaseOrder) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *PurchaseOrder) GetTotalDiscount() float64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *PurchaseOrder) GetTotalDiscountOk() (*float64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *PurchaseOrder) SetTotalDiscount(v float64)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscountField

`func (o *PurchaseOrder) HasTotalDiscountField() bool`

HasTotalDiscountField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *PurchaseOrder) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *PurchaseOrder) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *PurchaseOrder) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *PurchaseOrder) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *PurchaseOrder) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *PurchaseOrder) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *PurchaseOrder) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *PurchaseOrder) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *PurchaseOrder) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *PurchaseOrder) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *PurchaseOrder) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *PurchaseOrder) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *PurchaseOrder) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PurchaseOrder) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PurchaseOrder) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *PurchaseOrder) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.

### GetWarnings

`func (o *PurchaseOrder) GetWarnings() []ValidationError`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *PurchaseOrder) GetWarningsOk() (*[]ValidationError, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *PurchaseOrder) SetWarnings(v []ValidationError)`

SetWarnings sets Warnings field to given value.

### HasWarningsField

`func (o *PurchaseOrder) HasWarningsField() bool`

HasWarningsField returns a boolean if a field has been set.

### GetAttachments

`func (o *PurchaseOrder) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PurchaseOrder) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PurchaseOrder) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachmentsField

`func (o *PurchaseOrder) HasAttachmentsField() bool`

HasAttachmentsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


