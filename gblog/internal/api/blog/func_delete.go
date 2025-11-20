package blog

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 删除Blog
// @Summary 删除Blog
// @Description 删除Blog
// @Tags API.Blog
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/delete [post]
func (h *handler) Delete() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
