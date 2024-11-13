package natsiogaming

import "github.com/g-helper/natsio"

type ClientInterface interface {
	GetGameInfo(req GetGameInfoReq) (res GetGameInfoRes, err error)
}

func NewClient() ClientInterface {
	return &client{}
}

type client struct {
}

func (c client) GetGameInfo(req GetGameInfoReq) (res GetGameInfoRes, err error) {
	return natsio.ClientRequest[GetGameInfoReq, GetGameInfoRes](QueueName.GetGameInfo, req)
}
