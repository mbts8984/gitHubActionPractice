package types

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

type rateLimitCategory uint8

const (
	coreCategory rateLimitCategory = iota
	searchCategory

	categories // An array of this length will be able to contain all rate limit categories.
)

type ActionsService service
type ActivityService service
type AdminService service
type AppsService service
type AuthorizationsService service
type BillingService service
type ChecksService service
type CodeScanningService service
type EnterpriseService service
type GistsService service
type GitService service
type GitignoresService service
type InteractionsService service
type IssueImportService service
type IssuesService service
type LicensesService service
type MarketplaceService service
type MigrationService service
type OrganizationsService service
type ProjectsService service
type ReactionsService service
type RepositoriesService service
type SearchService service
type TeamsService service
type UsersService service

type Rate struct {
	// The number of requests per hour the client is currently limited to.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this hour.
	Remaining int `json:"remaining"`

	// The time at which the current rate limit will reset.
	Reset Timestamp `json:"reset"`
}

type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// Base URL for uploading files.
	UploadURL *url.URL

	// User agent used when communicating with the GitHub API.
	UserAgent string

	rateMu     sync.Mutex
	rateLimits [categories]Rate // Rate limits for the client as determined by the most recent API calls.

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the GitHub API.
	Actions        *ActionsService
	Activity       *ActivityService
	Admin          *AdminService
	Apps           *AppsService
	Authorizations *AuthorizationsService
	Billing        *BillingService
	Checks         *ChecksService
	CodeScanning   *CodeScanningService
	Enterprise     *EnterpriseService
	Gists          *GistsService
	Git            *GitService
	Gitignores     *GitignoresService
	Interactions   *InteractionsService
	IssueImport    *IssueImportService
	Issues         *IssuesService
	Licenses       *LicensesService
	Marketplace    *MarketplaceService
	Migrations     *MigrationService
	Organizations  *OrganizationsService
	Projects       *ProjectsService
	PullRequests   *PullRequestsService
	Reactions      *ReactionsService
	Repositories   *RepositoriesService
	Search         *SearchService
	Teams          *TeamsService
	Users          *UsersService
}

