package usecase_jira

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
	"strconv"
)

func (usecase *JiraUsecaseImpl) BuildDataCard(kosong interface{}, idRequest string, columnId, columnName string, dataCard dto.ResUpstreamGetAllCard) (dataOutput []dto.CardDownstreamGetAllCard) {
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
					Started:      usecase.GetStartDate(kosong, idRequest, card.ID),
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

func (usecase *JiraUsecaseImpl) GetStartDate(kosong interface{}, idRequest, issueId string) (startDate string) {
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
