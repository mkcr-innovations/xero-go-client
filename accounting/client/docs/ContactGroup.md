# ContactGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Name of the contact group. Required when creating a new contact  group | [optional] 
**Status** | Pointer to **string** | The Status of a contact group. To delete a contact group update the status to DELETED. Only contact groups with a status of ACTIVE are returned on GETs. | [optional] 
**ContactGroupID** | Pointer to **string** | The Xero identifier for an contact group – specified as a string following the endpoint name. e.g. /297c2dc5-cc47-4afd-8ec8-74990b8761e9 | [optional] 
**Contacts** | Pointer to [**[]Contact**](Contact.md) | The ContactID and Name of Contacts in a contact group. Returned on GETs when the ContactGroupID is supplied in the URL. | [optional] 

## Methods

### NewContactGroup

`func NewContactGroup() *ContactGroup`

NewContactGroup instantiates a new ContactGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactGroupWithDefaults

`func NewContactGroupWithDefaults() *ContactGroup`

NewContactGroupWithDefaults instantiates a new ContactGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContactGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactGroup) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *ContactGroup) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetStatus

`func (o *ContactGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContactGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContactGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *ContactGroup) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetContactGroupID

`func (o *ContactGroup) GetContactGroupID() string`

GetContactGroupID returns the ContactGroupID field if non-nil, zero value otherwise.

### GetContactGroupIDOk

`func (o *ContactGroup) GetContactGroupIDOk() (*string, bool)`

GetContactGroupIDOk returns a tuple with the ContactGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactGroupID

`func (o *ContactGroup) SetContactGroupID(v string)`

SetContactGroupID sets ContactGroupID field to given value.

### HasContactGroupIDField

`func (o *ContactGroup) HasContactGroupIDField() bool`

HasContactGroupIDField returns a boolean if a field has been set.

### GetContacts

`func (o *ContactGroup) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ContactGroup) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ContactGroup) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContactsField

`func (o *ContactGroup) HasContactsField() bool`

HasContactsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