type service struct {
	client *Client
}
type PullRequestsService service
type CodeOfConduct struct {
	Name *string `json:"name,omitempty"`
	Key  *string `json:"key,omitempty"`
	URL  *string `json:"url,omitempty"`
	Body *string `json:"body,omitempty"`
}
type License struct {
	Key  *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`

	SPDXID         *string   `json:"spdx_id,omitempty"`
	HTMLURL        *string   `json:"html_url,omitempty"`
	Featured       *bool     `json:"featured,omitempty"`
	Description    *string   `json:"description,omitempty"`
	Implementation *string   `json:"implementation,omitempty"`
	Permissions    *[]string `json:"permissions,omitempty"`
	Conditions     *[]string `json:"conditions,omitempty"`
	Limitations    *[]string `json:"limitations,omitempty"`
	Body           *string   `json:"body,omitempty"`
}

type Repository struct {
	ID                  *int64          `json:"id,omitempty"`
	NodeID              *string         `json:"node_id,omitempty"`
	Owner               *User           `json:"owner,omitempty"`
	Name                *string         `json:"name,omitempty"`
	FullName            *string         `json:"full_name,omitempty"`
	Description         *string         `json:"description,omitempty"`
	Homepage            *string         `json:"homepage,omitempty"`
	CodeOfConduct       *CodeOfConduct  `json:"code_of_conduct,omitempty"`
	DefaultBranch       *string         `json:"default_branch,omitempty"`
	MasterBranch        *string         `json:"master_branch,omitempty"`
	CreatedAt           *Timestamp      `json:"created_at,omitempty"`
	PushedAt            *Timestamp      `json:"pushed_at,omitempty"`
	UpdatedAt           *Timestamp      `json:"updated_at,omitempty"`
	HTMLURL             *string         `json:"html_url,omitempty"`
	CloneURL            *string         `json:"clone_url,omitempty"`
	GitURL              *string         `json:"git_url,omitempty"`
	MirrorURL           *string         `json:"mirror_url,omitempty"`
	SSHURL              *string         `json:"ssh_url,omitempty"`
	SVNURL              *string         `json:"svn_url,omitempty"`
	Language            *string         `json:"language,omitempty"`
	Fork                *bool           `json:"fork,omitempty"`
	ForksCount          *int            `json:"forks_count,omitempty"`
	NetworkCount        *int            `json:"network_count,omitempty"`
	OpenIssuesCount     *int            `json:"open_issues_count,omitempty"`
	OpenIssues          *int            `json:"open_issues,omitempty"` // Deprecated: Replaced by OpenIssuesCount. For backward compatibility OpenIssues is still populated.
	StargazersCount     *int            `json:"stargazers_count,omitempty"`
	SubscribersCount    *int            `json:"subscribers_count,omitempty"`
	WatchersCount       *int            `json:"watchers_count,omitempty"` // Deprecated: Replaced by StargazersCount. For backward compatibility WatchersCount is still populated.
	Watchers            *int            `json:"watchers,omitempty"`       // Deprecated: Replaced by StargazersCount. For backward compatibility Watchers is still populated.
	Size                *int            `json:"size,omitempty"`
	AutoInit            *bool           `json:"auto_init,omitempty"`
	Parent              *Repository     `json:"parent,omitempty"`
	Source              *Repository     `json:"source,omitempty"`
	TemplateRepository  *Repository     `json:"template_repository,omitempty"`
	Organization        *Organization   `json:"organization,omitempty"`
	Permissions         map[string]bool `json:"permissions,omitempty"`
	AllowRebaseMerge    *bool           `json:"allow_rebase_merge,omitempty"`
	AllowSquashMerge    *bool           `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit    *bool           `json:"allow_merge_commit,omitempty"`
	DeleteBranchOnMerge *bool           `json:"delete_branch_on_merge,omitempty"`
	Topics              []string        `json:"topics,omitempty"`
	Archived            *bool           `json:"archived,omitempty"`
	Disabled            *bool           `json:"disabled,omitempty"`

	// Only provided when using RepositoriesService.Get while in preview
	License *License `json:"license,omitempty"`

	// Additional mutable fields when creating and editing a repository
	Private           *bool   `json:"private,omitempty"`
	HasIssues         *bool   `json:"has_issues,omitempty"`
	HasWiki           *bool   `json:"has_wiki,omitempty"`
	HasPages          *bool   `json:"has_pages,omitempty"`
	HasProjects       *bool   `json:"has_projects,omitempty"`
	HasDownloads      *bool   `json:"has_downloads,omitempty"`
	IsTemplate        *bool   `json:"is_template,omitempty"`
	LicenseTemplate   *string `json:"license_template,omitempty"`
	GitignoreTemplate *string `json:"gitignore_template,omitempty"`

	// Creating an organization repository. Required for non-owners.
	TeamID *int64 `json:"team_id,omitempty"`

	// API URLs
	URL              *string `json:"url,omitempty"`
	ArchiveURL       *string `json:"archive_url,omitempty"`
	AssigneesURL     *string `json:"assignees_url,omitempty"`
	BlobsURL         *string `json:"blobs_url,omitempty"`
	BranchesURL      *string `json:"branches_url,omitempty"`
	CollaboratorsURL *string `json:"collaborators_url,omitempty"`
	CommentsURL      *string `json:"comments_url,omitempty"`
	CommitsURL       *string `json:"commits_url,omitempty"`
	CompareURL       *string `json:"compare_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	ContributorsURL  *string `json:"contributors_url,omitempty"`
	DeploymentsURL   *string `json:"deployments_url,omitempty"`
	DownloadsURL     *string `json:"downloads_url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	ForksURL         *string `json:"forks_url,omitempty"`
	GitCommitsURL    *string `json:"git_commits_url,omitempty"`
	GitRefsURL       *string `json:"git_refs_url,omitempty"`
	GitTagsURL       *string `json:"git_tags_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssueCommentURL  *string `json:"issue_comment_url,omitempty"`
	IssueEventsURL   *string `json:"issue_events_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	KeysURL          *string `json:"keys_url,omitempty"`
	LabelsURL        *string `json:"labels_url,omitempty"`
	LanguagesURL     *string `json:"languages_url,omitempty"`
	MergesURL        *string `json:"merges_url,omitempty"`
	MilestonesURL    *string `json:"milestones_url,omitempty"`
	NotificationsURL *string `json:"notifications_url,omitempty"`
	PullsURL         *string `json:"pulls_url,omitempty"`
	ReleasesURL      *string `json:"releases_url,omitempty"`
	StargazersURL    *string `json:"stargazers_url,omitempty"`
	StatusesURL      *string `json:"statuses_url,omitempty"`
	SubscribersURL   *string `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string `json:"subscription_url,omitempty"`
	TagsURL          *string `json:"tags_url,omitempty"`
	TreesURL         *string `json:"trees_url,omitempty"`
	TeamsURL         *string `json:"teams_url,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://docs.github.com/en/free-pro-team@latest/rest/reference/search/#text-match-metadata
	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	// Visibility is only used for Create and Edit endpoints. The visibility field
	// overrides the field parameter when both are used.
	// Can be one of public, private or internal.
	Visibility *string `json:"visibility,omitempty"`
}

