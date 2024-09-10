package helpers

import (
	"encoding/json"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func BuildActivityLog(param dto.ActivityLogParam) entity.ActivityLog {
	stringResponse, _ := json.Marshal(param.Response)
	return entity.ActivityLog{
		Id:           GenerateUUID(),
		IdRequest:    param.IdRequest,
		IdUser:       param.IdUser,
		Endpoint:     param.Endpoint,
		ResponseCode: param.Response.ResponseCode,
		BodyRequest:  TrimAll(param.Request),
		BodyResponse: string(stringResponse),
		ClientIp:     param.IpClient,
		CreatedAt:    Now(),
	}
}
