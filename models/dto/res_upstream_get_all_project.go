package dto

type ResUpstreamGetAllProject struct {
	Expand         string      `json:"expand"`
	Self           string      `json:"self"`
	Id             int         `json:"id"`
	Key            string      `json:"key"`
	Name           string      `json:"name"`
	AvatarUrls     AvatarUrls  `json:"avatarUrls"`
	ProjectTypeKey string      `json:"projectTypeKey"`
	Simplified     bool        `json:"simplified"`
	Style          string      `json:"style"`
	IsPrivate      bool        `json:"isPrivate"`
	Properties     interface{} `json:"properties"`
	EntityID       string      `json:"entityId"`
	Uuid           string      `json:"uuid"`
}
