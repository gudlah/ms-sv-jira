package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func BuildAttachmentData(cardId string, dataAttachment []dto.Attachment) (dataOutput []dto.AttachmentDownstreamGetAllCard) {
	dataOutput = make([]dto.AttachmentDownstreamGetAllCard, len(dataAttachment))
	for indexAttachment, attachment := range dataAttachment {
		dataOutput[indexAttachment] = dto.AttachmentDownstreamGetAllCard{
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
