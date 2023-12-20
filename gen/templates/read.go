package templates

const readContent = `package $package_name

import (
	"$prj_name/model"
	"$prj_name/service/apis/common"
	"$prj_name/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type $func_nameApi struct {
	*common.ApiCommon
	Data $func_nameRequest
}
type $func_nameRequest struct {
	$id_request
}
type $func_nameResponse struct {
	$data_response
}

func $func_name(c *gin.Context) {
	req := &$func_nameApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed()
		req.Reply.Response(c)
		return
	}
	req.checkParams()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	record := req.getRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	res := $func_nameResponse{
		Data:  record,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *$func_nameApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
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
