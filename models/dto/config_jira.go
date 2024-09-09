package dto

type JiraConfig struct {
	Url   string `env:"JIRA_URL"`
	Email string `env:"JIRA_EMAIL"`
	Token string `env:"JIRA_TOKEN"`
}
