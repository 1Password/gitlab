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
	AccessRequestsServiceMock               *mocks.MockAccessRequestsService
	AppearanceServiceMock                   *mocks.MockAppearanceService
	ApplicationsServiceMock                 *mocks.MockApplicationsService
	AuditEventsServiceMock                  *mocks.MockAuditEventsService
	AvatarRequestsServiceMock               *mocks.MockAvatarRequestsService
	AwardEmojiServiceMock                   *mocks.MockAwardEmojiService
	BranchesServiceMock                     *mocks.MockBranchesService
	BroadcastMessagesServiceMock            *mocks.MockBroadcastMessagesService
	CIYMLTemplatesServiceMock               *mocks.MockCIYMLTemplatesService
	ClusterAgentsServiceMock                *mocks.MockClusterAgentsService
	CommitsServiceMock                      *mocks.MockCommitsService
	ContainerRegistryServiceMock            *mocks.MockContainerRegistryService
	CustomAttributesServiceMock             *mocks.MockCustomAttributesService
	DORAMetricsServiceMock                  *mocks.MockDORAMetricsService
	DeployKeysServiceMock                   *mocks.MockDeployKeysService
	DeployTokensServiceMock                 *mocks.MockDeployTokensService
	DeploymentMergeRequestsServiceMock      *mocks.MockDeploymentMergeRequestsService
	DeploymentsServiceMock                  *mocks.MockDeploymentsService
	DiscussionsServiceMock                  *mocks.MockDiscussionsService
	DockerfileTemplatesServiceMock          *mocks.MockDockerfileTemplatesService
	DraftNotesServiceMock                   *mocks.MockDraftNotesService
	EnvironmentsServiceMock                 *mocks.MockEnvironmentsService
	EpicIssuesServiceMock                   *mocks.MockEpicIssuesService
	EpicsServiceMock                        *mocks.MockEpicsService
	ErrorTrackingServiceMock                *mocks.MockErrorTrackingService
	EventsServiceMock                       *mocks.MockEventsService
	ExternalStatusChecksServiceMock         *mocks.MockExternalStatusChecksService
	FeaturesServiceMock                     *mocks.MockFeaturesService
	FreezePeriodsServiceMock                *mocks.MockFreezePeriodsService
	GenericPackagesServiceMock              *mocks.MockGenericPackagesService
	GeoNodesServiceMock                     *mocks.MockGeoNodesService
	GitIgnoreTemplatesServiceMock           *mocks.MockGitIgnoreTemplatesService
	GroupAccessTokensServiceMock            *mocks.MockGroupAccessTokensService
	GroupBadgesServiceMock                  *mocks.MockGroupBadgesService
	GroupClustersServiceMock                *mocks.MockGroupClustersService
	GroupEpicBoardsServiceMock              *mocks.MockGroupEpicBoardsService
	GroupImportExportServiceMock            *mocks.MockGroupImportExportService
	GroupIssueBoardsServiceMock             *mocks.MockGroupIssueBoardsService
	GroupIterationsServiceMock              *mocks.MockGroupIterationsService
	GroupLabelsServiceMock                  *mocks.MockGroupLabelsService
	GroupMembersServiceMock                 *mocks.MockGroupMembersService
	GroupMilestonesServiceMock              *mocks.MockGroupMilestonesService
	GroupProtectedEnvironmentsServiceMock   *mocks.MockGroupProtectedEnvironmentsService
	GroupRepositoryStorageMoveServiceMock   *mocks.MockGroupRepositoryStorageMoveService
	GroupSSHCertificatesServiceMock         *mocks.MockGroupSSHCertificatesService
	GroupVariablesServiceMock               *mocks.MockGroupVariablesService
	GroupWikisServiceMock                   *mocks.MockGroupWikisService
	GroupsServiceMock                       *mocks.MockGroupsService
	ImportServiceMock                       *mocks.MockImportService
	InstanceClustersServiceMock             *mocks.MockInstanceClustersService
	InstanceVariablesServiceMock            *mocks.MockInstanceVariablesService
	InvitesServiceMock                      *mocks.MockInvitesService
	IssueBoardsServiceMock                  *mocks.MockIssueBoardsService
	IssueLinksServiceMock                   *mocks.MockIssueLinksService
	IssuesServiceMock                       *mocks.MockIssuesService
	IssuesStatisticsServiceMock             *mocks.MockIssuesStatisticsService
	JobTokenScopeServiceMock                *mocks.MockJobTokenScopeService
	JobsServiceMock                         *mocks.MockJobsService
	KeysServiceMock                         *mocks.MockKeysService
	LabelsServiceMock                       *mocks.MockLabelsService
	LicenseServiceMock                      *mocks.MockLicenseService
	LicenseTemplatesServiceMock             *mocks.MockLicenseTemplatesService
	ManagedLicensesServiceMock              *mocks.MockManagedLicensesService
	MarkdownServiceMock                     *mocks.MockMarkdownService
	MemberRolesServiceMock                  *mocks.MockMemberRolesService
	MergeRequestApprovalsServiceMock        *mocks.MockMergeRequestApprovalsService
	MergeRequestsServiceMock                *mocks.MockMergeRequestsService
	MergeTrainsServiceMock                  *mocks.MockMergeTrainsService
	MetadataServiceMock                     *mocks.MockMetadataService
	MilestonesServiceMock                   *mocks.MockMilestonesService
	NamespacesServiceMock                   *mocks.MockNamespacesService
	NotesServiceMock                        *mocks.MockNotesService
	NotificationSettingsServiceMock         *mocks.MockNotificationSettingsService
	PackagesServiceMock                     *mocks.MockPackagesService
	PagesDomainsServiceMock                 *mocks.MockPagesDomainsService
	PagesServiceMock                        *mocks.MockPagesService
	PersonalAccessTokensServiceMock         *mocks.MockPersonalAccessTokensService
	PipelineSchedulesServiceMock            *mocks.MockPipelineSchedulesService
	PipelineTriggersServiceMock             *mocks.MockPipelineTriggersService
	PipelinesServiceMock                    *mocks.MockPipelinesService
	PlanLimitsServiceMock                   *mocks.MockPlanLimitsService
	ProjectAccessTokensServiceMock          *mocks.MockProjectAccessTokensService
	ProjectBadgesServiceMock                *mocks.MockProjectBadgesService
	ProjectClustersServiceMock              *mocks.MockProjectClustersService
	ProjectFeatureFlagServiceMock           *mocks.MockProjectFeatureFlagService
	ProjectImportExportServiceMock          *mocks.MockProjectImportExportService
	ProjectIterationsServiceMock            *mocks.MockProjectIterationsService
	ProjectMembersServiceMock               *mocks.MockProjectMembersService
	ProjectMirrorServiceMock                *mocks.MockProjectMirrorService
	ProjectRepositoryStorageMoveServiceMock *mocks.MockProjectRepositoryStorageMoveService
	ProjectSnippetsServiceMock              *mocks.MockProjectSnippetsService
	ProjectTemplatesServiceMock             *mocks.MockProjectTemplatesService
	ProjectVariablesServiceMock             *mocks.MockProjectVariablesService
	ProjectVulnerabilitiesServiceMock       *mocks.MockProjectVulnerabilitiesService
	ProjectsServiceMock                     *mocks.MockProjectsService
	ProtectedBranchesServiceMock            *mocks.MockProtectedBranchesService
	ProtectedEnvironmentsServiceMock        *mocks.MockProtectedEnvironmentsService
	ProtectedTagsServiceMock                *mocks.MockProtectedTagsService
	ReleaseLinksServiceMock                 *mocks.MockReleaseLinksService
	ReleasesServiceMock                     *mocks.MockReleasesService
	RepositoriesServiceMock                 *mocks.MockRepositoriesService
	RepositoryFilesServiceMock              *mocks.MockRepositoryFilesService
	RepositorySubmodulesServiceMock         *mocks.MockRepositorySubmodulesService
	ResourceGroupServiceMock                *mocks.MockResourceGroupService
	ResourceIterationEventsServiceMock      *mocks.MockResourceIterationEventsService
	ResourceLabelEventsServiceMock          *mocks.MockResourceLabelEventsService
	ResourceMilestoneEventsServiceMock      *mocks.MockResourceMilestoneEventsService
	ResourceStateEventsServiceMock          *mocks.MockResourceStateEventsService
	ResourceWeightEventsServiceMock         *mocks.MockResourceWeightEventsService
	RunnersServiceMock                      *mocks.MockRunnersService
	SearchServiceMock                       *mocks.MockSearchService
	ServicesServiceMock                     *mocks.MockServicesService
	SettingsServiceMock                     *mocks.MockSettingsService
	SidekiqServiceMock                      *mocks.MockSidekiqService
	SnippetRepositoryStorageMoveServiceMock *mocks.MockSnippetRepositoryStorageMoveService
	SnippetsServiceMock                     *mocks.MockSnippetsService
	SystemHooksServiceMock                  *mocks.MockSystemHooksService
	TagsServiceMock                         *mocks.MockTagsService
	TodosServiceMock                        *mocks.MockTodosService
	TopicsServiceMock                       *mocks.MockTopicsService
	UsersServiceMock                        *mocks.MockUsersService
	ValidateServiceMock                     *mocks.MockValidateService
	VersionServiceMock                      *mocks.MockVersionService
	WikisServiceMock                        *mocks.MockWikisService
}

