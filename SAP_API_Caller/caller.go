package sap_api_caller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-purchase-contract-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-purchase-contract-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostPurchaseContract(
	header *requests.Header,
	item *requests.Item,
	itemAddress *requests.ItemAddress,
	itemCondition *requests.ItemCondition,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(item)
				wg.Done()
			}()
		case "ItemAddress":
			func() {
				c.ItemAddress(itemAddress)
				wg.Done()
			}()
		case "ItemCondition":
			func() {
				c.ItemCondition(itemCondition)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(header *requests.Header) {
	outputDataHeader, err := c.callPurchaseContractSrvAPIRequirementHeader("A_PurchaseContract", header)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataHeader)
}

func (c *SAPAPICaller) callPurchaseContractSrvAPIRequirementHeader(api string, header *requests.Header) (*sap_api_output_formatter.Header, error) {
	body, err := json.Marshal(header)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PURCHASECONTRACT_PROCESS_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(item *requests.Item) {
	url := fmt.Sprintf("A_PurchaseContract('%s')/to_Item", item.PurchaseContract)
	outputDataItem, err := c.callPurchaseContractSrvAPIRequirementItem(url, item)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataItem)
}

func (c *SAPAPICaller) callPurchaseContractSrvAPIRequirementItem(api string, item *requests.Item) (*sap_api_output_formatter.Item, error) {
	body, err := json.Marshal(item)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PURCHASECONTRACT_PROCESS_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemAddress(itemAddress *requests.ItemAddress) {
	outputDataItemAddress, err := c.callPurchaseContractSrvAPIRequirementItemAddress("A_PurCtrAddress", itemAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataItemAddress)
}

func (c *SAPAPICaller) callPurchaseContractSrvAPIRequirementItemAddress(api string, itemAddress *requests.ItemAddress) (*sap_api_output_formatter.ItemAddress, error) {
	body, err := json.Marshal(itemAddress)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PURCHASECONTRACT_PROCESS_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToItemAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemCondition(itemCondition *requests.ItemCondition) {
	outputDataItemCondition, err := c.callPurchaseContractSrvAPIRequirementItemCondition("A_PurContrItemCondition", itemCondition)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataItemCondition)
}

func (c *SAPAPICaller) callPurchaseContractSrvAPIRequirementItemCondition(api string, itemCondition *requests.ItemCondition) (*sap_api_output_formatter.ItemCondition, error) {
	body, err := json.Marshal(itemCondition)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PURCHASECONTRACT_PROCESS_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToItemCondition(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
