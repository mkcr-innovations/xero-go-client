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

// CurrencyCode 3 letter alpha code for the currency – see list of currency codes
type CurrencyCode string

// List of CurrencyCode
const (
	AED_CurrencyCode CurrencyCode = "AED"
	AFN_CurrencyCode CurrencyCode = "AFN"
	ALL_CurrencyCode CurrencyCode = "ALL"
	AMD_CurrencyCode CurrencyCode = "AMD"
	ANG_CurrencyCode CurrencyCode = "ANG"
	AOA_CurrencyCode CurrencyCode = "AOA"
	ARS_CurrencyCode CurrencyCode = "ARS"
	AUD_CurrencyCode CurrencyCode = "AUD"
	AWG_CurrencyCode CurrencyCode = "AWG"
	AZN_CurrencyCode CurrencyCode = "AZN"
	BAM_CurrencyCode CurrencyCode = "BAM"
	BBD_CurrencyCode CurrencyCode = "BBD"
	BDT_CurrencyCode CurrencyCode = "BDT"
	BGN_CurrencyCode CurrencyCode = "BGN"
	BHD_CurrencyCode CurrencyCode = "BHD"
	BIF_CurrencyCode CurrencyCode = "BIF"
	BMD_CurrencyCode CurrencyCode = "BMD"
	BND_CurrencyCode CurrencyCode = "BND"
	BOB_CurrencyCode CurrencyCode = "BOB"
	BRL_CurrencyCode CurrencyCode = "BRL"
	BSD_CurrencyCode CurrencyCode = "BSD"
	BTN_CurrencyCode CurrencyCode = "BTN"
	BWP_CurrencyCode CurrencyCode = "BWP"
	BYN_CurrencyCode CurrencyCode = "BYN"
	BYR_CurrencyCode CurrencyCode = "BYR"
	BZD_CurrencyCode CurrencyCode = "BZD"
	CAD_CurrencyCode CurrencyCode = "CAD"
	CDF_CurrencyCode CurrencyCode = "CDF"
	CHF_CurrencyCode CurrencyCode = "CHF"
	CLF_CurrencyCode CurrencyCode = "CLF"
	CLP_CurrencyCode CurrencyCode = "CLP"
	CNY_CurrencyCode CurrencyCode = "CNY"
	COP_CurrencyCode CurrencyCode = "COP"
	CRC_CurrencyCode CurrencyCode = "CRC"
	CUC_CurrencyCode CurrencyCode = "CUC"
	CUP_CurrencyCode CurrencyCode = "CUP"
	CVE_CurrencyCode CurrencyCode = "CVE"
	CZK_CurrencyCode CurrencyCode = "CZK"
	DJF_CurrencyCode CurrencyCode = "DJF"
	DKK_CurrencyCode CurrencyCode = "DKK"
	DOP_CurrencyCode CurrencyCode = "DOP"
	DZD_CurrencyCode CurrencyCode = "DZD"
	EEK_CurrencyCode CurrencyCode = "EEK"
	EGP_CurrencyCode CurrencyCode = "EGP"
	ERN_CurrencyCode CurrencyCode = "ERN"
	ETB_CurrencyCode CurrencyCode = "ETB"
	EUR_CurrencyCode CurrencyCode = "EUR"
	FJD_CurrencyCode CurrencyCode = "FJD"
	FKP_CurrencyCode CurrencyCode = "FKP"
	GBP_CurrencyCode CurrencyCode = "GBP"
	GEL_CurrencyCode CurrencyCode = "GEL"
	GHS_CurrencyCode CurrencyCode = "GHS"
	GIP_CurrencyCode CurrencyCode = "GIP"
	GMD_CurrencyCode CurrencyCode = "GMD"
	GNF_CurrencyCode CurrencyCode = "GNF"
	GTQ_CurrencyCode CurrencyCode = "GTQ"
	GYD_CurrencyCode CurrencyCode = "GYD"
	HKD_CurrencyCode CurrencyCode = "HKD"
	HNL_CurrencyCode CurrencyCode = "HNL"
	HRK_CurrencyCode CurrencyCode = "HRK"
	HTG_CurrencyCode CurrencyCode = "HTG"
	HUF_CurrencyCode CurrencyCode = "HUF"
	IDR_CurrencyCode CurrencyCode = "IDR"
	ILS_CurrencyCode CurrencyCode = "ILS"
	INR_CurrencyCode CurrencyCode = "INR"
	IQD_CurrencyCode CurrencyCode = "IQD"
	IRR_CurrencyCode CurrencyCode = "IRR"
	ISK_CurrencyCode CurrencyCode = "ISK"
	JMD_CurrencyCode CurrencyCode = "JMD"
	JOD_CurrencyCode CurrencyCode = "JOD"
	JPY_CurrencyCode CurrencyCode = "JPY"
	KES_CurrencyCode CurrencyCode = "KES"
	KGS_CurrencyCode CurrencyCode = "KGS"
	KHR_CurrencyCode CurrencyCode = "KHR"
	KMF_CurrencyCode CurrencyCode = "KMF"
	KPW_CurrencyCode CurrencyCode = "KPW"
	KRW_CurrencyCode CurrencyCode = "KRW"
	KWD_CurrencyCode CurrencyCode = "KWD"
	KYD_CurrencyCode CurrencyCode = "KYD"
	KZT_CurrencyCode CurrencyCode = "KZT"
	LAK_CurrencyCode CurrencyCode = "LAK"
	LBP_CurrencyCode CurrencyCode = "LBP"
	LKR_CurrencyCode CurrencyCode = "LKR"
	LRD_CurrencyCode CurrencyCode = "LRD"
	LSL_CurrencyCode CurrencyCode = "LSL"
	LTL_CurrencyCode CurrencyCode = "LTL"
	LVL_CurrencyCode CurrencyCode = "LVL"
	LYD_CurrencyCode CurrencyCode = "LYD"
	MAD_CurrencyCode CurrencyCode = "MAD"
	MDL_CurrencyCode CurrencyCode = "MDL"
	MGA_CurrencyCode CurrencyCode = "MGA"
	MKD_CurrencyCode CurrencyCode = "MKD"
	MMK_CurrencyCode CurrencyCode = "MMK"
	MNT_CurrencyCode CurrencyCode = "MNT"
	MOP_CurrencyCode CurrencyCode = "MOP"
	MRO_CurrencyCode CurrencyCode = "MRO"
	MRU_CurrencyCode CurrencyCode = "MRU"
	MUR_CurrencyCode CurrencyCode = "MUR"
	MVR_CurrencyCode CurrencyCode = "MVR"
	MWK_CurrencyCode CurrencyCode = "MWK"
	MXN_CurrencyCode CurrencyCode = "MXN"
	MXV_CurrencyCode CurrencyCode = "MXV"
	MYR_CurrencyCode CurrencyCode = "MYR"
	MZN_CurrencyCode CurrencyCode = "MZN"
	NAD_CurrencyCode CurrencyCode = "NAD"
	NGN_CurrencyCode CurrencyCode = "NGN"
	NIO_CurrencyCode CurrencyCode = "NIO"
	NOK_CurrencyCode CurrencyCode = "NOK"
	NPR_CurrencyCode CurrencyCode = "NPR"
	NZD_CurrencyCode CurrencyCode = "NZD"
	OMR_CurrencyCode CurrencyCode = "OMR"
	PAB_CurrencyCode CurrencyCode = "PAB"
	PEN_CurrencyCode CurrencyCode = "PEN"
	PGK_CurrencyCode CurrencyCode = "PGK"
	PHP_CurrencyCode CurrencyCode = "PHP"
	PKR_CurrencyCode CurrencyCode = "PKR"
	PLN_CurrencyCode CurrencyCode = "PLN"
	PYG_CurrencyCode CurrencyCode = "PYG"
	QAR_CurrencyCode CurrencyCode = "QAR"
	RON_CurrencyCode CurrencyCode = "RON"
	RSD_CurrencyCode CurrencyCode = "RSD"
	RUB_CurrencyCode CurrencyCode = "RUB"
	RWF_CurrencyCode CurrencyCode = "RWF"
	SAR_CurrencyCode CurrencyCode = "SAR"
	SBD_CurrencyCode CurrencyCode = "SBD"
	SCR_CurrencyCode CurrencyCode = "SCR"
	SDG_CurrencyCode CurrencyCode = "SDG"
	SEK_CurrencyCode CurrencyCode = "SEK"
	SGD_CurrencyCode CurrencyCode = "SGD"
	SHP_CurrencyCode CurrencyCode = "SHP"
	SKK_CurrencyCode CurrencyCode = "SKK"
	SLE_CurrencyCode CurrencyCode = "SLE"
	SLL_CurrencyCode CurrencyCode = "SLL"
	SOS_CurrencyCode CurrencyCode = "SOS"
	SRD_CurrencyCode CurrencyCode = "SRD"
	STN_CurrencyCode CurrencyCode = "STD"
	STD_CurrencyCode CurrencyCode = "STN"
	SVC_CurrencyCode CurrencyCode = "SVC"
	SYP_CurrencyCode CurrencyCode = "SYP"
	SZL_CurrencyCode CurrencyCode = "SZL"
	THB_CurrencyCode CurrencyCode = "THB"
	TJS_CurrencyCode CurrencyCode = "TJS"
	TMT_CurrencyCode CurrencyCode = "TMT"
	TND_CurrencyCode CurrencyCode = "TND"
	TOP_CurrencyCode CurrencyCode = "TOP"
	TRY_LIRA_CurrencyCode CurrencyCode = "TRY"
	TTD_CurrencyCode CurrencyCode = "TTD"
	TWD_CurrencyCode CurrencyCode = "TWD"
	TZS_CurrencyCode CurrencyCode = "TZS"
	UAH_CurrencyCode CurrencyCode = "UAH"
	UGX_CurrencyCode CurrencyCode = "UGX"
	USD_CurrencyCode CurrencyCode = "USD"
	UYU_CurrencyCode CurrencyCode = "UYU"
	UZS_CurrencyCode CurrencyCode = "UZS"
	VEF_CurrencyCode CurrencyCode = "VEF"
	VES_CurrencyCode CurrencyCode = "VES"
	VND_CurrencyCode CurrencyCode = "VND"
	VUV_CurrencyCode CurrencyCode = "VUV"
	WST_CurrencyCode CurrencyCode = "WST"
	XAF_CurrencyCode CurrencyCode = "XAF"
	XCD_CurrencyCode CurrencyCode = "XCD"
	XOF_CurrencyCode CurrencyCode = "XOF"
	XPF_CurrencyCode CurrencyCode = "XPF"
	YER_CurrencyCode CurrencyCode = "YER"
	ZAR_CurrencyCode CurrencyCode = "ZAR"
	ZMW_CurrencyCode CurrencyCode = "ZMW"
	ZMK_CurrencyCode CurrencyCode = "ZMK"
	ZWD_CurrencyCode CurrencyCode = "ZWD"
)

