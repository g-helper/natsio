package audit

type SendAuditReq struct {
	Data   []Audit `json:"data"`
	Source string  `json:"source"`
}

type Audit struct {
	TargetId string      `json:"targetId"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	ActionBy string      `json:"staff"`
}

type SendAuditRes struct {
	TotalSuccess int64   `json:"totalSuccess"`
	TotalFailed  int64   `json:"totalFailed"`
	ListFailed   []Audit `bson:"listFailed"`
}
