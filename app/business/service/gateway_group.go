package service

import (
	"fmt"
	"github.com/kuro-liang/slowcom-iot-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-iot-sdk/http"
)

type GatewayGroupRequest struct {
	IotClient *http.IotClient
}

// Save 添加或修改组
func (s *GatewayGroupRequest) Save(req *entity.GatewayGroupAddOrUpdate) (resData *entity.GatewayGroupAddOrUpdateRes, err error) {
	res, err := s.IotClient.PostJson("/device/group/split/save", req)
	if err == nil {
		resData = res.Data.(*entity.GatewayGroupAddOrUpdateRes)
	}
	return
}

// Delete 删除组
func (s *GatewayGroupRequest) Delete(req *entity.GatewayGroupDelete) (res *http.IotRes, err error) {
	res, err = s.IotClient.PostJson("/device/group/split/delete", req)
	return
}

// GroupDetail 查询组信息
func (s *GatewayGroupRequest) GroupDetail(groupId string) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprint("/device/group/room/detail/", groupId))
	return
}
