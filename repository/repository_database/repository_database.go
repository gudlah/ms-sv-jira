package repository_database

import (
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

type DatabaseRepository interface {
	GetUser(username string) (entity.ServiceUsers, error)
	InsertLogActivityRepository(log entity.ActivityLog) error
	InsertLogUpstreamRepository(log entity.UpstreamServiceRequestLog) error
	GetSymbolsRepository() ([]entity.Symbols, error)
	GetBlockedIpRepository(ipClient string) (entity.IpBlockeds, error)
	GetLastClientHitRepository(clientIp string) (dto.ResQueryGetLastHit, error)
	InsertIpBlockRepository(ipBlock entity.IpBlockeds) error
	SoftDeleteBlockedIpRepository(id string) error
	InsertJiraUserRepository(jiraUser []entity.JiraUsers) error
	GetJiraUserRepository(userId string) (entity.JiraUsers, error)
	UpdateJiraUserRepository(user entity.JiraUsers) error
	GetJiraProjectRepository(projectId string) (entity.JiraProjects, error)
	InsertJiraProjectRepository(jiraProject []entity.JiraProjects) error
	UpdateJiraProjectRepository(project entity.JiraProjects) error
	UpdateJiraBoardRepository(board entity.JiraBoards) error
	InsertJiraBoardRepository(jiraBoards []entity.JiraBoards) error
	GetJiraBoardRepository(boardId int) (entity.JiraBoards, error)
	GetJiraSprintRepository(sprintId int) (entity.JiraSprints, error)
	InsertJiraSprintRepository(JiraSprints []entity.JiraSprints) error
	UpdateJiraSprintRepository(sprint entity.JiraSprints) error
	GetJiraPrioritiesRepository(priorityId int) (entity.JiraPriorities, error)
	UpdateJiraPrioritiesRepository(priority entity.JiraPriorities) error
	InsertJiraPrioritiesRepository(JiraPriorities []entity.JiraPriorities) error
	GetJiraColumnRepository(columnId string) (res entity.JiraColumns, err error)
	InsertJiraColumnRepository(jiraColumns []entity.JiraColumns) error
	UpdateJiraColumnRepository(column entity.JiraColumns) error
	GetJiraCardRepository(cardId string) (res entity.JiraCards, err error)
	InsertJiraCardRepository(jiraCards []entity.JiraCards) error
	UpdateJiraCardRepository(card entity.JiraCards) error
	GetJiraCommentRepository(commentId string) (res entity.JiraComments, err error)
	InsertJiraCommentRepository(jiraComments []entity.JiraComments) error
	UpdateJiraCommentRepository(comment entity.JiraComments) error
	GetJiraSubTaskRepository(subTaskId string) (res entity.JiraSubTasks, err error)
	InsertJiraSubTaskRepository(jiraSubTasks []entity.JiraSubTasks) error
	UpdateJiraSubTaskRepository(subTask entity.JiraSubTasks) error
	GetJiraAttachmentRepository(attachmentId string) (res entity.JiraAttachments, err error)
	InsertJiraAttachmentRepository(jiraAttachment []entity.JiraAttachments) error
	UpdateJiraAttachmentRepository(attachment entity.JiraAttachments) error
}
