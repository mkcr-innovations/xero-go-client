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
	"fmt"
)

// QuoteStatusCodes The status of the quote.
type QuoteStatusCodes string

// List of QuoteStatusCodes
const (
	DRAFT_QuoteStatusCodes QuoteStatusCodes = "DRAFT"
	SENT_QuoteStatusCodes QuoteStatusCodes = "SENT"
	DECLINED_QuoteStatusCodes QuoteStatusCodes = "DECLINED"
	ACCEPTED_QuoteStatusCodes QuoteStatusCodes = "ACCEPTED"
	INVOICED_QuoteStatusCodes QuoteStatusCodes = "INVOICED"
	DELETED_QuoteStatusCodes QuoteStatusCodes = "DELETED"
)

// All allowed values of QuoteStatusCodes enum
var AllowedQuoteStatusCodesEnumValues = []QuoteStatusCodes{
	"DRAFT",
	"SENT",
	"DECLINED",
	"ACCEPTED",
	"INVOICED",
	"DELETED",
}

func (v *QuoteStatusCodes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QuoteStatusCodes(value)
	for _, existing := range AllowedQuoteStatusCodesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QuoteStatusCodes", value)
}

// NewQuoteStatusCodesFromValue returns a pointer to a valid QuoteStatusCodes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQuoteStatusCodesFromValue(v string) (*QuoteStatusCodes, error) {
	ev := QuoteStatusCodes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QuoteStatusCodes: valid values are %v", v, AllowedQuoteStatusCodesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuoteStatusCodes) IsValid() bool {
	for _, existing := range AllowedQuoteStatusCodesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuoteStatusCodes value
func (v QuoteStatusCodes) Ptr() *QuoteStatusCodes {
	return &v
}

type NullableQuoteStatusCodes struct {
	value *QuoteStatusCodes
	isSet bool
}

func (v NullableQuoteStatusCodes) Get() *QuoteStatusCodes {
	return v.value
}

func (v *NullableQuoteStatusCodes) Set(val *QuoteStatusCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteStatusCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteStatusCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteStatusCodes(val *QuoteStatusCodes) *NullableQuoteStatusCodes {
	return &NullableQuoteStatusCodes{value: val, isSet: true}
}

func (v NullableQuoteStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteStatusCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

