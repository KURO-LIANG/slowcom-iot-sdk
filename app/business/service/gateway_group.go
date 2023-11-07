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
		resData = new(entity.GatewayGroupAddOrUpdateRes)
		var mapData = res.Data.(map[string]interface{})
		resData.GroupId = mapData["groupId"].(string)
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

// GetGroupDetailByGatewayId 根据网关ID获取网关的绑定状态
func (s *GatewayGroupRequest) GetGroupDetailByGatewayId(gatewayId string) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprint("/device/group/detailFromGateway/", gatewayId))
	return
}

// GroupListByIds 根据组ID列表获取组详情
func (s *GatewayGroupRequest) GroupListByIds(groupIds []string) (res *http.IotRes, err error) {
	res, err = s.IotClient.PostJson("/device/group/room/groupListByIds", map[string]interface{}{
		"ids": groupIds,
	})
	return
}

// SceneList 查询组的情景列表
func (s *GatewayGroupRequest) SceneList(groupId string, isPrivate uint8) (res *http.IotRes, err error) {
	res, err = s.IotClient.Get(fmt.Sprintf("/scene/detail/%s/%d", groupId, isPrivate))
	return
}
