package main

import (
	"time"
)

// options represents global command options (global flags).
type options struct {
	timeout time.Duration
	acl     string
	verbose bool
}

// newDefaultOptions returns default options.
func newDefaultOptions() *options { _ = "STUB: not implemented"; return nil }
