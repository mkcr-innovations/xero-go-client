# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Date of receipt – YYYY-MM-DD | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Reference** | Pointer to **string** | Additional reference number | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**SubTotal** | Pointer to **float64** | Total of receipt excluding taxes | [optional] 
**TotalTax** | Pointer to **float64** | Total tax on receipt | [optional] 
**Total** | Pointer to **float64** | Total of receipt tax inclusive (i.e. SubTotal + TotalTax) | [optional] 
**ReceiptID** | Pointer to **string** | Xero generated unique identifier for receipt | [optional] 
**Status** | Pointer to **string** | Current status of receipt – see status types | [optional] 
**ReceiptNumber** | Pointer to **string** | Xero generated sequence number for receipt in current claim for a given user | [optional] [readonly] 
**UpdatedDateUTC** | Pointer to **string** | Last modified date UTC format | [optional] [readonly] 
**HasAttachments** | Pointer to **bool** | boolean to indicate if a receipt has an attachment | [optional] [readonly] [default to false]
**Url** | Pointer to **string** | URL link to a source document – shown as “Go to [appName]” in the Xero app | [optional] [readonly] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 
**Warnings** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of warning messages from the API | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 

## Methods

### NewReceipt

`func NewReceipt() *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *Receipt) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Receipt) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Receipt) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *Receipt) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetContact

`func (o *Receipt) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Receipt) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Receipt) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *Receipt) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetLineItems

`func (o *Receipt) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Receipt) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Receipt) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItemsField

`func (o *Receipt) HasLineItemsField() bool`

HasLineItemsField returns a boolean if a field has been set.

### GetUser

`func (o *Receipt) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Receipt) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Receipt) SetUser(v User)`

SetUser sets User field to given value.

### HasUserField

`func (o *Receipt) HasUserField() bool`

HasUserField returns a boolean if a field has been set.

### GetReference

`func (o *Receipt) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Receipt) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Receipt) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *Receipt) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *Receipt) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *Receipt) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *Receipt) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *Receipt) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetSubTotal

`func (o *Receipt) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *Receipt) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *Receipt) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *Receipt) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *Receipt) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *Receipt) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *Receipt) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *Receipt) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *Receipt) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Receipt) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Receipt) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *Receipt) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetReceiptID

`func (o *Receipt) GetReceiptID() string`

GetReceiptID returns the ReceiptID field if non-nil, zero value otherwise.

### GetReceiptIDOk

`func (o *Receipt) GetReceiptIDOk() (*string, bool)`

GetReceiptIDOk returns a tuple with the ReceiptID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptID

`func (o *Receipt) SetReceiptID(v string)`

SetReceiptID sets ReceiptID field to given value.

### HasReceiptIDField

`func (o *Receipt) HasReceiptIDField() bool`

HasReceiptIDField returns a boolean if a field has been set.

### GetStatus

`func (o *Receipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Receipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Receipt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *Receipt) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetReceiptNumber

`func (o *Receipt) GetReceiptNumber() string`

GetReceiptNumber returns the ReceiptNumber field if non-nil, zero value otherwise.

### GetReceiptNumberOk

`func (o *Receipt) GetReceiptNumberOk() (*string, bool)`

GetReceiptNumberOk returns a tuple with the ReceiptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptNumber

`func (o *Receipt) SetReceiptNumber(v string)`

SetReceiptNumber sets ReceiptNumber field to given value.

### HasReceiptNumberField

`func (o *Receipt) HasReceiptNumberField() bool`

HasReceiptNumberField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Receipt) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Receipt) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Receipt) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Receipt) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *Receipt) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Receipt) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Receipt) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *Receipt) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetUrl

`func (o *Receipt) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Receipt) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Receipt) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrlField

`func (o *Receipt) HasUrlField() bool`

HasUrlField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Receipt) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Receipt) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Receipt) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Receipt) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.

### GetWarnings

`func (o *Receipt) GetWarnings() []ValidationError`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Receipt) GetWarningsOk() (*[]ValidationError, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Receipt) SetWarnings(v []ValidationError)`

SetWarnings sets Warnings field to given value.

### HasWarningsField

`func (o *Receipt) HasWarningsField() bool`

HasWarningsField returns a boolean if a field has been set.

### GetAttachments

`func (o *Receipt) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Receipt) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Receipt) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachmentsField

`func (o *Receipt) HasAttachmentsField() bool`

HasAttachmentsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