// NewMockClient creates a new mock client for testing. This should be
// used in place of [NewClient] when testing.
func NewMockClient(t *testing.T) *MockClient {
	m := gomock.NewController(t)
	return &MockClient{
		AccessRequestsServiceMock:               mocks.NewMockAccessRequestsService(m),
		AppearanceServiceMock:                   mocks.NewMockAppearanceService(m),
		ApplicationsServiceMock:                 mocks.NewMockApplicationsService(m),
		AuditEventsServiceMock:                  mocks.NewMockAuditEventsService(m),
		AvatarRequestsServiceMock:               mocks.NewMockAvatarRequestsService(m),
		AwardEmojiServiceMock:                   mocks.NewMockAwardEmojiService(m),
		BranchesServiceMock:                     mocks.NewMockBranchesService(m),
		BroadcastMessagesServiceMock:            mocks.NewMockBroadcastMessagesService(m),
		CIYMLTemplatesServiceMock:               mocks.NewMockCIYMLTemplatesService(m),
		ClusterAgentsServiceMock:                mocks.NewMockClusterAgentsService(m),
		CommitsServiceMock:                      mocks.NewMockCommitsService(m),
		ContainerRegistryServiceMock:            mocks.NewMockContainerRegistryService(m),
		CustomAttributesServiceMock:             mocks.NewMockCustomAttributesService(m),
		DORAMetricsServiceMock:                  mocks.NewMockDORAMetricsService(m),
		DeployKeysServiceMock:                   mocks.NewMockDeployKeysService(m),
		DeployTokensServiceMock:                 mocks.NewMockDeployTokensService(m),
		DeploymentMergeRequestsServiceMock:      mocks.NewMockDeploymentMergeRequestsService(m),
		DeploymentsServiceMock:                  mocks.NewMockDeploymentsService(m),
		DiscussionsServiceMock:                  mocks.NewMockDiscussionsService(m),
		DockerfileTemplatesServiceMock:          mocks.NewMockDockerfileTemplatesService(m),
		DraftNotesServiceMock:                   mocks.NewMockDraftNotesService(m),
		EnvironmentsServiceMock:                 mocks.NewMockEnvironmentsService(m),
		EpicIssuesServiceMock:                   mocks.NewMockEpicIssuesService(m),
		EpicsServiceMock:                        mocks.NewMockEpicsService(m),
		ErrorTrackingServiceMock:                mocks.NewMockErrorTrackingService(m),
		EventsServiceMock:                       mocks.NewMockEventsService(m),
		ExternalStatusChecksServiceMock:         mocks.NewMockExternalStatusChecksService(m),
		FeaturesServiceMock:                     mocks.NewMockFeaturesService(m),
		FreezePeriodsServiceMock:                mocks.NewMockFreezePeriodsService(m),
		GenericPackagesServiceMock:              mocks.NewMockGenericPackagesService(m),
		GeoNodesServiceMock:                     mocks.NewMockGeoNodesService(m),
		GitIgnoreTemplatesServiceMock:           mocks.NewMockGitIgnoreTemplatesService(m),
		GroupAccessTokensServiceMock:            mocks.NewMockGroupAccessTokensService(m),
		GroupBadgesServiceMock:                  mocks.NewMockGroupBadgesService(m),
		GroupClustersServiceMock:                mocks.NewMockGroupClustersService(m),
		GroupEpicBoardsServiceMock:              mocks.NewMockGroupEpicBoardsService(m),
		GroupImportExportServiceMock:            mocks.NewMockGroupImportExportService(m),
		GroupIssueBoardsServiceMock:             mocks.NewMockGroupIssueBoardsService(m),
		GroupIterationsServiceMock:              mocks.NewMockGroupIterationsService(m),
		GroupLabelsServiceMock:                  mocks.NewMockGroupLabelsService(m),
		GroupMembersServiceMock:                 mocks.NewMockGroupMembersService(m),
		GroupMilestonesServiceMock:              mocks.NewMockGroupMilestonesService(m),
		GroupProtectedEnvironmentsServiceMock:   mocks.NewMockGroupProtectedEnvironmentsService(m),
		GroupRepositoryStorageMoveServiceMock:   mocks.NewMockGroupRepositoryStorageMoveService(m),
		GroupSSHCertificatesServiceMock:         mocks.NewMockGroupSSHCertificatesService(m),
		GroupVariablesServiceMock:               mocks.NewMockGroupVariablesService(m),
		GroupWikisServiceMock:                   mocks.NewMockGroupWikisService(m),
		GroupsServiceMock:                       mocks.NewMockGroupsService(m),
		ImportServiceMock:                       mocks.NewMockImportService(m),
		InstanceClustersServiceMock:             mocks.NewMockInstanceClustersService(m),
		InstanceVariablesServiceMock:            mocks.NewMockInstanceVariablesService(m),
		InvitesServiceMock:                      mocks.NewMockInvitesService(m),
		IssueBoardsServiceMock:                  mocks.NewMockIssueBoardsService(m),
		IssueLinksServiceMock:                   mocks.NewMockIssueLinksService(m),
		IssuesServiceMock:                       mocks.NewMockIssuesService(m),
		IssuesStatisticsServiceMock:             mocks.NewMockIssuesStatisticsService(m),
		JobTokenScopeServiceMock:                mocks.NewMockJobTokenScopeService(m),
		JobsServiceMock:                         mocks.NewMockJobsService(m),
		KeysServiceMock:                         mocks.NewMockKeysService(m),
		LabelsServiceMock:                       mocks.NewMockLabelsService(m),
		LicenseServiceMock:                      mocks.NewMockLicenseService(m),
		LicenseTemplatesServiceMock:             mocks.NewMockLicenseTemplatesService(m),
		ManagedLicensesServiceMock:              mocks.NewMockManagedLicensesService(m),
		MarkdownServiceMock:                     mocks.NewMockMarkdownService(m),
		MemberRolesServiceMock:                  mocks.NewMockMemberRolesService(m),
		MergeRequestApprovalsServiceMock:        mocks.NewMockMergeRequestApprovalsService(m),
		MergeRequestsServiceMock:                mocks.NewMockMergeRequestsService(m),
		MergeTrainsServiceMock:                  mocks.NewMockMergeTrainsService(m),
		MetadataServiceMock:                     mocks.NewMockMetadataService(m),
		MilestonesServiceMock:                   mocks.NewMockMilestonesService(m),
		NamespacesServiceMock:                   mocks.NewMockNamespacesService(m),
		NotesServiceMock:                        mocks.NewMockNotesService(m),
		NotificationSettingsServiceMock:         mocks.NewMockNotificationSettingsService(m),
		PackagesServiceMock:                     mocks.NewMockPackagesService(m),
		PagesDomainsServiceMock:                 mocks.NewMockPagesDomainsService(m),
		PagesServiceMock:                        mocks.NewMockPagesService(m),
		PersonalAccessTokensServiceMock:         mocks.NewMockPersonalAccessTokensService(m),
		PipelineSchedulesServiceMock:            mocks.NewMockPipelineSchedulesService(m),
		PipelineTriggersServiceMock:             mocks.NewMockPipelineTriggersService(m),
		PipelinesServiceMock:                    mocks.NewMockPipelinesService(m),
		PlanLimitsServiceMock:                   mocks.NewMockPlanLimitsService(m),
		ProjectAccessTokensServiceMock:          mocks.NewMockProjectAccessTokensService(m),
		ProjectBadgesServiceMock:                mocks.NewMockProjectBadgesService(m),
		ProjectClustersServiceMock:              mocks.NewMockProjectClustersService(m),
		ProjectFeatureFlagServiceMock:           mocks.NewMockProjectFeatureFlagService(m),
		ProjectImportExportServiceMock:          mocks.NewMockProjectImportExportService(m),
		ProjectIterationsServiceMock:            mocks.NewMockProjectIterationsService(m),
		ProjectMembersServiceMock:               mocks.NewMockProjectMembersService(m),
		ProjectMirrorServiceMock:                mocks.NewMockProjectMirrorService(m),
		ProjectRepositoryStorageMoveServiceMock: mocks.NewMockProjectRepositoryStorageMoveService(m),
		ProjectSnippetsServiceMock:              mocks.NewMockProjectSnippetsService(m),
		ProjectTemplatesServiceMock:             mocks.NewMockProjectTemplatesService(m),
		ProjectVariablesServiceMock:             mocks.NewMockProjectVariablesService(m),
		ProjectVulnerabilitiesServiceMock:       mocks.NewMockProjectVulnerabilitiesService(m),
		ProjectsServiceMock:                     mocks.NewMockProjectsService(m),
		ProtectedBranchesServiceMock:            mocks.NewMockProtectedBranchesService(m),
		ProtectedEnvironmentsServiceMock:        mocks.NewMockProtectedEnvironmentsService(m),
		ProtectedTagsServiceMock:                mocks.NewMockProtectedTagsService(m),
		ReleaseLinksServiceMock:                 mocks.NewMockReleaseLinksService(m),
		ReleasesServiceMock:                     mocks.NewMockReleasesService(m),
		RepositoriesServiceMock:                 mocks.NewMockRepositoriesService(m),
		RepositoryFilesServiceMock:              mocks.NewMockRepositoryFilesService(m),
		RepositorySubmodulesServiceMock:         mocks.NewMockRepositorySubmodulesService(m),
		ResourceGroupServiceMock:                mocks.NewMockResourceGroupService(m),
		ResourceIterationEventsServiceMock:      mocks.NewMockResourceIterationEventsService(m),
		ResourceLabelEventsServiceMock:          mocks.NewMockResourceLabelEventsService(m),
		ResourceMilestoneEventsServiceMock:      mocks.NewMockResourceMilestoneEventsService(m),
		ResourceStateEventsServiceMock:          mocks.NewMockResourceStateEventsService(m),
		ResourceWeightEventsServiceMock:         mocks.NewMockResourceWeightEventsService(m),
		RunnersServiceMock:                      mocks.NewMockRunnersService(m),
		SearchServiceMock:                       mocks.NewMockSearchService(m),
		ServicesServiceMock:                     mocks.NewMockServicesService(m),
		SettingsServiceMock:                     mocks.NewMockSettingsService(m),
		SidekiqServiceMock:                      mocks.NewMockSidekiqService(m),
		SnippetRepositoryStorageMoveServiceMock: mocks.NewMockSnippetRepositoryStorageMoveService(m),
		SnippetsServiceMock:                     mocks.NewMockSnippetsService(m),
		SystemHooksServiceMock:                  mocks.NewMockSystemHooksService(m),
		TagsServiceMock:                         mocks.NewMockTagsService(m),
		TodosServiceMock:                        mocks.NewMockTodosService(m),
		TopicsServiceMock:                       mocks.NewMockTopicsService(m),
		UsersServiceMock:                        mocks.NewMockUsersService(m),
		ValidateServiceMock:                     mocks.NewMockValidateService(m),
		VersionServiceMock:                      mocks.NewMockVersionService(m),
		WikisServiceMock:                        mocks.NewMockWikisService(m),
	}
}

