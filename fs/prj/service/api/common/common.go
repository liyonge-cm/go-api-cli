package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Controller struct {
	c      *gin.Context    `json:"-"`
	Logger *zap.Logger     `json:"-"`
	Reply  *CommonResponse `json:"-"`
}

type Page struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func NewRequest(c *gin.Context) *Controller {
	return &Controller{
		c:     c,
		Reply: NewResponse(),
	}
}
func (s *Controller) WithLogger(log *zap.Logger) {
	s.Logger = log.With(zap.String("uuid", uuid.NewString()))
}

func (r *Controller) BindRequest(req interface{}) (err error) {
	err = r.c.BindJSON(&req)
	if err != nil {
		r.Logger.Error("读取API入参失败", zap.Error(err))
		return err
	}
	return
}

func (r *Controller) Response() {
	reply := make(map[string]interface{})
	reply["status"] = r.Reply.Status
	reply["message"] = r.Reply.Message
	if r.Reply.Data != nil {
		for k, v := range r.Reply.Data {
			reply[k] = v
		}
	}

	r.c.JSON(http.StatusOK, reply)
}
