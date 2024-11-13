package natsiogaming

import "time"

type GetGameInfoReq struct {
	Ids        []string `json:"ids"`
	Code       []string `json:"code"`
	PartnerIds []string `json:"partnerIds"`
}

type GetGameInfoRes struct {
	Data []GameInfo `json:"data"`
}

type GameInfo struct {
	ID          string    `json:"_id"`
	Title       string    `json:"title"`
	PartnerId   string    `json:"partnerId"`
	Desc        string    `json:"desc"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	IsAvailable bool      `json:"isAvailable"`
}
