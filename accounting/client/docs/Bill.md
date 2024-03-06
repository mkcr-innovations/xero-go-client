# Bill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **int32** | Day of Month (0-31) | [optional] 
**Type** | Pointer to [**PaymentTermType**](PaymentTermType.md) |  | [optional] 

## Methods

### NewBill

`func NewBill() *Bill`

NewBill instantiates a new Bill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillWithDefaults

`func NewBillWithDefaults() *Bill`

NewBillWithDefaults instantiates a new Bill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *Bill) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Bill) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Bill) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDayField

`func (o *Bill) HasDayField() bool`

HasDayField returns a boolean if a field has been set.

### GetType

`func (o *Bill) GetType() PaymentTermType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Bill) GetTypeOk() (*PaymentTermType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Bill) SetType(v PaymentTermType)`

SetType sets Type field to given value.

### HasTypeField

`func (o *Bill) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


