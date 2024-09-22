package dto

type ResDownstreamGetAllFull struct {
	ProjectId      string                      `json:"projectId"`
	ProjectKey     string                      `json:"projectKey"`
	ProjectName    string                      `json:"projectName"`
	ProjectTypeKey string                      `json:"projectTypeKey"`
	Boards         []BoardDownstreamGetAllFull `json:"boards"`
}

type BoardDownstreamGetAllFull struct {
	BoardId   int                          `json:"boardId"`
	ProjectId string                       `json:"projectId"`
	BoardName string                       `json:"boardName"`
	BoardType string                       `json:"boardType"`
	Sprints   []SprintDownstreamGetAllFull `json:"sprints"`
}

type SprintDownstreamGetAllFull struct {
	SprintId    int                          `json:"sprintId"`
	BoardId     int                          `json:"boardId"`
	State       string                       `json:"state"`
	Name        string                       `json:"name"`
	StartDate   string                       `json:"startDate"`
	EndDate     string                       `json:"endDate"`
	CreatedDate string                       `json:"createdDate"`
	Goal        string                       `json:"goal"`
	Columns     []ColumnDownstreamGetAllFull `json:"columns"`
}

type ColumnDownstreamGetAllFull struct {
	ColumnId   string                     `json:"columnId"`
	SprintId   int                        `json:"sprintId"`
	ColumnName string                     `json:"columnName"`
	Cards      []CardDownstreamGetAllFull `json:"cards"`
}

type CardDownstreamGetAllFull struct {
	CardId       string                           `json:"cardId"`
	ColumnId     string                           `json:"columnId"`
	CardTitle    string                           `json:"cardTitle"`
	CardKey      string                           `json:"cardKey"`
	CardTypeId   string                           `json:"cardTypeId"`
	CardTypeName string                           `json:"cardType"`
	Created      string                           `json:"created"`
	Updated      string                           `json:"updated"`
	Resolved     string                           `json:"resolved"`
	PriorityId   int                              `json:"priorityId"`
	PriorityName string                           `json:"priorityName"`
	AssigneeId   string                           `json:"assigneeId"`
	AssigneName  string                           `json:"assigneName"`
	Description  string                           `json:"description"`
	CreatorId    string                           `json:"creatorId"`
	CreatorName  string                           `json:"creatorName"`
	ReporterId   string                           `json:"reporterId"`
	ReporterName string                           `json:"reporterName"`
	Comments     []CommentDownstreamGetAllFull    `json:"comments"`
	Attachments  []AttachmentDownstreamGetAllFull `json:"attachment"`
	SubTasks     []SubTaskDownstreamGetAllFull    `json:"subTasks"`
}

type CommentDownstreamGetAllFull struct {
	CommentId  string `json:"commentId"`
	CardId     string `json:"cardId"`
	AuthorId   string `json:"authorId"`
	AuthorName string `json:"authorName"`
	Body       string `json:"body"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

type AttachmentDownstreamGetAllFull struct {
	AttachmentId string `json:"attachmentId"`
	CardId       string `json:"cardId"`
	FileName     string `json:"fileName"`
	AuthorId     string `json:"authorId"`
	AuthorName   string `json:"authorName"`
	Created      string `json:"created"`
	Url          string `json:"url"`
}

type SubTaskDownstreamGetAllFull struct {
	SubTaskId          string `json:"subTaskId"`
	CardKey            string `json:"cardKey"`
	SubTaskKey         string `json:"subTaskKey"`
	SubTaskTitle       string `json:"subTaskTitle"`
	SubTaskDescription string `json:"subTaskDescription"`
	StatusId           string `json:"statusId"`
	StatusName         string `json:"statusName"`
	PriorityId         int    `json:"priorityId"`
	PriorityName       string `json:"priorityName"`
	Created            string `json:"created"`
	Updated            string `json:"updated"`
	Resolved           string `json:"resolved"`
	CreatorId          string `json:"creatorId"`
	CreatorName        string `json:"creatorName"`
	ReporterId         string `json:"reporterId"`
	ReporterName       string `json:"reporterName"`
	AssigneeId         string `json:"assigneeId"`
	AssigneeName       string `json:"assigneeName"`
}
