package logic

import (
	"context"

	"web3/internal/svc"
	"web3/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Web3Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWeb3Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Web3Logic {
	return &Web3Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Web3Logic) Web3(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
