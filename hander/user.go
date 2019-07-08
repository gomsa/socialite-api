package hander

import (
	"fmt"
	"errors"
	"context"

	"github.com/gomsa/socialite/client"
	userPB "github.com/gomsa/socialite/proto/user"
	"github.com/gomsa/tools/uitl"
	"github.com/micro/go-micro/metadata"

	pb "github.com/gomsa/socialite-api/proto/user"
)

// User 用户结构
type User struct {
}

// Build 获取用户绑定登录
func (srv *User) Build(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	// meta["user_id"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	fmt.Println(meta)
	if userID, ok := meta["user_id"]; ok {
		user := &userPB.User{
			Id: userID,
		}
		userRes, err := client.User.Build(ctx, user)
		if err != nil {
			return err
		}
		for _, user := range userRes.Users {
			res.Build = append(res.Build, user.Origin)
		}
		if err != nil {
			return err
		}
	} else {
		err = errors.New("获取用户绑定账号失败,未找到用户ID")
		return err
	}
	return err
}

// Delete 删除用户
func (srv *User) Delete(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	// meta["user_id"] 通过 meta 获取用户 id --- So this function needs token to use
	meta, _ := metadata.FromContext(ctx)
	if userID, ok := meta["user_id"]; ok {
		user := &userPB.User{
			Id: userID,
		}
		userRes, err := client.User.Delete(ctx, user)
		if err != nil {
			return err
		}
		err = uitl.Data2Data(userRes, res)
		if err != nil {
			return err
		}
	} else {
		err = errors.New("删除用户绑定账号失败,未找到用户ID")
		return err
	}
	return err
}
