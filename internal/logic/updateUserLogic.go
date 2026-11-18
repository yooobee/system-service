package logic

import (
	"context"
	"strings"
	"time"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUser 3. 修改用户
func (l *UpdateUserLogic) UpdateUser(in *system.UpdateUserReq) (*system.UpdateUserResp, error) {
	// 添加入口日志
	l.Logger.Infof("更新用户信息前信息，User: %v", in)
	if in.UserId <= 0 {
		l.Logger.Errorf("更新用户失败：用户ID为空")
		return nil, status.Error(codes.InvalidArgument, "用户ID不能为空")
	}
	// 构建set子句
	setClauses, baseArgs, err := buildUpdateUser(in)
	if err != nil {
		// 记录构建失败日志
		l.Logger.Errorf("构建更新字段失败, userId: %d, err: %v", in.UserId, err)
		return nil, err
	}
	// 执行事务
	err = l.svcCtx.DB.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		updateQuery := "UPDATE user_info SET " + strings.Join(setClauses, ", ") + " WHERE user_id = ?"
		args := append(baseArgs, in.UserId)
		if _, err := session.ExecCtx(ctx, updateQuery, args...); err != nil {
			// 详细错误只记录在服务端日志，避免泄露给调用方
			logx.WithContext(ctx).Errorf("更新用户失败，userId: %v, err: %v", in.UserId, err)
			return status.Error(codes.Internal, "更新用户失败")
		}
		// 2. 更新角色（如果有）
		if len(in.Roles) > 0 {
			if err := updateUserRolesInTx(ctx, session, in.UserId, in.Roles); err != nil {
				// 角色更新失败已经在子函数里记录了日志，此处直接返回错误触发回滚
				return err
			}
		}
		return nil // 提交事务
	})
	if err != nil {
		return nil, err
	}

	return &system.UpdateUserResp{}, nil
}

// updateUserRolesInTx 在事务中更新用户角色
func updateUserRolesInTx(ctx context.Context, session sqlx.Session, userId int64, roles []int64) error {
	// 删除旧角色
	if _, err := session.ExecCtx(ctx, "DELETE FROM sys_user_role WHERE user_id = ?", userId); err != nil {
		logx.WithContext(ctx).Errorf("删除用户角色失败, userId: %d, err: %v", userId, err)
		return status.Error(codes.Internal, "更新用户角色失败")
	}

	// 批量插入新角色
	if len(roles) == 0 {
		return nil
	}
	valuePlaceholders := make([]string, 0, len(roles))
	args := make([]interface{}, 0, len(roles)*2)
	for _, roleId := range roles {
		valuePlaceholders = append(valuePlaceholders, "(?, ?)")
		args = append(args, userId, roleId)
	}
	insertQuery := "INSERT INTO sys_user_role (user_id, role_id) VALUES " + strings.Join(valuePlaceholders, ", ")
	if _, err := session.ExecCtx(ctx, insertQuery, args...); err != nil {
		logx.WithContext(ctx).Errorf("批量插入用户角色失败, userId: %d, roles: %v, err: %v", userId, roles, err)
		return status.Error(codes.Internal, "更新用户角色失败")
	}
	return nil
}
func buildUpdateUser(in *system.UpdateUserReq) ([]string, []interface{}, error) {
	var setClauses []string
	var args []interface{}
	if in.DeptId != 0 {
		setClauses = append(setClauses, "dept_id=?")
		args = append(args, in.DeptId)
	}
	if in.NickName != "" {
		setClauses = append(setClauses, "nick_name=?")
		args = append(args, in.NickName)
	}
	if in.UserType != "" {
		setClauses = append(setClauses, "user_type=?")
		args = append(args, in.UserType)
	}
	if in.Email != "" {
		setClauses = append(setClauses, "email=?")
		args = append(args, in.Email)
	}
	if in.PhoneNumber != "" {
		setClauses = append(setClauses, "phonenumber=?")
		args = append(args, in.PhoneNumber)
	}
	if in.Sex != "" {
		setClauses = append(setClauses, "sex=?")
		args = append(args, in.Sex)
	}
	if in.Status != "" {
		setClauses = append(setClauses, "status=?")
		args = append(args, in.Status)
	}
	if len(setClauses) == 0 {
		return nil, nil, status.Error(codes.InvalidArgument, "没有需要更新的字段")
	}
	// 添加统一的更新时间
	setClauses = append(setClauses, "update_time=?")
	args = append(args, time.Now())
	return setClauses, args, nil
}
