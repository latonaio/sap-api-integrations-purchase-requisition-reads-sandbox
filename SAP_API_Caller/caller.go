package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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
		
func (c *SAPAPICaller) AsyncGetPurchaseRequisition(PurchaseRequisition, PurchaseRequisitionItem, PurchasingDocument, PurchasingDocumentItem string) {
	wg := &sync.WaitGroup{}

	wg.Add(5)
	go func() {
		c.Header(PurchaseRequisition)
		wg.Done()
	}()
	go func() {
		c.Item(PurchaseRequisition, PurchaseRequisitionItem)
		wg.Done()
	}()
	go func() {
		c.DeliveryAddress(PurchaseRequisition, PurchaseRequisitionItem)
		wg.Done()
	}()
	go func() {
		c.Account(PurchaseRequisition, PurchaseRequisitionItem)
		wg.Done()
	}()
	go func() {
		c.PurchasingDocument(PurchasingDocument, PurchasingDocumentItem)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Header(PurchaseRequisition string) {
	res, err := c.callPurchaseRequisitionSrvAPIRequirementHeader("A_PurchaseRequisitionHeader", PurchaseRequisition)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Item(PurchaseRequisition, PurchaseRequisitionItem string) {
	res, err := c.callPurchaseRequisitionSrvAPIRequirementItem("A_PurchaseRequisitionItem", PurchaseRequisition, PurchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) DeliveryAddress(PurchaseRequisition, PurchaseRequisitionItem string) {
	res, err := c.callPurchaseRequisitionSrvAPIRequirementDeliveryAddress("A_PurReqAddDelivery", PurchaseRequisition, PurchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Account(PurchaseRequisition, PurchaseRequisitionItem string) {
	res, err := c.callPurchaseRequisitionSrvAPIRequirementAccount("A_PurReqnAcctAssgmt", PurchaseRequisition, PurchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) PurchasingDocument(PurchasingDocument, PurchasingDocumentItem string) {
	res, err := c.callPurchaseRequisitionSrvAPIRequirementPurchasingDocument("A_PurchaseRequisitionItem", PurchasingDocument, PurchasingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementHeader(api, PurchaseRequisition string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchaseRequisition")
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s'", PurchaseRequisition))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementItem(api, PurchaseRequisition, PurchaseRequisitionItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchaseRequisition, PurchaseRequisitionItem")
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", PurchaseRequisition, PurchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementDeliveryAddress(api, PurchaseRequisition, PurchaseRequisitionItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchaseRequisition, PurchaseRequisitionItem")
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", PurchaseRequisition, PurchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementAccount(api, PurchaseRequisition, PurchaseRequisitionItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchaseRequisition, PurchaseRequisitionItem")
	params.Add("$filter", fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", PurchaseRequisition, PurchaseRequisitionItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callPurchaseRequisitionSrvAPIRequirementPurchasingDocument(api, PurchasingDocument, PurchasingDocumentItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEREQ_PROCESS_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PurchasingDocument, PurchasingDocumentItem")
	params.Add("$filter", fmt.Sprintf("PurchasingDocument eq '%s' and PurchasingDocumentItem eq '%s'", PurchasingDocument, PurchasingDocumentItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}
