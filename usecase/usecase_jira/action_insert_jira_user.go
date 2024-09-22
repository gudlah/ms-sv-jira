package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraUsecaseImpl) InsertJiraUserAction(kosong interface{}, users []dto.ResDownstreamGetAllUser, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataUsers := []entity.JiraUsers{}
	var errGetUser int
	for _, user := range users {
		cekUser, err := usecase.DatabaseRepository.GetJiraUserRepository(user.AccountID)
		dataInsertUser := entity.JiraUsers{
			UserID:          user.AccountID,
			UserEmail:       user.EmailAddress,
			UserDisplayName: user.DisplayName,
			UserActive:      user.Active,
			UserLocale:      user.Locale,
		}
		if err != nil {
			errGetUser += 1
		} else if cekUser.UserID == "" {
			dataUsers = append(dataUsers, dataInsertUser)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraUserRepository(dataInsertUser)
			if err != nil {
				errGetUser += 1
			}
		}
	}

	if errGetUser > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataUsers) > 0 {
		err := usecase.DatabaseRepository.InsertJiraUserRepository(dataUsers)
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = httpCodeAsal, resAsal
		}
	} else {
		httpCode, res = httpCodeAsal, resAsal
	}

	return
}
