package p

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AWSCodeCommitConnection struct {
	// AccessKeyID description: The AWS access key ID to use when listing and updating repositories from AWS CodeCommit. Must have the AWSCodeCommitReadOnly IAM policy.
	AccessKeyID string `json:"accessKeyID"`
	// InitialRepositoryEnablement description: Defines whether repositories from AWS CodeCommit should be enabled and cloned when they are first seen by Sourcegraph. If false, the site admin must explicitly enable AWS CodeCommit repositories (in the site admin area) to clone them and make them searchable on Sourcegraph. If true, they will be enabled and cloned immediately (subject to rate limiting by AWS); site admins can still disable them explicitly, and they'll remain disabled.
	InitialRepositoryEnablement bool `json:"initialRepositoryEnablement,omitempty"`
	// Region description: The AWS region in which to access AWS CodeCommit. See the list of supported regions at https://docs.aws.amazon.com/codecommit/latest/userguide/regions.html#regions-git.
	Region string `json:"region"`
	// RepositoryPathPattern description: The pattern used to generate a the corresponding Sourcegraph repository path for an AWS CodeCommit repository. In the pattern, the variable "{name}" is replaced with the repository's name.
	//
	// For example, if your Sourcegraph instance is at https://src.example.com, then a repositoryPathPattern of "awsrepos/{name}" would mean that a AWS CodeCommit repository named "myrepo" is available on Sourcegraph at https://src.example.com/awsrepos/myrepo.
	RepositoryPathPattern string `json:"repositoryPathPattern,omitempty"`
	// SecretAccessKey description: The AWS secret access key (that corresponds to the AWS access key ID set in `accessKeyID`).
	SecretAccessKey string `json:"secretAccessKey"`
}
type AuthProviders struct {
	Builtin       *BuiltinAuthProvider
	Saml          *SAMLAuthProvider
	Openidconnect *OpenIDConnectAuthProvider
	HttpHeader    *HTTPHeaderAuthProvider
}

func (v AuthProviders) MarshalJSON() ([]byte, error) {
	if v.Builtin != nil {
		return json.Marshal(v.Builtin)
	}
	if v.Saml != nil {
		return json.Marshal(v.Saml)
	}
	if v.Openidconnect != nil {
		return json.Marshal(v.Openidconnect)
	}
	if v.HttpHeader != nil {
		return json.Marshal(v.HttpHeader)
	}
	return nil, errors.New("tagged union type must have exactly 1 non-nil field value")
}
func (v *AuthProviders) UnmarshalJSON(data []byte) error {
	var d struct {
		DiscriminantProperty string `json:"type"`
	}
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	switch d.DiscriminantProperty {
	case "builtin":
		return json.Unmarshal(data, &v.Builtin)
	case "http-header":
		return json.Unmarshal(data, &v.HttpHeader)
	case "openidconnect":
		return json.Unmarshal(data, &v.Openidconnect)
	case "saml":
		return json.Unmarshal(data, &v.Saml)
	}
	return fmt.Errorf("tagged union type must have a %q property whose value is one of %s", "type", []string{"builtin", "saml", "openidconnect", "http-header"})
}

type BitbucketServerConnection struct {
	// Certificate description: TLS certificate of a Bitbucket Server instance.
	Certificate string `json:"certificate,omitempty"`
	// GitURLType description: The type of Git URLs to use for cloning and fetching Git repositories on this Bitbucket Server instance.
	//
	// If "http", Sourcegraph will access Bitbucket Server repositories using Git URLs of the form http(s)://bitbucket.example.com/scm/myproject/myrepo.git (using https: if the Bitbucket Server instance uses HTTPS).
	//
	// If "ssh", Sourcegraph will access Bitbucket Server repositories using Git URLs of the form ssh://git@example.bitbucket.com/myproject/myrepo.git. See the documentation for how to provide SSH private keys and known_hosts: https://about.sourcegraph.com/docs/config/repositories#repositories-that-need-https-or-ssh-authentication.
	GitURLType string `json:"gitURLType,omitempty"`
	// InitialRepositoryEnablement description: Defines whether repositories from this Bitbucket Server instance should be enabled and cloned when they are first seen by Sourcegraph. If false, the site admin must explicitly enable Bitbucket Server repositories (in the site admin area) to clone them and make them searchable on Sourcegraph. If true, they will be enabled and cloned immediately (subject to rate limiting by Bitbucket Server); site admins can still disable them explicitly, and they'll remain disabled.
	InitialRepositoryEnablement bool `json:"initialRepositoryEnablement,omitempty"`
	// Password description: The password to use when authenticating to the Bitbucket Server instance. Also set the corresponding "username" field.
	//
	// For Bitbucket Server instances that support personal access tokens (Bitbucket Server version 5.5 and newer), it is recommended to provide a token instead (in the "token" field).
	Password string `json:"password,omitempty"`
	// RepositoryPathPattern description: The pattern used to generate the corresponding Sourcegraph repository path for a Bitbucket Server repository.
	//
	//  - "{host}" is replaced with the Bitbucket Server URL's host (such as bitbucket.example.com)
	//  - "{projectKey}" is replaced with the Bitbucket repository's parent project key (such as "PRJ")
	//  - "{repositorySlug}" is replaced with the Bitbucket repository's slug key (such as "my-repo").
	//
	// For example, if your Bitbucket Server is https://bitbucket.example.com and your Sourcegraph is https://src.example.com, then a repositoryPathPattern of "{host}/{projectKey}/{repositorySlug}" would mean that a Bitbucket Server repository at https://bitbucket.example.com/projects/PRJ/repos/my-repo is available on Sourcegraph at https://src.example.com/bitbucket.example.com/PRJ/my-repo.
	RepositoryPathPattern string `json:"repositoryPathPattern,omitempty"`
	// Token description: A Bitbucket Server personal access token with Read scope. Create one at https://[your-bitbucket-hostname]/plugins/servlet/access-tokens/add.
	//
	// For Bitbucket Server instances that don't support personal access tokens (Bitbucket Server version 5.4 and older), specify user-password credentials in the "username" and "password" fields.
	Token string `json:"token,omitempty"`
	// Url description: URL of a Bitbucket Server instance, such as https://bitbucket.example.com
	Url string `json:"url"`
	// Username description: The username to use when authenticating to the Bitbucket Server instance. Also set the corresponding "password" field.
	//
	// For Bitbucket Server instances that support personal access tokens (Bitbucket Server version 5.5 and newer), it is recommended to provide a token instead (in the "token" field).
	Username string `json:"username,omitempty"`
}

