package blog

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/blog"
)

type SearchOneData struct {
	AID    int32  // BLOGID
	Author string // 作者
	Title  string // 标题
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *blog.Blog, err error) {

	qb := blog.NewQueryBuilder()

	if searchOneData.AID != 0 {
		qb.WhereId(mysql.EqualPredicate, searchOneData.AID)
	}

	if searchOneData.Author != "" {
		qb.WhereUsername(mysql.EqualPredicate, searchOneData.Author)
	}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
