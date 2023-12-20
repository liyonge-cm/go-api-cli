package user

import "go-cli-prj/router"

func New() *router.Group {
	group := router.NewGroup("user")
	group.NewRouter("/create", CreateUser)
	group.NewRouter("/getList", GetUsers)
	group.NewRouter("/get", GetUser)
	group.NewRouter("/update", UpdateUser)
	group.NewRouter("/delete", DeleteUser)
	return group
}
