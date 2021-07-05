package server

import (
	"aurora/blog/api/app/service/manager/internal/service"
	"aurora/blog/api/pkg/config"
	"aurora/blog/api/pkg/middleware"
	"aurora/blog/api/pkg/transport"
	"aurora/blog/api/pkg/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// ProvideSet for server package ...
var ProvideSet = wire.NewSet(NewHttpServer, NewHttpRouter)

func NewHttpServer(engine *gin.Engine, logger *zap.Logger, config config.Server) transport.Server {
	return http.NewServer(engine, logger, config)
}

// call chain: data -> biz -> service -> server
// NewHttpRouter set gin.Engine as http.Handler
func NewHttpRouter(
	account *service.AccountService,
	article *service.ArticleService,
	category *service.CategoryService,
	site *service.SiteService,
	tag *service.TagService,
	theme *service.ThemeService,
	user *service.UserService,
	logger *zap.Logger,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(middleware.Ginzap(logger))
	engine.Use(middleware.RecoveryWithZap(logger, true))
	// set router
	mng := engine.Group("/v1/mng")

	RegisterAccountRouter(mng, account)
	RegisterArticleRouter(mng, article)
	RegisterCategoryRouter(mng, category)
	RegisterSiteRouter(mng, site)
	RegisterTagRouter(mng, tag)
	RegisterThemeRouter(mng, theme)
	RegisterUserRouter(mng, user)

	return engine
}
