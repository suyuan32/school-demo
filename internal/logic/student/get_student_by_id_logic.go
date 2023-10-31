package student

import (
	"context"

	"github.com/suyuan32/simple-admin-school/internal/svc"
	"github.com/suyuan32/simple-admin-school/internal/types"
	"github.com/suyuan32/simple-admin-school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentByIdLogic {
	return &GetStudentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentByIdLogic) GetStudentById(req *types.UUIDReq) (*types.StudentInfoResp, error) {
	data, err := l.svcCtx.DB.Student.Get(l.ctx, uuidx.ParseUUIDString(req.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.StudentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.StudentInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(data.ID.String()),
				CreatedAt: pointy.GetPointer(data.CreatedAt.Unix()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.Unix()),
			},
			Name:    &data.Name,
			Age:     &data.Age,
			Address: &data.Address,
		},
	}, nil
}
