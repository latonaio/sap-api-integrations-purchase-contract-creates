package main

import (
	sap_api_caller "sap-api-integrations-purchase-contract-creates/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-purchase-contract-creates/SAP_API_Input_Reader"
	"sap-api-integrations-purchase-contract-creates/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	sap_api_time_value_converter "github.com/latonaio/sap-api-time-value-converter"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	pc := sap_api_request_client_header_setup.NewSAPRequestClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		pc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs/SDC_Purchase_Contract_Header_sample.json")
	sap_api_time_value_converter.ChangeTimeFormatToSAPFormatStruct(&inputSDC)

	accepter := getAccepter(inputSDC)
	header := inputSDC.ConvertToHeader()
	item := inputSDC.ConvertToItem()
	itemAddress := inputSDC.ConvertToItemAddress()
	itemCondition := inputSDC.ConvertToItemCondition()

	caller.AsyncPostPurchaseContract(
		header,
		item,
		itemAddress,
		itemCondition,
		accepter,
	)
}

func getAccepter(sdc sap_api_input_reader.SDC) []string {
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"Header", "Item", "ItemAddress", "ItemCondition",
		}
	}
	return accepter
}
