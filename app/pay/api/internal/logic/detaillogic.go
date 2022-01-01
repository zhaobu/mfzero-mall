package logic

import (
	"context"

	"mfzero-mall/app/pay/api/internal/svc"
	"mfzero-mall/app/pay/api/internal/types"
	"mfzero-mall/app/pay/rpc/pay"

	"github.com/tal-tech/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) DetailLogic {
	return DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req types.DetailRequest) (resp *types.DetailResponse, err error) {
	res, err := l.svcCtx.PayRpc.Detail(l.ctx, &pay.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     req.Id,
		Uid:    res.Uid,
		Oid:    res.Oid,
		Amount: res.Amount,
		Source: res.Source,
		Status: res.Status,
	}, nil
}