package user

import "go-api-cli-prj/router"

func init() {
	group := router.NewGroup("user")
	group.NewRouter("/create", CreateUser)
	group.NewRouter("/getList", GetUsers)
	group.NewRouter("/get", GetUser)
	group.NewRouter("/update", UpdateUser)
	group.NewRouter("/delete", DeleteUser)
	group.Register()
}
