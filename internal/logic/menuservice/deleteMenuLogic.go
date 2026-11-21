package menuservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMenu 2. 批量删除菜单 (逻辑删除)
func (l *DeleteMenuLogic) DeleteMenu(in *system.BatchDeleteMenuReq) (*system.BatchDeleteMenuResp, error) {
	// todo: add your logic here and delete this line

	return &system.BatchDeleteMenuResp{}, nil
}
