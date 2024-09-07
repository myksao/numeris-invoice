package utils

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
	Data      any    `json:"data"`
}

func StandardResponse(ctx *gin.Context, StatusCode int, data interface{}) {

	var response APIResponse

	switch StatusCode {
	case 200:
		response = APIResponse{
			Status:    "success",
			Message:   "Request successful",
			Reference: "200",
			Data:      data,
		}

	case 201:
		response = APIResponse{
			Status:    "success",
			Message:   "Resource created",
			Reference: "201",
			Data:      data,
		}

	case 400:
		response = APIResponse{
			Status:    "error",
			Message:   "Bad request",
			Reference: "400",
			Data:      data,
		}

	case 401:
		response = APIResponse{
			Status:    "error",
			Message:   "Unauthorized",
			Reference: "401",
			Data:      data,
		}

	case 403:
		response = APIResponse{
			Status:    "error",
			Message:   "Forbidden",
			Reference: "403",
			Data:      data,
		}
	case 404:
		response = APIResponse{
			Status:    "error",
			Message:   "Resource not found",
			Reference: "404",
			Data:      data,
		}
	case 409:
		response = APIResponse{
			Status:    "error",
			Message:   "Conflict",
			Reference: "409",
			Data:      data,
		}
	}

	ctx.JSON(StatusCode, response)
}

func Pagination(limit, offset, total uint) map[string]interface{} {
	pagination := make(map[string]interface{})

	pagination["limit"] = limit
	pagination["offset"] = offset
	pagination["total"] = total

	return pagination
}

type DefaultCreateRes struct {
	ID string `json:"id"`
}
