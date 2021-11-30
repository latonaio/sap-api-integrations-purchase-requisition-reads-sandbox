package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseRequisition     string `json:"PurchaseRequisition"`
			PurchaseRequisitionType string `json:"PurchaseRequisitionType"`
			SourceDetermination     string `json:"SourceDetermination"`
		} `json:"results"`
	} `json:"d"`
}
