# Employee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeID** | Pointer to **string** | The Xero identifier for an employee e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9 | [optional] 
**Status** | Pointer to **string** | Current status of an employee – see contact status types | [optional] 
**FirstName** | Pointer to **string** | First name of an employee (max length &#x3D; 255) | [optional] 
**LastName** | Pointer to **string** | Last name of an employee (max length &#x3D; 255) | [optional] 
**ExternalLink** | Pointer to [**ExternalLink**](ExternalLink.md) |  | [optional] 
**UpdatedDateUTC** | Pointer to **string** |  | [optional] [readonly] 
**StatusAttributeString** | Pointer to **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewEmployee

`func NewEmployee() *Employee`

NewEmployee instantiates a new Employee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeWithDefaults

`func NewEmployeeWithDefaults() *Employee`

NewEmployeeWithDefaults instantiates a new Employee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeID

`func (o *Employee) GetEmployeeID() string`

GetEmployeeID returns the EmployeeID field if non-nil, zero value otherwise.

### GetEmployeeIDOk

`func (o *Employee) GetEmployeeIDOk() (*string, bool)`

GetEmployeeIDOk returns a tuple with the EmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeID

`func (o *Employee) SetEmployeeID(v string)`

SetEmployeeID sets EmployeeID field to given value.

### HasEmployeeIDField

`func (o *Employee) HasEmployeeIDField() bool`

HasEmployeeIDField returns a boolean if a field has been set.

### GetStatus

`func (o *Employee) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Employee) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Employee) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatusField

`func (o *Employee) HasStatusField() bool`

HasStatusField returns a boolean if a field has been set.

### GetFirstName

`func (o *Employee) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Employee) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Employee) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstNameField

`func (o *Employee) HasFirstNameField() bool`

HasFirstNameField returns a boolean if a field has been set.

### GetLastName

`func (o *Employee) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Employee) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Employee) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastNameField

`func (o *Employee) HasLastNameField() bool`

HasLastNameField returns a boolean if a field has been set.

### GetExternalLink

`func (o *Employee) GetExternalLink() ExternalLink`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *Employee) GetExternalLinkOk() (*ExternalLink, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *Employee) SetExternalLink(v ExternalLink)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLinkField

`func (o *Employee) HasExternalLinkField() bool`

HasExternalLinkField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Employee) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Employee) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Employee) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Employee) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *Employee) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *Employee) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *Employee) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *Employee) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Employee) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Employee) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Employee) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Employee) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


