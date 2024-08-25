package user

import (
	"context"
	"sakura/aurora/api/errors"
	userTypes "sakura/aurora/rpc/user/types"

	"sakura/aurora/api/internal/svc"
	"sakura/aurora/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateLogic struct {
	logx.Logger
	ctx     context.Context
	service *svc.ServiceContext
}

func NewAuthenticateLogic(ctx context.Context, service *svc.ServiceContext) *AuthenticateLogic {
	return &AuthenticateLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		service: service,
	}
}

func (auth *AuthenticateLogic) Authenticate(req *types.AuthenticateRequest) (resp *types.AuthenticatedResponse, err error) {
	if username, err := auth.service.Authenticate.Username(auth.ctx, &userTypes.UsernameAndPassword{Password: req.Password, Username: req.Username}); err == nil {
		return &types.AuthenticatedResponse{
			Token: username.Token,
		}, nil
	}
	return nil, errors.UsernamePasswordNotMatchStatus
}
