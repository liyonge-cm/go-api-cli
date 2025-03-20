package templates

const deleteContent = `package $package_name

import (
	"time"

	"$prj_name/model"
	"$prj_name/service/api/common"
	"$prj_name/service/mysql"

	"go.uber.org/zap"
)

type $func_nameApi struct {
	*common.Controller
	Data $func_nameRequest
}
type $func_nameRequest struct {
	$id_request
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
	req.deleteRecord()
	if req.Reply.IsStatusFailed() {
		return
	}
}

func (req *$func_nameApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.WithParamMiss("id")
		return
	}
	var count int64
	err := mysql.DB.Model(&model.$model_name{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		Count(&count).Error
	if err != nil {
		req.Logger.Error("check record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	if count <= 0 {
		req.Reply.WithParamFailed("id")
		return
	}
}

func (req *$func_nameApi) deleteRecord() {
	now := time.Now().Unix()
	record := &model.$model_name{
		Status:    common.RecordStatusDeleted,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.$model_name{}).
		Where("id = ?", req.Data.Id).
		Updates(&record).Error
	if err != nil {
		req.Logger.Error("delete record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.DeleteFailed()
		return
	}
}
`
