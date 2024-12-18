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
	"testing"

	"github.com/jaredallard/gitlab/mocks"
	"go.uber.org/mock/gomock"
)

// _ ensures that [MockClient] implements [Client].
var _ Client = &MockClient{}

// MockClient is a mock of Client interface. Create with
// [NewMockClient].
type MockClient struct {
	AccessRequestsServiceM               *mocks.MockAccessRequestsService
	AppearanceServiceM                   *mocks.MockAppearanceService
	ApplicationsServiceM                 *mocks.MockApplicationsService
	AuditEventsServiceM                  *mocks.MockAuditEventsService
	AvatarRequestsServiceM               *mocks.MockAvatarRequestsService
	AwardEmojiServiceM                   *mocks.MockAwardEmojiService
	BranchesServiceM                     *mocks.MockBranchesService
	BroadcastMessagesServiceM            *mocks.MockBroadcastMessagesService
	CIYMLTemplatesServiceM               *mocks.MockCIYMLTemplatesService
	ClusterAgentsServiceM                *mocks.MockClusterAgentsService
	CommitsServiceM                      *mocks.MockCommitsService
	ContainerRegistryServiceM            *mocks.MockContainerRegistryService
	CustomAttributesServiceM             *mocks.MockCustomAttributesService
	DORAMetricsServiceM                  *mocks.MockDORAMetricsService
	DependencyListExportServiceM         *mocks.MockDependencyListExportService
	DeployKeysServiceM                   *mocks.MockDeployKeysService
	DeployTokensServiceM                 *mocks.MockDeployTokensService
	DeploymentMergeRequestsServiceM      *mocks.MockDeploymentMergeRequestsService
	DeploymentsServiceM                  *mocks.MockDeploymentsService
	DiscussionsServiceM                  *mocks.MockDiscussionsService
	DockerfileTemplatesServiceM          *mocks.MockDockerfileTemplatesService
	DraftNotesServiceM                   *mocks.MockDraftNotesService
	EnvironmentsServiceM                 *mocks.MockEnvironmentsService
	EpicIssuesServiceM                   *mocks.MockEpicIssuesService
	EpicsServiceM                        *mocks.MockEpicsService
	ErrorTrackingServiceM                *mocks.MockErrorTrackingService
	EventsServiceM                       *mocks.MockEventsService
	ExternalStatusChecksServiceM         *mocks.MockExternalStatusChecksService
	FeaturesServiceM                     *mocks.MockFeaturesService
	FreezePeriodsServiceM                *mocks.MockFreezePeriodsService
	GenericPackagesServiceM              *mocks.MockGenericPackagesService
	GeoNodesServiceM                     *mocks.MockGeoNodesService
	GitIgnoreTemplatesServiceM           *mocks.MockGitIgnoreTemplatesService
	GroupAccessTokensServiceM            *mocks.MockGroupAccessTokensService
	GroupBadgesServiceM                  *mocks.MockGroupBadgesService
	GroupClustersServiceM                *mocks.MockGroupClustersService
	GroupEpicBoardsServiceM              *mocks.MockGroupEpicBoardsService
	GroupImportExportServiceM            *mocks.MockGroupImportExportService
	GroupIssueBoardsServiceM             *mocks.MockGroupIssueBoardsService
	GroupIterationsServiceM              *mocks.MockGroupIterationsService
	GroupLabelsServiceM                  *mocks.MockGroupLabelsService
	GroupMembersServiceM                 *mocks.MockGroupMembersService
	GroupMilestonesServiceM              *mocks.MockGroupMilestonesService
	GroupProtectedEnvironmentsServiceM   *mocks.MockGroupProtectedEnvironmentsService
	GroupRepositoryStorageMoveServiceM   *mocks.MockGroupRepositoryStorageMoveService
	GroupSSHCertificatesServiceM         *mocks.MockGroupSSHCertificatesService
	GroupVariablesServiceM               *mocks.MockGroupVariablesService
	GroupWikisServiceM                   *mocks.MockGroupWikisService
	GroupsServiceM                       *mocks.MockGroupsService
	ImportServiceM                       *mocks.MockImportService
	InstanceClustersServiceM             *mocks.MockInstanceClustersService
	InstanceVariablesServiceM            *mocks.MockInstanceVariablesService
	InvitesServiceM                      *mocks.MockInvitesService
	IssueBoardsServiceM                  *mocks.MockIssueBoardsService
	IssueLinksServiceM                   *mocks.MockIssueLinksService
	IssuesServiceM                       *mocks.MockIssuesService
	IssuesStatisticsServiceM             *mocks.MockIssuesStatisticsService
	JobTokenScopeServiceM                *mocks.MockJobTokenScopeService
	JobsServiceM                         *mocks.MockJobsService
	KeysServiceM                         *mocks.MockKeysService
	LabelsServiceM                       *mocks.MockLabelsService
	LicenseServiceM                      *mocks.MockLicenseService
	LicenseTemplatesServiceM             *mocks.MockLicenseTemplatesService
	ManagedLicensesServiceM              *mocks.MockManagedLicensesService
	MarkdownServiceM                     *mocks.MockMarkdownService
	MemberRolesServiceM                  *mocks.MockMemberRolesService
	MergeRequestApprovalsServiceM        *mocks.MockMergeRequestApprovalsService
	MergeRequestsServiceM                *mocks.MockMergeRequestsService
	MergeTrainsServiceM                  *mocks.MockMergeTrainsService
	MetadataServiceM                     *mocks.MockMetadataService
	MilestonesServiceM                   *mocks.MockMilestonesService
	NamespacesServiceM                   *mocks.MockNamespacesService
	NotesServiceM                        *mocks.MockNotesService
	NotificationSettingsServiceM         *mocks.MockNotificationSettingsService
	PackagesServiceM                     *mocks.MockPackagesService
	PagesDomainsServiceM                 *mocks.MockPagesDomainsService
	PagesServiceM                        *mocks.MockPagesService
	PersonalAccessTokensServiceM         *mocks.MockPersonalAccessTokensService
	PipelineSchedulesServiceM            *mocks.MockPipelineSchedulesService
	PipelineTriggersServiceM             *mocks.MockPipelineTriggersService
	PipelinesServiceM                    *mocks.MockPipelinesService
	PlanLimitsServiceM                   *mocks.MockPlanLimitsService
	ProjectAccessTokensServiceM          *mocks.MockProjectAccessTokensService
	ProjectBadgesServiceM                *mocks.MockProjectBadgesService
	ProjectClustersServiceM              *mocks.MockProjectClustersService
	ProjectFeatureFlagServiceM           *mocks.MockProjectFeatureFlagService
	ProjectImportExportServiceM          *mocks.MockProjectImportExportService
	ProjectIterationsServiceM            *mocks.MockProjectIterationsService
	ProjectMarkdownUploadsServiceM       *mocks.MockProjectMarkdownUploadsService
	ProjectMembersServiceM               *mocks.MockProjectMembersService
	ProjectMirrorServiceM                *mocks.MockProjectMirrorService
	ProjectRepositoryStorageMoveServiceM *mocks.MockProjectRepositoryStorageMoveService
	ProjectSnippetsServiceM              *mocks.MockProjectSnippetsService
	ProjectTemplatesServiceM             *mocks.MockProjectTemplatesService
	ProjectVariablesServiceM             *mocks.MockProjectVariablesService
	ProjectVulnerabilitiesServiceM       *mocks.MockProjectVulnerabilitiesService
	ProjectsServiceM                     *mocks.MockProjectsService
	ProtectedBranchesServiceM            *mocks.MockProtectedBranchesService
	ProtectedEnvironmentsServiceM        *mocks.MockProtectedEnvironmentsService
	ProtectedTagsServiceM                *mocks.MockProtectedTagsService
	ReleaseLinksServiceM                 *mocks.MockReleaseLinksService
	ReleasesServiceM                     *mocks.MockReleasesService
	RepositoriesServiceM                 *mocks.MockRepositoriesService
	RepositoryFilesServiceM              *mocks.MockRepositoryFilesService
	RepositorySubmodulesServiceM         *mocks.MockRepositorySubmodulesService
	ResourceGroupServiceM                *mocks.MockResourceGroupService
	ResourceIterationEventsServiceM      *mocks.MockResourceIterationEventsService
	ResourceLabelEventsServiceM          *mocks.MockResourceLabelEventsService
	ResourceMilestoneEventsServiceM      *mocks.MockResourceMilestoneEventsService
	ResourceStateEventsServiceM          *mocks.MockResourceStateEventsService
	ResourceWeightEventsServiceM         *mocks.MockResourceWeightEventsService
	RunnersServiceM                      *mocks.MockRunnersService
	SearchServiceM                       *mocks.MockSearchService
	ServicesServiceM                     *mocks.MockServicesService
	SettingsServiceM                     *mocks.MockSettingsService
	SidekiqServiceM                      *mocks.MockSidekiqService
	SnippetRepositoryStorageMoveServiceM *mocks.MockSnippetRepositoryStorageMoveService
	SnippetsServiceM                     *mocks.MockSnippetsService
	SystemHooksServiceM                  *mocks.MockSystemHooksService
	TagsServiceM                         *mocks.MockTagsService
	TodosServiceM                        *mocks.MockTodosService
	TopicsServiceM                       *mocks.MockTopicsService
	UsersServiceM                        *mocks.MockUsersService
	ValidateServiceM                     *mocks.MockValidateService
	VersionServiceM                      *mocks.MockVersionService
	WikisServiceM                        *mocks.MockWikisService
}

