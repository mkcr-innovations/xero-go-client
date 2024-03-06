# Organisation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganisationID** | Pointer to **string** | Unique Xero identifier | [optional] 
**APIKey** | Pointer to **string** | Display a unique key used for Xero-to-Xero transactions | [optional] 
**Name** | Pointer to **string** | Display name of organisation shown in Xero | [optional] 
**LegalName** | Pointer to **string** | Organisation name shown on Reports | [optional] 
**PaysTax** | Pointer to **bool** | Boolean to describe if organisation is registered with a local tax authority i.e. true, false | [optional] 
**Version** | Pointer to **string** | See Version Types | [optional] 
**OrganisationType** | Pointer to **string** | Organisation Type | [optional] 
**BaseCurrency** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**CountryCode** | Pointer to [**CountryCode**](CountryCode.md) |  | [optional] 
**IsDemoCompany** | Pointer to **bool** | Boolean to describe if organisation is a demo company. | [optional] 
**OrganisationStatus** | Pointer to **string** | Will be set to ACTIVE if you can connect to organisation via the Xero API | [optional] 
**RegistrationNumber** | Pointer to **string** | Shows for New Zealand, Australian and UK organisations | [optional] 
**EmployerIdentificationNumber** | Pointer to **string** | Shown if set. US Only. | [optional] 
**TaxNumber** | Pointer to **string** | Shown if set. Displays in the Xero UI as Tax File Number (AU), GST Number (NZ), VAT Number (UK) and Tax ID Number (US &amp; Global). | [optional] 
**FinancialYearEndDay** | Pointer to **int32** | Calendar day e.g. 0-31 | [optional] 
**FinancialYearEndMonth** | Pointer to **int32** | Calendar Month e.g. 1-12 | [optional] 
**SalesTaxBasis** | Pointer to **string** | The accounting basis used for tax returns. See Sales Tax Basis | [optional] 
**SalesTaxPeriod** | Pointer to **string** | The frequency with which tax returns are processed. See Sales Tax Period | [optional] 
**DefaultSalesTax** | Pointer to **string** | The default for LineAmountTypes on sales transactions | [optional] 
**DefaultPurchasesTax** | Pointer to **string** | The default for LineAmountTypes on purchase transactions | [optional] 
**PeriodLockDate** | Pointer to **string** | Shown if set. See lock dates | [optional] 
**EndOfYearLockDate** | Pointer to **string** | Shown if set. See lock dates | [optional] 
**CreatedDateUTC** | Pointer to **string** | Timestamp when the organisation was created in Xero | [optional] [readonly] 
**Timezone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**OrganisationEntityType** | Pointer to **string** | Organisation Entity Type | [optional] 
**ShortCode** | Pointer to **string** | A unique identifier for the organisation. Potential uses. | [optional] 
**Class** | Pointer to **string** | Organisation Classes describe which plan the Xero organisation is on (e.g. DEMO, TRIAL, PREMIUM) | [optional] 
**Edition** | Pointer to **string** | BUSINESS or PARTNER. Partner edition organisations are sold exclusively through accounting partners and have restricted functionality (e.g. no access to invoicing) | [optional] 
**LineOfBusiness** | Pointer to **string** | Description of business type as defined in Organisation settings | [optional] 
**Addresses** | Pointer to [**[]AddressForOrganisation**](AddressForOrganisation.md) | Address details for organisation – see Addresses | [optional] 
**Phones** | Pointer to [**[]Phone**](Phone.md) | Phones details for organisation – see Phones | [optional] 
**ExternalLinks** | Pointer to [**[]ExternalLink**](ExternalLink.md) | Organisation profile links for popular services such as Facebook,Twitter, GooglePlus and LinkedIn. You can also add link to your website here. Shown if Organisation settings  is updated in Xero. See ExternalLinks below | [optional] 
**PaymentTerms** | Pointer to [**PaymentTerm**](PaymentTerm.md) |  | [optional] 

## Methods

### NewOrganisation

`func NewOrganisation() *Organisation`

NewOrganisation instantiates a new Organisation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationWithDefaults

`func NewOrganisationWithDefaults() *Organisation`

NewOrganisationWithDefaults instantiates a new Organisation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganisationID

