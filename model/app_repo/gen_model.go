package app_repo

import (
	"time"
)

// App
//go:generate gormgen -structs App -input .
type App struct {
	Id          string    // id字段,服务端自动生成
	Name        string    // app的名字
	Description string    // app的描述内容
	AppSecret   string    // app的密钥,可重置
	OwnerId     string    // app拥有者的id
	CreateDate  time.Time `gorm:"time"` // app创建日期
	UpdateDate  time.Time `gorm:"time"` // app的修改日期
}
