package requests

type Header struct {
	PurchaseContract               string  `json:"PurchaseContract"`
	PurchaseContractType           string  `json:"PurchaseContractType"`
	CompanyCode                    string `json:"CompanyCode"`
	PurchasingDocumentDeletionCode *string `json:"PurchasingDocumentDeletionCode"`
	CreationDate                   *string `json:"CreationDate"`
	Supplier                       string `json:"Supplier"`
	PurchasingOrganization         string `json:"PurchasingOrganization"`
	PurchasingGroup                string `json:"PurchasingGroup"`
	PaymentTerms                   *string `json:"PaymentTerms"`
	NetPaymentDays                 *string `json:"NetPaymentDays"`
	DocumentCurrency               *string `json:"DocumentCurrency"`
	ExchangeRate                   *string `json:"ExchangeRate"`
	ValidityStartDate              *string `json:"ValidityStartDate"`
	ValidityEndDate                *string `json:"ValidityEndDate"`
	SupplierRespSalesPersonName    *string `json:"SupplierRespSalesPersonName"`
	SupplierPhoneNumber            *string `json:"SupplierPhoneNumber"`
	IncotermsClassification        *string `json:"IncotermsClassification"`
	PurchaseContractTargetAmount   *string `json:"PurchaseContractTargetAmount"`
	InvoicingParty                 *string `json:"InvoicingParty"`
	ReleaseCode                    *string `json:"ReleaseCode"`
	LastChangeDateTime             *string `json:"LastChangeDateTime"`
	PurchasingProcessingStatus     *string `json:"PurchasingProcessingStatus"`
	PurchasingProcessingStatusName *string `json:"PurchasingProcessingStatusName"`
	PurgContractIsInPreparation    *bool   `json:"PurgContractIsInPreparation"`
}
