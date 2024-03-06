# ExternalLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkType** | Pointer to **string** | See External link types | [optional] 
**Url** | Pointer to **string** | URL for service e.g. http://twitter.com/xeroapi | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalLink

`func NewExternalLink() *ExternalLink`

NewExternalLink instantiates a new ExternalLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalLinkWithDefaults

`func NewExternalLinkWithDefaults() *ExternalLink`

NewExternalLinkWithDefaults instantiates a new ExternalLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkType

`func (o *ExternalLink) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *ExternalLink) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *ExternalLink) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkTypeField

`func (o *ExternalLink) HasLinkTypeField() bool`

HasLinkTypeField returns a boolean if a field has been set.

### GetUrl

`func (o *ExternalLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExternalLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExternalLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrlField

`func (o *ExternalLink) HasUrlField() bool`

HasUrlField returns a boolean if a field has been set.

### GetDescription

`func (o *ExternalLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalLink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescriptionField

`func (o *ExternalLink) HasDescriptionField() bool`

HasDescriptionField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