`func (o *Organisation) GetOrganisationID() string`

GetOrganisationID returns the OrganisationID field if non-nil, zero value otherwise.

### GetOrganisationIDOk

`func (o *Organisation) GetOrganisationIDOk() (*string, bool)`

GetOrganisationIDOk returns a tuple with the OrganisationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationID

`func (o *Organisation) SetOrganisationID(v string)`

SetOrganisationID sets OrganisationID field to given value.

### HasOrganisationIDField

`func (o *Organisation) HasOrganisationIDField() bool`

HasOrganisationIDField returns a boolean if a field has been set.

### GetAPIKey

`func (o *Organisation) GetAPIKey() string`

GetAPIKey returns the APIKey field if non-nil, zero value otherwise.

### GetAPIKeyOk

`func (o *Organisation) GetAPIKeyOk() (*string, bool)`

GetAPIKeyOk returns a tuple with the APIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIKey

`func (o *Organisation) SetAPIKey(v string)`

SetAPIKey sets APIKey field to given value.

### HasAPIKeyField

`func (o *Organisation) HasAPIKeyField() bool`

HasAPIKeyField returns a boolean if a field has been set.

### GetName

`func (o *Organisation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organisation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organisation) SetName(v string)`

SetName sets Name field to given value.

### HasNameField

`func (o *Organisation) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetLegalName

`func (o *Organisation) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *Organisation) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *Organisation) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalNameField

`func (o *Organisation) HasLegalNameField() bool`

HasLegalNameField returns a boolean if a field has been set.

### GetPaysTax

`func (o *Organisation) GetPaysTax() bool`

GetPaysTax returns the PaysTax field if non-nil, zero value otherwise.

### GetPaysTaxOk

`func (o *Organisation) GetPaysTaxOk() (*bool, bool)`

GetPaysTaxOk returns a tuple with the PaysTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysTax

`func (o *Organisation) SetPaysTax(v bool)`

SetPaysTax sets PaysTax field to given value.

### HasPaysTaxField

`func (o *Organisation) HasPaysTaxField() bool`

HasPaysTaxField returns a boolean if a field has been set.

### GetVersion

`func (o *Organisation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Organisation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Organisation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersionField

`func (o *Organisation) HasVersionField() bool`

HasVersionField returns a boolean if a field has been set.

### GetOrganisationType

`func (o *Organisation) GetOrganisationType() string`

GetOrganisationType returns the OrganisationType field if non-nil, zero value otherwise.

### GetOrganisationTypeOk

`func (o *Organisation) GetOrganisationTypeOk() (*string, bool)`

GetOrganisationTypeOk returns a tuple with the OrganisationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationType

`func (o *Organisation) SetOrganisationType(v string)`

SetOrganisationType sets OrganisationType field to given value.

### HasOrganisationTypeField

`func (o *Organisation) HasOrganisationTypeField() bool`

HasOrganisationTypeField returns a boolean if a field has been set.

### GetBaseCurrency

`func (o *Organisation) GetBaseCurrency() CurrencyCode`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *Organisation) GetBaseCurrencyOk() (*CurrencyCode, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *Organisation) SetBaseCurrency(v CurrencyCode)`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrencyField

`func (o *Organisation) HasBaseCurrencyField() bool`

HasBaseCurrencyField returns a boolean if a field has been set.

### GetCountryCode

`func (o *Organisation) GetCountryCode() CountryCode`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Organisation) GetCountryCodeOk() (*CountryCode, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Organisation) SetCountryCode(v CountryCode)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCodeField

`func (o *Organisation) HasCountryCodeField() bool`

HasCountryCodeField returns a boolean if a field has been set.

### GetIsDemoCompany

`func (o *Organisation) GetIsDemoCompany() bool`

GetIsDemoCompany returns the IsDemoCompany field if non-nil, zero value otherwise.

### GetIsDemoCompanyOk

`func (o *Organisation) GetIsDemoCompanyOk() (*bool, bool)`

GetIsDemoCompanyOk returns a tuple with the IsDemoCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemoCompany

`func (o *Organisation) SetIsDemoCompany(v bool)`

SetIsDemoCompany sets IsDemoCompany field to given value.

