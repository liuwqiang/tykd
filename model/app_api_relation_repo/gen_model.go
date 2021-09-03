package app_api_relation_repo

import (
	"time"

	"gorm.io/datatypes"
)

// AppApiRelation
//go:generate gormgen -structs AppApiRelation -input .
type AppApiRelation struct {
	ApiId            string         // api的id字段
	AppId            string         // app的id字段
	ExpiredDate      time.Time      `gorm:"time"` // 过期时间
	Rate             int32          // 速率
	Per              int32          // 和rate配合使用,每per秒允许rate个请求通过
	Quota            int32          // 配额
	QuotaRenewalRate int32          // 和quota配合使用,每quota_renewal_rate个周期更新一次配额
	ExtraData        datatypes.JSON `gorm:"json"` // 额外数据,传递给后端api
	CreateDate       time.Time      `gorm:"time"` // api和app映射的创建日期
	UpdateDate       time.Time      `gorm:"time"` //
}
