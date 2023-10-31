package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-school/internal/svc"
	"github.com/suyuan32/simple-admin-school/internal/types"
	"github.com/suyuan32/simple-admin-school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherByIdLogic {
	return &GetTeacherByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherByIdLogic) GetTeacherById(req *types.IDReq) (*types.TeacherInfoResp, error) {
	data, err := l.svcCtx.DB.Teacher.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.TeacherInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.TeacherInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.Unix()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.Unix()),
			},
			Name: &data.Name,
			Age:  &data.Age,
		},
	}, nil
}
