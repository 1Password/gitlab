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

// Description: Contains pulled out definitions of some of the embedded
// structs to make it easier to construct them in tests.

package gitlab

import (
	"time"
)

// These are compile time tests that the embedded structs below are
// correct.
var (
	_ MergeEvent = MergeEvent{
		ObjectAttributes: MergeEventObjectAttributes{},
		Changes:          MergeEventChanges{},
	}
)

// MergeEventObjectAttributes is a duplicate of the embedded struct at
// [MergeEvent.ObjectAttributes].
type MergeEventObjectAttributes struct {
	ID                       int          "json:\"id\""
	TargetBranch             string       "json:\"target_branch\""
	SourceBranch             string       "json:\"source_branch\""
	SourceProjectID          int          "json:\"source_project_id\""
	AuthorID                 int          "json:\"author_id\""
	AssigneeID               int          "json:\"assignee_id\""
	AssigneeIDs              []int        "json:\"assignee_ids\""
	ReviewerIDs              []int        "json:\"reviewer_ids\""
	Title                    string       "json:\"title\""
	CreatedAt                string       "json:\"created_at\""
	UpdatedAt                string       "json:\"updated_at\""
	StCommits                []*Commit    "json:\"st_commits\""
	StDiffs                  []*Diff      "json:\"st_diffs\""
	LastEditedAt             string       "json:\"last_edited_at\""
	LastEditedByID           int          "json:\"last_edited_by_id\""
	MilestoneID              int          "json:\"milestone_id\""
	StateID                  StateID      "json:\"state_id\""
	State                    string       "json:\"state\""
	MergeStatus              string       "json:\"merge_status\""
	TargetProjectID          int          "json:\"target_project_id\""
	IID                      int          "json:\"iid\""
	Description              string       "json:\"description\""
	Position                 int          "json:\"position\""
	LockedAt                 string       "json:\"locked_at\""
	UpdatedByID              int          "json:\"updated_by_id\""
	MergeError               string       "json:\"merge_error\""
	MergeParams              *MergeParams "json:\"merge_params\""
	MergeWhenBuildSucceeds   bool         "json:\"merge_when_build_succeeds\""
	MergeUserID              int          "json:\"merge_user_id\""
	MergeCommitSHA           string       "json:\"merge_commit_sha\""
	DeletedAt                string       "json:\"deleted_at\""
	ApprovalsBeforeMerge     string       "json:\"approvals_before_merge\""
	RebaseCommitSHA          string       "json:\"rebase_commit_sha\""
	InProgressMergeCommitSHA string       "json:\"in_progress_merge_commit_sha\""
	LockVersion              int          "json:\"lock_version\""
	TimeEstimate             int          "json:\"time_estimate\""
	Source                   *Repository  "json:\"source\""
	Target                   *Repository  "json:\"target\""
	HeadPipelineID           *int         "json:\"head_pipeline_id\""
	LastCommit               struct {
		ID        string     "json:\"id\""
		Message   string     "json:\"message\""
		Title     string     "json:\"title\""
		Timestamp *time.Time "json:\"timestamp\""
		URL       string     "json:\"url\""
		Author    struct {
			Name  string "json:\"name\""
			Email string "json:\"email\""
		} "json:\"author\""
	} "json:\"last_commit\""
	BlockingDiscussionsResolved bool          "json:\"blocking_discussions_resolved\""
	WorkInProgress              bool          "json:\"work_in_progress\""
	Draft                       bool          "json:\"draft\""
	TotalTimeSpent              int           "json:\"total_time_spent\""
	TimeChange                  int           "json:\"time_change\""
	HumanTotalTimeSpent         string        "json:\"human_total_time_spent\""
	HumanTimeChange             string        "json:\"human_time_change\""
	HumanTimeEstimate           string        "json:\"human_time_estimate\""
	FirstContribution           bool          "json:\"first_contribution\""
	URL                         string        "json:\"url\""
	Labels                      []*EventLabel "json:\"labels\""
	Action                      string        "json:\"action\""
	DetailedMergeStatus         string        "json:\"detailed_merge_status\""
	OldRev                      string        "json:\"oldrev\""
}

