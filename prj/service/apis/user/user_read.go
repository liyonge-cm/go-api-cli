package user

import (
	"go-api-cli-prj/model"
	"go-api-cli-prj/service/apis/common"
	"go-api-cli-prj/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetUserApi struct {
	*common.ApiCommon
	Data GetUserRequest
}
type GetUserRequest struct {
	Id int `json:"id"`
}
type GetUserResponse struct {
	Data *model.User `json:"data"`
}

func GetUser(c *gin.Context) {
	req := &GetUserApi{
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

	req.Reply.DataSet(GetUserResponse{Data: record})
	req.Reply.Response(c)
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
