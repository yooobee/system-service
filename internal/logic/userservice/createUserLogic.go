package userservicelogic

import (
	"context"
	"time"

	"system-service/internal/svc"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUser 1. 新增用户
func (l *CreateUserLogic) CreateUser(in *system.CreateUserReq) (*system.CreateUserResp, error) {
	// 1.参数校验
	if in.UserName == "" || in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "用户名和密码不能为空")
	}
	// 2.处理 optional 字段 （Avatar）
	var avatar string
	if in.Avatar != nil {
		avatar = *in.Avatar
	}
	// 3.密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		logx.Errorf("密码加密失败：%v", err)
		return nil, status.Error(codes.Internal, "新增用户失败")
	}
	// 4.准备 SQL 语句
	query := `INSERT INTO user_info (dept_id, user_name, nick_name, user_type, email, phonenumber, sex, avatar, password, status, remark, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := l.svcCtx.DB.ExecCtx(l.ctx, query, in.DeptId, in.UserName, in.NickName, in.UserType, in.Email, in.PhoneNumber, in.Sex,
		avatar, password, in.Status, in.Remark, time.Now(), time.Now())
	if err != nil {
		logx.Errorf("插入用户失败：%v", err)
		return nil, status.Error(codes.Internal, "新增用户失败")
	}
	// 6. 获取自增 ID
	_, err = result.LastInsertId()
	if err != nil {
		logx.Errorf("获取自增ID失败：%v", err)
		return nil, status.Error(codes.Internal, "新增用户失败")
	}
	return &system.CreateUserResp{}, nil
}