type Response struct {
	*http.Response

	// These fields provide the page values for paginating through a set of
	// results. Any or all of these may be set to the zero value for
	// responses that are not part of a paginated set, or for which there
	// are no additional pages.
	//
	// These fields support what is called "offset pagination" and should
	// be used with the ListOptions struct.
	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int

	// Additionally, some APIs support "cursor pagination" instead of offset.
	// This means that a token points directly to the next record which
	// can lead to O(1) performance compared to O(n) performance provided
	// by offset pagination.
	//
	// For APIs that support cursor pagination (such as
	// TeamsService.ListIDPGroupsInOrganization), the following field
	// will be populated to point to the next page.
	//
	// To use this token, set ListCursorOptions.Page to this value before
	// calling the endpoint again.
	NextPageToken string

	// Explicitly specify the Rate type so Rate's String() receiver doesn't
	// propagate to Response.
	Rate Rate
}

type Organization struct {
	Login                       *string    `json:"login,omitempty"`
	ID                          *int64     `json:"id,omitempty"`
	NodeID                      *string    `json:"node_id,omitempty"`
	AvatarURL                   *string    `json:"avatar_url,omitempty"`
	HTMLURL                     *string    `json:"html_url,omitempty"`
	Name                        *string    `json:"name,omitempty"`
	Company                     *string    `json:"company,omitempty"`
	Blog                        *string    `json:"blog,omitempty"`
	Location                    *string    `json:"location,omitempty"`
	Email                       *string    `json:"email,omitempty"`
	TwitterUsername             *string    `json:"twitter_username,omitempty"`
	Description                 *string    `json:"description,omitempty"`
	PublicRepos                 *int       `json:"public_repos,omitempty"`
	PublicGists                 *int       `json:"public_gists,omitempty"`
	Followers                   *int       `json:"followers,omitempty"`
	Following                   *int       `json:"following,omitempty"`
	CreatedAt                   *time.Time `json:"created_at,omitempty"`
	UpdatedAt                   *time.Time `json:"updated_at,omitempty"`
	TotalPrivateRepos           *int       `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos           *int       `json:"owned_private_repos,omitempty"`
	PrivateGists                *int       `json:"private_gists,omitempty"`
	DiskUsage                   *int       `json:"disk_usage,omitempty"`
	Collaborators               *int       `json:"collaborators,omitempty"`
	BillingEmail                *string    `json:"billing_email,omitempty"`
	Type                        *string    `json:"type,omitempty"`
	Plan                        *Plan      `json:"plan,omitempty"`
	TwoFactorRequirementEnabled *bool      `json:"two_factor_requirement_enabled,omitempty"`
	IsVerified                  *bool      `json:"is_verified,omitempty"`
	HasOrganizationProjects     *bool      `json:"has_organization_projects,omitempty"`
	HasRepositoryProjects       *bool      `json:"has_repository_projects,omitempty"`

	// DefaultRepoPermission can be one of: "read", "write", "admin", or "none". (Default: "read").
	// It is only used in OrganizationsService.Edit.
	DefaultRepoPermission *string `json:"default_repository_permission,omitempty"`
	// DefaultRepoSettings can be one of: "read", "write", "admin", or "none". (Default: "read").
	// It is only used in OrganizationsService.Get.
	DefaultRepoSettings *string `json:"default_repository_settings,omitempty"`

	// MembersCanCreateRepos default value is true and is only used in Organizations.Edit.
	MembersCanCreateRepos *bool `json:"members_can_create_repositories,omitempty"`

	// https://developer.github.com/changes/2019-12-03-internal-visibility-changes/#rest-v3-api
	MembersCanCreatePublicRepos   *bool `json:"members_can_create_public_repositories,omitempty"`
	MembersCanCreatePrivateRepos  *bool `json:"members_can_create_private_repositories,omitempty"`
	MembersCanCreateInternalRepos *bool `json:"members_can_create_internal_repositories,omitempty"`

	// MembersAllowedRepositoryCreationType denotes if organization members can create repositories
	// and the type of repositories they can create. Possible values are: "all", "private", or "none".
	//
	// Deprecated: Use MembersCanCreatePublicRepos, MembersCanCreatePrivateRepos, MembersCanCreateInternalRepos
	// instead. The new fields overrides the existing MembersAllowedRepositoryCreationType during 'edit'
	// operation and does not consider 'internal' repositories during 'get' operation
	MembersAllowedRepositoryCreationType *string `json:"members_allowed_repository_creation_type,omitempty"`

	// MembersCanCreatePages toggles whether organization members can create GitHub Pages sites.
	MembersCanCreatePages *bool `json:"members_can_create_pages,omitempty"`
	// MembersCanCreatePublicPages toggles whether organization members can create public GitHub Pages sites.
	MembersCanCreatePublicPages *bool `json:"members_can_create_public_pages,omitempty"`
	// MembersCanCreatePrivatePages toggles whether organization members can create private GitHub Pages sites.
	MembersCanCreatePrivatePages *bool `json:"members_can_create_private_pages,omitempty"`

	// API URLs
	URL              *string `json:"url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	MembersURL       *string `json:"members_url,omitempty"`
	PublicMembersURL *string `json:"public_members_url,omitempty"`
	ReposURL         *string `json:"repos_url,omitempty"`
}
type Match struct {
	Text    *string `json:"text,omitempty"`
	Indices []int   `json:"indices,omitempty"`
}

