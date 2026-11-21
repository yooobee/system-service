package deptservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeptLogic {
	return &DeleteDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteDept 2. 批量删除部门 (逻辑删除)
func (l *DeleteDeptLogic) DeleteDept(in *system.BatchDeleteDeptReq) (*system.BatchDeleteDeptResp, error) {
	// todo: add your logic here and delete this line

	return &system.BatchDeleteDeptResp{}, nil
}
