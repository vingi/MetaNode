package blog

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/blog"
)

type createRequest struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	Author  string `json:"author" form:"author" binding:"required"`
}

type createResponse struct {
	Id int32 `json:"id"`
}

// Create 创建Blog
// @Summary 创建Blog
// @Description 创建Blog
// @Tags API.Blog
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		createData := new(blog.CreateBlog)
		createData.Title = req.Title
		createData.Content = req.Content
		createData.Author = req.Author

		id, err := h.blogService.Create(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.BlogCreateError,
				code.Text(code.BlogCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