### HasIsDemoCompanyField

`func (o *Organisation) HasIsDemoCompanyField() bool`

HasIsDemoCompanyField returns a boolean if a field has been set.

### GetOrganisationStatus

`func (o *Organisation) GetOrganisationStatus() string`

GetOrganisationStatus returns the OrganisationStatus field if non-nil, zero value otherwise.

### GetOrganisationStatusOk

`func (o *Organisation) GetOrganisationStatusOk() (*string, bool)`

GetOrganisationStatusOk returns a tuple with the OrganisationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationStatus

`func (o *Organisation) SetOrganisationStatus(v string)`

SetOrganisationStatus sets OrganisationStatus field to given value.

### HasOrganisationStatusField

`func (o *Organisation) HasOrganisationStatusField() bool`

HasOrganisationStatusField returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *Organisation) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *Organisation) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *Organisation) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumberField

`func (o *Organisation) HasRegistrationNumberField() bool`

HasRegistrationNumberField returns a boolean if a field has been set.

### GetEmployerIdentificationNumber

`func (o *Organisation) GetEmployerIdentificationNumber() string`

GetEmployerIdentificationNumber returns the EmployerIdentificationNumber field if non-nil, zero value otherwise.

### GetEmployerIdentificationNumberOk

`func (o *Organisation) GetEmployerIdentificationNumberOk() (*string, bool)`

GetEmployerIdentificationNumberOk returns a tuple with the EmployerIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerIdentificationNumber

`func (o *Organisation) SetEmployerIdentificationNumber(v string)`

SetEmployerIdentificationNumber sets EmployerIdentificationNumber field to given value.

### HasEmployerIdentificationNumberField

`func (o *Organisation) HasEmployerIdentificationNumberField() bool`

HasEmployerIdentificationNumberField returns a boolean if a field has been set.

### GetTaxNumber

