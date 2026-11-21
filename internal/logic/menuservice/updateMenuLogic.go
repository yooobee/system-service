package menuservicelogic

import (
	"context"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMenu 3. 修改菜单
func (l *UpdateMenuLogic) UpdateMenu(in *system.UpdateMenuReq) (*system.UpdateMenuResp, error) {
	// todo: add your logic here and delete this line

	return &system.UpdateMenuResp{}, nil
}
