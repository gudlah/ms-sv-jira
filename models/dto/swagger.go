package dto

type SwaggerGetAllFullSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         *struct {
		ProjectID      string `json:"projectId" example:"10000"`
		ProjectKey     string `json:"projectKey" example:"SV"`
		ProjectName    string `json:"projectName" example:"Parsing-API"`
		ProjectTypeKey string `json:"projectTypeKey" example:"software"`
		Boards         []*struct {
			BoardID   int    `json:"boardId" example:"1"`
			ProjectID string `json:"projectId" example:"10000"`
			BoardName string `json:"boardName" example:"SCRUM board"`
			BoardType string `json:"boardType" example:"simple"`
			Sprints   []*struct {
				SprintID    int    `json:"sprintId" example:"1"`
				BoardID     int    `json:"boardId" example:"1"`
				State       string `json:"state" example:"active"`
				Name        string `json:"name" example:"SCRUM Sprint 1"`
				StartDate   string `json:"startDate" example:"2024-09-06 11:06:01"`
				EndDate     string `json:"endDate" example:"2024-09-13 11:06:48"`
				CreatedDate string `json:"createdDate" example:"2024-09-06 11:03:22"`
				Goal        string `json:"goal" example:"End the parsing"`
				Columns     []*struct {
					ColumnID   string `json:"columnId" example:"10001"`
					SprintID   int    `json:"sprintId" example:"1"`
					ColumnName string `json:"columnName" example:"In Progress"`
					Cards      []*struct {
						CardID       string `json:"cardId" example:"10032"`
						ColumnID     string `json:"columnId" example:"10001"`
						CardTitle    string `json:"cardTitle" example:"Development Fitur : Cek Hari Libur"`
						CardKey      string `json:"cardKey" example:"SV-11"`
						CardTypeID   string `json:"cardTypeId" example:"10001"`
						CardType     string `json:"cardType" example:"Task"`
						Created      string `json:"created" example:"2024-09-08 03:28:23"`
						Updated      string `json:"updated" example:"2024-09-16 02:30:07"`
						Resolved     string `json:"resolved" example:"2024-09-19 02:30:07"`
						PriorityID   int    `json:"priorityId" example:"3"`
						PriorityName string `json:"priorityName" example:"Medium"`
						AssigneeID   string `json:"assigneeId" example:"712020:1af3618d-09c5-40f0-991e-23e644b9f220"`
						AssigneeName string `json:"assigneeName" example:"Bagus Anjar Sadewa"`
						Description  string `json:"description" example:"Task details here"`
						Comments     []*struct {
							CommentID  string `json:"commentId" example:"12345"`
							CardID     string `json:"cardId" example:"10032"`
							AuthorID   string `json:"authorId" example:"712020:1af3618d-09c5-40f0-991e-23e644b9f220"`
							AuthorName string `json:"authorName" example:"Bagus Anjar Sadewa"`
							Body       string `json:"body" example:"This is a comment"`
							Created    string `json:"created" example:"2024-09-12 02:51:42"`
						} `json:"comments"`
						Attachments []*struct {
							AttachmentID string `json:"attachmentId" example:"10002"`
							CardID       string `json:"cardId" example:"10032"`
							FileName     string `json:"fileName" example:"New Project.png"`
							AuthorID     string `json:"authorId" example:"712020:1af3618d-09c5-40f0-991e-23e644b9f220"`
							AuthorName   string `json:"authorName" example:"Bagus Anjar Sadewa"`
							Created      string `json:"created" example:"2024-09-12 02:51:42"`
							URL          string `json:"url" example:"https://it-team-sharingvision.atlassian.net/rest/api/2/attachment/content/10002"`
						} `json:"attachment"`
						SubTasks []*struct {
							SubTaskID    string `json:"subTaskId" example:"10045"`
							CardKey      string `json:"cardKey" example:"SV-11"`
							SubTaskKey   string `json:"subTaskKey" example:"SV-24"`
							SubTaskTitle string `json:"subTaskTitle" example:"Collect LC Testing"`
							StatusName   string `json:"statusName" example:"To Do"`
							PriorityID   int    `json:"priorityId" example:"3"`
							PriorityName string `json:"priorityName" example:"Medium"`
							Created      string `json:"created" example:"2024-09-08 03:42:49"`
							Updated      string `json:"updated" example:"2024-09-08 03:43:07"`
						} `json:"subTasks"`
					} `json:"cards"`
				} `json:"columns"`
			} `json:"sprints"`
		} `json:"boards"`
	} `json:"data"`
}

type SwaggerGetAllUserSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         []*struct {
		AccountID    string `json:"accountId" example:"6212f56f468c2e00716df01c"`
		DisplayName  string `json:"displayName" example:"Taofik Rakhman Sudrajat"`
		Active       bool   `json:"active" example:"true"`
		Locale       string `json:"locale" example:"en_US"`
		EmailAddress string `json:"emailAddress" example:"bagus.anjar@it.sharingvision.com"`
	} `json:"data"`
}

type SwaggerGetAllProjectSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         []*struct {
		ProjectID      string `json:"projectId" example:"10000"`
		ProjectKey     string `json:"projectKey" example:"SV"`
		ProjectName    string `json:"projectName" example:"Parsing-API"`
		ProjectTypeKey string `json:"projectTypeKey" example:"software"`
		Boards         []*struct {
			BoardID   int    `json:"boardId" example:"1"`
			ProjectID string `json:"projectId" example:"10000"`
			BoardName string `json:"boardName" example:"SCRUM board"`
			BoardType string `json:"boardType" example:"simple"`
		} `json:"boards"`
	} `json:"data"`
}

type SwaggerGetAllPrioritiesSuccess struct {
	Success      bool   `json:"success" example:"false"`
	ResponseCode string `json:"responseCode" example:"1010"`
	Message      string `json:"message" example:"Succesfully"`
	Data         []*struct {
		PriorityID          int    `json:"priorityId" example:"1"`
		PriorityName        string `json:"priorityName" example:"Highest"`
		PriorityDescription string `json:"priorityDescription" example:"This problem will block progress."`
	} `json:"data"`
}

type SwaggerGetAllSprintSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         []*struct {
		SprintID    int    `json:"sprintId" example:"1"`
		BoardID     int    `json:"boardId" example:"1"`
		State       string `json:"state" example:"active"`
		Name        string `json:"name" example:"SCRUM Sprint 1"`
		StartDate   string `json:"startDate" example:"2024-09-06 11:06:01"`
		EndDate     string `json:"endDate" example:"2024-09-13 11:06:48"`
		CreatedDate string `json:"createdDate" example:"2024-09-06 11:03:22"`
		Goal        string `json:"goal" example:"End the parsing"`
	} `json:"data"`
}

type SwaggerGetAllSubTaskSuccess struct {
	Success      bool   `json:"success" example:"false"`
	ResponseCode string `json:"responseCode" example:"1009"`
	Message      string `json:"message" example:"Successfully"`
	Data         []*struct {
		SubTaskID          string `json:"subTaskId" example:"10045"`
		CardKey            string `json:"cardKey" example:"SV-11"`
		SubTaskKey         string `json:"subTaskKey" example:"SV-24"`
		SubTaskTitle       string `json:"subTaskTitle" example:"Collect LC Testing"`
		SubTaskDescription string `json:"subTaskDescription" example:"Lorem ipsum dolor"`
		StatusID           string `json:"statusId" example:"10000"`
		StatusName         string `json:"statusName" example:"To Do"`
		PriorityID         int    `json:"priorityId" example:"3"`
		PriorityName       string `json:"priorityName" example:"Medium"`
		Created            string `json:"created" example:"2024-09-08 03:42:49"`
		Updated            string `json:"updated" example:"2024-09-09 03:43:07"`
		Resolved           string `json:"resolved" example:"2024-09-10 03:43:07"`
		CreatorID          string `json:"creatorId" example:"712020:1af3618d-09c5-40f0-991e-23e644b9f220"`
		CreatorName        string `json:"creatorName" example:"Bagus Anjar Sadewa"`
		ReporterID         string `json:"reporterId" example:"712020:1af3618d-09c5-40f0-991e-23e644b9f220"`
		ReporterName       string `json:"reporterName" example:"Bagus Anjar Sadewa"`
		AssigneeID         string `json:"assigneeId" example:"712020:1af3618d-09c5-40f0-991e-23e644b9f220"`
		AssigneeName       string `json:"assigneeName" example:"Bagus Anjar Sadewa"`
	} `json:"data"`
}

type SwaggerGetAlCardSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         *struct {
		ColumnID   string `json:"columnId" example:"10001"`
		SprintID   int    `json:"sprintId" example:"1"`
		ColumnName string `json:"columnName" example:"In Progress"`
		Cards      []*struct {
			CardID       string `json:"cardId" example:"10000"`
			ColumnID     string `json:"columnId" example:"10001"`
			CardTitle    string `json:"cardTitle" example:"Parsing-Task"`
			CardKey      string `json:"cardKey" example:"SV-1"`
			CardTypeID   string `json:"cardTypeId" example:"10001"`
			CardType     string `json:"cardType" example:"Task"`
			Created      string `json:"created" example:"2024-09-06 11:04:14"`
			Updated      string `json:"updated" example:"2024-09-10 18:53:22"`
			Resolved     string `json:"resolved" example:"2024-09-14 18:53:22"`
			PriorityID   int    `json:"priorityId" example:"3"`
			PriorityName string `json:"priorityName" example:"Medium"`
			AssigneeID   string `json:"assigneeId" example:"6212f56f468c2e00716df01c"`
			AssigneeName string `json:"assigneName" example:"Taofik Rakhman Sudrajat"`
			Description  string `json:"description" example:"Something here is a description hehe…"`
			CreatorID    string `json:"creatorId" example:"6212f56f468c2e00716df01c"`
			CreatorName  string `json:"creatorName" example:"Taofik Rakhman Sudrajat"`
			ReporterID   string `json:"reporterId" example:"6212f56f468c2e00716df01c"`
			ReporterName string `json:"reporterName" example:"Taofik Rakhman Sudrajat"`
			Comments     []*struct {
				CommentID  string `json:"commentId" example:"10000"`
				CardID     string `json:"cardId" example:"10000"`
				AuthorID   string `json:"authorId" example:"6212f56f468c2e00716df01c"`
				AuthorName string `json:"authorName" example:"Taofik Rakhman Sudrajat"`
				Body       string `json:"body" example:"THIS IS COMMENT ASDASDASDASDASDASDAS"`
				Created    string `json:"created" example:"2024-09-10 18:09:33"`
				Updated    string `json:"updated" example:"2024-09-10 18:09:33"`
			} `json:"comments"`
			Attachments []*struct {
				AttachmentID string `json:"attachmentId" example:"10000"`
				CardID       string `json:"cardId" example:"10000"`
				FileName     string `json:"fileName" example:"Modul 3 Pengembangan TI.pdf"`
				AuthorID     string `json:"authorId" example:"6212f56f468c2e00716df01c"`
				AuthorName   string `json:"authorName" example:"Taofik Rakhman Sudrajat"`
				Created      string `json:"created" example:"2024-09-06 11:06:47"`
				URL          string `json:"url" example:"https://it-team-sharingvision.atlassian.net/rest/api/2/attachment/content/10000"`
			} `json:"attachment"`
		} `json:"cards"`
	} `json:"data"`
}

type SwaggerInvalidCredential struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"1001"`
	Message      string    `json:"message" example:"Invalid credential"`
	Data         *struct{} `json:"data"`
}

type SwaggerInvalidValue struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"1002"`
	Message      string    `json:"message" example:"Invalid value"`
	Data         *struct{} `json:"data"`
}

type SwaggerGeneralSystemError struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"2001"`
	Message      string    `json:"message" example:"General system error"`
	Data         *struct{} `json:"data"`
}

type SwaggerBackendError struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"2002"`
	Message      string    `json:"message" example:"Backend error"`
	Data         *struct{} `json:"data"`
}

type SwaggerIpBlocked struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"2004"`
	Message      string    `json:"message" example:"IP blocked"`
	Data         *struct{} `json:"data"`
}
