package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraCardRepository(cardId string) (res entity.JiraCards, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraCards{}).
		Where("card_id = ?", cardId).
		Scan(&res).
		Error
	return
}
