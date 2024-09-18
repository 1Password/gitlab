// Code generated by ifacemaker; DO NOT EDIT.

package gitlab

// AppearanceService is an interface for [gitlab.Client.Appearance]
type AppearanceService interface {
	// GetAppearance gets the current appearance configuration of the GitLab instance.
	//
	// Gitlab API docs:
	// https://docs.gitlab.com/ee/api/appearance.html#get-current-appearance-configuration
	GetAppearance(options ...RequestOptionFunc) (*Appearance, *Response, error)
	// ChangeAppearance changes the appearance configuration.
	//
	// Gitlab API docs:
	// https://docs.gitlab.com/ee/api/appearance.html#change-appearance-configuration
	ChangeAppearance(opt *ChangeAppearanceOptions, options ...RequestOptionFunc) (*Appearance, *Response, error)
}
