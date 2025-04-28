package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Respond(c *gin.Context, code int, data any) {
	msg, ok := messages[code]
	if !ok {
		msg = MessageDetail{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Unknown error",
		}
	}

	c.JSON(msg.HTTPCode, ResponseData{
		Code:    code,
		Message: msg.Message,
		Data:    data,
	})
}
