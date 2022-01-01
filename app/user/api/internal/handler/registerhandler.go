package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"mfzero-mall/app/user/api/internal/logic"
	"mfzero-mall/app/user/api/internal/svc"
	"mfzero-mall/app/user/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
