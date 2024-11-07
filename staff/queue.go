package natsiostaff

import (
	"encoding/json"
	"errors"

	"github.com/g-helper/natsio"
	"github.com/nats-io/nats.go"
)

var QueueName = struct {
	CheckPermission string
	GetPartnerInfo  string
}{
	CheckPermission: "check.permission",
	GetPartnerInfo:  "get_partner_info",
}

type QueueService interface {
	CheckPermission(req CheckPermissionReq) (res *CheckPermissionRes, err error)
	GetPartnerInfo(req GetPartnerInfoByCodeReq) (res *GetPartnerInfoRes, err error)
}

type UnimplementedStaffServer struct{}

func (UnimplementedStaffServer) CheckPermission(req CheckPermissionReq) (res *CheckPermissionRes, err error) {
	return nil, errors.New("method GetUserInfo not implemented")
}

func _Staff_Add_CheckPermission(q interface{}) nats.MsgHandler {
	queueStaff := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req CheckPermissionReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queueStaff.CheckPermission(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}

func (UnimplementedStaffServer) GetPartnerInfo(req GetPartnerInfoByCodeReq) (res *GetPartnerInfoRes, err error) {
	return nil, errors.New("method GetUserInfo not implemented")
}

// ADD METHOD
func _Staff_Add_GetPartnerInfo(q interface{}) nats.MsgHandler {
	queueStaff := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req GetPartnerInfoByCodeReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queueStaff.GetPartnerInfo(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}
