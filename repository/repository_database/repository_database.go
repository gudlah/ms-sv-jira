package repository_database

import (
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

type DatabaseRepository interface {
	GetUser(username string) (entity.Users, error)
	InsertLogActivityRepository(log entity.ActivityLog) error
	InsertLogUpstreamRepository(log entity.UpstreamServiceRequestLog) error
	GetMappingUpstreamResponse(endpoint, responseCodeBrigate string) (entity.UpstreamResponseMapping, error)
	InsertSecretKeyRepository(data entity.SecretKey) error
	GetSecretKeyRepository(jenisTransaksi, phoneNumber string) (entity.SecretKey, error)
	UpdateSecretKeyRepository(data entity.SecretKey) error
	GetJumlahSessionRepository(sessionId, aEndpoint, bEndpoint string) (dto.JumlahLog, error)
	GetSymbolsRepository() ([]entity.Symbols, error)
	GetBlockedIpRepository(ipClient string) (entity.IpBlockeds, error)
	GetLastClientHitRepository(clientIp string) (dto.ResQueryGetLastHit, error)
	InsertIpBlockRepository(ipBlock entity.IpBlockeds) error
	SoftDeleteBlockedIpRepository(id string) error
}
