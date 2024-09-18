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

	"github.com/jaredallard/gitlab"
	"github.com/jaredallard/gitlab/mocks"
	"go.uber.org/mock/gomock"
)

// _ ensures that [MockClient] implements [Client].
var _ Client = &MockClient{}

// MockClient is a mock of Client interface. Create with
// [NewMockClient].
type MockClient struct {
{{- range .Services }}
	{{ .Name }}Mock *mocks.Mock{{ .Name }}
{{- end }}
}

// NewMockClient creates a new mock client for testing. This should be
// used in place of [NewClient] when testing.
func NewMockClient(t *testing.T) *MockClient {
	m := gomock.NewController(t)
	return &MockClient{
		{{- range .Services }}
		{{ .Name }}Mock:                mocks.NewMock{{ .Name }}(m),
		{{- end }}
	}
}

{{- range .Services }}
// {{ .Accessor }} returns a mocked [{{ .Name }}] service.
func (m *MockClient) {{ .Accessor }}() {{ .Name }} {
	return m.{{ .Name }}Mock
}
{{- end }}
