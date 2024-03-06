# LineItemItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | User defined item code (max length &#x3D; 30) | [optional] 
**Name** | Pointer to **string** | The name of the item (max length &#x3D; 50) | [optional] 
**ItemID** | Pointer to **string** | The Xero identifier for an Item | [optional] 

## Methods

### NewLineItemItem

`func NewLineItemItem() *LineItemItem`

NewLineItemItem instantiates a new LineItemItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemItemWithDefaults

`func NewLineItemItemWithDefaults() *LineItemItem`

NewLineItemItemWithDefaults instantiates a new LineItemItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LineItemItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LineItemItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LineItemItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCodeField

`func (o *LineItemItem) HasCodeField() bool`

HasCodeField returns a boolean if a field has been set.

### GetName

`func (o *LineItemItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LineItemItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LineItemItem) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *LineItemItem) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetItemID

`func (o *LineItemItem) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *LineItemItem) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *LineItemItem) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemIDField

`func (o *LineItemItem) HasItemIDField() bool`

HasItemIDField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


