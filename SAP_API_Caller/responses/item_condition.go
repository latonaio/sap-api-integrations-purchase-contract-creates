package responses

type ItemCondition struct {
	D struct {
		PurchaseContract             string `json:"PurchaseContract"`
		PurchaseContractItem         string `json:"PurchaseContractItem"`
		ConditionValidityEndDate     string `json:"ConditionValidityEndDate"`
		ConditionType                string `json:"ConditionType"`
		ConditionRecord              string `json:"ConditionRecord"`
		ConditionSequentialNumber    string `json:"ConditionSequentialNumber"`
		ConditionValidityStartDate   string `json:"ConditionValidityStartDate"`
		PricingScaleType             string `json:"PricingScaleType"`
		PricingScaleBasis            string `json:"PricingScaleBasis"`
		ConditionScaleQuantity       string `json:"ConditionScaleQuantity"`
		ConditionScaleQuantityUnit   string `json:"ConditionScaleQuantityUnit"`
		ConditionScaleAmount         string `json:"ConditionScaleAmount"`
		ConditionScaleAmountCurrency string `json:"ConditionScaleAmountCurrency"`
		ConditionCalculationType     string `json:"ConditionCalculationType"`
		ConditionRateValue           string `json:"ConditionRateValue"`
		ConditionRateValueUnit       string `json:"ConditionRateValueUnit"`
		ConditionQuantity            string `json:"ConditionQuantity"`
		ConditionQuantityUnit        string `json:"ConditionQuantityUnit"`
		BaseUnit                     string `json:"BaseUnit"`
		ConditionIsDeleted           bool   `json:"ConditionIsDeleted"`
		PaymentTerms                 string `json:"PaymentTerms"`
		ConditionReleaseStatus       string `json:"ConditionReleaseStatus"`
	} `json:"d"`
}
