# TrackingCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingCategoryID** | Pointer to **string** | The Xero identifier for a tracking category e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9 | [optional] 
**TrackingOptionID** | Pointer to **string** | The Xero identifier for a tracking option e.g. dc54c220-0140-495a-b925-3246adc0075f | [optional] 
**Name** | Pointer to **string** | The name of the tracking category e.g. Department, Region (max length &#x3D; 100) | [optional] 
**Option** | Pointer to **string** | The option name of the tracking option e.g. East, West (max length &#x3D; 100) | [optional] 
**Status** | Pointer to **string** | The status of a tracking category | [optional] 
**Options** | Pointer to [**[]TrackingOption**](TrackingOption.md) | See Tracking Options | [optional] 

## Methods

### NewTrackingCategory

`func NewTrackingCategory() *TrackingCategory`

NewTrackingCategory instantiates a new TrackingCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingCategoryWithDefaults

`func NewTrackingCategoryWithDefaults() *TrackingCategory`

NewTrackingCategoryWithDefaults instantiates a new TrackingCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingCategoryID

`func (o *TrackingCategory) GetTrackingCategoryID() string`

GetTrackingCategoryID returns the TrackingCategoryID field if non-nil, zero value otherwise.

### GetTrackingCategoryIDOk

`func (o *TrackingCategory) GetTrackingCategoryIDOk() (*string, bool)`

GetTrackingCategoryIDOk returns a tuple with the TrackingCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategoryID

`func (o *TrackingCategory) SetTrackingCategoryID(v string)`

SetTrackingCategoryID sets TrackingCategoryID field to given value.

### HasTrackingCategoryIDField

`func (o *TrackingCategory) HasTrackingCategoryIDField() bool`

HasTrackingCategoryIDField returns a boolean if a field has been set.

### GetTrackingOptionID

`func (o *TrackingCategory) GetTrackingOptionID() string`

GetTrackingOptionID returns the TrackingOptionID field if non-nil, zero value otherwise.

### GetTrackingOptionIDOk

`func (o *TrackingCategory) GetTrackingOptionIDOk() (*string, bool)`

GetTrackingOptionIDOk returns a tuple with the TrackingOptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingOptionID

`func (o *TrackingCategory) SetTrackingOptionID(v string)`

SetTrackingOptionID sets TrackingOptionID field to given value.

### HasTrackingOptionIDField

`func (o *TrackingCategory) HasTrackingOptionIDField() bool`

HasTrackingOptionIDField returns a boolean if a field has been set.

### GetName

`func (o *TrackingCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackingCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackingCategory) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *TrackingCategory) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetOption

`func (o *TrackingCategory) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *TrackingCategory) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *TrackingCategory) SetOption(v string)`

SetOption sets Option field to given value.

### HasOptionField

`func (o *TrackingCategory) HasOptionField() bool`

HasOptionField returns a boolean if a field has been set.

### GetStatus

`func (o *TrackingCategory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrackingCategory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrackingCategory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *TrackingCategory) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetOptions

`func (o *TrackingCategory) GetOptions() []TrackingOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TrackingCategory) GetOptionsOk() (*[]TrackingOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TrackingCategory) SetOptions(v []TrackingOption)`

SetOptions sets Options field to given value.

### HasOptionsField

`func (o *TrackingCategory) HasOptionsField() bool`

HasOptionsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


