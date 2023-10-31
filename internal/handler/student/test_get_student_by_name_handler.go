package student

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-school/internal/logic/student"
	"github.com/suyuan32/simple-admin-school/internal/svc"
	"github.com/suyuan32/simple-admin-school/internal/types"
)

// swagger:route post /student/test student TestGetStudentByName
//
// Test get student by name | 测试通过名称获取单个学生
//
// Test get student by name | 测试通过名称获取单个学生
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TestReq
//
// Responses:
//  200: TestResp

func TestGetStudentByNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := student.NewTestGetStudentByNameLogic(r.Context(), svcCtx)
		resp, err := l.TestGetStudentByName(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
