///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package app_api_relation_repo

import (
	"fmt"
	"time"

	"github.com/liuwqiang/tykd/model"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func NewModel() *AppApiRelation {
	return new(AppApiRelation)
}

func NewQueryBuilder() *appApiRelationRepoQueryBuilder {
	return new(appApiRelationRepoQueryBuilder)
}

func (t *AppApiRelation) Create(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {
		return fmt.Errorf("%s create err", err)
	}
	return nil
}

type appApiRelationRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *appApiRelationRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *appApiRelationRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&AppApiRelation{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return fmt.Errorf("%s updates err", err)
	}
	return nil
}

func (qb *appApiRelationRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&AppApiRelation{}).Error; err != nil {
		return fmt.Errorf("%s delete err", err)
	}
	return nil
}

func (qb *appApiRelationRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&AppApiRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *appApiRelationRepoQueryBuilder) First(db *gorm.DB) (*AppApiRelation, error) {
	ret := &AppApiRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *appApiRelationRepoQueryBuilder) QueryOne(db *gorm.DB) (*AppApiRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *appApiRelationRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*AppApiRelation, error) {
	var ret []*AppApiRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *appApiRelationRepoQueryBuilder) Limit(limit int) *appApiRelationRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) Offset(offset int) *appApiRelationRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereApiId(p model.Predicate, value string) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api_id", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereApiIdIn(value []string) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api_id", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereApiIdNotIn(value []string) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByApiId(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "api_id "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereAppId(p model.Predicate, value string) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_id", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereAppIdIn(value []string) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_id", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereAppIdNotIn(value []string) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByAppId(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "app_id "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereExpiredDate(p model.Predicate, value time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "expired_date", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereExpiredDateIn(value []time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "expired_date", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereExpiredDateNotIn(value []time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "expired_date", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByExpiredDate(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "expired_date "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereRate(p model.Predicate, value int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "rate", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereRateIn(value []int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "rate", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereRateNotIn(value []int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "rate", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByRate(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "rate "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WherePer(p model.Predicate, value int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "per", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WherePerIn(value []int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "per", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WherePerNotIn(value []int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "per", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByPer(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "per "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereQuota(p model.Predicate, value int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "quota", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereQuotaIn(value []int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "quota", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereQuotaNotIn(value []int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "quota", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByQuota(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "quota "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereQuotaRenewalRate(p model.Predicate, value int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "quota_renewal_rate", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereQuotaRenewalRateIn(value []int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "quota_renewal_rate", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereQuotaRenewalRateNotIn(value []int32) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "quota_renewal_rate", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByQuotaRenewalRate(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "quota_renewal_rate "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereExtraData(p model.Predicate, value datatypes.JSON) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "extra_data", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereExtraDataIn(value []datatypes.JSON) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "extra_data", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereExtraDataNotIn(value []datatypes.JSON) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "extra_data", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByExtraData(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "extra_data "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereCreateDate(p model.Predicate, value time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_date", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereCreateDateIn(value []time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_date", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereCreateDateNotIn(value []time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_date", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByCreateDate(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_date "+order)
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereUpdateDate(p model.Predicate, value time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_date", p),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereUpdateDateIn(value []time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_date", "IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) WhereUpdateDateNotIn(value []time.Time) *appApiRelationRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_date", "NOT IN"),
		value,
	})
	return qb
}

func (qb *appApiRelationRepoQueryBuilder) OrderByUpdateDate(asc bool) *appApiRelationRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "update_date "+order)
	return qb
}