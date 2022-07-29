package serializers

// Resp : struct
type Resp struct {
	Result     interface{} `json:"result"`
	Error      interface{} `json:"error"`
	Pagination interface{} `json:"pagination"`
}

type IPResp struct {
	Country     string `json:"country_name"`
	CountryCode string `json:"country_code"`
	Region      string `json:"region_name"`
	RegionCode  string `json:"region_code"`
}
