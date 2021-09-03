package db

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/liuwqiang/tykd/config"
	"github.com/liuwqiang/tykd/model"
	"github.com/liuwqiang/tykd/model/api_repo"
	"github.com/liuwqiang/tykd/model/app_api_relation_repo"
)

const (
	OPEN      = "open"
	APPID     = "appid"
	SIGNATURE = "signature"
)

func TestApiCreate(t *testing.T) {
	globalConf := config.Default
	config.SetGlobal(globalConf)
	repo, err := New()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	api := api_repo.NewModel()
	uuid, err := GetUUID()
	if err != nil {
		panic(err)
	}
	api.Id = uuid
	api.Name = "ping"
	api.Description = "test"
	api.GroupId = "test"
	api.GroupName = "test"
	api.OnwerId = "test"
	api.ListenPath = "/com.wanfangdata.tyk.Hello"
	api.TargetUrl = "h2c://tyk-control-plane.gateway:8088"
	api.AuthMode = OPEN

	err = api.Create(repo.GetDb().WithContext(ctx))
	if err != nil {
		panic(err)
	}
}

func TestApiDelete(t *testing.T) {
	globalConf := config.Default
	config.SetGlobal(globalConf)
	repo, err := New()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	qb := api_repo.NewQueryBuilder()
	id := "61318c49cb3e23ed3d9c"
	qb.WhereId(model.EqualPredicate, id)
	err = qb.Delete(repo.GetDb().WithContext(ctx))
	if err != nil {
		panic(err)
	}
}

func TestApiUpdate(t *testing.T) {
	globalConf := config.Default
	config.SetGlobal(globalConf)
	repo, err := New()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	data := map[string]interface{}{
		"name":        "done",
		"description": "lwq",
		"update_date": time.Now(),
	}
	qb := api_repo.NewQueryBuilder()
	id := "61318e17b05b35bf09e7"
	qb.WhereId(model.EqualPredicate, id)
	err = qb.Updates(repo.GetDb().WithContext(ctx), data)
	if err != nil {
		panic(err)
	}
}

func TestApiList(t *testing.T) {
	globalConf := config.Default
	config.SetGlobal(globalConf)
	ctx := context.Background()
	repo, err := New()
	if err != nil {
		panic(err)
	}
	qb := api_repo.NewQueryBuilder()
	listdata, err := qb.QueryAll(repo.GetDb().WithContext(ctx))
	if err != nil {
		panic(err)
	}
	for _, v := range listdata {
		fmt.Printf("%s-%s-%s\n", v.Id, v.Name, v.Description)
	}
}

type Test struct {
	Name string
	Id   string
	Age  int
}

func TestApiAppRelationCreate(t *testing.T) {
	globalConf := config.Default
	config.SetGlobal(globalConf)
	ctx := context.Background()
	repo, err := New()
	if err != nil {
		panic(err)
	}
	relation := app_api_relation_repo.NewModel()
	relation.ApiId = "test"
	relation.AppId = "test"
	relation.ExtraData, err = json.Marshal(Test{"lwq", "lwq", 20})
	relation.Rate = 10
	relation.Per = 10
	relation.Quota = 10
	relation.QuotaRenewalRate = 10
	relation.ExpiredDate = time.Now()
	if err != nil {
		panic(err)
	}
	err = relation.Create(repo.GetDb().WithContext(ctx))
	if err != nil {
		panic(err)
	}
}

func TestApiAppRelationList(t *testing.T) {
	globalConf := config.Default
	config.SetGlobal(globalConf)
	ctx := context.Background()
	repo, err := New()
	if err != nil {
		panic(err)
	}
	qb := app_api_relation_repo.NewQueryBuilder()
	listdata, err := qb.QueryAll(repo.GetDb().WithContext(ctx))
	if err != nil {
		panic(err)
	}
	for _, v := range listdata {
		extra := v.ExtraData
		test := Test{}
		json.Unmarshal(extra, &test)
		fmt.Printf("%s-%s-%d-{%s-%s-%d}\n", v.AppId, v.ApiId, v.Quota, test.Name, test.Id, test.Age)
	}
}
