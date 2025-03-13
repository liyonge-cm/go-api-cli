package gen

import (
	"path"

	"github.com/liyonge-cm/go-api-cli/config"

	"gorm.io/gorm"
)

/*
1. 连接数据库
2. 查询表字段
3. 生成model
4. 生成API
*/
type GenServer struct {
	cfg          *config.Config
	db           *gorm.DB
	tableInfos   map[string]*TableInfo
	modelPath    string
	modelPkgName string
	apiPath      string
	isJsonCamel  bool

	RemaneTableFileName  func(name string) string
	RemaneTableModelName func(name string) string
}

func NewGenServer(tables []string) *GenServer {
	g := &GenServer{
		cfg: &config.Cfg,
	}
	if tables != nil {
		g.cfg.Api.Tables = tables
	}
	g.initCfgDefault()
	return g
}

func (s *GenServer) initCfgDefault() {
	if s.cfg.Frame.OutPath == "" {
		s.cfg.Frame.OutPath = "./"
	}
	if s.cfg.Frame.PrjName == "" {
		s.cfg.Frame.PrjName = "prj_aiee"
	}
	if s.cfg.Api.Tables == nil {
		s.cfg.Api.Tables = []string{}
	}

	s.modelPath = path.Join(s.cfg.Frame.OutPath, s.cfg.Frame.PrjName, "model")
	s.apiPath = path.Join(s.cfg.Frame.OutPath, s.cfg.Frame.PrjName, "service/api")

	s.modelPkgName = "model"
	s.isJsonCamel = s.cfg.Frame.JsonCase == "camel"
	s.tableInfos = make(map[string]*TableInfo)
}
