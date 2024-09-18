# gitlab

<a href="https://github.com/jaredallard/gitlab/releases">
	<img alt="Latest Version" src="https://img.shields.io/github/v/release/jaredallard/gitlab?style=for-the-badge">
</a>
<a href="https://github.com/jaredallard/gitlab/blob/main/LICENSE">
	<img alt="License" src="https://img.shields.io/github/license/jaredallard/gitlab?style=for-the-badge">
</a>
<a href="https://github.com/jaredallard/gitlab/actions/workflows/tests.yaml">
	<img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/jaredallard/gitlab/tests.yaml?style=for-the-badge">
</a>
<a href="https://app.codecov.io/gh/jaredallard/gitlab">
	<img alt="Codecov" src="https://img.shields.io/codecov/c/github/jaredallard/gitlab?style=for-the-badge">
</a>

<br />

Wrapper for [go-gitlab] that supports mocking.

## Usage

See our [Go docs](https://pkg.go.dev/github.com/jaredallard/gitlab) as
well as the upstream [go-gitlab] documentation which this package
provides.

### Differences

Due to the original [go-gitlab] `Client` struct using embedded structs
instead of interfaces, you must use the `NewClient` call from this
package.

## Development

All of the code in this repository is generated through the
`tools/codegen` CLI. To change anything, you must add it to that CLI
tool.

The templates used can be found in the `embed` directory in the same CLI
directory.

## Special Thanks

Huge special thanks to the [mockgen] and [ifacemaker] project for making
this possible and saving me a lot of pain w/ the ast package :)

## License

LGPL-3.0

[go-gitlab]: https://github.com/xanzy/go-gitlab
[mockgen]: https://pkg.go.dev/go.uber.org/mock/mockgen
[ifacemaker]: https://github.com/vburenin/ifacemaker@latest
