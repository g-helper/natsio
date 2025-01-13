package natsiostaff

import (
	"encoding/json"
	"errors"

	"github.com/nats-io/nats.go"

	"github.com/g-helper/natsio"
)

var QueueName = struct {
	CheckPermission           string
	GetStaffInfo              string
	GetPartnerInfo            string
	GetPartnerInfoByAccessKey string
}{
	GetStaffInfo:              "get.staff.info",
	CheckPermission:           "check.permission",
	GetPartnerInfo:            "get_partner_info",
	GetPartnerInfoByAccessKey: "get.partner.info.by.access.key",
}

type QueueService interface {
	GetStaffInfo(req GetStaffInfoReq) (res *GetStaffInfoRes, err error)
	CheckPermission(req CheckPermissionReq) (res *CheckPermissionRes, err error)
	GetPartnerInfo(req GetPartnerInfoByCodeReq) (res *GetPartnerInfoRes, err error)
	GetPartnerInfoByAccessKey(req GetPartnerInfoByAccessKeyReq) (res *GetPartnerInfoByAccessKeyRes, err error)
}

type UnimplementedStaffServer struct{}

func (UnimplementedStaffServer) GetStaffInfo(req GetStaffInfoReq) (res *GetStaffInfoRes, err error) {
	return nil, errors.New("method GetStaffInfo not implemented")
}

func _Staff_Get_Info(q interface{}) nats.MsgHandler {
	queueStaff := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req GetStaffInfoReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queueStaff.GetStaffInfo(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}

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

func (UnimplementedStaffServer) GetPartnerInfoByAccessKey(req GetPartnerInfoByAccessKeyReq) (res *GetPartnerInfoByAccessKeyRes, err error) {
	return nil, errors.New("method GetPartnerInfoByAccessKey not implemented")
}

// ADD METHOD
func _Staff_Add_GetPartnerInfoByAccessKey(q interface{}) nats.MsgHandler {
	queueStaff := q.(QueueService)
	return func(msg *nats.Msg) {
		var (
			req GetPartnerInfoByAccessKeyReq
		)
		_ = json.Unmarshal(msg.Data, &req)
		res, err := queueStaff.GetPartnerInfoByAccessKey(req)
		if err != nil {
			_ = natsio.Response(msg, nil, err.Error())
			return
		}
		_ = natsio.Response(msg, res, "")
		return
	}
}
