package ctxlog

import (
	"fmt"
	"io"
	"strings"

	"github.com/op/go-logging"
)

type Direct struct {
	io.Writer
	Prefix string
}

func (d Direct) Logger(module string) *logging.Logger {
	if d.Prefix == "" {
		return logger(module, directBackend{d.Writer})
	}
	return logger(module, prefixBackend{d.Writer, d.Prefix})
}

type directBackend struct {
	io.Writer
}

func (b directBackend) Log(_ logging.Level, _ int, r *logging.Record) error {
	_, err := fmt.Fprintf(b, "%s\n", r.Message())
	return err
}

type prefixBackend struct {
	io.Writer
	Prefix string
}

func (b prefixBackend) Log(_ logging.Level, _ int, r *logging.Record) error {
	message := r.Message()
	if strings.Contains(message, "\n") {
		message = strings.Replace(message, "\n", fmt.Sprintf("\n%s", b.Prefix), -1)
	}
	_, err := fmt.Fprintf(b.Writer, "%s%s\n", b.Prefix, message)
	return err
}

var Discard = discard{}

type discard struct{}

func (d discard) Logger(module string) *logging.Logger {
	return logger(module, discardBackend{})
}

type discardBackend struct{}

func (discardBackend) Log(_ logging.Level, _ int, _ *logging.Record) error {
	return nil
}

func logger(module string, backend logging.Backend) *logging.Logger {
	l := logging.Logger{Module: module}
	l.SetBackend(logging.AddModuleLevel(backend))
	return &l
}
