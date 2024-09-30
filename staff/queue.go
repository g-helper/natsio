package natsiostaff

import (
	"errors"
)

var QueueName = struct {
	CheckPermission string
}{
	CheckPermission: "check.permission",
}

type QueueService interface {
	CheckPermission(req CheckPermissionReq) (res *CheckPermissionRes, err error)
}

type UnimplementedStaffServer struct{}

func (UnimplementedStaffServer) CheckPermission(req CheckPermissionReq) (res *CheckPermissionRes, err error) {
	return nil, errors.New("method GetUserInfo not implemented")
}
