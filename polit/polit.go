package polit

import (
	"context"

	"github.com/liuwqiang/tykd/db"
	"github.com/liuwqiang/tykd/model/api_repo"
	pb "github.com/liuwqiang/tykd/protos/golang"
)

func (t *T) CreateApi(ctx context.Context, req *pb.Api) (*pb.ApiDefinition, error) {
	api := api_repo.NewModel()
	uuid, err := db.GetUUID()
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

	err = api.Create(t.db.GetDb().WithContext(ctx))
	return nil, err
}
func (t *T) DeleteApi(ctx context.Context, id *pb.ID) (*pb.ApiDefinition, error) {
	return nil, nil
}
func (t *T) UpdateApi(ctx context.Context, api *pb.ApiDefinition) (*pb.ID, error) {
	return nil, nil
}
func (t *T) CreateApp(ctx context.Context, app *pb.App) (*pb.AppDefinition, error) {
	return nil, nil
}
func (t *T) DeleteApp(ctx context.Context, id *pb.ID) (*pb.AppDefinition, error) {
	return nil, nil
}
func (t *T) UpdateApp(ctx context.Context, app *pb.AppDefinition) (*pb.ID, error) {
	return nil, nil
}
func (t *T) FlushAppSecret(ctx context.Context, id *pb.ID) (*pb.AppDefinition, error) {
	return nil, nil
}
