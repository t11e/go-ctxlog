package ctxlog_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t11e/go-ctxlog"
)

func TestDirect(t *testing.T) {
	buf := bytes.Buffer{}
	cl := ctxlog.Direct{Writer: &buf, Prefix: "[testprefix] "}.Logger("testmodule")
	cl.Info("foo bar")
	assert.Equal(t, "[testprefix] foo bar\n", buf.String())

	buf.Truncate(0)
	cl.Info("foo bar\nbaz")
	assert.Equal(t, "[testprefix] foo bar\n[testprefix] baz\n", buf.String())
}

func TestDirect_EmptyPrefix(t *testing.T) {
	buf := bytes.Buffer{}
	cl := ctxlog.Direct{Writer: &buf}.Logger("testmodule")
	cl.Info("foo bar")
	assert.Equal(t, "foo bar\n", buf.String())

	buf.Truncate(0)
	cl.Info("foo bar\nbaz")
	assert.Equal(t, "foo bar\nbaz\n", buf.String())
}

func TestDiscard(t *testing.T) {
	cl := ctxlog.Discard.Logger("testmodule")
	cl.Info("foo bar")
}
