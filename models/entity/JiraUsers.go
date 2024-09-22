package entity

type JiraUsers struct {
	UserID          string `json:"user_id" gorm:"column:user_id;size:100;primaryKey"`
	UserEmail       string `json:"user_email,omitempty" gorm:"column:user_email;size:100"`
	UserDisplayName string `json:"user_display_name,omitempty" gorm:"column:user_display_name;size:100"`
	UserActive      bool   `json:"user_active,omitempty" gorm:"column:user_active;default:true"`
	UserLocale      string `json:"user_locale,omitempty" gorm:"column:user_locale;size:10"`
}
