package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraFullUsecaseImpl) InsertJiraCommentAction(kosong interface{}, comments []dto.CommentDownstreamGetAllFull, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataComments := []entity.JiraComments{}
	var errGetComment int
	for _, comment := range comments {
		cekComment, err := usecase.DatabaseRepository.GetJiraCommentRepository(comment.CommentId)
		dataInsertComment := entity.JiraComments{
			CommentID:      comment.CommentId,
			CardID:         comment.CardId,
			AuthorID:       comment.AuthorId,
			CommentBody:    comment.Body,
			CommentCreated: helpers.KonversiTanggalDb(comment.Created),
			CommentUpdated: helpers.KonversiTanggalDb(comment.Updated),
		}
		if err != nil {
			errGetComment += 1
		} else if cekComment.CommentID == "" {
			dataComments = append(dataComments, dataInsertComment)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraCommentRepository(dataInsertComment)
			if err != nil {
				errGetComment += 1
			}
		}
	}

	if errGetComment > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataComments) > 0 {
		err := usecase.DatabaseRepository.InsertJiraCommentRepository(dataComments)
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
