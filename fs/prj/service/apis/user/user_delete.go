package user

import (
	"time"

	"github.com/liyonge-cm/go-api-cli-prj/model"
	"github.com/liyonge-cm/go-api-cli-prj/service/apis/common"
	"github.com/liyonge-cm/go-api-cli-prj/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteUserApi struct {
	*common.ApiCommon
	Data DeleteUserRequest
}
type DeleteUserRequest struct {
	Id int `json:"id"`
}

func DeleteUser(c *gin.Context) {
	req := &DeleteUserApi{
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
	req.deleteRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
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
