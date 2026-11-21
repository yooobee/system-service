package configservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListConfigLogic {
	return &ListConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 5. 查询参数列表 (分页)
func (l *ListConfigLogic) ListConfig(in *system.ListConfigReq) (*system.ListConfigResp, error) {
	// todo: add your logic here and delete this line

	return &system.ListConfigResp{}, nil
}
