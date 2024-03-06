# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneType** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** | max length &#x3D; 50 | [optional] 
**PhoneAreaCode** | Pointer to **string** | max length &#x3D; 10 | [optional] 
**PhoneCountryCode** | Pointer to **string** | max length &#x3D; 20 | [optional] 

## Methods

### NewPhone

`func NewPhone() *Phone`

NewPhone instantiates a new Phone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneWithDefaults

`func NewPhoneWithDefaults() *Phone`

NewPhoneWithDefaults instantiates a new Phone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneType

`func (o *Phone) GetPhoneType() string`

GetPhoneType returns the PhoneType field if non-nil, zero value otherwise.

### GetPhoneTypeOk

`func (o *Phone) GetPhoneTypeOk() (*string, bool)`

GetPhoneTypeOk returns a tuple with the PhoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneType

`func (o *Phone) SetPhoneType(v string)`

SetPhoneType sets PhoneType field to given value.

### HasPhoneTypeField

`func (o *Phone) HasPhoneTypeField() bool`

HasPhoneTypeField returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Phone) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Phone) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Phone) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumberField

`func (o *Phone) HasPhoneNumberField() bool`

HasPhoneNumberField returns a boolean if a field has been set.

### GetPhoneAreaCode

`func (o *Phone) GetPhoneAreaCode() string`

GetPhoneAreaCode returns the PhoneAreaCode field if non-nil, zero value otherwise.

### GetPhoneAreaCodeOk

`func (o *Phone) GetPhoneAreaCodeOk() (*string, bool)`

GetPhoneAreaCodeOk returns a tuple with the PhoneAreaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneAreaCode

`func (o *Phone) SetPhoneAreaCode(v string)`

SetPhoneAreaCode sets PhoneAreaCode field to given value.

### HasPhoneAreaCodeField

`func (o *Phone) HasPhoneAreaCodeField() bool`

HasPhoneAreaCodeField returns a boolean if a field has been set.

### GetPhoneCountryCode

`func (o *Phone) GetPhoneCountryCode() string`

GetPhoneCountryCode returns the PhoneCountryCode field if non-nil, zero value otherwise.

### GetPhoneCountryCodeOk

`func (o *Phone) GetPhoneCountryCodeOk() (*string, bool)`

GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountryCode

`func (o *Phone) SetPhoneCountryCode(v string)`

SetPhoneCountryCode sets PhoneCountryCode field to given value.

### HasPhoneCountryCodeField

`func (o *Phone) HasPhoneCountryCodeField() bool`

HasPhoneCountryCodeField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


