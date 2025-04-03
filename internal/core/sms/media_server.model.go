// Code generated by gowebx, DO AVOID EDIT.
package sms

import "github.com/ixugo/goweb/pkg/orm"

// DefaultMediaServerID 临时变量，待未来重构分布式流媒体时，移除
const DefaultMediaServerID = "local"

// MediaServer domain model
type MediaServer struct {
	ID                string           `gorm:"primaryKey" json:"id"`
	IP                string           `gorm:"column:ip;notNull;default:''" json:"ip"`
	CreatedAt         orm.Time         `gorm:"column:created_at;notNull;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         orm.Time         `gorm:"column:updated_at;notNull;default:CURRENT_TIMESTAMP" json:"updated_at"`
	HookIP            string           `gorm:"column:hook_ip;notNull;default:''" json:"hook_ip"`
	SDPIP             string           `gorm:"column:sdp_ip;notNull;default:''" json:"sdp_ip"`
	StreamIP          string           `gorm:"column:stream_ip;notNull;default:''" json:"stream_ip"`
	Ports             MediaServerPorts `gorm:"column:ports;notNull;default:'{}';type:jsonb" json:"ports"`
	AutoConfig        bool             `gorm:"column:auto_config;notNull;default:FALSE" json:"auto_config"`
	Secret            string           `gorm:"column:secret;notNull;default:''" json:"secret"`
	HookAliveInterval int              `gorm:"column:hook_alive_interval;notNull;default:0" json:"hook_alive_interval"`
	RTPEnable         bool             `gorm:"column:rtpenable;notNull;default:FALSE" json:"rtpenable"`
	Status            bool             `gorm:"column:status;notNull;default:FALSE" json:"status"`
	RTPPortRange      string           `gorm:"column:rtpport_range;notNull;default:''" json:"rtpport_range"`
	SendRTPPortRange  string           `gorm:"column:send_rtpport_range;notNull;default:''" json:"send_rtpport_range"`
	RecordAssistPort  int              `gorm:"column:record_assist_port;notNull;default:0" json:"record_assist_port"`
	LastKeepaliveAt   orm.Time         `gorm:"column:last_keepalive_at;notNull;default:CURRENT_TIMESTAMP" json:"last_keepalive_at"`
	RecordDay         int              `gorm:"column:record_day;notNull;default:0" json:"record_day"`
	RecordPath        string           `gorm:"column:record_path;notNull;default:''" json:"record_path"`
	Type              string           `gorm:"column:type;notNull;default:''" json:"type"`
	TranscodeSuffix   string           `gorm:"column:transcode_suffix;notNull;default:''" json:"transcode_suffix"`
}

// TableName database table name
func (*MediaServer) TableName() string {
	return "media_servers"
}

func (m *MediaServer) GetSDPIP() string {
	if m.SDPIP != "" {
		return m.SDPIP
	}
	return m.IP
}