`func (o *Organisation) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *Organisation) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *Organisation) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumberField

`func (o *Organisation) HasTaxNumberField() bool`

HasTaxNumberField returns a boolean if a field has been set.

### GetFinancialYearEndDay

`func (o *Organisation) GetFinancialYearEndDay() int32`

GetFinancialYearEndDay returns the FinancialYearEndDay field if non-nil, zero value otherwise.

### GetFinancialYearEndDayOk

`func (o *Organisation) GetFinancialYearEndDayOk() (*int32, bool)`

GetFinancialYearEndDayOk returns a tuple with the FinancialYearEndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialYearEndDay

`func (o *Organisation) SetFinancialYearEndDay(v int32)`

SetFinancialYearEndDay sets FinancialYearEndDay field to given value.

### HasFinancialYearEndDayField

`func (o *Organisation) HasFinancialYearEndDayField() bool`

HasFinancialYearEndDayField returns a boolean if a field has been set.

### GetFinancialYearEndMonth

`func (o *Organisation) GetFinancialYearEndMonth() int32`

GetFinancialYearEndMonth returns the FinancialYearEndMonth field if non-nil, zero value otherwise.

### GetFinancialYearEndMonthOk

`func (o *Organisation) GetFinancialYearEndMonthOk() (*int32, bool)`

GetFinancialYearEndMonthOk returns a tuple with the FinancialYearEndMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialYearEndMonth

`func (o *Organisation) SetFinancialYearEndMonth(v int32)`

SetFinancialYearEndMonth sets FinancialYearEndMonth field to given value.

### HasFinancialYearEndMonthField

`func (o *Organisation) HasFinancialYearEndMonthField() bool`

HasFinancialYearEndMonthField returns a boolean if a field has been set.

### GetSalesTaxBasis

`func (o *Organisation) GetSalesTaxBasis() string`

GetSalesTaxBasis returns the SalesTaxBasis field if non-nil, zero value otherwise.

### GetSalesTaxBasisOk

`func (o *Organisation) GetSalesTaxBasisOk() (*string, bool)`

GetSalesTaxBasisOk returns a tuple with the SalesTaxBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxBasis

`func (o *Organisation) SetSalesTaxBasis(v string)`

SetSalesTaxBasis sets SalesTaxBasis field to given value.

### HasSalesTaxBasisField

`func (o *Organisation) HasSalesTaxBasisField() bool`

HasSalesTaxBasisField returns a boolean if a field has been set.

### GetSalesTaxPeriod

`func (o *Organisation) GetSalesTaxPeriod() string`

GetSalesTaxPeriod returns the SalesTaxPeriod field if non-nil, zero value otherwise.

### GetSalesTaxPeriodOk

`func (o *Organisation) GetSalesTaxPeriodOk() (*string, bool)`

GetSalesTaxPeriodOk returns a tuple with the SalesTaxPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxPeriod

`func (o *Organisation) SetSalesTaxPeriod(v string)`

SetSalesTaxPeriod sets SalesTaxPeriod field to given value.

### HasSalesTaxPeriodField

`func (o *Organisation) HasSalesTaxPeriodField() bool`

HasSalesTaxPeriodField returns a boolean if a field has been set.

### GetDefaultSalesTax

`func (o *Organisation) GetDefaultSalesTax() string`

GetDefaultSalesTax returns the DefaultSalesTax field if non-nil, zero value otherwise.

### GetDefaultSalesTaxOk

`func (o *Organisation) GetDefaultSalesTaxOk() (*string, bool)`

GetDefaultSalesTaxOk returns a tuple with the DefaultSalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSalesTax

`func (o *Organisation) SetDefaultSalesTax(v string)`

SetDefaultSalesTax sets DefaultSalesTax field to given value.

### HasDefaultSalesTaxField

`func (o *Organisation) HasDefaultSalesTaxField() bool`

HasDefaultSalesTaxField returns a boolean if a field has been set.

### GetDefaultPurchasesTax

`func (o *Organisation) GetDefaultPurchasesTax() string`

GetDefaultPurchasesTax returns the DefaultPurchasesTax field if non-nil, zero value otherwise.

### GetDefaultPurchasesTaxOk

`func (o *Organisation) GetDefaultPurchasesTaxOk() (*string, bool)`

GetDefaultPurchasesTaxOk returns a tuple with the DefaultPurchasesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPurchasesTax

`func (o *Organisation) SetDefaultPurchasesTax(v string)`

SetDefaultPurchasesTax sets DefaultPurchasesTax field to given value.

### HasDefaultPurchasesTaxField

`func (o *Organisation) HasDefaultPurchasesTaxField() bool`

HasDefaultPurchasesTaxField returns a boolean if a field has been set.

### GetPeriodLockDate

`func (o *Organisation) GetPeriodLockDate() string`

GetPeriodLockDate returns the PeriodLockDate field if non-nil, zero value otherwise.

### GetPeriodLockDateOk

`func (o *Organisation) GetPeriodLockDateOk() (*string, bool)`

GetPeriodLockDateOk returns a tuple with the PeriodLockDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodLockDate

`func (o *Organisation) SetPeriodLockDate(v string)`

SetPeriodLockDate sets PeriodLockDate field to given value.

### HasPeriodLockDateField

`func (o *Organisation) HasPeriodLockDateField() bool`

HasPeriodLockDateField returns a boolean if a field has been set.

### GetEndOfYearLockDate

`func (o *Organisation) GetEndOfYearLockDate() string`

GetEndOfYearLockDate returns the EndOfYearLockDate field if non-nil, zero value otherwise.

### GetEndOfYearLockDateOk

`func (o *Organisation) GetEndOfYearLockDateOk() (*string, bool)`

GetEndOfYearLockDateOk returns a tuple with the EndOfYearLockDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfYearLockDate

`func (o *Organisation) SetEndOfYearLockDate(v string)`

SetEndOfYearLockDate sets EndOfYearLockDate field to given value.

### HasEndOfYearLockDateField

`func (o *Organisation) HasEndOfYearLockDateField() bool`

HasEndOfYearLockDateField returns a boolean if a field has been set.

### GetCreatedDateUTC

`func (o *Organisation) GetCreatedDateUTC() string`

GetCreatedDateUTC returns the CreatedDateUTC field if non-nil, zero value otherwise.

### GetCreatedDateUTCOk

`func (o *Organisation) GetCreatedDateUTCOk() (*string, bool)`

GetCreatedDateUTCOk returns a tuple with the CreatedDateUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateUTC

`func (o *Organisation) SetCreatedDateUTC(v string)`

SetCreatedDateUTC sets CreatedDateUTC field to given value.

### HasCreatedDateUTCField

`func (o *Organisation) HasCreatedDateUTCField() bool`

HasCreatedDateUTCField returns a boolean if a field has been set.

### GetTimezone

`func (o *Organisation) GetTimezone() TimeZone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Organisation) GetTimezoneOk() (*TimeZone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Organisation) SetTimezone(v TimeZone)`

SetTimezone sets Timezone field to given value.

### HasTimezoneField

`func (o *Organisation) HasTimezoneField() bool`

HasTimezoneField returns a boolean if a field has been set.

### GetOrganisationEntityType

`func (o *Organisation) GetOrganisationEntityType() string`

GetOrganisationEntityType returns the OrganisationEntityType field if non-nil, zero value otherwise.

### GetOrganisationEntityTypeOk

`func (o *Organisation) GetOrganisationEntityTypeOk() (*string, bool)`

GetOrganisationEntityTypeOk returns a tuple with the OrganisationEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationEntityType

`func (o *Organisation) SetOrganisationEntityType(v string)`

SetOrganisationEntityType sets OrganisationEntityType field to given value.

### HasOrganisationEntityTypeField

`func (o *Organisation) HasOrganisationEntityTypeField() bool`

HasOrganisationEntityTypeField returns a boolean if a field has been set.

### GetShortCode

`func (o *Organisation) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *Organisation) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *Organisation) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCodeField

