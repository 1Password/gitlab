# gitlab

<a href="https://github.com/1password/gitlab/releases">
  <img alt="Latest Version" src="https://img.shields.io/github/v/release/1password/gitlab?style=for-the-badge">
</a>
<a href="https://github.com/1password/gitlab/blob/main/LICENSE">
  <img alt="License" src="https://img.shields.io/github/license/1password/gitlab?style=for-the-badge">
</a>
<a href="https://github.com/1password/gitlab/actions/workflows/tests.yaml">
  <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/1password/gitlab/tests.yaml?style=for-the-badge">
</a>
<a href="https://app.codecov.io/gh/1password/gitlab">
  <img alt="Codecov" src="https://img.shields.io/codecov/c/github/1password/gitlab?style=for-the-badge">
</a>

<br />

Wrapper for [go-gitlab]
([gitlab.com/gitlab-org/api/client-go](gitlab.com/gitlab-org/api/client-go))
that supports mocking.

## Usage

See our [Go docs](https://pkg.go.dev/github.com/1password/gitlab) as
well as the upstream [go-gitlab] documentation which this package
provides.

### Differences

Due to the original [go-gitlab] `Client` struct using embedded structs
instead of interfaces, you must use the `NewClient` call from this
package.

Example:

```go
gl, err := gitlab.NewClient()
if err != nil {
  // handle err
}

// Original
gl.MergeRequests.GetMergeRequest()

// New
gl.MergeRequests().GetMergeRequest()
```

You will also need to dereference `gitlab.Client` as it is now and
interface instead of a struct.

```go
// Original
type MyStruct {
  gl *gitlab.Client
}

// New
type MyStruct {
  gl gitlab.Client
}
```

### Using the mocks

All of the mocks are generated via [mockgen] under the hood. You can
access them on the `MockClient` type.

Example:

```go
func TestCanGetMergeRequest(t *testing.T) {
  gl := gitlab.NewMockClient(t)

  // Should be called once w/ the given arguments and return the given
  // result.
  gl.MergeRequestsServiceMock.EXPECT().
    GetMergeRequest(1, 1, &gitlab.GetMergeRequestsOptions{}).
    Return(&gitlab.MergeRequest{
      ID: 1,
    }, nil, nil)

  mr, _, err := gl.MergeRequests().GetMergeRequest(1, 1, &gitlab.GetMergeRequestsOptions{})
  assert.NilError(t, err)
  assert.Equal(t, mr.ID, 1)
}
```

## Development

All of the code in this repository is generated through the
`tools/codegen` CLI. To change anything, you must add it to that CLI
tool.

The templates used can be found in the `embed` directory in the same CLI
directory.

When you bump dependency versions — specifically, the client-go package — be
sure to run `mise generate` to pull in the latest changes to the package.

## Special Thanks

Huge special thanks to the [mockgen] and [ifacemaker] project for making
this possible and saving me a lot of pain w/ the ast package :)

## License

LGPL-3.0

[go-gitlab]: gitlab.com/gitlab-org/api/client-go
[mockgen]: https://pkg.go.dev/go.uber.org/mock/mockgen
[ifacemaker]: https://github.com/vburenin/ifacemaker@latest
