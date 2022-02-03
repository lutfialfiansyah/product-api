package response

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Meta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type PaginationMeta struct {
	Status      bool        `json:"status"`
	Code        int64       `json:"code"`
	Message     string      `json:"message"`
	CurrentPage int         `json:"current_page"`
	NextPage    interface{} `json:"next_page"`
	PrevPage    interface{} `json:"prev_page"`
	PerPage     int         `json:"per_page"`
	PageCount   int         `json:"page_count"`
	TotalCount  int64       `json:"total_count"`
}
type ModelPaginationResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{Data: nil, Message: message})
}

func Json(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{Data: data, Message: "OK"})
}
func PaginationResponse(code, total int64, page, perPage string, data interface{}) *ModelPaginationResponse {
	res := new(ModelPaginationResponse)
	convPage, _ := strconv.Atoi(page)
	convPerPage, _ := strconv.Atoi(perPage)
	page_count := int(math.Ceil(float64(total) / float64(convPerPage)))
	hasNext := false
	if float64(convPage) < float64(page_count) {
		hasNext = true
	}
	meta := PaginationMeta{
		Message:     "success",
		Code:        code,
		Status:      true,
		CurrentPage: convPage,
		NextPage:    hasNext,
		PrevPage:    convPage > 1,
		PerPage:     convPerPage,
		PageCount:   page_count,
		TotalCount:  total,
	}
	res.Meta = meta
	res.Data = data

	return res
}
