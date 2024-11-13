package natsiogaming

import (
	"encoding/json"
	"errors"

	"github.com/g-helper/natsio"
	"github.com/nats-io/nats.go"
)

var QueueName = struct {
	GetGameInfo string
}{
	GetGameInfo: "get.game.info",
}

type QueueService interface {
	GetGameInfo(req GetGameInfoReq) (res *GetGameInfoRes, err error)
}

type UnimplementedServer struct{}

func (UnimplementedServer) GetGameInfo(req GetGameInfoReq) (res *GetGameInfoRes, err error) {
	return nil, errors.New("method GetGameInfo not implemented")
}

func _Gaming_Add_GetGameInfo(q interface{}) nats.MsgHandler {
	queue := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req GetGameInfoReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queue.GetGameInfo(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}
