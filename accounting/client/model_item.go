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
	"bytes"
	"fmt"
)

// checks if the Item type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Item{}

// Item struct for Item
type Item struct {
	// User defined item code (max length = 30)
	Code string `json:"Code"`
	// The inventory asset account for the item. The account must be of type INVENTORY. The  COGSAccountCode in PurchaseDetails is also required to create a tracked item
	InventoryAssetAccountCode *string `json:"InventoryAssetAccountCode,omitempty"`
	// The name of the item (max length = 50)
	Name *string `json:"Name,omitempty"`
	// Boolean value, defaults to true. When IsSold is true the item will be available on sales transactions in the Xero UI. If IsSold is updated to false then Description and SalesDetails values will be nulled.
	IsSold *bool `json:"IsSold,omitempty"`
	// Boolean value, defaults to true. When IsPurchased is true the item is available for purchase transactions in the Xero UI. If IsPurchased is updated to false then PurchaseDescription and PurchaseDetails values will be nulled.
	IsPurchased *bool `json:"IsPurchased,omitempty"`
	// The sales description of the item (max length = 4000)
	Description *string `json:"Description,omitempty"`
	// The purchase description of the item (max length = 4000)
	PurchaseDescription *string `json:"PurchaseDescription,omitempty"`
	PurchaseDetails *Purchase `json:"PurchaseDetails,omitempty"`
	SalesDetails *Purchase `json:"SalesDetails,omitempty"`
	// True for items that are tracked as inventory. An item will be tracked as inventory if the InventoryAssetAccountCode and COGSAccountCode are set.
	IsTrackedAsInventory *bool `json:"IsTrackedAsInventory,omitempty"`
	// The value of the item on hand. Calculated using average cost accounting.
	TotalCostPool *float64 `json:"TotalCostPool,omitempty"`
	// The quantity of the item on hand
	QuantityOnHand *float64 `json:"QuantityOnHand,omitempty"`
	// Last modified date in UTC format
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	// The Xero identifier for an Item
	ItemID *string `json:"ItemID,omitempty"`
	// Status of object
	StatusAttributeString *string `json:"StatusAttributeString,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
}

type _Item Item

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem(code string) *Item {
	this := Item{}
	this.Code = code
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetCode returns the Code field value
func (o *Item) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Item) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Item) SetCode(v string) {
	o.Code = v
}

// GetInventoryAssetAccountCode returns the InventoryAssetAccountCode field value if set, zero value otherwise.
func (o *Item) GetInventoryAssetAccountCode() string {
	if o == nil || IsNil(o.InventoryAssetAccountCode) {
		var ret string
		return ret
	}
	return *o.InventoryAssetAccountCode
}

// GetInventoryAssetAccountCodeOk returns a tuple with the InventoryAssetAccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetInventoryAssetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryAssetAccountCode) {
		return nil, false
	}
	return o.InventoryAssetAccountCode, true
}

// HasInventoryAssetAccountCodeField returns a boolean if a field has been set.
func (o *Item) HasInventoryAssetAccountCodeField() bool {
	if o != nil && !IsNil(o.InventoryAssetAccountCode) {
		return true
	}

	return false
}

// SetInventoryAssetAccountCode gets a reference to the given string and assigns it to the InventoryAssetAccountCode field.
func (o *Item) SetInventoryAssetAccountCode(v string) {
	o.InventoryAssetAccountCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Item) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasNameField returns a boolean if a field has been set.
func (o *Item) HasNameField() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Item) SetName(v string) {
	o.Name = &v
}

// GetIsSold returns the IsSold field value if set, zero value otherwise.
func (o *Item) GetIsSold() bool {
	if o == nil || IsNil(o.IsSold) {
		var ret bool
		return ret
	}
	return *o.IsSold
}

// GetIsSoldOk returns a tuple with the IsSold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetIsSoldOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSold) {
		return nil, false
	}
	return o.IsSold, true
}

// HasIsSoldField returns a boolean if a field has been set.
func (o *Item) HasIsSoldField() bool {
	if o != nil && !IsNil(o.IsSold) {
		return true
	}

	return false
}

// SetIsSold gets a reference to the given bool and assigns it to the IsSold field.
func (o *Item) SetIsSold(v bool) {
	o.IsSold = &v
}

// GetIsPurchased returns the IsPurchased field value if set, zero value otherwise.
func (o *Item) GetIsPurchased() bool {
	if o == nil || IsNil(o.IsPurchased) {
		var ret bool
		return ret
	}
	return *o.IsPurchased
}

// GetIsPurchasedOk returns a tuple with the IsPurchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetIsPurchasedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPurchased) {
		return nil, false
	}
	return o.IsPurchased, true
}

// HasIsPurchasedField returns a boolean if a field has been set.
func (o *Item) HasIsPurchasedField() bool {
	if o != nil && !IsNil(o.IsPurchased) {
		return true
	}

	return false
}

// SetIsPurchased gets a reference to the given bool and assigns it to the IsPurchased field.
func (o *Item) SetIsPurchased(v bool) {
	o.IsPurchased = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Item) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescriptionField returns a boolean if a field has been set.
func (o *Item) HasDescriptionField() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Item) SetDescription(v string) {
	o.Description = &v
}

// GetPurchaseDescription returns the PurchaseDescription field value if set, zero value otherwise.
func (o *Item) GetPurchaseDescription() string {
	if o == nil || IsNil(o.PurchaseDescription) {
		var ret string
		return ret
	}
	return *o.PurchaseDescription
}

// GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetPurchaseDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseDescription) {
		return nil, false
	}
	return o.PurchaseDescription, true
}

// HasPurchaseDescriptionField returns a boolean if a field has been set.
func (o *Item) HasPurchaseDescriptionField() bool {
	if o != nil && !IsNil(o.PurchaseDescription) {
		return true
	}

	return false
}

// SetPurchaseDescription gets a reference to the given string and assigns it to the PurchaseDescription field.
func (o *Item) SetPurchaseDescription(v string) {
	o.PurchaseDescription = &v
}

// GetPurchaseDetails returns the PurchaseDetails field value if set, zero value otherwise.
func (o *Item) GetPurchaseDetails() Purchase {
	if o == nil || IsNil(o.PurchaseDetails) {
		var ret Purchase
		return ret
	}
	return *o.PurchaseDetails
}

// GetPurchaseDetailsOk returns a tuple with the PurchaseDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetPurchaseDetailsOk() (*Purchase, bool) {
	if o == nil || IsNil(o.PurchaseDetails) {
		return nil, false
	}
	return o.PurchaseDetails, true
}

// HasPurchaseDetailsField returns a boolean if a field has been set.
func (o *Item) HasPurchaseDetailsField() bool {
	if o != nil && !IsNil(o.PurchaseDetails) {
		return true
	}

	return false
}

// SetPurchaseDetails gets a reference to the given Purchase and assigns it to the PurchaseDetails field.
func (o *Item) SetPurchaseDetails(v Purchase) {
	o.PurchaseDetails = &v
}

// GetSalesDetails returns the SalesDetails field value if set, zero value otherwise.
func (o *Item) GetSalesDetails() Purchase {
	if o == nil || IsNil(o.SalesDetails) {
		var ret Purchase
		return ret
	}
	return *o.SalesDetails
}

// GetSalesDetailsOk returns a tuple with the SalesDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetSalesDetailsOk() (*Purchase, bool) {
	if o == nil || IsNil(o.SalesDetails) {
		return nil, false
	}
	return o.SalesDetails, true
}

// HasSalesDetailsField returns a boolean if a field has been set.
func (o *Item) HasSalesDetailsField() bool {
	if o != nil && !IsNil(o.SalesDetails) {
		return true
	}

	return false
}

// SetSalesDetails gets a reference to the given Purchase and assigns it to the SalesDetails field.
func (o *Item) SetSalesDetails(v Purchase) {
	o.SalesDetails = &v
}

// GetIsTrackedAsInventory returns the IsTrackedAsInventory field value if set, zero value otherwise.
func (o *Item) GetIsTrackedAsInventory() bool {
	if o == nil || IsNil(o.IsTrackedAsInventory) {
		var ret bool
		return ret
	}
	return *o.IsTrackedAsInventory
}

// GetIsTrackedAsInventoryOk returns a tuple with the IsTrackedAsInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetIsTrackedAsInventoryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTrackedAsInventory) {
		return nil, false
	}
	return o.IsTrackedAsInventory, true
}

// HasIsTrackedAsInventoryField returns a boolean if a field has been set.
func (o *Item) HasIsTrackedAsInventoryField() bool {
	if o != nil && !IsNil(o.IsTrackedAsInventory) {
		return true
	}

	return false
}

// SetIsTrackedAsInventory gets a reference to the given bool and assigns it to the IsTrackedAsInventory field.
func (o *Item) SetIsTrackedAsInventory(v bool) {
	o.IsTrackedAsInventory = &v
}

// GetTotalCostPool returns the TotalCostPool field value if set, zero value otherwise.
func (o *Item) GetTotalCostPool() float64 {
	if o == nil || IsNil(o.TotalCostPool) {
		var ret float64
		return ret
	}
	return *o.TotalCostPool
}

// GetTotalCostPoolOk returns a tuple with the TotalCostPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetTotalCostPoolOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCostPool) {
		return nil, false
	}
	return o.TotalCostPool, true
}

// HasTotalCostPoolField returns a boolean if a field has been set.
func (o *Item) HasTotalCostPoolField() bool {
	if o != nil && !IsNil(o.TotalCostPool) {
		return true
	}

	return false
}

// SetTotalCostPool gets a reference to the given float64 and assigns it to the TotalCostPool field.
func (o *Item) SetTotalCostPool(v float64) {
	o.TotalCostPool = &v
}

// GetQuantityOnHand returns the QuantityOnHand field value if set, zero value otherwise.
func (o *Item) GetQuantityOnHand() float64 {
	if o == nil || IsNil(o.QuantityOnHand) {
		var ret float64
		return ret
	}
	return *o.QuantityOnHand
}

// GetQuantityOnHandOk returns a tuple with the QuantityOnHand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetQuantityOnHandOk() (*float64, bool) {
	if o == nil || IsNil(o.QuantityOnHand) {
		return nil, false
	}
	return o.QuantityOnHand, true
}

// HasQuantityOnHandField returns a boolean if a field has been set.
func (o *Item) HasQuantityOnHandField() bool {
	if o != nil && !IsNil(o.QuantityOnHand) {
		return true
	}

	return false
}

// SetQuantityOnHand gets a reference to the given float64 and assigns it to the QuantityOnHand field.
func (o *Item) SetQuantityOnHand(v float64) {
	o.QuantityOnHand = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *Item) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTCField returns a boolean if a field has been set.
func (o *Item) HasUpdatedDateUTCField() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *Item) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetItemID returns the ItemID field value if set, zero value otherwise.
func (o *Item) GetItemID() string {
	if o == nil || IsNil(o.ItemID) {
		var ret string
		return ret
	}
	return *o.ItemID
}

// GetItemIDOk returns a tuple with the ItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetItemIDOk() (*string, bool) {
	if o == nil || IsNil(o.ItemID) {
		return nil, false
	}
	return o.ItemID, true
}

// HasItemIDField returns a boolean if a field has been set.
func (o *Item) HasItemIDField() bool {
	if o != nil && !IsNil(o.ItemID) {
		return true
	}

	return false
}

// SetItemID gets a reference to the given string and assigns it to the ItemID field.
func (o *Item) SetItemID(v string) {
	o.ItemID = &v
}

// GetStatusAttributeString returns the StatusAttributeString field value if set, zero value otherwise.
func (o *Item) GetStatusAttributeString() string {
	if o == nil || IsNil(o.StatusAttributeString) {
		var ret string
		return ret
	}
	return *o.StatusAttributeString
}

// GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetStatusAttributeStringOk() (*string, bool) {
	if o == nil || IsNil(o.StatusAttributeString) {
		return nil, false
	}
	return o.StatusAttributeString, true
}

// HasStatusAttributeStringField returns a boolean if a field has been set.
func (o *Item) HasStatusAttributeStringField() bool {
	if o != nil && !IsNil(o.StatusAttributeString) {
		return true
	}

	return false
}

// SetStatusAttributeString gets a reference to the given string and assigns it to the StatusAttributeString field.
func (o *Item) SetStatusAttributeString(v string) {
	o.StatusAttributeString = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *Item) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrorsField returns a boolean if a field has been set.
func (o *Item) HasValidationErrorsField() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *Item) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

func (o Item) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Item) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Code"] = o.Code
	if !IsNil(o.InventoryAssetAccountCode) {
		toSerialize["InventoryAssetAccountCode"] = o.InventoryAssetAccountCode
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.IsSold) {
		toSerialize["IsSold"] = o.IsSold
	}
	if !IsNil(o.IsPurchased) {
		toSerialize["IsPurchased"] = o.IsPurchased
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.PurchaseDescription) {
		toSerialize["PurchaseDescription"] = o.PurchaseDescription
	}
	if !IsNil(o.PurchaseDetails) {
		toSerialize["PurchaseDetails"] = o.PurchaseDetails
	}
	if !IsNil(o.SalesDetails) {
		toSerialize["SalesDetails"] = o.SalesDetails
	}
	if !IsNil(o.IsTrackedAsInventory) {
		toSerialize["IsTrackedAsInventory"] = o.IsTrackedAsInventory
	}
	if !IsNil(o.TotalCostPool) {
		toSerialize["TotalCostPool"] = o.TotalCostPool
	}
	if !IsNil(o.QuantityOnHand) {
		toSerialize["QuantityOnHand"] = o.QuantityOnHand
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.ItemID) {
		toSerialize["ItemID"] = o.ItemID
	}
	if !IsNil(o.StatusAttributeString) {
		toSerialize["StatusAttributeString"] = o.StatusAttributeString
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

func (o *Item) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varItem := _Item{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varItem)

	if err != nil {
		return err
	}

	*o = Item(varItem)

	return err
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


