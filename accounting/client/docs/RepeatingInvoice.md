# RepeatingInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | See Invoice Types | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Schedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | See LineItems | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**Reference** | Pointer to **string** | ACCREC only – additional reference number | [optional] 
**BrandingThemeID** | Pointer to **string** | See BrandingThemes | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**Status** | Pointer to **string** | One of the following - DRAFT or AUTHORISED – See Invoice Status Codes | [optional] 
**SubTotal** | Pointer to **float64** | Total of invoice excluding taxes | [optional] 
**TotalTax** | Pointer to **float64** | Total tax on invoice | [optional] 
**Total** | Pointer to **float64** | Total of Invoice tax inclusive (i.e. SubTotal + TotalTax) | [optional] 
**RepeatingInvoiceID** | Pointer to **string** | Xero generated unique identifier for repeating invoice template | [optional] 
**ID** | Pointer to **string** | Xero generated unique identifier for repeating invoice template | [optional] 
**HasAttachments** | Pointer to **bool** | Boolean to indicate if an invoice has an attachment | [optional] [readonly] [default to false]
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 
**ApprovedForSending** | Pointer to **bool** | Boolean to indicate whether the invoice has been approved for sending | [optional] [default to false]
**SendCopy** | Pointer to **bool** | Boolean to indicate whether a copy is sent to sender&#39;s email | [optional] [default to false]
**MarkAsSent** | Pointer to **bool** | Boolean to indicate whether the invoice in the Xero app displays as \&quot;sent\&quot; | [optional] [default to false]
**IncludePDF** | Pointer to **bool** | Boolean to indicate whether to include PDF attachment | [optional] [default to false]

## Methods

### NewRepeatingInvoice

`func NewRepeatingInvoice() *RepeatingInvoice`

NewRepeatingInvoice instantiates a new RepeatingInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepeatingInvoiceWithDefaults

`func NewRepeatingInvoiceWithDefaults() *RepeatingInvoice`

NewRepeatingInvoiceWithDefaults instantiates a new RepeatingInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RepeatingInvoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepeatingInvoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepeatingInvoice) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *RepeatingInvoice) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetContact

`func (o *RepeatingInvoice) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *RepeatingInvoice) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *RepeatingInvoice) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *RepeatingInvoice) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetSchedule

`func (o *RepeatingInvoice) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RepeatingInvoice) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RepeatingInvoice) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.

### HasScheduleField

`func (o *RepeatingInvoice) HasScheduleField() bool`

HasScheduleField returns a boolean if a field has been set.

### GetLineItems

`func (o *RepeatingInvoice) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *RepeatingInvoice) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *RepeatingInvoice) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItemsField

`func (o *RepeatingInvoice) HasLineItemsField() bool`

HasLineItemsField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *RepeatingInvoice) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *RepeatingInvoice) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *RepeatingInvoice) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *RepeatingInvoice) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetReference

`func (o *RepeatingInvoice) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RepeatingInvoice) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RepeatingInvoice) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *RepeatingInvoice) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetBrandingThemeID

`func (o *RepeatingInvoice) GetBrandingThemeID() string`

GetBrandingThemeID returns the BrandingThemeID field if non-nil, zero value otherwise.

### GetBrandingThemeIDOk

`func (o *RepeatingInvoice) GetBrandingThemeIDOk() (*string, bool)`

GetBrandingThemeIDOk returns a tuple with the BrandingThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingThemeID

`func (o *RepeatingInvoice) SetBrandingThemeID(v string)`

SetBrandingThemeID sets BrandingThemeID field to given value.

### HasBrandingThemeIDField

`func (o *RepeatingInvoice) HasBrandingThemeIDField() bool`

HasBrandingThemeIDField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *RepeatingInvoice) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RepeatingInvoice) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RepeatingInvoice) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *RepeatingInvoice) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetStatus

`func (o *RepeatingInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RepeatingInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RepeatingInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *RepeatingInvoice) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetSubTotal

`func (o *RepeatingInvoice) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *RepeatingInvoice) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *RepeatingInvoice) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *RepeatingInvoice) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *RepeatingInvoice) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *RepeatingInvoice) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *RepeatingInvoice) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *RepeatingInvoice) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *RepeatingInvoice) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RepeatingInvoice) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RepeatingInvoice) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *RepeatingInvoice) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetRepeatingInvoiceID

