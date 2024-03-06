# TrackingOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingOptionID** | Pointer to **string** | The Xero identifier for a tracking option e.g. ae777a87-5ef3-4fa0-a4f0-d10e1f13073a | [optional] 
**Name** | Pointer to **string** | The name of the tracking option e.g. Marketing, East (max length &#x3D; 100) | [optional] 
**Status** | Pointer to **string** | The status of a tracking option | [optional] 
**TrackingCategoryID** | Pointer to **string** | Filter by a tracking category e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9 | [optional] 

## Methods

### NewTrackingOption

`func NewTrackingOption() *TrackingOption`

NewTrackingOption instantiates a new TrackingOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingOptionWithDefaults

`func NewTrackingOptionWithDefaults() *TrackingOption`

NewTrackingOptionWithDefaults instantiates a new TrackingOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingOptionID

`func (o *TrackingOption) GetTrackingOptionID() string`

GetTrackingOptionID returns the TrackingOptionID field if non-nil, zero value otherwise.

### GetTrackingOptionIDOk

`func (o *TrackingOption) GetTrackingOptionIDOk() (*string, bool)`

GetTrackingOptionIDOk returns a tuple with the TrackingOptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingOptionID

`func (o *TrackingOption) SetTrackingOptionID(v string)`

SetTrackingOptionID sets TrackingOptionID field to given value.

### HasTrackingOptionIDField

`func (o *TrackingOption) HasTrackingOptionIDField() bool`

HasTrackingOptionIDField returns a boolean if a field has been set.

### GetName

`func (o *TrackingOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackingOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackingOption) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *TrackingOption) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetStatus

`func (o *TrackingOption) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrackingOption) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrackingOption) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *TrackingOption) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetTrackingCategoryID

`func (o *TrackingOption) GetTrackingCategoryID() string`

GetTrackingCategoryID returns the TrackingCategoryID field if non-nil, zero value otherwise.

### GetTrackingCategoryIDOk

`func (o *TrackingOption) GetTrackingCategoryIDOk() (*string, bool)`

GetTrackingCategoryIDOk returns a tuple with the TrackingCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategoryID

`func (o *TrackingOption) SetTrackingCategoryID(v string)`

SetTrackingCategoryID sets TrackingCategoryID field to given value.

### HasTrackingCategoryIDField

`func (o *TrackingOption) HasTrackingCategoryIDField() bool`

HasTrackingCategoryIDField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