// NewMockClient creates a new mock client for testing. This should be
// used in place of [NewClient] when testing.
func NewMockClient(t *testing.T) *MockClient {
	m := gomock.NewController(t)
	return &MockClient{
		AccessRequestsServiceM:               mocks.NewMockAccessRequestsService(m),
		AppearanceServiceM:                   mocks.NewMockAppearanceService(m),
		ApplicationsServiceM:                 mocks.NewMockApplicationsService(m),
		AuditEventsServiceM:                  mocks.NewMockAuditEventsService(m),
		AvatarRequestsServiceM:               mocks.NewMockAvatarRequestsService(m),
		AwardEmojiServiceM:                   mocks.NewMockAwardEmojiService(m),
		BranchesServiceM:                     mocks.NewMockBranchesService(m),
		BroadcastMessagesServiceM:            mocks.NewMockBroadcastMessagesService(m),
		CIYMLTemplatesServiceM:               mocks.NewMockCIYMLTemplatesService(m),
		ClusterAgentsServiceM:                mocks.NewMockClusterAgentsService(m),
		CommitsServiceM:                      mocks.NewMockCommitsService(m),
		ContainerRegistryServiceM:            mocks.NewMockContainerRegistryService(m),
		CustomAttributesServiceM:             mocks.NewMockCustomAttributesService(m),
		DORAMetricsServiceM:                  mocks.NewMockDORAMetricsService(m),
		DependencyListExportServiceM:         mocks.NewMockDependencyListExportService(m),
		DeployKeysServiceM:                   mocks.NewMockDeployKeysService(m),
		DeployTokensServiceM:                 mocks.NewMockDeployTokensService(m),
		DeploymentMergeRequestsServiceM:      mocks.NewMockDeploymentMergeRequestsService(m),
		DeploymentsServiceM:                  mocks.NewMockDeploymentsService(m),
		DiscussionsServiceM:                  mocks.NewMockDiscussionsService(m),
		DockerfileTemplatesServiceM:          mocks.NewMockDockerfileTemplatesService(m),
		DraftNotesServiceM:                   mocks.NewMockDraftNotesService(m),
		EnvironmentsServiceM:                 mocks.NewMockEnvironmentsService(m),
		EpicIssuesServiceM:                   mocks.NewMockEpicIssuesService(m),
		EpicsServiceM:                        mocks.NewMockEpicsService(m),
		ErrorTrackingServiceM:                mocks.NewMockErrorTrackingService(m),
		EventsServiceM:                       mocks.NewMockEventsService(m),
		ExternalStatusChecksServiceM:         mocks.NewMockExternalStatusChecksService(m),
		FeaturesServiceM:                     mocks.NewMockFeaturesService(m),
		FreezePeriodsServiceM:                mocks.NewMockFreezePeriodsService(m),
		GenericPackagesServiceM:              mocks.NewMockGenericPackagesService(m),
		GeoNodesServiceM:                     mocks.NewMockGeoNodesService(m),
		GitIgnoreTemplatesServiceM:           mocks.NewMockGitIgnoreTemplatesService(m),
		GroupAccessTokensServiceM:            mocks.NewMockGroupAccessTokensService(m),
		GroupBadgesServiceM:                  mocks.NewMockGroupBadgesService(m),
		GroupClustersServiceM:                mocks.NewMockGroupClustersService(m),
		GroupEpicBoardsServiceM:              mocks.NewMockGroupEpicBoardsService(m),
		GroupImportExportServiceM:            mocks.NewMockGroupImportExportService(m),
		GroupIssueBoardsServiceM:             mocks.NewMockGroupIssueBoardsService(m),
		GroupIterationsServiceM:              mocks.NewMockGroupIterationsService(m),
		GroupLabelsServiceM:                  mocks.NewMockGroupLabelsService(m),
		GroupMembersServiceM:                 mocks.NewMockGroupMembersService(m),
		GroupMilestonesServiceM:              mocks.NewMockGroupMilestonesService(m),
		GroupProtectedEnvironmentsServiceM:   mocks.NewMockGroupProtectedEnvironmentsService(m),
		GroupRepositoryStorageMoveServiceM:   mocks.NewMockGroupRepositoryStorageMoveService(m),
		GroupSSHCertificatesServiceM:         mocks.NewMockGroupSSHCertificatesService(m),
		GroupVariablesServiceM:               mocks.NewMockGroupVariablesService(m),
		GroupWikisServiceM:                   mocks.NewMockGroupWikisService(m),
		GroupsServiceM:                       mocks.NewMockGroupsService(m),
		ImportServiceM:                       mocks.NewMockImportService(m),
		InstanceClustersServiceM:             mocks.NewMockInstanceClustersService(m),
		InstanceVariablesServiceM:            mocks.NewMockInstanceVariablesService(m),
		InvitesServiceM:                      mocks.NewMockInvitesService(m),
		IssueBoardsServiceM:                  mocks.NewMockIssueBoardsService(m),
		IssueLinksServiceM:                   mocks.NewMockIssueLinksService(m),
		IssuesServiceM:                       mocks.NewMockIssuesService(m),
		IssuesStatisticsServiceM:             mocks.NewMockIssuesStatisticsService(m),
		JobTokenScopeServiceM:                mocks.NewMockJobTokenScopeService(m),
		JobsServiceM:                         mocks.NewMockJobsService(m),
		KeysServiceM:                         mocks.NewMockKeysService(m),
		LabelsServiceM:                       mocks.NewMockLabelsService(m),
		LicenseServiceM:                      mocks.NewMockLicenseService(m),
		LicenseTemplatesServiceM:             mocks.NewMockLicenseTemplatesService(m),
		ManagedLicensesServiceM:              mocks.NewMockManagedLicensesService(m),
		MarkdownServiceM:                     mocks.NewMockMarkdownService(m),
		MemberRolesServiceM:                  mocks.NewMockMemberRolesService(m),
		MergeRequestApprovalsServiceM:        mocks.NewMockMergeRequestApprovalsService(m),
		MergeRequestsServiceM:                mocks.NewMockMergeRequestsService(m),
		MergeTrainsServiceM:                  mocks.NewMockMergeTrainsService(m),
		MetadataServiceM:                     mocks.NewMockMetadataService(m),
		MilestonesServiceM:                   mocks.NewMockMilestonesService(m),
		NamespacesServiceM:                   mocks.NewMockNamespacesService(m),
		NotesServiceM:                        mocks.NewMockNotesService(m),
		NotificationSettingsServiceM:         mocks.NewMockNotificationSettingsService(m),
		PackagesServiceM:                     mocks.NewMockPackagesService(m),
		PagesDomainsServiceM:                 mocks.NewMockPagesDomainsService(m),
		PagesServiceM:                        mocks.NewMockPagesService(m),
		PersonalAccessTokensServiceM:         mocks.NewMockPersonalAccessTokensService(m),
		PipelineSchedulesServiceM:            mocks.NewMockPipelineSchedulesService(m),
		PipelineTriggersServiceM:             mocks.NewMockPipelineTriggersService(m),
		PipelinesServiceM:                    mocks.NewMockPipelinesService(m),
		PlanLimitsServiceM:                   mocks.NewMockPlanLimitsService(m),
		ProjectAccessTokensServiceM:          mocks.NewMockProjectAccessTokensService(m),
		ProjectBadgesServiceM:                mocks.NewMockProjectBadgesService(m),
		ProjectClustersServiceM:              mocks.NewMockProjectClustersService(m),
		ProjectFeatureFlagServiceM:           mocks.NewMockProjectFeatureFlagService(m),
		ProjectImportExportServiceM:          mocks.NewMockProjectImportExportService(m),
		ProjectIterationsServiceM:            mocks.NewMockProjectIterationsService(m),
		ProjectMarkdownUploadsServiceM:       mocks.NewMockProjectMarkdownUploadsService(m),
		ProjectMembersServiceM:               mocks.NewMockProjectMembersService(m),
		ProjectMirrorServiceM:                mocks.NewMockProjectMirrorService(m),
		ProjectRepositoryStorageMoveServiceM: mocks.NewMockProjectRepositoryStorageMoveService(m),
		ProjectSnippetsServiceM:              mocks.NewMockProjectSnippetsService(m),
		ProjectTemplatesServiceM:             mocks.NewMockProjectTemplatesService(m),
		ProjectVariablesServiceM:             mocks.NewMockProjectVariablesService(m),
		ProjectVulnerabilitiesServiceM:       mocks.NewMockProjectVulnerabilitiesService(m),
		ProjectsServiceM:                     mocks.NewMockProjectsService(m),
		ProtectedBranchesServiceM:            mocks.NewMockProtectedBranchesService(m),
		ProtectedEnvironmentsServiceM:        mocks.NewMockProtectedEnvironmentsService(m),
		ProtectedTagsServiceM:                mocks.NewMockProtectedTagsService(m),
		ReleaseLinksServiceM:                 mocks.NewMockReleaseLinksService(m),
		ReleasesServiceM:                     mocks.NewMockReleasesService(m),
		RepositoriesServiceM:                 mocks.NewMockRepositoriesService(m),
		RepositoryFilesServiceM:              mocks.NewMockRepositoryFilesService(m),
		RepositorySubmodulesServiceM:         mocks.NewMockRepositorySubmodulesService(m),
		ResourceGroupServiceM:                mocks.NewMockResourceGroupService(m),
		ResourceIterationEventsServiceM:      mocks.NewMockResourceIterationEventsService(m),
		ResourceLabelEventsServiceM:          mocks.NewMockResourceLabelEventsService(m),
		ResourceMilestoneEventsServiceM:      mocks.NewMockResourceMilestoneEventsService(m),
		ResourceStateEventsServiceM:          mocks.NewMockResourceStateEventsService(m),
		ResourceWeightEventsServiceM:         mocks.NewMockResourceWeightEventsService(m),
		RunnersServiceM:                      mocks.NewMockRunnersService(m),
		SearchServiceM:                       mocks.NewMockSearchService(m),
		ServicesServiceM:                     mocks.NewMockServicesService(m),
		SettingsServiceM:                     mocks.NewMockSettingsService(m),
		SidekiqServiceM:                      mocks.NewMockSidekiqService(m),
		SnippetRepositoryStorageMoveServiceM: mocks.NewMockSnippetRepositoryStorageMoveService(m),
		SnippetsServiceM:                     mocks.NewMockSnippetsService(m),
		SystemHooksServiceM:                  mocks.NewMockSystemHooksService(m),
		TagsServiceM:                         mocks.NewMockTagsService(m),
		TodosServiceM:                        mocks.NewMockTodosService(m),
		TopicsServiceM:                       mocks.NewMockTopicsService(m),
		UsersServiceM:                        mocks.NewMockUsersService(m),
		ValidateServiceM:                     mocks.NewMockValidateService(m),
		VersionServiceM:                      mocks.NewMockVersionService(m),
		WikisServiceM:                        mocks.NewMockWikisService(m),
	}
}

