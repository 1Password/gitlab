// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// ProtectedTagsService is an interface for [gitlab.Client.ProtectedTags]
type ProtectedTagsService interface {
	// ListProtectedTags returns a list of protected tags from a project.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/protected_tags.html#list-protected-tags
	ListProtectedTags(pid interface{}, opt *ListProtectedTagsOptions, options ...RequestOptionFunc) ([]*ProtectedTag, *Response, error)
	// GetProtectedTag returns a single protected tag or wildcard protected tag.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/protected_tags.html#get-a-single-protected-tag-or-wildcard-protected-tag
	GetProtectedTag(pid interface{}, tag string, options ...RequestOptionFunc) (*ProtectedTag, *Response, error)
	// ProtectRepositoryTags protects a single repository tag or several project
	// repository tags using a wildcard protected tag.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/protected_tags.html#protect-repository-tags
	ProtectRepositoryTags(pid interface{}, opt *ProtectRepositoryTagsOptions, options ...RequestOptionFunc) (*ProtectedTag, *Response, error)
	// UnprotectRepositoryTags unprotects the given protected tag or wildcard
	// protected tag.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/protected_tags.html#unprotect-repository-tags
	UnprotectRepositoryTags(pid interface{}, tag string, options ...RequestOptionFunc) (*Response, error)
}
