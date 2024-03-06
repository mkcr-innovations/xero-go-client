# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | User defined item code (max length &#x3D; 30) | 
**InventoryAssetAccountCode** | Pointer to **string** | The inventory asset account for the item. The account must be of type INVENTORY. The  COGSAccountCode in PurchaseDetails is also required to create a tracked item | [optional] 
**Name** | Pointer to **string** | The name of the item (max length &#x3D; 50) | [optional] 
**IsSold** | Pointer to **bool** | Boolean value, defaults to true. When IsSold is true the item will be available on sales transactions in the Xero UI. If IsSold is updated to false then Description and SalesDetails values will be nulled. | [optional] 
**IsPurchased** | Pointer to **bool** | Boolean value, defaults to true. When IsPurchased is true the item is available for purchase transactions in the Xero UI. If IsPurchased is updated to false then PurchaseDescription and PurchaseDetails values will be nulled. | [optional] 
**Description** | Pointer to **string** | The sales description of the item (max length &#x3D; 4000) | [optional] 
**PurchaseDescription** | Pointer to **string** | The purchase description of the item (max length &#x3D; 4000) | [optional] 
**PurchaseDetails** | Pointer to [**Purchase**](Purchase.md) |  | [optional] 
**SalesDetails** | Pointer to [**Purchase**](Purchase.md) |  | [optional] 
**IsTrackedAsInventory** | Pointer to **bool** | True for items that are tracked as inventory. An item will be tracked as inventory if the InventoryAssetAccountCode and COGSAccountCode are set. | [optional] 
**TotalCostPool** | Pointer to **float64** | The value of the item on hand. Calculated using average cost accounting. | [optional] 
**QuantityOnHand** | Pointer to **float64** | The quantity of the item on hand | [optional] 
**UpdatedDateUTC** | Pointer to **string** | Last modified date in UTC format | [optional] [readonly] 
**ItemID** | Pointer to **string** | The Xero identifier for an Item | [optional] 
**StatusAttributeString** | Pointer to **string** | Status of object | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewItem

`func NewItem(code string, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Item) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Item) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Item) SetCode(v string)`

SetCode sets Code field to given value.


### GetInventoryAssetAccountCode

`func (o *Item) GetInventoryAssetAccountCode() string`

GetInventoryAssetAccountCode returns the InventoryAssetAccountCode field if non-nil, zero value otherwise.

### GetInventoryAssetAccountCodeOk

`func (o *Item) GetInventoryAssetAccountCodeOk() (*string, bool)`

GetInventoryAssetAccountCodeOk returns a tuple with the InventoryAssetAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAssetAccountCode

`func (o *Item) SetInventoryAssetAccountCode(v string)`

SetInventoryAssetAccountCode sets InventoryAssetAccountCode field to given value.

### HasInventoryAssetAccountCodeField

`func (o *Item) HasInventoryAssetAccountCodeField() bool`

HasInventoryAssetAccountCodeField returns a boolean if a field has been set.

### GetName

`func (o *Item) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Item) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Item) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *Item) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetIsSold

`func (o *Item) GetIsSold() bool`

GetIsSold returns the IsSold field if non-nil, zero value otherwise.

### GetIsSoldOk

`func (o *Item) GetIsSoldOk() (*bool, bool)`

GetIsSoldOk returns a tuple with the IsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSold

`func (o *Item) SetIsSold(v bool)`

SetIsSold sets IsSold field to given value.

### HasIsSoldField

`func (o *Item) HasIsSoldField() bool`

HasIsSoldField returns a boolean if a field has been set.

### GetIsPurchased

`func (o *Item) GetIsPurchased() bool`

GetIsPurchased returns the IsPurchased field if non-nil, zero value otherwise.

### GetIsPurchasedOk

`func (o *Item) GetIsPurchasedOk() (*bool, bool)`

GetIsPurchasedOk returns a tuple with the IsPurchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPurchased

`func (o *Item) SetIsPurchased(v bool)`

SetIsPurchased sets IsPurchased field to given value.

### HasIsPurchasedField

`func (o *Item) HasIsPurchasedField() bool`

HasIsPurchasedField returns a boolean if a field has been set.

### GetDescription

`func (o *Item) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Item) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Item) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescriptionField

`func (o *Item) HasDescriptionField() bool`

HasDescriptionField returns a boolean if a field has been set.

### GetPurchaseDescription

`func (o *Item) GetPurchaseDescription() string`

GetPurchaseDescription returns the PurchaseDescription field if non-nil, zero value otherwise.

### GetPurchaseDescriptionOk

`func (o *Item) GetPurchaseDescriptionOk() (*string, bool)`

GetPurchaseDescriptionOk returns a tuple with the PurchaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDescription

`func (o *Item) SetPurchaseDescription(v string)`

SetPurchaseDescription sets PurchaseDescription field to given value.

### HasPurchaseDescriptionField

`func (o *Item) HasPurchaseDescriptionField() bool`

HasPurchaseDescriptionField returns a boolean if a field has been set.

### GetPurchaseDetails

`func (o *Item) GetPurchaseDetails() Purchase`

GetPurchaseDetails returns the PurchaseDetails field if non-nil, zero value otherwise.

### GetPurchaseDetailsOk

`func (o *Item) GetPurchaseDetailsOk() (*Purchase, bool)`

GetPurchaseDetailsOk returns a tuple with the PurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDetails

`func (o *Item) SetPurchaseDetails(v Purchase)`

SetPurchaseDetails sets PurchaseDetails field to given value.

### HasPurchaseDetailsField

`func (o *Item) HasPurchaseDetailsField() bool`

HasPurchaseDetailsField returns a boolean if a field has been set.

### GetSalesDetails

`func (o *Item) GetSalesDetails() Purchase`

GetSalesDetails returns the SalesDetails field if non-nil, zero value otherwise.

### GetSalesDetailsOk

`func (o *Item) GetSalesDetailsOk() (*Purchase, bool)`

GetSalesDetailsOk returns a tuple with the SalesDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDetails

`func (o *Item) SetSalesDetails(v Purchase)`

SetSalesDetails sets SalesDetails field to given value.

### HasSalesDetailsField

`func (o *Item) HasSalesDetailsField() bool`

HasSalesDetailsField returns a boolean if a field has been set.

### GetIsTrackedAsInventory

`func (o *Item) GetIsTrackedAsInventory() bool`

GetIsTrackedAsInventory returns the IsTrackedAsInventory field if non-nil, zero value otherwise.

### GetIsTrackedAsInventoryOk

`func (o *Item) GetIsTrackedAsInventoryOk() (*bool, bool)`

GetIsTrackedAsInventoryOk returns a tuple with the IsTrackedAsInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackedAsInventory

`func (o *Item) SetIsTrackedAsInventory(v bool)`

SetIsTrackedAsInventory sets IsTrackedAsInventory field to given value.

### HasIsTrackedAsInventoryField

`func (o *Item) HasIsTrackedAsInventoryField() bool`

HasIsTrackedAsInventoryField returns a boolean if a field has been set.

### GetTotalCostPool

`func (o *Item) GetTotalCostPool() float64`

GetTotalCostPool returns the TotalCostPool field if non-nil, zero value otherwise.

### GetTotalCostPoolOk

`func (o *Item) GetTotalCostPoolOk() (*float64, bool)`

GetTotalCostPoolOk returns a tuple with the TotalCostPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostPool

`func (o *Item) SetTotalCostPool(v float64)`

SetTotalCostPool sets TotalCostPool field to given value.

### HasTotalCostPoolField

`func (o *Item) HasTotalCostPoolField() bool`

HasTotalCostPoolField returns a boolean if a field has been set.

### GetQuantityOnHand

`func (o *Item) GetQuantityOnHand() float64`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *Item) GetQuantityOnHandOk() (*float64, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *Item) SetQuantityOnHand(v float64)`

SetQuantityOnHand sets QuantityOnHand field to given value.

### HasQuantityOnHandField

`func (o *Item) HasQuantityOnHandField() bool`

HasQuantityOnHandField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Item) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Item) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Item) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Item) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetItemID

`func (o *Item) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *Item) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *Item) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemIDField

`func (o *Item) HasItemIDField() bool`

HasItemIDField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *Item) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *Item) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *Item) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *Item) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Item) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Item) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Item) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Item) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


