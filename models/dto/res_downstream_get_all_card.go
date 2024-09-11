package dto

type ResDownstreamGetAllCard struct {
	ColumnId   string                     `json:"columnId"`
	ColumnName string                     `json:"columnName"`
	Cards      []CardDownstreamGetAllCard `json:"cards"`
}

type CardDownstreamGetAllCard struct {
	Summary      string                           `json:"summary"`
	IssueTypeId  string                           `json:"issueTypeId"`
	IssueType    string                           `json:"issueType"`
	Created      string                           `json:"created"`
	Updated      string                           `json:"updated"`
	PriorityId   string                           `json:"PriorityId"`
	PriorityName string                           `json:"priorityName"`
	AssigneeId   string                           `json:"assigneeId"`
	AssigneName  string                           `json:"assigneName"`
	Description  string                           `json:"description"`
	CreatorId    string                           `json:"creatorId"`
	CreatorName  string                           `json:"creatorName"`
	ReporterId   string                           `json:"reporterId"`
	ReporterName string                           `json:"reporterName"`
	Comments     []CommentDownstreamGetAllCard    `json:"comments"`
	Attachments  []AttachmentDownstreamGetAllCard `json:"attachment"`
}

type CommentDownstreamGetAllCard struct {
	CommentId  string `json:"commentId"`
	AuthorId   string `json:"authorId"`
	AuthorName string `json:"authorName"`
	Body       string `json:"body"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

type AttachmentDownstreamGetAllCard struct {
	AttachmentId string `json:"attachmentId"`
	FileName     string `json:"fileName"`
	AuthorId     string `json:"authorId"`
	AuthorName   string `json:"authorName"`
	Created      string `json:"created"`
	Url          string `json:"url"`
}