// BuiltinAuthProvider description: Configures the builtin username-password authentication provider.
type BuiltinAuthProvider struct {
	// AllowSignup description: Allows new visitors to sign up for accounts. The sign-up page will be enabled and accessible to all visitors.
	//
	// SECURITY: If the site has no users (i.e., during initial setup), it will always allow the first user to sign up and become site admin **without any approval** (first user to sign up becomes the admin).
	AllowSignup bool   `json:"allowSignup,omitempty"`
	Type        string `json:"type"`
}

// ExperimentalFeatures description: Experimental features to enable or disable. Features that are now enabled by default are marked as deprecated.
type ExperimentalFeatures struct {
	// EnhancedSAML description: Enables or disables the enhanced SAML implementation (which supports better error reporting, configuration, and usage with other auth providers).
	EnhancedSAML string `json:"enhancedSAML,omitempty"`
	// HostSurveysLocally description: Enables or disables hosting user satisfaction surveys locally. Once stable, this feature will be default.
	HostSurveysLocally string `json:"hostSurveysLocally,omitempty"`
	// JumpToDefOSSIndex description: Enables or disables consulting the OSS package index on Sourcegraph.com for cross repository jump to definition. When enabled Sourcegraph.com will receive Code Intelligence requests when they fail to resolve locally. NOTE: disablePublicRepoRedirects must not be set, or should be set to false.
	JumpToDefOSSIndex string `json:"jumpToDefOSSIndex,omitempty"`
	// MultipleAuthProviders description: Enables or disables the use of multiple authentication providers and a publicly accessible web page displaying authentication options for unauthenticated users. (WARNING: Do not use this unless you know what you're doing.)
	MultipleAuthProviders string `json:"multipleAuthProviders,omitempty"`
	// SearchTimeoutParameter description: Enables or disables the `timeout:` parameter in searches.
	SearchTimeoutParameter string `json:"searchTimeoutParameter,omitempty"`
}
type GitHubConnection struct {
	// Certificate description: TLS certificate of a GitHub Enterprise instance.
	Certificate string `json:"certificate,omitempty"`
	// GitURLType description: The type of Git URLs to use for cloning and fetching Git repositories on this GitHub instance.
	//
	// If "http", Sourcegraph will access GitLab repositories using Git URLs of the form http(s)://github.com/myteam/myproject.git (using https: if the GitHub instance uses HTTPS).
	//
	// If "ssh", Sourcegraph will access GitHub repositories using Git URLs of the form git@github.com:myteam/myproject.git. See the documentation for how to provide SSH private keys and known_hosts: https://about.sourcegraph.com/docs/config/repositories#repositories-that-need-https-or-ssh-authentication.
	GitURLType string `json:"gitURLType,omitempty"`
	// InitialRepositoryEnablement description: Defines whether repositories from this GitHub instance should be enabled and cloned when they are first seen by Sourcegraph. If false, the site admin must explicitly enable GitHub repositories (in the site admin area) to clone them and make them searchable on Sourcegraph. If true, they will be enabled and cloned immediately (subject to rate limiting by GitHub); site admins can still disable them explicitly, and they'll remain disabled.
	InitialRepositoryEnablement bool `json:"initialRepositoryEnablement,omitempty"`
	// PreemptivelyClone description: Preemptively clone GitHub repositories added (instead of cloning on-demand when the repository is searched or viewed)
	//
	// DEPRECATED: Use initialRepositoryEnablement instead.
	PreemptivelyClone bool `json:"preemptivelyClone,omitempty"`
	// Repos description: An array of repository "owner/name" strings specifying which GitHub or GitHub Enterprise repositories to mirror on Sourcegraph.
	Repos []string `json:"repos,omitempty"`
	// RepositoryPathPattern description: The pattern used to generate a the corresponding Sourcegraph repository path for a GitHub or GitHub Enterprise repository. In the pattern, the variable "{host}" is replaced with the GitHub host (such as github.example.com), and "{nameWithOwner}" is replaced with the GitHub repository's "owner/path" (such as "myorg/myrepo").
	//
	// For example, if your GitHub Enterprise URL is https://github.example.com and your Sourcegraph URL is https://src.example.com, then a repositoryPathPattern of "{host}/{nameWithOwner}" would mean that a GitHub repository at https://github.example.com/myorg/myrepo is available on Sourcegraph at https://src.example.com/github.example.com/myorg/myrepo.
	RepositoryPathPattern string `json:"repositoryPathPattern,omitempty"`
	// RepositoryQuery description: An array of strings specifying which GitHub or GitHub Enterprise repositories to mirror on Sourcegraph. The valid values are:
	//
	// - `public` mirrors all public repositories for GitHub Enterprise and is the equivalent of `none` for GitHub
	//
	// - `affiliated` mirrors all repositories affiliated with the configured token's user:
	// 	- Private repositories with read access
	// 	- Public repositories owned by the user or their orgs
	// 	- Public repositories with write access
	//
	// - `none` mirrors no repositories (except those specified in the `repos` configuration property or added manually)
	//
	// If multiple values are provided, their results are unioned.
	//
	// If you need to narrow the set of mirrored repositories further (and don't want to enumerate the set in the "repos" configuration property), create a new bot/machine user on GitHub or GitHub Enterprise that is only affiliated with the desired repositories.
	RepositoryQuery []string `json:"repositoryQuery,omitempty"`
	// Token description: A GitHub personal access token with repo and org scope.
	Token string `json:"token"`
	// Url description: URL of a GitHub instance, such as https://github.com or https://github-enterprise.example.com
	Url string `json:"url,omitempty"`
}
type GitLabConnection struct {
	// Certificate description: TLS certificate of a GitLab instance.
	Certificate string `json:"certificate,omitempty"`
	// GitURLType description: The type of Git URLs to use for cloning and fetching Git repositories on this GitLab instance.
	//
	// If "http", Sourcegraph will access GitLab repositories using Git URLs of the form http(s)://gitlab.example.com/myteam/myproject.git (using https: if the GitLab instance uses HTTPS).
	//
	// If "ssh", Sourcegraph will access GitLab repositories using Git URLs of the form git@example.gitlab.com:myteam/myproject.git. See the documentation for how to provide SSH private keys and known_hosts: https://about.sourcegraph.com/docs/config/repositories#repositories-that-need-https-or-ssh-authentication.
	GitURLType string `json:"gitURLType,omitempty"`
	// InitialRepositoryEnablement description: Defines whether repositories from this GitLab instance should be enabled and cloned when they are first seen by Sourcegraph. If false, the site admin must explicitly enable GitLab repositories (in the site admin area) to clone them and make them searchable on Sourcegraph. If true, they will be enabled and cloned immediately (subject to rate limiting by GitLab); site admins can still disable them explicitly, and they'll remain disabled.
	InitialRepositoryEnablement bool `json:"initialRepositoryEnablement,omitempty"`
	// ProjectQuery description: An array of strings specifying which GitLab projects to mirror on Sourcegraph. Each string is a URL query string for the GitLab projects API, such as "?membership=true&search=foo".
	//
	// The query string is passed directly to GitLab to retrieve the list of projects. The special string "none" can be used as the only element to disable this feature. Projects matched by multiple query strings are only imported once. See https://docs.gitlab.com/ee/api/projects.html#list-all-projects for available query string options.
	ProjectQuery []string `json:"projectQuery,omitempty"`
	// RepositoryPathPattern description: The pattern used to generate a the corresponding Sourcegraph repository path for a GitLab project. In the pattern, the variable "{host}" is replaced with the GitLab URL's host (such as gitlab.example.com), and "{pathWithNamespace}" is replaced with the GitLab project's "namespace/path" (such as "myteam/myproject").
	//
	// For example, if your GitLab is https://gitlab.example.com and your Sourcegraph is https://src.example.com, then a repositoryPathPattern of "{host}/{pathWithNamespace}" would mean that a GitLab project at https://gitlab.example.com/myteam/myproject is available on Sourcegraph at https://src.example.com/gitlab.example.com/myteam/myproject.
	RepositoryPathPattern string `json:"repositoryPathPattern,omitempty"`
	// Token description: A GitLab personal access token with "api" scope.
	Token string `json:"token"`
	// Url description: URL of a GitLab instance, such as https://gitlab.example.com or (for GitLab.com) https://gitlab.com
	Url string `json:"url"`
}
type GitoliteConnection struct {
	// Blacklist description: Regular expression to filter repositories from auto-discovery, so they will not get cloned automatically.
	Blacklist string `json:"blacklist,omitempty"`
	// Host description: Gitolite host that stores the repositories (e.g., git@gitolite.example.com).
	Host string `json:"host"`
	// PhabricatorMetadataCommand description: Bash command that prints out the Phabricator callsign for a Gitolite repository. This will be run with environment variable $REPO set to the URI of the repository and used to obtain the Phabricator metadata for a Gitolite repository. (Note: this requires `bash` to be installed.)
	PhabricatorMetadataCommand string `json:"phabricatorMetadataCommand,omitempty"`
	// Prefix description: Repository URI prefix that will map to this Gitolite host. This should likely end with a trailing slash. E.g., "gitolite.example.com/".
	Prefix string `json:"prefix"`
}

