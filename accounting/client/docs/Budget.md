# Budget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetID** | Pointer to **string** | Xero identifier | [optional] 
**Type** | Pointer to **string** | Type of Budget. OVERALL or TRACKING | [optional] 
**Description** | Pointer to **string** | The Budget description | [optional] 
**UpdatedDateUTC** | Pointer to **string** | UTC timestamp of last update to budget | [optional] [readonly] 
**BudgetLines** | Pointer to [**[]BudgetLine**](BudgetLine.md) |  | [optional] 
**Tracking** | Pointer to [**[]TrackingCategory**](TrackingCategory.md) |  | [optional] 

## Methods

### NewBudget

`func NewBudget() *Budget`

NewBudget instantiates a new Budget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetWithDefaults

`func NewBudgetWithDefaults() *Budget`

NewBudgetWithDefaults instantiates a new Budget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetID

`func (o *Budget) GetBudgetID() string`

GetBudgetID returns the BudgetID field if non-nil, zero value otherwise.

### GetBudgetIDOk

`func (o *Budget) GetBudgetIDOk() (*string, bool)`

GetBudgetIDOk returns a tuple with the BudgetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetID

`func (o *Budget) SetBudgetID(v string)`

SetBudgetID sets BudgetID field to given value.

### HasBudgetIDField

`func (o *Budget) HasBudgetIDField() bool`

HasBudgetIDField returns a boolean if a field has been set.

### GetType

`func (o *Budget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Budget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Budget) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *Budget) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetDescription

`func (o *Budget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Budget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Budget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescriptionField

`func (o *Budget) HasDescriptionField() bool`

HasDescriptionField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Budget) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Budget) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Budget) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Budget) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetBudgetLines

`func (o *Budget) GetBudgetLines() []BudgetLine`

GetBudgetLines returns the BudgetLines field if non-nil, zero value otherwise.

### GetBudgetLinesOk

`func (o *Budget) GetBudgetLinesOk() (*[]BudgetLine, bool)`

GetBudgetLinesOk returns a tuple with the BudgetLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLines

`func (o *Budget) SetBudgetLines(v []BudgetLine)`

SetBudgetLines sets BudgetLines field to given value.

### HasBudgetLinesField

`func (o *Budget) HasBudgetLinesField() bool`

HasBudgetLinesField returns a boolean if a field has been set.

### GetTracking

`func (o *Budget) GetTracking() []TrackingCategory`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *Budget) GetTrackingOk() (*[]TrackingCategory, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *Budget) SetTracking(v []TrackingCategory)`

SetTracking sets Tracking field to given value.

### HasTrackingField

`func (o *Budget) HasTrackingField() bool`

HasTrackingField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