type TextMatch struct {
	ObjectURL  *string  `json:"object_url,omitempty"`
	ObjectType *string  `json:"object_type,omitempty"`
	Property   *string  `json:"property,omitempty"`
	Fragment   *string  `json:"fragment,omitempty"`
	Matches    []*Match `json:"matches,omitempty"`
}

type Timestamp struct {
	time.Time
}
type Plan struct {
	Name          *string `json:"name,omitempty"`
	Space         *int    `json:"space,omitempty"`
	Collaborators *int    `json:"collaborators,omitempty"`
	PrivateRepos  *int    `json:"private_repos,omitempty"`
	FilledSeats   *int    `json:"filled_seats,omitempty"`
	Seats         *int    `json:"seats,omitempty"`
}

type PRLink struct {
	HRef *string `json:"href,omitempty"`
}
type Label struct {
	ID          *int64  `json:"id,omitempty"`
	URL         *string `json:"url,omitempty"`
	Name        *string `json:"name,omitempty"`
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Default     *bool   `json:"default,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
}
type PRLinks struct {
	Self           *PRLink `json:"self,omitempty"`
	HTML           *PRLink `json:"html,omitempty"`
	Issue          *PRLink `json:"issue,omitempty"`
	Comments       *PRLink `json:"comments,omitempty"`
	ReviewComments *PRLink `json:"review_comments,omitempty"`
	ReviewComment  *PRLink `json:"review_comment,omitempty"`
	Commits        *PRLink `json:"commits,omitempty"`
	Statuses       *PRLink `json:"statuses,omitempty"`
}
type User struct {
	Login                   *string    `json:"login,omitempty"`
	ID                      *int64     `json:"id,omitempty"`
	NodeID                  *string    `json:"node_id,omitempty"`
	AvatarURL               *string    `json:"avatar_url,omitempty"`
	HTMLURL                 *string    `json:"html_url,omitempty"`
	GravatarID              *string    `json:"gravatar_id,omitempty"`
	Name                    *string    `json:"name,omitempty"`
	Company                 *string    `json:"company,omitempty"`
	Blog                    *string    `json:"blog,omitempty"`
	Location                *string    `json:"location,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	Hireable                *bool      `json:"hireable,omitempty"`
	Bio                     *string    `json:"bio,omitempty"`
	TwitterUsername         *string    `json:"twitter_username,omitempty"`
	PublicRepos             *int       `json:"public_repos,omitempty"`
	PublicGists             *int       `json:"public_gists,omitempty"`
	Followers               *int       `json:"followers,omitempty"`
	Following               *int       `json:"following,omitempty"`
	CreatedAt               *Timestamp `json:"created_at,omitempty"`
	UpdatedAt               *Timestamp `json:"updated_at,omitempty"`
	SuspendedAt             *Timestamp `json:"suspended_at,omitempty"`
	Type                    *string    `json:"type,omitempty"`
	SiteAdmin               *bool      `json:"site_admin,omitempty"`
	TotalPrivateRepos       *int       `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       *int       `json:"owned_private_repos,omitempty"`
	PrivateGists            *int       `json:"private_gists,omitempty"`
	DiskUsage               *int       `json:"disk_usage,omitempty"`
	Collaborators           *int       `json:"collaborators,omitempty"`
	TwoFactorAuthentication *bool      `json:"two_factor_authentication,omitempty"`
	Plan                    *Plan      `json:"plan,omitempty"`
	LdapDn                  *string    `json:"ldap_dn,omitempty"`

	// API URLs
	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://docs.github.com/en/free-pro-team@latest/rest/reference/search/#text-match-metadata
	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	// Permissions identifies the permissions that a user has on a given
	// repository. This is only populated when calling Repositories.ListCollaborators.
	Permissions map[string]bool `json:"permissions,omitempty"`
}
type Team struct {
	ID          *int64  `json:"id,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
	Slug        *string `json:"slug,omitempty"`

	// Permission specifies the default permission for repositories owned by the team.
	Permission *string `json:"permission,omitempty"`

	// Permissions identifies the permissions that a team has on a given
	// repository. This is only populated when calling Repositories.ListTeams.
	Permissions map[string]bool `json:"permissions,omitempty"`

	// Privacy identifies the level of privacy this team should have.
	// Possible values are:
	//     secret - only visible to organization owners and members of this team
	//     closed - visible to all members of this organization
	// Default is "secret".
	Privacy *string `json:"privacy,omitempty"`

	MembersCount    *int          `json:"members_count,omitempty"`
	ReposCount      *int          `json:"repos_count,omitempty"`
	Organization    *Organization `json:"organization,omitempty"`
	MembersURL      *string       `json:"members_url,omitempty"`
	RepositoriesURL *string       `json:"repositories_url,omitempty"`
	Parent          *Team         `json:"parent,omitempty"`

	// LDAPDN is only available in GitHub Enterprise and when the team
	// membership is synchronized with LDAP.
	LDAPDN *string `json:"ldap_dn,omitempty"`
}