// AccessRequests returns a mocked [AccessRequestsService] service.
func (m *MockClient) AccessRequests() AccessRequestsService {
	return m.AccessRequestsServiceMock
}

// Appearance returns a mocked [AppearanceService] service.
func (m *MockClient) Appearance() AppearanceService {
	return m.AppearanceServiceMock
}

// Applications returns a mocked [ApplicationsService] service.
func (m *MockClient) Applications() ApplicationsService {
	return m.ApplicationsServiceMock
}

// AuditEvents returns a mocked [AuditEventsService] service.
func (m *MockClient) AuditEvents() AuditEventsService {
	return m.AuditEventsServiceMock
}

// Avatar returns a mocked [AvatarRequestsService] service.
func (m *MockClient) Avatar() AvatarRequestsService {
	return m.AvatarRequestsServiceMock
}

// AwardEmoji returns a mocked [AwardEmojiService] service.
func (m *MockClient) AwardEmoji() AwardEmojiService {
	return m.AwardEmojiServiceMock
}

// Branches returns a mocked [BranchesService] service.
func (m *MockClient) Branches() BranchesService {
	return m.BranchesServiceMock
}

// BroadcastMessage returns a mocked [BroadcastMessagesService] service.
func (m *MockClient) BroadcastMessage() BroadcastMessagesService {
	return m.BroadcastMessagesServiceMock
}

