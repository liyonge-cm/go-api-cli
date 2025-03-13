package user

import (
	"time"

	"github.com/liyonge-cm/go-api-cli-prj/model"
	"github.com/liyonge-cm/go-api-cli-prj/service/api/common"
	"github.com/liyonge-cm/go-api-cli-prj/service/mysql"

	"go.uber.org/zap"
)

type UpdateUserApi struct {
	*common.Controller
	Data UpdateUserRequest
}
type UpdateUserRequest struct {
	model.User
}

func UpdateUser(c *common.Controller) {
	req := UpdateUserApi{
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
	req.updateRecord()
	if req.Reply.IsStatusFailed() {
		return
	}
}

func (req *UpdateUserApi) checkParams() {
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

func (req *UpdateUserApi) updateRecord() {
	now := time.Now().Unix()
	record := &model.User{
		Name:      req.Data.Name,
		Age:       req.Data.Age,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.User{}).
		Where("id = ?", req.Data.Id).
		Updates(&record).Error
	if err != nil {
		req.Logger.Error("update record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.UpdateFailed()
		return
	}
}
