package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-contract-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
	PurchaseContract:               data.PurchaseContract,
	PurchaseContractType:           data.PurchaseContractType,
	CompanyCode:                    data.CompanyCode,
	PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
	CreationDate:                   data.CreationDate,
	Supplier:                       data.Supplier,
	PurchasingOrganization:         data.PurchasingOrganization,
	PurchasingGroup:                data.PurchasingGroup,
	PaymentTerms:                   data.PaymentTerms,
	NetPaymentDays:                 data.NetPaymentDays,
	DocumentCurrency:               data.DocumentCurrency,
	ExchangeRate:                   data.ExchangeRate,
	ValidityStartDate:              data.ValidityStartDate,
	ValidityEndDate:                data.ValidityEndDate,
	SupplierRespSalesPersonName:    data.SupplierRespSalesPersonName,
	SupplierPhoneNumber:            data.SupplierPhoneNumber,
	IncotermsClassification:        data.IncotermsClassification,
	PurchaseContractTargetAmount:   data.PurchaseContractTargetAmount,
	InvoicingParty:                 data.InvoicingParty,
	ReleaseCode:                    data.ReleaseCode,
	LastChangeDateTime:             data.LastChangeDateTime,
	PurchasingProcessingStatus:     data.PurchasingProcessingStatus,
	PurchasingProcessingStatusName: data.PurchasingProcessingStatusName,
	PurgContractIsInPreparation:    data.PurgContractIsInPreparation,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
	PurchaseContract:               data.PurchaseContract,
	PurchaseContractItem:           data.PurchaseContractItem,
	PurchasingContractDeletionCode: data.PurchasingContractDeletionCode,
	PurchaseContractItemText:       data.PurchaseContractItemText,
	CompanyCode:                    data.CompanyCode,
	Plant:                          data.Plant,
	StorageLocation:                data.StorageLocation,
	MaterialGroup:                  data.MaterialGroup,
	SupplierMaterialNumber:         data.SupplierMaterialNumber,
	OrderQuantityUnit:              data.OrderQuantityUnit,
	TargetQuantity:                 data.TargetQuantity,
	PurgDocReleaseOrderQuantity:    data.PurgDocReleaseOrderQuantity,
	OrderPriceUnit:                 data.OrderPriceUnit,
	ContractNetPriceAmount:         data.ContractNetPriceAmount,
	DocumentCurrency:               data.DocumentCurrency,
	NetPriceQuantity:               data.NetPriceQuantity,
	TaxCode:                        data.TaxCode,
	TaxCountry:                     data.TaxCountry,
	StockType:                      data.StockType,
	IsInfoRecordUpdated:            data.IsInfoRecordUpdated,
	PriceIsToBePrinted:             data.PriceIsToBePrinted,
	PurgDocEstimatedPrice:          data.PurgDocEstimatedPrice,
	PlannedDeliveryDurationInDays:  data.PlannedDeliveryDurationInDays,
	OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
	UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
	UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
	PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
	AccountAssignmentCategory:      data.AccountAssignmentCategory,
	GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
	GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
	InvoiceIsExpected:              data.InvoiceIsExpected,
	InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
	ManualDeliveryAddressID:        data.ManualDeliveryAddressID,
	VolumeUnit:                     data.VolumeUnit,
	Subcontractor:                  data.Subcontractor,
	EvaldRcptSettlmtIsAllowed:      data.EvaldRcptSettlmtIsAllowed,
	Material:                       data.Material,
	ProductType:                    data.ProductType,
	MaterialType:                   data.MaterialType,
	}

	return item, nil
}

func ConvertToItemAddress(raw []byte, l *logger.Logger) (*ItemAddress, error) {
	pm := &responses.ItemAddress{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemAddress. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	itemAddress := &ItemAddress{
	PurchaseContract:           data.PurchaseContract,
	AddressID:                  data.AddressID,
	PurchaseContractItem:       data.PurchaseContractItem,
	CityName:                   data.CityName,
	PostalCode:                 data.PostalCode,
	StreetName:                 data.StreetName,
	Country:                    data.Country,
	CorrespondenceLanguage:     data.CorrespondenceLanguage,
	Region:                     data.Region,
	PhoneNumber:                data.PhoneNumber,
	FaxNumber:                  data.FaxNumber,
	}

	return itemAddress, nil
}

func ConvertToItemCondition(raw []byte, l *logger.Logger) (*ItemCondition, error) {
	pm := &responses.ItemCondition{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemCondition. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	itemCondition := &ItemCondition{
	PurchaseContract:             data.PurchaseContract,
	PurchaseContractItem:         data.PurchaseContractItem,
	ConditionValidityEndDate:     data.ConditionValidityEndDate,
	ConditionType:                data.ConditionType,
	ConditionRecord:              data.ConditionRecord,
	ConditionSequentialNumber:    data.ConditionSequentialNumber,
	ConditionValidityStartDate:   data.ConditionValidityStartDate,
	PricingScaleType:             data.PricingScaleType,
	PricingScaleBasis:            data.PricingScaleBasis,
	ConditionScaleQuantity:       data.ConditionScaleQuantity,
	ConditionScaleQuantityUnit:   data.ConditionScaleQuantityUnit,
	ConditionScaleAmount:         data.ConditionScaleAmount,
	ConditionScaleAmountCurrency: data.ConditionScaleAmountCurrency,
	ConditionCalculationType:     data.ConditionCalculationType,
	ConditionRateValue:           data.ConditionRateValue,
	ConditionRateValueUnit:       data.ConditionRateValueUnit,
	ConditionQuantity:            data.ConditionQuantity,
	ConditionQuantityUnit:        data.ConditionQuantityUnit,
	BaseUnit:                     data.BaseUnit,
	ConditionIsDeleted:           data.ConditionIsDeleted,
	PaymentTerms:                 data.PaymentTerms,
	ConditionReleaseStatus:       data.ConditionReleaseStatus,
	}

	return itemCondition, nil
}
