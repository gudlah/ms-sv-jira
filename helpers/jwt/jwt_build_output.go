package helpers

import (
	"ms-sv-jira/models/dto"
	"time"
)

func BuildOutputJWT(token string, durasi time.Duration) dto.ResTokenJWT {
	layout := "2006-01-02 15:04:05 MST"
	expiredAt := time.Now().Add(durasi).Format(layout)
	return dto.ResTokenJWT{
		Token:     token,
		ExpiredAt: expiredAt,
	}
}
