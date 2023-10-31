package student

import (
	"context"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-school/ent/student"
	"github.com/suyuan32/simple-admin-school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-school/internal/svc"
	"github.com/suyuan32/simple-admin-school/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestGetStudentByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestGetStudentByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestGetStudentByNameLogic {
	return &TestGetStudentByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *TestGetStudentByNameLogic) TestGetStudentByName(req *types.TestReq) (resp *types.TestResp, err error) {
	data, err := l.svcCtx.DB.Student.Query().Where(student.NameEQ(*req.Name)).First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.TestResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
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
