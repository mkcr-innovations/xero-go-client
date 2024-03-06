# BrandingTheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandingThemeID** | Pointer to **string** | Xero identifier | [optional] 
**Name** | Pointer to **string** | Name of branding theme | [optional] 
**LogoUrl** | Pointer to **string** | The location of the image file used as the logo on this branding theme | [optional] 
**Type** | Pointer to **string** | Always INVOICE | [optional] 
**SortOrder** | Pointer to **int32** | Integer – ranked order of branding theme. The default branding theme has a value of 0 | [optional] 
**CreatedDateUTC** | Pointer to **string** | UTC timestamp of creation date of branding theme | [optional] [readonly] 

## Methods

### NewBrandingTheme

`func NewBrandingTheme() *BrandingTheme`

NewBrandingTheme instantiates a new BrandingTheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingThemeWithDefaults

`func NewBrandingThemeWithDefaults() *BrandingTheme`

NewBrandingThemeWithDefaults instantiates a new BrandingTheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandingThemeID

`func (o *BrandingTheme) GetBrandingThemeID() string`

GetBrandingThemeID returns the BrandingThemeID field if non-nil, zero value otherwise.

### GetBrandingThemeIDOk

`func (o *BrandingTheme) GetBrandingThemeIDOk() (*string, bool)`

GetBrandingThemeIDOk returns a tuple with the BrandingThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingThemeID

`func (o *BrandingTheme) SetBrandingThemeID(v string)`

SetBrandingThemeID sets BrandingThemeID field to given value.

### HasBrandingThemeIDField

`func (o *BrandingTheme) HasBrandingThemeIDField() bool`

HasBrandingThemeIDField returns a boolean if a field has been set.

### GetName

`func (o *BrandingTheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrandingTheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrandingTheme) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *BrandingTheme) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetLogoUrl

`func (o *BrandingTheme) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *BrandingTheme) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *BrandingTheme) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrlField

`func (o *BrandingTheme) HasLogoUrlField() bool`

HasLogoUrlField returns a boolean if a field has been set.

### GetType

`func (o *BrandingTheme) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BrandingTheme) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BrandingTheme) SetType(v string)`

SetType sets Type field to given value.

### HasTypeField

`func (o *BrandingTheme) HasTypeField() bool`

HasTypeField returns a boolean if a field has been set.

### GetSortOrder

`func (o *BrandingTheme) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *BrandingTheme) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *BrandingTheme) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrderField

`func (o *BrandingTheme) HasSortOrderField() bool`

HasSortOrderField returns a boolean if a field has been set.

### GetCreatedDateUTC

`func (o *BrandingTheme) GetCreatedDateUTC() string`

GetCreatedDateUTC returns the CreatedDateUTC field if non-nil, zero value otherwise.

### GetCreatedDateUTCOk

`func (o *BrandingTheme) GetCreatedDateUTCOk() (*string, bool)`

GetCreatedDateUTCOk returns a tuple with the CreatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateUTC

`func (o *BrandingTheme) SetCreatedDateUTC(v string)`

SetCreatedDateUTC sets CreatedDateUTC field to given value.

### HasCreatedDateUTCField

`func (o *BrandingTheme) HasCreatedDateUTCField() bool`

HasCreatedDateUTCField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


