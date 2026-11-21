package configservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigLogic {
	return &GetConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 4. 根据ID查询参数
func (l *GetConfigLogic) GetConfig(in *system.GetConfigReq) (*system.ConfigInfoResp, error) {
	// todo: add your logic here and delete this line

	return &system.ConfigInfoResp{}, nil
}
