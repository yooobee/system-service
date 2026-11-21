package deptservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeptLogic {
	return &CreateDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateDept 1. 新增部门
func (l *CreateDeptLogic) CreateDept(in *system.CreateDeptReq) (*system.CreateDeptResp, error) {
	// todo: add your logic here and delete this line

	return &system.CreateDeptResp{}, nil
}