// AccessRequests returns a mocked [AccessRequestsService] service.
func (m *MockClient) AccessRequests() AccessRequestsService {
	return m.AccessRequestsServiceM
}

// Appearance returns a mocked [AppearanceService] service.
func (m *MockClient) Appearance() AppearanceService {
	return m.AppearanceServiceM
}

// Applications returns a mocked [ApplicationsService] service.
func (m *MockClient) Applications() ApplicationsService {
	return m.ApplicationsServiceM
}

// AuditEvents returns a mocked [AuditEventsService] service.
func (m *MockClient) AuditEvents() AuditEventsService {
	return m.AuditEventsServiceM
}

// Avatar returns a mocked [AvatarRequestsService] service.
func (m *MockClient) Avatar() AvatarRequestsService {
	return m.AvatarRequestsServiceM
}

// AwardEmoji returns a mocked [AwardEmojiService] service.
func (m *MockClient) AwardEmoji() AwardEmojiService {
	return m.AwardEmojiServiceM
}

// Branches returns a mocked [BranchesService] service.
func (m *MockClient) Branches() BranchesService {
	return m.BranchesServiceM
}

// BroadcastMessage returns a mocked [BroadcastMessagesService] service.
func (m *MockClient) BroadcastMessage() BroadcastMessagesService {
	return m.BroadcastMessagesServiceM
}

