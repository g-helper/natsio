package audit

import (
	"encoding/json"
	"errors"

	"github.com/nats-io/nats.go"

	"github.com/g-helper/natsio"
)

var QueueName = struct {
	SendAudit string
}{
	SendAudit: "send.audits",
}

type QueueService interface {
	SendAudit(req SendAuditReq) (res *SendAuditRes, err error)
}

type UnimplementedServer struct{}

func (UnimplementedServer) SendAudit(req SendAuditReq) (res *SendAuditRes, err error) {
	return nil, errors.New("method SendAudit not implemented")
}

func _Audit_Send_Audit(q interface{}) nats.MsgHandler {
	queue := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req SendAuditReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queue.SendAudit(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}
