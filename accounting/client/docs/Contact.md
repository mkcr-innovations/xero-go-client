# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactID** | Pointer to **string** | Xero identifier | [optional] 
**MergedToContactID** | Pointer to **string** | ID for the destination of a merged contact. Only returned when using paging or when fetching a contact by ContactId or ContactNumber. | [optional] 
**ContactNumber** | Pointer to **string** | This can be updated via the API only i.e. This field is read only on the Xero contact screen, used to identify contacts in external systems (max length &#x3D; 50). If the Contact Number is used, this is displayed as Contact Code in the Contacts UI in Xero. | [optional] 
**AccountNumber** | Pointer to **string** | A user defined account number. This can be updated via the API and the Xero UI (max length &#x3D; 50) | [optional] 
**ContactStatus** | Pointer to **string** | Current status of a contact – see contact status types | [optional] 
**Name** | Pointer to **string** | Full name of contact/organisation (max length &#x3D; 255) | [optional] 
**FirstName** | Pointer to **string** | First name of contact person (max length &#x3D; 255) | [optional] 
**LastName** | Pointer to **string** | Last name of contact person (max length &#x3D; 255) | [optional] 
**CompanyNumber** | Pointer to **string** | Company registration number (max length &#x3D; 50) | [optional] 
**EmailAddress** | Pointer to **string** | Email address of contact person (umlauts not supported) (max length  &#x3D; 255) | [optional] 
**ContactPersons** | Pointer to [**[]ContactPerson**](ContactPerson.md) | See contact persons | [optional] 
**BankAccountDetails** | Pointer to **string** | Bank account number of contact | [optional] 
**TaxNumber** | Pointer to **string** | Tax number of contact – this is also known as the ABN (Australia), GST Number (New Zealand), VAT Number (UK) or Tax ID Number (US and global) in the Xero UI depending on which regionalized version of Xero you are using (max length &#x3D; 50) | [optional] 
**AccountsReceivableTaxType** | Pointer to **string** | The tax type from TaxRates | [optional] 
**AccountsPayableTaxType** | Pointer to **string** | The tax type from TaxRates | [optional] 
**Addresses** | Pointer to [**[]Address**](Address.md) | Store certain address types for a contact – see address types | [optional] 
**Phones** | Pointer to [**[]Phone**](Phone.md) | Store certain phone types for a contact – see phone types | [optional] 
**IsSupplier** | Pointer to **bool** | true or false – Boolean that describes if a contact that has any AP  invoices entered against them. Cannot be set via PUT or POST – it is automatically set when an accounts payable invoice is generated against this contact. | [optional] 
**IsCustomer** | Pointer to **bool** | true or false – Boolean that describes if a contact has any AR invoices entered against them. Cannot be set via PUT or POST – it is automatically set when an accounts receivable invoice is generated against this contact. | [optional] 
**SalesDefaultLineAmountType** | Pointer to **string** | The default sales line amount type for a contact. Only available when summaryOnly parameter or paging is used, or when fetch by ContactId or ContactNumber. | [optional] 
**PurchasesDefaultLineAmountType** | Pointer to **string** | The default purchases line amount type for a contact Only available when summaryOnly parameter or paging is used, or when fetch by ContactId or ContactNumber. | [optional] 
**DefaultCurrency** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**XeroNetworkKey** | Pointer to **string** | Store XeroNetworkKey for contacts. | [optional] 
**SalesDefaultAccountCode** | Pointer to **string** | The default sales account code for contacts | [optional] 
**PurchasesDefaultAccountCode** | Pointer to **string** | The default purchases account code for contacts | [optional] 
**SalesTrackingCategories** | Pointer to [**[]SalesTrackingCategory**](SalesTrackingCategory.md) | The default sales tracking categories for contacts | [optional] 
**PurchasesTrackingCategories** | Pointer to [**[]SalesTrackingCategory**](SalesTrackingCategory.md) | The default purchases tracking categories for contacts | [optional] 
**TrackingCategoryName** | Pointer to **string** | The name of the Tracking Category assigned to the contact under SalesTrackingCategories and PurchasesTrackingCategories | [optional] 
**TrackingCategoryOption** | Pointer to **string** | The name of the Tracking Option assigned to the contact under SalesTrackingCategories and PurchasesTrackingCategories | [optional] 
**PaymentTerms** | Pointer to [**PaymentTerm**](PaymentTerm.md) |  | [optional] 
**UpdatedDateUTC** | Pointer to **string** | UTC timestamp of last update to contact | [optional] [readonly] 
**ContactGroups** | Pointer to [**[]ContactGroup**](ContactGroup.md) | Displays which contact groups a contact is included in | [optional] 
**Website** | Pointer to **string** | Website address for contact (read only) | [optional] [readonly] 
**BrandingTheme** | Pointer to [**BrandingTheme**](BrandingTheme.md) |  | [optional] 
**BatchPayments** | Pointer to [**BatchPaymentDetails**](BatchPaymentDetails.md) |  | [optional] 
**Discount** | Pointer to **float64** | The default discount rate for the contact (read only) | [optional] [readonly] 
**Balances** | Pointer to [**Balances**](Balances.md) |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 
**HasAttachments** | Pointer to **bool** | A boolean to indicate if a contact has an attachment | [optional] [default to false]
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | Displays validation errors returned from the API | [optional] 
**HasValidationErrors** | Pointer to **bool** | A boolean to indicate if a contact has an validation errors | [optional] [default to false]
**StatusAttributeString** | Pointer to **string** | Status of object | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactID

`func (o *Contact) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *Contact) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *Contact) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactIDField

`func (o *Contact) HasContactIDField() bool`

HasContactIDField returns a boolean if a field has been set.

### GetMergedToContactID

`func (o *Contact) GetMergedToContactID() string`

GetMergedToContactID returns the MergedToContactID field if non-nil, zero value otherwise.

### GetMergedToContactIDOk

`func (o *Contact) GetMergedToContactIDOk() (*string, bool)`

GetMergedToContactIDOk returns a tuple with the MergedToContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedToContactID

`func (o *Contact) SetMergedToContactID(v string)`

SetMergedToContactID sets MergedToContactID field to given value.

### HasMergedToContactIDField

`func (o *Contact) HasMergedToContactIDField() bool`

HasMergedToContactIDField returns a boolean if a field has been set.

### GetContactNumber

`func (o *Contact) GetContactNumber() string`

GetContactNumber returns the ContactNumber field if non-nil, zero value otherwise.

### GetContactNumberOk

`func (o *Contact) GetContactNumberOk() (*string, bool)`

GetContactNumberOk returns a tuple with the ContactNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactNumber

`func (o *Contact) SetContactNumber(v string)`

SetContactNumber sets ContactNumber field to given value.

### HasContactNumberField

`func (o *Contact) HasContactNumberField() bool`

HasContactNumberField returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Contact) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Contact) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Contact) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumberField

`func (o *Contact) HasAccountNumberField() bool`

HasAccountNumberField returns a boolean if a field has been set.

### GetContactStatus

`func (o *Contact) GetContactStatus() string`

GetContactStatus returns the ContactStatus field if non-nil, zero value otherwise.

### GetContactStatusOk

`func (o *Contact) GetContactStatusOk() (*string, bool)`

GetContactStatusOk returns a tuple with the ContactStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactStatus

`func (o *Contact) SetContactStatus(v string)`

SetContactStatus sets ContactStatus field to given value.

### HasContactStatusField

`func (o *Contact) HasContactStatusField() bool`

HasContactStatusField returns a boolean if a field has been set.

### GetName

`func (o *Contact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contact) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *Contact) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetFirstName

`func (o *Contact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstNameField

`func (o *Contact) HasFirstNameField() bool`

HasFirstNameField returns a boolean if a field has been set.

### GetLastName

`func (o *Contact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastNameField

`func (o *Contact) HasLastNameField() bool`

HasLastNameField returns a boolean if a field has been set.

### GetCompanyNumber

`func (o *Contact) GetCompanyNumber() string`

GetCompanyNumber returns the CompanyNumber field if non-nil, zero value otherwise.

### GetCompanyNumberOk

`func (o *Contact) GetCompanyNumberOk() (*string, bool)`

GetCompanyNumberOk returns a tuple with the CompanyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNumber

`func (o *Contact) SetCompanyNumber(v string)`

SetCompanyNumber sets CompanyNumber field to given value.

### HasCompanyNumberField

`func (o *Contact) HasCompanyNumberField() bool`

HasCompanyNumberField returns a boolean if a field has been set.

### GetEmailAddress

`func (o *Contact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Contact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Contact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddressField

`func (o *Contact) HasEmailAddressField() bool`

HasEmailAddressField returns a boolean if a field has been set.

### GetContactPersons

`func (o *Contact) GetContactPersons() []ContactPerson`

GetContactPersons returns the ContactPersons field if non-nil, zero value otherwise.

### GetContactPersonsOk

`func (o *Contact) GetContactPersonsOk() (*[]ContactPerson, bool)`

GetContactPersonsOk returns a tuple with the ContactPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersons

`func (o *Contact) SetContactPersons(v []ContactPerson)`

SetContactPersons sets ContactPersons field to given value.

### HasContactPersonsField

`func (o *Contact) HasContactPersonsField() bool`

HasContactPersonsField returns a boolean if a field has been set.

### GetBankAccountDetails

`func (o *Contact) GetBankAccountDetails() string`

GetBankAccountDetails returns the BankAccountDetails field if non-nil, zero value otherwise.

### GetBankAccountDetailsOk

`func (o *Contact) GetBankAccountDetailsOk() (*string, bool)`

GetBankAccountDetailsOk returns a tuple with the BankAccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountDetails

`func (o *Contact) SetBankAccountDetails(v string)`

SetBankAccountDetails sets BankAccountDetails field to given value.

### HasBankAccountDetailsField

`func (o *Contact) HasBankAccountDetailsField() bool`

HasBankAccountDetailsField returns a boolean if a field has been set.

### GetTaxNumber

`func (o *Contact) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *Contact) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *Contact) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumberField

`func (o *Contact) HasTaxNumberField() bool`

HasTaxNumberField returns a boolean if a field has been set.

### GetAccountsReceivableTaxType

`func (o *Contact) GetAccountsReceivableTaxType() string`

GetAccountsReceivableTaxType returns the AccountsReceivableTaxType field if non-nil, zero value otherwise.

### GetAccountsReceivableTaxTypeOk

`func (o *Contact) GetAccountsReceivableTaxTypeOk() (*string, bool)`

GetAccountsReceivableTaxTypeOk returns a tuple with the AccountsReceivableTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsReceivableTaxType

`func (o *Contact) SetAccountsReceivableTaxType(v string)`

SetAccountsReceivableTaxType sets AccountsReceivableTaxType field to given value.

### HasAccountsReceivableTaxTypeField

`func (o *Contact) HasAccountsReceivableTaxTypeField() bool`

HasAccountsReceivableTaxTypeField returns a boolean if a field has been set.

### GetAccountsPayableTaxType

`func (o *Contact) GetAccountsPayableTaxType() string`

GetAccountsPayableTaxType returns the AccountsPayableTaxType field if non-nil, zero value otherwise.

### GetAccountsPayableTaxTypeOk

`func (o *Contact) GetAccountsPayableTaxTypeOk() (*string, bool)`

GetAccountsPayableTaxTypeOk returns a tuple with the AccountsPayableTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsPayableTaxType

`func (o *Contact) SetAccountsPayableTaxType(v string)`

SetAccountsPayableTaxType sets AccountsPayableTaxType field to given value.

### HasAccountsPayableTaxTypeField

`func (o *Contact) HasAccountsPayableTaxTypeField() bool`

HasAccountsPayableTaxTypeField returns a boolean if a field has been set.

### GetAddresses

`func (o *Contact) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Contact) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Contact) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddressesField

`func (o *Contact) HasAddressesField() bool`

HasAddressesField returns a boolean if a field has been set.

### GetPhones

`func (o *Contact) GetPhones() []Phone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *Contact) GetPhonesOk() (*[]Phone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *Contact) SetPhones(v []Phone)`

SetPhones sets Phones field to given value.

### HasPhonesField

`func (o *Contact) HasPhonesField() bool`

HasPhonesField returns a boolean if a field has been set.

### GetIsSupplier

`func (o *Contact) GetIsSupplier() bool`

GetIsSupplier returns the IsSupplier field if non-nil, zero value otherwise.

### GetIsSupplierOk

`func (o *Contact) GetIsSupplierOk() (*bool, bool)`

GetIsSupplierOk returns a tuple with the IsSupplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupplier

`func (o *Contact) SetIsSupplier(v bool)`

SetIsSupplier sets IsSupplier field to given value.

### HasIsSupplierField

`func (o *Contact) HasIsSupplierField() bool`

HasIsSupplierField returns a boolean if a field has been set.

### GetIsCustomer

`func (o *Contact) GetIsCustomer() bool`

GetIsCustomer returns the IsCustomer field if non-nil, zero value otherwise.

### GetIsCustomerOk

`func (o *Contact) GetIsCustomerOk() (*bool, bool)`

GetIsCustomerOk returns a tuple with the IsCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomer

`func (o *Contact) SetIsCustomer(v bool)`

SetIsCustomer sets IsCustomer field to given value.

### HasIsCustomerField

`func (o *Contact) HasIsCustomerField() bool`

HasIsCustomerField returns a boolean if a field has been set.

### GetSalesDefaultLineAmountType

`func (o *Contact) GetSalesDefaultLineAmountType() string`

GetSalesDefaultLineAmountType returns the SalesDefaultLineAmountType field if non-nil, zero value otherwise.

### GetSalesDefaultLineAmountTypeOk

`func (o *Contact) GetSalesDefaultLineAmountTypeOk() (*string, bool)`

GetSalesDefaultLineAmountTypeOk returns a tuple with the SalesDefaultLineAmountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDefaultLineAmountType

`func (o *Contact) SetSalesDefaultLineAmountType(v string)`

SetSalesDefaultLineAmountType sets SalesDefaultLineAmountType field to given value.

### HasSalesDefaultLineAmountTypeField

`func (o *Contact) HasSalesDefaultLineAmountTypeField() bool`

HasSalesDefaultLineAmountTypeField returns a boolean if a field has been set.

### GetPurchasesDefaultLineAmountType

`func (o *Contact) GetPurchasesDefaultLineAmountType() string`

GetPurchasesDefaultLineAmountType returns the PurchasesDefaultLineAmountType field if non-nil, zero value otherwise.

### GetPurchasesDefaultLineAmountTypeOk

`func (o *Contact) GetPurchasesDefaultLineAmountTypeOk() (*string, bool)`

GetPurchasesDefaultLineAmountTypeOk returns a tuple with the PurchasesDefaultLineAmountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasesDefaultLineAmountType

`func (o *Contact) SetPurchasesDefaultLineAmountType(v string)`

SetPurchasesDefaultLineAmountType sets PurchasesDefaultLineAmountType field to given value.

### HasPurchasesDefaultLineAmountTypeField

`func (o *Contact) HasPurchasesDefaultLineAmountTypeField() bool`

HasPurchasesDefaultLineAmountTypeField returns a boolean if a field has been set.

### GetDefaultCurrency

`func (o *Contact) GetDefaultCurrency() CurrencyCode`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *Contact) GetDefaultCurrencyOk() (*CurrencyCode, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *Contact) SetDefaultCurrency(v CurrencyCode)`

SetDefaultCurrency sets DefaultCurrency field to given value.

### HasDefaultCurrencyField

`func (o *Contact) HasDefaultCurrencyField() bool`

HasDefaultCurrencyField returns a boolean if a field has been set.

### GetXeroNetworkKey

`func (o *Contact) GetXeroNetworkKey() string`

GetXeroNetworkKey returns the XeroNetworkKey field if non-nil, zero value otherwise.

### GetXeroNetworkKeyOk

`func (o *Contact) GetXeroNetworkKeyOk() (*string, bool)`

GetXeroNetworkKeyOk returns a tuple with the XeroNetworkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXeroNetworkKey

`func (o *Contact) SetXeroNetworkKey(v string)`

SetXeroNetworkKey sets XeroNetworkKey field to given value.

### HasXeroNetworkKeyField

`func (o *Contact) HasXeroNetworkKeyField() bool`

HasXeroNetworkKeyField returns a boolean if a field has been set.

### GetSalesDefaultAccountCode

`func (o *Contact) GetSalesDefaultAccountCode() string`

GetSalesDefaultAccountCode returns the SalesDefaultAccountCode field if non-nil, zero value otherwise.

### GetSalesDefaultAccountCodeOk

`func (o *Contact) GetSalesDefaultAccountCodeOk() (*string, bool)`

GetSalesDefaultAccountCodeOk returns a tuple with the SalesDefaultAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDefaultAccountCode

`func (o *Contact) SetSalesDefaultAccountCode(v string)`

SetSalesDefaultAccountCode sets SalesDefaultAccountCode field to given value.

### HasSalesDefaultAccountCodeField

`func (o *Contact) HasSalesDefaultAccountCodeField() bool`

HasSalesDefaultAccountCodeField returns a boolean if a field has been set.

### GetPurchasesDefaultAccountCode

`func (o *Contact) GetPurchasesDefaultAccountCode() string`

GetPurchasesDefaultAccountCode returns the PurchasesDefaultAccountCode field if non-nil, zero value otherwise.

### GetPurchasesDefaultAccountCodeOk

`func (o *Contact) GetPurchasesDefaultAccountCodeOk() (*string, bool)`

GetPurchasesDefaultAccountCodeOk returns a tuple with the PurchasesDefaultAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasesDefaultAccountCode

`func (o *Contact) SetPurchasesDefaultAccountCode(v string)`

SetPurchasesDefaultAccountCode sets PurchasesDefaultAccountCode field to given value.

### HasPurchasesDefaultAccountCodeField

`func (o *Contact) HasPurchasesDefaultAccountCodeField() bool`

HasPurchasesDefaultAccountCodeField returns a boolean if a field has been set.

### GetSalesTrackingCategories

`func (o *Contact) GetSalesTrackingCategories() []SalesTrackingCategory`

GetSalesTrackingCategories returns the SalesTrackingCategories field if non-nil, zero value otherwise.

### GetSalesTrackingCategoriesOk

`func (o *Contact) GetSalesTrackingCategoriesOk() (*[]SalesTrackingCategory, bool)`

GetSalesTrackingCategoriesOk returns a tuple with the SalesTrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTrackingCategories

`func (o *Contact) SetSalesTrackingCategories(v []SalesTrackingCategory)`

SetSalesTrackingCategories sets SalesTrackingCategories field to given value.

### HasSalesTrackingCategoriesField

`func (o *Contact) HasSalesTrackingCategoriesField() bool`

HasSalesTrackingCategoriesField returns a boolean if a field has been set.

### GetPurchasesTrackingCategories

`func (o *Contact) GetPurchasesTrackingCategories() []SalesTrackingCategory`

GetPurchasesTrackingCategories returns the PurchasesTrackingCategories field if non-nil, zero value otherwise.

### GetPurchasesTrackingCategoriesOk

`func (o *Contact) GetPurchasesTrackingCategoriesOk() (*[]SalesTrackingCategory, bool)`

GetPurchasesTrackingCategoriesOk returns a tuple with the PurchasesTrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasesTrackingCategories

`func (o *Contact) SetPurchasesTrackingCategories(v []SalesTrackingCategory)`

SetPurchasesTrackingCategories sets PurchasesTrackingCategories field to given value.

### HasPurchasesTrackingCategoriesField

`func (o *Contact) HasPurchasesTrackingCategoriesField() bool`

HasPurchasesTrackingCategoriesField returns a boolean if a field has been set.

### GetTrackingCategoryName

`func (o *Contact) GetTrackingCategoryName() string`

GetTrackingCategoryName returns the TrackingCategoryName field if non-nil, zero value otherwise.

### GetTrackingCategoryNameOk

`func (o *Contact) GetTrackingCategoryNameOk() (*string, bool)`

GetTrackingCategoryNameOk returns a tuple with the TrackingCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategoryName

`func (o *Contact) SetTrackingCategoryName(v string)`

SetTrackingCategoryName sets TrackingCategoryName field to given value.

### HasTrackingCategoryNameField

`func (o *Contact) HasTrackingCategoryNameField() bool`

HasTrackingCategoryNameField returns a boolean if a field has been set.

### GetTrackingCategoryOption

`func (o *Contact) GetTrackingCategoryOption() string`

GetTrackingCategoryOption returns the TrackingCategoryOption field if non-nil, zero value otherwise.

### GetTrackingCategoryOptionOk

`func (o *Contact) GetTrackingCategoryOptionOk() (*string, bool)`

GetTrackingCategoryOptionOk returns a tuple with the TrackingCategoryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategoryOption

`func (o *Contact) SetTrackingCategoryOption(v string)`

SetTrackingCategoryOption sets TrackingCategoryOption field to given value.

### HasTrackingCategoryOptionField

`func (o *Contact) HasTrackingCategoryOptionField() bool`

HasTrackingCategoryOptionField returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *Contact) GetPaymentTerms() PaymentTerm`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *Contact) GetPaymentTermsOk() (*PaymentTerm, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *Contact) SetPaymentTerms(v PaymentTerm)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTermsField

`func (o *Contact) HasPaymentTermsField() bool`

HasPaymentTermsField returns a boolean if a field has been set.

### GetUpdatedDateUTC

`func (o *Contact) GetUpdatedDateUTC() string`

GetUpdatedDateUTC returns the UpdatedDateUTC field if non-nil, zero value otherwise.

### GetUpdatedDateUTCOk

`func (o *Contact) GetUpdatedDateUTCOk() (*string, bool)`

GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateUTC

`func (o *Contact) SetUpdatedDateUTC(v string)`

SetUpdatedDateUTC sets UpdatedDateUTC field to given value.

### HasUpdatedDateUTCField

`func (o *Contact) HasUpdatedDateUTCField() bool`

HasUpdatedDateUTCField returns a boolean if a field has been set.

### GetContactGroups

`func (o *Contact) GetContactGroups() []ContactGroup`

GetContactGroups returns the ContactGroups field if non-nil, zero value otherwise.

### GetContactGroupsOk

`func (o *Contact) GetContactGroupsOk() (*[]ContactGroup, bool)`

GetContactGroupsOk returns a tuple with the ContactGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactGroups

`func (o *Contact) SetContactGroups(v []ContactGroup)`

SetContactGroups sets ContactGroups field to given value.

### HasContactGroupsField

`func (o *Contact) HasContactGroupsField() bool`

HasContactGroupsField returns a boolean if a field has been set.

### GetWebsite

`func (o *Contact) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Contact) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Contact) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsiteField

`func (o *Contact) HasWebsiteField() bool`

HasWebsiteField returns a boolean if a field has been set.

### GetBrandingTheme

`func (o *Contact) GetBrandingTheme() BrandingTheme`

GetBrandingTheme returns the BrandingTheme field if non-nil, zero value otherwise.

### GetBrandingThemeOk

`func (o *Contact) GetBrandingThemeOk() (*BrandingTheme, bool)`

GetBrandingThemeOk returns a tuple with the BrandingTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTheme

`func (o *Contact) SetBrandingTheme(v BrandingTheme)`

SetBrandingTheme sets BrandingTheme field to given value.

### HasBrandingThemeField

`func (o *Contact) HasBrandingThemeField() bool`

HasBrandingThemeField returns a boolean if a field has been set.

### GetBatchPayments

`func (o *Contact) GetBatchPayments() BatchPaymentDetails`

GetBatchPayments returns the BatchPayments field if non-nil, zero value otherwise.

### GetBatchPaymentsOk

`func (o *Contact) GetBatchPaymentsOk() (*BatchPaymentDetails, bool)`

GetBatchPaymentsOk returns a tuple with the BatchPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchPayments

`func (o *Contact) SetBatchPayments(v BatchPaymentDetails)`

SetBatchPayments sets BatchPayments field to given value.

### HasBatchPaymentsField

`func (o *Contact) HasBatchPaymentsField() bool`

HasBatchPaymentsField returns a boolean if a field has been set.

### GetDiscount

`func (o *Contact) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Contact) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Contact) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscountField

`func (o *Contact) HasDiscountField() bool`

HasDiscountField returns a boolean if a field has been set.

### GetBalances

`func (o *Contact) GetBalances() Balances`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *Contact) GetBalancesOk() (*Balances, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *Contact) SetBalances(v Balances)`

SetBalances sets Balances field to given value.

### HasBalancesField

`func (o *Contact) HasBalancesField() bool`

HasBalancesField returns a boolean if a field has been set.

### GetAttachments

`func (o *Contact) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Contact) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Contact) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachmentsField

`func (o *Contact) HasAttachmentsField() bool`

HasAttachmentsField returns a boolean if a field has been set.

### GetHasAttachments

`func (o *Contact) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Contact) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Contact) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachmentsField

