package main

import (
	"context"

	"github.com/spf13/cobra"
)

const reindexDesc = `This command performs a reindex of the repository.

'helm s3 push' takes one argument:
- REPO - target repository.
`

const reindexExample = `  helm s3 reindex my-repo - performs a reindex of the repository with name 'my-repo'.`

func newReindexCommand(opts *options) *cobra.Command { _ = "STUB: not implemented"; return nil }

// No completions for the REPO argument.

type reindexAction struct {
	printer printer

	// global flags

	acl     string
	verbose bool

	// args

	repoName string

	// flags

	relative bool
}

func (act *reindexAction) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
