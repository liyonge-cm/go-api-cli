package user

import (
	"github.com/liyonge-cm/go-api-cli-prj/service/api/common"
)

func init() {
	group := common.NewGroup("user")
	group.NewRouter("/create", CreateUser)
	group.NewRouter("/getList", GetUsers)
	group.NewRouter("/get", GetUser)
	group.NewRouter("/update", UpdateUser)
	group.NewRouter("/delete", DeleteUser)
	group.Set()
}
