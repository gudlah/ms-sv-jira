package repository_database

import (
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

type DatabaseRepository interface {
	GetServiceUserRepository(username string) (entity.ServiceUsers, error)
	InsertLogActivityRepository(log entity.ActivityLog) error
	InsertLogUpstreamRepository(log entity.UpstreamServiceRequestLog) error
	GetSymbolsRepository() ([]entity.Symbols, error)
	GetBlockedIpRepository(ipClient string) (entity.IpBlockeds, error)
	GetLastClientHitRepository(clientIp string) (dto.ResQueryGetLastHit, error)
	InsertIpBlockRepository(ipBlock entity.IpBlockeds) error
	SoftDeleteBlockedIpRepository(id string) error
	GetAllJiraUsersRepository() (res []entity.JiraUsers, err error)
	GetAllJiraProjectRepository() ([]entity.JiraProjects, error)
	GetAllJiraBoardsRepository(projectId string) ([]entity.JiraBoards, error)
}
