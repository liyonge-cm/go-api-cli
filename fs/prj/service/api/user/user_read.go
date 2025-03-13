package user

import (
	"github.com/liyonge-cm/go-api-cli-prj/model"
	"github.com/liyonge-cm/go-api-cli-prj/service/api/common"
	"github.com/liyonge-cm/go-api-cli-prj/service/mysql"

	"go.uber.org/zap"
)

type GetUserApi struct {
	*common.Controller
	Data GetUserRequest
}
type GetUserRequest struct {
	Id int `json:"id"`
}
type GetUserResponse struct {
	Data *model.User `json:"data"`
}

func GetUser(c *common.Controller) {
	req := GetUserApi{
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

	req.Reply.DataSet(GetUserResponse{Data: record})
}

func (req *GetUserApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
}

func (req *GetUserApi) getRecord() (record *model.User) {
	err := mysql.DB.Model(&model.User{}).
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