// CIYMLTemplate returns a mocked [CIYMLTemplatesService] service.
func (m *MockClient) CIYMLTemplate() CIYMLTemplatesService {
	return m.CIYMLTemplatesServiceM
}

// ClusterAgents returns a mocked [ClusterAgentsService] service.
func (m *MockClient) ClusterAgents() ClusterAgentsService {
	return m.ClusterAgentsServiceM
}

// Commits returns a mocked [CommitsService] service.
func (m *MockClient) Commits() CommitsService {
	return m.CommitsServiceM
}

// ContainerRegistry returns a mocked [ContainerRegistryService] service.
func (m *MockClient) ContainerRegistry() ContainerRegistryService {
	return m.ContainerRegistryServiceM
}

// CustomAttribute returns a mocked [CustomAttributesService] service.
func (m *MockClient) CustomAttribute() CustomAttributesService {
	return m.CustomAttributesServiceM
}

// DORAMetrics returns a mocked [DORAMetricsService] service.
func (m *MockClient) DORAMetrics() DORAMetricsService {
	return m.DORAMetricsServiceM
}

// DependencyListExport returns a mocked [DependencyListExportService] service.
func (m *MockClient) DependencyListExport() DependencyListExportService {
	return m.DependencyListExportServiceM
}

// DeployKeys returns a mocked [DeployKeysService] service.
func (m *MockClient) DeployKeys() DeployKeysService {
	return m.DeployKeysServiceM
}

