package natsiostaff

import "github.com/g-helper/natsio"

type ClientInterface interface {
	CheckPermission(req CheckPermissionReq) (res CheckPermissionRes, err error)
	GetPartnerInfo(req GetPartnerInfoByCodeReq) (res GetPartnerInfoRes, err error)
}

func NewClient() ClientInterface {
	return &client{}
}

type client struct {
}

func (c client) GetPartnerInfo(req GetPartnerInfoByCodeReq) (res GetPartnerInfoRes, err error) {
	return natsio.ClientRequest[GetPartnerInfoByCodeReq, GetPartnerInfoRes](QueueName.CheckPermission, req)
}

// CheckPermission ...
func (c client) CheckPermission(req CheckPermissionReq) (res CheckPermissionRes, err error) {
	return natsio.ClientRequest[CheckPermissionReq, CheckPermissionRes](QueueName.CheckPermission, req)
}
