package pkg

import (
	"gogin/config"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code        int         `json:"-"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data,omitempty"`
	Meta        interface{} `json:"meta,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
	ctx.Abort()
	return
}

func NewRes(code int, data *config.Result) *Response {
	var response = Response{
		Code:   code,
		Status: getStatus(code),
	}

	if response.Code >= 400 {
		response.Description = data.Data
	} else if data.Message != nil {
		response.Description = data.Message
	} else {
		response.Data = data.Data
	}

	if data.Meta != nil {
		response.Meta = data.Meta
	}

	return &response
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
		break
	case 201:
		desc = "Created"
		break
	case 400:
		desc = "Bad Request"
		break
	case 401:
		desc = "Unauthorized"
		break
	case 403:
		desc = "Forbidden"
		break
	case 404:
		desc = "Not Found"
		break
	case 500:
		desc = "Internal Server Error"
		break
	case 501:
		desc = "Bad Gateway"
		break
	case 304:
		desc = "Not Modified"
		break
	default:
		desc = ""
	}

	return desc
}
