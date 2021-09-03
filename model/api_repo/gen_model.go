package api_repo

import (
	"time"
)

// Api
//go:generate gormgen -structs Api -input .
type Api struct {
	Id          string    // id字段,服务端自动生成
	Name        string    // api的名字
	Description string    // api的描述内容
	ListenPath  string    // 监听的地址
	TargetUrl   string    // 目标服务的地址
	OnwerId     string    // api拥有者的id
	GroupId     string    // api分组的id
	GroupName   string    // api分组的名称
	CreateDate  time.Time `gorm:"time"` // api创建日期
	UpdateDate  time.Time `gorm:"time"` // api的修改日期
	AuthMode    string    // 认证模式 [OPEN(无认证),APPID(简单认证),SIGNATURE(签名认证)]
}
