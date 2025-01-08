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

package main

import (
	_ "embed" // Used by go:embed
	"os"
	"text/template"

	"go.rgst.io/stencil/v2/pkg/slogext"
)

var (
	//go:embed embed/mock_client.go.tpl
	mockTplString string
	mockTpl       = template.Must(template.New("mock_client.go").Parse(mockTplString))

	//go:embed embed/interface.go.tpl
	interfaceTplString string
	interfaceTpl       = template.Must(template.New("interface.go").Parse(interfaceTplString))

	//go:embed embed/client.go.tpl
	clientTplString string
	clientTpl       = template.Must(template.New("client.go").Parse(clientTplString))
)

type TemplateInput struct {
	// Services is a list of services to generate clients for.
	Services []*Service
}

// generateClients generates the following clients; a mock client, a
// real client, and an interface based on the provided services.
func generateClients(log slogext.Logger, services []*Service) error {
	input := TemplateInput{Services: services}

	log.Info("Generating clients")

	if err := writeTemplateToFile(mockTpl, "mock_client.go", input); err != nil {
		return err
	}

	if err := writeTemplateToFile(interfaceTpl, "interface.go", input); err != nil {
		return err
	}

	if err := writeTemplateToFile(clientTpl, "client.go", input); err != nil {
		return err
	}

	return nil
}

// writeTemplateToFile executes the provided template with the provided
// data and writes the output to the provided path.
func writeTemplateToFile(tpl *template.Template, path string, data interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tpl.Execute(f, data); err != nil {
		return err
	}

	return nil
}
