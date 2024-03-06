# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteID** | Pointer to **string** | QuoteID GUID is automatically generated and is returned after create or GET. | [optional] 
**QuoteNumber** | Pointer to **string** | Unique alpha numeric code identifying a quote (Max Length &#x3D; 255) | [optional] 
**Reference** | Pointer to **string** | Additional reference number | [optional] 
**Terms** | Pointer to **string** | Terms of the quote | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | See LineItems | [optional] 
**Date** | Pointer to **string** | Date quote was issued – YYYY-MM-DD. If the Date element is not specified it will default to the current date based on the timezone setting of the organisation | [optional] 
**DateString** | Pointer to **string** | Date the quote was issued (YYYY-MM-DD) | [optional] 
**ExpiryDate** | Pointer to **string** | Date the quote expires – YYYY-MM-DD. | [optional] 
**ExpiryDateString** | Pointer to **string** | Date the quote expires – YYYY-MM-DD. | [optional] 
**Status** | Pointer to [**QuoteStatusCodes**](QuoteStatusCodes.md) |  | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**CurrencyRate** | Pointer to **float64** | The currency rate for a multicurrency quote | [optional] 
**SubTotal** | Pointer to **float64** | Total of quote excluding taxes. | [optional] [readonly] 
**TotalTax** | Pointer to **float64** | Total tax on quote | [optional] [readonly] 
**Total** | Pointer to **float64** | Total of Quote tax inclusive (i.e. SubTotal + TotalTax). This will be ignored if it doesn’t equal the sum of the LineAmounts | [optional] [readonly] 
**TotalDiscount** | Pointer to **float64** | Total of discounts applied on the quote line items | [optional] [readonly] 
**Title** | Pointer to **string** | Title text for the quote | [optional] 
**Summary** | Pointer to **string** | Summary text for the quote | [optional] 
**BrandingThemeID** | Pointer to **string** | See BrandingThemes | [optional] 
**UpdatedDateUTC** | Pointer to **string** | Last modified date UTC format | [optional] [readonly] 
**LineAmountTypes** | Pointer to [**QuoteLineAmountTypes**](QuoteLineAmountTypes.md) |  | [optional] 
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewQuote

`func NewQuote() *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteID

`func (o *Quote) GetQuoteID() string`

GetQuoteID returns the QuoteID field if non-nil, zero value otherwise.

### GetQuoteIDOk

`func (o *Quote) GetQuoteIDOk() (*string, bool)`

GetQuoteIDOk returns a tuple with the QuoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteID

`func (o *Quote) SetQuoteID(v string)`

SetQuoteID sets QuoteID field to given value.

### HasQuoteIDField

`func (o *Quote) HasQuoteIDField() bool`

HasQuoteIDField returns a boolean if a field has been set.

### GetQuoteNumber

`func (o *Quote) GetQuoteNumber() string`

GetQuoteNumber returns the QuoteNumber field if non-nil, zero value otherwise.

### GetQuoteNumberOk

`func (o *Quote) GetQuoteNumberOk() (*string, bool)`

GetQuoteNumberOk returns a tuple with the QuoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteNumber

`func (o *Quote) SetQuoteNumber(v string)`

SetQuoteNumber sets QuoteNumber field to given value.

### HasQuoteNumberField

`func (o *Quote) HasQuoteNumberField() bool`

HasQuoteNumberField returns a boolean if a field has been set.

### GetReference

`func (o *Quote) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Quote) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Quote) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *Quote) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetTerms

`func (o *Quote) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Quote) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Quote) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTermsField

`func (o *Quote) HasTermsField() bool`

HasTermsField returns a boolean if a field has been set.

### GetContact

`func (o *Quote) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Quote) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Quote) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContactField

`func (o *Quote) HasContactField() bool`

HasContactField returns a boolean if a field has been set.

### GetLineItems

`func (o *Quote) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Quote) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Quote) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItemsField

`func (o *Quote) HasLineItemsField() bool`

HasLineItemsField returns a boolean if a field has been set.

### GetDate

`func (o *Quote) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Quote) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Quote) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *Quote) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetDateString

`func (o *Quote) GetDateString() string`

GetDateString returns the DateString field if non-nil, zero value otherwise.

### GetDateStringOk

`func (o *Quote) GetDateStringOk() (*string, bool)`

