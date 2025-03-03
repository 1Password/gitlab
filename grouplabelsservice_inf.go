// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// GroupLabelsService is an interface for [gitlab.Client.GroupLabels]
type GroupLabelsService interface {
	// ListGroupLabels gets all labels for given group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_labels.html#list-group-labels
	ListGroupLabels(gid interface{}, opt *ListGroupLabelsOptions, options ...RequestOptionFunc) ([]*GroupLabel, *Response, error)
	// GetGroupLabel get a single label for a given group.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_labels.html#get-a-single-group-label
	GetGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*GroupLabel, *Response, error)
	// CreateGroupLabel creates a new label for given group with given name and
	// color.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_labels.html#create-a-new-group-label
	CreateGroupLabel(gid interface{}, opt *CreateGroupLabelOptions, options ...RequestOptionFunc) (*GroupLabel, *Response, error)
	// DeleteGroupLabel deletes a group label given by its name or ID.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_labels.html#delete-a-group-label
	DeleteGroupLabel(gid interface{}, lid interface{}, opt *DeleteGroupLabelOptions, options ...RequestOptionFunc) (*Response, error)
	// UpdateGroupLabel updates an existing label with new name or now color. At least
	// one parameter is required, to update the label.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_labels.html#update-a-group-label
	UpdateGroupLabel(gid interface{}, lid interface{}, opt *UpdateGroupLabelOptions, options ...RequestOptionFunc) (*GroupLabel, *Response, error)
	// SubscribeToGroupLabel subscribes the authenticated user to a label to receive
	// notifications. If the user is already subscribed to the label, the status
	// code 304 is returned.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_labels.html#subscribe-to-a-group-label
	SubscribeToGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*GroupLabel, *Response, error)
	// UnsubscribeFromGroupLabel unsubscribes the authenticated user from a label to not
	// receive notifications from it. If the user is not subscribed to the label, the
	// status code 304 is returned.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/ee/api/group_labels.html#unsubscribe-from-a-group-label
	UnsubscribeFromGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*Response, error)
}
