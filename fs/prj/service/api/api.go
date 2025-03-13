package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/liyonge-cm/go-api-cli-prj/config"
	"github.com/liyonge-cm/go-api-cli-prj/service/api/common"
	_ "github.com/liyonge-cm/go-api-cli-prj/service/api/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Http = &HttpService{}

type HttpService struct {
	ctx    context.Context
	config *config.Service
	Logger *zap.Logger
}

func NewHttp(ctx context.Context, config *config.Service) *HttpService {
	return &HttpService{
		ctx:    ctx,
		config: config,
	}
}

func (s *HttpService) WithLogger(log *zap.Logger) {
	s.Logger = log.With(zap.String("service", "http"))
}

func (s *HttpService) Start() {
	s.initRouter()
}

func (s *HttpService) Close() {
}

// 初始化路由
func (s *HttpService) initRouter() {
	// 设置为发布模式（初始化路由之前设置）
	gin.SetMode(gin.ReleaseMode)
	// gin 默认中间件
	r := gin.Default()

	// 访问一个错误网站时，返回404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404 ,page not exists!",
		})
	})

	for _, group := range common.Groups {
		gr := r.Group(group.Path)
		for _, router := range group.Routers {
			gr.POST(router.Path, func(c *gin.Context) {
				controller := common.NewRequest(c)
				controller.WithLogger(s.Logger)
				router.Function(controller)
			})
		}
	}

	port := s.config.Port
	s.Logger.Info(fmt.Sprintf("start service on: %d", port))

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
