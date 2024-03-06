# Journal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JournalID** | Pointer to **string** | Xero identifier | [optional] 
**JournalDate** | Pointer to **string** | Date the journal was posted | [optional] 
**JournalNumber** | Pointer to **int32** | Xero generated journal number | [optional] 
**CreatedDateUTC** | Pointer to **string** | Created date UTC format | [optional] [readonly] 
**Reference** | Pointer to **string** | reference field for additional indetifying information | [optional] 
**SourceID** | Pointer to **string** | The identifier for the source transaction (e.g. InvoiceID) | [optional] 
**SourceType** | Pointer to **string** | The journal source type. The type of transaction that created the journal | [optional] 
**JournalLines** | Pointer to [**[]JournalLine**](JournalLine.md) | See JournalLines | [optional] 

## Methods

### NewJournal

`func NewJournal() *Journal`

NewJournal instantiates a new Journal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalWithDefaults

`func NewJournalWithDefaults() *Journal`

NewJournalWithDefaults instantiates a new Journal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJournalID

`func (o *Journal) GetJournalID() string`

GetJournalID returns the JournalID field if non-nil, zero value otherwise.

### GetJournalIDOk

`func (o *Journal) GetJournalIDOk() (*string, bool)`

GetJournalIDOk returns a tuple with the JournalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalID

`func (o *Journal) SetJournalID(v string)`

SetJournalID sets JournalID field to given value.

### HasJournalIDField

`func (o *Journal) HasJournalIDField() bool`

HasJournalIDField returns a boolean if a field has been set.

### GetJournalDate

`func (o *Journal) GetJournalDate() string`

GetJournalDate returns the JournalDate field if non-nil, zero value otherwise.

### GetJournalDateOk

`func (o *Journal) GetJournalDateOk() (*string, bool)`

GetJournalDateOk returns a tuple with the JournalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalDate

`func (o *Journal) SetJournalDate(v string)`

SetJournalDate sets JournalDate field to given value.

### HasJournalDateField

`func (o *Journal) HasJournalDateField() bool`

HasJournalDateField returns a boolean if a field has been set.

### GetJournalNumber

`func (o *Journal) GetJournalNumber() int32`

GetJournalNumber returns the JournalNumber field if non-nil, zero value otherwise.

### GetJournalNumberOk

`func (o *Journal) GetJournalNumberOk() (*int32, bool)`

GetJournalNumberOk returns a tuple with the JournalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalNumber

`func (o *Journal) SetJournalNumber(v int32)`

SetJournalNumber sets JournalNumber field to given value.

### HasJournalNumberField

`func (o *Journal) HasJournalNumberField() bool`

HasJournalNumberField returns a boolean if a field has been set.

### GetCreatedDateUTC

`func (o *Journal) GetCreatedDateUTC() string`

GetCreatedDateUTC returns the CreatedDateUTC field if non-nil, zero value otherwise.

### GetCreatedDateUTCOk

`func (o *Journal) GetCreatedDateUTCOk() (*string, bool)`

GetCreatedDateUTCOk returns a tuple with the CreatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateUTC

`func (o *Journal) SetCreatedDateUTC(v string)`

SetCreatedDateUTC sets CreatedDateUTC field to given value.

### HasCreatedDateUTCField

`func (o *Journal) HasCreatedDateUTCField() bool`

HasCreatedDateUTCField returns a boolean if a field has been set.

### GetReference

`func (o *Journal) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Journal) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Journal) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReferenceField

`func (o *Journal) HasReferenceField() bool`

HasReferenceField returns a boolean if a field has been set.

### GetSourceID

`func (o *Journal) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *Journal) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *Journal) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceIDField

`func (o *Journal) HasSourceIDField() bool`

HasSourceIDField returns a boolean if a field has been set.

### GetSourceType

`func (o *Journal) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Journal) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Journal) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceTypeField

`func (o *Journal) HasSourceTypeField() bool`

HasSourceTypeField returns a boolean if a field has been set.

### GetJournalLines

`func (o *Journal) GetJournalLines() []JournalLine`

GetJournalLines returns the JournalLines field if non-nil, zero value otherwise.

### GetJournalLinesOk

`func (o *Journal) GetJournalLinesOk() (*[]JournalLine, bool)`

GetJournalLinesOk returns a tuple with the JournalLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalLines

`func (o *Journal) SetJournalLines(v []JournalLine)`

SetJournalLines sets JournalLines field to given value.

### HasJournalLinesField

`func (o *Journal) HasJournalLinesField() bool`

HasJournalLinesField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