// CIYMLTemplate returns a mocked [CIYMLTemplatesService] service.
func (m *MockClient) CIYMLTemplate() CIYMLTemplatesService {
	return m.CIYMLTemplatesServiceMock
}

// ClusterAgents returns a mocked [ClusterAgentsService] service.
func (m *MockClient) ClusterAgents() ClusterAgentsService {
	return m.ClusterAgentsServiceMock
}

// Commits returns a mocked [CommitsService] service.
func (m *MockClient) Commits() CommitsService {
	return m.CommitsServiceMock
}

// ContainerRegistry returns a mocked [ContainerRegistryService] service.
func (m *MockClient) ContainerRegistry() ContainerRegistryService {
	return m.ContainerRegistryServiceMock
}

// CustomAttribute returns a mocked [CustomAttributesService] service.
func (m *MockClient) CustomAttribute() CustomAttributesService {
	return m.CustomAttributesServiceMock
}

// DORAMetrics returns a mocked [DORAMetricsService] service.
func (m *MockClient) DORAMetrics() DORAMetricsService {
	return m.DORAMetricsServiceMock
}

// DeployKeys returns a mocked [DeployKeysService] service.
func (m *MockClient) DeployKeys() DeployKeysService {
	return m.DeployKeysServiceMock
}

// DeployTokens returns a mocked [DeployTokensService] service.
func (m *MockClient) DeployTokens() DeployTokensService {
	return m.DeployTokensServiceMock
}

