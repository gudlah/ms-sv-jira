package usecase_jira_full

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (usecase *JiraFullUsecaseImpl) InsertJiraBoardAction(kosong interface{}, boards []dto.BoardDownstreamGetAllFull, httpCodeAsal int, resAsal dto.Res) (httpCode int, res dto.Res) {
	dataBoards := []entity.JiraBoards{}
	var errBoard int
	for _, board := range boards {
		cekBoard, err := usecase.DatabaseRepository.GetJiraBoardRepository(board.BoardId)
		dataInserBoards := entity.JiraBoards{
			BoardID:   board.BoardId,
			ProjectID: board.ProjectId,
			BoardName: board.BoardName,
			BoardType: board.BoardType,
		}
		if err != nil {
			errBoard += 1
		} else if cekBoard.BoardID == 0 {
			dataBoards = append(dataBoards, dataInserBoards)
		} else {
			err = usecase.DatabaseRepository.UpdateJiraBoardRepository(dataInserBoards)
			if err != nil {
				errBoard += 1
			}
		}
	}

	if errBoard > 0 {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if len(dataBoards) > 0 {
		err := usecase.DatabaseRepository.InsertJiraBoardRepository(dataBoards)
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
		}
	} else {
		httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
	}

	return
}
