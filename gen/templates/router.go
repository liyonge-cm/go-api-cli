package templates

const routerContent = `package $package_name

import "$prj_name/service/api/common"

func init() {
	group := common.NewGroup("$group_router")
	group.NewRouter("/create", $create_func_name)
	group.NewRouter("/getList", $read_all_func_name)
	group.NewRouter("/get", $read_func_name)
	group.NewRouter("/update", $update_func_name)
	group.NewRouter("/delete", $delete_func_name)
	group.Set()
}
`
