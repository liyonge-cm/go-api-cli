package gen_frame

import (
	"fmt"
	"path"
	"strings"

	"github.com/liyonge-cm/go-api-cli/config"
	"github.com/liyonge-cm/go-api-cli/fs"
	"github.com/liyonge-cm/go-api-cli/utils"
)

type GenFrameConfig struct {
	OutPath  string
	PrjName  string
	JsonCase string
}
type GenFrameService struct {
	cfg         *GenFrameConfig
	genPrjName  string
	genPrjPath  string
	genJsonCase string
	outPrjPath  string
}

func NewGenFrameConfig(outPath, prjName string) *GenFrameConfig {
	cfg := &GenFrameConfig{
		OutPath: outPath,
		PrjName: prjName,
	}
	if outPath == "" {
		cfg.OutPath = "./"
	}
	if prjName == "" {
		cfg.PrjName = "prj_aiee"
	}
	return cfg
}

func NewGenFrameService(cfg *GenFrameConfig) *GenFrameService {
	if cfg == nil {
		cfg = &GenFrameConfig{
			OutPath:  config.Cfg.Frame.OutPath,
			PrjName:  config.Cfg.Frame.PrjName,
			JsonCase: config.Cfg.Frame.JsonCase,
		}
	}
	s := &GenFrameService{
		cfg:         cfg,
		genPrjName:  "github.com/liyonge-cm/go-api-cli-prj",
		genPrjPath:  "./prj",
		genJsonCase: "snake",
		outPrjPath:  path.Join(cfg.OutPath, cfg.PrjName),
	}
	return s
}

func (s *GenFrameService) GenFrame() (err error) {
	if err = utils.CreateDir(s.outPrjPath); err != nil {
		return err
	}
	if err = s.genMod(); err != nil {
		return err
	}
	if err = s.genWithPrjFile("main.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("config/config.go", nil); err != nil {
		return err
	}
	if err = s.genWithPrjFile("config/config.yml", map[string]string{
		s.genPrjName:  s.cfg.PrjName,
		s.genPrjPath:  "..",
		s.genJsonCase: s.cfg.JsonCase,
	}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("app/app.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("app/run.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/logger/logger.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/mysql/mysql.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/api/api.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/api/common/common.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/api/common/router.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/api/common/const.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/api/common/reply.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("service/api/router/groups.go", map[string]string{s.genPrjName: s.cfg.PrjName}); err != nil {
		return err
	}
	if err = s.genWithPrjFile("README.md", nil); err != nil {
		return err
	}
	return nil
}

func (s *GenFrameService) replaceContent(content string, oldStr, newStr string) string {
	content = strings.ReplaceAll(content, oldStr, newStr)
	return content
}

func (s *GenFrameService) genWithPrjFile(fileName string, replace map[string]string) error {
	filePath := path.Join(s.outPrjPath, fileName)
	contentPath := path.Join(s.genPrjPath, fileName)
	datab, err := fs.ReadFile(contentPath)
	if err != nil {
		return err
	}
	content := string(datab)
	for oldStr, newStr := range replace {
		content = s.replaceContent(content, oldStr, newStr)
	}
	if s.cfg.JsonCase == "camel" {
		content = utils.JsonToCamel(content)
	}
	err = utils.SaveFile(filePath, []byte(content))
	if err != nil {
		fmt.Println("filePath err", err.Error())
		return err
	}

	return err
}

func (s *GenFrameService) genMod() error {
	fileName := "go.mod.bac"
	replace := map[string]string{s.genPrjName: s.cfg.PrjName}
	newPath := path.Join(s.outPrjPath, "go.mod")
	contentPath := path.Join(s.genPrjPath, fileName)
	datab, err := fs.ReadFile(contentPath)
	if err != nil {
		return err
	}
	content := string(datab)
	for oldStr, newStr := range replace {
		content = s.replaceContent(content, oldStr, newStr)
	}
	if s.cfg.JsonCase == "camel" {
		content = utils.JsonToCamel(content)
	}
	err = utils.SaveFile(newPath, []byte(content))
	if err != nil {
		fmt.Println("filePath err", err.Error())
		return err
	}

	return err
}