`func (o *Organisation) HasShortCodeField() bool`

HasShortCodeField returns a boolean if a field has been set.

### GetClass

`func (o *Organisation) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Organisation) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Organisation) SetClass(v string)`

SetClass sets Class field to given value.

### HasClassField

`func (o *Organisation) HasClassField() bool`

HasClassField returns a boolean if a field has been set.

### GetEdition

`func (o *Organisation) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *Organisation) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *Organisation) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEditionField

`func (o *Organisation) HasEditionField() bool`

HasEditionField returns a boolean if a field has been set.

### GetLineOfBusiness

`func (o *Organisation) GetLineOfBusiness() string`

GetLineOfBusiness returns the LineOfBusiness field if non-nil, zero value otherwise.

### GetLineOfBusinessOk

`func (o *Organisation) GetLineOfBusinessOk() (*string, bool)`

GetLineOfBusinessOk returns a tuple with the LineOfBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineOfBusiness

`func (o *Organisation) SetLineOfBusiness(v string)`

SetLineOfBusiness sets LineOfBusiness field to given value.

### HasLineOfBusinessField

`func (o *Organisation) HasLineOfBusinessField() bool`

HasLineOfBusinessField returns a boolean if a field has been set.

### GetAddresses

`func (o *Organisation) GetAddresses() []AddressForOrganisation`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Organisation) GetAddressesOk() (*[]AddressForOrganisation, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Organisation) SetAddresses(v []AddressForOrganisation)`

SetAddresses sets Addresses field to given value.

### HasAddressesField

`func (o *Organisation) HasAddressesField() bool`

HasAddressesField returns a boolean if a field has been set.

### GetPhones

`func (o *Organisation) GetPhones() []Phone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *Organisation) GetPhonesOk() (*[]Phone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *Organisation) SetPhones(v []Phone)`

SetPhones sets Phones field to given value.

### HasPhonesField

`func (o *Organisation) HasPhonesField() bool`

HasPhonesField returns a boolean if a field has been set.

### GetExternalLinks

`func (o *Organisation) GetExternalLinks() []ExternalLink`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *Organisation) GetExternalLinksOk() (*[]ExternalLink, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *Organisation) SetExternalLinks(v []ExternalLink)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinksField

`func (o *Organisation) HasExternalLinksField() bool`

HasExternalLinksField returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *Organisation) GetPaymentTerms() PaymentTerm`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *Organisation) GetPaymentTermsOk() (*PaymentTerm, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *Organisation) SetPaymentTerms(v PaymentTerm)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTermsField

`func (o *Organisation) HasPaymentTermsField() bool`

HasPaymentTermsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


