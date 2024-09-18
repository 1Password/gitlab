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
 "github.com/xanzy/go-gitlab"
)

// client is a wrapper around [gitlab.Client] that implements the
// [Client] interface by turning struct fields into accessors.
type client struct {
 *gitlab.Client

 {{- range .Services }}
 {{ .PrivateName }} {{ .Name }}
 {{- end }}
}

{{- range .Services }}
// {{ .Accessor }} returns the [{{ .Name }}] service for the client.
func (c *client) {{ .Accessor }}() {{ .Name }} {
 return c.{{ .PrivateName }}
}
{{- end }}

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
		Client:            gl,
		{{- range .Services }}
		{{ .PrivateName }}: gl.{{ .Accessor }},
		{{- end }}
	}
}
