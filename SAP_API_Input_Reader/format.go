package sap_api_input_reader

import (
	"sap-api-integrations-purchase-contract-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.PurchaseContract
	return &requests.Header{
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
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataPurchaseContract := sdc.PurchaseContract
	data := sdc.PurchaseContract.PurchaseContractItem
	return &requests.Item{
		PurchaseContract:               dataPurchaseContract.PurchaseContract,
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
}

func (sdc *SDC) ConvertToItemAddress() *requests.ItemAddress {
	dataPurchaseContract := sdc.PurchaseContract
	dataPurchaseContractItem := sdc.PurchaseContract.PurchaseContractItem
	data := sdc.PurchaseContract.PurchaseContractItem.ItemAddress
	return &requests.ItemAddress{
		PurchaseContract:       dataPurchaseContract.PurchaseContract,
		AddressID:              data.AddressID,
		PurchaseContractItem:   dataPurchaseContractItem.PurchaseContractItem,
		CityName:               data.CityName,
		PostalCode:             data.PostalCode,
		StreetName:             data.StreetName,
		Country:                data.Country,
		CorrespondenceLanguage: data.CorrespondenceLanguage,
		Region:                 data.Region,
		PhoneNumber:            data.PhoneNumber,
		FaxNumber:              data.FaxNumber,
	}
}

func (sdc *SDC) ConvertToItemCondition() *requests.ItemCondition {
	dataPurchaseContract := sdc.PurchaseContract
	dataPurchaseContractItem := sdc.PurchaseContract.PurchaseContractItem
	data := sdc.PurchaseContract.PurchaseContractItem.ItemCondition
	return &requests.ItemCondition{
		PurchaseContract:             dataPurchaseContract.PurchaseContract,
		PurchaseContractItem:         dataPurchaseContractItem.PurchaseContractItem,
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
}
