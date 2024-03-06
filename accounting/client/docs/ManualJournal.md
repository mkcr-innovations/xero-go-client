# ManualJournal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Narration** | **string** | Description of journal being posted | 
**JournalLines** | Pointer to [**[]ManualJournalLine**](ManualJournalLine.md) | See JournalLines | [optional] 
**Date** | Pointer to **string** | Date journal was posted – YYYY-MM-DD | [optional] 
**LineAmountTypes** | Pointer to [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**Status** | Pointer to **string** | See Manual Journal Status Codes | [optional] 
**Url** | Pointer to **string** | Url link to a source document – shown as “Go to [appName]” in the Xero app | [optional] 
**ShowOnCashBasisReports** | Pointer to **bool** | Boolean – default is true if not specified | [optional] 
**HasAttachments** | Pointer to **bool** | Boolean to indicate if a manual journal has an attachment | [optional] [readonly] [default to false]
**UpdatedDateUTC** | Pointer to **string** | Last modified date UTC format | [optional] [readonly] 
**ManualJournalID** | Pointer to **string** | The Xero identifier for a Manual Journal | [optional] 
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**Warnings** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of warning messages from the API | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 

## Methods

### NewManualJournal

`func NewManualJournal(narration string, ) *ManualJournal`

NewManualJournal instantiates a new ManualJournal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalWithDefaults

`func NewManualJournalWithDefaults() *ManualJournal`

NewManualJournalWithDefaults instantiates a new ManualJournal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNarration

`func (o *ManualJournal) GetNarration() string`

GetNarration returns the Narration field if non-nil, zero value otherwise.

### GetNarrationOk

`func (o *ManualJournal) GetNarrationOk() (*string, bool)`

GetNarrationOk returns a tuple with the Narration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarration

`func (o *ManualJournal) SetNarration(v string)`

SetNarration sets Narration field to given value.


### GetJournalLines

`func (o *ManualJournal) GetJournalLines() []ManualJournalLine`

GetJournalLines returns the JournalLines field if non-nil, zero value otherwise.

### GetJournalLinesOk

`func (o *ManualJournal) GetJournalLinesOk() (*[]ManualJournalLine, bool)`

GetJournalLinesOk returns a tuple with the JournalLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalLines

`func (o *ManualJournal) SetJournalLines(v []ManualJournalLine)`

SetJournalLines sets JournalLines field to given value.

### HasJournalLinesField

`func (o *ManualJournal) HasJournalLinesField() bool`

HasJournalLinesField returns a boolean if a field has been set.

### GetDate

`func (o *ManualJournal) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ManualJournal) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ManualJournal) SetDate(v string)`

SetDate sets Date field to given value.

### HasDateField

`func (o *ManualJournal) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetLineAmountTypes

`func (o *ManualJournal) GetLineAmountTypes() LineAmountTypes`

GetLineAmountTypes returns the LineAmountTypes field if non-nil, zero value otherwise.

### GetLineAmountTypesOk

`func (o *ManualJournal) GetLineAmountTypesOk() (*LineAmountTypes, bool)`

GetLineAmountTypesOk returns a tuple with the LineAmountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmountTypes

`func (o *ManualJournal) SetLineAmountTypes(v LineAmountTypes)`

SetLineAmountTypes sets LineAmountTypes field to given value.

### HasLineAmountTypesField

`func (o *ManualJournal) HasLineAmountTypesField() bool`

HasLineAmountTypesField returns a boolean if a field has been set.

### GetStatus

`func (o *ManualJournal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManualJournal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManualJournal) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *ManualJournal) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetUrl

`func (o *ManualJournal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ManualJournal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ManualJournal) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrlField

`func (o *ManualJournal) HasUrlField() bool`

HasUrlField returns a boolean if a field has been set.

### GetShowOnCashBasisReports

`func (o *ManualJournal) GetShowOnCashBasisReports() bool`

GetShowOnCashBasisReports returns the ShowOnCashBasisReports field if non-nil, zero value otherwise.

### GetShowOnCashBasisReportsOk

`func (o *ManualJournal) GetShowOnCashBasisReportsOk() (*bool, bool)`

GetShowOnCashBasisReportsOk returns a tuple with the ShowOnCashBasisReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCashBasisReports

`func (o *ManualJournal) SetShowOnCashBasisReports(v bool)`

SetShowOnCashBasisReports sets ShowOnCashBasisReports field to given value.

### HasShowOnCashBasisReportsField

`func (o *ManualJournal) HasShowOnCashBasisReportsField() bool`

HasShowOnCashBasisReportsField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *ManualJournal) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *ManualJournal) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *ManualJournal) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *ManualJournal) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *ManualJournal) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *ManualJournal) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *ManualJournal) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *ManualJournal) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetManualJournalID

`func (o *ManualJournal) GetManualJournalID() string`

GetManualJournalID returns the ManualJournalID field if non-nil, zero value otherwise.

### GetManualJournalIDOk

`func (o *ManualJournal) GetManualJournalIDOk() (*string, bool)`

GetManualJournalIDOk returns a tuple with the ManualJournalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualJournalID

`func (o *ManualJournal) SetManualJournalID(v string)`

SetManualJournalID sets ManualJournalID field to given value.

### HasManualJournalIDField

`func (o *ManualJournal) HasManualJournalIDField() bool`

HasManualJournalIDField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *ManualJournal) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *ManualJournal) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *ManualJournal) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *ManualJournal) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetWarnings

`func (o *ManualJournal) GetWarnings() []ValidationError`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ManualJournal) GetWarningsOk() (*[]ValidationError, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ManualJournal) SetWarnings(v []ValidationError)`

SetWarnings sets Warnings field to given value.

### HasWarningsField

`func (o *ManualJournal) HasWarningsField() bool`

HasWarningsField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ManualJournal) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ManualJournal) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ManualJournal) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *ManualJournal) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.

### GetAttachments

`func (o *ManualJournal) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ManualJournal) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ManualJournal) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachmentsField

`func (o *ManualJournal) HasAttachmentsField() bool`

HasAttachmentsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


