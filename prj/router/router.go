package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var ApiGroups = []*Group{}

type Router struct {
	Router   string
	Function func(c *gin.Context)
}

type Group struct {
	Router  string
	Routers []*Router
}

func NewGroup(groupRouter string) *Group {
	return &Group{
		Router:  groupRouter,
		Routers: []*Router{},
	}
}

func (g *Group) NewRouter(router string, function func(c *gin.Context)) {
	g.Routers = append(g.Routers, &Router{
		Router:   router,
		Function: function,
	})
}

func (g *Group) Register() {
	ApiGroups = append(ApiGroups, g)
}

func Init() {
	// 设置为发布模式（初始化路由之前设置）
	gin.SetMode(gin.ReleaseMode)
	// gin 默认中间件
	r := gin.Default()

	// 访问一个错误路由时，返回404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "404, page not exists!",
		})
	})

	// 注册路由
	for _, group := range ApiGroups {
		gr := r.Group(group.Router)
		for _, router := range group.Routers {
			gr.POST(router.Router, router.Function)
		}
	}

	// 启动API服务
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
