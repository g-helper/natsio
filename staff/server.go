package natsiostaff

import (
	"github.com/g-helper/natsio"
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
		{
			Subject: QueueName.GetPartnerInfo,
			Worker:  Worker,
			Handle:  _Staff_Add_GetPartnerInfo,
		},
		{
			Subject: QueueName.GetPartnerInfoByAccessKey,
			Worker:  Worker,
			Handle:  _Staff_Add_GetPartnerInfoByAccessKey,
		},
	},
}

// RegisterServer ...
func RegisterServer(queue QueueService) error {
	return serviceDesc.Register(queue)
}