GetDateStringOk returns a tuple with the DateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateString

`func (o *Quote) SetDateString(v string)`

SetDateString sets DateString field to given value.

### HasDateStringField

`func (o *Quote) HasDateStringField() bool`

HasDateStringField returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Quote) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Quote) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Quote) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDateField

`func (o *Quote) HasExpiryDateField() bool`

HasExpiryDateField returns a boolean if a field has been set.

### GetExpiryDateString

`func (o *Quote) GetExpiryDateString() string`

GetExpiryDateString returns the ExpiryDateString field if non-nil, zero value otherwise.

### GetExpiryDateStringOk

`func (o *Quote) GetExpiryDateStringOk() (*string, bool)`

GetExpiryDateStringOk returns a tuple with the ExpiryDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateString

`func (o *Quote) SetExpiryDateString(v string)`

SetExpiryDateString sets ExpiryDateString field to given value.

### HasExpiryDateStringField

`func (o *Quote) HasExpiryDateStringField() bool`

HasExpiryDateStringField returns a boolean if a field has been set.

### GetStatus

`func (o *Quote) GetStatus() QuoteStatusCodes`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Quote) GetStatusOk() (*QuoteStatusCodes, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Quote) SetStatus(v QuoteStatusCodes)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *Quote) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Quote) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Quote) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Quote) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCodeField

`func (o *Quote) HasCurrencyCodeField() bool`

HasCurrencyCodeField returns a boolean if a field has been set.

### GetCurrencyRate

`func (o *Quote) GetCurrencyRate() float64`

GetCurrencyRate returns the CurrencyRate field if non-nil, zero value otherwise.

### GetCurrencyRateOk

`func (o *Quote) GetCurrencyRateOk() (*float64, bool)`

GetCurrencyRateOk returns a tuple with the CurrencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyRate

`func (o *Quote) SetCurrencyRate(v float64)`

SetCurrencyRate sets CurrencyRate field to given value.

### HasCurrencyRateField

`func (o *Quote) HasCurrencyRateField() bool`

HasCurrencyRateField returns a boolean if a field has been set.

### GetSubTotal

`func (o *Quote) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *Quote) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *Quote) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotalField

`func (o *Quote) HasSubTotalField() bool`

HasSubTotalField returns a boolean if a field has been set.

### GetTotalTax

`func (o *Quote) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *Quote) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *Quote) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTaxField

`func (o *Quote) HasTotalTaxField() bool`

HasTotalTaxField returns a boolean if a field has been set.

### GetTotal

`func (o *Quote) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Quote) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Quote) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotalField

`func (o *Quote) HasTotalField() bool`

HasTotalField returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *Quote) GetTotalDiscount() float64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *Quote) GetTotalDiscountOk() (*float64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *Quote) SetTotalDiscount(v float64)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscountField

`func (o *Quote) HasTotalDiscountField() bool`

HasTotalDiscountField returns a boolean if a field has been set.

### GetTitle

`func (o *Quote) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Quote) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Quote) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitleField

`func (o *Quote) HasTitleField() bool`

HasTitleField returns a boolean if a field has been set.

### GetSummary

`func (o *Quote) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Quote) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Quote) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummaryField

`func (o *Quote) HasSummaryField() bool`

HasSummaryField returns a boolean if a field has been set.

### GetBrandingThemeID

`func (o *Quote) GetBrandingThemeID() string`

GetBrandingThemeID returns the BrandingThemeID field if non-nil, zero value otherwise.

### GetBrandingThemeIDOk

`func (o *Quote) GetBrandingThemeIDOk() (*string, bool)`

GetBrandingThemeIDOk returns a tuple with the BrandingThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingThemeID

`func (o *Quote) SetBrandingThemeID(v string)`

SetBrandingThemeID sets BrandingThemeID field to given value.

### HasBrandingThemeIDField

`func (o *Quote) HasBrandingThemeIDField() bool`

HasBrandingThemeIDField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Quote) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Quote) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Quote) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Quote) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *Quote) GetLineAmountTypes() QuoteLineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *Quote) GetLineAmountTypesOk() (*QuoteLineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *Quote) SetLineAmountTypes(v QuoteLineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *Quote) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *Quote) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *Quote) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *Quote) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *Quote) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Quote) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Quote) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Quote) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Quote) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


