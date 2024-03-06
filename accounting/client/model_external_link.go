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

// checks if the ExternalLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalLink{}

// ExternalLink struct for ExternalLink
type ExternalLink struct {
	// See External link types
	LinkType *string `json:"LinkType,omitempty"`
	// URL for service e.g. http://twitter.com/xeroapi
	Url *string `json:"Url,omitempty"`
	Description *string `json:"Description,omitempty"`
}

// NewExternalLink instantiates a new ExternalLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalLink() *ExternalLink {
	this := ExternalLink{}
	return &this
}

// NewExternalLinkWithDefaults instantiates a new ExternalLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalLinkWithDefaults() *ExternalLink {
	this := ExternalLink{}
	return &this
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *ExternalLink) GetLinkType() string {
	if o == nil || IsNil(o.LinkType) {
		var ret string
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetLinkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LinkType) {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkTypeField returns a boolean if a field has been set.
func (o *ExternalLink) HasLinkTypeField() bool {
	if o != nil && !IsNil(o.LinkType) {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given string and assigns it to the LinkType field.
func (o *ExternalLink) SetLinkType(v string) {
	o.LinkType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ExternalLink) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrlField returns a boolean if a field has been set.
func (o *ExternalLink) HasUrlField() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ExternalLink) SetUrl(v string) {
	o.Url = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExternalLink) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalLink) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescriptionField returns a boolean if a field has been set.
func (o *ExternalLink) HasDescriptionField() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExternalLink) SetDescription(v string) {
	o.Description = &v
}

func (o ExternalLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkType) {
		toSerialize["LinkType"] = o.LinkType
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	return toSerialize, nil
}

type NullableExternalLink struct {
	value *ExternalLink
	isSet bool
}

func (v NullableExternalLink) Get() *ExternalLink {
	return v.value
}

func (v *NullableExternalLink) Set(val *ExternalLink) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalLink) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalLink(val *ExternalLink) *NullableExternalLink {
	return &NullableExternalLink{value: val, isSet: true}
}

func (v NullableExternalLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


