# AddressForOrganisation

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

### NewAddressForOrganisation

`func NewAddressForOrganisation() *AddressForOrganisation`

NewAddressForOrganisation instantiates a new AddressForOrganisation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressForOrganisationWithDefaults

`func NewAddressForOrganisationWithDefaults() *AddressForOrganisation`

NewAddressForOrganisationWithDefaults instantiates a new AddressForOrganisation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressType

`func (o *AddressForOrganisation) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *AddressForOrganisation) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *AddressForOrganisation) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressTypeField

`func (o *AddressForOrganisation) HasAddressTypeField() bool`

HasAddressTypeField returns a boolean if a field has been set.

### GetAddressLine1

`func (o *AddressForOrganisation) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressForOrganisation) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressForOrganisation) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1Field

`func (o *AddressForOrganisation) HasAddressLine1Field() bool`

HasAddressLine1Field returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AddressForOrganisation) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressForOrganisation) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressForOrganisation) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2Field

`func (o *AddressForOrganisation) HasAddressLine2Field() bool`

HasAddressLine2Field returns a boolean if a field has been set.

### GetAddressLine3

`func (o *AddressForOrganisation) GetAddressLine3() string`

GetAddressLine3 returns the AddressLine3 field if non-nil, zero value otherwise.

### GetAddressLine3Ok

`func (o *AddressForOrganisation) GetAddressLine3Ok() (*string, bool)`

GetAddressLine3Ok returns a tuple with the AddressLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine3

`func (o *AddressForOrganisation) SetAddressLine3(v string)`

SetAddressLine3 sets AddressLine3 field to given value.

### HasAddressLine3Field

`func (o *AddressForOrganisation) HasAddressLine3Field() bool`

HasAddressLine3Field returns a boolean if a field has been set.

### GetAddressLine4

`func (o *AddressForOrganisation) GetAddressLine4() string`

GetAddressLine4 returns the AddressLine4 field if non-nil, zero value otherwise.

### GetAddressLine4Ok

`func (o *AddressForOrganisation) GetAddressLine4Ok() (*string, bool)`

GetAddressLine4Ok returns a tuple with the AddressLine4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine4

`func (o *AddressForOrganisation) SetAddressLine4(v string)`

SetAddressLine4 sets AddressLine4 field to given value.

### HasAddressLine4Field

`func (o *AddressForOrganisation) HasAddressLine4Field() bool`

HasAddressLine4Field returns a boolean if a field has been set.

### GetCity

`func (o *AddressForOrganisation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressForOrganisation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressForOrganisation) SetCity(v string)`

SetCity sets City field to given value.

### HasCityField

`func (o *AddressForOrganisation) HasCityField() bool`

HasCityField returns a boolean if a field has been set.

### GetRegion

`func (o *AddressForOrganisation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddressForOrganisation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddressForOrganisation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegionField

`func (o *AddressForOrganisation) HasRegionField() bool`

HasRegionField returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressForOrganisation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressForOrganisation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressForOrganisation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCodeField

`func (o *AddressForOrganisation) HasPostalCodeField() bool`

HasPostalCodeField returns a boolean if a field has been set.

### GetCountry

`func (o *AddressForOrganisation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressForOrganisation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressForOrganisation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountryField

`func (o *AddressForOrganisation) HasCountryField() bool`

HasCountryField returns a boolean if a field has been set.

### GetAttentionTo

`func (o *AddressForOrganisation) GetAttentionTo() string`

GetAttentionTo returns the AttentionTo field if non-nil, zero value otherwise.

### GetAttentionToOk

`func (o *AddressForOrganisation) GetAttentionToOk() (*string, bool)`

GetAttentionToOk returns a tuple with the AttentionTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttentionTo

`func (o *AddressForOrganisation) SetAttentionTo(v string)`

SetAttentionTo sets AttentionTo field to given value.

### HasAttentionToField

`func (o *AddressForOrganisation) HasAttentionToField() bool`

HasAttentionToField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


