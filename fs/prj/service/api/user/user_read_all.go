package user

import (
	"github.com/liyonge-cm/go-api-cli-prj/model"
	"github.com/liyonge-cm/go-api-cli-prj/service/api/common"
	"github.com/liyonge-cm/go-api-cli-prj/service/mysql"

	"go.uber.org/zap"
)

type GetUsersApi struct {
	*common.Controller
	Data GetUsersRequest
}
type GetUsersRequest struct {
	common.Page
}
type GetUsersResponse struct {
	Data  []*model.User `json:"data"`
	Count int64         `json:"count"`
}

func GetUsers(c *common.Controller) {
	req := GetUsersApi{
		Controller: c,
	}
	defer req.Response()

	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed()
		return
	}

	records, count := req.getRecords()
	if req.Reply.IsStatusFailed() {
		return
	}

	res := GetUsersResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
}

func (req *GetUsersApi) getRecords() (records []*model.User, count int64) {
	records = make([]*model.User, 0)
	tx := mysql.DB.Model(&model.User{}).Where("status != ?", common.RecordStatusDeleted)

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
