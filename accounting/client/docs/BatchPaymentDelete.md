# BatchPaymentDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchPaymentID** | **string** | The Xero generated unique identifier for the bank transaction (read-only) | 
**Status** | **string** | The status of the batch payment. | [default to "DELETED"]

## Methods

### NewBatchPaymentDelete

`func NewBatchPaymentDelete(batchPaymentID string, status string, ) *BatchPaymentDelete`

NewBatchPaymentDelete instantiates a new BatchPaymentDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentDeleteWithDefaults

`func NewBatchPaymentDeleteWithDefaults() *BatchPaymentDelete`

NewBatchPaymentDeleteWithDefaults instantiates a new BatchPaymentDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchPaymentID

`func (o *BatchPaymentDelete) GetBatchPaymentID() string`

GetBatchPaymentID returns the BatchPaymentID field if non-nil, zero value otherwise.

### GetBatchPaymentIDOk

`func (o *BatchPaymentDelete) GetBatchPaymentIDOk() (*string, bool)`

GetBatchPaymentIDOk returns a tuple with the BatchPaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchPaymentID

`func (o *BatchPaymentDelete) SetBatchPaymentID(v string)`

SetBatchPaymentID sets BatchPaymentID field to given value.


### GetStatus

`func (o *BatchPaymentDelete) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchPaymentDelete) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchPaymentDelete) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


