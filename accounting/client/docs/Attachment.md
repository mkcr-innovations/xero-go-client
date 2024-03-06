# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentID** | Pointer to **string** | Unique ID for the file | [optional] 
**FileName** | Pointer to **string** | Name of the file | [optional] 
**Url** | Pointer to **string** | URL to the file on xero.com | [optional] 
**MimeType** | Pointer to **string** | Type of file | [optional] 
**ContentLength** | Pointer to **int32** | Length of the file content | [optional] 
**IncludeOnline** | Pointer to **bool** | Include the file with the online invoice | [optional] 

## Methods

### NewAttachment

`func NewAttachment() *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentID

`func (o *Attachment) GetAttachmentID() string`

GetAttachmentID returns the AttachmentID field if non-nil, zero value otherwise.

### GetAttachmentIDOk

`func (o *Attachment) GetAttachmentIDOk() (*string, bool)`

GetAttachmentIDOk returns a tuple with the AttachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentID

`func (o *Attachment) SetAttachmentID(v string)`

SetAttachmentID sets AttachmentID field to given value.

### HasAttachmentIDField

`func (o *Attachment) HasAttachmentIDField() bool`

HasAttachmentIDField returns a boolean if a field has been set.

### GetFileName

`func (o *Attachment) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Attachment) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Attachment) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileNameField

`func (o *Attachment) HasFileNameField() bool`

HasFileNameField returns a boolean if a field has been set.

### GetUrl

`func (o *Attachment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Attachment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Attachment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrlField

`func (o *Attachment) HasUrlField() bool`

HasUrlField returns a boolean if a field has been set.

### GetMimeType

`func (o *Attachment) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Attachment) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Attachment) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeTypeField

`func (o *Attachment) HasMimeTypeField() bool`

HasMimeTypeField returns a boolean if a field has been set.

### GetContentLength

`func (o *Attachment) GetContentLength() int32`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *Attachment) GetContentLengthOk() (*int32, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *Attachment) SetContentLength(v int32)`

SetContentLength sets ContentLength field to given value.

### HasContentLengthField

`func (o *Attachment) HasContentLengthField() bool`

HasContentLengthField returns a boolean if a field has been set.

### GetIncludeOnline

`func (o *Attachment) GetIncludeOnline() bool`

GetIncludeOnline returns the IncludeOnline field if non-nil, zero value otherwise.

### GetIncludeOnlineOk

`func (o *Attachment) GetIncludeOnlineOk() (*bool, bool)`

GetIncludeOnlineOk returns a tuple with the IncludeOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOnline

`func (o *Attachment) SetIncludeOnline(v bool)`

SetIncludeOnline sets IncludeOnline field to given value.

### HasIncludeOnlineField

`func (o *Attachment) HasIncludeOnlineField() bool`

HasIncludeOnlineField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


