# RequestEmpty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Need at least one field to create an empty JSON payload | [optional] 

## Methods

### NewRequestEmpty

`func NewRequestEmpty() *RequestEmpty`

NewRequestEmpty instantiates a new RequestEmpty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestEmptyWithDefaults

`func NewRequestEmptyWithDefaults() *RequestEmpty`

NewRequestEmptyWithDefaults instantiates a new RequestEmpty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestEmpty) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestEmpty) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestEmpty) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *RequestEmpty) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


