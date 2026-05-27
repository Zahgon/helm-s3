package main

import (
	"context"

	"github.com/spf13/cobra"
)

const pushDesc = `This command uploads a chart to the repository.

'helm s3 push' takes two arguments:
- PATH - path to the chart file,
- REPO - target repository.

[Provenance]

If the chart is signed, the provenance file is uploaded to the repository as well.
`

const pushExample = `  helm s3 push ./epicservice-0.5.1.tgz my-repo - uploads chart file 'epicservice-0.5.1.tgz' from the current directory to the repository with name 'my-repo'.`

func newPushCommand(opts *options) *cobra.Command { _ = "STUB: not implemented"; return nil }

// Allow file completion for the PATH argument.

// No completions for the REPO argument.

// We don't use cobra's feature
//
//  cmd.MarkFlagsMutuallyExclusive("force", "ignore-if-exists")
//
// because the error message is confusing. Instead, we check the flags
// manually in run().

type pushAction struct {
	printer printer

	// global args

	acl string

	// args

	chartPath string
	repoName  string

	// flags

	contentType    string
	dryRun         bool
	force          bool
	ignoreIfExists bool
	relative       bool
}

func (act *pushAction) run(ctx context.Context) error {
	_ = "STUB: not implemented" //nolint:gocyclo // Maybe refactor later.
	return nil
}

// Sanity check.

// Load chart, calculate required params like hash,
// and upload the chart right away.

// if cached index exists, check if the same chart version exists in it.

// fallthrough on --force

// No provenance file, ignore it.

// fallthrough on --force

// The gap between index fetching and uploading should be as small as
// possible to make the best effort to avoid race conditions.
// See https://github.com/hypnoglow/helm-s3/issues/18 for more info.

// Fetch current index, update it and upload it back.

func (act *pushAction) ignoreIfExistsError() error { _ = "STUB: not implemented"; return nil }

func (act *pushAction) chartExistsError() error { _ = "STUB: not implemented"; return nil }
