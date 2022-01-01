package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"mfzero-mall/app/product/api/internal/logic"
	"mfzero-mall/app/product/api/internal/svc"
	"mfzero-mall/app/product/api/internal/types"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
