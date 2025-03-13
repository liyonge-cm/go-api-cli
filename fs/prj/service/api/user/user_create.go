package user

import (
	"time"

	"github.com/liyonge-cm/go-api-cli-prj/model"
	"github.com/liyonge-cm/go-api-cli-prj/service/api/common"
	"github.com/liyonge-cm/go-api-cli-prj/service/mysql"

	"go.uber.org/zap"
)

type CreateUserApi struct {
	*common.Controller
	Data *model.User
}

func CreateUser(c *common.Controller) {
	req := CreateUserApi{
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
	req.createRecord()
	if req.Reply.IsStatusFailed() {
		return
	}
}

func (req *CreateUserApi) checkParams() {
}

func (req *CreateUserApi) createRecord() {
	now := time.Now().Unix()
	record := &model.User{
		Name:      req.Data.Name,
		Age:       req.Data.Age,
		Status:    common.RecordStatusInit,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.User{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
