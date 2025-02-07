package audit

import "github.com/g-helper/natsio"

const (
	Worker = "audit_workers"
)

var serviceDesc = natsio.ServiceDesc{
	Queues: []natsio.QueueDesc{
		{
			Subject: QueueName.SendAudit,
			Worker:  Worker,
			Handle:  _Audit_Send_Audit,
		},
	},
}

// RegisterServer ...
func RegisterServer(queue QueueService) error {
	return serviceDesc.Register(queue)
}
