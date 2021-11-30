package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-requisition-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToPurchaseRequisition(raw []byte, l *logger.Logger) *PurchaseRequisition {
	pm := &responses.PurchaseRequisition{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PurchaseRequisition{
		PurchaseRequisition            data.PurchaseRequisition,
		PurchaseRequisitionType        data.PurchaseRequisitionType,
		SourceDetermination            data.SourceDetermination,
		PurchaseRequisitionItem        data.PurchaseRequisitionItem,
		PurchasingDocument             data.PurchasingDocument,
		PurchasingDocumentItem         data.PurchasingDocumentItem,
		PurReqnReleaseStatus           data.PurReqnReleaseStatus,
		PurchasingDocumentItemCategory data.PurchasingDocumentItemCategory,
		PurchaseRequisitionItemText    data.PurchaseRequisitionItemText,
		AccountAssignmentCategory      data.AccountAssignmentCategory,
		Material                       data.Material,
		MaterialGroup                  data.MaterialGroup,
		PurchasingDocumentCategory     data.PurchasingDocumentCategory,
		RequestedQuantity              data.RequestedQuantity,
		BaseUnit                       data.BaseUnit,
		PurchaseRequisitionPrice       data.PurchaseRequisitionPrice,
		PurReqnPriceQuantity           data.PurReqnPriceQuantity,
		MaterialGoodsReceiptDuration   data.MaterialGoodsReceiptDuration,
		ReleaseCode                    data.ReleaseCode,
		PurchaseRequisitionReleaseDate data.PurchaseRequisitionReleaseDate,
		PurchasingOrganization         data.PurchasingOrganization,
		PurchasingGroup                data.PurchasingGroup,
		Plant                          data.Plant,
		CompanyCode                    data.CompanyCode,
		SourceOfSupplyIsAssigned       data.SourceOfSupplyIsAssigned,
		SupplyingPlant                 data.SupplyingPlant,
		OrderedQuantity                data.OrderedQuantity,
		DeliveryDate                   data.DeliveryDate,
		CreationDate                   data.CreationDate,
		ProcessingStatus               data.ProcessingStatus,
		ExternalApprovalStatus         data.ExternalApprovalStatus,
		PurchasingInfoRecord           data.PurchasingInfoRecord,
		Supplier                       data.Supplier,
		FixedSupplier                  data.FixedSupplier,
		RequisitionerName              data.RequisitionerName,
		PurReqnItemCurrency            data.PurReqnItemCurrency,
		MaterialPlannedDeliveryDurn    data.MaterialPlannedDeliveryDurn,
		StorageLocation                data.StorageLocation,
		PurReqnSourceOfSupplyType      data.PurReqnSourceOfSupplyType,
		ConsumptionPosting             data.ConsumptionPosting,
		PurReqnOrigin                  data.PurReqnOrigin,
		IsPurReqnBlocked               data.IsPurReqnBlocked,
		PurchaseRequisitionStatus      data.PurchaseRequisitionStatus,
		Batch                          data.Batch,
		GoodsReceiptIsExpected         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated      data.GoodsReceiptIsNonValuated,
		RequirementTracking            data.RequirementTracking,
		MRPController                  data.MRPController,
		Reservation                    data.Reservation,
		LastChangeDateTime             data.LastChangeDateTime,
		IsDeleted                      data.IsDeleted,
		PurchaseReqnAcctAssgmtNumber   data.PurchaseReqnAcctAssgmtNumber,
		CostCenter                     data.CostCenter,
		MasterFixedAsset               data.MasterFixedAsset,
		FixedAsset                     data.FixedAsset,
		ProjectNetwork                 data.ProjectNetwork,
		CostElement                    data.CostElement,
		CostObject                     data.CostObject,
		GLAccount                      data.GLAccount,
		BusinessArea                   data.BusinessArea,
		SalesOrder                     data.SalesOrder,
		SalesOrderItem                 data.SalesOrderItem,
		SalesOrderScheduleLine         data.SalesOrderScheduleLine,
		OrderID                        data.OrderID,
		UnloadingPointName             data.UnloadingPointName,
		ControllingArea                data.ControllingArea,
		ProfitabilitySegment           data.ProfitabilitySegment,
		ProfitCenter                   data.ProfitCenter,
		FunctionalArea                 data.FunctionalArea,
		GoodsRecipientName             data.GoodsRecipientName,
		CostCtrActivityType            data.CostCtrActivityType,
		WBSElement                     data.WBSElement,
		IsDeleted                      data.IsDeleted,
		AddressID                      data.AddressID,
		Country                        data.Country,
		Region                         data.Region,
		StreetName                     data.StreetName,
		CityName                       data.CityName,
		CorrespondenceLanguage         data.CorrespondenceLanguage,
		FaxNumber                      data.FaxNumber,
		PhoneNumber                    data.PhoneNumber,
	}
}
