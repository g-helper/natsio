package natsiostaff

import (
	"encoding/json"

	"github.com/g-helper/natsio"

	"github.com/nats-io/nats.go"
)

const (
	Worker = "staff_workers"
)

var serviceDesc = natsio.ServiceDesc{
	Queues: []natsio.QueueDesc{
		{
			Subject: QueueName.CheckPermission,
			Worker:  Worker,
			Handle:  _Staff_Add_CheckPermission,
		},
	},
}

func _Staff_Add_CheckPermission(q interface{}) nats.MsgHandler {
	queueStaff := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req CheckPermissionReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queueStaff.CheckPermission(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}

// RegisterServer ...
func RegisterServer(queue QueueService) error {
	return serviceDesc.Register(queue)
}
