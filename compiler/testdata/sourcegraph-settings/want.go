package p

type SearchSavedQueries struct {
// Description description: Description of this saved query
	Description string `json:"description"`
// Key description: Unique key for this query in this file
	Key string `json:"key"`
// Notify description: Notify the owner of this configuration file when new results are available
	Notify bool `json:"notify,omitempty"`
// NotifySlack description: Notify Slack via the organization's Slack webhook URL when new results are available
	NotifySlack bool `json:"notifySlack,omitempty"`
// Query description: Query string
	Query string `json:"query"`
// ShowOnHomepage description: Show this saved query on the homepage
	ShowOnHomepage bool `json:"showOnHomepage,omitempty"`
}
type SearchScope struct {
// Description description: A description for this search scope
	Description string `json:"description,omitempty"`
// Id description: A unique identifier for the search scope.
// 
// If set, a scoped search page is available at https://[sourcegraph-hostname]/search/scope/ID, where ID is this value.
	Id string `json:"id,omitempty"`
// Name description: The human-readable name for this search scope
	Name string `json:"name"`
// Value description: The query string of this search scope
	Value string `json:"value"`
}

// Settings description: Configuration settings for users and organizations on Sourcegraph.
type Settings struct {
// Motd description: An array of messages (often with just one element) to display at the top of all pages, including for unauthenticated users. Users may dismiss a message (and any message with the same string value will remain dismissed for the user).
// 
// Markdown formatting is supported.
// 
// Usually this setting is used in global and organization settings. If set in user settings, the message will only be displayed to that user. (This is useful for testing the correctness of the message's Markdown formatting.)
// 
// MOTD stands for "message of the day" (which is the conventional Unix name for this type of message).
	Motd               []string                  `json:"motd,omitempty"`
	NotificationsSlack *SlackNotificationsConfig `json:"notifications.slack,omitempty"`
// SearchRepositoryGroups description: Named groups of repositories that can be referenced in a search query using the repogroup: operator.
	SearchRepositoryGroups map[string][]string `json:"search.repositoryGroups,omitempty"`
// SearchSavedQueries description: Saved search queries
	SearchSavedQueries []*SearchSavedQueries `json:"search.savedQueries,omitempty"`
// SearchScopes description: Predefined search scopes
	SearchScopes []*SearchScope `json:"search.scopes,omitempty"`
}

// SlackNotificationsConfig description: Configuration for sending notifications to Slack.
type SlackNotificationsConfig struct {
// WebhookURL description: The Slack webhook URL used to post notification messages to a Slack channel. To obtain this URL, go to: https://YOUR-WORKSPACE-NAME.slack.com/apps/new/A0F7XDUAZ-incoming-webhooks
	WebhookURL string `json:"webhookURL"`
}
