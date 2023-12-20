package templates

const readAllContent = `package $package_name

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
	common.Page
}
type $func_nameResponse struct {
	$all_data_response
	$count_response
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

	records, count := req.getRecords()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	res := $func_nameResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *$func_nameApi) getRecords() (records []*model.$model_name, count int64) {
	records = make([]*model.$model_name, 0)
	tx := mysql.DB.Model(&model.$model_name{}).Where("status != ?", common.RecordStatusDeleted)

	tx = tx.Count(&count)
	if count == 0 {
		return
	}

	if req.Data.Limit > 0 && req.Data.Offset > 0 {
		tx = tx.Limit(req.Data.Limit).Offset(req.Data.Limit * (req.Data.Offset - 1))
	}
	err := tx.Order("created_at desc").Find(&records).Error
	if err != nil {
		req.Logger.Error("get records failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}

	return records, count
}
`
