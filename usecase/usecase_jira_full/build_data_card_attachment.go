package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func BuildDataCardAttachment(cardId string, dataAttachment []dto.Attachment) (dataOutput []dto.AttachmentDownstreamGetAllFull) {
	dataOutput = make([]dto.AttachmentDownstreamGetAllFull, len(dataAttachment))
	for indexAttachment, attachment := range dataAttachment {
		dataOutput[indexAttachment] = dto.AttachmentDownstreamGetAllFull{
			AttachmentId: attachment.Id,
			CardId:       cardId,
			FileName:     attachment.Filename,
			AuthorId:     attachment.Author.AccountId,
			AuthorName:   attachment.Author.DisplayName,
			Created:      helpers.KonversiTanggalIssueOutput(attachment.Created),
			Url:          attachment.Content,
		}
	}
	return
}
