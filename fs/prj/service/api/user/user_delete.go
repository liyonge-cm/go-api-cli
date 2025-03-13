package user

import (
	"time"

	"github.com/liyonge-cm/go-api-cli-prj/model"
	"github.com/liyonge-cm/go-api-cli-prj/service/api/common"
	"github.com/liyonge-cm/go-api-cli-prj/service/mysql"

	"go.uber.org/zap"
)

type DeleteUserApi struct {
	*common.Controller
	Data DeleteUserRequest
}
type DeleteUserRequest struct {
	Id int `json:"id"`
}

func DeleteUser(c *common.Controller) {
	req := DeleteUserApi{
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

func (req *DeleteUserApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
	var count int64
	err := mysql.DB.Model(&model.User{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		Count(&count).Error
	if err != nil {
		req.Logger.Error("check record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	if count <= 0 {
		req.Reply.MsgSet(common.ReplyStatusBindRequestFailed, common.ReplyMessageBindRequestFailed)
		return
	}
}

func (req *DeleteUserApi) deleteRecord() {
	now := time.Now().Unix()
	record := &model.User{
		Status:    common.RecordStatusDeleted,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.User{}).
		Where("id = ?", req.Data.Id).
		Updates(&record).Error
	if err != nil {
		req.Logger.Error("delete record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.DeleteFailed()
		return
	}
}
