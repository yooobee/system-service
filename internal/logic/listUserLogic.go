package logic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 5. 查询用户列表 (分页)
func (l *ListUserLogic) ListUser(in *system.ListUserReq) (*system.ListUserResp, error) {
	// todo: add your logic here and delete this line

	return &system.ListUserResp{}, nil
}