// MergeEventProject is a duplicate of the embedded struct at
// [MergeEvent.Project].
type MergeEventProject struct {
	ID                int             "json:\"id\""
	Name              string          "json:\"name\""
	Description       string          "json:\"description\""
	AvatarURL         string          "json:\"avatar_url\""
	GitSSHURL         string          "json:\"git_ssh_url\""
	GitHTTPURL        string          "json:\"git_http_url\""
	Namespace         string          "json:\"namespace\""
	PathWithNamespace string          "json:\"path_with_namespace\""
	DefaultBranch     string          "json:\"default_branch\""
	CIConfigPath      string          "json:\"ci_config_path\""
	Homepage          string          "json:\"homepage\""
	URL               string          "json:\"url\""
	SSHURL            string          "json:\"ssh_url\""
	HTTPURL           string          "json:\"http_url\""
	WebURL            string          "json:\"web_url\""
	Visibility        VisibilityValue "json:\"visibility\""
}

// MergeEventChangesString is a duplicate of multiple embedded structs
// that were under [MergeEvent.Changes]. This is for string types.
type MergeEventChangesString struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
}

// MergeEventChangesBool is a duplicate of multiple embedded structs
// that were under [MergeEvent.Changes]. This is for bool types.
type MergeEventChangesBool struct {
	Previous bool `json:"previous"`
	Current  bool `json:"current"`
}

// MergeEventChangesInt is a duplicate of multiple embedded structs
// that were under [MergeEvent.Changes]. This is for int types.
type MergeEventChangesInt struct {
	Previous int `json:"previous"`
	Current  int `json:"current"`
}

// MergeEventChanges is a duplicate of the embedded struct at
// [MergeEvent.Changes].
type MergeEventChanges struct {
	Assignees struct {
		Previous []*EventUser `json:"previous"`
		Current  []*EventUser `json:"current"`
	} `json:"assignees"`
	Reviewers struct {
		Previous []*EventUser `json:"previous"`
		Current  []*EventUser `json:"current"`
	} `json:"reviewers"`
	Description struct {
		Previous string `json:"previous"`
		Current  string `json:"current"`
	} `json:"description"`
	Draft struct {
		Previous bool `json:"previous"`
		Current  bool `json:"current"`
	} `json:"draft"`
	Labels struct {
		Previous []*EventLabel `json:"previous"`
		Current  []*EventLabel `json:"current"`
	} `json:"labels"`
	LastEditedAt struct {
		Previous string `json:"previous"`
		Current  string `json:"current"`
	} `json:"last_edited_at"`
	LastEditedByID struct {
		Previous int `json:"previous"`
		Current  int `json:"current"`
	} `json:"last_edited_by_id"`
	MergeStatus struct {
		Previous string `json:"previous"`
		Current  string `json:"current"`
	} `json:"merge_status"`
	MilestoneID struct {
		Previous int `json:"previous"`
		Current  int `json:"current"`
	} `json:"milestone_id"`
	SourceBranch struct {
		Previous string `json:"previous"`
		Current  string `json:"current"`
	} `json:"source_branch"`
	SourceProjectID struct {
		Previous int `json:"previous"`
		Current  int `json:"current"`
	} `json:"source_project_id"`
	StateID struct {
		Previous StateID `json:"previous"`
		Current  StateID `json:"current"`
	} `json:"state_id"`
	TargetBranch struct {
		Previous string `json:"previous"`
		Current  string `json:"current"`
	} `json:"target_branch"`
	TargetProjectID struct {
		Previous int `json:"previous"`
		Current  int `json:"current"`
	} `json:"target_project_id"`
	Title struct {
		Previous string `json:"previous"`
		Current  string `json:"current"`
	} `json:"title"`
	UpdatedAt struct {
		Previous string `json:"previous"`
		Current  string `json:"current"`
	} `json:"updated_at"`
	UpdatedByID struct {
		Previous int `json:"previous"`
		Current  int `json:"current"`
	} `json:"updated_by_id"`
}
