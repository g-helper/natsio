package natsiouser

type GetUserInfoReq struct {
	SourceIds []string `json:"sourceIds"`
	Ids       []string `json:"ids"`
	PartnerId string   `json:"partnerId"`
}

type GetUserInfoRes struct {
	Data []UserInfo `json:"data"`
}

type UserInfo struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	SourceId    string `json:"sourceId"`
	Avatar      string `json:"avatar"`
	PartnerId   string `json:"partnerId"`
	PartnerCode string `json:"partnerCode"`
}
