package ctxlog_test

import (
	"net/http"
	"os"

	"github.com/t11e/go-ctxlog"
	"golang.org/x/net/context"
)

func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	log := ctxlog.FromContext(ctx)
	log.Info("Hello, World!")
	w.WriteHeader(http.StatusOK)
}

func Example() {
	http.ListenAndServe("localhost:8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := ctxlog.Direct{Writer: os.Stdout, Prefix: "[Example] "}.Logger("example")
		ctx := ctxlog.NewContext(context.Background(), log)
		Handler(ctx, w, r)
	}))
}
