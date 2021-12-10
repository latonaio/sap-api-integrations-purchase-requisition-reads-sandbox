package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-purchase-requisition-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchaseRequisition(purchaseRequisition, purchaseRequisitionItem, purchasingDocument, purchasingDocumentItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(purchaseRequisition)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "DeliveryAddress":
			func() {
				c.DeliveryAddress(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "Account":
			func() {
				c.Account(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		case "PurchasingDocument":
			func() {
				c.PurchasingDocument(purchasingDocument, purchasingDocumentItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(purchaseRequisition string) {
	data, err := c.callPurchaseRequisitionSrvAPIRequirementHeader("A_PurchaseRequisitionHeader", purchaseRequisition)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementHeader(api, purchaseRequisition string) (*sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, purchaseRequisition)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(purchaseRequisition, purchaseRequisitionItem string) {
	data, err := c.callPurchaseRequisitionSrvAPIRequirementItem("A_PurchaseRequisitionItem", purchaseRequisition, purchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementItem(api, purchaseRequisition, purchaseRequisitionItem string) (*sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, purchaseRequisition, purchaseRequisitionItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) DeliveryAddress(purchaseRequisition, purchaseRequisitionItem string) {
	data, err := c.callPurchaseRequisitionSrvAPIRequirementDeliveryAddress("A_PurReqAddDelivery", purchaseRequisition, purchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementDeliveryAddress(api, purchaseRequisition, purchaseRequisitionItem string) (*sap_api_output_formatter.DeliveryAddress, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithDeliveryAddress(req, purchaseRequisition, purchaseRequisitionItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToDeliveryAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Account(purchaseRequisition, purchaseRequisitionItem string) {
	data, err := c.callPurchaseRequisitionSrvAPIRequirementAccount("A_PurReqnAcctAssgmt", purchaseRequisition, purchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementAccount(api, purchaseRequisition, purchaseRequisitionItem string) (*sap_api_output_formatter.Account, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithAccount(req, purchaseRequisition, purchaseRequisitionItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchasingDocument(purchasingDocument, purchasingDocumentItem string) {
	data, err := c.callPurchaseRequisitionSrvAPIRequirementPurchasingDocument("A_PurchaseRequisitionItem", purchasingDocument, purchasingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementPurchasingDocument(api, purchasingDocument, purchasingDocumentItem string) (*sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPurchasingDocument(req, purchasingDocument, purchasingDocumentItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, purchaseRequisition string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s'", purchaseRequisition))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, purchaseRequisition, purchaseRequisitionItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithAccount(req *http.Request, purchaseRequisition, purchaseRequisitionItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithDeliveryAddress(req *http.Request, purchaseRequisition, purchaseRequisitionItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPurchasingDocument(req *http.Request, purchasingDocument, purchasingDocumentItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PurchasingDocument eq '%s' and PurchasingDocumentItem eq '%s'", purchasingDocument, purchasingDocumentItem))
	req.URL.RawQuery = params.Encode()
}
