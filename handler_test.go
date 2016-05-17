package ctxlog_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t11e/go-ctxlog"
	"golang.org/x/net/context"
)

func TestFromContext(t *testing.T) {
	log := ctxlog.FromContext(context.Background())
	assert.NotNil(t, log)
	log.Info("should not blow up")

	log = ctxlog.Direct{Prefix: "testprefix"}.Logger("testmodule")
	ctx := ctxlog.NewContext(context.Background(), log)
	assert.Equal(t, log, ctxlog.FromContext(ctx))
}
