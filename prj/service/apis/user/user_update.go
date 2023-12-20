package user

import (
	"time"

	"go-cli-prj/model"
	"go-cli-prj/service/apis/common"
	"go-cli-prj/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateUserApi struct {
	*common.ApiCommon
	Data UpdateUserRequest
}
type UpdateUserRequest struct {
	model.User
}

func UpdateUser(c *gin.Context) {
	req := &UpdateUserApi{
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
	req.updateRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
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
