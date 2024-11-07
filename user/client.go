package natsiouser

import "github.com/g-helper/natsio"

type ClientInterface interface {
	GetUserInfo(req GetUserInfoReq) (res GetUserInfoRes, err error)
}

func NewClient() ClientInterface {
	return &client{}
}

type client struct {
}

func (c client) GetUserInfo(req GetUserInfoReq) (res GetUserInfoRes, err error) {
	return natsio.ClientRequest[GetUserInfoReq, GetUserInfoRes](QueueName.GetUserInfo, req)
}