type Milestone struct {
	URL          *string    `json:"url,omitempty"`
	HTMLURL      *string    `json:"html_url,omitempty"`
	LabelsURL    *string    `json:"labels_url,omitempty"`
	ID           *int64     `json:"id,omitempty"`
	Number       *int       `json:"number,omitempty"`
	State        *string    `json:"state,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Creator      *User      `json:"creator,omitempty"`
	OpenIssues   *int       `json:"open_issues,omitempty"`
	ClosedIssues *int       `json:"closed_issues,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
	DueOn        *time.Time `json:"due_on,omitempty"`
	NodeID       *string    `json:"node_id,omitempty"`
}
type PullRequestBranch struct {
	Label *string     `json:"label,omitempty"`
	Ref   *string     `json:"ref,omitempty"`
	SHA   *string     `json:"sha,omitempty"`
	Repo  *Repository `json:"repo,omitempty"`
	User  *User       `json:"user,omitempty"`
}

// PullRequest represents a GitHub pull request on a repository.
type PullRequest struct {
	ID                  *int64                `json:"id,omitempty"`
	Number              *int                  `json:"number,omitempty"`
	State               *string               `json:"state,omitempty"`
	Locked              *bool                 `json:"locked,omitempty"`
	Title               *string               `json:"title,omitempty"`
	Body                *string               `json:"body,omitempty"`
	CreatedAt           *time.Time            `json:"created_at,omitempty"`
	UpdatedAt           *time.Time            `json:"updated_at,omitempty"`
	ClosedAt            *time.Time            `json:"closed_at,omitempty"`
	MergedAt            *time.Time            `json:"merged_at,omitempty"`
	Labels              []*Label              `json:"labels,omitempty"`
	User                *User                 `json:"user,omitempty"`
	Draft               *bool                 `json:"draft,omitempty"`
	Merged              *bool                 `json:"merged,omitempty"`
	Mergeable           *bool                 `json:"mergeable,omitempty"`
	MergeableState      *string               `json:"mergeable_state,omitempty"`
	MergedBy            *User                 `json:"merged_by,omitempty"`
	MergeCommitSHA      *string               `json:"merge_commit_sha,omitempty"`
	Rebaseable          *bool                 `json:"rebaseable,omitempty"`
	Comments            *int                  `json:"comments,omitempty"`
	Commits             *int                  `json:"commits,omitempty"`
	Additions           *int                  `json:"additions,omitempty"`
	Deletions           *int                  `json:"deletions,omitempty"`
	ChangedFiles        *int                  `json:"changed_files,omitempty"`
	URL                 *string               `json:"url,omitempty"`
	HTMLURL             *string               `json:"html_url,omitempty"`
	IssueURL            *string               `json:"issue_url,omitempty"`
	StatusesURL         *string               `json:"statuses_url,omitempty"`
	DiffURL             *string               `json:"diff_url,omitempty"`
	PatchURL            *string               `json:"patch_url,omitempty"`
	CommitsURL          *string               `json:"commits_url,omitempty"`
	CommentsURL         *string               `json:"comments_url,omitempty"`
	ReviewCommentsURL   *string               `json:"review_comments_url,omitempty"`
	ReviewCommentURL    *string               `json:"review_comment_url,omitempty"`
	ReviewComments      *int                  `json:"review_comments,omitempty"`
	Assignee            *User                 `json:"assignee,omitempty"`
	Assignees           []*User               `json:"assignees,omitempty"`
	Milestone           *Milestone            `json:"milestone,omitempty"`
	MaintainerCanModify *bool                 `json:"maintainer_can_modify,omitempty"`
	AuthorAssociation   *string               `json:"author_association,omitempty"`
	NodeID              *string               `json:"node_id,omitempty"`
	RequestedReviewers  []*User               `json:"requested_reviewers,omitempty"`
	AutoMerge           *PullRequestAutoMerge `json:"auto_merge,omitempty"`

	// RequestedTeams is populated as part of the PullRequestEvent.
	// See, https://docs.github.com/en/developers/webhooks-and-events/github-event-types#pullrequestevent for an example.
	RequestedTeams []*Team `json:"requested_teams,omitempty"`

	Links *PRLinks           `json:"_links,omitempty"`
	Head  *PullRequestBranch `json:"head,omitempty"`
	Base  *PullRequestBranch `json:"base,omitempty"`

	// ActiveLockReason is populated only when LockReason is provided while locking the pull request.
	// Possible values are: "off-topic", "too heated", "resolved", and "spam".
	ActiveLockReason *string `json:"active_lock_reason,omitempty"`
}
type PullRequestAutoMerge struct {
	EnabledBy     *User   `json:"enabled_by,omitempty"`
	MergeMethod   *string `json:"merge_method,omitempty"`
	CommitTitle   *string `json:"commit_title,omitempty"`
	CommitMessage *string `json:"commit_message,omitempty"`
}