// DeployTokens returns a mocked [DeployTokensService] service.
func (m *MockClient) DeployTokens() DeployTokensService {
	return m.DeployTokensServiceM
}

// DeploymentMergeRequests returns a mocked [DeploymentMergeRequestsService] service.
func (m *MockClient) DeploymentMergeRequests() DeploymentMergeRequestsService {
	return m.DeploymentMergeRequestsServiceM
}

// Deployments returns a mocked [DeploymentsService] service.
func (m *MockClient) Deployments() DeploymentsService {
	return m.DeploymentsServiceM
}

// Discussions returns a mocked [DiscussionsService] service.
func (m *MockClient) Discussions() DiscussionsService {
	return m.DiscussionsServiceM
}

// DockerfileTemplate returns a mocked [DockerfileTemplatesService] service.
func (m *MockClient) DockerfileTemplate() DockerfileTemplatesService {
	return m.DockerfileTemplatesServiceM
}

// DraftNotes returns a mocked [DraftNotesService] service.
func (m *MockClient) DraftNotes() DraftNotesService {
	return m.DraftNotesServiceM
}

// Environments returns a mocked [EnvironmentsService] service.
func (m *MockClient) Environments() EnvironmentsService {
	return m.EnvironmentsServiceM
}

// EpicIssues returns a mocked [EpicIssuesService] service.
func (m *MockClient) EpicIssues() EpicIssuesService {
	return m.EpicIssuesServiceM
}

// Epics returns a mocked [EpicsService] service.
func (m *MockClient) Epics() EpicsService {
	return m.EpicsServiceM
}

// ErrorTracking returns a mocked [ErrorTrackingService] service.
func (m *MockClient) ErrorTracking() ErrorTrackingService {
	return m.ErrorTrackingServiceM
}

// Events returns a mocked [EventsService] service.
func (m *MockClient) Events() EventsService {
	return m.EventsServiceM
}

// ExternalStatusChecks returns a mocked [ExternalStatusChecksService] service.
func (m *MockClient) ExternalStatusChecks() ExternalStatusChecksService {
	return m.ExternalStatusChecksServiceM
}

// Features returns a mocked [FeaturesService] service.
func (m *MockClient) Features() FeaturesService {
	return m.FeaturesServiceM
}

// FreezePeriods returns a mocked [FreezePeriodsService] service.
func (m *MockClient) FreezePeriods() FreezePeriodsService {
	return m.FreezePeriodsServiceM
}

// GenericPackages returns a mocked [GenericPackagesService] service.
func (m *MockClient) GenericPackages() GenericPackagesService {
	return m.GenericPackagesServiceM
}

// GeoNodes returns a mocked [GeoNodesService] service.
func (m *MockClient) GeoNodes() GeoNodesService {
	return m.GeoNodesServiceM
}

// GitIgnoreTemplates returns a mocked [GitIgnoreTemplatesService] service.
func (m *MockClient) GitIgnoreTemplates() GitIgnoreTemplatesService {
	return m.GitIgnoreTemplatesServiceM
}

// GroupAccessTokens returns a mocked [GroupAccessTokensService] service.
func (m *MockClient) GroupAccessTokens() GroupAccessTokensService {
	return m.GroupAccessTokensServiceM
}

// GroupBadges returns a mocked [GroupBadgesService] service.
func (m *MockClient) GroupBadges() GroupBadgesService {
	return m.GroupBadgesServiceM
}

