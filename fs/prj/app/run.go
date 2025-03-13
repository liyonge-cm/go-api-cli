package app

import (
	"github.com/liyonge-cm/go-api-cli-prj/service/api"
	"github.com/liyonge-cm/go-api-cli-prj/service/logger"
	"github.com/liyonge-cm/go-api-cli-prj/service/mysql"
)

func (s *App) Run() {

	s.WithLogger(logger.Logger)

	// regist and start services
	// 由于服务之间有依赖，根据依赖关系指定先后顺序的，api应该是最后的
	s.RegistService(mysql.NewMySQL(s.ctx, &s.config.MySQL))
	s.RegistService(api.NewHttp(s.ctx, &s.config.Service))

	s.StartServices()

}
