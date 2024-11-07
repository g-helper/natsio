package natsiouser

import "github.com/g-helper/natsio"

const (
	Worker = "user_workers"
)

var serviceDesc = natsio.ServiceDesc{
	Queues: []natsio.QueueDesc{
		{
			Subject: QueueName.GetUserInfo,
			Worker:  Worker,
			Handle:  _User_Add_GetUserInfo,
		},
	},
}

// RegisterServer ...
func RegisterServer(queue QueueService) error {
	return serviceDesc.Register(queue)
}
