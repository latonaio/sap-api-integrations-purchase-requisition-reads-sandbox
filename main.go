package main

import (
	sap_api_caller "sap-api-integrations-purchase-requisition-reads/SAP_API_Caller"
	"sap-api-integrations-purchase-requisition-reads/file_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := file_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Purchase_Requisition_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetPurchaseRequisition(
		inoutSDC.PurchaseRequisition,
		inoutSDC.PurchaseRequisition.PurchaseRequisitionItem,
		inoutSDC.PurchasingDocument.PurchasingDocumentItem,
	)
}