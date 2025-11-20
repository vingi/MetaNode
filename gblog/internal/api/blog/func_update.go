package blog

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 更新Blogs
// @Summary 更新Blogs
// @Description 更新Blogs
// @Tags API.Blog
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} updateResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/update [post]
func (h *handler) Update() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