`func (o *Contact) HasHasAttachmentsField() bool`

HasHasAttachmentsField returns a boolean if a field has been set.

### GetValidationErrors

`func (o *Contact) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Contact) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Contact) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrorsField

`func (o *Contact) HasValidationErrorsField() bool`

HasValidationErrorsField returns a boolean if a field has been set.

### GetHasValidationErrors

`func (o *Contact) GetHasValidationErrors() bool`

GetHasValidationErrors returns the HasValidationErrors field if non-nil, zero value otherwise.

### GetHasValidationErrorsOk

`func (o *Contact) GetHasValidationErrorsOk() (*bool, bool)`

GetHasValidationErrorsOk returns a tuple with the HasValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValidationErrors

`func (o *Contact) SetHasValidationErrors(v bool)`

SetHasValidationErrors sets HasValidationErrors field to given value.

### HasHasValidationErrorsField

`func (o *Contact) HasHasValidationErrorsField() bool`

HasHasValidationErrorsField returns a boolean if a field has been set.

### GetStatusAttributeString

`func (o *Contact) GetStatusAttributeString() string`

GetStatusAttributeString returns the StatusAttributeString field if non-nil, zero value otherwise.

### GetStatusAttributeStringOk

`func (o *Contact) GetStatusAttributeStringOk() (*string, bool)`

GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAttributeString

`func (o *Contact) SetStatusAttributeString(v string)`

SetStatusAttributeString sets StatusAttributeString field to given value.

### HasStatusAttributeStringField

`func (o *Contact) HasStatusAttributeStringField() bool`

HasStatusAttributeStringField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


