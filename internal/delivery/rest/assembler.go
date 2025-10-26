package rest

import "github.com/iqbalnzls/sistem-manajemen-armada/internal/dto"

func toBaseResponse(data interface{}) dto.BaseResponse {
	return dto.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	}
}
