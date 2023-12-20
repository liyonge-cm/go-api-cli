package templates

var createContent = `package $package_name

import (
	"time"

	"$prj_name/model"
	"$prj_name/service/apis/common"
	"$prj_name/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type $func_nameApi struct {
	*common.ApiCommon
	Data *model.$model_name
}

func $func_name(c *gin.Context) {
	req := &$func_nameApi{
		ApiCommon: common.NewRequest(c),
	}

	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed().Response(c)
		return
	}
	req.checkParams()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	req.createRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
}

func (req *$func_nameApi) checkParams() {
}

func (req *$func_nameApi) createRecord() {
	now := time.Now().Unix()
	record := &model.$model_name{$params_content
		Status:    common.RecordStatusInit,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.$model_name{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
`
