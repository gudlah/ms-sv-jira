package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraUsecaseImpl) InsertJiraCardAction(kosong interface{}, cards []dto.CardDownstreamGetAllCard, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataCards := []entity.JiraCards{}
	var errGetCard int
	for _, card := range cards {
		cekCard, err := usecase.DatabaseRepository.GetJiraCardRepository(card.CardId)
		dataInsertCard := entity.JiraCards{
			CardID:          card.CardId,
			ColumnID:        card.ColumnId,
			CardTitle:       card.CardTitle,
			CardKey:         card.CardKey,
			CardDescription: card.Description,
			CardTypeID:      card.CardTypeId,
			PriorityID:      card.PriorityId,
			AssigneeID:      card.AssigneeId,
			CreatorID:       card.CreatorId,
			ReporterID:      card.ReporterId,
			CardCreated:     helpers.KonversiTanggalDb(card.Created),
			CardUpdated:     helpers.KonversiTanggalDb(card.Updated),
			CardResolved:    helpers.KonversiTanggalDb(card.Resolved),
		}
		if err != nil {
			errGetCard += 1
		} else if cekCard.CardID == "" {
			dataCards = append(dataCards, dataInsertCard)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraCardRepository(dataInsertCard)
			if err != nil {
				errGetCard += 1
			}
		}
	}

	if errGetCard > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataCards) > 0 {
		err := usecase.DatabaseRepository.InsertJiraCardRepository(dataCards)
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = httpCodeAsal, resAsal
		}
	} else {
		httpCode, res = httpCodeAsal, resAsal
	}

	var errInsertComment int
	for _, card := range cards {
		_, resComment := usecase.InsertJiraCommentAction(kosong, card.Comments, httpCode, res)
		if resComment.ResponseCode != "0000" {
			errInsertComment += 1
		}
	}

	var errInserAttachment int
	for _, card := range cards {
		_, resAttachment := usecase.InsertJiraAttachmentAction(kosong, card.Attachments, httpCode, res)
		if resAttachment.ResponseCode != "0000" {
			errInserAttachment += 1
		}
	}

	return
}
