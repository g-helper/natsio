package natsiostaff

import "github.com/g-helper/natsio"

type ClientInterface interface {
	CheckPermission(req CheckPermissionReq) (res CheckPermissionRes, err error)
	GetPartnerInfo(req GetPartnerInfoByCodeReq) (res GetPartnerInfoRes, err error)
	GetPartnerInfoByAccessKey(req GetPartnerInfoByAccessKeyReq) (res GetPartnerInfoByAccessKeyRes, err error)
}

func NewClient() ClientInterface {
	return &client{}
}

type client struct {
}

func (c client) GetPartnerInfoByAccessKey(req GetPartnerInfoByAccessKeyReq) (res GetPartnerInfoByAccessKeyRes, err error) {
	return natsio.ClientRequest[GetPartnerInfoByAccessKeyReq, GetPartnerInfoByAccessKeyRes](QueueName.GetPartnerInfoByAccessKey, req)
}

func (c client) GetPartnerInfo(req GetPartnerInfoByCodeReq) (res GetPartnerInfoRes, err error) {
	return natsio.ClientRequest[GetPartnerInfoByCodeReq, GetPartnerInfoRes](QueueName.GetPartnerInfo, req)
}

// CheckPermission ...
func (c client) CheckPermission(req CheckPermissionReq) (res CheckPermissionRes, err error) {
	return natsio.ClientRequest[CheckPermissionReq, CheckPermissionRes](QueueName.CheckPermission, req)
}
