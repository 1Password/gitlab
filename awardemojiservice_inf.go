// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// AwardEmojiService is an interface for [gitlab.Client.AwardEmoji]
type AwardEmojiService interface {
	// ListMergeRequestAwardEmoji gets a list of all award emoji on the merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#list-an-awardables-award-emojis
	ListMergeRequestAwardEmoji(pid interface{}, mergeRequestIID int, opt *ListAwardEmojiOptions, options ...RequestOptionFunc) ([]*AwardEmoji, *Response, error)
	// ListIssueAwardEmoji gets a list of all award emoji on the issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#list-an-awardables-award-emojis
	ListIssueAwardEmoji(pid interface{}, issueIID int, opt *ListAwardEmojiOptions, options ...RequestOptionFunc) ([]*AwardEmoji, *Response, error)
	// ListSnippetAwardEmoji gets a list of all award emoji on the snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#list-an-awardables-award-emojis
	ListSnippetAwardEmoji(pid interface{}, snippetID int, opt *ListAwardEmojiOptions, options ...RequestOptionFunc) ([]*AwardEmoji, *Response, error)
	// GetMergeRequestAwardEmoji get an award emoji from merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#get-single-award-emoji
	GetMergeRequestAwardEmoji(pid interface{}, mergeRequestIID, awardID int, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// GetIssueAwardEmoji get an award emoji from issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#get-single-award-emoji
	GetIssueAwardEmoji(pid interface{}, issueIID, awardID int, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// GetSnippetAwardEmoji get an award emoji from snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#get-single-award-emoji
	GetSnippetAwardEmoji(pid interface{}, snippetID, awardID int, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// CreateMergeRequestAwardEmoji get an award emoji from merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#award-a-new-emoji
	CreateMergeRequestAwardEmoji(pid interface{}, mergeRequestIID int, opt *CreateAwardEmojiOptions, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// CreateIssueAwardEmoji get an award emoji from issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#award-a-new-emoji
	CreateIssueAwardEmoji(pid interface{}, issueIID int, opt *CreateAwardEmojiOptions, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// CreateSnippetAwardEmoji get an award emoji from snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#award-a-new-emoji
	CreateSnippetAwardEmoji(pid interface{}, snippetID int, opt *CreateAwardEmojiOptions, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// DeleteIssueAwardEmoji delete award emoji on an issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#delete-an-award-emoji
	DeleteIssueAwardEmoji(pid interface{}, issueIID, awardID int, options ...RequestOptionFunc) (*Response, error)
	// DeleteMergeRequestAwardEmoji delete award emoji on a merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#delete-an-award-emoji
	DeleteMergeRequestAwardEmoji(pid interface{}, mergeRequestIID, awardID int, options ...RequestOptionFunc) (*Response, error)
	// DeleteSnippetAwardEmoji delete award emoji on a snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#delete-an-award-emoji
	DeleteSnippetAwardEmoji(pid interface{}, snippetID, awardID int, options ...RequestOptionFunc) (*Response, error)
	// ListIssuesAwardEmojiOnNote gets a list of all award emoji on a note from the
	// issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#list-a-comments-award-emojis
	ListIssuesAwardEmojiOnNote(pid interface{}, issueID, noteID int, opt *ListAwardEmojiOptions, options ...RequestOptionFunc) ([]*AwardEmoji, *Response, error)
	// ListMergeRequestAwardEmojiOnNote gets a list of all award emoji on a note
	// from the merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#list-a-comments-award-emojis
	ListMergeRequestAwardEmojiOnNote(pid interface{}, mergeRequestIID, noteID int, opt *ListAwardEmojiOptions, options ...RequestOptionFunc) ([]*AwardEmoji, *Response, error)
	// ListSnippetAwardEmojiOnNote gets a list of all award emoji on a note from the
	// snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#list-a-comments-award-emojis
	ListSnippetAwardEmojiOnNote(pid interface{}, snippetIID, noteID int, opt *ListAwardEmojiOptions, options ...RequestOptionFunc) ([]*AwardEmoji, *Response, error)
	// GetIssuesAwardEmojiOnNote gets an award emoji on a note from an issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#get-an-award-emoji-for-a-comment
	GetIssuesAwardEmojiOnNote(pid interface{}, issueID, noteID, awardID int, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// GetMergeRequestAwardEmojiOnNote gets an award emoji on a note from a
	// merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#get-an-award-emoji-for-a-comment
	GetMergeRequestAwardEmojiOnNote(pid interface{}, mergeRequestIID, noteID, awardID int, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// GetSnippetAwardEmojiOnNote gets an award emoji on a note from a snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#get-an-award-emoji-for-a-comment
	GetSnippetAwardEmojiOnNote(pid interface{}, snippetIID, noteID, awardID int, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// CreateIssuesAwardEmojiOnNote gets an award emoji on a note from an issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#award-a-new-emoji-on-a-comment
	CreateIssuesAwardEmojiOnNote(pid interface{}, issueID, noteID int, opt *CreateAwardEmojiOptions, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// CreateMergeRequestAwardEmojiOnNote gets an award emoji on a note from a
	// merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#award-a-new-emoji-on-a-comment
	CreateMergeRequestAwardEmojiOnNote(pid interface{}, mergeRequestIID, noteID int, opt *CreateAwardEmojiOptions, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// CreateSnippetAwardEmojiOnNote gets an award emoji on a note from a snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#award-a-new-emoji-on-a-comment
	CreateSnippetAwardEmojiOnNote(pid interface{}, snippetIID, noteID int, opt *CreateAwardEmojiOptions, options ...RequestOptionFunc) (*AwardEmoji, *Response, error)
	// DeleteIssuesAwardEmojiOnNote deletes an award emoji on a note from an issue.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#delete-an-award-emoji-from-a-comment
	DeleteIssuesAwardEmojiOnNote(pid interface{}, issueID, noteID, awardID int, options ...RequestOptionFunc) (*Response, error)
	// DeleteMergeRequestAwardEmojiOnNote deletes an award emoji on a note from a
	// merge request.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#delete-an-award-emoji-from-a-comment
	DeleteMergeRequestAwardEmojiOnNote(pid interface{}, mergeRequestIID, noteID, awardID int, options ...RequestOptionFunc) (*Response, error)
	// DeleteSnippetAwardEmojiOnNote deletes an award emoji on a note from a snippet.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/award_emoji.html#delete-an-award-emoji-from-a-comment
	DeleteSnippetAwardEmojiOnNote(pid interface{}, snippetIID, noteID, awardID int, options ...RequestOptionFunc) (*Response, error)
}
