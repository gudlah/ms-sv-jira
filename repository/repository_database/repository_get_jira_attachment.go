package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraAttachmentRepository(attachmentId string) (res entity.JiraAttachments, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraAttachments{}).
		Where("attachment_id = ?", attachmentId).
		Scan(&res).
		Error
	return
}
