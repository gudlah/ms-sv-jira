package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func BuildDataCardComment(cardId string, dataComment dto.Comment) (dataOutput []dto.CommentDownstreamGetAllFull) {
	dataOutput = make([]dto.CommentDownstreamGetAllFull, dataComment.Total)
	for indexComent, comment := range dataComment.Comments {
		dataOutput[indexComent] = dto.CommentDownstreamGetAllFull{
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
