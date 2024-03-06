/*
Xero Accounting API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.1
Contact: api@xero.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ContactGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactGroup{}

// ContactGroup struct for ContactGroup
type ContactGroup struct {
	// The Name of the contact group. Required when creating a new contact  group
	Name *string `json:"Name,omitempty"`
	// The Status of a contact group. To delete a contact group update the status to DELETED. Only contact groups with a status of ACTIVE are returned on GETs.
	Status *string `json:"Status,omitempty"`
	// The Xero identifier for an contact group – specified as a string following the endpoint name. e.g. /297c2dc5-cc47-4afd-8ec8-74990b8761e9
	ContactGroupID *string `json:"ContactGroupID,omitempty"`
	// The ContactID and Name of Contacts in a contact group. Returned on GETs when the ContactGroupID is supplied in the URL.
	Contacts []Contact `json:"Contacts,omitempty"`
}

// NewContactGroup instantiates a new ContactGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactGroup() *ContactGroup {
	this := ContactGroup{}
	return &this
}

// NewContactGroupWithDefaults instantiates a new ContactGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactGroupWithDefaults() *ContactGroup {
	this := ContactGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContactGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasNameField returns a boolean if a field has been set.
func (o *ContactGroup) HasNameField() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContactGroup) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContactGroup) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactGroup) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatusField returns a boolean if a field has been set.
func (o *ContactGroup) HasStatusField() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ContactGroup) SetStatus(v string) {
	o.Status = &v
}

// GetContactGroupID returns the ContactGroupID field value if set, zero value otherwise.
func (o *ContactGroup) GetContactGroupID() string {
	if o == nil || IsNil(o.ContactGroupID) {
		var ret string
		return ret
	}
	return *o.ContactGroupID
}

// GetContactGroupIDOk returns a tuple with the ContactGroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactGroup) GetContactGroupIDOk() (*string, bool) {
	if o == nil || IsNil(o.ContactGroupID) {
		return nil, false
	}
	return o.ContactGroupID, true
}

// HasContactGroupIDField returns a boolean if a field has been set.
func (o *ContactGroup) HasContactGroupIDField() bool {
	if o != nil && !IsNil(o.ContactGroupID) {
		return true
	}

	return false
}

// SetContactGroupID gets a reference to the given string and assigns it to the ContactGroupID field.
func (o *ContactGroup) SetContactGroupID(v string) {
	o.ContactGroupID = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *ContactGroup) GetContacts() []Contact {
	if o == nil || IsNil(o.Contacts) {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactGroup) GetContactsOk() ([]Contact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContactsField returns a boolean if a field has been set.
func (o *ContactGroup) HasContactsField() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *ContactGroup) SetContacts(v []Contact) {
	o.Contacts = v
}

func (o ContactGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.ContactGroupID) {
		toSerialize["ContactGroupID"] = o.ContactGroupID
	}
	if !IsNil(o.Contacts) {
		toSerialize["Contacts"] = o.Contacts
	}
	return toSerialize, nil
}

type NullableContactGroup struct {
	value *ContactGroup
	isSet bool
}

func (v NullableContactGroup) Get() *ContactGroup {
	return v.value
}

func (v *NullableContactGroup) Set(val *ContactGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableContactGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableContactGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactGroup(val *ContactGroup) *NullableContactGroup {
	return &NullableContactGroup{value: val, isSet: true}
}

func (v NullableContactGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

