package userservicelogic

import (
	"context"
	"database/sql"
	"time"

	"system-service/internal/svc"
	"system-service/model"
	"system-service/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUser 4. 根据ID查询用户
func (l *GetUserLogic) GetUser(in *system.GetUserReq) (*system.UserInfoResp, error) {
	if in.UserId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "用户ID不能为空")
	}

	user, err := model.NewUserInfoModel(l.svcCtx.DB).FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		l.Errorf("查询用户失败, userId: %d, err: %v", in.UserId, err)
		return nil, status.Error(codes.Internal, "查询用户失败")
	}
	// 不返回 password 等敏感字段
	return &system.UserInfoResp{
		UserId:      user.UserId,
		DeptId:      user.DeptId.Int64,
		UserName:    user.UserName,
		NickName:    user.NickName,
		UserType:    user.UserType,
		Email:       user.Email,
		PhoneNumber: user.Phonenumber,
		Sex:         user.Sex,
		Avatar:      user.Avatar,
		Status:      user.Status,
		DelFlag:     user.DelFlag,
		LoginIp:     user.LoginIp,
		LoginDate:   nullTimeToString(user.LoginDate),
		CreateTime:  nullTimeToString(user.CreateTime),
		UpdateTime:  nullTimeToString(user.UpdateTime),
		CreateBy:    user.CreateBy,
		UpdateBy:    user.UpdateBy,
		Remark:      user.Remark.String,
	}, nil
}

// nullTimeToString 把可空的数据库时间转成 proto 里的字符串格式
func nullTimeToString(t sql.NullTime) string {
	if !t.Valid {
		return ""
	}
	return t.Time.Format(time.DateTime)
}
