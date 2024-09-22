package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraUsecaseImpl) InsertJiraAttachmentAction(kosong interface{}, attachments []dto.AttachmentDownstreamGetAllCard, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataAttachments := []entity.JiraAttachments{}
	var errGetComment int
	for _, attachment := range attachments {
		cekAttachment, err := usecase.DatabaseRepository.GetJiraAttachmentRepository(attachment.AttachmentId)
		dataInsertAttachment := entity.JiraAttachments{
			AttachmentID:       attachment.AttachmentId,
			CardID:             attachment.CardId,
			AttachmentFileName: attachment.FileName,
			AttachmentURL:      attachment.Url,
			AuthorID:           attachment.AuthorId,
			AttachmentCreated:  helpers.KonversiTanggalDb(attachment.Created),
		}
		if err != nil {
			errGetComment += 1
		} else if cekAttachment.AttachmentID == "" {
			dataAttachments = append(dataAttachments, dataInsertAttachment)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraAttachmentRepository(dataInsertAttachment)
			if err != nil {
				errGetComment += 1
			}
		}
	}

	if errGetComment > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataAttachments) > 0 {
		err := usecase.DatabaseRepository.InsertJiraAttachmentRepository(dataAttachments)
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = httpCodeAsal, resAsal
		}
	} else {
		httpCode, res = httpCodeAsal, resAsal
	}

	return
}
