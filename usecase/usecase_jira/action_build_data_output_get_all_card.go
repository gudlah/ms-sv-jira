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
		}

		if dataCard.Total == 0 {
			dataOutput[indexColumn].Cards = make([]dto.CardDownstreamGetAllCard, dataCard.Total)
		} else {
			issues := []dto.CardDownstreamGetAllCard{}
			for _, card := range dataCard.Issues {
				field := card.Fields
				if field.Status.Name == column.Name {
					dataIssue := dto.CardDownstreamGetAllCard{
						Summary:      field.Summary,
						IssueTypeId:  field.IssueType.ID,
						IssueType:    field.IssueType.Name,
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
					}
					dataIssue.Comments = make([]dto.CommentDownstreamGetAllCard, field.Comment.Total)
					for indexComent, comment := range field.Comment.Comments {
						dataIssue.Comments[indexComent] = dto.CommentDownstreamGetAllCard{
							CommentId:  comment.Id,
							AuthorId:   comment.Author.AccountId,
							AuthorName: comment.Author.DisplayName,
							Body:       comment.Body,
							Created:    comment.Created,
							Updated:    comment.Updated,
						}
					}
					dataIssue.Attachments = make([]dto.AttachmentDownstreamGetAllCard, len(field.Attachment))
					for indexAttachment, attachment := range field.Attachment {
						dataIssue.Attachments[indexAttachment] = dto.AttachmentDownstreamGetAllCard{
							AttachmentId: attachment.Id,
							FileName:     attachment.Filename,
							AuthorId:     attachment.Author.AccountId,
							AuthorName:   attachment.Author.DisplayName,
							Created:      attachment.Created,
							Url:          attachment.Content,
						}
					}

					issues = append(issues, dataIssue)
				}
			}

			dataOutput[indexColumn].Cards = issues
		}
	}
	return
}
