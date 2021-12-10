package main

import (
	sap_api_caller "sap-api-integrations-purchase-requisition-reads/SAP_API_Caller"
	"sap-api-integrations-purchase-requisition-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Purchase_Requisition_Purchasing_Document_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item", "DeliveryAddress", "Account",
			"PurchasingDocument",
		}
	}

	caller.AsyncGetPurchaseRequisition(
		inoutSDC.PurchaseRequisition.PurchaseRequisition,
		inoutSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchaseRequisitionItem,
		inoutSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchasingDocument,
		inoutSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchasingDocumentItem,
		accepter,
	)
}
