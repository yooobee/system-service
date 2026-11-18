package logic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 2. 删除用户 (逻辑删除)
func (l *DeleteUserLogic) DeleteUser(in *system.BatchDeleteUserReq) (*system.BatchDeleteUserResp, error) {
	// todo: add your logic here and delete this line

	return &system.BatchDeleteUserResp{}, nil
}
