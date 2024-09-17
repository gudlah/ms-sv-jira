package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraCommentRepository(commentId string) (res entity.JiraComments, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraComments{}).
		Where("comment_id = ?", commentId).
		Scan(&res).
		Error
	return
}
