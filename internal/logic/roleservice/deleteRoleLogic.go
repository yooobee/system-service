package roleservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 2. 批量删除角色 (逻辑删除)
func (l *DeleteRoleLogic) DeleteRole(in *system.BatchDeleteRoleReq) (*system.BatchDeleteRoleResp, error) {
	// todo: add your logic here and delete this line

	return &system.BatchDeleteRoleResp{}, nil
}