`func (o *RepeatingInvoice) GetRepeatingInvoiceID() string`

GetRepeatingInvoiceID returns the RepeatingInvoiceID field if non-nil, zero value otherwise.

### GetRepeatingInvoiceIDOk

`func (o *RepeatingInvoice) GetRepeatingInvoiceIDOk() (*string, bool)`

GetRepeatingInvoiceIDOk returns a tuple with the RepeatingInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatingInvoiceID

`func (o *RepeatingInvoice) SetRepeatingInvoiceID(v string)`

SetRepeatingInvoiceID sets RepeatingInvoiceID field to given value.

### HasRepeatingInvoiceIDField

`func (o *RepeatingInvoice) HasRepeatingInvoiceIDField() bool`

HasRepeatingInvoiceIDField returns a boolean if a field has been set.

### GetID

`func (o *RepeatingInvoice) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RepeatingInvoice) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RepeatingInvoice) SetID(v string)`

SetID sets ID field to given value.

### HasIDField

`func (o *RepeatingInvoice) HasIDField() bool`

HasIDField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *RepeatingInvoice) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *RepeatingInvoice) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *RepeatingInvoice) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *RepeatingInvoice) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetAttachments

`func (o *RepeatingInvoice) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *RepeatingInvoice) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *RepeatingInvoice) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachmentsField

`func (o *RepeatingInvoice) HasAttachmentsField() bool`

HasAttachmentsField returns a boolean if a field has been set.

### GetApprovedForSending

`func (o *RepeatingInvoice) GetApprovedForSending() bool`

GetApprovedForSending returns the ApprovedForSending field if non-nil, zero value otherwise.

### GetApprovedForSendingOk

`func (o *RepeatingInvoice) GetApprovedForSendingOk() (*bool, bool)`

GetApprovedForSendingOk returns a tuple with the ApprovedForSending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedForSending

`func (o *RepeatingInvoice) SetApprovedForSending(v bool)`

SetApprovedForSending sets ApprovedForSending field to given value.

### HasApprovedForSendingField

`func (o *RepeatingInvoice) HasApprovedForSendingField() bool`

HasApprovedForSendingField returns a boolean if a field has been set.

### GetSendCopy

`func (o *RepeatingInvoice) GetSendCopy() bool`

GetSendCopy returns the SendCopy field if non-nil, zero value otherwise.

### GetSendCopyOk

`func (o *RepeatingInvoice) GetSendCopyOk() (*bool, bool)`

GetSendCopyOk returns a tuple with the SendCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCopy

`func (o *RepeatingInvoice) SetSendCopy(v bool)`

SetSendCopy sets SendCopy field to given value.

### HasSendCopyField

`func (o *RepeatingInvoice) HasSendCopyField() bool`

HasSendCopyField returns a boolean if a field has been set.

### GetMarkAsSent

`func (o *RepeatingInvoice) GetMarkAsSent() bool`

GetMarkAsSent returns the MarkAsSent field if non-nil, zero value otherwise.

### GetMarkAsSentOk

`func (o *RepeatingInvoice) GetMarkAsSentOk() (*bool, bool)`

GetMarkAsSentOk returns a tuple with the MarkAsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsSent

`func (o *RepeatingInvoice) SetMarkAsSent(v bool)`

SetMarkAsSent sets MarkAsSent field to given value.

### HasMarkAsSentField

`func (o *RepeatingInvoice) HasMarkAsSentField() bool`

HasMarkAsSentField returns a boolean if a field has been set.

### GetIncludePDF

`func (o *RepeatingInvoice) GetIncludePDF() bool`

GetIncludePDF returns the IncludePDF field if non-nil, zero value otherwise.

### GetIncludePDFOk

`func (o *RepeatingInvoice) GetIncludePDFOk() (*bool, bool)`

GetIncludePDFOk returns a tuple with the IncludePDF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePDF

`func (o *RepeatingInvoice) SetIncludePDF(v bool)`

SetIncludePDF sets IncludePDF field to given value.

### HasIncludePDFField

`func (o *RepeatingInvoice) HasIncludePDFField() bool`

HasIncludePDFField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


