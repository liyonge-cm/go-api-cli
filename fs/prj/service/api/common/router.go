package common

type Router struct {
	Path     string
	Function func(*Controller)
}

type Group struct {
	Path    string
	Routers []*Router
}

var Groups []*Group

func NewGroup(path string) *Group {
	return &Group{
		Path:    path,
		Routers: []*Router{},
	}
}

func (g *Group) NewRouter(path string, function func(*Controller)) {
	g.Routers = append(g.Routers, &Router{
		Path:     path,
		Function: function,
	})
}

func (g *Group) Set() {
	Groups = append(Groups, g)
}

// func (g *Groups) Set() {
// 	fmt.Println("set", g.Router)
// 	Http.routers = append(Http.routers, g)
// }

// type Api interface {
// 	Response()
// 	// GetUsers()
// 	// GetUser()
// 	// UpdateUser()
// 	// DeleteUser()
// 	//ChatCompletions()

// }

// func newApi(c *gin.Context, log *zap.Logger) Api {
// 	return common.NewRequest(c, log)
// }

// func RegisterRouter(r *gin.Engine, log *zap.Logger) {
// 	r.POST("demo", func(c *gin.Context) {
// 		common.NewRequest(c, log).CreateUser()
// 	})
// }
