///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package api_repo

import (
	"fmt"
	"time"

	"github.com/liuwqiang/tykd/model"

	"gorm.io/gorm"
)

func NewModel() *Api {
	return new(Api)
}

func NewQueryBuilder() *apiRepoQueryBuilder {
	return new(apiRepoQueryBuilder)
}

func (t *Api) Create(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {
		return fmt.Errorf("%s create err", err)
	}
	return nil
}

type apiRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *apiRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *apiRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Api{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return fmt.Errorf("%s updates err", err)
	}
	return nil
}

func (qb *apiRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Api{}).Error; err != nil {
		return fmt.Errorf("%s delete err", err)
	}
	return nil
}

func (qb *apiRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Api{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *apiRepoQueryBuilder) First(db *gorm.DB) (*Api, error) {
	ret := &Api{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *apiRepoQueryBuilder) QueryOne(db *gorm.DB) (*Api, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *apiRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*Api, error) {
	var ret []*Api
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *apiRepoQueryBuilder) Limit(limit int) *apiRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *apiRepoQueryBuilder) Offset(offset int) *apiRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *apiRepoQueryBuilder) WhereId(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereIdIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereIdNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderById(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereName(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereNameIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereNameNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByName(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereDescription(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereDescriptionIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereDescriptionNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByDescription(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "description "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereListenPath(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "listen_path", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereListenPathIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "listen_path", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereListenPathNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "listen_path", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByListenPath(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "listen_path "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereTargetUrl(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "target_url", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereTargetUrlIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "target_url", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereTargetUrlNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "target_url", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByTargetUrl(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "target_url "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereOnwerId(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "onwer_id", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereOnwerIdIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "onwer_id", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereOnwerIdNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "onwer_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByOnwerId(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "onwer_id "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereGroupId(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "group_id", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereGroupIdIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "group_id", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereGroupIdNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "group_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByGroupId(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "group_id "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereGroupName(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "group_name", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereGroupNameIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "group_name", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereGroupNameNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "group_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByGroupName(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "group_name "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereCreateDate(p model.Predicate, value time.Time) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_date", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereCreateDateIn(value []time.Time) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_date", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereCreateDateNotIn(value []time.Time) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_date", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByCreateDate(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_date "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereUpdateDate(p model.Predicate, value time.Time) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_date", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereUpdateDateIn(value []time.Time) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_date", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereUpdateDateNotIn(value []time.Time) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_date", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByUpdateDate(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "update_date "+order)
	return qb
}

func (qb *apiRepoQueryBuilder) WhereAuthMode(p model.Predicate, value string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "auth_mode", p),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereAuthModeIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "auth_mode", "IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) WhereAuthModeNotIn(value []string) *apiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "auth_mode", "NOT IN"),
		value,
	})
	return qb
}

func (qb *apiRepoQueryBuilder) OrderByAuthMode(asc bool) *apiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "auth_mode "+order)
	return qb
}
