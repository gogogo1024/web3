package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"web3/internal/logic"
	"web3/internal/svc"
	"web3/internal/types"
)

func Web3Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewWeb3Logic(r.Context(), svcCtx)
		resp, err := l.Web3(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