// DeploymentMergeRequests returns a mocked [DeploymentMergeRequestsService] service.
func (m *MockClient) DeploymentMergeRequests() DeploymentMergeRequestsService {
	return m.DeploymentMergeRequestsServiceMock
}

// Deployments returns a mocked [DeploymentsService] service.
func (m *MockClient) Deployments() DeploymentsService {
	return m.DeploymentsServiceMock
}

// Discussions returns a mocked [DiscussionsService] service.
func (m *MockClient) Discussions() DiscussionsService {
	return m.DiscussionsServiceMock
}

// DockerfileTemplate returns a mocked [DockerfileTemplatesService] service.
func (m *MockClient) DockerfileTemplate() DockerfileTemplatesService {
	return m.DockerfileTemplatesServiceMock
}

// DraftNotes returns a mocked [DraftNotesService] service.
func (m *MockClient) DraftNotes() DraftNotesService {
	return m.DraftNotesServiceMock
}

// Environments returns a mocked [EnvironmentsService] service.
func (m *MockClient) Environments() EnvironmentsService {
	return m.EnvironmentsServiceMock
}

// EpicIssues returns a mocked [EpicIssuesService] service.
func (m *MockClient) EpicIssues() EpicIssuesService {
	return m.EpicIssuesServiceMock
}

// Epics returns a mocked [EpicsService] service.
func (m *MockClient) Epics() EpicsService {
	return m.EpicsServiceMock
}

// ErrorTracking returns a mocked [ErrorTrackingService] service.
func (m *MockClient) ErrorTracking() ErrorTrackingService {
	return m.ErrorTrackingServiceMock
}

// Events returns a mocked [EventsService] service.
func (m *MockClient) Events() EventsService {
	return m.EventsServiceMock
}

// ExternalStatusChecks returns a mocked [ExternalStatusChecksService] service.
func (m *MockClient) ExternalStatusChecks() ExternalStatusChecksService {
	return m.ExternalStatusChecksServiceMock
}

// Features returns a mocked [FeaturesService] service.
func (m *MockClient) Features() FeaturesService {
	return m.FeaturesServiceMock
}

// FreezePeriods returns a mocked [FreezePeriodsService] service.
func (m *MockClient) FreezePeriods() FreezePeriodsService {
	return m.FreezePeriodsServiceMock
}

// GenericPackages returns a mocked [GenericPackagesService] service.
func (m *MockClient) GenericPackages() GenericPackagesService {
	return m.GenericPackagesServiceMock
}

// GeoNodes returns a mocked [GeoNodesService] service.
func (m *MockClient) GeoNodes() GeoNodesService {
	return m.GeoNodesServiceMock
}

// GitIgnoreTemplates returns a mocked [GitIgnoreTemplatesService] service.
func (m *MockClient) GitIgnoreTemplates() GitIgnoreTemplatesService {
	return m.GitIgnoreTemplatesServiceMock
}

// GroupAccessTokens returns a mocked [GroupAccessTokensService] service.
func (m *MockClient) GroupAccessTokens() GroupAccessTokensService {
	return m.GroupAccessTokensServiceMock
}

// GroupBadges returns a mocked [GroupBadgesService] service.
func (m *MockClient) GroupBadges() GroupBadgesService {
	return m.GroupBadgesServiceMock
}

// GroupCluster returns a mocked [GroupClustersService] service.
func (m *MockClient) GroupCluster() GroupClustersService {
	return m.GroupClustersServiceMock
}

