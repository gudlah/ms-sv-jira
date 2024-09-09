package dto

type ResUpstreamGetAllProject struct {
	Expand         string      `json:"expand"`
	Self           string      `json:"self"`
	Id             string      `json:"id"`
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

type AvatarUrls struct {
	X48x48 string `json:"48x48"`
	X24x24 string `json:"24x24"`
	X16x16 string `json:"16x16"`
	X32x32 string `json:"32x32"`
}
