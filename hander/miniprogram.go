package hander

import (
	"context"

	pb "github.com/gomsa/socialite-api/proto/miniprogram"
	"github.com/gomsa/socialite/client"
	srvPB "github.com/gomsa/socialite/proto/miniprogram"
	"github.com/gomsa/tools/uitl"
)

// Miniprogram 小程序
type Miniprogram struct {
}

// Auth 小程序登录授权
func (srv *Miniprogram) Auth(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	query := &srvPB.Request{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	mpRes, err := client.Miniprogram.Auth(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(mpRes, res)
	if err != nil {
		return err
	}
	return err
}
