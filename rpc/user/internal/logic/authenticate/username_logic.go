package authenticatelogic

import (
	"context"
	"errors"

	"sakura/aurora/rpc/user/internal/svc"
	"sakura/aurora/rpc/user/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsernameLogic {
	return &UsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsernameLogic) Username(in *types.UsernameAndPassword) (*types.Authenticated, error) {
	if in.Username != "leaf" || in.Password != "123456" {
		return nil, errors.New("登陆失败")
	}
	return &types.Authenticated{Token: "登陆成功"}, nil
}
