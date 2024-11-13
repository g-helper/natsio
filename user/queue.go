package natsiouser

import (
	"encoding/json"
	"errors"

	"github.com/g-helper/natsio"
	"github.com/nats-io/nats.go"
)

var QueueName = struct {
	GetUserInfo string
	UpsertUser  string
}{
	GetUserInfo: "get.user.info",
	UpsertUser:  "upsert.user",
}

type QueueService interface {
	GetUserInfo(req GetUserInfoReq) (res *GetUserInfoRes, err error)
	UpsertUser(req UpsertUserReq) (res *UserInfo, err error)
}

type UnimplementedServer struct{}

func (UnimplementedServer) GetUserInfo(req GetUserInfoReq) (res *GetUserInfoRes, err error) {
	return nil, errors.New("method GetUserInfo not implemented")
}

func (UnimplementedServer) UpsertUser(req UpsertUserReq) (res *UserInfo, err error) {
	return nil, errors.New("method UpsertUser not implemented")
}

func _User_Add_UpsertUser(q interface{}) nats.MsgHandler {
	queueStaff := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req UpsertUserReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queueStaff.UpsertUser(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}

func _User_Add_GetUserInfo(q interface{}) nats.MsgHandler {
	queueStaff := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req GetUserInfoReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queueStaff.GetUserInfo(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}
