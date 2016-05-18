# go-ctxlog

[![Build Status](https://semaphoreci.com/api/v1/projects/717c1b05-a11f-47ba-a8f2-dec48dfbb1a2/814664/badge.svg)](https://semaphoreci.com/t11e/go-ctxlog)

Context aware logging library.
Wraps [github.com/op/go-logging](http://github.com/op/go-logging)
with some convenience methods.

## Usage

Add this library as a dependency to your project.

```bash
glide get github.com/t11e/go-ctxlog
```

In each handler that needs a logger:

```go
func handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
        log := ctxlog.FromContext(ctx)
        // ...
        log.Infof("Something interesting happened for %s", r.URL)
        // ...
}
```

Be sure to supply a logger directly or via middleware when configuring the http server.

For a full stack example, see [./example_test.go](./example_test.go).

## Development

```bash
brew install go glide
glide install
go test $(go list ./... | grep -v /vendor/)
```

Be sure to install our standard [go commit hook](https://github.com/t11e/development-environment#golang-checks).

To run goimports without messing up `vendor/` use `goimports -w $(find . -name '*.go' | grep -v /vendor/)`.
