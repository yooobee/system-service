package menuservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMenu 4. 根据ID查询菜单
func (l *GetMenuLogic) GetMenu(in *system.GetMenuReq) (*system.MenuInfoResp, error) {
	// todo: add your logic here and delete this line

	return &system.MenuInfoResp{}, nil
}