// GroupCluster returns a mocked [GroupClustersService] service.
func (m *MockClient) GroupCluster() GroupClustersService {
	return m.GroupClustersServiceM
}

// GroupEpicBoards returns a mocked [GroupEpicBoardsService] service.
func (m *MockClient) GroupEpicBoards() GroupEpicBoardsService {
	return m.GroupEpicBoardsServiceM
}

// GroupImportExport returns a mocked [GroupImportExportService] service.
func (m *MockClient) GroupImportExport() GroupImportExportService {
	return m.GroupImportExportServiceM
}

// GroupIssueBoards returns a mocked [GroupIssueBoardsService] service.
func (m *MockClient) GroupIssueBoards() GroupIssueBoardsService {
	return m.GroupIssueBoardsServiceM
}

// GroupIterations returns a mocked [GroupIterationsService] service.
func (m *MockClient) GroupIterations() GroupIterationsService {
	return m.GroupIterationsServiceM
}

// GroupLabels returns a mocked [GroupLabelsService] service.
func (m *MockClient) GroupLabels() GroupLabelsService {
	return m.GroupLabelsServiceM
}

// GroupMembers returns a mocked [GroupMembersService] service.
func (m *MockClient) GroupMembers() GroupMembersService {
	return m.GroupMembersServiceM
}

// GroupMilestones returns a mocked [GroupMilestonesService] service.
func (m *MockClient) GroupMilestones() GroupMilestonesService {
	return m.GroupMilestonesServiceM
}

// GroupProtectedEnvironments returns a mocked [GroupProtectedEnvironmentsService] service.
func (m *MockClient) GroupProtectedEnvironments() GroupProtectedEnvironmentsService {
	return m.GroupProtectedEnvironmentsServiceM
}

// GroupRepositoryStorageMove returns a mocked [GroupRepositoryStorageMoveService] service.
func (m *MockClient) GroupRepositoryStorageMove() GroupRepositoryStorageMoveService {
	return m.GroupRepositoryStorageMoveServiceM
}

// GroupSSHCertificates returns a mocked [GroupSSHCertificatesService] service.
func (m *MockClient) GroupSSHCertificates() GroupSSHCertificatesService {
	return m.GroupSSHCertificatesServiceM
}

// GroupVariables returns a mocked [GroupVariablesService] service.
func (m *MockClient) GroupVariables() GroupVariablesService {
	return m.GroupVariablesServiceM
}

// GroupWikis returns a mocked [GroupWikisService] service.
func (m *MockClient) GroupWikis() GroupWikisService {
	return m.GroupWikisServiceM
}

// Groups returns a mocked [GroupsService] service.
func (m *MockClient) Groups() GroupsService {
	return m.GroupsServiceM
}

// Import returns a mocked [ImportService] service.
func (m *MockClient) Import() ImportService {
	return m.ImportServiceM
}

// InstanceCluster returns a mocked [InstanceClustersService] service.
func (m *MockClient) InstanceCluster() InstanceClustersService {
	return m.InstanceClustersServiceM
}

// InstanceVariables returns a mocked [InstanceVariablesService] service.
func (m *MockClient) InstanceVariables() InstanceVariablesService {
	return m.InstanceVariablesServiceM
}

// Invites returns a mocked [InvitesService] service.
func (m *MockClient) Invites() InvitesService {
	return m.InvitesServiceM
}

// Boards returns a mocked [IssueBoardsService] service.
func (m *MockClient) Boards() IssueBoardsService {
	return m.IssueBoardsServiceM
}

// IssueLinks returns a mocked [IssueLinksService] service.
func (m *MockClient) IssueLinks() IssueLinksService {
	return m.IssueLinksServiceM
}

// Issues returns a mocked [IssuesService] service.
func (m *MockClient) Issues() IssuesService {
	return m.IssuesServiceM
}

// IssuesStatistics returns a mocked [IssuesStatisticsService] service.
func (m *MockClient) IssuesStatistics() IssuesStatisticsService {
	return m.IssuesStatisticsServiceM
}

// JobTokenScope returns a mocked [JobTokenScopeService] service.
func (m *MockClient) JobTokenScope() JobTokenScopeService {
	return m.JobTokenScopeServiceM
}

// Jobs returns a mocked [JobsService] service.
func (m *MockClient) Jobs() JobsService {
	return m.JobsServiceM
}

// Keys returns a mocked [KeysService] service.
func (m *MockClient) Keys() KeysService {
	return m.KeysServiceM
}

// Labels returns a mocked [LabelsService] service.
func (m *MockClient) Labels() LabelsService {
	return m.LabelsServiceM
}

// License returns a mocked [LicenseService] service.
func (m *MockClient) License() LicenseService {
	return m.LicenseServiceM
}

// LicenseTemplates returns a mocked [LicenseTemplatesService] service.
func (m *MockClient) LicenseTemplates() LicenseTemplatesService {
	return m.LicenseTemplatesServiceM
}

// ManagedLicenses returns a mocked [ManagedLicensesService] service.
func (m *MockClient) ManagedLicenses() ManagedLicensesService {
	return m.ManagedLicensesServiceM
}

// Markdown returns a mocked [MarkdownService] service.
func (m *MockClient) Markdown() MarkdownService {
	return m.MarkdownServiceM
}

// MemberRolesService returns a mocked [MemberRolesService] service.
func (m *MockClient) MemberRolesService() MemberRolesService {
	return m.MemberRolesServiceM
}

