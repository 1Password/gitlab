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

// Description: Contains helper functions.

package gitlab

import _gitlab "github.com/xanzy/go-gitlab"

// Ptr returns a pointer to the value passed in.
func Ptr[T any](v T) *T {
	return &v
}

// WithBaseURL is an alias to [_gitlab.WithBaseURL]
var WithBaseURL = _gitlab.WithBaseURL
