package dto

type IssueType struct {
	ID             string `json:"id"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Name           string `json:"name"`
	Subtask        bool   `json:"subtask"`
	AvatarID       int    `json:"avatarId"`
	EntityID       string `json:"entityId"`
	HierarchyLevel int    `json:"hierarchyLevel"`
}

type StatusCategory struct {
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}

type Status struct {
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconURL        string         `json:"iconUrl"`
	Name           string         `json:"name"`
	ID             string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

type Priority struct {
	Self    string `json:"self"`
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

type ParentFields struct {
	Summary   string    `json:"summary"`
	Status    Status    `json:"status"`
	Priority  Priority  `json:"priority"`
	IssueType IssueType `json:"issuetype"`
}

type Parent struct {
	ID     string       `json:"id"`
	Key    string       `json:"key"`
	Self   string       `json:"self"`
	Fields ParentFields `json:"fields"`
}

type Assignee struct {
	AccountID   string            `json:"accountId"`
	AvatarURLs  map[string]string `json:"avatarUrls"`
	DisplayName string            `json:"displayName"`
	Active      bool              `json:"active"`
	TimeZone    string            `json:"timeZone"`
	AccountType string            `json:"accountType"`
}

type Watchers struct {
	Self       string `json:"self"`
	WatchCount int    `json:"watchCount"`
	IsWatching bool   `json:"isWatching"`
}

type Issuerestriction struct {
	Issuerestrictions map[string]interface{} `json:"issuerestrictions"`
	ShouldDisplay     bool                   `json:"shouldDisplay"`
}

type Progress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type Subtask struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Self   string `json:"self"`
	Fields struct {
		Summary   string    `json:"summary"`
		Status    Status    `json:"status"`
		Priority  Priority  `json:"priority"`
		IssueType IssueType `json:"issuetype"`
	} `json:"fields"`
}

type Reporter struct {
	Self        string            `json:"self"`
	AccountID   string            `json:"accountId"`
	AvatarURLs  map[string]string `json:"avatarUrls"`
	DisplayName string            `json:"displayName"`
	Active      bool              `json:"active"`
	TimeZone    string            `json:"timeZone"`
	AccountType string            `json:"accountType"`
}

type Comment struct {
	Comments   []CommentValue `json:"comments"`
	Self       string         `json:"self"`
	MaxResults int            `json:"maxResults"`
	Total      int            `json:"total"`
	StartAt    int            `json:"startAt"`
}

type CommentValue struct {
	Id      string        `json:"id"`
	Body    string        `json:"body"`
	Created string        `json:"created"`
	Updated string        `json:"updated"`
	Author  CommentAuthor `json:"author"`
}

type CommentAuthor struct {
	AccountId   string `json:"accountId"`
	DisplayName string `json:"displayNAme"`
}

type Votes struct {
	Self     string `json:"self"`
	Votes    int    `json:"votes"`
	HasVoted bool   `json:"hasVoted"`
}

type Worklog struct {
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Worklogs   []interface{} `json:"worklogs"`
}

type IssueFields struct {
	Expand                   string           `json:"expand"`
	StatusCategoryChangeDate string           `json:"statuscategorychangedate"`
	IssueType                IssueType        `json:"issuetype"`
	Parent                   *Parent          `json:"parent,omitempty"`
	Timespent                interface{}      `json:"timespent"`
	CustomField10031         interface{}      `json:"customfield_10031"`
	CustomField10033         interface{}      `json:"customfield_10033"`
	FixVersions              []interface{}    `json:"fixVersions"`
	Aggregatetimespent       interface{}      `json:"aggregatetimespent"`
	Resolution               interface{}      `json:"resolution"`
	CustomField10036         interface{}      `json:"customfield_10036"`
	CustomField10027         interface{}      `json:"customfield_10027"`
	ResolutionDate           string           `json:"resolutiondate"`
	Workratio                int              `json:"workratio"`
	Issuerestriction         Issuerestriction `json:"issuerestriction"`
	LastViewed               string           `json:"lastViewed"`
	Watchers                 Watchers         `json:"watches"`
	Created                  string           `json:"created"`
	CustomField10021         interface{}      `json:"customfield_10021"`
	Epic                     interface{}      `json:"epic"`
	CustomField10022         interface{}      `json:"customfield_10022"`
	Priority                 Priority         `json:"priority"`
	Labels                   []string         `json:"labels"`
	Assignee                 Assignee         `json:"assignee"`
	Updated                  string           `json:"updated"`
	Status                   Status           `json:"status"`
	Components               []interface{}    `json:"components"`
	Summary                  string           `json:"summary"`
	Creator                  Reporter         `json:"creator"`
	Subtasks                 []Subtask        `json:"subtasks"`
	Reporter                 Reporter         `json:"reporter"`
	AggregateProgress        Progress         `json:"aggregateprogress"`
	Progress                 Progress         `json:"progress"`
	Comment                  Comment          `json:"comment"`
	Votes                    Votes            `json:"votes"`
	Worklog                  Worklog          `json:"worklog"`
	Description              string           `json:"description"`
	Attachment               []Attachment     `json:"attachment"`
}

type Attachment struct {
	Id       string           `json:"id"`
	Filename string           `json:"filename"`
	Created  string           `json:"created"`
	Content  string           `json:"content"`
	Author   AttachmentAuthor `json:"author"`
}

type AttachmentAuthor struct {
	AccountId   string `json:"accountId"`
	DisplayName string `json:"displayName"`
}

type Issue struct {
	ID     string      `json:"id"`
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type ResUpstreamGetAllCard struct {
	Expand     string  `json:"expand"`
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}
