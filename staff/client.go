package natsiostaff

import "github.com/g-helper/natsio"

type ClientInterface interface {
	CheckPermission(req CheckPermissionReq) (res CheckPermissionRes, err error)
}

func NewClient() ClientInterface {
	return &client{}
}

type client struct {
}

// CheckPermission ...
func (c client) CheckPermission(req CheckPermissionReq) (res CheckPermissionRes, err error) {
	return natsio.ClientRequest[CheckPermissionReq, CheckPermissionRes](QueueName.CheckPermission, req)
}
