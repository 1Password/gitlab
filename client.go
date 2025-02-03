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

import (
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// client is a wrapper around [gitlab.Client] that implements the
// [Client] interface by turning struct fields into accessors.
type client struct {
	*gitlab.Client
	accessRequestsService               AccessRequestsService
	appearanceService                   AppearanceService
	applicationsService                 ApplicationsService
	auditEventsService                  AuditEventsService
	avatarRequestsService               AvatarRequestsService
	awardEmojiService                   AwardEmojiService
	branchesService                     BranchesService
	broadcastMessagesService            BroadcastMessagesService
	cIYMLTemplatesService               CIYMLTemplatesService
	clusterAgentsService                ClusterAgentsService
	commitsService                      CommitsService
	containerRegistryService            ContainerRegistryService
	customAttributesService             CustomAttributesService
	dORAMetricsService                  DORAMetricsService
	dependencyListExportService         DependencyListExportService
	deployKeysService                   DeployKeysService
	deployTokensService                 DeployTokensService
	deploymentMergeRequestsService      DeploymentMergeRequestsService
	deploymentsService                  DeploymentsService
	discussionsService                  DiscussionsService
	dockerfileTemplatesService          DockerfileTemplatesService
	draftNotesService                   DraftNotesService
	environmentsService                 EnvironmentsService
	epicIssuesService                   EpicIssuesService
	epicsService                        EpicsService
	errorTrackingService                ErrorTrackingService
	eventsService                       EventsService
	externalStatusChecksService         ExternalStatusChecksService
	featuresService                     FeaturesService
	freezePeriodsService                FreezePeriodsService
	genericPackagesService              GenericPackagesService
	geoNodesService                     GeoNodesService
	gitIgnoreTemplatesService           GitIgnoreTemplatesService
	groupAccessTokensService            GroupAccessTokensService
	groupBadgesService                  GroupBadgesService
	groupClustersService                GroupClustersService
	groupEpicBoardsService              GroupEpicBoardsService
	groupImportExportService            GroupImportExportService
	groupIssueBoardsService             GroupIssueBoardsService
	groupIterationsService              GroupIterationsService
	groupLabelsService                  GroupLabelsService
	groupMembersService                 GroupMembersService
	groupMilestonesService              GroupMilestonesService
	groupProtectedEnvironmentsService   GroupProtectedEnvironmentsService
	groupRepositoryStorageMoveService   GroupRepositoryStorageMoveService
	groupSSHCertificatesService         GroupSSHCertificatesService
	groupSecuritySettingsService        GroupSecuritySettingsService
	groupVariablesService               GroupVariablesService
	groupWikisService                   GroupWikisService
	groupsService                       GroupsService
	importService                       ImportService
	instanceClustersService             InstanceClustersService
	instanceVariablesService            InstanceVariablesService
	invitesService                      InvitesService
	issueBoardsService                  IssueBoardsService
	issueLinksService                   IssueLinksService
	issuesService                       IssuesService
	issuesStatisticsService             IssuesStatisticsService
	jobTokenScopeService                JobTokenScopeService
	jobsService                         JobsService
	keysService                         KeysService
	labelsService                       LabelsService
	licenseService                      LicenseService
	licenseTemplatesService             LicenseTemplatesService
	managedLicensesService              ManagedLicensesService
	markdownService                     MarkdownService
	memberRolesService                  MemberRolesService
	mergeRequestApprovalsService        MergeRequestApprovalsService
	mergeRequestsService                MergeRequestsService
	mergeTrainsService                  MergeTrainsService
	metadataService                     MetadataService
	milestonesService                   MilestonesService
	namespacesService                   NamespacesService
	notesService                        NotesService
	notificationSettingsService         NotificationSettingsService
	packagesService                     PackagesService
	pagesDomainsService                 PagesDomainsService
	pagesService                        PagesService
	personalAccessTokensService         PersonalAccessTokensService
	pipelineSchedulesService            PipelineSchedulesService
	pipelineTriggersService             PipelineTriggersService
	pipelinesService                    PipelinesService
	planLimitsService                   PlanLimitsService
	projectAccessTokensService          ProjectAccessTokensService
	projectBadgesService                ProjectBadgesService
	projectClustersService              ProjectClustersService
	projectFeatureFlagService           ProjectFeatureFlagService
	projectImportExportService          ProjectImportExportService
	projectIterationsService            ProjectIterationsService
	projectMarkdownUploadsService       ProjectMarkdownUploadsService
	projectMembersService               ProjectMembersService
	projectMirrorService                ProjectMirrorService
	projectRepositoryStorageMoveService ProjectRepositoryStorageMoveService
	projectSnippetsService              ProjectSnippetsService
	projectTemplatesService             ProjectTemplatesService
	projectVariablesService             ProjectVariablesService
	projectVulnerabilitiesService       ProjectVulnerabilitiesService
	projectsService                     ProjectsService
	protectedBranchesService            ProtectedBranchesService
	protectedEnvironmentsService        ProtectedEnvironmentsService
	protectedTagsService                ProtectedTagsService
	releaseLinksService                 ReleaseLinksService
	releasesService                     ReleasesService
	repositoriesService                 RepositoriesService
	repositoryFilesService              RepositoryFilesService
	repositorySubmodulesService         RepositorySubmodulesService
	resourceGroupService                ResourceGroupService
	resourceIterationEventsService      ResourceIterationEventsService
	resourceLabelEventsService          ResourceLabelEventsService
	resourceMilestoneEventsService      ResourceMilestoneEventsService
	resourceStateEventsService          ResourceStateEventsService
	resourceWeightEventsService         ResourceWeightEventsService
	runnersService                      RunnersService
	searchService                       SearchService
	servicesService                     ServicesService
	settingsService                     SettingsService
	sidekiqService                      SidekiqService
	snippetRepositoryStorageMoveService SnippetRepositoryStorageMoveService
	snippetsService                     SnippetsService
	systemHooksService                  SystemHooksService
	tagsService                         TagsService
	todosService                        TodosService
	topicsService                       TopicsService
	usersService                        UsersService
	validateService                     ValidateService
	versionService                      VersionService
	wikisService                        WikisService
}

// AccessRequests returns the [AccessRequestsService] service for the client.
func (c *client) AccessRequests() AccessRequestsService {
	return c.accessRequestsService
}

// Appearance returns the [AppearanceService] service for the client.
func (c *client) Appearance() AppearanceService {
	return c.appearanceService
}

// Applications returns the [ApplicationsService] service for the client.
func (c *client) Applications() ApplicationsService {
	return c.applicationsService
}

// AuditEvents returns the [AuditEventsService] service for the client.
func (c *client) AuditEvents() AuditEventsService {
	return c.auditEventsService
}

// Avatar returns the [AvatarRequestsService] service for the client.
func (c *client) Avatar() AvatarRequestsService {
	return c.avatarRequestsService
}

// AwardEmoji returns the [AwardEmojiService] service for the client.
func (c *client) AwardEmoji() AwardEmojiService {
	return c.awardEmojiService
}

// Branches returns the [BranchesService] service for the client.
func (c *client) Branches() BranchesService {
	return c.branchesService
}

// BroadcastMessage returns the [BroadcastMessagesService] service for the client.
func (c *client) BroadcastMessage() BroadcastMessagesService {
	return c.broadcastMessagesService
}

// CIYMLTemplate returns the [CIYMLTemplatesService] service for the client.
func (c *client) CIYMLTemplate() CIYMLTemplatesService {
	return c.cIYMLTemplatesService
}

// ClusterAgents returns the [ClusterAgentsService] service for the client.
func (c *client) ClusterAgents() ClusterAgentsService {
	return c.clusterAgentsService
}

// Commits returns the [CommitsService] service for the client.
func (c *client) Commits() CommitsService {
	return c.commitsService
}

// ContainerRegistry returns the [ContainerRegistryService] service for the client.
func (c *client) ContainerRegistry() ContainerRegistryService {
	return c.containerRegistryService
}

// CustomAttribute returns the [CustomAttributesService] service for the client.
func (c *client) CustomAttribute() CustomAttributesService {
	return c.customAttributesService
}

// DORAMetrics returns the [DORAMetricsService] service for the client.
func (c *client) DORAMetrics() DORAMetricsService {
	return c.dORAMetricsService
}

// DependencyListExport returns the [DependencyListExportService] service for the client.
func (c *client) DependencyListExport() DependencyListExportService {
	return c.dependencyListExportService
}

// DeployKeys returns the [DeployKeysService] service for the client.
func (c *client) DeployKeys() DeployKeysService {
	return c.deployKeysService
}

// DeployTokens returns the [DeployTokensService] service for the client.
func (c *client) DeployTokens() DeployTokensService {
	return c.deployTokensService
}

// DeploymentMergeRequests returns the [DeploymentMergeRequestsService] service for the client.
func (c *client) DeploymentMergeRequests() DeploymentMergeRequestsService {
	return c.deploymentMergeRequestsService
}

// Deployments returns the [DeploymentsService] service for the client.
func (c *client) Deployments() DeploymentsService {
	return c.deploymentsService
}

// Discussions returns the [DiscussionsService] service for the client.
func (c *client) Discussions() DiscussionsService {
	return c.discussionsService
}

// DockerfileTemplate returns the [DockerfileTemplatesService] service for the client.
func (c *client) DockerfileTemplate() DockerfileTemplatesService {
	return c.dockerfileTemplatesService
}

// DraftNotes returns the [DraftNotesService] service for the client.
func (c *client) DraftNotes() DraftNotesService {
	return c.draftNotesService
}

// Environments returns the [EnvironmentsService] service for the client.
func (c *client) Environments() EnvironmentsService {
	return c.environmentsService
}

// EpicIssues returns the [EpicIssuesService] service for the client.
func (c *client) EpicIssues() EpicIssuesService {
	return c.epicIssuesService
}

// Epics returns the [EpicsService] service for the client.
func (c *client) Epics() EpicsService {
	return c.epicsService
}

// ErrorTracking returns the [ErrorTrackingService] service for the client.
func (c *client) ErrorTracking() ErrorTrackingService {
	return c.errorTrackingService
}

// Events returns the [EventsService] service for the client.
func (c *client) Events() EventsService {
	return c.eventsService
}

// ExternalStatusChecks returns the [ExternalStatusChecksService] service for the client.
func (c *client) ExternalStatusChecks() ExternalStatusChecksService {
	return c.externalStatusChecksService
}

// Features returns the [FeaturesService] service for the client.
func (c *client) Features() FeaturesService {
	return c.featuresService
}

// FreezePeriods returns the [FreezePeriodsService] service for the client.
func (c *client) FreezePeriods() FreezePeriodsService {
	return c.freezePeriodsService
}

// GenericPackages returns the [GenericPackagesService] service for the client.
func (c *client) GenericPackages() GenericPackagesService {
	return c.genericPackagesService
}

// GeoNodes returns the [GeoNodesService] service for the client.
func (c *client) GeoNodes() GeoNodesService {
	return c.geoNodesService
}

// GitIgnoreTemplates returns the [GitIgnoreTemplatesService] service for the client.
func (c *client) GitIgnoreTemplates() GitIgnoreTemplatesService {
	return c.gitIgnoreTemplatesService
}

// GroupAccessTokens returns the [GroupAccessTokensService] service for the client.
func (c *client) GroupAccessTokens() GroupAccessTokensService {
	return c.groupAccessTokensService
}

// GroupBadges returns the [GroupBadgesService] service for the client.
func (c *client) GroupBadges() GroupBadgesService {
	return c.groupBadgesService
}

// GroupCluster returns the [GroupClustersService] service for the client.
func (c *client) GroupCluster() GroupClustersService {
	return c.groupClustersService
}

// GroupEpicBoards returns the [GroupEpicBoardsService] service for the client.
func (c *client) GroupEpicBoards() GroupEpicBoardsService {
	return c.groupEpicBoardsService
}

// GroupImportExport returns the [GroupImportExportService] service for the client.
func (c *client) GroupImportExport() GroupImportExportService {
	return c.groupImportExportService
}

// GroupIssueBoards returns the [GroupIssueBoardsService] service for the client.
func (c *client) GroupIssueBoards() GroupIssueBoardsService {
	return c.groupIssueBoardsService
}

// GroupIterations returns the [GroupIterationsService] service for the client.
func (c *client) GroupIterations() GroupIterationsService {
	return c.groupIterationsService
}

// GroupLabels returns the [GroupLabelsService] service for the client.
func (c *client) GroupLabels() GroupLabelsService {
	return c.groupLabelsService
}

// GroupMembers returns the [GroupMembersService] service for the client.
func (c *client) GroupMembers() GroupMembersService {
	return c.groupMembersService
}

// GroupMilestones returns the [GroupMilestonesService] service for the client.
func (c *client) GroupMilestones() GroupMilestonesService {
	return c.groupMilestonesService
}

// GroupProtectedEnvironments returns the [GroupProtectedEnvironmentsService] service for the client.
func (c *client) GroupProtectedEnvironments() GroupProtectedEnvironmentsService {
	return c.groupProtectedEnvironmentsService
}

// GroupRepositoryStorageMove returns the [GroupRepositoryStorageMoveService] service for the client.
func (c *client) GroupRepositoryStorageMove() GroupRepositoryStorageMoveService {
	return c.groupRepositoryStorageMoveService
}

// GroupSSHCertificates returns the [GroupSSHCertificatesService] service for the client.
func (c *client) GroupSSHCertificates() GroupSSHCertificatesService {
	return c.groupSSHCertificatesService
}

// GroupSecuritySettings returns the [GroupSecuritySettingsService] service for the client.
func (c *client) GroupSecuritySettings() GroupSecuritySettingsService {
	return c.groupSecuritySettingsService
}

// GroupVariables returns the [GroupVariablesService] service for the client.
func (c *client) GroupVariables() GroupVariablesService {
	return c.groupVariablesService
}

// GroupWikis returns the [GroupWikisService] service for the client.
func (c *client) GroupWikis() GroupWikisService {
	return c.groupWikisService
}

// Groups returns the [GroupsService] service for the client.
func (c *client) Groups() GroupsService {
	return c.groupsService
}

// Import returns the [ImportService] service for the client.
func (c *client) Import() ImportService {
	return c.importService
}

// InstanceCluster returns the [InstanceClustersService] service for the client.
func (c *client) InstanceCluster() InstanceClustersService {
	return c.instanceClustersService
}

// InstanceVariables returns the [InstanceVariablesService] service for the client.
func (c *client) InstanceVariables() InstanceVariablesService {
	return c.instanceVariablesService
}

// Invites returns the [InvitesService] service for the client.
func (c *client) Invites() InvitesService {
	return c.invitesService
}

// Boards returns the [IssueBoardsService] service for the client.
func (c *client) Boards() IssueBoardsService {
	return c.issueBoardsService
}

// IssueLinks returns the [IssueLinksService] service for the client.
func (c *client) IssueLinks() IssueLinksService {
	return c.issueLinksService
}

// Issues returns the [IssuesService] service for the client.
func (c *client) Issues() IssuesService {
	return c.issuesService
}

// IssuesStatistics returns the [IssuesStatisticsService] service for the client.
func (c *client) IssuesStatistics() IssuesStatisticsService {
	return c.issuesStatisticsService
}

// JobTokenScope returns the [JobTokenScopeService] service for the client.
func (c *client) JobTokenScope() JobTokenScopeService {
	return c.jobTokenScopeService
}

// Jobs returns the [JobsService] service for the client.
func (c *client) Jobs() JobsService {
	return c.jobsService
}

// Keys returns the [KeysService] service for the client.
func (c *client) Keys() KeysService {
	return c.keysService
}

// Labels returns the [LabelsService] service for the client.
func (c *client) Labels() LabelsService {
	return c.labelsService
}

// License returns the [LicenseService] service for the client.
func (c *client) License() LicenseService {
	return c.licenseService
}

// LicenseTemplates returns the [LicenseTemplatesService] service for the client.
func (c *client) LicenseTemplates() LicenseTemplatesService {
	return c.licenseTemplatesService
}

// ManagedLicenses returns the [ManagedLicensesService] service for the client.
func (c *client) ManagedLicenses() ManagedLicensesService {
	return c.managedLicensesService
}

// Markdown returns the [MarkdownService] service for the client.
func (c *client) Markdown() MarkdownService {
	return c.markdownService
}

// MemberRolesService returns the [MemberRolesService] service for the client.
func (c *client) MemberRolesService() MemberRolesService {
	return c.memberRolesService
}

// MergeRequestApprovals returns the [MergeRequestApprovalsService] service for the client.
func (c *client) MergeRequestApprovals() MergeRequestApprovalsService {
	return c.mergeRequestApprovalsService
}

// MergeRequests returns the [MergeRequestsService] service for the client.
func (c *client) MergeRequests() MergeRequestsService {
	return c.mergeRequestsService
}

// MergeTrains returns the [MergeTrainsService] service for the client.
func (c *client) MergeTrains() MergeTrainsService {
	return c.mergeTrainsService
}

// Metadata returns the [MetadataService] service for the client.
func (c *client) Metadata() MetadataService {
	return c.metadataService
}

// Milestones returns the [MilestonesService] service for the client.
func (c *client) Milestones() MilestonesService {
	return c.milestonesService
}

// Namespaces returns the [NamespacesService] service for the client.
func (c *client) Namespaces() NamespacesService {
	return c.namespacesService
}

// Notes returns the [NotesService] service for the client.
func (c *client) Notes() NotesService {
	return c.notesService
}

// NotificationSettings returns the [NotificationSettingsService] service for the client.
func (c *client) NotificationSettings() NotificationSettingsService {
	return c.notificationSettingsService
}

// Packages returns the [PackagesService] service for the client.
func (c *client) Packages() PackagesService {
	return c.packagesService
}

// PagesDomains returns the [PagesDomainsService] service for the client.
func (c *client) PagesDomains() PagesDomainsService {
	return c.pagesDomainsService
}

// Pages returns the [PagesService] service for the client.
func (c *client) Pages() PagesService {
	return c.pagesService
}

// PersonalAccessTokens returns the [PersonalAccessTokensService] service for the client.
func (c *client) PersonalAccessTokens() PersonalAccessTokensService {
	return c.personalAccessTokensService
}

// PipelineSchedules returns the [PipelineSchedulesService] service for the client.
func (c *client) PipelineSchedules() PipelineSchedulesService {
	return c.pipelineSchedulesService
}

// PipelineTriggers returns the [PipelineTriggersService] service for the client.
func (c *client) PipelineTriggers() PipelineTriggersService {
	return c.pipelineTriggersService
}

// Pipelines returns the [PipelinesService] service for the client.
func (c *client) Pipelines() PipelinesService {
	return c.pipelinesService
}

// PlanLimits returns the [PlanLimitsService] service for the client.
func (c *client) PlanLimits() PlanLimitsService {
	return c.planLimitsService
}

// ProjectAccessTokens returns the [ProjectAccessTokensService] service for the client.
func (c *client) ProjectAccessTokens() ProjectAccessTokensService {
	return c.projectAccessTokensService
}

// ProjectBadges returns the [ProjectBadgesService] service for the client.
func (c *client) ProjectBadges() ProjectBadgesService {
	return c.projectBadgesService
}

// ProjectCluster returns the [ProjectClustersService] service for the client.
func (c *client) ProjectCluster() ProjectClustersService {
	return c.projectClustersService
}

// ProjectFeatureFlags returns the [ProjectFeatureFlagService] service for the client.
func (c *client) ProjectFeatureFlags() ProjectFeatureFlagService {
	return c.projectFeatureFlagService
}

// ProjectImportExport returns the [ProjectImportExportService] service for the client.
func (c *client) ProjectImportExport() ProjectImportExportService {
	return c.projectImportExportService
}

// ProjectIterations returns the [ProjectIterationsService] service for the client.
func (c *client) ProjectIterations() ProjectIterationsService {
	return c.projectIterationsService
}

// ProjectMarkdownUploads returns the [ProjectMarkdownUploadsService] service for the client.
func (c *client) ProjectMarkdownUploads() ProjectMarkdownUploadsService {
	return c.projectMarkdownUploadsService
}

// ProjectMembers returns the [ProjectMembersService] service for the client.
func (c *client) ProjectMembers() ProjectMembersService {
	return c.projectMembersService
}

// ProjectMirrors returns the [ProjectMirrorService] service for the client.
func (c *client) ProjectMirrors() ProjectMirrorService {
	return c.projectMirrorService
}

// ProjectRepositoryStorageMove returns the [ProjectRepositoryStorageMoveService] service for the client.
func (c *client) ProjectRepositoryStorageMove() ProjectRepositoryStorageMoveService {
	return c.projectRepositoryStorageMoveService
}

// ProjectSnippets returns the [ProjectSnippetsService] service for the client.
func (c *client) ProjectSnippets() ProjectSnippetsService {
	return c.projectSnippetsService
}

// ProjectTemplates returns the [ProjectTemplatesService] service for the client.
func (c *client) ProjectTemplates() ProjectTemplatesService {
	return c.projectTemplatesService
}

// ProjectVariables returns the [ProjectVariablesService] service for the client.
func (c *client) ProjectVariables() ProjectVariablesService {
	return c.projectVariablesService
}

// ProjectVulnerabilities returns the [ProjectVulnerabilitiesService] service for the client.
func (c *client) ProjectVulnerabilities() ProjectVulnerabilitiesService {
	return c.projectVulnerabilitiesService
}

// Projects returns the [ProjectsService] service for the client.
func (c *client) Projects() ProjectsService {
	return c.projectsService
}

// ProtectedBranches returns the [ProtectedBranchesService] service for the client.
func (c *client) ProtectedBranches() ProtectedBranchesService {
	return c.protectedBranchesService
}

// ProtectedEnvironments returns the [ProtectedEnvironmentsService] service for the client.
func (c *client) ProtectedEnvironments() ProtectedEnvironmentsService {
	return c.protectedEnvironmentsService
}

// ProtectedTags returns the [ProtectedTagsService] service for the client.
func (c *client) ProtectedTags() ProtectedTagsService {
	return c.protectedTagsService
}

// ReleaseLinks returns the [ReleaseLinksService] service for the client.
func (c *client) ReleaseLinks() ReleaseLinksService {
	return c.releaseLinksService
}

// Releases returns the [ReleasesService] service for the client.
func (c *client) Releases() ReleasesService {
	return c.releasesService
}

// Repositories returns the [RepositoriesService] service for the client.
func (c *client) Repositories() RepositoriesService {
	return c.repositoriesService
}

// RepositoryFiles returns the [RepositoryFilesService] service for the client.
func (c *client) RepositoryFiles() RepositoryFilesService {
	return c.repositoryFilesService
}

// RepositorySubmodules returns the [RepositorySubmodulesService] service for the client.
func (c *client) RepositorySubmodules() RepositorySubmodulesService {
	return c.repositorySubmodulesService
}

// ResourceGroup returns the [ResourceGroupService] service for the client.
func (c *client) ResourceGroup() ResourceGroupService {
	return c.resourceGroupService
}

// ResourceIterationEvents returns the [ResourceIterationEventsService] service for the client.
func (c *client) ResourceIterationEvents() ResourceIterationEventsService {
	return c.resourceIterationEventsService
}

// ResourceLabelEvents returns the [ResourceLabelEventsService] service for the client.
func (c *client) ResourceLabelEvents() ResourceLabelEventsService {
	return c.resourceLabelEventsService
}

// ResourceMilestoneEvents returns the [ResourceMilestoneEventsService] service for the client.
func (c *client) ResourceMilestoneEvents() ResourceMilestoneEventsService {
	return c.resourceMilestoneEventsService
}

// ResourceStateEvents returns the [ResourceStateEventsService] service for the client.
func (c *client) ResourceStateEvents() ResourceStateEventsService {
	return c.resourceStateEventsService
}

// ResourceWeightEvents returns the [ResourceWeightEventsService] service for the client.
func (c *client) ResourceWeightEvents() ResourceWeightEventsService {
	return c.resourceWeightEventsService
}

// Runners returns the [RunnersService] service for the client.
func (c *client) Runners() RunnersService {
	return c.runnersService
}

// Search returns the [SearchService] service for the client.
func (c *client) Search() SearchService {
	return c.searchService
}

// Services returns the [ServicesService] service for the client.
func (c *client) Services() ServicesService {
	return c.servicesService
}

// Settings returns the [SettingsService] service for the client.
func (c *client) Settings() SettingsService {
	return c.settingsService
}

// Sidekiq returns the [SidekiqService] service for the client.
func (c *client) Sidekiq() SidekiqService {
	return c.sidekiqService
}

// SnippetRepositoryStorageMove returns the [SnippetRepositoryStorageMoveService] service for the client.
func (c *client) SnippetRepositoryStorageMove() SnippetRepositoryStorageMoveService {
	return c.snippetRepositoryStorageMoveService
}

// Snippets returns the [SnippetsService] service for the client.
func (c *client) Snippets() SnippetsService {
	return c.snippetsService
}

// SystemHooks returns the [SystemHooksService] service for the client.
func (c *client) SystemHooks() SystemHooksService {
	return c.systemHooksService
}

// Tags returns the [TagsService] service for the client.
func (c *client) Tags() TagsService {
	return c.tagsService
}

// Todos returns the [TodosService] service for the client.
func (c *client) Todos() TodosService {
	return c.todosService
}

// Topics returns the [TopicsService] service for the client.
func (c *client) Topics() TopicsService {
	return c.topicsService
}

// Users returns the [UsersService] service for the client.
func (c *client) Users() UsersService {
	return c.usersService
}

// Validate returns the [ValidateService] service for the client.
func (c *client) Validate() ValidateService {
	return c.validateService
}

// Version returns the [VersionService] service for the client.
func (c *client) Version() VersionService {
	return c.versionService
}

// Wikis returns the [WikisService] service for the client.
func (c *client) Wikis() WikisService {
	return c.wikisService
}

// NewClient creates a new [Client] with the provided token and options.
func NewClient(token string, opts ...ClientOptionFunc) (Client, error) {
	gl, err := gitlab.NewOAuthClient(token, opts...)
	if err != nil {
		return nil, err
	}

	return FromClient(gl), nil
}

// FromClient creates a new [Client] from an existing [gitlab.Client].
func FromClient(gl *gitlab.Client) Client {
	return &client{
		Client:                              gl,
		accessRequestsService:               gl.AccessRequests,
		appearanceService:                   gl.Appearance,
		applicationsService:                 gl.Applications,
		auditEventsService:                  gl.AuditEvents,
		avatarRequestsService:               gl.Avatar,
		awardEmojiService:                   gl.AwardEmoji,
		branchesService:                     gl.Branches,
		broadcastMessagesService:            gl.BroadcastMessage,
		cIYMLTemplatesService:               gl.CIYMLTemplate,
		clusterAgentsService:                gl.ClusterAgents,
		commitsService:                      gl.Commits,
		containerRegistryService:            gl.ContainerRegistry,
		customAttributesService:             gl.CustomAttribute,
		dORAMetricsService:                  gl.DORAMetrics,
		dependencyListExportService:         gl.DependencyListExport,
		deployKeysService:                   gl.DeployKeys,
		deployTokensService:                 gl.DeployTokens,
		deploymentMergeRequestsService:      gl.DeploymentMergeRequests,
		deploymentsService:                  gl.Deployments,
		discussionsService:                  gl.Discussions,
		dockerfileTemplatesService:          gl.DockerfileTemplate,
		draftNotesService:                   gl.DraftNotes,
		environmentsService:                 gl.Environments,
		epicIssuesService:                   gl.EpicIssues,
		epicsService:                        gl.Epics,
		errorTrackingService:                gl.ErrorTracking,
		eventsService:                       gl.Events,
		externalStatusChecksService:         gl.ExternalStatusChecks,
		featuresService:                     gl.Features,
		freezePeriodsService:                gl.FreezePeriods,
		genericPackagesService:              gl.GenericPackages,
		geoNodesService:                     gl.GeoNodes,
		gitIgnoreTemplatesService:           gl.GitIgnoreTemplates,
		groupAccessTokensService:            gl.GroupAccessTokens,
		groupBadgesService:                  gl.GroupBadges,
		groupClustersService:                gl.GroupCluster,
		groupEpicBoardsService:              gl.GroupEpicBoards,
		groupImportExportService:            gl.GroupImportExport,
		groupIssueBoardsService:             gl.GroupIssueBoards,
		groupIterationsService:              gl.GroupIterations,
		groupLabelsService:                  gl.GroupLabels,
		groupMembersService:                 gl.GroupMembers,
		groupMilestonesService:              gl.GroupMilestones,
		groupProtectedEnvironmentsService:   gl.GroupProtectedEnvironments,
		groupRepositoryStorageMoveService:   gl.GroupRepositoryStorageMove,
		groupSSHCertificatesService:         gl.GroupSSHCertificates,
		groupSecuritySettingsService:        gl.GroupSecuritySettings,
		groupVariablesService:               gl.GroupVariables,
		groupWikisService:                   gl.GroupWikis,
		groupsService:                       gl.Groups,
		importService:                       gl.Import,
		instanceClustersService:             gl.InstanceCluster,
		instanceVariablesService:            gl.InstanceVariables,
		invitesService:                      gl.Invites,
		issueBoardsService:                  gl.Boards,
		issueLinksService:                   gl.IssueLinks,
		issuesService:                       gl.Issues,
		issuesStatisticsService:             gl.IssuesStatistics,
		jobTokenScopeService:                gl.JobTokenScope,
		jobsService:                         gl.Jobs,
		keysService:                         gl.Keys,
		labelsService:                       gl.Labels,
		licenseService:                      gl.License,
		licenseTemplatesService:             gl.LicenseTemplates,
		managedLicensesService:              gl.ManagedLicenses,
		markdownService:                     gl.Markdown,
		memberRolesService:                  gl.MemberRolesService,
		mergeRequestApprovalsService:        gl.MergeRequestApprovals,
		mergeRequestsService:                gl.MergeRequests,
		mergeTrainsService:                  gl.MergeTrains,
		metadataService:                     gl.Metadata,
		milestonesService:                   gl.Milestones,
		namespacesService:                   gl.Namespaces,
		notesService:                        gl.Notes,
		notificationSettingsService:         gl.NotificationSettings,
		packagesService:                     gl.Packages,
		pagesDomainsService:                 gl.PagesDomains,
		pagesService:                        gl.Pages,
		personalAccessTokensService:         gl.PersonalAccessTokens,
		pipelineSchedulesService:            gl.PipelineSchedules,
		pipelineTriggersService:             gl.PipelineTriggers,
		pipelinesService:                    gl.Pipelines,
		planLimitsService:                   gl.PlanLimits,
		projectAccessTokensService:          gl.ProjectAccessTokens,
		projectBadgesService:                gl.ProjectBadges,
		projectClustersService:              gl.ProjectCluster,
		projectFeatureFlagService:           gl.ProjectFeatureFlags,
		projectImportExportService:          gl.ProjectImportExport,
		projectIterationsService:            gl.ProjectIterations,
		projectMarkdownUploadsService:       gl.ProjectMarkdownUploads,
		projectMembersService:               gl.ProjectMembers,
		projectMirrorService:                gl.ProjectMirrors,
		projectRepositoryStorageMoveService: gl.ProjectRepositoryStorageMove,
		projectSnippetsService:              gl.ProjectSnippets,
		projectTemplatesService:             gl.ProjectTemplates,
		projectVariablesService:             gl.ProjectVariables,
		projectVulnerabilitiesService:       gl.ProjectVulnerabilities,
		projectsService:                     gl.Projects,
		protectedBranchesService:            gl.ProtectedBranches,
		protectedEnvironmentsService:        gl.ProtectedEnvironments,
		protectedTagsService:                gl.ProtectedTags,
		releaseLinksService:                 gl.ReleaseLinks,
		releasesService:                     gl.Releases,
		repositoriesService:                 gl.Repositories,
		repositoryFilesService:              gl.RepositoryFiles,
		repositorySubmodulesService:         gl.RepositorySubmodules,
		resourceGroupService:                gl.ResourceGroup,
		resourceIterationEventsService:      gl.ResourceIterationEvents,
		resourceLabelEventsService:          gl.ResourceLabelEvents,
		resourceMilestoneEventsService:      gl.ResourceMilestoneEvents,
		resourceStateEventsService:          gl.ResourceStateEvents,
		resourceWeightEventsService:         gl.ResourceWeightEvents,
		runnersService:                      gl.Runners,
		searchService:                       gl.Search,
		servicesService:                     gl.Services,
		settingsService:                     gl.Settings,
		sidekiqService:                      gl.Sidekiq,
		snippetRepositoryStorageMoveService: gl.SnippetRepositoryStorageMove,
		snippetsService:                     gl.Snippets,
		systemHooksService:                  gl.SystemHooks,
		tagsService:                         gl.Tags,
		todosService:                        gl.Todos,
		topicsService:                       gl.Topics,
		usersService:                        gl.Users,
		validateService:                     gl.Validate,
		versionService:                      gl.Version,
		wikisService:                        gl.Wikis,
	}
}
