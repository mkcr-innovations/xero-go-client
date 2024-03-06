# LinkedTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceTransactionID** | Pointer to **string** | Filter by the SourceTransactionID. Get all the linked transactions created from a particular ACCPAY invoice | [optional] 
**SourceLineItemID** | Pointer to **string** | The line item identifier from the source transaction. | [optional] 
**ContactID** | Pointer to **string** | Filter by the combination of ContactID and Status. Get all the linked transactions that have been assigned to a particular customer and have a particular status e.g. GET /LinkedTransactions?ContactID&#x3D;4bb34b03-3378-4bb2-a0ed-6345abf3224e&amp;Status&#x3D;APPROVED. | [optional] 
**TargetTransactionID** | Pointer to **string** | Filter by the TargetTransactionID. Get all the linked transactions  allocated to a particular ACCREC invoice | [optional] 
**TargetLineItemID** | Pointer to **string** | The line item identifier from the target transaction. It is possible  to link multiple billable expenses to the same TargetLineItemID. | [optional] 
**LinkedTransactionID** | Pointer to **string** | The Xero identifier for an Linked Transaction e.g./LinkedTransactions/297c2dc5-cc47-4afd-8ec8-74990b8761e9 | [optional] 
**Status** | Pointer to **string** | Filter by the combination of ContactID and Status. Get all the linked transactions that have been assigned to a particular customer and have a particular status e.g. GET /LinkedTransactions?ContactID&#x3D;4bb34b03-3378-4bb2-a0ed-6345abf3224e&amp;Status&#x3D;APPROVED. | [optional] 
**Type** | Pointer to **string** | This will always be BILLABLEEXPENSE. More types may be added in future. | [optional] 
**UpdatedDateUTC** | Pointer to **string** | The last modified date in UTC format | [optional] [readonly] 
**SourceTransactionTypeCode** | Pointer to **string** | The Type of the source tranasction. This will be ACCPAY if the linked transaction was created from an invoice and SPEND if it was created from a bank transaction. | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewLinkedTransaction

`func NewLinkedTransaction() *LinkedTransaction`

NewLinkedTransaction instantiates a new LinkedTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedTransactionWithDefaults

`func NewLinkedTransactionWithDefaults() *LinkedTransaction`

NewLinkedTransactionWithDefaults instantiates a new LinkedTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceTransactionID

`func (o *LinkedTransaction) GetSourceTransactionID() string`

GetSourceTransactionID returns the SourceTransactionID field if non-nil, zero value otherwise.

### GetSourceTransactionIDOk

`func (o *LinkedTransaction) GetSourceTransactionIDOk() (*string, bool)`

GetSourceTransactionIDOk returns a tuple with the SourceTransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTransactionID

`func (o *LinkedTransaction) SetSourceTransactionID(v string)`

SetSourceTransactionID sets SourceTransactionID field to given value.

### HasSourceTransactionIDField

`func (o *LinkedTransaction) HasSourceTransactionIDField() bool`

HasSourceTransactionIDField returns a boolean if a field has been set.

### GetSourceLineItemID

`func (o *LinkedTransaction) GetSourceLineItemID() string`

GetSourceLineItemID returns the SourceLineItemID field if non-nil, zero value otherwise.

### GetSourceLineItemIDOk

`func (o *LinkedTransaction) GetSourceLineItemIDOk() (*string, bool)`

GetSourceLineItemIDOk returns a tuple with the SourceLineItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLineItemID

`func (o *LinkedTransaction) SetSourceLineItemID(v string)`

SetSourceLineItemID sets SourceLineItemID field to given value.

### HasSourceLineItemIDField

`func (o *LinkedTransaction) HasSourceLineItemIDField() bool`

HasSourceLineItemIDField returns a boolean if a field has been set.

### GetContactID

`func (o *LinkedTransaction) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *LinkedTransaction) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *LinkedTransaction) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactIDField

`func (o *LinkedTransaction) HasContactIDField() bool`

HasContactIDField returns a boolean if a field has been set.

### GetTargetTransactionID

`func (o *LinkedTransaction) GetTargetTransactionID() string`

GetTargetTransactionID returns the TargetTransactionID field if non-nil, zero value otherwise.

### GetTargetTransactionIDOk

`func (o *LinkedTransaction) GetTargetTransactionIDOk() (*string, bool)`

GetTargetTransactionIDOk returns a tuple with the TargetTransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTransactionID

`func (o *LinkedTransaction) SetTargetTransactionID(v string)`

SetTargetTransactionID sets TargetTransactionID field to given value.

### HasTargetTransactionIDField

`func (o *LinkedTransaction) HasTargetTransactionIDField() bool`

HasTargetTransactionIDField returns a boolean if a field has been set.

### GetTargetLineItemID

`func (o *LinkedTransaction) GetTargetLineItemID() string`

GetTargetLineItemID returns the TargetLineItemID field if non-nil, zero value otherwise.

### GetTargetLineItemIDOk

`func (o *LinkedTransaction) GetTargetLineItemIDOk() (*string, bool)`

GetTargetLineItemIDOk returns a tuple with the TargetLineItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLineItemID

`func (o *LinkedTransaction) SetTargetLineItemID(v string)`

SetTargetLineItemID sets TargetLineItemID field to given value.

### HasTargetLineItemIDField

`func (o *LinkedTransaction) HasTargetLineItemIDField() bool`

HasTargetLineItemIDField returns a boolean if a field has been set.

### GetLinkedTransactionID

`func (o *LinkedTransaction) GetLinkedTransactionID() string`

GetLinkedTransactionID returns the LinkedTransactionID field if non-nil, zero value otherwise.

### GetLinkedTransactionIDOk

`func (o *LinkedTransaction) GetLinkedTransactionIDOk() (*string, bool)`

GetLinkedTransactionIDOk returns a tuple with the LinkedTransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactionID

`func (o *LinkedTransaction) SetLinkedTransactionID(v string)`

SetLinkedTransactionID sets LinkedTransactionID field to given value.

### HasLinkedTransactionIDField

`func (o *LinkedTransaction) HasLinkedTransactionIDField() bool`

HasLinkedTransactionIDField returns a boolean if a field has been set.

### GetStatus

`func (o *LinkedTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkedTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkedTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *LinkedTransaction) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetType

`func (o *LinkedTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *LinkedTransaction) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *LinkedTransaction) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *LinkedTransaction) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *LinkedTransaction) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *LinkedTransaction) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetSourceTransactionTypeCode

`func (o *LinkedTransaction) GetSourceTransactionTypeCode() string`

GetSourceTransactionTypeCode returns the SourceTransactionTypeCode field if non-nil, zero value otherwise.

### GetSourceTransactionTypeCodeOk

`func (o *LinkedTransaction) GetSourceTransactionTypeCodeOk() (*string, bool)`

GetSourceTransactionTypeCodeOk returns a tuple with the SourceTransactionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTransactionTypeCode

`func (o *LinkedTransaction) SetSourceTransactionTypeCode(v string)`

SetSourceTransactionTypeCode sets SourceTransactionTypeCode field to given value.

### HasSourceTransactionTypeCodeField

`func (o *LinkedTransaction) HasSourceTransactionTypeCodeField() bool`

HasSourceTransactionTypeCodeField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *LinkedTransaction) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *LinkedTransaction) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *LinkedTransaction) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *LinkedTransaction) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