// HTTPHeaderAuthProvider description: Configures the HTTP header authentication provider (which authenticates users by consulting an HTTP request header set by an authentication proxy such as https://github.com/bitly/oauth2_proxy).
type HTTPHeaderAuthProvider struct {
	Type string `json:"type"`
	// UsernameHeader description: The name (case-insensitive) of an HTTP header whose value is taken to be the username of the client requesting the page. Set this value when using an HTTP proxy that authenticates requests, and you don't want the extra configurability of the other authentication methods.
	//
	// Requires auth.provider=="http-header".
	UsernameHeader string `json:"usernameHeader"`
}
type Langservers struct {
	// Address description: TCP address of the language server. Required for Sourcegraph Server; do not set for Sourcegraph Data Center.
	Address string `json:"address,omitempty"`
	// Disabled description: Whether or not this language server is disabled.
	Disabled bool `json:"disabled,omitempty"`
	// InitializationOptions description: LSP initialization options. This object will be set as the `initializationOptions` field in LSP initialize requests (https://microsoft.github.io/language-server-protocol/specification#initialize).
	InitializationOptions map[string]any `json:"initializationOptions,omitempty"`
	// Language description: Name of the language mode for the langserver (e.g. go, java)
	Language string `json:"language"`
	// Metadata description: Language server metadata. Used to populate various UI elements.
	Metadata *Metadata `json:"metadata,omitempty"`
}
type Links struct {
	// Blob description: URL template for specifying how to link to files at an external location. Use "{path}" as a placeholder for a given path and "{rev}" as a placeholder for a given revision e.g. "https://example.com/myrepo@{rev}/browse/{path}"
	Blob string `json:"blob,omitempty"`
	// Commit description: URL template for specifying how to link to commits at an external location. Use "{commit}" as a placeholder for a given commit ID e.g. "https://example.com/myrepo/view-commit/{commit}"
	Commit string `json:"commit,omitempty"`
	// Repository description: URL specifying where to view the repository at an external location e.g. "https://example.com/myrepo"
	Repository string `json:"repository,omitempty"`
	// Tree description: URL template for specifying how to link to paths at an external location. Use "{path}" as a placeholder for a given path and "{rev}" as a placeholder for a given revision e.g. "https://example.com/myrepo@{rev}/browse/{path}"
	Tree string `json:"tree,omitempty"`
}

