package usecase_jira_full

import "ms-sv-jira/models/dto"

func BuildDataSubTaskDescription(dataDescription dto.DescriptionUpstreamGetAllSubTask) (outputDescription string) {
	if dataDescription.Type != "" {
		for _, descriptionContent := range dataDescription.Content {
			if descriptionContent.Type == "paragraph" {
				for _, valueDescriptionContent := range descriptionContent.Content {
					if valueDescriptionContent.Type == "text" {
						outputDescription = valueDescriptionContent.Text
					}
				}
			}
		}
	}
	return
}
