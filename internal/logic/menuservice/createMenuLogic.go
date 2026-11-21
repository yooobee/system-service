package menuservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateMenu 1. 新增菜单
func (l *CreateMenuLogic) CreateMenu(in *system.CreateMenuReq) (*system.CreateMenuResp, error) {
	// todo: add your logic here and delete this line

	return &system.CreateMenuResp{}, nil
}
