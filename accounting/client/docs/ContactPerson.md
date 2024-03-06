# ContactPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First name of person | [optional] 
**LastName** | Pointer to **string** | Last name of person | [optional] 
**EmailAddress** | Pointer to **string** | Email address of person | [optional] 
**IncludeInEmails** | Pointer to **bool** | boolean to indicate whether contact should be included on emails with invoices etc. | [optional] 

## Methods

### NewContactPerson

`func NewContactPerson() *ContactPerson`

NewContactPerson instantiates a new ContactPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactPersonWithDefaults

`func NewContactPersonWithDefaults() *ContactPerson`

NewContactPersonWithDefaults instantiates a new ContactPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ContactPerson) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactPerson) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactPerson) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstNameField

`func (o *ContactPerson) HasFirstNameField() bool`

HasFirstNameField returns a boolean if a field has been set.

### GetLastName

`func (o *ContactPerson) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactPerson) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactPerson) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastNameField

`func (o *ContactPerson) HasLastNameField() bool`

HasLastNameField returns a boolean if a field has been set.

### GetEmailAddress

`func (o *ContactPerson) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ContactPerson) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ContactPerson) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddressField

`func (o *ContactPerson) HasEmailAddressField() bool`

HasEmailAddressField returns a boolean if a field has been set.

### GetIncludeInEmails

`func (o *ContactPerson) GetIncludeInEmails() bool`

GetIncludeInEmails returns the IncludeInEmails field if non-nil, zero value otherwise.

### GetIncludeInEmailsOk

`func (o *ContactPerson) GetIncludeInEmailsOk() (*bool, bool)`

GetIncludeInEmailsOk returns a tuple with the IncludeInEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInEmails

`func (o *ContactPerson) SetIncludeInEmails(v bool)`

SetIncludeInEmails sets IncludeInEmails field to given value.

### HasIncludeInEmailsField

`func (o *ContactPerson) HasIncludeInEmailsField() bool`

HasIncludeInEmailsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


