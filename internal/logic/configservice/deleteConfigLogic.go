package configservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigLogic {
	return &DeleteConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 2. 批量删除参数
func (l *DeleteConfigLogic) DeleteConfig(in *system.BatchDeleteConfigReq) (*system.BatchDeleteConfigResp, error) {
	// todo: add your logic here and delete this line

	return &system.BatchDeleteConfigResp{}, nil
}
