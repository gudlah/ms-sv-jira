package repository_external

import (
	"github.com/go-resty/resty/v2"
)

type ExternalRepository interface {
	GetAllProjetRepository() (*resty.Response, error)
	GetAllBoardRepository() (*resty.Response, error)
	GetAllSprintRepository(boardId string) (*resty.Response, error)
	GetAllColumnRepository(boardId string) (*resty.Response, error)
	GetAllCardRepository(sprintId string) (*resty.Response, error)
}