// Metadata description: Language server metadata. Used to populate various UI elements.
type Metadata struct {
	// DocsURL description: URL to the language server's documentation, if available.
	DocsURL string `json:"docsURL,omitempty"`
	// Experimental description: Whether or not this language server should be considered experimental. Has no effect on behavior, only effects how the language server is presented e.g. in the UI.
	Experimental bool `json:"experimental,omitempty"`
	// HomepageURL description: URL to the language server's homepage, if available.
	HomepageURL string `json:"homepageURL,omitempty"`
	// IssuesURL description: URL to the language server's open/known issues, if available.
	IssuesURL string `json:"issuesURL,omitempty"`
}

// OpenIDConnectAuthProvider description: Configures the OpenID Connect authentication provider for SSO.
type OpenIDConnectAuthProvider struct {
	// ClientID description: The client ID for the OpenID Connect client for this site.
	//
	// For Google Apps: obtain this value from the API console (https://console.developers.google.com), as described at https://developers.google.com/identity/protocols/OpenIDConnect#getcredentials
	ClientID string `json:"clientID"`
	// ClientSecret description: The client secret for the OpenID Connect client for this site.
	//
	// For Google Apps: obtain this value from the API console (https://console.developers.google.com), as described at https://developers.google.com/identity/protocols/OpenIDConnect#getcredentials
	ClientSecret string `json:"clientSecret"`
	// Issuer description: The URL of the OpenID Connect issuer.
	//
	// For Google Apps: https://accounts.google.com
	Issuer string `json:"issuer"`
	// OverrideToken description: (For testing and development only) A token used to circumvent the OpenID Connect authentication layer.
	//
	// DEPRECATED: This override-auth feature will be removed, and no replacement is planned.
	OverrideToken string `json:"overrideToken,omitempty"`
	// RequireEmailDomain description: Only allow users to authenticate if their email domain is equal to this value (example: mycompany.com). Do not include a leading "@". If not set, all users on this OpenID Connect provider can authenticate to Sourcegraph.
	RequireEmailDomain string `json:"requireEmailDomain,omitempty"`
	Type               string `json:"type"`
}
type Phabricator struct {
	// Repos description: The list of repositories available on Phabricator.
	Repos []*Repos `json:"repos,omitempty"`
	// Token description: API token for the Phabricator instance.
	Token string `json:"token,omitempty"`
	// Url description: URL of a Phabricator instance, such as https://phabricator.example.com
	Url string `json:"url,omitempty"`
}
type Repos struct {
	// Callsign description: The unique Phabricator identifier for the repository, like 'MUX'.
	Callsign string `json:"callsign"`
	// Path description: Display path for the url e.g. gitolite/my/repo
	Path string `json:"path"`
}
type Repository struct {
	Links *Links `json:"links,omitempty"`
	// Path description: Display path on Sourcegraph for the repository, such as my/repo
	Path string `json:"path"`
	// Type description: Type of the version control system for this repository, such as "git"
	Type string `json:"type,omitempty"`
	// Url description: Clone URL for the repository, such as git@example.com:my/repo.git
	Url string `json:"url"`
}

