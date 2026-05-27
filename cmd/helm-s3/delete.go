package main

import (
	"context"

	"github.com/spf13/cobra"
)

const deleteDesc = `This command removes a chart from the repository.

'helm s3 delete' takes two arguments:
- NAME - name of the chart to delete,
- REPO - target repository.

[Provenance]

If the chart is signed, the provenance file is removed from the repository as well.
`

const deleteExample = `  helm s3 delete epicservice --version 0.5.1 my-repo - removes the chart with name 'epicservice' and version 0.5.1 from the repository with name 'my-repo'.`

func newDeleteCommand(opts *options) *cobra.Command { _ = "STUB: not implemented"; return nil }

// No completions for the NAME and REPO arguments.

type deleteAction struct {
	printer printer

	// global flags

	acl string

	// args

	chartName string
	repoName  string

	// flags

	version string
}

func (act *deleteAction) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Fetch current index.

// Update index.

// Delete the file from S3 and replace index file.

// For relative URLs we need to prepend base URL.
