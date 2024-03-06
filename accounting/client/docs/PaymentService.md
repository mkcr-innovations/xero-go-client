# PaymentService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentServiceID** | Pointer to **string** | Xero identifier | [optional] 
**PaymentServiceName** | Pointer to **string** | Name of payment service | [optional] 
**PaymentServiceUrl** | Pointer to **string** | The custom payment URL | [optional] 
**PayNowText** | Pointer to **string** | The text displayed on the Pay Now button in Xero Online Invoicing. If this is not set it will default to Pay by credit card | [optional] 
**PaymentServiceType** | Pointer to **string** | This will always be CUSTOM for payment services created via the API. | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

## Methods

### NewPaymentService

`func NewPaymentService() *PaymentService`

NewPaymentService instantiates a new PaymentService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceWithDefaults

`func NewPaymentServiceWithDefaults() *PaymentService`

NewPaymentServiceWithDefaults instantiates a new PaymentService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentServiceID

`func (o *PaymentService) GetPaymentServiceID() string`

GetPaymentServiceID returns the PaymentServiceID field if non-nil, zero value otherwise.

### GetPaymentServiceIDOk

`func (o *PaymentService) GetPaymentServiceIDOk() (*string, bool)`

GetPaymentServiceIDOk returns a tuple with the PaymentServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceID

`func (o *PaymentService) SetPaymentServiceID(v string)`

SetPaymentServiceID sets PaymentServiceID field to given value.

### HasPaymentServiceIDField

`func (o *PaymentService) HasPaymentServiceIDField() bool`

HasPaymentServiceIDField returns a boolean if a field has been set.

### GetPaymentServiceName

`func (o *PaymentService) GetPaymentServiceName() string`

GetPaymentServiceName returns the PaymentServiceName field if non-nil, zero value otherwise.

### GetPaymentServiceNameOk

`func (o *PaymentService) GetPaymentServiceNameOk() (*string, bool)`

GetPaymentServiceNameOk returns a tuple with the PaymentServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceName

`func (o *PaymentService) SetPaymentServiceName(v string)`

SetPaymentServiceName sets PaymentServiceName field to given value.

### HasPaymentServiceNameField

`func (o *PaymentService) HasPaymentServiceNameField() bool`

HasPaymentServiceNameField returns a boolean if a field has been set.

### GetPaymentServiceUrl

`func (o *PaymentService) GetPaymentServiceUrl() string`

GetPaymentServiceUrl returns the PaymentServiceUrl field if non-nil, zero value otherwise.

### GetPaymentServiceUrlOk

`func (o *PaymentService) GetPaymentServiceUrlOk() (*string, bool)`

GetPaymentServiceUrlOk returns a tuple with the PaymentServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceUrl

`func (o *PaymentService) SetPaymentServiceUrl(v string)`

SetPaymentServiceUrl sets PaymentServiceUrl field to given value.

### HasPaymentServiceUrlField

`func (o *PaymentService) HasPaymentServiceUrlField() bool`

HasPaymentServiceUrlField returns a boolean if a field has been set.

### GetPayNowText

`func (o *PaymentService) GetPayNowText() string`

GetPayNowText returns the PayNowText field if non-nil, zero value otherwise.

### GetPayNowTextOk

`func (o *PaymentService) GetPayNowTextOk() (*string, bool)`

GetPayNowTextOk returns a tuple with the PayNowText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayNowText

`func (o *PaymentService) SetPayNowText(v string)`

SetPayNowText sets PayNowText field to given value.

### HasPayNowTextField

`func (o *PaymentService) HasPayNowTextField() bool`

HasPayNowTextField returns a boolean if a field has been set.

### GetPaymentServiceType

`func (o *PaymentService) GetPaymentServiceType() string`

GetPaymentServiceType returns the PaymentServiceType field if non-nil, zero value otherwise.

### GetPaymentServiceTypeOk

`func (o *PaymentService) GetPaymentServiceTypeOk() (*string, bool)`

GetPaymentServiceTypeOk returns a tuple with the PaymentServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceType

`func (o *PaymentService) SetPaymentServiceType(v string)`

SetPaymentServiceType sets PaymentServiceType field to given value.

### HasPaymentServiceTypeField

`func (o *PaymentService) HasPaymentServiceTypeField() bool`

HasPaymentServiceTypeField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *PaymentService) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PaymentService) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PaymentService) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *PaymentService) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


