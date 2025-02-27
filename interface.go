// Copyright (C) 2024 gitlab contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public
// License along with this program. If not, see
// <https://www.gnu.org/licenses/>.
//
// SPDX-License-Identifier: LGPL-3.0

package gitlab

// Client is a client for interacting with the GitLab API. To create a
// new client, use the [NewClient] function. For testing, use the
// [NewMockClient] function instead.
type Client interface {
	AccessRequests() AccessRequestsService
	Appearance() AppearanceService
	Applications() ApplicationsService
	AuditEvents() AuditEventsService
	Avatar() AvatarRequestsService
	AwardEmoji() AwardEmojiService
	Branches() BranchesService
	BroadcastMessage() BroadcastMessagesService
	BulkImports() BulkImportsService
	CIYMLTemplate() CIYMLTemplatesService
	ClusterAgents() ClusterAgentsService
	Commits() CommitsService
	ContainerRegistry() ContainerRegistryService
	CustomAttribute() CustomAttributesService
	DORAMetrics() DORAMetricsService
	DependencyListExport() DependencyListExportService
	DeployKeys() DeployKeysService
	DeployTokens() DeployTokensService
	DeploymentMergeRequests() DeploymentMergeRequestsService
	Deployments() DeploymentsService
	Discussions() DiscussionsService
	DockerfileTemplate() DockerfileTemplatesService
	DraftNotes() DraftNotesService
	Environments() EnvironmentsService
	EpicIssues() EpicIssuesService
	Epics() EpicsService
	ErrorTracking() ErrorTrackingService
	Events() EventsService
	ExternalStatusChecks() ExternalStatusChecksService
	Features() FeaturesService
	FreezePeriods() FreezePeriodsService
	GenericPackages() GenericPackagesService
	GeoNodes() GeoNodesService
	GitIgnoreTemplates() GitIgnoreTemplatesService
	GroupAccessTokens() GroupAccessTokensService
	GroupBadges() GroupBadgesService
	GroupCluster() GroupClustersService
	GroupEpicBoards() GroupEpicBoardsService
	GroupImportExport() GroupImportExportService
	GroupIssueBoards() GroupIssueBoardsService
	GroupIterations() GroupIterationsService
	GroupLabels() GroupLabelsService
	GroupMembers() GroupMembersService
	GroupMilestones() GroupMilestonesService
	GroupProtectedEnvironments() GroupProtectedEnvironmentsService
	GroupRepositoryStorageMove() GroupRepositoryStorageMoveService
	GroupSSHCertificates() GroupSSHCertificatesService
	GroupSecuritySettings() GroupSecuritySettingsService
	GroupVariables() GroupVariablesService
	GroupWikis() GroupWikisService
	Groups() GroupsService
	Import() ImportService
	InstanceCluster() InstanceClustersService
	InstanceVariables() InstanceVariablesService
	Invites() InvitesService
	Boards() IssueBoardsService
	IssueLinks() IssueLinksService
	Issues() IssuesService
	IssuesStatistics() IssuesStatisticsService
	JobTokenScope() JobTokenScopeService
	Jobs() JobsService
	Keys() KeysService
	Labels() LabelsService
	License() LicenseService
	LicenseTemplates() LicenseTemplatesService
	ManagedLicenses() ManagedLicensesService
	Markdown() MarkdownService
	MemberRolesService() MemberRolesService
	MergeRequestApprovals() MergeRequestApprovalsService
	MergeRequests() MergeRequestsService
	MergeTrains() MergeTrainsService
	Metadata() MetadataService
	Milestones() MilestonesService
	Namespaces() NamespacesService
	Notes() NotesService
	NotificationSettings() NotificationSettingsService
	Packages() PackagesService
	PagesDomains() PagesDomainsService
	Pages() PagesService
	PersonalAccessTokens() PersonalAccessTokensService
	PipelineSchedules() PipelineSchedulesService
	PipelineTriggers() PipelineTriggersService
	Pipelines() PipelinesService
	PlanLimits() PlanLimitsService
	ProjectAccessTokens() ProjectAccessTokensService
	ProjectBadges() ProjectBadgesService
	ProjectCluster() ProjectClustersService
	ProjectFeatureFlags() ProjectFeatureFlagService
	ProjectImportExport() ProjectImportExportService
	ProjectIterations() ProjectIterationsService
	ProjectMarkdownUploads() ProjectMarkdownUploadsService
	ProjectMembers() ProjectMembersService
	ProjectMirrors() ProjectMirrorService
	ProjectRepositoryStorageMove() ProjectRepositoryStorageMoveService
	ProjectSnippets() ProjectSnippetsService
	ProjectTemplates() ProjectTemplatesService
	ProjectVariables() ProjectVariablesService
	ProjectVulnerabilities() ProjectVulnerabilitiesService
	Projects() ProjectsService
	ProtectedBranches() ProtectedBranchesService
	ProtectedEnvironments() ProtectedEnvironmentsService
	ProtectedTags() ProtectedTagsService
	ReleaseLinks() ReleaseLinksService
	Releases() ReleasesService
	Repositories() RepositoriesService
	RepositoryFiles() RepositoryFilesService
	RepositorySubmodules() RepositorySubmodulesService
	ResourceGroup() ResourceGroupService
	ResourceIterationEvents() ResourceIterationEventsService
	ResourceLabelEvents() ResourceLabelEventsService
	ResourceMilestoneEvents() ResourceMilestoneEventsService
	ResourceStateEvents() ResourceStateEventsService
	ResourceWeightEvents() ResourceWeightEventsService
	Runners() RunnersService
	Search() SearchService
	Services() ServicesService
	Settings() SettingsService
	Sidekiq() SidekiqService
	SnippetRepositoryStorageMove() SnippetRepositoryStorageMoveService
	Snippets() SnippetsService
	SystemHooks() SystemHooksService
	Tags() TagsService
	Todos() TodosService
	Topics() TopicsService
	Users() UsersService
	Validate() ValidateService
	Version() VersionService
	Wikis() WikisService
}
