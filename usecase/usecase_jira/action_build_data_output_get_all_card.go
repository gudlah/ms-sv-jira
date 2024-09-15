package usecase_jira

import (
	"ms-sv-jira/models/dto"
)

func (usecase *JiraUsecaseImpl) BuildDataOuputGetAllCardAction(dataColumn dto.ResUpstreamGetAllColumn, dataCard dto.ResUpstreamGetAllCard) (dataOutput []dto.ResDownstreamGetAllCard) {
	dataOutput = make([]dto.ResDownstreamGetAllCard, len(dataColumn.ColumnConfig.Columns))
	for indexColumn, column := range dataColumn.ColumnConfig.Columns {
		dataOutput[indexColumn] = dto.ResDownstreamGetAllCard{
			ColumnId:   column.Statuses[0].Id,
			ColumnName: column.Name,
			Cards:      buildDataCard(column.Name, dataCard),
		}
	}
	return
}

func buildDataCard(columnName string, dataCard dto.ResUpstreamGetAllCard) (dataOutput []dto.CardDownstreamGetAllCard) {
	if dataCard.Total == 0 {
		dataOutput = make([]dto.CardDownstreamGetAllCard, dataCard.Total)
	} else {
		cards := []dto.CardDownstreamGetAllCard{}
		for _, card := range dataCard.Issues {
			field := card.Fields
			if field.Status.Name == columnName {
				dataBuildCard := dto.CardDownstreamGetAllCard{
					CardId:       card.ID,
					CardTitle:    field.Summary,
					CardKey:      card.Key,
					CardTypeId:   field.IssueType.ID,
					CardTypeName: field.IssueType.Name,
					Created:      field.Created,
					Updated:      field.Updated,
					PriorityId:   field.Priority.ID,
					PriorityName: field.Priority.Name,
					AssigneeId:   field.Assignee.AccountID,
					AssigneName:  field.Assignee.DisplayName,
					Description:  field.Description,
					CreatorId:    field.Creator.AccountID,
					CreatorName:  field.Creator.DisplayName,
					ReporterId:   field.Reporter.AccountID,
					ReporterName: field.Reporter.DisplayName,
					Resolved:     field.ResolutionDate,
				}
				dataBuildCard.Comments = buildCommentData(field.Comment)
				dataBuildCard.Attachments = buildAttachmentData(field.Attachment)

				cards = append(cards, dataBuildCard)
			}
		}

		dataOutput = cards
	}
	return
}

func buildCommentData(dataComment dto.Comment) (dataOutput []dto.CommentDownstreamGetAllCard) {
	dataOutput = make([]dto.CommentDownstreamGetAllCard, dataComment.Total)
	for indexComent, comment := range dataComment.Comments {
		dataOutput[indexComent] = dto.CommentDownstreamGetAllCard{
			CommentId:  comment.Id,
			AuthorId:   comment.Author.AccountId,
			AuthorName: comment.Author.DisplayName,
			Body:       comment.Body,
			Created:    comment.Created,
			Updated:    comment.Updated,
		}
	}
	return
}

func buildAttachmentData(dataAttachment []dto.Attachment) (dataOutput []dto.AttachmentDownstreamGetAllCard) {
	dataOutput = make([]dto.AttachmentDownstreamGetAllCard, len(dataAttachment))
	for indexAttachment, attachment := range dataAttachment {
		dataOutput[indexAttachment] = dto.AttachmentDownstreamGetAllCard{
			AttachmentId: attachment.Id,
			FileName:     attachment.Filename,
			AuthorId:     attachment.Author.AccountId,
			AuthorName:   attachment.Author.DisplayName,
			Created:      attachment.Created,
			Url:          attachment.Content,
		}
	}
	return
}
