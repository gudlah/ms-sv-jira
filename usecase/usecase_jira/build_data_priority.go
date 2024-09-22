package usecase_jira

import (
	"ms-sv-jira/models/dto"
	"strconv"
)

func BuildDataPriority(priorities []dto.ResUpstreamGetAllPriority) (dataOutput []dto.ResDownstreamGetAllPriority) {
	dataOutput = make([]dto.ResDownstreamGetAllPriority, len(priorities))
	for indexPriority, priority := range priorities {
		priorityIdInt, _ := strconv.Atoi(priority.Id)
		dataOutput[indexPriority] = dto.ResDownstreamGetAllPriority{
			PriorityId:          priorityIdInt,
			PriorityName:        priority.Name,
			PriorityDescription: priority.Description,
		}
	}

	return
}
