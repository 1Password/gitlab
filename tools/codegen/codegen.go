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
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/jaredallard/cmdexec"
	"go.rgst.io/stencil/v2/pkg/slogext"
	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"
)

func main() {
	log := slogext.New()
	if err := entrypoint(log); err != nil {
		log.WithError(err).Error("failed to run codegen")
		os.Exit(1)
	}
}

type Service struct {
	// Name of the service as it appeared on the Type.
	Name string

	// Accessor is the name of the field on the client to access the
	// service. This is the name of the field on the go-gitlab.Client
	// struct.
	Accessor string
}

// PrivateName returns a private version of the service name.
func (s *Service) PrivateName() string {
	return strings.ToLower(s.Name[:1]) + s.Name[1:]
}

// getServices returns all of the "Service" fields on the
// [gitlab.Client] provided by go-gitlab.
func getServices(log slogext.Logger, pkgs []*packages.Package) []*Service {
	services := make([]*Service, 0)
	servicesByName := make(map[string]struct{})
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				decl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}

				for _, spec := range decl.Specs {
					spec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}

					if spec.Name.Name != "Client" {
						continue
					}

					st, ok := spec.Type.(*ast.StructType)
					if !ok {
						continue
					}

					for _, field := range st.Fields.List {
						for _, name := range field.Names {
							// Check if the type is named "Service"
							typeName := pkg.TypesInfo.TypeOf(name)
							if typeName == nil {
								continue
							}

							spl := strings.Split(typeName.String(), ".")
							shortTypeName := spl[len(spl)-1]
							if !strings.HasSuffix(shortTypeName, "Service") {
								continue
							}

							log.Debug("Discovered service", "service.name", shortTypeName, "service.accessor", name.Name)
							if _, ok := servicesByName[shortTypeName]; ok {
								continue
							}

							servicesByName[shortTypeName] = struct{}{}
							services = append(services, &Service{
								Name:     shortTypeName,
								Accessor: name.Name,
							})
						}
					}
				}
			}
		}
	}

	// Sort the services by name.
	slices.SortFunc(services, func(i, j *Service) int {
		return strings.Compare(i.Name, j.Name)
	})
	return services
}

// generateInterfaces creates an interface for each of the service types
// passed by name to this function. The interfaces are generated using
// 'ifacemaker' and are generated into the root directory of
// the project.
func generateInterfaces(log slogext.Logger, goGitlabDir string, services []*Service) error {
	files, err := filepath.Glob("*_inf.go")
	if err != nil {
		return fmt.Errorf("failed to glob existing interfaces: %w", err)
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("failed to remove existing interface: %w", err)
		}
	}

	for _, service := range services {
		log.Infof("Generating interface for %s", service.Name)
		cmd := cmdexec.Command(
			"ifacemaker",
			"-f", goGitlabDir+"/*.go",
			"-s", service.Name,
			"-i", service.Name,
			"-p", "gitlab",
			"-y", service.Name+" is an interface for [gitlab.Client."+service.Accessor+"]",
			"-o", strings.ToLower(service.Name)+"_inf.go",
		)
		cmd.UseOSStreams(true)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to generate interface for %s: %w", service, err)
		}
	}

	return nil
}

// generateMocks generates mocks for all of the _inf.go files that were
// generated by the generateInterfaces function. The mocks are generated
// using 'mockgen' and are placed in the mocks/ directory.
func generateMocks(log slogext.Logger) error {
	os.RemoveAll("mocks")
	if err := os.Mkdir("mocks", 0o755); err != nil {
		return fmt.Errorf("failed to create mocks directory: %w", err)
	}

	files, err := filepath.Glob("*_inf.go")
	if err != nil {
		return fmt.Errorf("failed to glob existing interfaces: %w", err)
	}

	for _, file := range files {
		log.Infof("Generating mock for %s", file)
		cmd := cmdexec.Command(
			"mockgen",
			"-package=mocks",
			"-source="+file,
			"-destination="+filepath.Join("mocks", file),
		)
		cmd.UseOSStreams(true)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to generate mock for %s: %w", file, err)
		}

		// Replace the gitlab import with the original package.
		// TODO(jaredallard): Better version of this
		b, err := os.ReadFile(filepath.Join("mocks", file))
		if err != nil {
			return fmt.Errorf("failed to read mock file: %w", err)
		}

		b = []byte(strings.ReplaceAll(string(b), `gitlab "github.com/1password/gitlab"`, `gitlab "gitlab.com/gitlab-org/api/client-go"`))
		//nolint:gosec // Why: OK.
		if err := os.WriteFile(filepath.Join("mocks", file), b, 0o644); err != nil {
			return fmt.Errorf("failed to write mock file: %w", err)
		}
	}

	return nil
}

func entrypoint(log slogext.Logger) error {
	// Read the version of "gitlab.com/gitlab-org/api/client-go" from the top level
	// go.mod.
	b, err := os.ReadFile("go.mod")
	if err != nil {
		return fmt.Errorf("failed to read go.mod: %w", err)
	}

	mf, err := modfile.Parse("go.mod", b, nil)
	if err != nil {
		return fmt.Errorf("failed to parse go.mod: %w", err)
	}

	var goGitlabVersion string
	for _, r := range mf.Require {
		if r.Mod.Path == "gitlab.com/gitlab-org/api/client-go" {
			goGitlabVersion = r.Mod.Version
			break
		}
	}
	if goGitlabVersion == "" {
		return fmt.Errorf("failed to find go-gitlab version")
	}

	log.Info("Detected go-gitlab version", "version", goGitlabVersion)

	// Embedding in the directory was causing issues with the `packages`
	// package.
	tmpDir := filepath.Join(os.TempDir(), "go-gitlab")
	if _, err := os.Stat(tmpDir); err == nil {
		if err := os.RemoveAll(tmpDir); err != nil {
			return fmt.Errorf("failed to remove existing tmpDir: %w", err)
		}
	}

	log.Info("Cloning go-gitlab")
	cmd := cmdexec.Command(
		"git", "clone", "https://gitlab.com/gitlab-org/api/client-go",
		"--single-branch", "--branch", goGitlabVersion,
		tmpDir,
	)
	cmd.UseOSStreams(true)
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Infof("Using go-gitlab from %s", tmpDir)

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedFiles,
		Dir:  tmpDir,
	})
	if err != nil {
		return err
	}
	if len(pkgs) == 0 {
		return fmt.Errorf("no packages found")
	}

	services := getServices(log, pkgs)

	if err := generateAliases(log, services, pkgs); err != nil {
		return err
	}

	if err := generateInterfaces(log, tmpDir, services); err != nil {
		return err
	}

	if err := generateMocks(log); err != nil {
		return err
	}

	if err := generateClients(log, services); err != nil {
		return err
	}

	log.Info("Formatting generated code")
	cmd = cmdexec.Command("mise", "run", "fmt")
	cmd.UseOSStreams(true)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to format generated code: %w", err)
	}

	return nil
}
