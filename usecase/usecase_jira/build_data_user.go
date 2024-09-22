package usecase_jira

import "ms-sv-jira/models/dto"

func BuildDataUser(users []dto.ResUpstreamGetAllUser) (dataOutput []dto.ResDownstreamGetAllUser) {
	for _, user := range users {
		if user.AccountType == "atlassian" {
			dataOutput = append(dataOutput, dto.ResDownstreamGetAllUser{
				AccountID:    user.AccountID,
				DisplayName:  user.DisplayName,
				Active:       user.Active,
				Locale:       user.Locale,
				EmailAddress: user.EmailAddress,
			})
		}
	}

	return
}
