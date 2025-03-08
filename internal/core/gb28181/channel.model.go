// Code generated by gowebx, DO AVOID EDIT.
package gb28181

// Channel domain model
type Channel struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	DID       string    `gorm:"column:did;index;notNull;default:'';comment:父级 ID" json:"did"`
	DeviceID  string    `gorm:"column:device_id;index;notNull;default:'';comment:国标编码" json:"device_id"`   // 国标编码
	ChannelID string    `gorm:"column:channel_id;index;notNull;default:'';comment:国标编码" json:"channel_id"` // 国标编码
	Name      string    `gorm:"column:name;notNull;default:'';comment:通道名称" json:"name"`                   // 通道名称
	PTZType   int       `gorm:"column:ptztype;notNull;default:0;comment:云台类型" json:"ptztype"`              // 云台类型
	IsOnline  bool      `gorm:"column:is_online;notNull;default:FALSE;comment:是否在线" json:"is_online"`      // 是否在线
	Ext       DeviceExt `gorm:"column:ext;notNull;default:'{}';type:jsonb" json:"ext"`
}

// TableName database table name
func (*Channel) TableName() string {
	return "channels"
}
