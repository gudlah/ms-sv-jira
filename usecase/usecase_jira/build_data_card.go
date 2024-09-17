package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"strconv"
)

func BuildDataCard(columnId, columnName string, dataCard dto.ResUpstreamGetAllCard) (dataOutput []dto.CardDownstreamGetAllCard) {
	if dataCard.Total == 0 {
		dataOutput = make([]dto.CardDownstreamGetAllCard, dataCard.Total)
	} else {
		cards := []dto.CardDownstreamGetAllCard{}
		for _, card := range dataCard.Issues {
			field := card.Fields
			if field.Status.Name == columnName {
				priorityIdInt, _ := strconv.Atoi(field.Priority.ID)
				dataBuildCard := dto.CardDownstreamGetAllCard{
					CardId:       card.ID,
					ColumnId:     columnId,
					CardTitle:    field.Summary,
					CardKey:      card.Key,
					CardTypeId:   field.IssueType.ID,
					CardTypeName: field.IssueType.Name,
					Created:      helpers.KonversiTanggalIssueOutput(field.Created),
					Updated:      helpers.KonversiTanggalIssueOutput(field.Updated),
					PriorityId:   priorityIdInt,
					PriorityName: field.Priority.Name,
					AssigneeId:   field.Assignee.AccountID,
					AssigneName:  field.Assignee.DisplayName,
					Description:  field.Description,
					CreatorId:    field.Creator.AccountID,
					CreatorName:  field.Creator.DisplayName,
					ReporterId:   field.Reporter.AccountID,
					ReporterName: field.Reporter.DisplayName,
					Resolved:     helpers.KonversiTanggalIssueOutput(field.ResolutionDate),
				}
				dataBuildCard.Comments = BuildCommentData(card.ID, field.Comment)
				dataBuildCard.Attachments = BuildAttachmentData(card.ID, field.Attachment)

				cards = append(cards, dataBuildCard)
			}
		}

		dataOutput = cards
	}
	return
}
