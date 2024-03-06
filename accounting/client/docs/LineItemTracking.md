# LineItemTracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingCategoryID** | Pointer to **string** | The Xero identifier for a tracking category | [optional] 
**TrackingOptionID** | Pointer to **string** | The Xero identifier for a tracking category option | [optional] 
**Name** | Pointer to **string** | The name of the tracking category | [optional] 
**Option** | Pointer to **string** | See Tracking Options | [optional] 

## Methods

### NewLineItemTracking

`func NewLineItemTracking() *LineItemTracking`

NewLineItemTracking instantiates a new LineItemTracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemTrackingWithDefaults

`func NewLineItemTrackingWithDefaults() *LineItemTracking`

NewLineItemTrackingWithDefaults instantiates a new LineItemTracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingCategoryID

`func (o *LineItemTracking) GetTrackingCategoryID() string`

GetTrackingCategoryID returns the TrackingCategoryID field if non-nil, zero value otherwise.

### GetTrackingCategoryIDOk

`func (o *LineItemTracking) GetTrackingCategoryIDOk() (*string, bool)`

GetTrackingCategoryIDOk returns a tuple with the TrackingCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategoryID

`func (o *LineItemTracking) SetTrackingCategoryID(v string)`

SetTrackingCategoryID sets TrackingCategoryID field to given value.

### HasTrackingCategoryIDField

`func (o *LineItemTracking) HasTrackingCategoryIDField() bool`

HasTrackingCategoryIDField returns a boolean if a field has been set.

### GetTrackingOptionID

`func (o *LineItemTracking) GetTrackingOptionID() string`

GetTrackingOptionID returns the TrackingOptionID field if non-nil, zero value otherwise.

### GetTrackingOptionIDOk

`func (o *LineItemTracking) GetTrackingOptionIDOk() (*string, bool)`

GetTrackingOptionIDOk returns a tuple with the TrackingOptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingOptionID

`func (o *LineItemTracking) SetTrackingOptionID(v string)`

SetTrackingOptionID sets TrackingOptionID field to given value.

### HasTrackingOptionIDField

`func (o *LineItemTracking) HasTrackingOptionIDField() bool`

HasTrackingOptionIDField returns a boolean if a field has been set.

### GetName

`func (o *LineItemTracking) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LineItemTracking) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LineItemTracking) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *LineItemTracking) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetOption

`func (o *LineItemTracking) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *LineItemTracking) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *LineItemTracking) SetOption(v string)`

SetOption sets Option field to given value.

### HasOptionField

`func (o *LineItemTracking) HasOptionField() bool`

HasOptionField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


