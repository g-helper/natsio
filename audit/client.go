package audit

import "github.com/g-helper/natsio"

type ClientInterface interface {
	SendAudit(req SendAuditReq) (res SendAuditRes, err error)
}

func NewClient() ClientInterface {
	return &client{}
}

type client struct {
}

func (c client) SendAudit(req SendAuditReq) (res SendAuditRes, err error) {
	return natsio.ClientRequest[SendAuditReq, SendAuditRes](QueueName.SendAudit, req)
}
