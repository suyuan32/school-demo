package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-school/internal/svc"
	"github.com/suyuan32/simple-admin-school/internal/types"
	"github.com/suyuan32/simple-admin-school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTeacherLogic {
	return &CreateTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTeacherLogic) CreateTeacher(req *types.TeacherInfo) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.Teacher.Create().
		SetNotNilName(req.Name).
		SetNotNilAge(req.Age).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
