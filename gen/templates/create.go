package templates

var createContent = `package $package_name

import (
	"time"

	"$prj_name/model"
	"$prj_name/service/api/common"
	"$prj_name/service/mysql"

	"go.uber.org/zap"
)

type $func_nameApi struct {
	*common.Controller
	Data *model.$model_name
}

func $func_name(c *common.Controller) {
	req := &$func_nameApi{
		Controller: c,
	}
	defer req.Response()	

	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed()
		return
	}
	req.checkParams()
	if req.Reply.IsStatusFailed() {
		return
	}
	req.createRecord()
	if req.Reply.IsStatusFailed() {
		return
	}
}

func (req *$func_nameApi) checkParams() {
}

func (req *$func_nameApi) createRecord() {
	now := time.Now().Unix()
	req.Data.Status = common.RecordStatusInit
	req.Data.CreatedAt = now
	req.Data.UpdatedAt = now

	err := mysql.DB.Model(&model.$model_name{}).Create(&req.Data).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", req.Data), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
`
