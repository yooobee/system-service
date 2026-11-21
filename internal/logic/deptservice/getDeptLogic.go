package deptservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeptLogic {
	return &GetDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetDept 4. 根据ID查询部门
func (l *GetDeptLogic) GetDept(in *system.GetDeptReq) (*system.DeptInfoResp, error) {
	// todo: add your logic here and delete this line

	return &system.DeptInfoResp{}, nil
}