// SAMLAuthProvider description: Configures the SAML authentication provider for SSO.
type SAMLAuthProvider struct {
	// IdentityProviderMetadata description: SAML Identity Provider metadata XML contents (for static configuration of the SAML Service Provider). The value of this field should be an XML document whose root element is `<EntityDescriptor>`.
	IdentityProviderMetadata string `json:"identityProviderMetadata,omitempty"`
	// IdentityProviderMetadataURL description: SAML Identity Provider metadata URL (for dynamic configuration of the SAML Service Provider).
	IdentityProviderMetadataURL string `json:"identityProviderMetadataURL,omitempty"`
	// ServiceProviderCertificate description: SAML Service Provider certificate in X.509 encoding (begins with "-----BEGIN CERTIFICATE-----").
	ServiceProviderCertificate string `json:"serviceProviderCertificate"`
	// ServiceProviderPrivateKey description: SAML Service Provider private key in PKCS#8 encoding (begins with "-----BEGIN PRIVATE KEY-----").
	ServiceProviderPrivateKey string `json:"serviceProviderPrivateKey"`
	Type                      string `json:"type"`
}

// SMTPServerConfig description: The SMTP server used to send transactional emails (such as email verifications, reset-password emails, and notifications).
type SMTPServerConfig struct {
	// Authentication description: The type of authentication to use for the SMTP server.
	Authentication string `json:"authentication"`
	// Domain description: The HELO domain to provide to the SMTP server (if needed).
	Domain string `json:"domain,omitempty"`
	// Host description: The SMTP server host.
	Host string `json:"host"`
	// Password description: The username to use when communicating with the SMTP server.
	Password string `json:"password,omitempty"`
	// Port description: The SMTP server port.
	Port int `json:"port"`
	// Username description: The username to use when communicating with the SMTP server.
	Username string `json:"username,omitempty"`
}
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
	Additional   map[string]any `json:"-"` // additionalProperties not explicitly defined in the schema
}

func (v Settings) MarshalJSON() ([]byte, error) {
	m := make(map[string]any, len(v.Additional))
	for k, v := range v.Additional {
		m[k] = v
	}
	type wrapper Settings
	b, err := json.Marshal(wrapper(v))
	if err != nil {
		return nil, err
	}
	var m2 map[string]any
	if err := json.Unmarshal(b, &m2); err != nil {
		return nil, err
	}
	for k, v := range m2 {
		m[k] = v
	}
	return json.Marshal(m)
}
func (v *Settings) UnmarshalJSON(data []byte) error {
	type wrapper Settings
	var s wrapper
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*v = Settings(s)
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	delete(m, "motd")
	delete(m, "notifications.slack")
	delete(m, "search.repositoryGroups")
	delete(m, "search.savedQueries")
	delete(m, "search.scopes")
	if len(m) > 0 {
		v.Additional = make(map[string]any, len(m))
	}
	for k, vv := range m {
		v.Additional[k] = vv
	}
	return nil
}

