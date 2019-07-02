package hander

import (
	"context"

	"github.com/micro/go-micro/metadata"

	"github.com/gomsa/socialite/client"
	userPB "github.com/gomsa/socialite/proto/user"
	"github.com/gomsa/tools/uitl"

	pb "github.com/gomsa/socialite-api/proto/user"
)

// User 用户结构
type User struct {
}

// List 用户列表
func (srv *User) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	query := &userPB.ListQuery{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	userRes, err := client.User.List(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Get 获取用户
func (srv *User) Get(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	userRes, err := client.User.Get(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	meta, _ := metadata.FromContext(ctx)
	user.Origin = meta["service"]
	userRes, err := client.User.Create(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Update 更新用户
func (srv *User) Update(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{}
	err = uitl.Data2Data(req, user)
	if err != nil {
		return err
	}
	userRes, err := client.User.Update(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}

// Delete 删除用户
func (srv *User) Delete(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user := &userPB.User{
		Id: req.Id,
	}
	if err != nil {
		return err
	}
	userRes, err := client.User.Delete(ctx, user)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(userRes, res)
	if err != nil {
		return err
	}
	return err
}
