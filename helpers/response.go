package helpers

import "ms-sv-jira/models/dto"

func ResSuccess(isSuccess bool, responseCode string, message string, data interface{}) (int, dto.Res) {
	return 200, dto.Res{
		Success:      isSuccess,
		ResponseCode: responseCode,
		Message:      message,
		Data:         data,
	}
}

func ResInvalidCredential(kosong interface{}) (int, dto.Res) {
	return 401, dto.Res{
		Success:      false,
		ResponseCode: "1001",
		Message:      "Invalid credential",
		Data:         kosong,
	}
}

func ResInvalidValue(kosong interface{}) (int, dto.Res) {
	return 400, dto.Res{
		Success:      false,
		ResponseCode: "1002",
		Message:      "Invalid value",
		Data:         kosong,
	}
}

func ResGeneralSystemError(kosong interface{}) (int, dto.Res) {
	return 500, dto.Res{
		Success:      false,
		ResponseCode: "2001",
		Message:      "General system error",
		Data:         kosong,
	}
}

func ResBackendError(kosong interface{}) (int, dto.Res) {
	return 502, dto.Res{
		Success:      false,
		ResponseCode: "2002",
		Message:      "Backend error",
		Data:         kosong,
	}
}

func ResPageNotFound() (int, dto.Res) {
	return 404, dto.Res{
		Success:      false,
		ResponseCode: "2003",
		Message:      "Page not found",
		Data:         EmptyObject(),
	}
}

func ResIpBlocked(kosong interface{}) (int, dto.Res) {
	return 403, dto.Res{
		Success:      false,
		ResponseCode: "2004",
		Message:      "IP blocked",
		Data:         kosong,
	}
}
