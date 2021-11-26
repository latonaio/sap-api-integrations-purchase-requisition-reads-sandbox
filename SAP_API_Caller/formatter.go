package sap_api_caller

type PurchaseRequisitionReads struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	APISchema           string `json:"api_schema"`
	PurchaseRequisition string `json:"purchase_requisition"`
	Deleted             string `json:"deleted"`
}

type PurchaseRequisition struct {
	PurchaseRequisition     string `json:"PurchaseRequisition"`
	PurchaseRequisitionType string `json:"PurchaseRequisitionType"`
	SourceDetermination     string `json:"SourceDetermination"`
}

type PurchaseRequisitionItem struct {
	PurchaseRequisition            string `json:"PurchaseRequisition"`
    PurchaseRequisitionItem        int    `json:"PurchaseRequisitionItem"`
    PurchasingDocument             string `json:"PurchasingDocument"`
    PurchasingDocumentItem         int    `json:"PurchasingDocumentItem"`
    PurReqnReleaseStatus           string `json:"PurReqnReleaseStatus"`
    PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
    PurchaseRequisitionItemText    string `json:"PurchaseRequisitionItemText"`
    AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
    Material                       string `json:"Material"`
    MaterialGroup                  string `json:"MaterialGroup"`
    PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
    RequestedQuantity              float64 `json:"RequestedQuantity"`
    BaseUnit                       string `json:"BaseUnit"`
    PurchaseRequisitionPrice       float64 `json:"PurchaseRequisitionPrice"`
    PurReqnPriceQuantity           float64 `json:"PurReqnPriceQuantity"`
    MaterialGoodsReceiptDuration   int    `json:"MaterialGoodsReceiptDuration"`
    ReleaseCode                    string `json:"ReleaseCode"`
    PurchaseRequisitionReleaseDate string `json:"PurchaseRequisitionReleaseDate"`
    PurchasingOrganization         string `json:"PurchasingOrganization"`
    PurchasingGroup                string `json:"PurchasingGroup"`
    Plant                          string `json:"Plant"`
    CompanyCode                    string `json:"CompanyCode"`
    SourceOfSupplyIsAssigned       string `json:"SourceOfSupplyIsAssigned"`
    SupplyingPlant                 string `json:"SupplyingPlant"`
    OrderedQuantity                float64 `json:"OrderedQuantity"`
    DeliveryDate                   string `json:"DeliveryDate"`
    CreationDate                   string `json:"CreationDate"`
    ProcessingStatus               string `json:"ProcessingStatus"`
    ExternalApprovalStatus         string `json:"ExternalApprovalStatus"`
    PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
    Supplier                       string `json:"Supplier"`
    FixedSupplier                  string `json:"FixedSupplier"`
    RequisitionerName              string `json:"RequisitionerName"`
    PurReqnItemCurrency            string `json:"PurReqnItemCurrency"`
    MaterialPlannedDeliveryDurn    int    `json:"MaterialPlannedDeliveryDurn"`
    StorageLocation                string `json:"StorageLocation"`
    PurReqnSourceOfSupplyType      string `json:"PurReqnSourceOfSupplyType"`
    ConsumptionPosting             string `json:"ConsumptionPosting"`
    PurReqnOrigin                  string `json:"PurReqnOrigin"`
    IsPurReqnBlocked               string `json:"IsPurReqnBlocked"`
    PurchaseRequisitionStatus      string `json:"PurchaseRequisitionStatus"`
    Batch                          string `json:"Batch"`
    GoodsReceiptIsExpected         string `json:"GoodsReceiptIsExpected"`
    GoodsReceiptIsNonValuated      string `json:"GoodsReceiptIsNonValuated"`
    RequirementTracking            string `json:"RequirementTracking"`
    MRPController                  string `json:"MRPController"`
    Reservation                    int    `json:"Reservation"`
    LastChangeDateTime             string `json:"LastChangeDateTime"`
    IsDeleted                      string `json:"IsDeleted"`
}

type PurchaseReqnAcctAssgmtNumber  struct {
	PurchaseRequisition            string `json:"PurchaseRequisition"`
    PurchaseRequisitionItem        int    `json:"PurchaseRequisitionItem"`
    PurchaseReqnAcctAssgmtNumber   int    `json:"PurchaseReqnAcctAssgmtNumber"`
    CostCenter                     string `json:"CostCenter"`
    MasterFixedAsset               string `json:"MasterFixedAsset"`
    FixedAsset                     string `json:"FixedAsset"`
    ProjectNetwork                 string `json:"ProjectNetwork"`
    CostElement                    string `json:"CostElement"`
    CostObject                     string `json:"CostObject"`
    GLAccount                      string `json:"GLAccount"`
    BusinessArea                   string `json:"BusinessArea"`
    SalesOrder                     string `json:"SalesOrder"`
    SalesOrderItem                 int    `json:"SalesOrderItem"`
    SalesOrderScheduleLine         int    `json:"SalesOrderScheduleLine"`
    OrderID                        string `json:"OrderID"`
    UnloadingPointName             string `json:"UnloadingPointName"`
    ControllingArea                string `json:"ControllingArea"`
    ProfitabilitySegment           int    `json:"ProfitabilitySegment"`
    ProfitCenter                   string `json:"ProfitCenter"`
    FunctionalArea                 string `json:"FunctionalArea"`
    GoodsRecipientName             string `json:"GoodsRecipientName"`
    CostCtrActivityType            string `json:"CostCtrActivityType"`
    WBSElement                     int    `json:"WBSElement"`
    IsDeleted                      string `json:"IsDeleted"`
}   

type DeliveryAddressID struct {
	PurchaseRequisition            string `json:"PurchaseRequisition"`
    PurchaseRequisitionItem        int    `json:"PurchaseRequisitionItem"`
    AddressID                      string `json:"AddressID"`
    Country                        string `json:"Country"`
    Region                         string `json:"Region"`
    StreetName                     string `json:"StreetName"`
    CityName                       string `json:"CityName"`
    CorrespondenceLanguage         string `json:"CorrespondenceLanguage"`
    FaxNumber                      string `json:"FaxNumber"`
    PhoneNumber                    string `json:"PhoneNumber"`
}
