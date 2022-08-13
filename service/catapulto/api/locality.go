package api

type LocalityRequest struct {
	Term  string `json:"term"`
	ISO   string `json:"iso,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type LocalityResponse struct {
	Id          string `json:"id"`
	Region1     string `json:"region_1"`
	Region1Type string `json:"region1_type"`
	Region2     string `json:"region_2,omitempty"`
	Region2Type string `json:"region_2_type,omitempty"`
	Region3     string `json:"region_3,omitempty"`
	Locality    string `json:"locality"`
	Zip         string `json:"zip"`
	AOID        string `json:"aoid"`
	ShortName   string `json:"shortname"`
	Slug        string `json:"slug"`
	ISO         string `json:"iso,omitempty"`
}
