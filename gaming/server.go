package natsiogaming

import "github.com/g-helper/natsio"

const (
	Worker = "gaming_workers"
)

var serviceDesc = natsio.ServiceDesc{
	Queues: []natsio.QueueDesc{
		{
			Subject: QueueName.GetGameInfo,
			Worker:  Worker,
			Handle:  _Gaming_Add_GetGameInfo,
		},
	},
}

// RegisterServer ...
func RegisterServer(queue QueueService) error {
	return serviceDesc.Register(queue)
}