// GroupEpicBoards returns a mocked [GroupEpicBoardsService] service.
func (m *MockClient) GroupEpicBoards() GroupEpicBoardsService {
	return m.GroupEpicBoardsServiceMock
}

// GroupImportExport returns a mocked [GroupImportExportService] service.
func (m *MockClient) GroupImportExport() GroupImportExportService {
	return m.GroupImportExportServiceMock
}

// GroupIssueBoards returns a mocked [GroupIssueBoardsService] service.
func (m *MockClient) GroupIssueBoards() GroupIssueBoardsService {
	return m.GroupIssueBoardsServiceMock
}

// GroupIterations returns a mocked [GroupIterationsService] service.
func (m *MockClient) GroupIterations() GroupIterationsService {
	return m.GroupIterationsServiceMock
}

// GroupLabels returns a mocked [GroupLabelsService] service.
func (m *MockClient) GroupLabels() GroupLabelsService {
	return m.GroupLabelsServiceMock
}

// GroupMembers returns a mocked [GroupMembersService] service.
func (m *MockClient) GroupMembers() GroupMembersService {
	return m.GroupMembersServiceMock
}

// GroupMilestones returns a mocked [GroupMilestonesService] service.
func (m *MockClient) GroupMilestones() GroupMilestonesService {
	return m.GroupMilestonesServiceMock
}

// GroupProtectedEnvironments returns a mocked [GroupProtectedEnvironmentsService] service.
func (m *MockClient) GroupProtectedEnvironments() GroupProtectedEnvironmentsService {
	return m.GroupProtectedEnvironmentsServiceMock
}

// GroupRepositoryStorageMove returns a mocked [GroupRepositoryStorageMoveService] service.
func (m *MockClient) GroupRepositoryStorageMove() GroupRepositoryStorageMoveService {
	return m.GroupRepositoryStorageMoveServiceMock
}

// GroupSSHCertificates returns a mocked [GroupSSHCertificatesService] service.
func (m *MockClient) GroupSSHCertificates() GroupSSHCertificatesService {
	return m.GroupSSHCertificatesServiceMock
}

// GroupVariables returns a mocked [GroupVariablesService] service.
func (m *MockClient) GroupVariables() GroupVariablesService {
	return m.GroupVariablesServiceMock
}

// GroupWikis returns a mocked [GroupWikisService] service.
func (m *MockClient) GroupWikis() GroupWikisService {
	return m.GroupWikisServiceMock
}

// Groups returns a mocked [GroupsService] service.
func (m *MockClient) Groups() GroupsService {
	return m.GroupsServiceMock
}

// Import returns a mocked [ImportService] service.
func (m *MockClient) Import() ImportService {
	return m.ImportServiceMock
}

// InstanceCluster returns a mocked [InstanceClustersService] service.
func (m *MockClient) InstanceCluster() InstanceClustersService {
	return m.InstanceClustersServiceMock
}

// InstanceVariables returns a mocked [InstanceVariablesService] service.
func (m *MockClient) InstanceVariables() InstanceVariablesService {
	return m.InstanceVariablesServiceMock
}

// Invites returns a mocked [InvitesService] service.
func (m *MockClient) Invites() InvitesService {
	return m.InvitesServiceMock
}

// Boards returns a mocked [IssueBoardsService] service.
func (m *MockClient) Boards() IssueBoardsService {
	return m.IssueBoardsServiceMock
}

// IssueLinks returns a mocked [IssueLinksService] service.
func (m *MockClient) IssueLinks() IssueLinksService {
	return m.IssueLinksServiceMock
}

// Issues returns a mocked [IssuesService] service.
func (m *MockClient) Issues() IssuesService {
	return m.IssuesServiceMock
}

// IssuesStatistics returns a mocked [IssuesStatisticsService] service.
func (m *MockClient) IssuesStatistics() IssuesStatisticsService {
	return m.IssuesStatisticsServiceMock
}

// JobTokenScope returns a mocked [JobTokenScopeService] service.
func (m *MockClient) JobTokenScope() JobTokenScopeService {
	return m.JobTokenScopeServiceMock
}

// Jobs returns a mocked [JobsService] service.
func (m *MockClient) Jobs() JobsService {
	return m.JobsServiceMock
}

// Keys returns a mocked [KeysService] service.
func (m *MockClient) Keys() KeysService {
	return m.KeysServiceMock
}

// Labels returns a mocked [LabelsService] service.
func (m *MockClient) Labels() LabelsService {
	return m.LabelsServiceMock
}

// License returns a mocked [LicenseService] service.
func (m *MockClient) License() LicenseService {
	return m.LicenseServiceMock
}

// LicenseTemplates returns a mocked [LicenseTemplatesService] service.
func (m *MockClient) LicenseTemplates() LicenseTemplatesService {
	return m.LicenseTemplatesServiceMock
}

// ManagedLicenses returns a mocked [ManagedLicensesService] service.
func (m *MockClient) ManagedLicenses() ManagedLicensesService {
	return m.ManagedLicensesServiceMock
}

// Markdown returns a mocked [MarkdownService] service.
func (m *MockClient) Markdown() MarkdownService {
	return m.MarkdownServiceMock
}

// MemberRolesService returns a mocked [MemberRolesService] service.
func (m *MockClient) MemberRolesService() MemberRolesService {
	return m.MemberRolesServiceMock
}

