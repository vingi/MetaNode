package blog

import (
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/blog"
)

type CreateBlog struct {
	Author   string // 作者
	Category string // 分类
	Title    string // 标题
	Content  string // 内容
}

func (s *service) Create(ctx core.Context, blogData *CreateBlog) (id int32, err error) {
	model := blog.NewModel()
	model.Author = blogData.Author
	model.Category = blogData.Category
	model.Title = blogData.Title
	model.Content = blogData.Content
	model.AddTime = time.Now()
	model.UpdateTime = time.Now()
	model.Status = 1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
