// Code generated by gowebx, DO AVOID EDIT.
package gb28181

import "github.com/ixugo/goweb/pkg/web"

type FindChannelInput struct {
	web.PagerFilter
	DID      string `form:"did"`       // 设备 id
	DeviceID string `form:"device_id"` // 国标编码
	Key      string `form:"key"`       // 名称/国标编码 模糊搜索，id 精确搜索
	// Name     string    `form:"name"`      // 通道名称
	// PTZType  int       `form:"ptztype"`   // 云台类型
	IsOnline bool `form:"is_online"` // 是否在线
}

type EditChannelInput struct {
	DeviceID string    `json:"device_id"` // 国标编码
	Name     string    `json:"name"`      // 通道名称
	PTZType  int       `json:"ptztype"`   // 云台类型
	IsOnline bool      `json:"is_online"` // 是否在线
	Ext      DeviceExt `json:"ext"`
}

type AddChannelInput struct {
	DeviceID string    `json:"device_id"` // 国标编码
	Name     string    `json:"name"`      // 通道名称
	PTZType  int       `json:"ptztype"`   // 云台类型
	IsOnline bool      `json:"is_online"` // 是否在线
	Ext      DeviceExt `json:"ext"`
}
