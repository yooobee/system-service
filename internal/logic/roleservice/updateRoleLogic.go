package roleservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 3. 修改角色
func (l *UpdateRoleLogic) UpdateRole(in *system.UpdateRoleReq) (*system.UpdateRoleResp, error) {
	// todo: add your logic here and delete this line

	return &system.UpdateRoleResp{}, nil
}
