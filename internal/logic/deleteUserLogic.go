package logic

import (
	"context"
	"fmt"
	"strings"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// DeleteUser 2. 删除用户 (逻辑删除)
func (l *DeleteUserLogic) DeleteUser(in *system.BatchDeleteUserReq) (*system.BatchDeleteUserResp, error) {
	// todo: add your logic here and delete this line
	ids := in.UserIds
	if len(ids) == 0 {
		return nil, status.Error(codes.Internal, "删除用户失败")
	}
	err := l.svcCtx.DB.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		placeholders := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
		args := make([]interface{}, len(ids))
		for i, id := range ids {
			args[i] = id
		}
		query := fmt.Sprintf("UPDATE user_info SET del_flag = '2' WHERE user_id IN (%s) AND del_flag = '0'", placeholders)
		_, err := session.ExecCtx(ctx, query, args...)
		if err != nil {
			return status.Error(codes.Internal, "删除用户失败")
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "删除用户失败")
	}
	return &system.BatchDeleteUserResp{}, nil
}
