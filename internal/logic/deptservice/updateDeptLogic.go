package deptservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptLogic {
	return &UpdateDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDept 3. 修改部门
func (l *UpdateDeptLogic) UpdateDept(in *system.UpdateDeptReq) (*system.UpdateDeptResp, error) {
	// todo: add your logic here and delete this line

	return &system.UpdateDeptResp{}, nil
}
