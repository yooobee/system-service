package server

import (
	"context"

	deptservicelogic "system-service/internal/logic/deptservice"
	"system-service/internal/svc"
	"system-service/rpc/system"
)

type DeptServiceServer struct {
	svcCtx *svc.ServiceContext
	system.UnimplementedDeptServiceServer
}

func NewDeptServiceServer(svcCtx *svc.ServiceContext) *DeptServiceServer {
	return &DeptServiceServer{
		svcCtx: svcCtx,
	}
}

// 1. 新增部门
func (s *DeptServiceServer) CreateDept(ctx context.Context, in *system.CreateDeptReq) (*system.CreateDeptResp, error) {
	l := deptservicelogic.NewCreateDeptLogic(ctx, s.svcCtx)
	return l.CreateDept(in)
}

// 2. 批量删除部门 (逻辑删除)
func (s *DeptServiceServer) DeleteDept(ctx context.Context, in *system.BatchDeleteDeptReq) (*system.BatchDeleteDeptResp, error) {
	l := deptservicelogic.NewDeleteDeptLogic(ctx, s.svcCtx)
	return l.DeleteDept(in)
}

// 3. 修改部门
func (s *DeptServiceServer) UpdateDept(ctx context.Context, in *system.UpdateDeptReq) (*system.UpdateDeptResp, error) {
	l := deptservicelogic.NewUpdateDeptLogic(ctx, s.svcCtx)
	return l.UpdateDept(in)
}

// 4. 根据ID查询部门
func (s *DeptServiceServer) GetDept(ctx context.Context, in *system.GetDeptReq) (*system.DeptInfoResp, error) {
	l := deptservicelogic.NewGetDeptLogic(ctx, s.svcCtx)
	return l.GetDept(in)
}

// 5. 查询部门列表
func (s *DeptServiceServer) ListDept(ctx context.Context, in *system.ListDeptReq) (*system.ListDeptResp, error) {
	l := deptservicelogic.NewListDeptLogic(ctx, s.svcCtx)
	return l.ListDept(in)
}
