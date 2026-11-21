package roleservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 5. 查询角色列表 (分页)
func (l *ListRoleLogic) ListRole(in *system.ListRoleReq) (*system.ListRoleResp, error) {
	// todo: add your logic here and delete this line

	return &system.ListRoleResp{}, nil
}
