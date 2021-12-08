package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-requisition-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Header{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionType:        data.PurchaseRequisitionType,
		SourceDetermination:            data.SourceDetermination,
	}, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Item{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		PurReqnReleaseStatus:           data.PurReqnReleaseStatus,
		PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
		PurchaseRequisitionItemText:    data.PurchaseRequisitionItemText,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		Material:                       data.Material,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
		RequestedQuantity:              data.RequestedQuantity,
		BaseUnit:                       data.BaseUnit,
		PurchaseRequisitionPrice:       data.PurchaseRequisitionPrice,
		PurReqnPriceQuantity:           data.PurReqnPriceQuantity,
		MaterialGoodsReceiptDuration:   data.MaterialGoodsReceiptDuration,
		ReleaseCode:                    data.ReleaseCode,
		PurchaseRequisitionReleaseDate: data.PurchaseRequisitionReleaseDate,
		PurchasingOrganization:         data.PurchasingOrganization,
		PurchasingGroup:                data.PurchasingGroup,
		Plant:                          data.Plant,
		CompanyCode:                    data.CompanyCode,
		SourceOfSupplyIsAssigned:       data.SourceOfSupplyIsAssigned,
		SupplyingPlant:                 data.SupplyingPlant,
		OrderedQuantity:                data.OrderedQuantity,
		DeliveryDate:                   data.DeliveryDate,
		CreationDate:                   data.CreationDate,
		ProcessingStatus:               data.ProcessingStatus,
		ExternalApprovalStatus:         data.ExternalApprovalStatus,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		Supplier:                       data.Supplier,
		FixedSupplier:                  data.FixedSupplier,
		RequisitionerName:              data.RequisitionerName,
		PurReqnItemCurrency:            data.PurReqnItemCurrency,
		MaterialPlannedDeliveryDurn:    data.MaterialPlannedDeliveryDurn,
		StorageLocation:                data.StorageLocation,
		PurReqnSourceOfSupplyType:      data.PurReqnSourceOfSupplyType,
		ConsumptionPosting:             data.ConsumptionPosting,
		PurReqnOrigin:                  data.PurReqnOrigin,
		IsPurReqnBlocked:               data.IsPurReqnBlocked,
		PurchaseRequisitionStatus:      data.PurchaseRequisitionStatus,
		Batch:                          data.Batch,
		GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
		RequirementTracking:            data.RequirementTracking,
		MRPController:                  data.MRPController,
		Reservation:                    data.Reservation,
		LastChangeDateTime:             data.LastChangeDateTime,
		IsDeleted:                      data.IsDeleted,
	}, nil
}

func ConvertToDeliveryAddress(raw []byte, l *logger.Logger) (*DeliveryAddress, error) {
	pm := &responses.DeliveryAddress{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to DeliveryAddress. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &DeliveryAddress{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		AddressID:                      data.AddressID,
		Country:                        data.Country,
		Region:                         data.Region,
		StreetName:                     data.StreetName,
		CityName:                       data.CityName,
		CorrespondenceLanguage:         data.CorrespondenceLanguage,
		FaxNumber:                      data.FaxNumber,
		PhoneNumber:                    data.PhoneNumber,
	}, nil
}

func ConvertToAccount(raw []byte, l *logger.Logger) (*Account, error) {
	pm := &responses.Account{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Account. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Account{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
        PurchaseReqnAcctAssgmtNumber:   data.PurchaseReqnAcctAssgmtNumber,
		CostCenter:                     data.CostCenter,
		MasterFixedAsset:               data.MasterFixedAsset,
		FixedAsset:                     data.FixedAsset,
		ProjectNetwork:                 data.ProjectNetwork,
		CostElement:                    data.CostElement,
		CostObject:                     data.CostObject,
		GLAccount:                      data.GLAccount,
		BusinessArea:                   data.BusinessArea,
		SalesOrder:                     data.SalesOrder,
		SalesOrderItem:                 data.SalesOrderItem,
		SalesOrderScheduleLine:         data.SalesOrderScheduleLine,
		OrderID:                        data.OrderID,
		UnloadingPointName:             data.UnloadingPointName,
		ControllingArea:                data.ControllingArea,
		ProfitabilitySegment:           data.ProfitabilitySegment,
		ProfitCenter:                   data.ProfitCenter,
		FunctionalArea:                 data.FunctionalArea,
		GoodsRecipientName:             data.GoodsRecipientName,
		CostCtrActivityType:            data.CostCtrActivityType,
		WBSElement:                     data.WBSElement,
		IsDeleted:                      data.IsDeleted,
	}, nil
}
