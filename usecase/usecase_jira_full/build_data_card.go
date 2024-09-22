package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"strconv"
)

func (usecase *JiraFullUsecaseImpl) BuildDataCard(kosong interface{}, idRequest, columnId, columnName string, dataCard dto.ResUpstreamGetAllCard) (httpCode int, res dto.Res, dataOutput []dto.CardDownstreamGetAllFull) {
	if dataCard.Total == 0 {
		dataOutput = make([]dto.CardDownstreamGetAllFull, dataCard.Total)
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
	} else {
		cards := []dto.CardDownstreamGetAllFull{}
		for _, card := range dataCard.Issues {
			field := card.Fields
			dataComments := BuildDataCardComment(card.ID, field.Comment)
			dataAttachment := BuildDataCardAttachment(card.ID, field.Attachment)
			var dataSubTasks []dto.SubTaskDownstreamGetAllFull
			if field.Status.Name == columnName {
				httpCode, res, dataSubTasks = usecase.GetDataSubTaskAction(kosong, idRequest, card.Key)
				if res.ResponseCode != "0000" {
					return
				} else {
					priorityIdInt, _ := strconv.Atoi(field.Priority.ID)
					dataBuildCard := dto.CardDownstreamGetAllFull{
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
						Comments:     dataComments,
						Attachments:  dataAttachment,
						SubTasks:     dataSubTasks,
					}
					cards = append(cards, dataBuildCard)
					httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", dataOutput)
				}
			} else {
				httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
			}
			httpCode, res = usecase.InsertJiraCommentAction(kosong, dataComments, httpCode, res)
			if res.ResponseCode != "0000" {
				return
			}
			httpCode, res = usecase.InsertJiraAttachmentAction(kosong, dataAttachment, httpCode, res)
			if res.ResponseCode != "0000" {
				return
			}
			httpCode, res = usecase.InsertJiraSubTaskAction(kosong, dataSubTasks, httpCode, res)
			if res.ResponseCode != "0000" {
				return
			}
		}
		dataOutput = cards
		httpCode, res = usecase.InsertJiraCardAction(kosong, dataOutput, httpCode, res)
	}
	return
}
