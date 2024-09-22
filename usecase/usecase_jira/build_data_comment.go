package usecase_jira

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func BuildCommentData(cardId string, dataComment dto.Comment) (dataOutput []dto.CommentDownstreamGetAllCard) {
	dataOutput = make([]dto.CommentDownstreamGetAllCard, dataComment.Total)
	for indexComent, comment := range dataComment.Comments {
		dataOutput[indexComent] = dto.CommentDownstreamGetAllCard{
			CommentId:  comment.Id,
			CardId:     cardId,
			AuthorId:   comment.Author.AccountId,
			AuthorName: comment.Author.DisplayName,
			Body:       comment.Body,
			Created:    helpers.KonversiTanggalIssueOutput(comment.Created),
			Updated:    helpers.KonversiTanggalIssueOutput(comment.Updated),
		}
	}
	return
}
