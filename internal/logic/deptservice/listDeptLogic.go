package deptservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDeptLogic {
	return &ListDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListDept 5. 查询部门列表
func (l *ListDeptLogic) ListDept(in *system.ListDeptReq) (*system.ListDeptResp, error) {
	// todo: add your logic here and delete this line

	return &system.ListDeptResp{}, nil
}