// MergeRequestApprovals returns a mocked [MergeRequestApprovalsService] service.
func (m *MockClient) MergeRequestApprovals() MergeRequestApprovalsService {
	return m.MergeRequestApprovalsServiceM
}

// MergeRequests returns a mocked [MergeRequestsService] service.
func (m *MockClient) MergeRequests() MergeRequestsService {
	return m.MergeRequestsServiceM
}

// MergeTrains returns a mocked [MergeTrainsService] service.
func (m *MockClient) MergeTrains() MergeTrainsService {
	return m.MergeTrainsServiceM
}

// Metadata returns a mocked [MetadataService] service.
func (m *MockClient) Metadata() MetadataService {
	return m.MetadataServiceM
}

// Milestones returns a mocked [MilestonesService] service.
func (m *MockClient) Milestones() MilestonesService {
	return m.MilestonesServiceM
}

// Namespaces returns a mocked [NamespacesService] service.
func (m *MockClient) Namespaces() NamespacesService {
	return m.NamespacesServiceM
}

// Notes returns a mocked [NotesService] service.
func (m *MockClient) Notes() NotesService {
	return m.NotesServiceM
}

// NotificationSettings returns a mocked [NotificationSettingsService] service.
func (m *MockClient) NotificationSettings() NotificationSettingsService {
	return m.NotificationSettingsServiceM
}

// Packages returns a mocked [PackagesService] service.
func (m *MockClient) Packages() PackagesService {
	return m.PackagesServiceM
}

// PagesDomains returns a mocked [PagesDomainsService] service.
func (m *MockClient) PagesDomains() PagesDomainsService {
	return m.PagesDomainsServiceM
}

// Pages returns a mocked [PagesService] service.
func (m *MockClient) Pages() PagesService {
	return m.PagesServiceM
}

// PersonalAccessTokens returns a mocked [PersonalAccessTokensService] service.
func (m *MockClient) PersonalAccessTokens() PersonalAccessTokensService {
	return m.PersonalAccessTokensServiceM
}

// PipelineSchedules returns a mocked [PipelineSchedulesService] service.
func (m *MockClient) PipelineSchedules() PipelineSchedulesService {
	return m.PipelineSchedulesServiceM
}

// PipelineTriggers returns a mocked [PipelineTriggersService] service.
func (m *MockClient) PipelineTriggers() PipelineTriggersService {
	return m.PipelineTriggersServiceM
}

// Pipelines returns a mocked [PipelinesService] service.
func (m *MockClient) Pipelines() PipelinesService {
	return m.PipelinesServiceM
}

// PlanLimits returns a mocked [PlanLimitsService] service.
func (m *MockClient) PlanLimits() PlanLimitsService {
	return m.PlanLimitsServiceM
}

// ProjectAccessTokens returns a mocked [ProjectAccessTokensService] service.
func (m *MockClient) ProjectAccessTokens() ProjectAccessTokensService {
	return m.ProjectAccessTokensServiceM
}

// ProjectBadges returns a mocked [ProjectBadgesService] service.
func (m *MockClient) ProjectBadges() ProjectBadgesService {
	return m.ProjectBadgesServiceM
}

// ProjectCluster returns a mocked [ProjectClustersService] service.
func (m *MockClient) ProjectCluster() ProjectClustersService {
	return m.ProjectClustersServiceM
}

// ProjectFeatureFlags returns a mocked [ProjectFeatureFlagService] service.
func (m *MockClient) ProjectFeatureFlags() ProjectFeatureFlagService {
	return m.ProjectFeatureFlagServiceM
}

// ProjectImportExport returns a mocked [ProjectImportExportService] service.
func (m *MockClient) ProjectImportExport() ProjectImportExportService {
	return m.ProjectImportExportServiceM
}

// ProjectIterations returns a mocked [ProjectIterationsService] service.
func (m *MockClient) ProjectIterations() ProjectIterationsService {
	return m.ProjectIterationsServiceM
}

// ProjectMarkdownUploads returns a mocked [ProjectMarkdownUploadsService] service.
func (m *MockClient) ProjectMarkdownUploads() ProjectMarkdownUploadsService {
	return m.ProjectMarkdownUploadsServiceM
}

// ProjectMembers returns a mocked [ProjectMembersService] service.
func (m *MockClient) ProjectMembers() ProjectMembersService {
	return m.ProjectMembersServiceM
}

// ProjectMirrors returns a mocked [ProjectMirrorService] service.
func (m *MockClient) ProjectMirrors() ProjectMirrorService {
	return m.ProjectMirrorServiceM
}

// ProjectRepositoryStorageMove returns a mocked [ProjectRepositoryStorageMoveService] service.
func (m *MockClient) ProjectRepositoryStorageMove() ProjectRepositoryStorageMoveService {
	return m.ProjectRepositoryStorageMoveServiceM
}

// ProjectSnippets returns a mocked [ProjectSnippetsService] service.
func (m *MockClient) ProjectSnippets() ProjectSnippetsService {
	return m.ProjectSnippetsServiceM
}

// ProjectTemplates returns a mocked [ProjectTemplatesService] service.
func (m *MockClient) ProjectTemplates() ProjectTemplatesService {
	return m.ProjectTemplatesServiceM
}

// ProjectVariables returns a mocked [ProjectVariablesService] service.
func (m *MockClient) ProjectVariables() ProjectVariablesService {
	return m.ProjectVariablesServiceM
}

