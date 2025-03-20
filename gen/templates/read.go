package templates

const readContent = `package $package_name

import (
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
type $func_nameResponse struct {
	$data_response
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
	record := req.getRecord()
	if req.Reply.IsStatusFailed() {
		return
	}
	res := $func_nameResponse{
		Data:  record,
	}
	req.Reply.DataSet(res)
}

func (req *$func_nameApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.WithParamMiss("id")
		return
	}
}

func (req *$func_nameApi) getRecord() (record *model.$model_name) {
	err := mysql.DB.Model(&model.$model_name{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		First(&record).Error
	if err != nil {
		req.Logger.Error("get record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	return record
}
`
