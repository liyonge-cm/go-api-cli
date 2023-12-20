package common

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"-"`
}

func NewResponse() *CommonResponse {
	r := &CommonResponse{
		Data: make(map[string]interface{}),
	}
	r.Success()
	return r
}

func (r *CommonResponse) MsgSet(s int, m string) *CommonResponse {
	r.Status = s
	r.Message = m
	return r
}
func (r *CommonResponse) DataSet(data interface{}) *CommonResponse {
	resData := make(map[string]interface{})
	b, _ := json.Marshal(data)
	_ = json.Unmarshal(b, &resData)
	for key, value := range resData {
		r.Data[key] = value
	}
	return r
}
func (r *CommonResponse) IsStatusFailed() bool {
	return r.Status > ReplyStatusSuccess
}

func (r *CommonResponse) Success() *CommonResponse {
	r.Status = ReplyStatusSuccess
	r.Message = ReplyMessageSuccess
	return r
}
func (r *CommonResponse) NewSuccess(data interface{}) *CommonResponse {
	r.Status = ReplyStatusSuccess
	r.Message = ReplyMessageSuccess
	return r
}
func (r *CommonResponse) BindRequestFailed() *CommonResponse {
	r.Status = ReplyStatusBindRequestFailed
	r.Message = ReplyMessageBindRequestFailed
	return r
}
func (r *CommonResponse) CreateFailed() *CommonResponse {
	r.Status = ReplyStatusCreateFailed
	r.Message = ReplyMessageCreateFailed
	return r
}
func (r *CommonResponse) ReadFailed() *CommonResponse {
	r.Status = ReplyStatusReadFailed
	r.Message = ReplyMessageReadFailed
	return r
}
func (r *CommonResponse) UpdateFailed() *CommonResponse {
	r.Status = ReplyStatusUpdateFailed
	r.Message = ReplyMessageUpdateFailed
	return r
}
func (r *CommonResponse) DeleteFailed() *CommonResponse {
	r.Status = ReplyStatusDeleteFailed
	r.Message = ReplyMessageDeleteFailed
	return r
}

func (r *CommonResponse) Response(c *gin.Context) {
	reply := make(map[string]interface{})
	reply["status"] = r.Status
	reply["message"] = r.Message
	if r.Data != nil {
		for k, v := range r.Data {
			reply[k] = v
		}
	}

	c.JSON(http.StatusOK, reply)
}