// All allowed values of CurrencyCode enum
var AllowedCurrencyCodeEnumValues = []CurrencyCode{
	"AED",
	"AFN",
	"ALL",
	"AMD",
	"ANG",
	"AOA",
	"ARS",
	"AUD",
	"AWG",
	"AZN",
	"BAM",
	"BBD",
	"BDT",
	"BGN",
	"BHD",
	"BIF",
	"BMD",
	"BND",
	"BOB",
	"BRL",
	"BSD",
	"BTN",
	"BWP",
	"BYN",
	"BYR",
	"BZD",
	"CAD",
	"CDF",
	"CHF",
	"CLF",
	"CLP",
	"CNY",
	"COP",
	"CRC",
	"CUC",
	"CUP",
	"CVE",
	"CZK",
	"DJF",
	"DKK",
	"DOP",
	"DZD",
	"EEK",
	"EGP",
	"ERN",
	"ETB",
	"EUR",
	"FJD",
	"FKP",
	"GBP",
	"GEL",
	"GHS",
	"GIP",
	"GMD",
	"GNF",
	"GTQ",
	"GYD",
	"HKD",
	"HNL",
	"HRK",
	"HTG",
	"HUF",
	"IDR",
	"ILS",
	"INR",
	"IQD",
	"IRR",
	"ISK",
	"JMD",
	"JOD",
	"JPY",
	"KES",
	"KGS",
	"KHR",
	"KMF",
	"KPW",
	"KRW",
	"KWD",
	"KYD",
	"KZT",
	"LAK",
	"LBP",
	"LKR",
	"LRD",
	"LSL",
	"LTL",
	"LVL",
	"LYD",
	"MAD",
	"MDL",
	"MGA",
	"MKD",
	"MMK",
	"MNT",
	"MOP",
	"MRO",
	"MRU",
	"MUR",
	"MVR",
	"MWK",
	"MXN",
	"MXV",
	"MYR",
	"MZN",
	"NAD",
	"NGN",
	"NIO",
	"NOK",
	"NPR",
	"NZD",
	"OMR",
	"PAB",
	"PEN",
	"PGK",
	"PHP",
	"PKR",
	"PLN",
	"PYG",
	"QAR",
	"RON",
	"RSD",
	"RUB",
	"RWF",
	"SAR",
	"SBD",
	"SCR",
	"SDG",
	"SEK",
	"SGD",
	"SHP",
	"SKK",
	"SLE",
	"SLL",
	"SOS",
	"SRD",
	"STD",
	"STN",
	"SVC",
	"SYP",
	"SZL",
	"THB",
	"TJS",
	"TMT",
	"TND",
	"TOP",
	"TRY",
	"TTD",
	"TWD",
	"TZS",
	"UAH",
	"UGX",
	"USD",
	"UYU",
	"UZS",
	"VEF",
	"VES",
	"VND",
	"VUV",
	"WST",
	"XAF",
	"XCD",
	"XOF",
	"XPF",
	"YER",
	"ZAR",
	"ZMW",
	"ZMK",
	"ZWD",
}

func (v *CurrencyCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CurrencyCode(value)
	for _, existing := range AllowedCurrencyCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CurrencyCode", value)
}

// NewCurrencyCodeFromValue returns a pointer to a valid CurrencyCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCurrencyCodeFromValue(v string) (*CurrencyCode, error) {
	ev := CurrencyCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CurrencyCode: valid values are %v", v, AllowedCurrencyCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CurrencyCode) IsValid() bool {
	for _, existing := range AllowedCurrencyCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CurrencyCode value
func (v CurrencyCode) Ptr() *CurrencyCode {
	return &v
}

type NullableCurrencyCode struct {
	value *CurrencyCode
	isSet bool
}

func (v NullableCurrencyCode) Get() *CurrencyCode {
	return v.value
}

func (v *NullableCurrencyCode) Set(val *CurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyCode(val *CurrencyCode) *NullableCurrencyCode {
	return &NullableCurrencyCode{value: val, isSet: true}
}

func (v NullableCurrencyCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

