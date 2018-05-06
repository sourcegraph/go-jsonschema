package p

// SearchSavedQueries
type SearchSavedQueries struct {
	Description    string `json:"description"`
	Key            string `json:"key"`
	Notify         bool   `json:"notify,omitempty"`
	NotifySlack    bool   `json:"notifySlack,omitempty"`
	Query          string `json:"query"`
	ShowOnHomepage bool   `json:"showOnHomepage,omitempty"`
}

// SearchScope
type SearchScope struct {
	Description string `json:"description,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

// Settings Configuration settings for users and organizations on Sourcegraph.
type Settings struct {
	Motd                   []string                  `json:"motd,omitempty"`
	NotificationsSlack     *SlackNotificationsConfig `json:"notifications.slack,omitempty"`
	SearchRepositoryGroups map[string][]string       `json:"search.repositoryGroups,omitempty"`
	SearchSavedQueries     []*SearchSavedQueries     `json:"search.savedQueries,omitempty"`
	SearchScopes           []*SearchScope            `json:"search.scopes,omitempty"`
}

// SlackNotificationsConfig Configuration for sending notifications to Slack.
type SlackNotificationsConfig struct {
	WebhookURL string `json:"webhookURL"`
}
