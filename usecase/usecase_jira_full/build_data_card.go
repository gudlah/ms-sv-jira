package usecase_jira_full

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
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
						Started:      usecase.GetStartDate(kosong, idRequest, card.ID),
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

func (usecase *JiraFullUsecaseImpl) GetStartDate(kosong interface{}, idRequest, issueId string) (startDate string) {
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   "",
		RequestTimestamp: helpers.Now(),
	}
	resUpstream, err := usecase.ExternalRepository.GetAllChangelogRepository(issueId)
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
	} else {
		logUpstream.ResponsePayload = resUpstream.String()
		logUpstream.IsSuccess = 1
		var resStruct dto.ResUpstreamIssueChangelog
		json.Unmarshal(resUpstream.Body(), &resStruct)
		histories := resStruct.Changelog.Histories
		if len(histories) > 0 {
			for _, history := range histories {
				item := history.Items[0]
				if item.FromString == "To Do" && item.ToString == "In Progress" {
					startDate = helpers.KonversiTanggalIssueOutput(history.Created)
				}
			}
		}
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, 200, dto.Res{}, kosong)
	usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