// ProjectVulnerabilities returns a mocked [ProjectVulnerabilitiesService] service.
func (m *MockClient) ProjectVulnerabilities() ProjectVulnerabilitiesService {
	return m.ProjectVulnerabilitiesServiceM
}

// Projects returns a mocked [ProjectsService] service.
func (m *MockClient) Projects() ProjectsService {
	return m.ProjectsServiceM
}

// ProtectedBranches returns a mocked [ProtectedBranchesService] service.
func (m *MockClient) ProtectedBranches() ProtectedBranchesService {
	return m.ProtectedBranchesServiceM
}

// ProtectedEnvironments returns a mocked [ProtectedEnvironmentsService] service.
func (m *MockClient) ProtectedEnvironments() ProtectedEnvironmentsService {
	return m.ProtectedEnvironmentsServiceM
}

// ProtectedTags returns a mocked [ProtectedTagsService] service.
func (m *MockClient) ProtectedTags() ProtectedTagsService {
	return m.ProtectedTagsServiceM
}

// ReleaseLinks returns a mocked [ReleaseLinksService] service.
func (m *MockClient) ReleaseLinks() ReleaseLinksService {
	return m.ReleaseLinksServiceM
}

// Releases returns a mocked [ReleasesService] service.
func (m *MockClient) Releases() ReleasesService {
	return m.ReleasesServiceM
}

// Repositories returns a mocked [RepositoriesService] service.
func (m *MockClient) Repositories() RepositoriesService {
	return m.RepositoriesServiceM
}

// RepositoryFiles returns a mocked [RepositoryFilesService] service.
func (m *MockClient) RepositoryFiles() RepositoryFilesService {
	return m.RepositoryFilesServiceM
}

// RepositorySubmodules returns a mocked [RepositorySubmodulesService] service.
func (m *MockClient) RepositorySubmodules() RepositorySubmodulesService {
	return m.RepositorySubmodulesServiceM
}

// ResourceGroup returns a mocked [ResourceGroupService] service.
func (m *MockClient) ResourceGroup() ResourceGroupService {
	return m.ResourceGroupServiceM
}

// ResourceIterationEvents returns a mocked [ResourceIterationEventsService] service.
func (m *MockClient) ResourceIterationEvents() ResourceIterationEventsService {
	return m.ResourceIterationEventsServiceM
}

// ResourceLabelEvents returns a mocked [ResourceLabelEventsService] service.
func (m *MockClient) ResourceLabelEvents() ResourceLabelEventsService {
	return m.ResourceLabelEventsServiceM
}

// ResourceMilestoneEvents returns a mocked [ResourceMilestoneEventsService] service.
func (m *MockClient) ResourceMilestoneEvents() ResourceMilestoneEventsService {
	return m.ResourceMilestoneEventsServiceM
}

// ResourceStateEvents returns a mocked [ResourceStateEventsService] service.
func (m *MockClient) ResourceStateEvents() ResourceStateEventsService {
	return m.ResourceStateEventsServiceM
}

// ResourceWeightEvents returns a mocked [ResourceWeightEventsService] service.
func (m *MockClient) ResourceWeightEvents() ResourceWeightEventsService {
	return m.ResourceWeightEventsServiceM
}

// Runners returns a mocked [RunnersService] service.
func (m *MockClient) Runners() RunnersService {
	return m.RunnersServiceM
}

// Search returns a mocked [SearchService] service.
func (m *MockClient) Search() SearchService {
	return m.SearchServiceM
}

// Services returns a mocked [ServicesService] service.
func (m *MockClient) Services() ServicesService {
	return m.ServicesServiceM
}

// Settings returns a mocked [SettingsService] service.
func (m *MockClient) Settings() SettingsService {
	return m.SettingsServiceM
}

// Sidekiq returns a mocked [SidekiqService] service.
func (m *MockClient) Sidekiq() SidekiqService {
	return m.SidekiqServiceM
}

// SnippetRepositoryStorageMove returns a mocked [SnippetRepositoryStorageMoveService] service.
func (m *MockClient) SnippetRepositoryStorageMove() SnippetRepositoryStorageMoveService {
	return m.SnippetRepositoryStorageMoveServiceM
}

// Snippets returns a mocked [SnippetsService] service.
func (m *MockClient) Snippets() SnippetsService {
	return m.SnippetsServiceM
}

// SystemHooks returns a mocked [SystemHooksService] service.
func (m *MockClient) SystemHooks() SystemHooksService {
	return m.SystemHooksServiceM
}

// Tags returns a mocked [TagsService] service.
func (m *MockClient) Tags() TagsService {
	return m.TagsServiceM
}

// Todos returns a mocked [TodosService] service.
func (m *MockClient) Todos() TodosService {
	return m.TodosServiceM
}

// Topics returns a mocked [TopicsService] service.
func (m *MockClient) Topics() TopicsService {
	return m.TopicsServiceM
}

// Users returns a mocked [UsersService] service.
func (m *MockClient) Users() UsersService {
	return m.UsersServiceM
}

// Validate returns a mocked [ValidateService] service.
func (m *MockClient) Validate() ValidateService {
	return m.ValidateServiceM
}

// Version returns a mocked [VersionService] service.
func (m *MockClient) Version() VersionService {
	return m.VersionServiceM
}

// Wikis returns a mocked [WikisService] service.
func (m *MockClient) Wikis() WikisService {
	return m.WikisServiceM
}
