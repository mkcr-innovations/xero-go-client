# \AccountingAPI

All URIs are relative to *https://api.xero.com/api.xro/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountingAPI.md#CreateAccount) | **Put** /Accounts | Creates a new chart of accounts
[**CreateAccountAttachmentByFileName**](AccountingAPI.md#CreateAccountAttachmentByFileName) | **Put** /Accounts/{AccountID}/Attachments/{FileName} | Creates an attachment on a specific account
[**CreateBankTransactionAttachmentByFileName**](AccountingAPI.md#CreateBankTransactionAttachmentByFileName) | **Put** /BankTransactions/{BankTransactionID}/Attachments/{FileName} | Creates an attachment for a specific bank transaction by filename
[**CreateBankTransactionHistoryRecord**](AccountingAPI.md#CreateBankTransactionHistoryRecord) | **Put** /BankTransactions/{BankTransactionID}/History | Creates a history record for a specific bank transactions
[**CreateBankTransactions**](AccountingAPI.md#CreateBankTransactions) | **Put** /BankTransactions | Creates one or more spent or received money transaction
[**CreateBankTransfer**](AccountingAPI.md#CreateBankTransfer) | **Put** /BankTransfers | Creates a bank transfer
[**CreateBankTransferAttachmentByFileName**](AccountingAPI.md#CreateBankTransferAttachmentByFileName) | **Put** /BankTransfers/{BankTransferID}/Attachments/{FileName} | 
[**CreateBankTransferHistoryRecord**](AccountingAPI.md#CreateBankTransferHistoryRecord) | **Put** /BankTransfers/{BankTransferID}/History | Creates a history record for a specific bank transfer
[**CreateBatchPayment**](AccountingAPI.md#CreateBatchPayment) | **Put** /BatchPayments | Creates one or many batch payments for invoices
[**CreateBatchPaymentHistoryRecord**](AccountingAPI.md#CreateBatchPaymentHistoryRecord) | **Put** /BatchPayments/{BatchPaymentID}/History | Creates a history record for a specific batch payment
[**CreateBrandingThemePaymentServices**](AccountingAPI.md#CreateBrandingThemePaymentServices) | **Post** /BrandingThemes/{BrandingThemeID}/PaymentServices | Creates a new custom payment service for a specific branding theme
[**CreateContactAttachmentByFileName**](AccountingAPI.md#CreateContactAttachmentByFileName) | **Put** /Contacts/{ContactID}/Attachments/{FileName} | 
[**CreateContactGroup**](AccountingAPI.md#CreateContactGroup) | **Put** /ContactGroups | Creates a contact group
[**CreateContactGroupContacts**](AccountingAPI.md#CreateContactGroupContacts) | **Put** /ContactGroups/{ContactGroupID}/Contacts | Creates contacts to a specific contact group
[**CreateContactHistory**](AccountingAPI.md#CreateContactHistory) | **Put** /Contacts/{ContactID}/History | Creates a new history record for a specific contact
[**CreateContacts**](AccountingAPI.md#CreateContacts) | **Put** /Contacts | Creates multiple contacts (bulk) in a Xero organisation
[**CreateCreditNoteAllocation**](AccountingAPI.md#CreateCreditNoteAllocation) | **Put** /CreditNotes/{CreditNoteID}/Allocations | Creates allocation for a specific credit note
[**CreateCreditNoteAttachmentByFileName**](AccountingAPI.md#CreateCreditNoteAttachmentByFileName) | **Put** /CreditNotes/{CreditNoteID}/Attachments/{FileName} | Creates an attachment for a specific credit note
[**CreateCreditNoteHistory**](AccountingAPI.md#CreateCreditNoteHistory) | **Put** /CreditNotes/{CreditNoteID}/History | Retrieves history records of a specific credit note
[**CreateCreditNotes**](AccountingAPI.md#CreateCreditNotes) | **Put** /CreditNotes | Creates a new credit note
[**CreateCurrency**](AccountingAPI.md#CreateCurrency) | **Put** /Currencies | Create a new currency for a Xero organisation
[**CreateEmployees**](AccountingAPI.md#CreateEmployees) | **Put** /Employees | Creates new employees used in Xero payrun
[**CreateExpenseClaimHistory**](AccountingAPI.md#CreateExpenseClaimHistory) | **Put** /ExpenseClaims/{ExpenseClaimID}/History | Creates a history record for a specific expense claim
[**CreateExpenseClaims**](AccountingAPI.md#CreateExpenseClaims) | **Put** /ExpenseClaims | Creates expense claims
[**CreateInvoiceAttachmentByFileName**](AccountingAPI.md#CreateInvoiceAttachmentByFileName) | **Put** /Invoices/{InvoiceID}/Attachments/{FileName} | Creates an attachment for a specific invoice or purchase bill by filename
[**CreateInvoiceHistory**](AccountingAPI.md#CreateInvoiceHistory) | **Put** /Invoices/{InvoiceID}/History | Creates a history record for a specific invoice
[**CreateInvoices**](AccountingAPI.md#CreateInvoices) | **Put** /Invoices | Creates one or more sales invoices or purchase bills
[**CreateItemHistory**](AccountingAPI.md#CreateItemHistory) | **Put** /Items/{ItemID}/History | Creates a history record for a specific item
[**CreateItems**](AccountingAPI.md#CreateItems) | **Put** /Items | Creates one or more items
[**CreateLinkedTransaction**](AccountingAPI.md#CreateLinkedTransaction) | **Put** /LinkedTransactions | Creates linked transactions (billable expenses)
[**CreateManualJournalAttachmentByFileName**](AccountingAPI.md#CreateManualJournalAttachmentByFileName) | **Put** /ManualJournals/{ManualJournalID}/Attachments/{FileName} | Creates a specific attachment for a specific manual journal by file name
[**CreateManualJournalHistoryRecord**](AccountingAPI.md#CreateManualJournalHistoryRecord) | **Put** /ManualJournals/{ManualJournalID}/History | Creates a history record for a specific manual journal
[**CreateManualJournals**](AccountingAPI.md#CreateManualJournals) | **Put** /ManualJournals | Creates one or more manual journals
[**CreateOverpaymentAllocations**](AccountingAPI.md#CreateOverpaymentAllocations) | **Put** /Overpayments/{OverpaymentID}/Allocations | Creates a single allocation for a specific overpayment
[**CreateOverpaymentHistory**](AccountingAPI.md#CreateOverpaymentHistory) | **Put** /Overpayments/{OverpaymentID}/History | Creates a history record for a specific overpayment
[**CreatePayment**](AccountingAPI.md#CreatePayment) | **Post** /Payments | Creates a single payment for invoice or credit notes
[**CreatePaymentHistory**](AccountingAPI.md#CreatePaymentHistory) | **Put** /Payments/{PaymentID}/History | Creates a history record for a specific payment
[**CreatePaymentService**](AccountingAPI.md#CreatePaymentService) | **Put** /PaymentServices | Creates a payment service
[**CreatePayments**](AccountingAPI.md#CreatePayments) | **Put** /Payments | Creates multiple payments for invoices or credit notes
[**CreatePrepaymentAllocations**](AccountingAPI.md#CreatePrepaymentAllocations) | **Put** /Prepayments/{PrepaymentID}/Allocations | Allows you to create an Allocation for prepayments
[**CreatePrepaymentHistory**](AccountingAPI.md#CreatePrepaymentHistory) | **Put** /Prepayments/{PrepaymentID}/History | Creates a history record for a specific prepayment
[**CreatePurchaseOrderAttachmentByFileName**](AccountingAPI.md#CreatePurchaseOrderAttachmentByFileName) | **Put** /PurchaseOrders/{PurchaseOrderID}/Attachments/{FileName} | Creates attachment for a specific purchase order
[**CreatePurchaseOrderHistory**](AccountingAPI.md#CreatePurchaseOrderHistory) | **Put** /PurchaseOrders/{PurchaseOrderID}/History | Creates a history record for a specific purchase orders
[**CreatePurchaseOrders**](AccountingAPI.md#CreatePurchaseOrders) | **Put** /PurchaseOrders | Creates one or more purchase orders
[**CreateQuoteAttachmentByFileName**](AccountingAPI.md#CreateQuoteAttachmentByFileName) | **Put** /Quotes/{QuoteID}/Attachments/{FileName} | Creates attachment for a specific quote
[**CreateQuoteHistory**](AccountingAPI.md#CreateQuoteHistory) | **Put** /Quotes/{QuoteID}/History | Creates a history record for a specific quote
[**CreateQuotes**](AccountingAPI.md#CreateQuotes) | **Put** /Quotes | Create one or more quotes
[**CreateReceipt**](AccountingAPI.md#CreateReceipt) | **Put** /Receipts | Creates draft expense claim receipts for any user
[**CreateReceiptAttachmentByFileName**](AccountingAPI.md#CreateReceiptAttachmentByFileName) | **Put** /Receipts/{ReceiptID}/Attachments/{FileName} | Creates an attachment on a specific expense claim receipts by file name
[**CreateReceiptHistory**](AccountingAPI.md#CreateReceiptHistory) | **Put** /Receipts/{ReceiptID}/History | Creates a history record for a specific receipt
[**CreateRepeatingInvoiceAttachmentByFileName**](AccountingAPI.md#CreateRepeatingInvoiceAttachmentByFileName) | **Put** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments/{FileName} | Creates an attachment from a specific repeating invoices by file name
[**CreateRepeatingInvoiceHistory**](AccountingAPI.md#CreateRepeatingInvoiceHistory) | **Put** /RepeatingInvoices/{RepeatingInvoiceID}/History | Creates a  history record for a specific repeating invoice
[**CreateRepeatingInvoices**](AccountingAPI.md#CreateRepeatingInvoices) | **Put** /RepeatingInvoices | Creates one or more repeating invoice templates
[**CreateTaxRates**](AccountingAPI.md#CreateTaxRates) | **Put** /TaxRates | Creates one or more tax rates
[**CreateTrackingCategory**](AccountingAPI.md#CreateTrackingCategory) | **Put** /TrackingCategories | Create tracking categories
[**CreateTrackingOptions**](AccountingAPI.md#CreateTrackingOptions) | **Put** /TrackingCategories/{TrackingCategoryID}/Options | Creates options for a specific tracking category
[**DeleteAccount**](AccountingAPI.md#DeleteAccount) | **Delete** /Accounts/{AccountID} | Deletes a chart of accounts
[**DeleteBatchPayment**](AccountingAPI.md#DeleteBatchPayment) | **Post** /BatchPayments | Updates a specific batch payment for invoices and credit notes
[**DeleteBatchPaymentByUrlParam**](AccountingAPI.md#DeleteBatchPaymentByUrlParam) | **Post** /BatchPayments/{BatchPaymentID} | Updates a specific batch payment for invoices and credit notes
[**DeleteContactGroupContact**](AccountingAPI.md#DeleteContactGroupContact) | **Delete** /ContactGroups/{ContactGroupID}/Contacts/{ContactID} | Deletes a specific contact from a contact group using a unique contact Id
[**DeleteContactGroupContacts**](AccountingAPI.md#DeleteContactGroupContacts) | **Delete** /ContactGroups/{ContactGroupID}/Contacts | Deletes all contacts from a specific contact group
[**DeleteCreditNoteAllocations**](AccountingAPI.md#DeleteCreditNoteAllocations) | **Delete** /CreditNotes/{CreditNoteID}/Allocations/{AllocationID} | Deletes an Allocation from a Credit Note
[**DeleteItem**](AccountingAPI.md#DeleteItem) | **Delete** /Items/{ItemID} | Deletes a specific item
[**DeleteLinkedTransaction**](AccountingAPI.md#DeleteLinkedTransaction) | **Delete** /LinkedTransactions/{LinkedTransactionID} | Deletes a specific linked transactions (billable expenses)
[**DeleteOverpaymentAllocations**](AccountingAPI.md#DeleteOverpaymentAllocations) | **Delete** /Overpayments/{OverpaymentID}/Allocations/{AllocationID} | Deletes an Allocation from an overpayment
[**DeletePayment**](AccountingAPI.md#DeletePayment) | **Post** /Payments/{PaymentID} | Updates a specific payment for invoices and credit notes
[**DeletePrepaymentAllocations**](AccountingAPI.md#DeletePrepaymentAllocations) | **Delete** /Prepayments/{PrepaymentID}/Allocations/{AllocationID} | Deletes an Allocation from a Prepayment
[**DeleteTrackingCategory**](AccountingAPI.md#DeleteTrackingCategory) | **Delete** /TrackingCategories/{TrackingCategoryID} | Deletes a specific tracking category
[**DeleteTrackingOptions**](AccountingAPI.md#DeleteTrackingOptions) | **Delete** /TrackingCategories/{TrackingCategoryID}/Options/{TrackingOptionID} | Deletes a specific option for a specific tracking category
[**EmailInvoice**](AccountingAPI.md#EmailInvoice) | **Post** /Invoices/{InvoiceID}/Email | Sends a copy of a specific invoice to related contact via email
[**GetAccount**](AccountingAPI.md#GetAccount) | **Get** /Accounts/{AccountID} | Retrieves a single chart of accounts by using a unique account Id
[**GetAccountAttachmentByFileName**](AccountingAPI.md#GetAccountAttachmentByFileName) | **Get** /Accounts/{AccountID}/Attachments/{FileName} | Retrieves an attachment for a specific account by filename
[**GetAccountAttachmentById**](AccountingAPI.md#GetAccountAttachmentById) | **Get** /Accounts/{AccountID}/Attachments/{AttachmentID} | Retrieves a specific attachment from a specific account using a unique attachment Id
[**GetAccountAttachments**](AccountingAPI.md#GetAccountAttachments) | **Get** /Accounts/{AccountID}/Attachments | Retrieves attachments for a specific accounts by using a unique account Id
[**GetAccounts**](AccountingAPI.md#GetAccounts) | **Get** /Accounts | Retrieves the full chart of accounts
[**GetBankTransaction**](AccountingAPI.md#GetBankTransaction) | **Get** /BankTransactions/{BankTransactionID} | Retrieves a single spent or received money transaction by using a unique bank transaction Id
[**GetBankTransactionAttachmentByFileName**](AccountingAPI.md#GetBankTransactionAttachmentByFileName) | **Get** /BankTransactions/{BankTransactionID}/Attachments/{FileName} | Retrieves a specific attachment from a specific bank transaction by filename
[**GetBankTransactionAttachmentById**](AccountingAPI.md#GetBankTransactionAttachmentById) | **Get** /BankTransactions/{BankTransactionID}/Attachments/{AttachmentID} | Retrieves specific attachments from a specific BankTransaction using a unique attachment Id
[**GetBankTransactionAttachments**](AccountingAPI.md#GetBankTransactionAttachments) | **Get** /BankTransactions/{BankTransactionID}/Attachments | Retrieves any attachments from a specific bank transactions
[**GetBankTransactions**](AccountingAPI.md#GetBankTransactions) | **Get** /BankTransactions | Retrieves any spent or received money transactions
[**GetBankTransactionsHistory**](AccountingAPI.md#GetBankTransactionsHistory) | **Get** /BankTransactions/{BankTransactionID}/History | Retrieves history from a specific bank transaction using a unique bank transaction Id
[**GetBankTransfer**](AccountingAPI.md#GetBankTransfer) | **Get** /BankTransfers/{BankTransferID} | Retrieves specific bank transfers by using a unique bank transfer Id
[**GetBankTransferAttachmentByFileName**](AccountingAPI.md#GetBankTransferAttachmentByFileName) | **Get** /BankTransfers/{BankTransferID}/Attachments/{FileName} | Retrieves a specific attachment on a specific bank transfer by file name
[**GetBankTransferAttachmentById**](AccountingAPI.md#GetBankTransferAttachmentById) | **Get** /BankTransfers/{BankTransferID}/Attachments/{AttachmentID} | Retrieves a specific attachment from a specific bank transfer using a unique attachment ID
[**GetBankTransferAttachments**](AccountingAPI.md#GetBankTransferAttachments) | **Get** /BankTransfers/{BankTransferID}/Attachments | Retrieves attachments from a specific bank transfer
[**GetBankTransferHistory**](AccountingAPI.md#GetBankTransferHistory) | **Get** /BankTransfers/{BankTransferID}/History | Retrieves history from a specific bank transfer using a unique bank transfer Id
[**GetBankTransfers**](AccountingAPI.md#GetBankTransfers) | **Get** /BankTransfers | Retrieves all bank transfers
[**GetBatchPayment**](AccountingAPI.md#GetBatchPayment) | **Get** /BatchPayments/{BatchPaymentID} | Retrieves a specific batch payment using a unique batch payment Id
[**GetBatchPaymentHistory**](AccountingAPI.md#GetBatchPaymentHistory) | **Get** /BatchPayments/{BatchPaymentID}/History | Retrieves history from a specific batch payment
[**GetBatchPayments**](AccountingAPI.md#GetBatchPayments) | **Get** /BatchPayments | Retrieves either one or many batch payments for invoices
[**GetBrandingTheme**](AccountingAPI.md#GetBrandingTheme) | **Get** /BrandingThemes/{BrandingThemeID} | Retrieves a specific branding theme using a unique branding theme Id
[**GetBrandingThemePaymentServices**](AccountingAPI.md#GetBrandingThemePaymentServices) | **Get** /BrandingThemes/{BrandingThemeID}/PaymentServices | Retrieves the payment services for a specific branding theme
[**GetBrandingThemes**](AccountingAPI.md#GetBrandingThemes) | **Get** /BrandingThemes | Retrieves all the branding themes
[**GetBudget**](AccountingAPI.md#GetBudget) | **Get** /Budgets/{BudgetID} | Retrieves a specific budget, which includes budget lines
[**GetBudgets**](AccountingAPI.md#GetBudgets) | **Get** /Budgets | Retrieve a list of budgets
[**GetContact**](AccountingAPI.md#GetContact) | **Get** /Contacts/{ContactID} | Retrieves a specific contacts in a Xero organisation using a unique contact Id
[**GetContactAttachmentByFileName**](AccountingAPI.md#GetContactAttachmentByFileName) | **Get** /Contacts/{ContactID}/Attachments/{FileName} | Retrieves a specific attachment from a specific contact by file name
[**GetContactAttachmentById**](AccountingAPI.md#GetContactAttachmentById) | **Get** /Contacts/{ContactID}/Attachments/{AttachmentID} | Retrieves a specific attachment from a specific contact using a unique attachment Id
[**GetContactAttachments**](AccountingAPI.md#GetContactAttachments) | **Get** /Contacts/{ContactID}/Attachments | Retrieves attachments for a specific contact in a Xero organisation
[**GetContactByContactNumber**](AccountingAPI.md#GetContactByContactNumber) | **Get** /Contacts/{ContactNumber} | Retrieves a specific contact by contact number in a Xero organisation
[**GetContactCISSettings**](AccountingAPI.md#GetContactCISSettings) | **Get** /Contacts/{ContactID}/CISSettings | Retrieves CIS settings for a specific contact in a Xero organisation
[**GetContactGroup**](AccountingAPI.md#GetContactGroup) | **Get** /ContactGroups/{ContactGroupID} | Retrieves a specific contact group by using a unique contact group Id
[**GetContactGroups**](AccountingAPI.md#GetContactGroups) | **Get** /ContactGroups | Retrieves the contact Id and name of each contact group
[**GetContactHistory**](AccountingAPI.md#GetContactHistory) | **Get** /Contacts/{ContactID}/History | Retrieves history records for a specific contact
[**GetContacts**](AccountingAPI.md#GetContacts) | **Get** /Contacts | Retrieves all contacts in a Xero organisation
[**GetCreditNote**](AccountingAPI.md#GetCreditNote) | **Get** /CreditNotes/{CreditNoteID} | Retrieves a specific credit note using a unique credit note Id
[**GetCreditNoteAsPdf**](AccountingAPI.md#GetCreditNoteAsPdf) | **Get** /CreditNotes/{CreditNoteID}/pdf | Retrieves credit notes as PDF files
[**GetCreditNoteAttachmentByFileName**](AccountingAPI.md#GetCreditNoteAttachmentByFileName) | **Get** /CreditNotes/{CreditNoteID}/Attachments/{FileName} | Retrieves a specific attachment on a specific credit note by file name
[**GetCreditNoteAttachmentById**](AccountingAPI.md#GetCreditNoteAttachmentById) | **Get** /CreditNotes/{CreditNoteID}/Attachments/{AttachmentID} | Retrieves a specific attachment from a specific credit note using a unique attachment Id
[**GetCreditNoteAttachments**](AccountingAPI.md#GetCreditNoteAttachments) | **Get** /CreditNotes/{CreditNoteID}/Attachments | Retrieves attachments for a specific credit notes
[**GetCreditNoteHistory**](AccountingAPI.md#GetCreditNoteHistory) | **Get** /CreditNotes/{CreditNoteID}/History | Retrieves history records of a specific credit note
[**GetCreditNotes**](AccountingAPI.md#GetCreditNotes) | **Get** /CreditNotes | Retrieves any credit notes
[**GetCurrencies**](AccountingAPI.md#GetCurrencies) | **Get** /Currencies | Retrieves currencies for your Xero organisation
[**GetEmployee**](AccountingAPI.md#GetEmployee) | **Get** /Employees/{EmployeeID} | Retrieves a specific employee used in Xero payrun using a unique employee Id
[**GetEmployees**](AccountingAPI.md#GetEmployees) | **Get** /Employees | Retrieves employees used in Xero payrun
[**GetExpenseClaim**](AccountingAPI.md#GetExpenseClaim) | **Get** /ExpenseClaims/{ExpenseClaimID} | Retrieves a specific expense claim using a unique expense claim Id
[**GetExpenseClaimHistory**](AccountingAPI.md#GetExpenseClaimHistory) | **Get** /ExpenseClaims/{ExpenseClaimID}/History | Retrieves history records of a specific expense claim
[**GetExpenseClaims**](AccountingAPI.md#GetExpenseClaims) | **Get** /ExpenseClaims | Retrieves expense claims
[**GetInvoice**](AccountingAPI.md#GetInvoice) | **Get** /Invoices/{InvoiceID} | Retrieves a specific sales invoice or purchase bill using a unique invoice Id
[**GetInvoiceAsPdf**](AccountingAPI.md#GetInvoiceAsPdf) | **Get** /Invoices/{InvoiceID}/pdf | Retrieves invoices or purchase bills as PDF files
[**GetInvoiceAttachmentByFileName**](AccountingAPI.md#GetInvoiceAttachmentByFileName) | **Get** /Invoices/{InvoiceID}/Attachments/{FileName} | Retrieves an attachment from a specific invoice or purchase bill by filename
[**GetInvoiceAttachmentById**](AccountingAPI.md#GetInvoiceAttachmentById) | **Get** /Invoices/{InvoiceID}/Attachments/{AttachmentID} | Retrieves a specific attachment from a specific invoices or purchase bills by using a unique attachment Id
[**GetInvoiceAttachments**](AccountingAPI.md#GetInvoiceAttachments) | **Get** /Invoices/{InvoiceID}/Attachments | Retrieves attachments for a specific invoice or purchase bill
[**GetInvoiceHistory**](AccountingAPI.md#GetInvoiceHistory) | **Get** /Invoices/{InvoiceID}/History | Retrieves history records for a specific invoice
[**GetInvoiceReminders**](AccountingAPI.md#GetInvoiceReminders) | **Get** /InvoiceReminders/Settings | Retrieves invoice reminder settings
[**GetInvoices**](AccountingAPI.md#GetInvoices) | **Get** /Invoices | Retrieves sales invoices or purchase bills
[**GetItem**](AccountingAPI.md#GetItem) | **Get** /Items/{ItemID} | Retrieves a specific item using a unique item Id
[**GetItemHistory**](AccountingAPI.md#GetItemHistory) | **Get** /Items/{ItemID}/History | Retrieves history for a specific item
[**GetItems**](AccountingAPI.md#GetItems) | **Get** /Items | Retrieves items
[**GetJournal**](AccountingAPI.md#GetJournal) | **Get** /Journals/{JournalID} | Retrieves a specific journal using a unique journal Id.
[**GetJournalByNumber**](AccountingAPI.md#GetJournalByNumber) | **Get** /Journals/{JournalNumber} | Retrieves a specific journal using a unique journal number.
[**GetJournals**](AccountingAPI.md#GetJournals) | **Get** /Journals | Retrieves journals
[**GetLinkedTransaction**](AccountingAPI.md#GetLinkedTransaction) | **Get** /LinkedTransactions/{LinkedTransactionID} | Retrieves a specific linked transaction (billable expenses) using a unique linked transaction Id
[**GetLinkedTransactions**](AccountingAPI.md#GetLinkedTransactions) | **Get** /LinkedTransactions | Retrieves linked transactions (billable expenses)
[**GetManualJournal**](AccountingAPI.md#GetManualJournal) | **Get** /ManualJournals/{ManualJournalID} | Retrieves a specific manual journal
[**GetManualJournalAttachmentByFileName**](AccountingAPI.md#GetManualJournalAttachmentByFileName) | **Get** /ManualJournals/{ManualJournalID}/Attachments/{FileName} | Retrieves a specific attachment from a specific manual journal by file name
[**GetManualJournalAttachmentById**](AccountingAPI.md#GetManualJournalAttachmentById) | **Get** /ManualJournals/{ManualJournalID}/Attachments/{AttachmentID} | Allows you to retrieve a specific attachment from a specific manual journal using a unique attachment Id
[**GetManualJournalAttachments**](AccountingAPI.md#GetManualJournalAttachments) | **Get** /ManualJournals/{ManualJournalID}/Attachments | Retrieves attachment for a specific manual journal
[**GetManualJournals**](AccountingAPI.md#GetManualJournals) | **Get** /ManualJournals | Retrieves manual journals
[**GetManualJournalsHistory**](AccountingAPI.md#GetManualJournalsHistory) | **Get** /ManualJournals/{ManualJournalID}/History | Retrieves history for a specific manual journal
[**GetOnlineInvoice**](AccountingAPI.md#GetOnlineInvoice) | **Get** /Invoices/{InvoiceID}/OnlineInvoice | Retrieves a URL to an online invoice
[**GetOrganisationActions**](AccountingAPI.md#GetOrganisationActions) | **Get** /Organisation/Actions | Retrieves a list of the key actions your app has permission to perform in the connected Xero organisation.
[**GetOrganisationCISSettings**](AccountingAPI.md#GetOrganisationCISSettings) | **Get** /Organisation/{OrganisationID}/CISSettings | Retrieves the CIS settings for the Xero organistaion.
[**GetOrganisations**](AccountingAPI.md#GetOrganisations) | **Get** /Organisation | Retrieves Xero organisation details
[**GetOverpayment**](AccountingAPI.md#GetOverpayment) | **Get** /Overpayments/{OverpaymentID} | Retrieves a specific overpayment using a unique overpayment Id
[**GetOverpaymentHistory**](AccountingAPI.md#GetOverpaymentHistory) | **Get** /Overpayments/{OverpaymentID}/History | Retrieves history records of a specific overpayment
[**GetOverpayments**](AccountingAPI.md#GetOverpayments) | **Get** /Overpayments | Retrieves overpayments
[**GetPayment**](AccountingAPI.md#GetPayment) | **Get** /Payments/{PaymentID} | Retrieves a specific payment for invoices and credit notes using a unique payment Id
[**GetPaymentHistory**](AccountingAPI.md#GetPaymentHistory) | **Get** /Payments/{PaymentID}/History | Retrieves history records of a specific payment
[**GetPaymentServices**](AccountingAPI.md#GetPaymentServices) | **Get** /PaymentServices | Retrieves payment services
[**GetPayments**](AccountingAPI.md#GetPayments) | **Get** /Payments | Retrieves payments for invoices and credit notes
[**GetPrepayment**](AccountingAPI.md#GetPrepayment) | **Get** /Prepayments/{PrepaymentID} | Allows you to retrieve a specified prepayments
[**GetPrepaymentHistory**](AccountingAPI.md#GetPrepaymentHistory) | **Get** /Prepayments/{PrepaymentID}/History | Retrieves history record for a specific prepayment
[**GetPrepayments**](AccountingAPI.md#GetPrepayments) | **Get** /Prepayments | Retrieves prepayments
[**GetPurchaseOrder**](AccountingAPI.md#GetPurchaseOrder) | **Get** /PurchaseOrders/{PurchaseOrderID} | Retrieves a specific purchase order using a unique purchase order Id
[**GetPurchaseOrderAsPdf**](AccountingAPI.md#GetPurchaseOrderAsPdf) | **Get** /PurchaseOrders/{PurchaseOrderID}/pdf | Retrieves specific purchase order as PDF files using a unique purchase order Id
[**GetPurchaseOrderAttachmentByFileName**](AccountingAPI.md#GetPurchaseOrderAttachmentByFileName) | **Get** /PurchaseOrders/{PurchaseOrderID}/Attachments/{FileName} | Retrieves a specific attachment for a specific purchase order by filename
[**GetPurchaseOrderAttachmentById**](AccountingAPI.md#GetPurchaseOrderAttachmentById) | **Get** /PurchaseOrders/{PurchaseOrderID}/Attachments/{AttachmentID} | Retrieves specific attachment for a specific purchase order using a unique attachment Id
[**GetPurchaseOrderAttachments**](AccountingAPI.md#GetPurchaseOrderAttachments) | **Get** /PurchaseOrders/{PurchaseOrderID}/Attachments | Retrieves attachments for a specific purchase order
[**GetPurchaseOrderByNumber**](AccountingAPI.md#GetPurchaseOrderByNumber) | **Get** /PurchaseOrders/{PurchaseOrderNumber} | Retrieves a specific purchase order using purchase order number
[**GetPurchaseOrderHistory**](AccountingAPI.md#GetPurchaseOrderHistory) | **Get** /PurchaseOrders/{PurchaseOrderID}/History | Retrieves history for a specific purchase order
[**GetPurchaseOrders**](AccountingAPI.md#GetPurchaseOrders) | **Get** /PurchaseOrders | Retrieves purchase orders
[**GetQuote**](AccountingAPI.md#GetQuote) | **Get** /Quotes/{QuoteID} | Retrieves a specific quote using a unique quote Id
[**GetQuoteAsPdf**](AccountingAPI.md#GetQuoteAsPdf) | **Get** /Quotes/{QuoteID}/pdf | Retrieves a specific quote as a PDF file using a unique quote Id
[**GetQuoteAttachmentByFileName**](AccountingAPI.md#GetQuoteAttachmentByFileName) | **Get** /Quotes/{QuoteID}/Attachments/{FileName} | Retrieves a specific attachment from a specific quote by filename
[**GetQuoteAttachmentById**](AccountingAPI.md#GetQuoteAttachmentById) | **Get** /Quotes/{QuoteID}/Attachments/{AttachmentID} | Retrieves a specific attachment from a specific quote using a unique attachment Id
[**GetQuoteAttachments**](AccountingAPI.md#GetQuoteAttachments) | **Get** /Quotes/{QuoteID}/Attachments | Retrieves attachments for a specific quote
[**GetQuoteHistory**](AccountingAPI.md#GetQuoteHistory) | **Get** /Quotes/{QuoteID}/History | Retrieves history records of a specific quote
[**GetQuotes**](AccountingAPI.md#GetQuotes) | **Get** /Quotes | Retrieves sales quotes
[**GetReceipt**](AccountingAPI.md#GetReceipt) | **Get** /Receipts/{ReceiptID} | Retrieves a specific draft expense claim receipt by using a unique receipt Id
[**GetReceiptAttachmentByFileName**](AccountingAPI.md#GetReceiptAttachmentByFileName) | **Get** /Receipts/{ReceiptID}/Attachments/{FileName} | Retrieves a specific attachment from a specific expense claim receipts by file name
[**GetReceiptAttachmentById**](AccountingAPI.md#GetReceiptAttachmentById) | **Get** /Receipts/{ReceiptID}/Attachments/{AttachmentID} | Retrieves a specific attachments from a specific expense claim receipts by using a unique attachment Id
[**GetReceiptAttachments**](AccountingAPI.md#GetReceiptAttachments) | **Get** /Receipts/{ReceiptID}/Attachments | Retrieves attachments for a specific expense claim receipt
[**GetReceiptHistory**](AccountingAPI.md#GetReceiptHistory) | **Get** /Receipts/{ReceiptID}/History | Retrieves a history record for a specific receipt
[**GetReceipts**](AccountingAPI.md#GetReceipts) | **Get** /Receipts | Retrieves draft expense claim receipts for any user
[**GetRepeatingInvoice**](AccountingAPI.md#GetRepeatingInvoice) | **Get** /RepeatingInvoices/{RepeatingInvoiceID} | Retrieves a specific repeating invoice by using a unique repeating invoice Id
[**GetRepeatingInvoiceAttachmentByFileName**](AccountingAPI.md#GetRepeatingInvoiceAttachmentByFileName) | **Get** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments/{FileName} | Retrieves a specific attachment from a specific repeating invoices by file name
[**GetRepeatingInvoiceAttachmentById**](AccountingAPI.md#GetRepeatingInvoiceAttachmentById) | **Get** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments/{AttachmentID} | Retrieves a specific attachment from a specific repeating invoice
[**GetRepeatingInvoiceAttachments**](AccountingAPI.md#GetRepeatingInvoiceAttachments) | **Get** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments | Retrieves attachments from a specific repeating invoice
[**GetRepeatingInvoiceHistory**](AccountingAPI.md#GetRepeatingInvoiceHistory) | **Get** /RepeatingInvoices/{RepeatingInvoiceID}/History | Retrieves history record for a specific repeating invoice
[**GetRepeatingInvoices**](AccountingAPI.md#GetRepeatingInvoices) | **Get** /RepeatingInvoices | Retrieves repeating invoices
[**GetReportAgedPayablesByContact**](AccountingAPI.md#GetReportAgedPayablesByContact) | **Get** /Reports/AgedPayablesByContact | Retrieves report for aged payables by contact
[**GetReportAgedReceivablesByContact**](AccountingAPI.md#GetReportAgedReceivablesByContact) | **Get** /Reports/AgedReceivablesByContact | Retrieves report for aged receivables by contact
[**GetReportBalanceSheet**](AccountingAPI.md#GetReportBalanceSheet) | **Get** /Reports/BalanceSheet | Retrieves report for balancesheet
[**GetReportBankSummary**](AccountingAPI.md#GetReportBankSummary) | **Get** /Reports/BankSummary | Retrieves report for bank summary
[**GetReportBudgetSummary**](AccountingAPI.md#GetReportBudgetSummary) | **Get** /Reports/BudgetSummary | Retrieves report for budget summary
[**GetReportExecutiveSummary**](AccountingAPI.md#GetReportExecutiveSummary) | **Get** /Reports/ExecutiveSummary | Retrieves report for executive summary
[**GetReportFromId**](AccountingAPI.md#GetReportFromId) | **Get** /Reports/{ReportID} | Retrieves a specific report using a unique ReportID
[**GetReportProfitAndLoss**](AccountingAPI.md#GetReportProfitAndLoss) | **Get** /Reports/ProfitAndLoss | Retrieves report for profit and loss
[**GetReportTenNinetyNine**](AccountingAPI.md#GetReportTenNinetyNine) | **Get** /Reports/TenNinetyNine | Retrieve reports for 1099
[**GetReportTrialBalance**](AccountingAPI.md#GetReportTrialBalance) | **Get** /Reports/TrialBalance | Retrieves report for trial balance
[**GetReportsList**](AccountingAPI.md#GetReportsList) | **Get** /Reports | Retrieves a list of the organistaions unique reports that require a uuid to fetch
[**GetTaxRateByTaxType**](AccountingAPI.md#GetTaxRateByTaxType) | **Get** /TaxRates/{TaxType} | Retrieves a specific tax rate according to given TaxType code
[**GetTaxRates**](AccountingAPI.md#GetTaxRates) | **Get** /TaxRates | Retrieves tax rates
[**GetTrackingCategories**](AccountingAPI.md#GetTrackingCategories) | **Get** /TrackingCategories | Retrieves tracking categories and options
[**GetTrackingCategory**](AccountingAPI.md#GetTrackingCategory) | **Get** /TrackingCategories/{TrackingCategoryID} | Retrieves specific tracking categories and options using a unique tracking category Id
[**GetUser**](AccountingAPI.md#GetUser) | **Get** /Users/{UserID} | Retrieves a specific user
[**GetUsers**](AccountingAPI.md#GetUsers) | **Get** /Users | Retrieves users
[**PostSetup**](AccountingAPI.md#PostSetup) | **Post** /Setup | Sets the chart of accounts, the conversion date and conversion balances
[**UpdateAccount**](AccountingAPI.md#UpdateAccount) | **Post** /Accounts/{AccountID} | Updates a chart of accounts
[**UpdateAccountAttachmentByFileName**](AccountingAPI.md#UpdateAccountAttachmentByFileName) | **Post** /Accounts/{AccountID}/Attachments/{FileName} | Updates attachment on a specific account by filename
[**UpdateBankTransaction**](AccountingAPI.md#UpdateBankTransaction) | **Post** /BankTransactions/{BankTransactionID} | Updates a single spent or received money transaction
[**UpdateBankTransactionAttachmentByFileName**](AccountingAPI.md#UpdateBankTransactionAttachmentByFileName) | **Post** /BankTransactions/{BankTransactionID}/Attachments/{FileName} | Updates a specific attachment from a specific bank transaction by filename
[**UpdateBankTransferAttachmentByFileName**](AccountingAPI.md#UpdateBankTransferAttachmentByFileName) | **Post** /BankTransfers/{BankTransferID}/Attachments/{FileName} | 
[**UpdateContact**](AccountingAPI.md#UpdateContact) | **Post** /Contacts/{ContactID} | Updates a specific contact in a Xero organisation
[**UpdateContactAttachmentByFileName**](AccountingAPI.md#UpdateContactAttachmentByFileName) | **Post** /Contacts/{ContactID}/Attachments/{FileName} | 
[**UpdateContactGroup**](AccountingAPI.md#UpdateContactGroup) | **Post** /ContactGroups/{ContactGroupID} | Updates a specific contact group
[**UpdateCreditNote**](AccountingAPI.md#UpdateCreditNote) | **Post** /CreditNotes/{CreditNoteID} | Updates a specific credit note
[**UpdateCreditNoteAttachmentByFileName**](AccountingAPI.md#UpdateCreditNoteAttachmentByFileName) | **Post** /CreditNotes/{CreditNoteID}/Attachments/{FileName} | Updates attachments on a specific credit note by file name
[**UpdateExpenseClaim**](AccountingAPI.md#UpdateExpenseClaim) | **Post** /ExpenseClaims/{ExpenseClaimID} | Updates a specific expense claims
[**UpdateInvoice**](AccountingAPI.md#UpdateInvoice) | **Post** /Invoices/{InvoiceID} | Updates a specific sales invoices or purchase bills
[**UpdateInvoiceAttachmentByFileName**](AccountingAPI.md#UpdateInvoiceAttachmentByFileName) | **Post** /Invoices/{InvoiceID}/Attachments/{FileName} | Updates an attachment from a specific invoices or purchase bill by filename
[**UpdateItem**](AccountingAPI.md#UpdateItem) | **Post** /Items/{ItemID} | Updates a specific item
[**UpdateLinkedTransaction**](AccountingAPI.md#UpdateLinkedTransaction) | **Post** /LinkedTransactions/{LinkedTransactionID} | Updates a specific linked transactions (billable expenses)
[**UpdateManualJournal**](AccountingAPI.md#UpdateManualJournal) | **Post** /ManualJournals/{ManualJournalID} | Updates a specific manual journal
[**UpdateManualJournalAttachmentByFileName**](AccountingAPI.md#UpdateManualJournalAttachmentByFileName) | **Post** /ManualJournals/{ManualJournalID}/Attachments/{FileName} | Updates a specific attachment from a specific manual journal by file name
[**UpdateOrCreateBankTransactions**](AccountingAPI.md#UpdateOrCreateBankTransactions) | **Post** /BankTransactions | Updates or creates one or more spent or received money transaction
[**UpdateOrCreateContacts**](AccountingAPI.md#UpdateOrCreateContacts) | **Post** /Contacts | Updates or creates one or more contacts in a Xero organisation
[**UpdateOrCreateCreditNotes**](AccountingAPI.md#UpdateOrCreateCreditNotes) | **Post** /CreditNotes | Updates or creates one or more credit notes
[**UpdateOrCreateEmployees**](AccountingAPI.md#UpdateOrCreateEmployees) | **Post** /Employees | Creates a single new employees used in Xero payrun
[**UpdateOrCreateInvoices**](AccountingAPI.md#UpdateOrCreateInvoices) | **Post** /Invoices | Updates or creates one or more sales invoices or purchase bills
[**UpdateOrCreateItems**](AccountingAPI.md#UpdateOrCreateItems) | **Post** /Items | Updates or creates one or more items
[**UpdateOrCreateManualJournals**](AccountingAPI.md#UpdateOrCreateManualJournals) | **Post** /ManualJournals | Updates or creates a single manual journal
[**UpdateOrCreatePurchaseOrders**](AccountingAPI.md#UpdateOrCreatePurchaseOrders) | **Post** /PurchaseOrders | Updates or creates one or more purchase orders
[**UpdateOrCreateQuotes**](AccountingAPI.md#UpdateOrCreateQuotes) | **Post** /Quotes | Updates or creates one or more quotes
[**UpdateOrCreateRepeatingInvoices**](AccountingAPI.md#UpdateOrCreateRepeatingInvoices) | **Post** /RepeatingInvoices | Creates or deletes one or more repeating invoice templates
[**UpdatePurchaseOrder**](AccountingAPI.md#UpdatePurchaseOrder) | **Post** /PurchaseOrders/{PurchaseOrderID} | Updates a specific purchase order
[**UpdatePurchaseOrderAttachmentByFileName**](AccountingAPI.md#UpdatePurchaseOrderAttachmentByFileName) | **Post** /PurchaseOrders/{PurchaseOrderID}/Attachments/{FileName} | Updates a specific attachment for a specific purchase order by filename
[**UpdateQuote**](AccountingAPI.md#UpdateQuote) | **Post** /Quotes/{QuoteID} | Updates a specific quote
[**UpdateQuoteAttachmentByFileName**](AccountingAPI.md#UpdateQuoteAttachmentByFileName) | **Post** /Quotes/{QuoteID}/Attachments/{FileName} | Updates a specific attachment from a specific quote by filename
[**UpdateReceipt**](AccountingAPI.md#UpdateReceipt) | **Post** /Receipts/{ReceiptID} | Updates a specific draft expense claim receipts
[**UpdateReceiptAttachmentByFileName**](AccountingAPI.md#UpdateReceiptAttachmentByFileName) | **Post** /Receipts/{ReceiptID}/Attachments/{FileName} | Updates a specific attachment on a specific expense claim receipts by file name
[**UpdateRepeatingInvoice**](AccountingAPI.md#UpdateRepeatingInvoice) | **Post** /RepeatingInvoices/{RepeatingInvoiceID} | Deletes a specific repeating invoice template
[**UpdateRepeatingInvoiceAttachmentByFileName**](AccountingAPI.md#UpdateRepeatingInvoiceAttachmentByFileName) | **Post** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments/{FileName} | Updates a specific attachment from a specific repeating invoices by file name
[**UpdateTaxRate**](AccountingAPI.md#UpdateTaxRate) | **Post** /TaxRates | Updates tax rates
[**UpdateTrackingCategory**](AccountingAPI.md#UpdateTrackingCategory) | **Post** /TrackingCategories/{TrackingCategoryID} | Updates a specific tracking category
[**UpdateTrackingOptions**](AccountingAPI.md#UpdateTrackingOptions) | **Post** /TrackingCategories/{TrackingCategoryID}/Options/{TrackingOptionID} | Updates a specific option for a specific tracking category



## CreateAccount

> Accounts CreateAccount(ctx).XeroTenantId(xeroTenantId).Account(account).IdempotencyKey(idempotencyKey).Execute()

Creates a new chart of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	account := *openapiclient.NewAccount() // Account | Account object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateAccount(context.Background()).XeroTenantId(xeroTenantId).Account(account).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: Accounts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **account** | [**Account**](Account.md) | Account object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountAttachmentByFileName

> Attachments CreateAccountAttachmentByFileName(ctx, accountID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates an attachment on a specific account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	accountID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Account object
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateAccountAttachmentByFileName(context.Background(), accountID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateAccountAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateAccountAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Unique identifier for Account object | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransactionAttachmentByFileName

> Attachments CreateBankTransactionAttachmentByFileName(ctx, bankTransactionID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates an attachment for a specific bank transaction by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBankTransactionAttachmentByFileName(context.Background(), bankTransactionID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBankTransactionAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankTransactionAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBankTransactionAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankTransactionAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransactionHistoryRecord

> HistoryRecords CreateBankTransactionHistoryRecord(ctx, bankTransactionID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific bank transactions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBankTransactionHistoryRecord(context.Background(), bankTransactionID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBankTransactionHistoryRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankTransactionHistoryRecord`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBankTransactionHistoryRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankTransactionHistoryRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransactions

> BankTransactions CreateBankTransactions(ctx).XeroTenantId(xeroTenantId).BankTransactions(bankTransactions).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Creates one or more spent or received money transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactions := *openapiclient.NewBankTransactions() // BankTransactions | BankTransactions with an array of BankTransaction objects in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBankTransactions(context.Background()).XeroTenantId(xeroTenantId).BankTransactions(bankTransactions).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBankTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankTransactions`: BankTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBankTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **bankTransactions** | [**BankTransactions**](BankTransactions.md) | BankTransactions with an array of BankTransaction objects in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransfer

> BankTransfers CreateBankTransfer(ctx).XeroTenantId(xeroTenantId).BankTransfers(bankTransfers).IdempotencyKey(idempotencyKey).Execute()

Creates a bank transfer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransfers := *openapiclient.NewBankTransfers() // BankTransfers | BankTransfers with array of BankTransfer objects in request body
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBankTransfer(context.Background()).XeroTenantId(xeroTenantId).BankTransfers(bankTransfers).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBankTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankTransfer`: BankTransfers
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBankTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **bankTransfers** | [**BankTransfers**](BankTransfers.md) | BankTransfers with array of BankTransfer objects in request body | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**BankTransfers**](BankTransfers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransferAttachmentByFileName

> Attachments CreateBankTransferAttachmentByFileName(ctx, bankTransferID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransferID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transfer
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBankTransferAttachmentByFileName(context.Background(), bankTransferID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBankTransferAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankTransferAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBankTransferAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransferID** | **string** | Xero generated unique identifier for a bank transfer | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankTransferAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransferHistoryRecord

> HistoryRecords CreateBankTransferHistoryRecord(ctx, bankTransferID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific bank transfer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransferID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transfer
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBankTransferHistoryRecord(context.Background(), bankTransferID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBankTransferHistoryRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankTransferHistoryRecord`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBankTransferHistoryRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransferID** | **string** | Xero generated unique identifier for a bank transfer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankTransferHistoryRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchPayment

> BatchPayments CreateBatchPayment(ctx).XeroTenantId(xeroTenantId).BatchPayments(batchPayments).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates one or many batch payments for invoices

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	batchPayments := *openapiclient.NewBatchPayments() // BatchPayments | BatchPayments with an array of Payments in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBatchPayment(context.Background()).XeroTenantId(xeroTenantId).BatchPayments(batchPayments).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBatchPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchPayment`: BatchPayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBatchPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **batchPayments** | [**BatchPayments**](BatchPayments.md) | BatchPayments with an array of Payments in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**BatchPayments**](BatchPayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchPaymentHistoryRecord

> HistoryRecords CreateBatchPaymentHistoryRecord(ctx, batchPaymentID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific batch payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	batchPaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for BatchPayment
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBatchPaymentHistoryRecord(context.Background(), batchPaymentID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBatchPaymentHistoryRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchPaymentHistoryRecord`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBatchPaymentHistoryRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchPaymentID** | **string** | Unique identifier for BatchPayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchPaymentHistoryRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBrandingThemePaymentServices

> PaymentServices CreateBrandingThemePaymentServices(ctx, brandingThemeID).XeroTenantId(xeroTenantId).PaymentServices(paymentServices).IdempotencyKey(idempotencyKey).Execute()

Creates a new custom payment service for a specific branding theme

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	brandingThemeID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Branding Theme
	paymentServices := *openapiclient.NewPaymentServices() // PaymentServices | PaymentServices array with PaymentService object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateBrandingThemePaymentServices(context.Background(), brandingThemeID).XeroTenantId(xeroTenantId).PaymentServices(paymentServices).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateBrandingThemePaymentServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBrandingThemePaymentServices`: PaymentServices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateBrandingThemePaymentServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandingThemeID** | **string** | Unique identifier for a Branding Theme | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingThemePaymentServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **paymentServices** | [**PaymentServices**](PaymentServices.md) | PaymentServices array with PaymentService object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactAttachmentByFileName

> Attachments CreateContactAttachmentByFileName(ctx, contactID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateContactAttachmentByFileName(context.Background(), contactID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateContactAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateContactAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactGroup

> ContactGroups CreateContactGroup(ctx).XeroTenantId(xeroTenantId).ContactGroups(contactGroups).IdempotencyKey(idempotencyKey).Execute()

Creates a contact group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactGroups := *openapiclient.NewContactGroups() // ContactGroups | ContactGroups with an array of names in request body
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateContactGroup(context.Background()).XeroTenantId(xeroTenantId).ContactGroups(contactGroups).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateContactGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactGroup`: ContactGroups
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateContactGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **contactGroups** | [**ContactGroups**](ContactGroups.md) | ContactGroups with an array of names in request body | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**ContactGroups**](ContactGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactGroupContacts

> Contacts CreateContactGroupContacts(ctx, contactGroupID).XeroTenantId(xeroTenantId).Contacts(contacts).IdempotencyKey(idempotencyKey).Execute()

Creates contacts to a specific contact group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactGroupID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact Group
	contacts := *openapiclient.NewContacts() // Contacts | Contacts with array of contacts specifying the ContactID to be added to ContactGroup in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateContactGroupContacts(context.Background(), contactGroupID).XeroTenantId(xeroTenantId).Contacts(contacts).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateContactGroupContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactGroupContacts`: Contacts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateContactGroupContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactGroupID** | **string** | Unique identifier for a Contact Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactGroupContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **contacts** | [**Contacts**](Contacts.md) | Contacts with array of contacts specifying the ContactID to be added to ContactGroup in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactHistory

> HistoryRecords CreateContactHistory(ctx, contactID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a new history record for a specific contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateContactHistory(context.Background(), contactID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateContactHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateContactHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContacts

> Contacts CreateContacts(ctx).XeroTenantId(xeroTenantId).Contacts(contacts).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates multiple contacts (bulk) in a Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contacts := *openapiclient.NewContacts() // Contacts | Contacts with an array of Contact objects to create in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateContacts(context.Background()).XeroTenantId(xeroTenantId).Contacts(contacts).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContacts`: Contacts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **contacts** | [**Contacts**](Contacts.md) | Contacts with an array of Contact objects to create in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditNoteAllocation

> Allocations CreateCreditNoteAllocation(ctx, creditNoteID).XeroTenantId(xeroTenantId).Allocations(allocations).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates allocation for a specific credit note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	allocations := *openapiclient.NewAllocations() // Allocations | Allocations with array of Allocation object in body of request.
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateCreditNoteAllocation(context.Background(), creditNoteID).XeroTenantId(xeroTenantId).Allocations(allocations).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateCreditNoteAllocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCreditNoteAllocation`: Allocations
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateCreditNoteAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditNoteAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **allocations** | [**Allocations**](Allocations.md) | Allocations with array of Allocation object in body of request. | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Allocations**](Allocations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditNoteAttachmentByFileName

> Attachments CreateCreditNoteAttachmentByFileName(ctx, creditNoteID, fileName).XeroTenantId(xeroTenantId).Body(body).IncludeOnline(includeOnline).IdempotencyKey(idempotencyKey).Execute()

Creates an attachment for a specific credit note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	includeOnline := true // bool | Allows an attachment to be seen by the end customer within their online invoice (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateCreditNoteAttachmentByFileName(context.Background(), creditNoteID, fileName).XeroTenantId(xeroTenantId).Body(body).IncludeOnline(includeOnline).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateCreditNoteAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCreditNoteAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateCreditNoteAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditNoteAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **includeOnline** | **bool** | Allows an attachment to be seen by the end customer within their online invoice | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditNoteHistory

> HistoryRecords CreateCreditNoteHistory(ctx, creditNoteID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Retrieves history records of a specific credit note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateCreditNoteHistory(context.Background(), creditNoteID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateCreditNoteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCreditNoteHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateCreditNoteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditNoteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditNotes

> CreditNotes CreateCreditNotes(ctx).XeroTenantId(xeroTenantId).CreditNotes(creditNotes).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Creates a new credit note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNotes := *openapiclient.NewCreditNotes() // CreditNotes | Credit Notes with array of CreditNote object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateCreditNotes(context.Background()).XeroTenantId(xeroTenantId).CreditNotes(creditNotes).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateCreditNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCreditNotes`: CreditNotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateCreditNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **creditNotes** | [**CreditNotes**](CreditNotes.md) | Credit Notes with array of CreditNote object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCurrency

> Currencies CreateCurrency(ctx).XeroTenantId(xeroTenantId).Currency(currency).IdempotencyKey(idempotencyKey).Execute()

Create a new currency for a Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	currency := *openapiclient.NewCurrency() // Currency | Currency object in the body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateCurrency(context.Background()).XeroTenantId(xeroTenantId).Currency(currency).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateCurrency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCurrency`: Currencies
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateCurrency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **currency** | [**Currency**](Currency.md) | Currency object in the body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Currencies**](Currencies.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmployees

> Employees CreateEmployees(ctx).XeroTenantId(xeroTenantId).Employees(employees).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates new employees used in Xero payrun

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	employees := *openapiclient.NewEmployees() // Employees | Employees with array of Employee object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateEmployees(context.Background()).XeroTenantId(xeroTenantId).Employees(employees).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateEmployees``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmployees`: Employees
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateEmployees`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmployeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **employees** | [**Employees**](Employees.md) | Employees with array of Employee object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Employees**](Employees.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExpenseClaimHistory

> HistoryRecords CreateExpenseClaimHistory(ctx, expenseClaimID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific expense claim

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	expenseClaimID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ExpenseClaim
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateExpenseClaimHistory(context.Background(), expenseClaimID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateExpenseClaimHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExpenseClaimHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateExpenseClaimHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseClaimID** | **string** | Unique identifier for a ExpenseClaim | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExpenseClaimHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExpenseClaims

> ExpenseClaims CreateExpenseClaims(ctx).XeroTenantId(xeroTenantId).ExpenseClaims(expenseClaims).IdempotencyKey(idempotencyKey).Execute()

Creates expense claims

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	expenseClaims := *openapiclient.NewExpenseClaims() // ExpenseClaims | ExpenseClaims with array of ExpenseClaim object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateExpenseClaims(context.Background()).XeroTenantId(xeroTenantId).ExpenseClaims(expenseClaims).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateExpenseClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExpenseClaims`: ExpenseClaims
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateExpenseClaims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExpenseClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **expenseClaims** | [**ExpenseClaims**](ExpenseClaims.md) | ExpenseClaims with array of ExpenseClaim object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**ExpenseClaims**](ExpenseClaims.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceAttachmentByFileName

> Attachments CreateInvoiceAttachmentByFileName(ctx, invoiceID, fileName).XeroTenantId(xeroTenantId).Body(body).IncludeOnline(includeOnline).IdempotencyKey(idempotencyKey).Execute()

Creates an attachment for a specific invoice or purchase bill by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	includeOnline := true // bool | Allows an attachment to be seen by the end customer within their online invoice (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateInvoiceAttachmentByFileName(context.Background(), invoiceID, fileName).XeroTenantId(xeroTenantId).Body(body).IncludeOnline(includeOnline).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateInvoiceAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateInvoiceAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **includeOnline** | **bool** | Allows an attachment to be seen by the end customer within their online invoice | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceHistory

> HistoryRecords CreateInvoiceHistory(ctx, invoiceID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific invoice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateInvoiceHistory(context.Background(), invoiceID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateInvoiceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateInvoiceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoices

> Invoices CreateInvoices(ctx).XeroTenantId(xeroTenantId).Invoices(invoices).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Creates one or more sales invoices or purchase bills

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoices := *openapiclient.NewInvoices() // Invoices | Invoices with an array of invoice objects in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateInvoices(context.Background()).XeroTenantId(xeroTenantId).Invoices(invoices).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoices`: Invoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **invoices** | [**Invoices**](Invoices.md) | Invoices with an array of invoice objects in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateItemHistory

> HistoryRecords CreateItemHistory(ctx, itemID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific item

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	itemID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Item
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateItemHistory(context.Background(), itemID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateItemHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateItemHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemID** | **string** | Unique identifier for an Item | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateItems

> Items CreateItems(ctx).XeroTenantId(xeroTenantId).Items(items).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Creates one or more items

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	items := *openapiclient.NewItems() // Items | Items with an array of Item objects in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateItems(context.Background()).XeroTenantId(xeroTenantId).Items(items).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItems`: Items
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **items** | [**Items**](Items.md) | Items with an array of Item objects in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLinkedTransaction

> LinkedTransactions CreateLinkedTransaction(ctx).XeroTenantId(xeroTenantId).LinkedTransaction(linkedTransaction).IdempotencyKey(idempotencyKey).Execute()

Creates linked transactions (billable expenses)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	linkedTransaction := *openapiclient.NewLinkedTransaction() // LinkedTransaction | LinkedTransaction object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateLinkedTransaction(context.Background()).XeroTenantId(xeroTenantId).LinkedTransaction(linkedTransaction).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateLinkedTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLinkedTransaction`: LinkedTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateLinkedTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **linkedTransaction** | [**LinkedTransaction**](LinkedTransaction.md) | LinkedTransaction object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**LinkedTransactions**](LinkedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManualJournalAttachmentByFileName

> Attachments CreateManualJournalAttachmentByFileName(ctx, manualJournalID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates a specific attachment for a specific manual journal by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateManualJournalAttachmentByFileName(context.Background(), manualJournalID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateManualJournalAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManualJournalAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateManualJournalAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualJournalAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManualJournalHistoryRecord

> HistoryRecords CreateManualJournalHistoryRecord(ctx, manualJournalID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific manual journal

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateManualJournalHistoryRecord(context.Background(), manualJournalID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateManualJournalHistoryRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManualJournalHistoryRecord`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateManualJournalHistoryRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualJournalHistoryRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManualJournals

> ManualJournals CreateManualJournals(ctx).XeroTenantId(xeroTenantId).ManualJournals(manualJournals).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates one or more manual journals

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournals := *openapiclient.NewManualJournals() // ManualJournals | ManualJournals array with ManualJournal object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateManualJournals(context.Background()).XeroTenantId(xeroTenantId).ManualJournals(manualJournals).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateManualJournals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManualJournals`: ManualJournals
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateManualJournals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualJournalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **manualJournals** | [**ManualJournals**](ManualJournals.md) | ManualJournals array with ManualJournal object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOverpaymentAllocations

> Allocations CreateOverpaymentAllocations(ctx, overpaymentID).XeroTenantId(xeroTenantId).Allocations(allocations).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates a single allocation for a specific overpayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	overpaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Overpayment
	allocations := *openapiclient.NewAllocations() // Allocations | Allocations array with Allocation object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateOverpaymentAllocations(context.Background(), overpaymentID).XeroTenantId(xeroTenantId).Allocations(allocations).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateOverpaymentAllocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOverpaymentAllocations`: Allocations
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateOverpaymentAllocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**overpaymentID** | **string** | Unique identifier for a Overpayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOverpaymentAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **allocations** | [**Allocations**](Allocations.md) | Allocations array with Allocation object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Allocations**](Allocations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOverpaymentHistory

> HistoryRecords CreateOverpaymentHistory(ctx, overpaymentID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific overpayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	overpaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Overpayment
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateOverpaymentHistory(context.Background(), overpaymentID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateOverpaymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOverpaymentHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateOverpaymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**overpaymentID** | **string** | Unique identifier for a Overpayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOverpaymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayment

> Payments CreatePayment(ctx).XeroTenantId(xeroTenantId).Payment(payment).IdempotencyKey(idempotencyKey).Execute()

Creates a single payment for invoice or credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	payment := *openapiclient.NewPayment() // Payment | Request body with a single Payment object
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePayment(context.Background()).XeroTenantId(xeroTenantId).Payment(payment).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePayment`: Payments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **payment** | [**Payment**](Payment.md) | Request body with a single Payment object | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentHistory

> HistoryRecords CreatePaymentHistory(ctx, paymentID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	paymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Payment
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePaymentHistory(context.Background(), paymentID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePaymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePaymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentID** | **string** | Unique identifier for a Payment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentService

> PaymentServices CreatePaymentService(ctx).XeroTenantId(xeroTenantId).PaymentServices(paymentServices).IdempotencyKey(idempotencyKey).Execute()

Creates a payment service

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	paymentServices := *openapiclient.NewPaymentServices() // PaymentServices | PaymentServices array with PaymentService object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePaymentService(context.Background()).XeroTenantId(xeroTenantId).PaymentServices(paymentServices).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePaymentService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentService`: PaymentServices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePaymentService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **paymentServices** | [**PaymentServices**](PaymentServices.md) | PaymentServices array with PaymentService object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayments

> Payments CreatePayments(ctx).XeroTenantId(xeroTenantId).Payments(payments).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates multiple payments for invoices or credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	payments := *openapiclient.NewPayments() // Payments | Payments array with Payment object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePayments(context.Background()).XeroTenantId(xeroTenantId).Payments(payments).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePayments`: Payments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **payments** | [**Payments**](Payments.md) | Payments array with Payment object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrepaymentAllocations

> Allocations CreatePrepaymentAllocations(ctx, prepaymentID).XeroTenantId(xeroTenantId).Allocations(allocations).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Allows you to create an Allocation for prepayments

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	prepaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a PrePayment
	allocations := *openapiclient.NewAllocations() // Allocations | Allocations with an array of Allocation object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePrepaymentAllocations(context.Background(), prepaymentID).XeroTenantId(xeroTenantId).Allocations(allocations).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePrepaymentAllocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrepaymentAllocations`: Allocations
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePrepaymentAllocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prepaymentID** | **string** | Unique identifier for a PrePayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrepaymentAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **allocations** | [**Allocations**](Allocations.md) | Allocations with an array of Allocation object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Allocations**](Allocations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrepaymentHistory

> HistoryRecords CreatePrepaymentHistory(ctx, prepaymentID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific prepayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	prepaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a PrePayment
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePrepaymentHistory(context.Background(), prepaymentID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePrepaymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrepaymentHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePrepaymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prepaymentID** | **string** | Unique identifier for a PrePayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrepaymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePurchaseOrderAttachmentByFileName

> Attachments CreatePurchaseOrderAttachmentByFileName(ctx, purchaseOrderID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates attachment for a specific purchase order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePurchaseOrderAttachmentByFileName(context.Background(), purchaseOrderID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePurchaseOrderAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePurchaseOrderAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePurchaseOrderAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePurchaseOrderAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePurchaseOrderHistory

> HistoryRecords CreatePurchaseOrderHistory(ctx, purchaseOrderID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific purchase orders

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePurchaseOrderHistory(context.Background(), purchaseOrderID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePurchaseOrderHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePurchaseOrderHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePurchaseOrderHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePurchaseOrderHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePurchaseOrders

> PurchaseOrders CreatePurchaseOrders(ctx).XeroTenantId(xeroTenantId).PurchaseOrders(purchaseOrders).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates one or more purchase orders

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrders := *openapiclient.NewPurchaseOrders() // PurchaseOrders | PurchaseOrders with an array of PurchaseOrder object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreatePurchaseOrders(context.Background()).XeroTenantId(xeroTenantId).PurchaseOrders(purchaseOrders).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreatePurchaseOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePurchaseOrders`: PurchaseOrders
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreatePurchaseOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePurchaseOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **purchaseOrders** | [**PurchaseOrders**](PurchaseOrders.md) | PurchaseOrders with an array of PurchaseOrder object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuoteAttachmentByFileName

> Attachments CreateQuoteAttachmentByFileName(ctx, quoteID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates attachment for a specific quote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateQuoteAttachmentByFileName(context.Background(), quoteID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateQuoteAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuoteAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateQuoteAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuoteAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuoteHistory

> HistoryRecords CreateQuoteHistory(ctx, quoteID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific quote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateQuoteHistory(context.Background(), quoteID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateQuoteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuoteHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateQuoteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuoteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuotes

> Quotes CreateQuotes(ctx).XeroTenantId(xeroTenantId).Quotes(quotes).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Create one or more quotes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quotes := *openapiclient.NewQuotes() // Quotes | Quotes with an array of Quote object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateQuotes(context.Background()).XeroTenantId(xeroTenantId).Quotes(quotes).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuotes`: Quotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **quotes** | [**Quotes**](Quotes.md) | Quotes with an array of Quote object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReceipt

> Receipts CreateReceipt(ctx).XeroTenantId(xeroTenantId).Receipts(receipts).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Creates draft expense claim receipts for any user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receipts := *openapiclient.NewReceipts() // Receipts | Receipts with an array of Receipt object in body of request
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateReceipt(context.Background()).XeroTenantId(xeroTenantId).Receipts(receipts).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReceipt`: Receipts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **receipts** | [**Receipts**](Receipts.md) | Receipts with an array of Receipt object in body of request | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Receipts**](Receipts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReceiptAttachmentByFileName

> Attachments CreateReceiptAttachmentByFileName(ctx, receiptID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates an attachment on a specific expense claim receipts by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateReceiptAttachmentByFileName(context.Background(), receiptID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateReceiptAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReceiptAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateReceiptAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReceiptHistory

> HistoryRecords CreateReceiptHistory(ctx, receiptID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a history record for a specific receipt

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateReceiptHistory(context.Background(), receiptID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateReceiptHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReceiptHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateReceiptHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepeatingInvoiceAttachmentByFileName

> Attachments CreateRepeatingInvoiceAttachmentByFileName(ctx, repeatingInvoiceID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates an attachment from a specific repeating invoices by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateRepeatingInvoiceAttachmentByFileName(context.Background(), repeatingInvoiceID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateRepeatingInvoiceAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRepeatingInvoiceAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateRepeatingInvoiceAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepeatingInvoiceAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepeatingInvoiceHistory

> HistoryRecords CreateRepeatingInvoiceHistory(ctx, repeatingInvoiceID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()

Creates a  history record for a specific repeating invoice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice
	historyRecords := *openapiclient.NewHistoryRecords() // HistoryRecords | HistoryRecords containing an array of HistoryRecord objects in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateRepeatingInvoiceHistory(context.Background(), repeatingInvoiceID).XeroTenantId(xeroTenantId).HistoryRecords(historyRecords).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateRepeatingInvoiceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRepeatingInvoiceHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateRepeatingInvoiceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepeatingInvoiceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **historyRecords** | [**HistoryRecords**](HistoryRecords.md) | HistoryRecords containing an array of HistoryRecord objects in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepeatingInvoices

> RepeatingInvoices CreateRepeatingInvoices(ctx).XeroTenantId(xeroTenantId).RepeatingInvoices(repeatingInvoices).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates one or more repeating invoice templates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoices := *openapiclient.NewRepeatingInvoices() // RepeatingInvoices | RepeatingInvoices with an array of repeating invoice objects in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateRepeatingInvoices(context.Background()).XeroTenantId(xeroTenantId).RepeatingInvoices(repeatingInvoices).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateRepeatingInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRepeatingInvoices`: RepeatingInvoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateRepeatingInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepeatingInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **repeatingInvoices** | [**RepeatingInvoices**](RepeatingInvoices.md) | RepeatingInvoices with an array of repeating invoice objects in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**RepeatingInvoices**](RepeatingInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaxRates

> TaxRates CreateTaxRates(ctx).XeroTenantId(xeroTenantId).TaxRates(taxRates).IdempotencyKey(idempotencyKey).Execute()

Creates one or more tax rates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	taxRates := *openapiclient.NewTaxRates() // TaxRates | TaxRates array with TaxRate object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateTaxRates(context.Background()).XeroTenantId(xeroTenantId).TaxRates(taxRates).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateTaxRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaxRates`: TaxRates
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateTaxRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaxRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **taxRates** | [**TaxRates**](TaxRates.md) | TaxRates array with TaxRate object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**TaxRates**](TaxRates.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackingCategory

> TrackingCategories CreateTrackingCategory(ctx).XeroTenantId(xeroTenantId).TrackingCategory(trackingCategory).IdempotencyKey(idempotencyKey).Execute()

Create tracking categories

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	trackingCategory := *openapiclient.NewTrackingCategory() // TrackingCategory | TrackingCategory object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateTrackingCategory(context.Background()).XeroTenantId(xeroTenantId).TrackingCategory(trackingCategory).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateTrackingCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrackingCategory`: TrackingCategories
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateTrackingCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackingCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **trackingCategory** | [**TrackingCategory**](TrackingCategory.md) | TrackingCategory object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackingOptions

> TrackingOptions CreateTrackingOptions(ctx, trackingCategoryID).XeroTenantId(xeroTenantId).TrackingOption(trackingOption).IdempotencyKey(idempotencyKey).Execute()

Creates options for a specific tracking category

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	trackingCategoryID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a TrackingCategory
	trackingOption := *openapiclient.NewTrackingOption() // TrackingOption | TrackingOption object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.CreateTrackingOptions(context.Background(), trackingCategoryID).XeroTenantId(xeroTenantId).TrackingOption(trackingOption).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.CreateTrackingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrackingOptions`: TrackingOptions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.CreateTrackingOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingCategoryID** | **string** | Unique identifier for a TrackingCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackingOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **trackingOption** | [**TrackingOption**](TrackingOption.md) | TrackingOption object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**TrackingOptions**](TrackingOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> Accounts DeleteAccount(ctx, accountID).XeroTenantId(xeroTenantId).Execute()

Deletes a chart of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	accountID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Account object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeleteAccount(context.Background(), accountID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccount`: Accounts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Unique identifier for Account object | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBatchPayment

> BatchPayments DeleteBatchPayment(ctx).XeroTenantId(xeroTenantId).BatchPaymentDelete(batchPaymentDelete).IdempotencyKey(idempotencyKey).Execute()

Updates a specific batch payment for invoices and credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	batchPaymentDelete := *openapiclient.NewBatchPaymentDelete("BatchPaymentID_example", "Status_example") // BatchPaymentDelete | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeleteBatchPayment(context.Background()).XeroTenantId(xeroTenantId).BatchPaymentDelete(batchPaymentDelete).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteBatchPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBatchPayment`: BatchPayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeleteBatchPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **batchPaymentDelete** | [**BatchPaymentDelete**](BatchPaymentDelete.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**BatchPayments**](BatchPayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBatchPaymentByUrlParam

> BatchPayments DeleteBatchPaymentByUrlParam(ctx, batchPaymentID).XeroTenantId(xeroTenantId).BatchPaymentDeleteByUrlParam(batchPaymentDeleteByUrlParam).IdempotencyKey(idempotencyKey).Execute()

Updates a specific batch payment for invoices and credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	batchPaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for BatchPayment
	batchPaymentDeleteByUrlParam := *openapiclient.NewBatchPaymentDeleteByUrlParam("Status_example") // BatchPaymentDeleteByUrlParam | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeleteBatchPaymentByUrlParam(context.Background(), batchPaymentID).XeroTenantId(xeroTenantId).BatchPaymentDeleteByUrlParam(batchPaymentDeleteByUrlParam).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteBatchPaymentByUrlParam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBatchPaymentByUrlParam`: BatchPayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeleteBatchPaymentByUrlParam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchPaymentID** | **string** | Unique identifier for BatchPayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchPaymentByUrlParamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **batchPaymentDeleteByUrlParam** | [**BatchPaymentDeleteByUrlParam**](BatchPaymentDeleteByUrlParam.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**BatchPayments**](BatchPayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactGroupContact

> DeleteContactGroupContact(ctx, contactGroupID, contactID).XeroTenantId(xeroTenantId).Execute()

Deletes a specific contact from a contact group using a unique contact Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactGroupID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact Group
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountingAPI.DeleteContactGroupContact(context.Background(), contactGroupID, contactID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteContactGroupContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactGroupID** | **string** | Unique identifier for a Contact Group | 
**contactID** | **string** | Unique identifier for a Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactGroupContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactGroupContacts

> DeleteContactGroupContacts(ctx, contactGroupID).XeroTenantId(xeroTenantId).Execute()

Deletes all contacts from a specific contact group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactGroupID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountingAPI.DeleteContactGroupContacts(context.Background(), contactGroupID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteContactGroupContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactGroupID** | **string** | Unique identifier for a Contact Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactGroupContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCreditNoteAllocations

> Allocation DeleteCreditNoteAllocations(ctx, creditNoteID, allocationID).XeroTenantId(xeroTenantId).Execute()

Deletes an Allocation from a Credit Note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	allocationID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Allocation object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeleteCreditNoteAllocations(context.Background(), creditNoteID, allocationID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteCreditNoteAllocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCreditNoteAllocations`: Allocation
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeleteCreditNoteAllocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 
**allocationID** | **string** | Unique identifier for Allocation object | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCreditNoteAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 



### Return type

[**Allocation**](Allocation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItem

> DeleteItem(ctx, itemID).XeroTenantId(xeroTenantId).Execute()

Deletes a specific item

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	itemID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountingAPI.DeleteItem(context.Background(), itemID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemID** | **string** | Unique identifier for an Item | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLinkedTransaction

> DeleteLinkedTransaction(ctx, linkedTransactionID).XeroTenantId(xeroTenantId).Execute()

Deletes a specific linked transactions (billable expenses)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	linkedTransactionID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a LinkedTransaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountingAPI.DeleteLinkedTransaction(context.Background(), linkedTransactionID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteLinkedTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkedTransactionID** | **string** | Unique identifier for a LinkedTransaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOverpaymentAllocations

> Allocation DeleteOverpaymentAllocations(ctx, overpaymentID, allocationID).XeroTenantId(xeroTenantId).Execute()

Deletes an Allocation from an overpayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	overpaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Overpayment
	allocationID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Allocation object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeleteOverpaymentAllocations(context.Background(), overpaymentID, allocationID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteOverpaymentAllocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOverpaymentAllocations`: Allocation
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeleteOverpaymentAllocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**overpaymentID** | **string** | Unique identifier for a Overpayment | 
**allocationID** | **string** | Unique identifier for Allocation object | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOverpaymentAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 



### Return type

[**Allocation**](Allocation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePayment

> Payments DeletePayment(ctx, paymentID).XeroTenantId(xeroTenantId).PaymentDelete(paymentDelete).IdempotencyKey(idempotencyKey).Execute()

Updates a specific payment for invoices and credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	paymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Payment
	paymentDelete := *openapiclient.NewPaymentDelete("Status_example") // PaymentDelete | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeletePayment(context.Background(), paymentID).XeroTenantId(xeroTenantId).PaymentDelete(paymentDelete).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeletePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePayment`: Payments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeletePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentID** | **string** | Unique identifier for a Payment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **paymentDelete** | [**PaymentDelete**](PaymentDelete.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrepaymentAllocations

> Allocation DeletePrepaymentAllocations(ctx, prepaymentID, allocationID).XeroTenantId(xeroTenantId).Execute()

Deletes an Allocation from a Prepayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	prepaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a PrePayment
	allocationID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Allocation object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeletePrepaymentAllocations(context.Background(), prepaymentID, allocationID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeletePrepaymentAllocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePrepaymentAllocations`: Allocation
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeletePrepaymentAllocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prepaymentID** | **string** | Unique identifier for a PrePayment | 
**allocationID** | **string** | Unique identifier for Allocation object | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrepaymentAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 



### Return type

[**Allocation**](Allocation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrackingCategory

> TrackingCategories DeleteTrackingCategory(ctx, trackingCategoryID).XeroTenantId(xeroTenantId).Execute()

Deletes a specific tracking category

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	trackingCategoryID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a TrackingCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeleteTrackingCategory(context.Background(), trackingCategoryID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteTrackingCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrackingCategory`: TrackingCategories
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeleteTrackingCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingCategoryID** | **string** | Unique identifier for a TrackingCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrackingCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrackingOptions

> TrackingOptions DeleteTrackingOptions(ctx, trackingCategoryID, trackingOptionID).XeroTenantId(xeroTenantId).Execute()

Deletes a specific option for a specific tracking category

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	trackingCategoryID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a TrackingCategory
	trackingOptionID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Tracking Option

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.DeleteTrackingOptions(context.Background(), trackingCategoryID, trackingOptionID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.DeleteTrackingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrackingOptions`: TrackingOptions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.DeleteTrackingOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingCategoryID** | **string** | Unique identifier for a TrackingCategory | 
**trackingOptionID** | **string** | Unique identifier for a Tracking Option | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrackingOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 



### Return type

[**TrackingOptions**](TrackingOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailInvoice

> EmailInvoice(ctx, invoiceID).XeroTenantId(xeroTenantId).RequestEmpty(requestEmpty).IdempotencyKey(idempotencyKey).Execute()

Sends a copy of a specific invoice to related contact via email

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice
	requestEmpty := *openapiclient.NewRequestEmpty() // RequestEmpty | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountingAPI.EmailInvoice(context.Background(), invoiceID).XeroTenantId(xeroTenantId).RequestEmpty(requestEmpty).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.EmailInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **requestEmpty** | [**RequestEmpty**](RequestEmpty.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Accounts GetAccount(ctx, accountID).XeroTenantId(xeroTenantId).Execute()

Retrieves a single chart of accounts by using a unique account Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	accountID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Account object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetAccount(context.Background(), accountID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Accounts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Unique identifier for Account object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAttachmentByFileName

> *os.File GetAccountAttachmentByFileName(ctx, accountID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves an attachment for a specific account by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	accountID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Account object
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetAccountAttachmentByFileName(context.Background(), accountID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetAccountAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetAccountAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Unique identifier for Account object | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAttachmentById

> *os.File GetAccountAttachmentById(ctx, accountID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific account using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	accountID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Account object
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetAccountAttachmentById(context.Background(), accountID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetAccountAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetAccountAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Unique identifier for Account object | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAttachments

> Attachments GetAccountAttachments(ctx, accountID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments for a specific accounts by using a unique account Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	accountID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Account object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetAccountAttachments(context.Background(), accountID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetAccountAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetAccountAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Unique identifier for Account object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> Accounts GetAccounts(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()

Retrieves the full chart of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status==&quot;ACTIVE&quot; AND Type==&quot;BANK&quot;" // string | Filter by an any element (optional)
	order := "Name ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetAccounts(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccounts`: Accounts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransaction

> BankTransactions GetBankTransaction(ctx, bankTransactionID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()

Retrieves a single spent or received money transaction by using a unique bank transaction Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransaction(context.Background(), bankTransactionID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransaction`: BankTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionAttachmentByFileName

> *os.File GetBankTransactionAttachmentByFileName(ctx, bankTransactionID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific bank transaction by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransactionAttachmentByFileName(context.Background(), bankTransactionID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransactionAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactionAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransactionAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionAttachmentById

> *os.File GetBankTransactionAttachmentById(ctx, bankTransactionID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves specific attachments from a specific BankTransaction using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransactionAttachmentById(context.Background(), bankTransactionID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransactionAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactionAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransactionAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionAttachments

> Attachments GetBankTransactionAttachments(ctx, bankTransactionID).XeroTenantId(xeroTenantId).Execute()

Retrieves any attachments from a specific bank transactions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransactionAttachments(context.Background(), bankTransactionID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransactionAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactionAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransactionAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactions

> BankTransactions GetBankTransactions(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Unitdp(unitdp).Execute()

Retrieves any spent or received money transactions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="AUTHORISED"" // string | Filter by an any element (optional)
	order := "Type ASC" // string | Order by an any element (optional)
	page := int32(1) // int32 | Up to 100 bank transactions will be returned in a single API call with line items details (optional)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransactions(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactions`: BankTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **page** | **int32** | Up to 100 bank transactions will be returned in a single API call with line items details | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionsHistory

> HistoryRecords GetBankTransactionsHistory(ctx, bankTransactionID).XeroTenantId(xeroTenantId).Execute()

Retrieves history from a specific bank transaction using a unique bank transaction Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransactionsHistory(context.Background(), bankTransactionID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransactionsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactionsHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransactionsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransfer

> BankTransfers GetBankTransfer(ctx, bankTransferID).XeroTenantId(xeroTenantId).Execute()

Retrieves specific bank transfers by using a unique bank transfer Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransferID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transfer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransfer(context.Background(), bankTransferID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransfer`: BankTransfers
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransferID** | **string** | Xero generated unique identifier for a bank transfer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**BankTransfers**](BankTransfers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransferAttachmentByFileName

> *os.File GetBankTransferAttachmentByFileName(ctx, bankTransferID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment on a specific bank transfer by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransferID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transfer
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransferAttachmentByFileName(context.Background(), bankTransferID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransferAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransferAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransferAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransferID** | **string** | Xero generated unique identifier for a bank transfer | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransferAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransferAttachmentById

> *os.File GetBankTransferAttachmentById(ctx, bankTransferID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific bank transfer using a unique attachment ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransferID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transfer
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransferAttachmentById(context.Background(), bankTransferID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransferAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransferAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransferAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransferID** | **string** | Xero generated unique identifier for a bank transfer | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransferAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransferAttachments

> Attachments GetBankTransferAttachments(ctx, bankTransferID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments from a specific bank transfer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransferID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transfer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransferAttachments(context.Background(), bankTransferID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransferAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransferAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransferAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransferID** | **string** | Xero generated unique identifier for a bank transfer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransferAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransferHistory

> HistoryRecords GetBankTransferHistory(ctx, bankTransferID).XeroTenantId(xeroTenantId).Execute()

Retrieves history from a specific bank transfer using a unique bank transfer Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransferID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transfer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransferHistory(context.Background(), bankTransferID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransferHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransferHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransferHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransferID** | **string** | Xero generated unique identifier for a bank transfer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransferHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransfers

> BankTransfers GetBankTransfers(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()

Retrieves all bank transfers

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "HasAttachments==true" // string | Filter by an any element (optional)
	order := "Amount ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBankTransfers(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBankTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransfers`: BankTransfers
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBankTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**BankTransfers**](BankTransfers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPayment

> BatchPayments GetBatchPayment(ctx, batchPaymentID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific batch payment using a unique batch payment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	batchPaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for BatchPayment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBatchPayment(context.Background(), batchPaymentID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBatchPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchPayment`: BatchPayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBatchPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchPaymentID** | **string** | Unique identifier for BatchPayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**BatchPayments**](BatchPayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPaymentHistory

> HistoryRecords GetBatchPaymentHistory(ctx, batchPaymentID).XeroTenantId(xeroTenantId).Execute()

Retrieves history from a specific batch payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	batchPaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for BatchPayment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBatchPaymentHistory(context.Background(), batchPaymentID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBatchPaymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchPaymentHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBatchPaymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchPaymentID** | **string** | Unique identifier for BatchPayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchPaymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPayments

> BatchPayments GetBatchPayments(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()

Retrieves either one or many batch payments for invoices

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="AUTHORISED"" // string | Filter by an any element (optional)
	order := "Date ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBatchPayments(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBatchPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchPayments`: BatchPayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBatchPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**BatchPayments**](BatchPayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingTheme

> BrandingThemes GetBrandingTheme(ctx, brandingThemeID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific branding theme using a unique branding theme Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	brandingThemeID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Branding Theme

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBrandingTheme(context.Background(), brandingThemeID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBrandingTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandingTheme`: BrandingThemes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBrandingTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandingThemeID** | **string** | Unique identifier for a Branding Theme | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**BrandingThemes**](BrandingThemes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingThemePaymentServices

> PaymentServices GetBrandingThemePaymentServices(ctx, brandingThemeID).XeroTenantId(xeroTenantId).Execute()

Retrieves the payment services for a specific branding theme

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	brandingThemeID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Branding Theme

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBrandingThemePaymentServices(context.Background(), brandingThemeID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBrandingThemePaymentServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandingThemePaymentServices`: PaymentServices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBrandingThemePaymentServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandingThemeID** | **string** | Unique identifier for a Branding Theme | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingThemePaymentServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingThemes

> BrandingThemes GetBrandingThemes(ctx).XeroTenantId(xeroTenantId).Execute()

Retrieves all the branding themes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBrandingThemes(context.Background()).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBrandingThemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandingThemes`: BrandingThemes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBrandingThemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingThemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

### Return type

[**BrandingThemes**](BrandingThemes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudget

> Budgets GetBudget(ctx, budgetID).XeroTenantId(xeroTenantId).DateTo(dateTo).DateFrom(dateFrom).Execute()

Retrieves a specific budget, which includes budget lines

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	budgetID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Budgets
	dateTo := time.Now() // string | Filter by start date (optional)
	dateFrom := time.Now() // string | Filter by end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBudget(context.Background(), budgetID).XeroTenantId(xeroTenantId).DateTo(dateTo).DateFrom(dateFrom).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudget`: Budgets
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetID** | **string** | Unique identifier for Budgets | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **dateTo** | **string** | Filter by start date | 
 **dateFrom** | **string** | Filter by end date | 

### Return type

[**Budgets**](Budgets.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgets

> Budgets GetBudgets(ctx).XeroTenantId(xeroTenantId).IDs(iDs).DateTo(dateTo).DateFrom(dateFrom).Execute()

Retrieve a list of budgets

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	iDs := []string{"Inner_example"} // []string | Filter by BudgetID. Allows you to retrieve a specific individual budget. (optional)
	dateTo := time.Now() // string | Filter by start date (optional)
	dateFrom := time.Now() // string | Filter by end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetBudgets(context.Background()).XeroTenantId(xeroTenantId).IDs(iDs).DateTo(dateTo).DateFrom(dateFrom).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgets`: Budgets
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **iDs** | **[]string** | Filter by BudgetID. Allows you to retrieve a specific individual budget. | 
 **dateTo** | **string** | Filter by start date | 
 **dateFrom** | **string** | Filter by end date | 

### Return type

[**Budgets**](Budgets.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContact

> Contacts GetContact(ctx, contactID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific contacts in a Xero organisation using a unique contact Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContact(context.Background(), contactID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContact`: Contacts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactAttachmentByFileName

> *os.File GetContactAttachmentByFileName(ctx, contactID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific contact by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContactAttachmentByFileName(context.Background(), contactID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContactAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContactAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactAttachmentById

> *os.File GetContactAttachmentById(ctx, contactID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific contact using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContactAttachmentById(context.Background(), contactID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContactAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContactAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactAttachments

> Attachments GetContactAttachments(ctx, contactID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments for a specific contact in a Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContactAttachments(context.Background(), contactID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContactAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContactAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactByContactNumber

> Contacts GetContactByContactNumber(ctx, contactNumber).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific contact by contact number in a Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactNumber := "SB2" // string | This field is read only on the Xero contact screen, used to identify contacts in external systems (max length = 50).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContactByContactNumber(context.Background(), contactNumber).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContactByContactNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactByContactNumber`: Contacts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContactByContactNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactNumber** | **string** | This field is read only on the Xero contact screen, used to identify contacts in external systems (max length &#x3D; 50). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactByContactNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactCISSettings

> CISSettings GetContactCISSettings(ctx, contactID).XeroTenantId(xeroTenantId).Execute()

Retrieves CIS settings for a specific contact in a Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContactCISSettings(context.Background(), contactID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContactCISSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactCISSettings`: CISSettings
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContactCISSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactCISSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**CISSettings**](CISSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactGroup

> ContactGroups GetContactGroup(ctx, contactGroupID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific contact group by using a unique contact group Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactGroupID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContactGroup(context.Background(), contactGroupID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContactGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactGroup`: ContactGroups
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContactGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactGroupID** | **string** | Unique identifier for a Contact Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**ContactGroups**](ContactGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactGroups

> ContactGroups GetContactGroups(ctx).XeroTenantId(xeroTenantId).Where(where).Order(order).Execute()

Retrieves the contact Id and name of each contact group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	where := "Status=="ACTIVE"" // string | Filter by an any element (optional)
	order := "Name ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContactGroups(context.Background()).XeroTenantId(xeroTenantId).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContactGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactGroups`: ContactGroups
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContactGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**ContactGroups**](ContactGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactHistory

> HistoryRecords GetContactHistory(ctx, contactID).XeroTenantId(xeroTenantId).Execute()

Retrieves history records for a specific contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContactHistory(context.Background(), contactID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContactHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContactHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContacts

> Contacts GetContacts(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).IDs(iDs).Page(page).IncludeArchived(includeArchived).SummaryOnly(summaryOnly).SearchTerm(searchTerm).Execute()

Retrieves all contacts in a Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "ContactStatus==&quot;ACTIVE&quot;" // string | Filter by an any element (optional)
	order := "Name ASC" // string | Order by an any element (optional)
	iDs := []string{"Inner_example"} // []string | Filter by a comma separated list of ContactIDs. Allows you to retrieve a specific set of contacts in a single call. (optional)
	page := int32(1) // int32 | e.g. page=1 - Up to 100 contacts will be returned in a single API call. (optional)
	includeArchived := true // bool | e.g. includeArchived=true - Contacts with a status of ARCHIVED will be included in the response (optional)
	summaryOnly := true // bool | Use summaryOnly=true in GET Contacts and Invoices endpoint to retrieve a smaller version of the response object. This returns only lightweight fields, excluding computation-heavy fields from the response, making the API calls quick and efficient. (optional) (default to false)
	searchTerm := "Joe Bloggs" // string | Search parameter that performs a case-insensitive text search across the Name, FirstName, LastName, ContactNumber and EmailAddress fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetContacts(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).IDs(iDs).Page(page).IncludeArchived(includeArchived).SummaryOnly(summaryOnly).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContacts`: Contacts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **iDs** | **[]string** | Filter by a comma separated list of ContactIDs. Allows you to retrieve a specific set of contacts in a single call. | 
 **page** | **int32** | e.g. page&#x3D;1 - Up to 100 contacts will be returned in a single API call. | 
 **includeArchived** | **bool** | e.g. includeArchived&#x3D;true - Contacts with a status of ARCHIVED will be included in the response | 
 **summaryOnly** | **bool** | Use summaryOnly&#x3D;true in GET Contacts and Invoices endpoint to retrieve a smaller version of the response object. This returns only lightweight fields, excluding computation-heavy fields from the response, making the API calls quick and efficient. | [default to false]
 **searchTerm** | **string** | Search parameter that performs a case-insensitive text search across the Name, FirstName, LastName, ContactNumber and EmailAddress fields. | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNote

> CreditNotes GetCreditNote(ctx, creditNoteID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()

Retrieves a specific credit note using a unique credit note Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetCreditNote(context.Background(), creditNoteID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetCreditNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditNote`: CreditNotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteAsPdf

> *os.File GetCreditNoteAsPdf(ctx, creditNoteID).XeroTenantId(xeroTenantId).Execute()

Retrieves credit notes as PDF files

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetCreditNoteAsPdf(context.Background(), creditNoteID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetCreditNoteAsPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditNoteAsPdf`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetCreditNoteAsPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNoteAsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteAttachmentByFileName

> *os.File GetCreditNoteAttachmentByFileName(ctx, creditNoteID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment on a specific credit note by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetCreditNoteAttachmentByFileName(context.Background(), creditNoteID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetCreditNoteAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditNoteAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetCreditNoteAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNoteAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteAttachmentById

> *os.File GetCreditNoteAttachmentById(ctx, creditNoteID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific credit note using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetCreditNoteAttachmentById(context.Background(), creditNoteID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetCreditNoteAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditNoteAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetCreditNoteAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNoteAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteAttachments

> Attachments GetCreditNoteAttachments(ctx, creditNoteID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments for a specific credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetCreditNoteAttachments(context.Background(), creditNoteID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetCreditNoteAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditNoteAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetCreditNoteAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNoteAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteHistory

> HistoryRecords GetCreditNoteHistory(ctx, creditNoteID).XeroTenantId(xeroTenantId).Execute()

Retrieves history records of a specific credit note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetCreditNoteHistory(context.Background(), creditNoteID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetCreditNoteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditNoteHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetCreditNoteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNoteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNotes

> CreditNotes GetCreditNotes(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Unitdp(unitdp).Execute()

Retrieves any credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="DRAFT"" // string | Filter by an any element (optional)
	order := "CreditNoteNumber ASC" // string | Order by an any element (optional)
	page := int32(1) // int32 | e.g. page=1 – Up to 100 credit notes will be returned in a single API call with line items shown for each credit note (optional)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetCreditNotes(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetCreditNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreditNotes`: CreditNotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetCreditNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **page** | **int32** | e.g. page&#x3D;1 – Up to 100 credit notes will be returned in a single API call with line items shown for each credit note | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrencies

> Currencies GetCurrencies(ctx).XeroTenantId(xeroTenantId).Where(where).Order(order).Execute()

Retrieves currencies for your Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	where := "Code=="USD"" // string | Filter by an any element (optional)
	order := "Code ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetCurrencies(context.Background()).XeroTenantId(xeroTenantId).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrencies`: Currencies
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetCurrencies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**Currencies**](Currencies.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployee

> Employees GetEmployee(ctx, employeeID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific employee used in Xero payrun using a unique employee Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	employeeID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Employee

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetEmployee(context.Background(), employeeID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetEmployee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployee`: Employees
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetEmployee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employeeID** | **string** | Unique identifier for a Employee | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Employees**](Employees.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployees

> Employees GetEmployees(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()

Retrieves employees used in Xero payrun

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="ACTIVE"" // string | Filter by an any element (optional)
	order := "LastName ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetEmployees(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetEmployees``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployees`: Employees
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetEmployees`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**Employees**](Employees.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseClaim

> ExpenseClaims GetExpenseClaim(ctx, expenseClaimID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific expense claim using a unique expense claim Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	expenseClaimID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ExpenseClaim

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetExpenseClaim(context.Background(), expenseClaimID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetExpenseClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenseClaim`: ExpenseClaims
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetExpenseClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseClaimID** | **string** | Unique identifier for a ExpenseClaim | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**ExpenseClaims**](ExpenseClaims.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseClaimHistory

> HistoryRecords GetExpenseClaimHistory(ctx, expenseClaimID).XeroTenantId(xeroTenantId).Execute()

Retrieves history records of a specific expense claim

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	expenseClaimID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ExpenseClaim

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetExpenseClaimHistory(context.Background(), expenseClaimID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetExpenseClaimHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenseClaimHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetExpenseClaimHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseClaimID** | **string** | Unique identifier for a ExpenseClaim | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseClaimHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseClaims

> ExpenseClaims GetExpenseClaims(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()

Retrieves expense claims

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="SUBMITTED"" // string | Filter by an any element (optional)
	order := "Status ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetExpenseClaims(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetExpenseClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenseClaims`: ExpenseClaims
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetExpenseClaims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**ExpenseClaims**](ExpenseClaims.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> Invoices GetInvoice(ctx, invoiceID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()

Retrieves a specific sales invoice or purchase bill using a unique invoice Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetInvoice(context.Background(), invoiceID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoice`: Invoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAsPdf

> *os.File GetInvoiceAsPdf(ctx, invoiceID).XeroTenantId(xeroTenantId).Execute()

Retrieves invoices or purchase bills as PDF files

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetInvoiceAsPdf(context.Background(), invoiceID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetInvoiceAsPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceAsPdf`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetInvoiceAsPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceAsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAttachmentByFileName

> *os.File GetInvoiceAttachmentByFileName(ctx, invoiceID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves an attachment from a specific invoice or purchase bill by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetInvoiceAttachmentByFileName(context.Background(), invoiceID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetInvoiceAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetInvoiceAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAttachmentById

> *os.File GetInvoiceAttachmentById(ctx, invoiceID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific invoices or purchase bills by using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetInvoiceAttachmentById(context.Background(), invoiceID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetInvoiceAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetInvoiceAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAttachments

> Attachments GetInvoiceAttachments(ctx, invoiceID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments for a specific invoice or purchase bill

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetInvoiceAttachments(context.Background(), invoiceID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetInvoiceAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetInvoiceAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceHistory

> HistoryRecords GetInvoiceHistory(ctx, invoiceID).XeroTenantId(xeroTenantId).Execute()

Retrieves history records for a specific invoice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetInvoiceHistory(context.Background(), invoiceID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetInvoiceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetInvoiceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceReminders

> InvoiceReminders GetInvoiceReminders(ctx).XeroTenantId(xeroTenantId).Execute()

Retrieves invoice reminder settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetInvoiceReminders(context.Background()).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetInvoiceReminders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceReminders`: InvoiceReminders
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetInvoiceReminders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRemindersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

### Return type

[**InvoiceReminders**](InvoiceReminders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> Invoices GetInvoices(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).IDs(iDs).InvoiceNumbers(invoiceNumbers).ContactIDs(contactIDs).Statuses(statuses).Page(page).IncludeArchived(includeArchived).CreatedByMyApp(createdByMyApp).Unitdp(unitdp).SummaryOnly(summaryOnly).Execute()

Retrieves sales invoices or purchase bills

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="DRAFT"" // string | Filter by an any element (optional)
	order := "InvoiceNumber ASC" // string | Order by an any element (optional)
	iDs := []string{"Inner_example"} // []string | Filter by a comma-separated list of InvoicesIDs. (optional)
	invoiceNumbers := []string{"Inner_example"} // []string | Filter by a comma-separated list of InvoiceNumbers. (optional)
	contactIDs := []string{"Inner_example"} // []string | Filter by a comma-separated list of ContactIDs. (optional)
	statuses := []string{"Inner_example"} // []string | Filter by a comma-separated list Statuses. For faster response times we recommend using these explicit parameters instead of passing OR conditions into the Where filter. (optional)
	page := int32(1) // int32 | e.g. page=1 – Up to 100 invoices will be returned in a single API call with line items shown for each invoice (optional)
	includeArchived := true // bool | e.g. includeArchived=true - Invoices with a status of ARCHIVED will be included in the response (optional)
	createdByMyApp := false // bool | When set to true you'll only retrieve Invoices created by your app (optional)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	summaryOnly := true // bool | Use summaryOnly=true in GET Contacts and Invoices endpoint to retrieve a smaller version of the response object. This returns only lightweight fields, excluding computation-heavy fields from the response, making the API calls quick and efficient. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetInvoices(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).IDs(iDs).InvoiceNumbers(invoiceNumbers).ContactIDs(contactIDs).Statuses(statuses).Page(page).IncludeArchived(includeArchived).CreatedByMyApp(createdByMyApp).Unitdp(unitdp).SummaryOnly(summaryOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoices`: Invoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **iDs** | **[]string** | Filter by a comma-separated list of InvoicesIDs. | 
 **invoiceNumbers** | **[]string** | Filter by a comma-separated list of InvoiceNumbers. | 
 **contactIDs** | **[]string** | Filter by a comma-separated list of ContactIDs. | 
 **statuses** | **[]string** | Filter by a comma-separated list Statuses. For faster response times we recommend using these explicit parameters instead of passing OR conditions into the Where filter. | 
 **page** | **int32** | e.g. page&#x3D;1 – Up to 100 invoices will be returned in a single API call with line items shown for each invoice | 
 **includeArchived** | **bool** | e.g. includeArchived&#x3D;true - Invoices with a status of ARCHIVED will be included in the response | 
 **createdByMyApp** | **bool** | When set to true you&#39;ll only retrieve Invoices created by your app | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **summaryOnly** | **bool** | Use summaryOnly&#x3D;true in GET Contacts and Invoices endpoint to retrieve a smaller version of the response object. This returns only lightweight fields, excluding computation-heavy fields from the response, making the API calls quick and efficient. | [default to false]

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem

> Items GetItem(ctx, itemID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()

Retrieves a specific item using a unique item Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	itemID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Item
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetItem(context.Background(), itemID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItem`: Items
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemID** | **string** | Unique identifier for an Item | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemHistory

> HistoryRecords GetItemHistory(ctx, itemID).XeroTenantId(xeroTenantId).Execute()

Retrieves history for a specific item

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	itemID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetItemHistory(context.Background(), itemID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetItemHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetItemHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemID** | **string** | Unique identifier for an Item | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> Items GetItems(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Unitdp(unitdp).Execute()

Retrieves items

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "IsSold==true" // string | Filter by an any element (optional)
	order := "Code ASC" // string | Order by an any element (optional)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetItems(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItems`: Items
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournal

> Journals GetJournal(ctx, journalID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific journal using a unique journal Id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	journalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Journal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetJournal(context.Background(), journalID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetJournal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournal`: Journals
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetJournal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalID** | **string** | Unique identifier for a Journal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Journals**](Journals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournalByNumber

> Journals GetJournalByNumber(ctx, journalNumber).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific journal using a unique journal number.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	journalNumber := int32(1000) // int32 | Number of a Journal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetJournalByNumber(context.Background(), journalNumber).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetJournalByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournalByNumber`: Journals
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetJournalByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalNumber** | **int32** | Number of a Journal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Journals**](Journals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournals

> Journals GetJournals(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Offset(offset).PaymentsOnly(paymentsOnly).Execute()

Retrieves journals

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	offset := int32(10) // int32 | Offset by a specified journal number. e.g. journals with a JournalNumber greater than the offset will be returned (optional)
	paymentsOnly := true // bool | Filter to retrieve journals on a cash basis. Journals are returned on an accrual basis by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetJournals(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Offset(offset).PaymentsOnly(paymentsOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetJournals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournals`: Journals
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetJournals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **offset** | **int32** | Offset by a specified journal number. e.g. journals with a JournalNumber greater than the offset will be returned | 
 **paymentsOnly** | **bool** | Filter to retrieve journals on a cash basis. Journals are returned on an accrual basis by default. | 

### Return type

[**Journals**](Journals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedTransaction

> LinkedTransactions GetLinkedTransaction(ctx, linkedTransactionID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific linked transaction (billable expenses) using a unique linked transaction Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	linkedTransactionID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a LinkedTransaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetLinkedTransaction(context.Background(), linkedTransactionID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetLinkedTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedTransaction`: LinkedTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetLinkedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkedTransactionID** | **string** | Unique identifier for a LinkedTransaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**LinkedTransactions**](LinkedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedTransactions

> LinkedTransactions GetLinkedTransactions(ctx).XeroTenantId(xeroTenantId).Page(page).LinkedTransactionID(linkedTransactionID).SourceTransactionID(sourceTransactionID).ContactID(contactID).Status(status).TargetTransactionID(targetTransactionID).Execute()

Retrieves linked transactions (billable expenses)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	page := int32(1) // int32 | Up to 100 linked transactions will be returned in a single API call. Use the page parameter to specify the page to be returned e.g. page=1. (optional)
	linkedTransactionID := "00000000-0000-0000-0000-000000000000" // string | The Xero identifier for an Linked Transaction (optional)
	sourceTransactionID := "00000000-0000-0000-0000-000000000000" // string | Filter by the SourceTransactionID. Get the linked transactions created from a particular ACCPAY invoice (optional)
	contactID := "00000000-0000-0000-0000-000000000000" // string | Filter by the ContactID. Get all the linked transactions that have been assigned to a particular customer. (optional)
	status := "APPROVED" // string | Filter by the combination of ContactID and Status. Get  the linked transactions associated to a  customer and with a status (optional)
	targetTransactionID := "00000000-0000-0000-0000-000000000000" // string | Filter by the TargetTransactionID. Get all the linked transactions allocated to a particular ACCREC invoice (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetLinkedTransactions(context.Background()).XeroTenantId(xeroTenantId).Page(page).LinkedTransactionID(linkedTransactionID).SourceTransactionID(sourceTransactionID).ContactID(contactID).Status(status).TargetTransactionID(targetTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetLinkedTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedTransactions`: LinkedTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetLinkedTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **page** | **int32** | Up to 100 linked transactions will be returned in a single API call. Use the page parameter to specify the page to be returned e.g. page&#x3D;1. | 
 **linkedTransactionID** | **string** | The Xero identifier for an Linked Transaction | 
 **sourceTransactionID** | **string** | Filter by the SourceTransactionID. Get the linked transactions created from a particular ACCPAY invoice | 
 **contactID** | **string** | Filter by the ContactID. Get all the linked transactions that have been assigned to a particular customer. | 
 **status** | **string** | Filter by the combination of ContactID and Status. Get  the linked transactions associated to a  customer and with a status | 
 **targetTransactionID** | **string** | Filter by the TargetTransactionID. Get all the linked transactions allocated to a particular ACCREC invoice | 

### Return type

[**LinkedTransactions**](LinkedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournal

> ManualJournals GetManualJournal(ctx, manualJournalID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific manual journal

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetManualJournal(context.Background(), manualJournalID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetManualJournal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualJournal`: ManualJournals
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetManualJournal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournalAttachmentByFileName

> *os.File GetManualJournalAttachmentByFileName(ctx, manualJournalID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific manual journal by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetManualJournalAttachmentByFileName(context.Background(), manualJournalID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetManualJournalAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualJournalAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetManualJournalAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualJournalAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournalAttachmentById

> *os.File GetManualJournalAttachmentById(ctx, manualJournalID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Allows you to retrieve a specific attachment from a specific manual journal using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetManualJournalAttachmentById(context.Background(), manualJournalID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetManualJournalAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualJournalAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetManualJournalAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualJournalAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournalAttachments

> Attachments GetManualJournalAttachments(ctx, manualJournalID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachment for a specific manual journal

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetManualJournalAttachments(context.Background(), manualJournalID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetManualJournalAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualJournalAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetManualJournalAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualJournalAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournals

> ManualJournals GetManualJournals(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Execute()

Retrieves manual journals

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="DRAFT"" // string | Filter by an any element (optional)
	order := "Date ASC" // string | Order by an any element (optional)
	page := int32(1) // int32 | e.g. page=1 – Up to 100 manual journals will be returned in a single API call with line items shown for each overpayment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetManualJournals(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetManualJournals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualJournals`: ManualJournals
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetManualJournals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManualJournalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **page** | **int32** | e.g. page&#x3D;1 – Up to 100 manual journals will be returned in a single API call with line items shown for each overpayment | 

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournalsHistory

> HistoryRecords GetManualJournalsHistory(ctx, manualJournalID).XeroTenantId(xeroTenantId).Execute()

Retrieves history for a specific manual journal

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetManualJournalsHistory(context.Background(), manualJournalID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetManualJournalsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualJournalsHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetManualJournalsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualJournalsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnlineInvoice

> OnlineInvoices GetOnlineInvoice(ctx, invoiceID).XeroTenantId(xeroTenantId).Execute()

Retrieves a URL to an online invoice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetOnlineInvoice(context.Background(), invoiceID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetOnlineInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnlineInvoice`: OnlineInvoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetOnlineInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnlineInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**OnlineInvoices**](OnlineInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisationActions

> Actions GetOrganisationActions(ctx).XeroTenantId(xeroTenantId).Execute()

Retrieves a list of the key actions your app has permission to perform in the connected Xero organisation.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetOrganisationActions(context.Background()).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetOrganisationActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganisationActions`: Actions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetOrganisationActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

### Return type

[**Actions**](Actions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisationCISSettings

> CISOrgSettings GetOrganisationCISSettings(ctx, organisationID).XeroTenantId(xeroTenantId).Execute()

Retrieves the CIS settings for the Xero organistaion.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	organisationID := "00000000-0000-0000-0000-000000000000" // string | The unique Xero identifier for an organisation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetOrganisationCISSettings(context.Background(), organisationID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetOrganisationCISSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganisationCISSettings`: CISOrgSettings
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetOrganisationCISSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationID** | **string** | The unique Xero identifier for an organisation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationCISSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**CISOrgSettings**](CISOrgSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisations

> Organisations GetOrganisations(ctx).XeroTenantId(xeroTenantId).Execute()

Retrieves Xero organisation details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetOrganisations(context.Background()).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganisations`: Organisations
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetOrganisations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

### Return type

[**Organisations**](Organisations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverpayment

> Overpayments GetOverpayment(ctx, overpaymentID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific overpayment using a unique overpayment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	overpaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Overpayment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetOverpayment(context.Background(), overpaymentID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetOverpayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverpayment`: Overpayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetOverpayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**overpaymentID** | **string** | Unique identifier for a Overpayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOverpaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Overpayments**](Overpayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverpaymentHistory

> HistoryRecords GetOverpaymentHistory(ctx, overpaymentID).XeroTenantId(xeroTenantId).Execute()

Retrieves history records of a specific overpayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	overpaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Overpayment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetOverpaymentHistory(context.Background(), overpaymentID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetOverpaymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverpaymentHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetOverpaymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**overpaymentID** | **string** | Unique identifier for a Overpayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOverpaymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverpayments

> Overpayments GetOverpayments(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Unitdp(unitdp).Execute()

Retrieves overpayments

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="AUTHORISED"" // string | Filter by an any element (optional)
	order := "Status ASC" // string | Order by an any element (optional)
	page := int32(1) // int32 | e.g. page=1 – Up to 100 overpayments will be returned in a single API call with line items shown for each overpayment (optional)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetOverpayments(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetOverpayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverpayments`: Overpayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetOverpayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOverpaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **page** | **int32** | e.g. page&#x3D;1 – Up to 100 overpayments will be returned in a single API call with line items shown for each overpayment | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Overpayments**](Overpayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayment

> Payments GetPayment(ctx, paymentID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific payment for invoices and credit notes using a unique payment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	paymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPayment(context.Background(), paymentID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayment`: Payments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentID** | **string** | Unique identifier for a Payment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentHistory

> HistoryRecords GetPaymentHistory(ctx, paymentID).XeroTenantId(xeroTenantId).Execute()

Retrieves history records of a specific payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	paymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPaymentHistory(context.Background(), paymentID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPaymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPaymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentID** | **string** | Unique identifier for a Payment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentServices

> PaymentServices GetPaymentServices(ctx).XeroTenantId(xeroTenantId).Execute()

Retrieves payment services

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPaymentServices(context.Background()).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPaymentServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentServices`: PaymentServices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPaymentServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayments

> Payments GetPayments(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Execute()

Retrieves payments for invoices and credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="AUTHORISED"" // string | Filter by an any element (optional)
	order := "Amount ASC" // string | Order by an any element (optional)
	page := int32(1) // int32 | Up to 100 payments will be returned in a single API call (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPayments(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayments`: Payments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **page** | **int32** | Up to 100 payments will be returned in a single API call | 

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrepayment

> Prepayments GetPrepayment(ctx, prepaymentID).XeroTenantId(xeroTenantId).Execute()

Allows you to retrieve a specified prepayments

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	prepaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a PrePayment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPrepayment(context.Background(), prepaymentID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPrepayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrepayment`: Prepayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPrepayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prepaymentID** | **string** | Unique identifier for a PrePayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrepaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Prepayments**](Prepayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrepaymentHistory

> HistoryRecords GetPrepaymentHistory(ctx, prepaymentID).XeroTenantId(xeroTenantId).Execute()

Retrieves history record for a specific prepayment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	prepaymentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a PrePayment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPrepaymentHistory(context.Background(), prepaymentID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPrepaymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrepaymentHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPrepaymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prepaymentID** | **string** | Unique identifier for a PrePayment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrepaymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrepayments

> Prepayments GetPrepayments(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Unitdp(unitdp).Execute()

Retrieves prepayments

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="AUTHORISED"" // string | Filter by an any element (optional)
	order := "Reference ASC" // string | Order by an any element (optional)
	page := int32(1) // int32 | e.g. page=1 – Up to 100 prepayments will be returned in a single API call with line items shown for each overpayment (optional)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPrepayments(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Page(page).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPrepayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrepayments`: Prepayments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPrepayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrepaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **page** | **int32** | e.g. page&#x3D;1 – Up to 100 prepayments will be returned in a single API call with line items shown for each overpayment | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Prepayments**](Prepayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrder

> PurchaseOrders GetPurchaseOrder(ctx, purchaseOrderID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific purchase order using a unique purchase order Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPurchaseOrder(context.Background(), purchaseOrderID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPurchaseOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPurchaseOrder`: PurchaseOrders
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPurchaseOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderAsPdf

> *os.File GetPurchaseOrderAsPdf(ctx, purchaseOrderID).XeroTenantId(xeroTenantId).Execute()

Retrieves specific purchase order as PDF files using a unique purchase order Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPurchaseOrderAsPdf(context.Background(), purchaseOrderID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPurchaseOrderAsPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPurchaseOrderAsPdf`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPurchaseOrderAsPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseOrderAsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderAttachmentByFileName

> *os.File GetPurchaseOrderAttachmentByFileName(ctx, purchaseOrderID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment for a specific purchase order by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPurchaseOrderAttachmentByFileName(context.Background(), purchaseOrderID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPurchaseOrderAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPurchaseOrderAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPurchaseOrderAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseOrderAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderAttachmentById

> *os.File GetPurchaseOrderAttachmentById(ctx, purchaseOrderID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves specific attachment for a specific purchase order using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPurchaseOrderAttachmentById(context.Background(), purchaseOrderID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPurchaseOrderAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPurchaseOrderAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPurchaseOrderAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseOrderAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderAttachments

> Attachments GetPurchaseOrderAttachments(ctx, purchaseOrderID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments for a specific purchase order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPurchaseOrderAttachments(context.Background(), purchaseOrderID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPurchaseOrderAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPurchaseOrderAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPurchaseOrderAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseOrderAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderByNumber

> PurchaseOrders GetPurchaseOrderByNumber(ctx, purchaseOrderNumber).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific purchase order using purchase order number

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderNumber := "PO1234" // string | Unique identifier for a PurchaseOrder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPurchaseOrderByNumber(context.Background(), purchaseOrderNumber).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPurchaseOrderByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPurchaseOrderByNumber`: PurchaseOrders
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPurchaseOrderByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderNumber** | **string** | Unique identifier for a PurchaseOrder | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseOrderByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderHistory

> HistoryRecords GetPurchaseOrderHistory(ctx, purchaseOrderID).XeroTenantId(xeroTenantId).Execute()

Retrieves history for a specific purchase order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPurchaseOrderHistory(context.Background(), purchaseOrderID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPurchaseOrderHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPurchaseOrderHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPurchaseOrderHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseOrderHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrders

> PurchaseOrders GetPurchaseOrders(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Status(status).DateFrom(dateFrom).DateTo(dateTo).Order(order).Page(page).Execute()

Retrieves purchase orders

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	status := "SUBMITTED" // string | Filter by purchase order status (optional)
	dateFrom := "2019-12-01" // string | Filter by purchase order date (e.g. GET https://.../PurchaseOrders?DateFrom=2015-12-01&DateTo=2015-12-31 (optional)
	dateTo := "2019-12-31" // string | Filter by purchase order date (e.g. GET https://.../PurchaseOrders?DateFrom=2015-12-01&DateTo=2015-12-31 (optional)
	order := "PurchaseOrderNumber ASC" // string | Order by an any element (optional)
	page := int32(1) // int32 | To specify a page, append the page parameter to the URL e.g. ?page=1. If there are 100 records in the response you will need to check if there is any more data by fetching the next page e.g ?page=2 and continuing this process until no more results are returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetPurchaseOrders(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Status(status).DateFrom(dateFrom).DateTo(dateTo).Order(order).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetPurchaseOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPurchaseOrders`: PurchaseOrders
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetPurchaseOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **status** | **string** | Filter by purchase order status | 
 **dateFrom** | **string** | Filter by purchase order date (e.g. GET https://.../PurchaseOrders?DateFrom&#x3D;2015-12-01&amp;DateTo&#x3D;2015-12-31 | 
 **dateTo** | **string** | Filter by purchase order date (e.g. GET https://.../PurchaseOrders?DateFrom&#x3D;2015-12-01&amp;DateTo&#x3D;2015-12-31 | 
 **order** | **string** | Order by an any element | 
 **page** | **int32** | To specify a page, append the page parameter to the URL e.g. ?page&#x3D;1. If there are 100 records in the response you will need to check if there is any more data by fetching the next page e.g ?page&#x3D;2 and continuing this process until no more results are returned. | 

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuote

> Quotes GetQuote(ctx, quoteID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific quote using a unique quote Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetQuote(context.Background(), quoteID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuote`: Quotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteAsPdf

> *os.File GetQuoteAsPdf(ctx, quoteID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific quote as a PDF file using a unique quote Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetQuoteAsPdf(context.Background(), quoteID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetQuoteAsPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteAsPdf`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetQuoteAsPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteAsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteAttachmentByFileName

> *os.File GetQuoteAttachmentByFileName(ctx, quoteID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific quote by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetQuoteAttachmentByFileName(context.Background(), quoteID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetQuoteAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetQuoteAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteAttachmentById

> *os.File GetQuoteAttachmentById(ctx, quoteID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific quote using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetQuoteAttachmentById(context.Background(), quoteID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetQuoteAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetQuoteAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteAttachments

> Attachments GetQuoteAttachments(ctx, quoteID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments for a specific quote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetQuoteAttachments(context.Background(), quoteID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetQuoteAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetQuoteAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteHistory

> HistoryRecords GetQuoteHistory(ctx, quoteID).XeroTenantId(xeroTenantId).Execute()

Retrieves history records of a specific quote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetQuoteHistory(context.Background(), quoteID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetQuoteHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetQuoteHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotes

> Quotes GetQuotes(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).DateFrom(dateFrom).DateTo(dateTo).ExpiryDateFrom(expiryDateFrom).ExpiryDateTo(expiryDateTo).ContactID(contactID).Status(status).Page(page).Order(order).QuoteNumber(quoteNumber).Execute()

Retrieves sales quotes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	dateFrom := time.Now() // string | Filter for quotes after a particular date (optional)
	dateTo := time.Now() // string | Filter for quotes before a particular date (optional)
	expiryDateFrom := time.Now() // string | Filter for quotes expiring after a particular date (optional)
	expiryDateTo := time.Now() // string | Filter for quotes before a particular date (optional)
	contactID := "00000000-0000-0000-0000-000000000000" // string | Filter for quotes belonging to a particular contact (optional)
	status := "DRAFT" // string | Filter for quotes of a particular Status (optional)
	page := int32(1) // int32 | e.g. page=1 – Up to 100 Quotes will be returned in a single API call with line items shown for each quote (optional)
	order := "Status ASC" // string | Order by an any element (optional)
	quoteNumber := "QU-0001" // string | Filter by quote number (e.g. GET https://.../Quotes?QuoteNumber=QU-0001) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetQuotes(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).DateFrom(dateFrom).DateTo(dateTo).ExpiryDateFrom(expiryDateFrom).ExpiryDateTo(expiryDateTo).ContactID(contactID).Status(status).Page(page).Order(order).QuoteNumber(quoteNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotes`: Quotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **dateFrom** | **string** | Filter for quotes after a particular date | 
 **dateTo** | **string** | Filter for quotes before a particular date | 
 **expiryDateFrom** | **string** | Filter for quotes expiring after a particular date | 
 **expiryDateTo** | **string** | Filter for quotes before a particular date | 
 **contactID** | **string** | Filter for quotes belonging to a particular contact | 
 **status** | **string** | Filter for quotes of a particular Status | 
 **page** | **int32** | e.g. page&#x3D;1 – Up to 100 Quotes will be returned in a single API call with line items shown for each quote | 
 **order** | **string** | Order by an any element | 
 **quoteNumber** | **string** | Filter by quote number (e.g. GET https://.../Quotes?QuoteNumber&#x3D;QU-0001) | 

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceipt

> Receipts GetReceipt(ctx, receiptID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()

Retrieves a specific draft expense claim receipt by using a unique receipt Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReceipt(context.Background(), receiptID).XeroTenantId(xeroTenantId).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceipt`: Receipts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Receipts**](Receipts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptAttachmentByFileName

> *os.File GetReceiptAttachmentByFileName(ctx, receiptID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific expense claim receipts by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReceiptAttachmentByFileName(context.Background(), receiptID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReceiptAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceiptAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReceiptAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptAttachmentById

> *os.File GetReceiptAttachmentById(ctx, receiptID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachments from a specific expense claim receipts by using a unique attachment Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReceiptAttachmentById(context.Background(), receiptID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReceiptAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceiptAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReceiptAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptAttachments

> Attachments GetReceiptAttachments(ctx, receiptID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments for a specific expense claim receipt

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReceiptAttachments(context.Background(), receiptID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReceiptAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceiptAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReceiptAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptHistory

> HistoryRecords GetReceiptHistory(ctx, receiptID).XeroTenantId(xeroTenantId).Execute()

Retrieves a history record for a specific receipt

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReceiptHistory(context.Background(), receiptID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReceiptHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceiptHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReceiptHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceipts

> Receipts GetReceipts(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Unitdp(unitdp).Execute()

Retrieves draft expense claim receipts for any user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "Status=="DRAFT"" // string | Filter by an any element (optional)
	order := "ReceiptNumber ASC" // string | Order by an any element (optional)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReceipts(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Unitdp(unitdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReceipts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceipts`: Receipts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReceipts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Receipts**](Receipts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoice

> RepeatingInvoices GetRepeatingInvoice(ctx, repeatingInvoiceID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific repeating invoice by using a unique repeating invoice Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetRepeatingInvoice(context.Background(), repeatingInvoiceID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetRepeatingInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepeatingInvoice`: RepeatingInvoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetRepeatingInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepeatingInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**RepeatingInvoices**](RepeatingInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoiceAttachmentByFileName

> *os.File GetRepeatingInvoiceAttachmentByFileName(ctx, repeatingInvoiceID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific repeating invoices by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice
	fileName := "xero-dev.jpg" // string | Name of the attachment
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetRepeatingInvoiceAttachmentByFileName(context.Background(), repeatingInvoiceID, fileName).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetRepeatingInvoiceAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepeatingInvoiceAttachmentByFileName`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetRepeatingInvoiceAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepeatingInvoiceAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoiceAttachmentById

> *os.File GetRepeatingInvoiceAttachmentById(ctx, repeatingInvoiceID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()

Retrieves a specific attachment from a specific repeating invoice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice
	attachmentID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Attachment object
	contentType := "image/jpg" // string | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetRepeatingInvoiceAttachmentById(context.Background(), repeatingInvoiceID, attachmentID).XeroTenantId(xeroTenantId).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetRepeatingInvoiceAttachmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepeatingInvoiceAttachmentById`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetRepeatingInvoiceAttachmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 
**attachmentID** | **string** | Unique identifier for Attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepeatingInvoiceAttachmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **contentType** | **string** | The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoiceAttachments

> Attachments GetRepeatingInvoiceAttachments(ctx, repeatingInvoiceID).XeroTenantId(xeroTenantId).Execute()

Retrieves attachments from a specific repeating invoice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetRepeatingInvoiceAttachments(context.Background(), repeatingInvoiceID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetRepeatingInvoiceAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepeatingInvoiceAttachments`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetRepeatingInvoiceAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepeatingInvoiceAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoiceHistory

> HistoryRecords GetRepeatingInvoiceHistory(ctx, repeatingInvoiceID).XeroTenantId(xeroTenantId).Execute()

Retrieves history record for a specific repeating invoice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetRepeatingInvoiceHistory(context.Background(), repeatingInvoiceID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetRepeatingInvoiceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepeatingInvoiceHistory`: HistoryRecords
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetRepeatingInvoiceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepeatingInvoiceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoices

> RepeatingInvoices GetRepeatingInvoices(ctx).XeroTenantId(xeroTenantId).Where(where).Order(order).Execute()

Retrieves repeating invoices

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	where := "Status=="DRAFT"" // string | Filter by an any element (optional)
	order := "Total ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetRepeatingInvoices(context.Background()).XeroTenantId(xeroTenantId).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetRepeatingInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepeatingInvoices`: RepeatingInvoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetRepeatingInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepeatingInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**RepeatingInvoices**](RepeatingInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportAgedPayablesByContact

> ReportWithRows GetReportAgedPayablesByContact(ctx).XeroTenantId(xeroTenantId).ContactId(contactId).Date(date).FromDate(fromDate).ToDate(toDate).Execute()

Retrieves report for aged payables by contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactId := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact
	date := time.Now() // string | The date of the Aged Payables By Contact report (optional)
	fromDate := time.Now() // string | filter by the from date of the report e.g. 2021-02-01 (optional)
	toDate := time.Now() // string | filter by the to date of the report e.g. 2021-02-28 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportAgedPayablesByContact(context.Background()).XeroTenantId(xeroTenantId).ContactId(contactId).Date(date).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportAgedPayablesByContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportAgedPayablesByContact`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportAgedPayablesByContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportAgedPayablesByContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **contactId** | **string** | Unique identifier for a Contact | 
 **date** | **string** | The date of the Aged Payables By Contact report | 
 **fromDate** | **string** | filter by the from date of the report e.g. 2021-02-01 | 
 **toDate** | **string** | filter by the to date of the report e.g. 2021-02-28 | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportAgedReceivablesByContact

> ReportWithRows GetReportAgedReceivablesByContact(ctx).XeroTenantId(xeroTenantId).ContactId(contactId).Date(date).FromDate(fromDate).ToDate(toDate).Execute()

Retrieves report for aged receivables by contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactId := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact
	date := time.Now() // string | The date of the Aged Receivables By Contact report (optional)
	fromDate := time.Now() // string | filter by the from date of the report e.g. 2021-02-01 (optional)
	toDate := time.Now() // string | filter by the to date of the report e.g. 2021-02-28 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportAgedReceivablesByContact(context.Background()).XeroTenantId(xeroTenantId).ContactId(contactId).Date(date).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportAgedReceivablesByContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportAgedReceivablesByContact`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportAgedReceivablesByContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportAgedReceivablesByContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **contactId** | **string** | Unique identifier for a Contact | 
 **date** | **string** | The date of the Aged Receivables By Contact report | 
 **fromDate** | **string** | filter by the from date of the report e.g. 2021-02-01 | 
 **toDate** | **string** | filter by the to date of the report e.g. 2021-02-28 | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportBalanceSheet

> ReportWithRows GetReportBalanceSheet(ctx).XeroTenantId(xeroTenantId).Date(date).Periods(periods).Timeframe(timeframe).TrackingOptionID1(trackingOptionID1).TrackingOptionID2(trackingOptionID2).StandardLayout(standardLayout).PaymentsOnly(paymentsOnly).Execute()

Retrieves report for balancesheet

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	date := time.Now() // string | The date of the Balance Sheet report (optional)
	periods := int32(3) // int32 | The number of periods for the Balance Sheet report (optional)
	timeframe := "MONTH" // string | The period size to compare to (MONTH, QUARTER, YEAR) (optional)
	trackingOptionID1 := "00000000-0000-0000-0000-000000000000" // string | The tracking option 1 for the Balance Sheet report (optional)
	trackingOptionID2 := "00000000-0000-0000-0000-000000000000" // string | The tracking option 2 for the Balance Sheet report (optional)
	standardLayout := true // bool | The standard layout boolean for the Balance Sheet report (optional)
	paymentsOnly := false // bool | return a cash basis for the Balance Sheet report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportBalanceSheet(context.Background()).XeroTenantId(xeroTenantId).Date(date).Periods(periods).Timeframe(timeframe).TrackingOptionID1(trackingOptionID1).TrackingOptionID2(trackingOptionID2).StandardLayout(standardLayout).PaymentsOnly(paymentsOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportBalanceSheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportBalanceSheet`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportBalanceSheet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportBalanceSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **date** | **string** | The date of the Balance Sheet report | 
 **periods** | **int32** | The number of periods for the Balance Sheet report | 
 **timeframe** | **string** | The period size to compare to (MONTH, QUARTER, YEAR) | 
 **trackingOptionID1** | **string** | The tracking option 1 for the Balance Sheet report | 
 **trackingOptionID2** | **string** | The tracking option 2 for the Balance Sheet report | 
 **standardLayout** | **bool** | The standard layout boolean for the Balance Sheet report | 
 **paymentsOnly** | **bool** | return a cash basis for the Balance Sheet report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportBankSummary

> ReportWithRows GetReportBankSummary(ctx).XeroTenantId(xeroTenantId).FromDate(fromDate).ToDate(toDate).Execute()

Retrieves report for bank summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	fromDate := time.Now() // string | filter by the from date of the report e.g. 2021-02-01 (optional)
	toDate := time.Now() // string | filter by the to date of the report e.g. 2021-02-28 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportBankSummary(context.Background()).XeroTenantId(xeroTenantId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportBankSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportBankSummary`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportBankSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportBankSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **fromDate** | **string** | filter by the from date of the report e.g. 2021-02-01 | 
 **toDate** | **string** | filter by the to date of the report e.g. 2021-02-28 | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportBudgetSummary

> ReportWithRows GetReportBudgetSummary(ctx).XeroTenantId(xeroTenantId).Date(date).Periods(periods).Timeframe(timeframe).Execute()

Retrieves report for budget summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	date := time.Now() // string | The date for the Bank Summary report e.g. 2018-03-31 (optional)
	periods := int32(2) // int32 | The number of periods to compare (integer between 1 and 12) (optional)
	timeframe := int32(3) // int32 | The period size to compare to (1=month, 3=quarter, 12=year) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportBudgetSummary(context.Background()).XeroTenantId(xeroTenantId).Date(date).Periods(periods).Timeframe(timeframe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportBudgetSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportBudgetSummary`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportBudgetSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportBudgetSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **date** | **string** | The date for the Bank Summary report e.g. 2018-03-31 | 
 **periods** | **int32** | The number of periods to compare (integer between 1 and 12) | 
 **timeframe** | **int32** | The period size to compare to (1&#x3D;month, 3&#x3D;quarter, 12&#x3D;year) | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportExecutiveSummary

> ReportWithRows GetReportExecutiveSummary(ctx).XeroTenantId(xeroTenantId).Date(date).Execute()

Retrieves report for executive summary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	date := time.Now() // string | The date for the Bank Summary report e.g. 2018-03-31 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportExecutiveSummary(context.Background()).XeroTenantId(xeroTenantId).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportExecutiveSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportExecutiveSummary`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportExecutiveSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportExecutiveSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **date** | **string** | The date for the Bank Summary report e.g. 2018-03-31 | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportFromId

> ReportWithRows GetReportFromId(ctx, reportID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific report using a unique ReportID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	reportID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Report

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportFromId(context.Background(), reportID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportFromId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportFromId`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportFromId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportID** | **string** | Unique identifier for a Report | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportFromIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportProfitAndLoss

> ReportWithRows GetReportProfitAndLoss(ctx).XeroTenantId(xeroTenantId).FromDate(fromDate).ToDate(toDate).Periods(periods).Timeframe(timeframe).TrackingCategoryID(trackingCategoryID).TrackingCategoryID2(trackingCategoryID2).TrackingOptionID(trackingOptionID).TrackingOptionID2(trackingOptionID2).StandardLayout(standardLayout).PaymentsOnly(paymentsOnly).Execute()

Retrieves report for profit and loss

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	fromDate := time.Now() // string | filter by the from date of the report e.g. 2021-02-01 (optional)
	toDate := time.Now() // string | filter by the to date of the report e.g. 2021-02-28 (optional)
	periods := int32(3) // int32 | The number of periods to compare (integer between 1 and 12) (optional)
	timeframe := "MONTH" // string | The period size to compare to (MONTH, QUARTER, YEAR) (optional)
	trackingCategoryID := "00000000-0000-0000-0000-000000000000" // string | The trackingCategory 1 for the ProfitAndLoss report (optional)
	trackingCategoryID2 := "00000000-0000-0000-0000-000000000000" // string | The trackingCategory 2 for the ProfitAndLoss report (optional)
	trackingOptionID := "00000000-0000-0000-0000-000000000000" // string | The tracking option 1 for the ProfitAndLoss report (optional)
	trackingOptionID2 := "00000000-0000-0000-0000-000000000000" // string | The tracking option 2 for the ProfitAndLoss report (optional)
	standardLayout := true // bool | Return the standard layout for the ProfitAndLoss report (optional)
	paymentsOnly := false // bool | Return cash only basis for the ProfitAndLoss report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportProfitAndLoss(context.Background()).XeroTenantId(xeroTenantId).FromDate(fromDate).ToDate(toDate).Periods(periods).Timeframe(timeframe).TrackingCategoryID(trackingCategoryID).TrackingCategoryID2(trackingCategoryID2).TrackingOptionID(trackingOptionID).TrackingOptionID2(trackingOptionID2).StandardLayout(standardLayout).PaymentsOnly(paymentsOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportProfitAndLoss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportProfitAndLoss`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportProfitAndLoss`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportProfitAndLossRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **fromDate** | **string** | filter by the from date of the report e.g. 2021-02-01 | 
 **toDate** | **string** | filter by the to date of the report e.g. 2021-02-28 | 
 **periods** | **int32** | The number of periods to compare (integer between 1 and 12) | 
 **timeframe** | **string** | The period size to compare to (MONTH, QUARTER, YEAR) | 
 **trackingCategoryID** | **string** | The trackingCategory 1 for the ProfitAndLoss report | 
 **trackingCategoryID2** | **string** | The trackingCategory 2 for the ProfitAndLoss report | 
 **trackingOptionID** | **string** | The tracking option 1 for the ProfitAndLoss report | 
 **trackingOptionID2** | **string** | The tracking option 2 for the ProfitAndLoss report | 
 **standardLayout** | **bool** | Return the standard layout for the ProfitAndLoss report | 
 **paymentsOnly** | **bool** | Return cash only basis for the ProfitAndLoss report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportTenNinetyNine

> Reports GetReportTenNinetyNine(ctx).XeroTenantId(xeroTenantId).ReportYear(reportYear).Execute()

Retrieve reports for 1099

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	reportYear := "2019" // string | The year of the 1099 report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportTenNinetyNine(context.Background()).XeroTenantId(xeroTenantId).ReportYear(reportYear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportTenNinetyNine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportTenNinetyNine`: Reports
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportTenNinetyNine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportTenNinetyNineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **reportYear** | **string** | The year of the 1099 report | 

### Return type

[**Reports**](Reports.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportTrialBalance

> ReportWithRows GetReportTrialBalance(ctx).XeroTenantId(xeroTenantId).Date(date).PaymentsOnly(paymentsOnly).Execute()

Retrieves report for trial balance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	date := time.Now() // string | The date for the Trial Balance report e.g. 2018-03-31 (optional)
	paymentsOnly := true // bool | Return cash only basis for the Trial Balance report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportTrialBalance(context.Background()).XeroTenantId(xeroTenantId).Date(date).PaymentsOnly(paymentsOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportTrialBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportTrialBalance`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportTrialBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportTrialBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **date** | **string** | The date for the Trial Balance report e.g. 2018-03-31 | 
 **paymentsOnly** | **bool** | Return cash only basis for the Trial Balance report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsList

> ReportWithRows GetReportsList(ctx).XeroTenantId(xeroTenantId).Execute()

Retrieves a list of the organistaions unique reports that require a uuid to fetch

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetReportsList(context.Background()).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetReportsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsList`: ReportWithRows
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetReportsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxRateByTaxType

> TaxRates GetTaxRateByTaxType(ctx, taxType).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific tax rate according to given TaxType code

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	taxType := "INPUT2" // string | A valid TaxType code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetTaxRateByTaxType(context.Background(), taxType).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetTaxRateByTaxType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxRateByTaxType`: TaxRates
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetTaxRateByTaxType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxType** | **string** | A valid TaxType code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxRateByTaxTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**TaxRates**](TaxRates.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxRates

> TaxRates GetTaxRates(ctx).XeroTenantId(xeroTenantId).Where(where).Order(order).Execute()

Retrieves tax rates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	where := "Status=="ACTIVE"" // string | Filter by an any element (optional)
	order := "Name ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetTaxRates(context.Background()).XeroTenantId(xeroTenantId).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetTaxRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxRates`: TaxRates
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetTaxRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**TaxRates**](TaxRates.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackingCategories

> TrackingCategories GetTrackingCategories(ctx).XeroTenantId(xeroTenantId).Where(where).Order(order).IncludeArchived(includeArchived).Execute()

Retrieves tracking categories and options

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	where := "Status=="ACTIVE"" // string | Filter by an any element (optional)
	order := "Name ASC" // string | Order by an any element (optional)
	includeArchived := true // bool | e.g. includeArchived=true - Categories and options with a status of ARCHIVED will be included in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetTrackingCategories(context.Background()).XeroTenantId(xeroTenantId).Where(where).Order(order).IncludeArchived(includeArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetTrackingCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrackingCategories`: TrackingCategories
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetTrackingCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackingCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 
 **includeArchived** | **bool** | e.g. includeArchived&#x3D;true - Categories and options with a status of ARCHIVED will be included in the response | 

### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackingCategory

> TrackingCategories GetTrackingCategory(ctx, trackingCategoryID).XeroTenantId(xeroTenantId).Execute()

Retrieves specific tracking categories and options using a unique tracking category Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	trackingCategoryID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a TrackingCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetTrackingCategory(context.Background(), trackingCategoryID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetTrackingCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrackingCategory`: TrackingCategories
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetTrackingCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingCategoryID** | **string** | Unique identifier for a TrackingCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackingCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> Users GetUser(ctx, userID).XeroTenantId(xeroTenantId).Execute()

Retrieves a specific user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	userID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a User

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetUser(context.Background(), userID).XeroTenantId(xeroTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: Users
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | Unique identifier for a User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


### Return type

[**Users**](Users.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> Users GetUsers(ctx).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()

Retrieves users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	ifModifiedSince := time.Now() // time.Time | Only records created or modified since this timestamp will be returned (optional)
	where := "IsSubscriber==true" // string | Filter by an any element (optional)
	order := "LastName ASC" // string | Order by an any element (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.GetUsers(context.Background()).XeroTenantId(xeroTenantId).IfModifiedSince(ifModifiedSince).Where(where).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: Users
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **ifModifiedSince** | **time.Time** | Only records created or modified since this timestamp will be returned | 
 **where** | **string** | Filter by an any element | 
 **order** | **string** | Order by an any element | 

### Return type

[**Users**](Users.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetup

> ImportSummaryObject PostSetup(ctx).XeroTenantId(xeroTenantId).Setup(setup).IdempotencyKey(idempotencyKey).Execute()

Sets the chart of accounts, the conversion date and conversion balances

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	setup := *openapiclient.NewSetup() // Setup | Object including an accounts array, a conversion balances array and a conversion date object in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.PostSetup(context.Background()).XeroTenantId(xeroTenantId).Setup(setup).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.PostSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSetup`: ImportSummaryObject
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.PostSetup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **setup** | [**Setup**](Setup.md) | Object including an accounts array, a conversion balances array and a conversion date object in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**ImportSummaryObject**](ImportSummaryObject.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> Accounts UpdateAccount(ctx, accountID).XeroTenantId(xeroTenantId).Accounts(accounts).IdempotencyKey(idempotencyKey).Execute()

Updates a chart of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	accountID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Account object
	accounts := *openapiclient.NewAccounts() // Accounts | Request of type Accounts array with one Account
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateAccount(context.Background(), accountID).XeroTenantId(xeroTenantId).Accounts(accounts).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: Accounts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Unique identifier for Account object | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **accounts** | [**Accounts**](Accounts.md) | Request of type Accounts array with one Account | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountAttachmentByFileName

> Attachments UpdateAccountAttachmentByFileName(ctx, accountID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates attachment on a specific account by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	accountID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for Account object
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateAccountAttachmentByFileName(context.Background(), accountID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateAccountAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateAccountAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | Unique identifier for Account object | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankTransaction

> BankTransactions UpdateBankTransaction(ctx, bankTransactionID).XeroTenantId(xeroTenantId).BankTransactions(bankTransactions).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates a single spent or received money transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction
	bankTransactions := *openapiclient.NewBankTransactions() // BankTransactions | 
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateBankTransaction(context.Background(), bankTransactionID).XeroTenantId(xeroTenantId).BankTransactions(bankTransactions).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateBankTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBankTransaction`: BankTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateBankTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **bankTransactions** | [**BankTransactions**](BankTransactions.md) |  | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankTransactionAttachmentByFileName

> Attachments UpdateBankTransactionAttachmentByFileName(ctx, bankTransactionID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates a specific attachment from a specific bank transaction by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactionID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transaction
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateBankTransactionAttachmentByFileName(context.Background(), bankTransactionID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateBankTransactionAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBankTransactionAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateBankTransactionAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransactionID** | **string** | Xero generated unique identifier for a bank transaction | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankTransactionAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankTransferAttachmentByFileName

> Attachments UpdateBankTransferAttachmentByFileName(ctx, bankTransferID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransferID := "00000000-0000-0000-0000-000000000000" // string | Xero generated unique identifier for a bank transfer
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateBankTransferAttachmentByFileName(context.Background(), bankTransferID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateBankTransferAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBankTransferAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateBankTransferAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankTransferID** | **string** | Xero generated unique identifier for a bank transfer | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankTransferAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> Contacts UpdateContact(ctx, contactID).XeroTenantId(xeroTenantId).Contacts(contacts).IdempotencyKey(idempotencyKey).Execute()

Updates a specific contact in a Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact
	contacts := *openapiclient.NewContacts() // Contacts | an array of Contacts containing single Contact object with properties to update
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateContact(context.Background(), contactID).XeroTenantId(xeroTenantId).Contacts(contacts).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: Contacts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **contacts** | [**Contacts**](Contacts.md) | an array of Contacts containing single Contact object with properties to update | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactAttachmentByFileName

> Attachments UpdateContactAttachmentByFileName(ctx, contactID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateContactAttachmentByFileName(context.Background(), contactID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateContactAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContactAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateContactAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Unique identifier for a Contact | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactGroup

> ContactGroups UpdateContactGroup(ctx, contactGroupID).XeroTenantId(xeroTenantId).ContactGroups(contactGroups).IdempotencyKey(idempotencyKey).Execute()

Updates a specific contact group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contactGroupID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Contact Group
	contactGroups := *openapiclient.NewContactGroups() // ContactGroups | an array of Contact groups with Name of specific group to update
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateContactGroup(context.Background(), contactGroupID).XeroTenantId(xeroTenantId).ContactGroups(contactGroups).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateContactGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContactGroup`: ContactGroups
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateContactGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactGroupID** | **string** | Unique identifier for a Contact Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **contactGroups** | [**ContactGroups**](ContactGroups.md) | an array of Contact groups with Name of specific group to update | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**ContactGroups**](ContactGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreditNote

> CreditNotes UpdateCreditNote(ctx, creditNoteID).XeroTenantId(xeroTenantId).CreditNotes(creditNotes).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates a specific credit note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	creditNotes := *openapiclient.NewCreditNotes() // CreditNotes | an array of Credit Notes containing credit note details to update
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateCreditNote(context.Background(), creditNoteID).XeroTenantId(xeroTenantId).CreditNotes(creditNotes).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateCreditNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCreditNote`: CreditNotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **creditNotes** | [**CreditNotes**](CreditNotes.md) | an array of Credit Notes containing credit note details to update | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreditNoteAttachmentByFileName

> Attachments UpdateCreditNoteAttachmentByFileName(ctx, creditNoteID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates attachments on a specific credit note by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Credit Note
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateCreditNoteAttachmentByFileName(context.Background(), creditNoteID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateCreditNoteAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCreditNoteAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateCreditNoteAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteID** | **string** | Unique identifier for a Credit Note | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreditNoteAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExpenseClaim

> ExpenseClaims UpdateExpenseClaim(ctx, expenseClaimID).XeroTenantId(xeroTenantId).ExpenseClaims(expenseClaims).IdempotencyKey(idempotencyKey).Execute()

Updates a specific expense claims

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	expenseClaimID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ExpenseClaim
	expenseClaims := *openapiclient.NewExpenseClaims() // ExpenseClaims | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateExpenseClaim(context.Background(), expenseClaimID).XeroTenantId(xeroTenantId).ExpenseClaims(expenseClaims).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateExpenseClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExpenseClaim`: ExpenseClaims
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateExpenseClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseClaimID** | **string** | Unique identifier for a ExpenseClaim | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExpenseClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **expenseClaims** | [**ExpenseClaims**](ExpenseClaims.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**ExpenseClaims**](ExpenseClaims.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> Invoices UpdateInvoice(ctx, invoiceID).XeroTenantId(xeroTenantId).Invoices(invoices).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates a specific sales invoices or purchase bills

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice
	invoices := *openapiclient.NewInvoices() // Invoices | 
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateInvoice(context.Background(), invoiceID).XeroTenantId(xeroTenantId).Invoices(invoices).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoice`: Invoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **invoices** | [**Invoices**](Invoices.md) |  | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoiceAttachmentByFileName

> Attachments UpdateInvoiceAttachmentByFileName(ctx, invoiceID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates an attachment from a specific invoices or purchase bill by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Invoice
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateInvoiceAttachmentByFileName(context.Background(), invoiceID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateInvoiceAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoiceAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateInvoiceAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceID** | **string** | Unique identifier for an Invoice | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItem

> Items UpdateItem(ctx, itemID).XeroTenantId(xeroTenantId).Items(items).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates a specific item

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	itemID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Item
	items := *openapiclient.NewItems() // Items | 
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateItem(context.Background(), itemID).XeroTenantId(xeroTenantId).Items(items).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItem`: Items
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemID** | **string** | Unique identifier for an Item | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **items** | [**Items**](Items.md) |  | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinkedTransaction

> LinkedTransactions UpdateLinkedTransaction(ctx, linkedTransactionID).XeroTenantId(xeroTenantId).LinkedTransactions(linkedTransactions).IdempotencyKey(idempotencyKey).Execute()

Updates a specific linked transactions (billable expenses)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	linkedTransactionID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a LinkedTransaction
	linkedTransactions := *openapiclient.NewLinkedTransactions() // LinkedTransactions | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateLinkedTransaction(context.Background(), linkedTransactionID).XeroTenantId(xeroTenantId).LinkedTransactions(linkedTransactions).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateLinkedTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLinkedTransaction`: LinkedTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateLinkedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkedTransactionID** | **string** | Unique identifier for a LinkedTransaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinkedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **linkedTransactions** | [**LinkedTransactions**](LinkedTransactions.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**LinkedTransactions**](LinkedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManualJournal

> ManualJournals UpdateManualJournal(ctx, manualJournalID).XeroTenantId(xeroTenantId).ManualJournals(manualJournals).IdempotencyKey(idempotencyKey).Execute()

Updates a specific manual journal

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal
	manualJournals := *openapiclient.NewManualJournals() // ManualJournals | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateManualJournal(context.Background(), manualJournalID).XeroTenantId(xeroTenantId).ManualJournals(manualJournals).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateManualJournal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateManualJournal`: ManualJournals
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateManualJournal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManualJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **manualJournals** | [**ManualJournals**](ManualJournals.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManualJournalAttachmentByFileName

> Attachments UpdateManualJournalAttachmentByFileName(ctx, manualJournalID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates a specific attachment from a specific manual journal by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournalID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a ManualJournal
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateManualJournalAttachmentByFileName(context.Background(), manualJournalID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateManualJournalAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateManualJournalAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateManualJournalAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manualJournalID** | **string** | Unique identifier for a ManualJournal | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManualJournalAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateBankTransactions

> BankTransactions UpdateOrCreateBankTransactions(ctx).XeroTenantId(xeroTenantId).BankTransactions(bankTransactions).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates or creates one or more spent or received money transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	bankTransactions := *openapiclient.NewBankTransactions() // BankTransactions | 
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateBankTransactions(context.Background()).XeroTenantId(xeroTenantId).BankTransactions(bankTransactions).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateBankTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateBankTransactions`: BankTransactions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateBankTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateBankTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **bankTransactions** | [**BankTransactions**](BankTransactions.md) |  | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateContacts

> Contacts UpdateOrCreateContacts(ctx).XeroTenantId(xeroTenantId).Contacts(contacts).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Updates or creates one or more contacts in a Xero organisation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	contacts := *openapiclient.NewContacts() // Contacts | 
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateContacts(context.Background()).XeroTenantId(xeroTenantId).Contacts(contacts).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateContacts`: Contacts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **contacts** | [**Contacts**](Contacts.md) |  | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateCreditNotes

> CreditNotes UpdateOrCreateCreditNotes(ctx).XeroTenantId(xeroTenantId).CreditNotes(creditNotes).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates or creates one or more credit notes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	creditNotes := *openapiclient.NewCreditNotes() // CreditNotes | an array of Credit Notes with a single CreditNote object.
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateCreditNotes(context.Background()).XeroTenantId(xeroTenantId).CreditNotes(creditNotes).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateCreditNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateCreditNotes`: CreditNotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateCreditNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateCreditNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **creditNotes** | [**CreditNotes**](CreditNotes.md) | an array of Credit Notes with a single CreditNote object. | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateEmployees

> Employees UpdateOrCreateEmployees(ctx).XeroTenantId(xeroTenantId).Employees(employees).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates a single new employees used in Xero payrun

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	employees := *openapiclient.NewEmployees() // Employees | Employees with array of Employee object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateEmployees(context.Background()).XeroTenantId(xeroTenantId).Employees(employees).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateEmployees``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateEmployees`: Employees
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateEmployees`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateEmployeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **employees** | [**Employees**](Employees.md) | Employees with array of Employee object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Employees**](Employees.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateInvoices

> Invoices UpdateOrCreateInvoices(ctx).XeroTenantId(xeroTenantId).Invoices(invoices).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates or creates one or more sales invoices or purchase bills

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	invoices := *openapiclient.NewInvoices() // Invoices | 
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateInvoices(context.Background()).XeroTenantId(xeroTenantId).Invoices(invoices).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateInvoices`: Invoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **invoices** | [**Invoices**](Invoices.md) |  | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateItems

> Items UpdateOrCreateItems(ctx).XeroTenantId(xeroTenantId).Items(items).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates or creates one or more items

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	items := *openapiclient.NewItems() // Items | 
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateItems(context.Background()).XeroTenantId(xeroTenantId).Items(items).SummarizeErrors(summarizeErrors).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateItems`: Items
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **items** | [**Items**](Items.md) |  | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateManualJournals

> ManualJournals UpdateOrCreateManualJournals(ctx).XeroTenantId(xeroTenantId).ManualJournals(manualJournals).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Updates or creates a single manual journal

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	manualJournals := *openapiclient.NewManualJournals() // ManualJournals | ManualJournals array with ManualJournal object in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateManualJournals(context.Background()).XeroTenantId(xeroTenantId).ManualJournals(manualJournals).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateManualJournals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateManualJournals`: ManualJournals
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateManualJournals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateManualJournalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **manualJournals** | [**ManualJournals**](ManualJournals.md) | ManualJournals array with ManualJournal object in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreatePurchaseOrders

> PurchaseOrders UpdateOrCreatePurchaseOrders(ctx).XeroTenantId(xeroTenantId).PurchaseOrders(purchaseOrders).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Updates or creates one or more purchase orders

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrders := *openapiclient.NewPurchaseOrders() // PurchaseOrders | 
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreatePurchaseOrders(context.Background()).XeroTenantId(xeroTenantId).PurchaseOrders(purchaseOrders).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreatePurchaseOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreatePurchaseOrders`: PurchaseOrders
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreatePurchaseOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreatePurchaseOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **purchaseOrders** | [**PurchaseOrders**](PurchaseOrders.md) |  | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateQuotes

> Quotes UpdateOrCreateQuotes(ctx).XeroTenantId(xeroTenantId).Quotes(quotes).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Updates or creates one or more quotes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quotes := *openapiclient.NewQuotes() // Quotes | 
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateQuotes(context.Background()).XeroTenantId(xeroTenantId).Quotes(quotes).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateQuotes`: Quotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **quotes** | [**Quotes**](Quotes.md) |  | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateRepeatingInvoices

> RepeatingInvoices UpdateOrCreateRepeatingInvoices(ctx).XeroTenantId(xeroTenantId).RepeatingInvoices(repeatingInvoices).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()

Creates or deletes one or more repeating invoice templates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoices := *openapiclient.NewRepeatingInvoices() // RepeatingInvoices | RepeatingInvoices with an array of repeating invoice objects in body of request
	summarizeErrors := true // bool | If false return 200 OK and mix of successfully created objects and any with validation errors (optional) (default to false)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateOrCreateRepeatingInvoices(context.Background()).XeroTenantId(xeroTenantId).RepeatingInvoices(repeatingInvoices).SummarizeErrors(summarizeErrors).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateOrCreateRepeatingInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrCreateRepeatingInvoices`: RepeatingInvoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateOrCreateRepeatingInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrCreateRepeatingInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **repeatingInvoices** | [**RepeatingInvoices**](RepeatingInvoices.md) | RepeatingInvoices with an array of repeating invoice objects in body of request | 
 **summarizeErrors** | **bool** | If false return 200 OK and mix of successfully created objects and any with validation errors | [default to false]
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**RepeatingInvoices**](RepeatingInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePurchaseOrder

> PurchaseOrders UpdatePurchaseOrder(ctx, purchaseOrderID).XeroTenantId(xeroTenantId).PurchaseOrders(purchaseOrders).IdempotencyKey(idempotencyKey).Execute()

Updates a specific purchase order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order
	purchaseOrders := *openapiclient.NewPurchaseOrders() // PurchaseOrders | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdatePurchaseOrder(context.Background(), purchaseOrderID).XeroTenantId(xeroTenantId).PurchaseOrders(purchaseOrders).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdatePurchaseOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePurchaseOrder`: PurchaseOrders
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdatePurchaseOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePurchaseOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **purchaseOrders** | [**PurchaseOrders**](PurchaseOrders.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePurchaseOrderAttachmentByFileName

> Attachments UpdatePurchaseOrderAttachmentByFileName(ctx, purchaseOrderID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates a specific attachment for a specific purchase order by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	purchaseOrderID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Purchase Order
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdatePurchaseOrderAttachmentByFileName(context.Background(), purchaseOrderID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdatePurchaseOrderAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePurchaseOrderAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdatePurchaseOrderAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseOrderID** | **string** | Unique identifier for an Purchase Order | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePurchaseOrderAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuote

> Quotes UpdateQuote(ctx, quoteID).XeroTenantId(xeroTenantId).Quotes(quotes).IdempotencyKey(idempotencyKey).Execute()

Updates a specific quote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote
	quotes := *openapiclient.NewQuotes() // Quotes | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateQuote(context.Background(), quoteID).XeroTenantId(xeroTenantId).Quotes(quotes).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQuote`: Quotes
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **quotes** | [**Quotes**](Quotes.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuoteAttachmentByFileName

> Attachments UpdateQuoteAttachmentByFileName(ctx, quoteID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates a specific attachment from a specific quote by filename

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	quoteID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for an Quote
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateQuoteAttachmentByFileName(context.Background(), quoteID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateQuoteAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQuoteAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateQuoteAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteID** | **string** | Unique identifier for an Quote | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuoteAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReceipt

> Receipts UpdateReceipt(ctx, receiptID).XeroTenantId(xeroTenantId).Receipts(receipts).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()

Updates a specific draft expense claim receipts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt
	receipts := *openapiclient.NewReceipts() // Receipts | 
	unitdp := int32(4) // int32 | e.g. unitdp=4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts (optional)
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateReceipt(context.Background(), receiptID).XeroTenantId(xeroTenantId).Receipts(receipts).Unitdp(unitdp).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReceipt`: Receipts
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **receipts** | [**Receipts**](Receipts.md) |  | 
 **unitdp** | **int32** | e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Receipts**](Receipts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReceiptAttachmentByFileName

> Attachments UpdateReceiptAttachmentByFileName(ctx, receiptID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates a specific attachment on a specific expense claim receipts by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	receiptID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Receipt
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateReceiptAttachmentByFileName(context.Background(), receiptID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateReceiptAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReceiptAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateReceiptAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptID** | **string** | Unique identifier for a Receipt | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReceiptAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepeatingInvoice

> RepeatingInvoices UpdateRepeatingInvoice(ctx, repeatingInvoiceID).XeroTenantId(xeroTenantId).RepeatingInvoices(repeatingInvoices).IdempotencyKey(idempotencyKey).Execute()

Deletes a specific repeating invoice template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice
	repeatingInvoices := *openapiclient.NewRepeatingInvoices() // RepeatingInvoices | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateRepeatingInvoice(context.Background(), repeatingInvoiceID).XeroTenantId(xeroTenantId).RepeatingInvoices(repeatingInvoices).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateRepeatingInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRepeatingInvoice`: RepeatingInvoices
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateRepeatingInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepeatingInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **repeatingInvoices** | [**RepeatingInvoices**](RepeatingInvoices.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**RepeatingInvoices**](RepeatingInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepeatingInvoiceAttachmentByFileName

> Attachments UpdateRepeatingInvoiceAttachmentByFileName(ctx, repeatingInvoiceID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()

Updates a specific attachment from a specific repeating invoices by file name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	repeatingInvoiceID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Repeating Invoice
	fileName := "xero-dev.jpg" // string | Name of the attachment
	body := string(BYTE_ARRAY_DATA_HERE) // string | Byte array of file in body of request
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateRepeatingInvoiceAttachmentByFileName(context.Background(), repeatingInvoiceID, fileName).XeroTenantId(xeroTenantId).Body(body).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateRepeatingInvoiceAttachmentByFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRepeatingInvoiceAttachmentByFileName`: Attachments
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateRepeatingInvoiceAttachmentByFileName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repeatingInvoiceID** | **string** | Unique identifier for a Repeating Invoice | 
**fileName** | **string** | Name of the attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepeatingInvoiceAttachmentByFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **body** | **string** | Byte array of file in body of request | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaxRate

> TaxRates UpdateTaxRate(ctx).XeroTenantId(xeroTenantId).TaxRates(taxRates).IdempotencyKey(idempotencyKey).Execute()

Updates tax rates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	taxRates := *openapiclient.NewTaxRates() // TaxRates | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateTaxRate(context.Background()).XeroTenantId(xeroTenantId).TaxRates(taxRates).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateTaxRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTaxRate`: TaxRates
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateTaxRate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaxRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 
 **taxRates** | [**TaxRates**](TaxRates.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**TaxRates**](TaxRates.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrackingCategory

> TrackingCategories UpdateTrackingCategory(ctx, trackingCategoryID).XeroTenantId(xeroTenantId).TrackingCategory(trackingCategory).IdempotencyKey(idempotencyKey).Execute()

Updates a specific tracking category

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	trackingCategoryID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a TrackingCategory
	trackingCategory := *openapiclient.NewTrackingCategory() // TrackingCategory | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateTrackingCategory(context.Background(), trackingCategoryID).XeroTenantId(xeroTenantId).TrackingCategory(trackingCategory).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateTrackingCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTrackingCategory`: TrackingCategories
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateTrackingCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingCategoryID** | **string** | Unique identifier for a TrackingCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrackingCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 

 **trackingCategory** | [**TrackingCategory**](TrackingCategory.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrackingOptions

> TrackingOptions UpdateTrackingOptions(ctx, trackingCategoryID, trackingOptionID).XeroTenantId(xeroTenantId).TrackingOption(trackingOption).IdempotencyKey(idempotencyKey).Execute()

Updates a specific option for a specific tracking category

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mkcr-innovations/xero-go-client/accounting/client"
)

func main() {
	xeroTenantId := "YOUR_XERO_TENANT_ID" // string | Xero identifier for Tenant
	trackingCategoryID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a TrackingCategory
	trackingOptionID := "00000000-0000-0000-0000-000000000000" // string | Unique identifier for a Tracking Option
	trackingOption := *openapiclient.NewTrackingOption() // TrackingOption | 
	idempotencyKey := "KEY_VALUE" // string | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingAPI.UpdateTrackingOptions(context.Background(), trackingCategoryID, trackingOptionID).XeroTenantId(xeroTenantId).TrackingOption(trackingOption).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingAPI.UpdateTrackingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTrackingOptions`: TrackingOptions
	fmt.Fprintf(os.Stdout, "Response from `AccountingAPI.UpdateTrackingOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingCategoryID** | **string** | Unique identifier for a TrackingCategory | 
**trackingOptionID** | **string** | Unique identifier for a Tracking Option | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrackingOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xeroTenantId** | **string** | Xero identifier for Tenant | 


 **trackingOption** | [**TrackingOption**](TrackingOption.md) |  | 
 **idempotencyKey** | **string** | This allows you to safely retry requests without the risk of duplicate processing. 128 character max. | 

### Return type

[**TrackingOptions**](TrackingOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

