# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorNumber** | Pointer to **int32** | Exception number | [optional] 
**Type** | Pointer to **string** | Exception type | [optional] 
**Message** | Pointer to **string** | Exception message | [optional] 
**Elements** | Pointer to [**[]Element**](Element.md) | Array of Elements of validation Errors | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorNumber

`func (o *Error) GetErrorNumber() int32`

GetErrorNumber returns the ErrorNumber field if non-nil, zero value otherwise.

### GetErrorNumberOk

`func (o *Error) GetErrorNumberOk() (*int32, bool)`

GetErrorNumberOk returns a tuple with the ErrorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNumber

`func (o *Error) SetErrorNumber(v int32)`

SetErrorNumber sets ErrorNumber field to given value.

### HasErrorNumberField

`func (o *Error) HasErrorNumberField() bool`

HasErrorNumberField returns a boolean if a field has been set.

### GetType

`func (o *Error) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Error) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Error) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *Error) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessageField

`func (o *Error) HasMessageField() bool`

HasMessageField returns a boolean if a field has been set.

### GetElements

`func (o *Error) GetElements() []Element`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *Error) GetElementsOk() (*[]Element, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *Error) SetElements(v []Element)`

SetElements sets Elements field to given value.

### HasElementsField

`func (o *Error) HasElementsField() bool`

HasElementsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


