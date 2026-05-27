package main

import (
	"github.com/spf13/cobra"
)

type errorType int

const (
	errorTypeBadUsage errorType = iota + 1
	errorTypeSilent
)

func (t errorType) Is(err error) bool { _ = "STUB: not implemented"; return false }

type customError struct {
	errType errorType
	err     error
}

func (c customError) Error() string { _ = "STUB: not implemented"; return "" }

func (c customError) errorType() errorType { _ = "STUB: not implemented"; return *new(errorType) }

func newBadUsageError(err error) error { _ = "STUB: not implemented"; return nil }

func newSilentError() error { _ = "STUB: not implemented"; return nil }

func wrapPositionalArgsBadUsage(f cobra.PositionalArgs) cobra.PositionalArgs {
	_ = "STUB: not implemented"
	return *new(cobra.PositionalArgs)
}
