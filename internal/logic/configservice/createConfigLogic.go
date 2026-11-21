package configservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateConfigLogic {
	return &CreateConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 1. 新增参数
func (l *CreateConfigLogic) CreateConfig(in *system.CreateConfigReq) (*system.CreateConfigResp, error) {
	// todo: add your logic here and delete this line

	return &system.CreateConfigResp{}, nil
}
