package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-school/ent/predicate"
	"github.com/suyuan32/simple-admin-school/ent/teacher"
	"github.com/suyuan32/simple-admin-school/internal/svc"
	"github.com/suyuan32/simple-admin-school/internal/types"
	"github.com/suyuan32/simple-admin-school/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherListLogic {
	return &GetTeacherListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherListLogic) GetTeacherList(req *types.TeacherListReq) (*types.TeacherListResp, error) {
	var predicates []predicate.Teacher
	if req.Name != nil {
		predicates = append(predicates, teacher.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.Teacher.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.TeacherListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.TeacherInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.Unix()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.Unix()),
				},
				Name: &v.Name,
				Age:  &v.Age,
			})
	}

	return resp, nil
}
