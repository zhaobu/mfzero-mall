package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"mfzero-mall/app/order/api/internal/logic"
	"mfzero-mall/app/order/api/internal/svc"
	"mfzero-mall/app/order/api/internal/types"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