// SiteConfiguration description: Configuration for a Sourcegraph site.
type SiteConfiguration struct {
	// AppURL description: Publicly accessible URL to web app (e.g., what you type into your browser).
	AppURL string `json:"appURL,omitempty"`
	// AuthAllowSignup description: Allows new visitors to sign up for accounts. The sign-up page will be enabled and accessible to all visitors.
	//
	// SECURITY: If the site has no users (i.e., during initial setup), it will always allow the first user to sign up and become site admin **without any approval** (first user to sign up becomes the admin).
	//
	// Requires auth.provider == "builtin"
	//
	// DEPRECATED: Use "auth.providers" with an entry of the form {"type": "builtin", "allowSignup": true} instead.
	AuthAllowSignup bool `json:"auth.allowSignup,omitempty"`
	// AuthDisableAccessTokens description: Prevents users from creating access tokens, which enable external tools to access the Sourcegraph API with the privileges of the user.
	AuthDisableAccessTokens bool                       `json:"auth.disableAccessTokens,omitempty"`
	AuthOpenIDConnect       *OpenIDConnectAuthProvider `json:"auth.openIDConnect,omitempty"`
	// AuthProvider description: The authentication provider to use for identifying and signing in users. Defaults to "builtin" authentication.
	//
	// DEPRECATED: Use "auth.providers" instead. During the deprecation period (before this property is removed), provider set here will be added as an entry in "auth.providers".
	AuthProvider string `json:"auth.provider,omitempty"`
	// AuthProviders description: The authentication providers to use for identifying and signing in users.
	//
	// Only one authentication provider is supported. If you set the deprecated field "auth.provider", then that value is used as the authentication provider, and you can't set another one here.
	AuthProviders []AuthProviders `json:"auth.providers,omitempty"`
	// AuthPublic description: Allows anonymous visitors full read access to repositories, code files, search, and other data (except site configuration).
	//
	// SECURITY WARNING: If you enable this, you must ensure that only authorized users can access the server (using firewall rules or an external proxy, for example).
	//
	// Requires usage of the builtin authentication provider.
	AuthPublic bool              `json:"auth.public,omitempty"`
	AuthSaml   *SAMLAuthProvider `json:"auth.saml,omitempty"`
	// AuthUserIdentityHTTPHeader description: The name (case-insensitive) of an HTTP header whose value is taken to be the username of the client requesting the page. Set this value when using an HTTP proxy that authenticates requests, and you don't want the extra configurability of the other authentication methods.
	//
	// Requires auth.provider=="http-header".
	//
	// DEPRECATED: Use "auth.providers" with an entry of the form {"type": "http-header", "usernameHeader": "..."} instead.
	AuthUserIdentityHTTPHeader string `json:"auth.userIdentityHTTPHeader,omitempty"`
	// AuthUserOrgMap description: Ensure that matching users are members of the specified orgs (auto-joining users to the orgs if they are not already a member). Provide a JSON object of the form `{"*": ["org1", "org2"]}`, where org1 and org2 are orgs that all users are automatically joined to. Currently the only supported key is `"*"`.
	AuthUserOrgMap map[string][]string `json:"auth.userOrgMap,omitempty"`
	// AwsCodeCommit description: JSON array of configuration for AWS CodeCommit endpoints.
	AwsCodeCommit []*AWSCodeCommitConnection `json:"awsCodeCommit,omitempty"`
	// BitbucketServer description: JSON array of configuration for Bitbucket Server hosts.
	BitbucketServer []*BitbucketServerConnection `json:"bitbucketServer,omitempty"`
	// BlacklistGoGet description: List of domains to blacklist dependency fetching from. Separated by ','.
	//
	// Unlike `noGoGetDomains` (which tries to use a hueristic to determine where to clone the dependencies from), this option outright prevents fetching of dependencies with the given domain name. This will prevent code intelligence from working on these dependencies, so most users should not use this option.
	BlacklistGoGet []string `json:"blacklistGoGet,omitempty"`
	// CorsOrigin description: Value for the Access-Control-Allow-Origin header returned with all requests.
	CorsOrigin string `json:"corsOrigin,omitempty"`
	// DisableAutoGitUpdates description: Disable periodically fetching git contents for existing repositories.
	DisableAutoGitUpdates bool `json:"disableAutoGitUpdates,omitempty"`
	// DisableBrowserExtension description: Disable incoming connections from the Sourcegraph browser extension.
	DisableBrowserExtension bool `json:"disableBrowserExtension,omitempty"`
	// DisableBuiltInSearches description: Whether built-in searches should be hidden on the Searches page.
	DisableBuiltInSearches bool `json:"disableBuiltInSearches,omitempty"`
	// DisableExampleSearches description: (Deprecated: use disableBuiltInSearches) Whether built-in searches should be hidden on the Searches page.
	DisableExampleSearches bool `json:"disableExampleSearches,omitempty"`
	// DisablePublicRepoRedirects description: Disable redirects to sourcegraph.com when visiting public repositories that can't exist on this server.
	DisablePublicRepoRedirects bool `json:"disablePublicRepoRedirects,omitempty"`
	// DisableTelemetry description: (Deprecated: event-level telemetry is now always disabled) Prevent usage data from being sent back to Sourcegraph (no private code is sent and URLs are sanitized to prevent leakage of private data).
	DisableTelemetry bool `json:"disableTelemetry,omitempty"`
	// DontIncludeSymbolResultsByDefault description: Set to `true` to not include symbol results if no `type:` filter was given
	DontIncludeSymbolResultsByDefault bool `json:"dontIncludeSymbolResultsByDefault,omitempty"`
	// EmailAddress description: The "from" address for emails sent by this server.
	EmailAddress string            `json:"email.address,omitempty"`
	EmailSmtp    *SMTPServerConfig `json:"email.smtp,omitempty"`
	// ExecuteGradleOriginalRootPaths description: Java: A comma-delimited list of patterns that selects repository revisions for which to execute Gradle scripts, rather than extracting Gradle metadata statically. **Security note:** these should be restricted to repositories within your own organization. A percent sign ('%') can be used to prefix-match. For example, `git://my.internal.host/org1/%,git://my.internal.host/org2/repoA?%` would select all revisions of all repositories in org1 and all revisions of repoA in org2.
	// Note: this field is misnamed, as it matches against the originalRootURI LSP initialize parameter, rather than the no-longer-used originalRootPath parameter.
	ExecuteGradleOriginalRootPaths string `json:"executeGradleOriginalRootPaths,omitempty"`
	// ExperimentalFeatures description: Experimental features to enable or disable. Features that are now enabled by default are marked as deprecated.
	ExperimentalFeatures *ExperimentalFeatures `json:"experimentalFeatures,omitempty"`
	// GitMaxConcurrentClones description: Maximum number of git clone processes that will be run concurrently to update repositories.
	GitMaxConcurrentClones int `json:"gitMaxConcurrentClones,omitempty"`
	// GitOriginMap description: Space separated list of mappings from repo name prefix to origin url, for example "github.com/!https://github.com/%.git".
	GitOriginMap string `json:"gitOriginMap,omitempty"`
	// Github description: JSON array of configuration for GitHub hosts. See GitHub Configuration section for more information.
	Github []*GitHubConnection `json:"github,omitempty"`
	// GithubClientID description: Client ID for GitHub.
	GithubClientID string `json:"githubClientID,omitempty"`
	// GithubClientSecret description: Client secret for GitHub.
	GithubClientSecret string `json:"githubClientSecret,omitempty"`
	// GithubEnterpriseAccessToken description: (Deprecated: Use GitHub) Access token to authenticate to GitHub Enterprise API.
	GithubEnterpriseAccessToken string `json:"githubEnterpriseAccessToken,omitempty"`
	// GithubEnterpriseCert description: (Deprecated: Use GitHub) TLS certificate of GitHub Enterprise instance, if from a CA that's not part of the standard certificate chain.
	GithubEnterpriseCert string `json:"githubEnterpriseCert,omitempty"`
	// GithubEnterpriseURL description: (Deprecated: Use GitHub) URL of GitHub Enterprise instance from which to sync repositories.
	GithubEnterpriseURL string `json:"githubEnterpriseURL,omitempty"`
	// GithubPersonalAccessToken description: (Deprecated: Use GitHub) Personal access token for GitHub.
	GithubPersonalAccessToken string `json:"githubPersonalAccessToken,omitempty"`
	// Gitlab description: JSON array of configuration for GitLab hosts.
	Gitlab []*GitLabConnection `json:"gitlab,omitempty"`
	// Gitolite description: JSON array of configuration for Gitolite hosts.
	Gitolite []*GitoliteConnection `json:"gitolite,omitempty"`
	// HtmlBodyBottom description: HTML to inject at the bottom of the `<body>` element on each page, for analytics scripts
	HtmlBodyBottom string `json:"htmlBodyBottom,omitempty"`
	// HtmlBodyTop description: HTML to inject at the top of the `<body>` element on each page, for analytics scripts
	HtmlBodyTop string `json:"htmlBodyTop,omitempty"`
	// HtmlHeadBottom description: HTML to inject at the bottom of the `<head>` element on each page, for analytics scripts
	HtmlHeadBottom string `json:"htmlHeadBottom,omitempty"`
	// HtmlHeadTop description: HTML to inject at the top of the `<head>` element on each page, for analytics scripts
	HtmlHeadTop string `json:"htmlHeadTop,omitempty"`
	// HttpStrictTransportSecurity description: The value of the Strict-Transport-Security HTTP header sent by Sourcegraph, if non-empty
	HttpStrictTransportSecurity string `json:"httpStrictTransportSecurity,omitempty"`
	// HttpToHttpsRedirect description: Redirect users from HTTP to HTTPS. Accepted values are "on", "off", and "load-balanced" (boolean values true and false are also accepted and equivalent to "on" and "off" respectively). If "load-balanced" then additionally we use "X-Forwarded-Proto" to determine if on HTTP.
	HttpToHttpsRedirect any `json:"httpToHttpsRedirect,omitempty"`
	// Langservers description: Language server configuration.
	Langservers []*Langservers `json:"langservers,omitempty"`
	// LightstepAccessToken description: Access token for sending traces to LightStep.
	LightstepAccessToken string `json:"lightstepAccessToken,omitempty"`
	// LightstepProject description: The project ID on LightStep that corresponds to the `lightstepAccessToken`, only for generating links to traces. For example, if `lightstepProject` is `mycompany-prod`, all HTTP responses from Sourcegraph will include an X-Trace header with the URL to the trace on LightStep, of the form `https://app.lightstep.com/mycompany-prod/trace?span_guid=...&at_micros=...`.
	LightstepProject string `json:"lightstepProject,omitempty"`
	// MaxReposToSearch description: The maximum number of repositories to search across. The user is prompted to narrow their query if exceeded. The value -1 means unlimited.
	MaxReposToSearch int `json:"maxReposToSearch,omitempty"`
	// NoGoGetDomains description: List of domains in import paths to NOT perform `go get` on, but instead treat as standard Git repositories. Separated by ','.
	//
	// For example, if your code imports non-go-gettable packages like `"mygitolite.aws.me.org/mux.git/subpkg"` you may set this option to `"mygitolite.aws.me.org"` and Sourcegraph will effectively run `git clone mygitolite.aws.me.org/mux.git` instead of performing the usual `go get` dependency resolution behavior.
	NoGoGetDomains string `json:"noGoGetDomains,omitempty"`
	// OidcClientID description: OIDC Client ID
	//
	// DEPRECATED: Use auth.provider=="openidconnect" and auth.openidconnect's "clientID" property instead.
	OidcClientID string `json:"oidcClientID,omitempty"`
	// OidcClientSecret description: OIDC Client Secret
	//
	// DEPRECATED: Use auth.provider=="openidconnect" and auth.openidconnect's "clientSecret" property instead.
	OidcClientSecret string `json:"oidcClientSecret,omitempty"`
	// OidcEmailDomain description: Whitelisted email domain for logins, e.g. 'mycompany.com'
	//
	// DEPRECATED: Use auth.provider=="openidconnect" and auth.openidconnect's "requireEmailDomain" property instead.
	OidcEmailDomain string `json:"oidcEmailDomain,omitempty"`
	// OidcProvider description: The URL of the OpenID Connect Provider
	//
	// DEPRECATED: Use auth.provider=="openidconnect" and auth.openidconnect's "issuer" property instead.
	OidcProvider string `json:"oidcProvider,omitempty"`
	// Phabricator description: JSON array of configuration for Phabricator hosts. See Phabricator Configuration section for more information.
	Phabricator []*Phabricator `json:"phabricator,omitempty"`
	// PhabricatorURL description: (Deprecated: Use Phabricator) URL of Phabricator instance.
	PhabricatorURL string `json:"phabricatorURL,omitempty"`
	// PrivateArtifactRepoID description: Java: Private artifact repository ID in your build files. If you do not explicitly include the private artifact repository, then set this to some unique string (e.g,. "my-repository").
	PrivateArtifactRepoID string `json:"privateArtifactRepoID,omitempty"`
	// PrivateArtifactRepoPassword description: Java: The password to authenticate to the private Artifactory.
	PrivateArtifactRepoPassword string `json:"privateArtifactRepoPassword,omitempty"`
	// PrivateArtifactRepoURL description: Java: The URL that corresponds to privateArtifactRepoID (e.g., http://my.artifactory.local/artifactory/root).
	PrivateArtifactRepoURL string `json:"privateArtifactRepoURL,omitempty"`
	// PrivateArtifactRepoUsername description: Java: The username to authenticate to the private Artifactory.
	PrivateArtifactRepoUsername string `json:"privateArtifactRepoUsername,omitempty"`
	// RepoListUpdateInterval description: Interval (in minutes) for checking code hosts (such as GitHub, Gitolite, etc.) for new repositories.
	RepoListUpdateInterval int `json:"repoListUpdateInterval,omitempty"`
	// ReposList description: JSON array of configuration for external repositories.
	ReposList []*Repository `json:"repos.list,omitempty"`
	// SamlIDProviderMetadataURL description: SAML Identity Provider metadata URL (for dyanmic configuration of SAML Service Provider)
	//
	// DEPRECATED: Use auth.provider=="saml" and auth.saml's "identityProviderMetadataURL" property instead.
	SamlIDProviderMetadataURL string `json:"samlIDProviderMetadataURL,omitempty"`
	// SamlSPCert description: SAML Service Provider certificate
	//
	// DEPRECATED: Use auth.provider=="saml" and auth.saml's "serviceProviderCertificate" property instead.
	SamlSPCert string `json:"samlSPCert,omitempty"`
	// SamlSPKey description: SAML Service Provider private key
	//
	// DEPRECATED: Use auth.provider=="saml" and auth.saml's "serviceProviderPrivateKey" property instead.
	SamlSPKey string `json:"samlSPKey,omitempty"`
	// SearchScopes description: JSON array of custom search scopes (e.g., [{"name":"Text Files","value":"file:\.txt$"}]).
	//
	// DEPRECATED: Values should be moved to the "settings" field's "search.scopes" property.
	SearchScopes []*SearchScope `json:"searchScopes,omitempty"`
	// SecretKey description: A base64-encoded secret key for this site, used for generating links to invite users to organizations. If you have `openssl` installed, you can generate a valid key by running `openssl rand -base64 32`.
	SecretKey string `json:"secretKey,omitempty"`
	// Settings description: Site settings. Organization and user settings override site settings.
	Settings *Settings `json:"settings,omitempty"`
	// SiteID description: The identifier for this site. A Sourcegraph site is a collection of one or more Sourcegraph instances that are all part of the same logical site. If the site ID is not set here, it is stored in the database the first time the server is run.
	SiteID string `json:"siteID,omitempty"`
	// TlsLetsencrypt description: Toggles ACME functionality for automatically using a TLS certificate issued by the Let's Encrypt Certificate Authority.
	// The default value is auto, which uses the following conditions to switch on:
	//  - tlsCert and tlsKey are unset.
	//  - appURL is a https:// URL
	//  - Can successfully bind to port 443
	TlsLetsencrypt string `json:"tls.letsencrypt,omitempty"`
	// TlsCert description: The contents of the PEM-encoded TLS certificate for the external HTTP server (for the web app and API).
	//
	// See https://about.sourcegraph.com/docs/config/tlsssl/ for more information.
	TlsCert string `json:"tlsCert,omitempty"`
	// TlsKey description: The contents of the PEM-encoded TLS key for the external HTTP server (for the web app and API).
	//
	// See https://about.sourcegraph.com/docs/config/tlsssl/ for more information.
	TlsKey string `json:"tlsKey,omitempty"`
	// UpdateChannel description: The channel on which to automatically check for Sourcegraph updates.
	UpdateChannel string `json:"update.channel,omitempty"`
	// UseJaeger description: Use local Jaeger instance for tracing. Data Center only.
	//
	// After enabling Jaeger and updating your Kubernetes cluster, `kubectl get pods`
	// should display pods prefixed with `jaeger-cassandra`,
	// `jaeger-collector`, and `jaeger-query`. `jaeger-collector` will start
	// crashing until you initialize the Cassandra DB. To do so, do the
	// following:
	//
	// 1. Install [`cqlsh`](https://pypi.python.org/pypi/cqlsh).
	// 1. `kubectl port-forward $(kubectl get pods | grep jaeger-cassandra | awk '{ print $1 }') 9042`
	// 1. `git clone https://github.com/uber/jaeger && cd jaeger && MODE=test ./plugin/storage/cassandra/schema/create.sh | cqlsh`
	// 1. `kubectl port-forward $(kubectl get pods | grep jaeger-query | awk '{ print $1 }') 16686`
	// 1. Go to http://localhost:16686 to view the Jaeger dashboard.
	UseJaeger bool `json:"useJaeger,omitempty"`
}

// SlackNotificationsConfig description: Configuration for sending notifications to Slack.
type SlackNotificationsConfig struct {
	// WebhookURL description: The Slack webhook URL used to post notification messages to a Slack channel. To obtain this URL, go to: https://YOUR-WORKSPACE-NAME.slack.com/apps/new/A0F7XDUAZ-incoming-webhooks
	WebhookURL string `json:"webhookURL"`
}
