package server

import (
	"context"

	menuservicelogic "system-service/internal/logic/menuservice"
	"system-service/internal/svc"
	"system-service/rpc/system"
)

type MenuServiceServer struct {
	svcCtx *svc.ServiceContext
	system.UnimplementedMenuServiceServer
}

func NewMenuServiceServer(svcCtx *svc.ServiceContext) *MenuServiceServer {
	return &MenuServiceServer{
		svcCtx: svcCtx,
	}
}

// 1. 新增菜单
func (s *MenuServiceServer) CreateMenu(ctx context.Context, in *system.CreateMenuReq) (*system.CreateMenuResp, error) {
	l := menuservicelogic.NewCreateMenuLogic(ctx, s.svcCtx)
	return l.CreateMenu(in)
}

// 2. 批量删除菜单 (逻辑删除)
func (s *MenuServiceServer) DeleteMenu(ctx context.Context, in *system.BatchDeleteMenuReq) (*system.BatchDeleteMenuResp, error) {
	l := menuservicelogic.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

// 3. 修改菜单
func (s *MenuServiceServer) UpdateMenu(ctx context.Context, in *system.UpdateMenuReq) (*system.UpdateMenuResp, error) {
	l := menuservicelogic.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

// 4. 根据ID查询菜单
func (s *MenuServiceServer) GetMenu(ctx context.Context, in *system.GetMenuReq) (*system.MenuInfoResp, error) {
	l := menuservicelogic.NewGetMenuLogic(ctx, s.svcCtx)
	return l.GetMenu(in)
}

// 5. 查询菜单列表
func (s *MenuServiceServer) ListMenu(ctx context.Context, in *system.ListMenuReq) (*system.ListMenuResp, error) {
	l := menuservicelogic.NewListMenuLogic(ctx, s.svcCtx)
	return l.ListMenu(in)
}
