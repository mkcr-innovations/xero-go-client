# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserID** | Pointer to **string** | Xero identifier | [optional] 
**EmailAddress** | Pointer to **string** | Email address of user | [optional] 
**FirstName** | Pointer to **string** | First name of user | [optional] 
**LastName** | Pointer to **string** | Last name of user | [optional] 
**UpdatedDateUTC** | Pointer to **string** | Timestamp of last change to user | [optional] [readonly] 
**IsSubscriber** | Pointer to **bool** | Boolean to indicate if user is the subscriber | [optional] 
**OrganisationRole** | Pointer to **string** | User role that defines permissions in Xero and via API (READONLY, INVOICEONLY, STANDARD, FINANCIALADVISER, etc) | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserID

`func (o *User) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *User) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *User) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserIDField

`func (o *User) HasUserIDField() bool`

HasUserIDField returns a boolean if a field has been set.

### GetEmailAddress

`func (o *User) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *User) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *User) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddressField

`func (o *User) HasEmailAddressField() bool`

HasEmailAddressField returns a boolean if a field has been set.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstNameField

`func (o *User) HasFirstNameField() bool`

HasFirstNameField returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastNameField

`func (o *User) HasLastNameField() bool`

HasLastNameField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *User) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *User) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *User) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *User) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetIsSubscriber

`func (o *User) GetIsSubscriber() bool`

GetIsSubscriber returns the IsSubscriber field if non-nil, zero value otherwise.

### GetIsSubscriberOk

`func (o *User) GetIsSubscriberOk() (*bool, bool)`

GetIsSubscriberOk returns a tuple with the IsSubscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscriber

`func (o *User) SetIsSubscriber(v bool)`

SetIsSubscriber sets IsSubscriber field to given value.

### HasIsSubscriberField

`func (o *User) HasIsSubscriberField() bool`

HasIsSubscriberField returns a boolean if a field has been set.

### GetOrganisationRole

`func (o *User) GetOrganisationRole() string`

GetOrganisationRole returns the OrganisationRole field if non-nil, zero value otherwise.

### GetOrganisationRoleOk

`func (o *User) GetOrganisationRoleOk() (*string, bool)`

GetOrganisationRoleOk returns a tuple with the OrganisationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationRole

`func (o *User) SetOrganisationRole(v string)`

SetOrganisationRole sets OrganisationRole field to given value.

### HasOrganisationRoleField

`func (o *User) HasOrganisationRoleField() bool`

HasOrganisationRoleField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


