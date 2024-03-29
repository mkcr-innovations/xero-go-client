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

// checks if the PaymentService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentService{}

// PaymentService struct for PaymentService
type PaymentService struct {
	// Xero identifier
	PaymentServiceID *string `json:"PaymentServiceID,omitempty"`
	// Name of payment service
	PaymentServiceName *string `json:"PaymentServiceName,omitempty"`
	// The custom payment URL
	PaymentServiceUrl *string `json:"PaymentServiceUrl,omitempty"`
	// The text displayed on the Pay Now button in Xero Online Invoicing. If this is not set it will default to Pay by credit card
	PayNowText *string `json:"PayNowText,omitempty"`
	// This will always be CUSTOM for payment services created via the API.
	PaymentServiceType *string `json:"PaymentServiceType,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
}

// NewPaymentService instantiates a new PaymentService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentService() *PaymentService {
	this := PaymentService{}
	return &this
}

// NewPaymentServiceWithDefaults instantiates a new PaymentService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceWithDefaults() *PaymentService {
	this := PaymentService{}
	return &this
}

// GetPaymentServiceID returns the PaymentServiceID field value if set, zero value otherwise.
func (o *PaymentService) GetPaymentServiceID() string {
	if o == nil || IsNil(o.PaymentServiceID) {
		var ret string
		return ret
	}
	return *o.PaymentServiceID
}

// GetPaymentServiceIDOk returns a tuple with the PaymentServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentService) GetPaymentServiceIDOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceID) {
		return nil, false
	}
	return o.PaymentServiceID, true
}

// HasPaymentServiceIDField returns a boolean if a field has been set.
func (o *PaymentService) HasPaymentServiceIDField() bool {
	if o != nil && !IsNil(o.PaymentServiceID) {
		return true
	}

	return false
}

// SetPaymentServiceID gets a reference to the given string and assigns it to the PaymentServiceID field.
func (o *PaymentService) SetPaymentServiceID(v string) {
	o.PaymentServiceID = &v
}

// GetPaymentServiceName returns the PaymentServiceName field value if set, zero value otherwise.
func (o *PaymentService) GetPaymentServiceName() string {
	if o == nil || IsNil(o.PaymentServiceName) {
		var ret string
		return ret
	}
	return *o.PaymentServiceName
}

// GetPaymentServiceNameOk returns a tuple with the PaymentServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentService) GetPaymentServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceName) {
		return nil, false
	}
	return o.PaymentServiceName, true
}

// HasPaymentServiceNameField returns a boolean if a field has been set.
func (o *PaymentService) HasPaymentServiceNameField() bool {
	if o != nil && !IsNil(o.PaymentServiceName) {
		return true
	}

	return false
}

// SetPaymentServiceName gets a reference to the given string and assigns it to the PaymentServiceName field.
func (o *PaymentService) SetPaymentServiceName(v string) {
	o.PaymentServiceName = &v
}

// GetPaymentServiceUrl returns the PaymentServiceUrl field value if set, zero value otherwise.
func (o *PaymentService) GetPaymentServiceUrl() string {
	if o == nil || IsNil(o.PaymentServiceUrl) {
		var ret string
		return ret
	}
	return *o.PaymentServiceUrl
}

// GetPaymentServiceUrlOk returns a tuple with the PaymentServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentService) GetPaymentServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceUrl) {
		return nil, false
	}
	return o.PaymentServiceUrl, true
}

// HasPaymentServiceUrlField returns a boolean if a field has been set.
func (o *PaymentService) HasPaymentServiceUrlField() bool {
	if o != nil && !IsNil(o.PaymentServiceUrl) {
		return true
	}

	return false
}

// SetPaymentServiceUrl gets a reference to the given string and assigns it to the PaymentServiceUrl field.
func (o *PaymentService) SetPaymentServiceUrl(v string) {
	o.PaymentServiceUrl = &v
}

// GetPayNowText returns the PayNowText field value if set, zero value otherwise.
func (o *PaymentService) GetPayNowText() string {
	if o == nil || IsNil(o.PayNowText) {
		var ret string
		return ret
	}
	return *o.PayNowText
}

// GetPayNowTextOk returns a tuple with the PayNowText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentService) GetPayNowTextOk() (*string, bool) {
	if o == nil || IsNil(o.PayNowText) {
		return nil, false
	}
	return o.PayNowText, true
}

// HasPayNowTextField returns a boolean if a field has been set.
func (o *PaymentService) HasPayNowTextField() bool {
	if o != nil && !IsNil(o.PayNowText) {
		return true
	}

	return false
}

// SetPayNowText gets a reference to the given string and assigns it to the PayNowText field.
func (o *PaymentService) SetPayNowText(v string) {
	o.PayNowText = &v
}

// GetPaymentServiceType returns the PaymentServiceType field value if set, zero value otherwise.
func (o *PaymentService) GetPaymentServiceType() string {
	if o == nil || IsNil(o.PaymentServiceType) {
		var ret string
		return ret
	}
	return *o.PaymentServiceType
}

// GetPaymentServiceTypeOk returns a tuple with the PaymentServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentService) GetPaymentServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceType) {
		return nil, false
	}
	return o.PaymentServiceType, true
}

// HasPaymentServiceTypeField returns a boolean if a field has been set.
func (o *PaymentService) HasPaymentServiceTypeField() bool {
	if o != nil && !IsNil(o.PaymentServiceType) {
		return true
	}

	return false
}

// SetPaymentServiceType gets a reference to the given string and assigns it to the PaymentServiceType field.
func (o *PaymentService) SetPaymentServiceType(v string) {
	o.PaymentServiceType = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *PaymentService) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentService) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrorsField returns a boolean if a field has been set.
func (o *PaymentService) HasValidationErrorsField() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *PaymentService) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

func (o PaymentService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentServiceID) {
		toSerialize["PaymentServiceID"] = o.PaymentServiceID
	}
	if !IsNil(o.PaymentServiceName) {
		toSerialize["PaymentServiceName"] = o.PaymentServiceName
	}
	if !IsNil(o.PaymentServiceUrl) {
		toSerialize["PaymentServiceUrl"] = o.PaymentServiceUrl
	}
	if !IsNil(o.PayNowText) {
		toSerialize["PayNowText"] = o.PayNowText
	}
	if !IsNil(o.PaymentServiceType) {
		toSerialize["PaymentServiceType"] = o.PaymentServiceType
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullablePaymentService struct {
	value *PaymentService
	isSet bool
}

func (v NullablePaymentService) Get() *PaymentService {
	return v.value
}

func (v *NullablePaymentService) Set(val *PaymentService) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentService) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentService(val *PaymentService) *NullablePaymentService {
	return &NullablePaymentService{value: val, isSet: true}
}

func (v NullablePaymentService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


