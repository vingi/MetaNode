package blog

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/blog"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建Blog
	// @Tags API.Blog
	// @Router /api/order/create [post]
	Create() core.HandlerFunc

	// Update 更新Blogs
	// @Tags API.Blog
	// @Router /api/order/update [post]
	Update() core.HandlerFunc

	// Delete 删除Blog
	// @Tags API.Blog
	// @Router /api/order/delete [post]
	Delete() core.HandlerFunc

	// Detail Blog
	// @Tags API.blog
	// @Router /api/order/{id} [get]
	Detail() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	db          mysql.Repo
	cache       redis.Repo
	hashids     hash.Hash
	blogService blog.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:  logger,
		db:      db,
		cache:   cache,
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
