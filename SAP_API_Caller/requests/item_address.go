package requests

type ItemAddress struct {
	PurchaseContract       string  `json:"PurchaseContract"`
	AddressID              string  `json:"AddressID"`
	PurchaseContractItem   string  `json:"PurchaseContractItem"`
	CityName               *string `json:"CityName"`
	PostalCode             *string `json:"PostalCode"`
	StreetName             *string `json:"StreetName"`
	Country                *string `json:"Country"`
	CorrespondenceLanguage *string `json:"CorrespondenceLanguage"`
	Region                 *string `json:"Region"`
	PhoneNumber            *string `json:"PhoneNumber"`
	FaxNumber              *string `json:"FaxNumber"`
}