// MergeRequestApprovals returns a mocked [MergeRequestApprovalsService] service.
func (m *MockClient) MergeRequestApprovals() MergeRequestApprovalsService {
	return m.MergeRequestApprovalsServiceMock
}

// MergeRequests returns a mocked [MergeRequestsService] service.
func (m *MockClient) MergeRequests() MergeRequestsService {
	return m.MergeRequestsServiceMock
}

// MergeTrains returns a mocked [MergeTrainsService] service.
func (m *MockClient) MergeTrains() MergeTrainsService {
	return m.MergeTrainsServiceMock
}

// Metadata returns a mocked [MetadataService] service.
func (m *MockClient) Metadata() MetadataService {
	return m.MetadataServiceMock
}

// Milestones returns a mocked [MilestonesService] service.
func (m *MockClient) Milestones() MilestonesService {
	return m.MilestonesServiceMock
}

// Namespaces returns a mocked [NamespacesService] service.
func (m *MockClient) Namespaces() NamespacesService {
	return m.NamespacesServiceMock
}

// Notes returns a mocked [NotesService] service.
func (m *MockClient) Notes() NotesService {
	return m.NotesServiceMock
}

// NotificationSettings returns a mocked [NotificationSettingsService] service.
func (m *MockClient) NotificationSettings() NotificationSettingsService {
	return m.NotificationSettingsServiceMock
}

// Packages returns a mocked [PackagesService] service.
func (m *MockClient) Packages() PackagesService {
	return m.PackagesServiceMock
}

// PagesDomains returns a mocked [PagesDomainsService] service.
func (m *MockClient) PagesDomains() PagesDomainsService {
	return m.PagesDomainsServiceMock
}

// Pages returns a mocked [PagesService] service.
func (m *MockClient) Pages() PagesService {
	return m.PagesServiceMock
}

// PersonalAccessTokens returns a mocked [PersonalAccessTokensService] service.
func (m *MockClient) PersonalAccessTokens() PersonalAccessTokensService {
	return m.PersonalAccessTokensServiceMock
}

// PipelineSchedules returns a mocked [PipelineSchedulesService] service.
func (m *MockClient) PipelineSchedules() PipelineSchedulesService {
	return m.PipelineSchedulesServiceMock
}

// PipelineTriggers returns a mocked [PipelineTriggersService] service.
func (m *MockClient) PipelineTriggers() PipelineTriggersService {
	return m.PipelineTriggersServiceMock
}

// Pipelines returns a mocked [PipelinesService] service.
func (m *MockClient) Pipelines() PipelinesService {
	return m.PipelinesServiceMock
}

// PlanLimits returns a mocked [PlanLimitsService] service.
func (m *MockClient) PlanLimits() PlanLimitsService {
	return m.PlanLimitsServiceMock
}

// ProjectAccessTokens returns a mocked [ProjectAccessTokensService] service.
func (m *MockClient) ProjectAccessTokens() ProjectAccessTokensService {
	return m.ProjectAccessTokensServiceMock
}

// ProjectBadges returns a mocked [ProjectBadgesService] service.
func (m *MockClient) ProjectBadges() ProjectBadgesService {
	return m.ProjectBadgesServiceMock
}

// ProjectCluster returns a mocked [ProjectClustersService] service.
func (m *MockClient) ProjectCluster() ProjectClustersService {
	return m.ProjectClustersServiceMock
}

// ProjectFeatureFlags returns a mocked [ProjectFeatureFlagService] service.
func (m *MockClient) ProjectFeatureFlags() ProjectFeatureFlagService {
	return m.ProjectFeatureFlagServiceMock
}

// ProjectImportExport returns a mocked [ProjectImportExportService] service.
func (m *MockClient) ProjectImportExport() ProjectImportExportService {
	return m.ProjectImportExportServiceMock
}

// ProjectIterations returns a mocked [ProjectIterationsService] service.
func (m *MockClient) ProjectIterations() ProjectIterationsService {
	return m.ProjectIterationsServiceMock
}

// ProjectMembers returns a mocked [ProjectMembersService] service.
func (m *MockClient) ProjectMembers() ProjectMembersService {
	return m.ProjectMembersServiceMock
}

// ProjectMirrors returns a mocked [ProjectMirrorService] service.
func (m *MockClient) ProjectMirrors() ProjectMirrorService {
	return m.ProjectMirrorServiceMock
}

// ProjectRepositoryStorageMove returns a mocked [ProjectRepositoryStorageMoveService] service.
func (m *MockClient) ProjectRepositoryStorageMove() ProjectRepositoryStorageMoveService {
	return m.ProjectRepositoryStorageMoveServiceMock
}

// ProjectSnippets returns a mocked [ProjectSnippetsService] service.
func (m *MockClient) ProjectSnippets() ProjectSnippetsService {
	return m.ProjectSnippetsServiceMock
}

// ProjectTemplates returns a mocked [ProjectTemplatesService] service.
func (m *MockClient) ProjectTemplates() ProjectTemplatesService {
	return m.ProjectTemplatesServiceMock
}

// ProjectVariables returns a mocked [ProjectVariablesService] service.
func (m *MockClient) ProjectVariables() ProjectVariablesService {
	return m.ProjectVariablesServiceMock
}

