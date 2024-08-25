package user

import (
	"net/http"
	"sakura/aurora/api/errors"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sakura/aurora/api/internal/logic/user"
	"sakura/aurora/api/internal/svc"
	"sakura/aurora/api/internal/types"
)

func AuthenticateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthenticateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, errors.WrapStatus(errors.ParamParseErrorStatus))
			return
		}
		users := user.NewAuthenticateLogic(r.Context(), ctx)
		if resp, err := users.Authenticate(&req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, errors.WrapStatus(err))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
