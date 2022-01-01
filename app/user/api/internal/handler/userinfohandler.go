package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"mfzero-mall/app/user/api/internal/logic"
	"mfzero-mall/app/user/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
