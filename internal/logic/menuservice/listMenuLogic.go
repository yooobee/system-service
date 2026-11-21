package menuservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMenuLogic {
	return &ListMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListMenu 5. 查询菜单列表
func (l *ListMenuLogic) ListMenu(in *system.ListMenuReq) (*system.ListMenuResp, error) {
	// todo: add your logic here and delete this line

	return &system.ListMenuResp{}, nil
}