// ProjectVulnerabilities returns a mocked [ProjectVulnerabilitiesService] service.
func (m *MockClient) ProjectVulnerabilities() ProjectVulnerabilitiesService {
	return m.ProjectVulnerabilitiesServiceMock
}

// Projects returns a mocked [ProjectsService] service.
func (m *MockClient) Projects() ProjectsService {
	return m.ProjectsServiceMock
}

// ProtectedBranches returns a mocked [ProtectedBranchesService] service.
func (m *MockClient) ProtectedBranches() ProtectedBranchesService {
	return m.ProtectedBranchesServiceMock
}

// ProtectedEnvironments returns a mocked [ProtectedEnvironmentsService] service.
func (m *MockClient) ProtectedEnvironments() ProtectedEnvironmentsService {
	return m.ProtectedEnvironmentsServiceMock
}

// ProtectedTags returns a mocked [ProtectedTagsService] service.
func (m *MockClient) ProtectedTags() ProtectedTagsService {
	return m.ProtectedTagsServiceMock
}

// ReleaseLinks returns a mocked [ReleaseLinksService] service.
func (m *MockClient) ReleaseLinks() ReleaseLinksService {
	return m.ReleaseLinksServiceMock
}

// Releases returns a mocked [ReleasesService] service.
func (m *MockClient) Releases() ReleasesService {
	return m.ReleasesServiceMock
}

// Repositories returns a mocked [RepositoriesService] service.
func (m *MockClient) Repositories() RepositoriesService {
	return m.RepositoriesServiceMock
}

// RepositoryFiles returns a mocked [RepositoryFilesService] service.
func (m *MockClient) RepositoryFiles() RepositoryFilesService {
	return m.RepositoryFilesServiceMock
}

// RepositorySubmodules returns a mocked [RepositorySubmodulesService] service.
func (m *MockClient) RepositorySubmodules() RepositorySubmodulesService {
	return m.RepositorySubmodulesServiceMock
}

// ResourceGroup returns a mocked [ResourceGroupService] service.
func (m *MockClient) ResourceGroup() ResourceGroupService {
	return m.ResourceGroupServiceMock
}

// ResourceIterationEvents returns a mocked [ResourceIterationEventsService] service.
func (m *MockClient) ResourceIterationEvents() ResourceIterationEventsService {
	return m.ResourceIterationEventsServiceMock
}

// ResourceLabelEvents returns a mocked [ResourceLabelEventsService] service.
func (m *MockClient) ResourceLabelEvents() ResourceLabelEventsService {
	return m.ResourceLabelEventsServiceMock
}

// ResourceMilestoneEvents returns a mocked [ResourceMilestoneEventsService] service.
func (m *MockClient) ResourceMilestoneEvents() ResourceMilestoneEventsService {
	return m.ResourceMilestoneEventsServiceMock
}

// ResourceStateEvents returns a mocked [ResourceStateEventsService] service.
func (m *MockClient) ResourceStateEvents() ResourceStateEventsService {
	return m.ResourceStateEventsServiceMock
}

// ResourceWeightEvents returns a mocked [ResourceWeightEventsService] service.
func (m *MockClient) ResourceWeightEvents() ResourceWeightEventsService {
	return m.ResourceWeightEventsServiceMock
}

// Runners returns a mocked [RunnersService] service.
func (m *MockClient) Runners() RunnersService {
	return m.RunnersServiceMock
}

// Search returns a mocked [SearchService] service.
func (m *MockClient) Search() SearchService {
	return m.SearchServiceMock
}

// Services returns a mocked [ServicesService] service.
func (m *MockClient) Services() ServicesService {
	return m.ServicesServiceMock
}

// Settings returns a mocked [SettingsService] service.
func (m *MockClient) Settings() SettingsService {
	return m.SettingsServiceMock
}

// Sidekiq returns a mocked [SidekiqService] service.
func (m *MockClient) Sidekiq() SidekiqService {
	return m.SidekiqServiceMock
}

// SnippetRepositoryStorageMove returns a mocked [SnippetRepositoryStorageMoveService] service.
func (m *MockClient) SnippetRepositoryStorageMove() SnippetRepositoryStorageMoveService {
	return m.SnippetRepositoryStorageMoveServiceMock
}

// Snippets returns a mocked [SnippetsService] service.
func (m *MockClient) Snippets() SnippetsService {
	return m.SnippetsServiceMock
}

// SystemHooks returns a mocked [SystemHooksService] service.
func (m *MockClient) SystemHooks() SystemHooksService {
	return m.SystemHooksServiceMock
}

// Tags returns a mocked [TagsService] service.
func (m *MockClient) Tags() TagsService {
	return m.TagsServiceMock
}

// Todos returns a mocked [TodosService] service.
func (m *MockClient) Todos() TodosService {
	return m.TodosServiceMock
}

// Topics returns a mocked [TopicsService] service.
func (m *MockClient) Topics() TopicsService {
	return m.TopicsServiceMock
}

// Users returns a mocked [UsersService] service.
func (m *MockClient) Users() UsersService {
	return m.UsersServiceMock
}

// Validate returns a mocked [ValidateService] service.
func (m *MockClient) Validate() ValidateService {
	return m.ValidateServiceMock
}

// Version returns a mocked [VersionService] service.
func (m *MockClient) Version() VersionService {
	return m.VersionServiceMock
}

// Wikis returns a mocked [WikisService] service.
func (m *MockClient) Wikis() WikisService {
	return m.WikisServiceMock
}
