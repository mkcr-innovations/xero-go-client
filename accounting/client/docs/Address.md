# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType** | Pointer to **string** | define the type of address | [optional] 
**AddressLine1** | Pointer to **string** | max length &#x3D; 500 | [optional] 
**AddressLine2** | Pointer to **string** | max length &#x3D; 500 | [optional] 
**AddressLine3** | Pointer to **string** | max length &#x3D; 500 | [optional] 
**AddressLine4** | Pointer to **string** | max length &#x3D; 500 | [optional] 
**City** | Pointer to **string** | max length &#x3D; 255 | [optional] 
**Region** | Pointer to **string** | max length &#x3D; 255 | [optional] 
**PostalCode** | Pointer to **string** | max length &#x3D; 50 | [optional] 
**Country** | Pointer to **string** | max length &#x3D; 50, [A-Z], [a-z] only | [optional] 
**AttentionTo** | Pointer to **string** | max length &#x3D; 255 | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressType

`func (o *Address) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *Address) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *Address) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressTypeField

`func (o *Address) HasAddressTypeField() bool`

HasAddressTypeField returns a boolean if a field has been set.

### GetAddressLine1

`func (o *Address) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Address) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Address) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1Field

`func (o *Address) HasAddressLine1Field() bool`

HasAddressLine1Field returns a boolean if a field has been set.

### GetAddressLine2

`func (o *Address) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Address) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Address) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2Field

`func (o *Address) HasAddressLine2Field() bool`

HasAddressLine2Field returns a boolean if a field has been set.

### GetAddressLine3

`func (o *Address) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *Address) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *Address) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3Field

`func (o *Address) HasAddressLine3Field() bool`

HasAddressLine3Field returns a boolean if a field has been set.

### GetAddressLine4

`func (o *Address) GetAddressLine4() string`

GetAddressLine4 returns the AddressLine4 field if non-nil, zero value otherwise.

### GetAddressLine4Ok

`func (o *Address) GetAddressLine4Ok() (*string, bool)`

GetAddressLine4Ok returns a tuple with the AddressLine4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine4

`func (o *Address) SetAddressLine4(v string)`

SetAddressLine4 sets AddressLine4 field to given value.

### HasAddressLine4Field

`func (o *Address) HasAddressLine4Field() bool`

HasAddressLine4Field returns a boolean if a field has been set.

### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCityField

`func (o *Address) HasCityField() bool`

HasCityField returns a boolean if a field has been set.

### GetRegion

`func (o *Address) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Address) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Address) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegionField

`func (o *Address) HasRegionField() bool`

HasRegionField returns a boolean if a field has been set.

### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCodeField

`func (o *Address) HasPostalCodeField() bool`

HasPostalCodeField returns a boolean if a field has been set.

### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountryField

`func (o *Address) HasCountryField() bool`

HasCountryField returns a boolean if a field has been set.

### GetAttentionTo

`func (o *Address) GetAttentionTo() string`

GetAttentionTo returns the AttentionTo field if non-nil, zero value otherwise.

### GetAttentionToOk

`func (o *Address) GetAttentionToOk() (*string, bool)`

GetAttentionToOk returns a tuple with the AttentionTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttentionTo

`func (o *Address) SetAttentionTo(v string)`

SetAttentionTo sets AttentionTo field to given value.

### HasAttentionToField

`func (o *Address) HasAttentionToField() bool`

HasAttentionToField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


