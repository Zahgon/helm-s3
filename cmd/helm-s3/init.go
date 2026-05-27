package main

import (
	"context"

	"github.com/spf13/cobra"
)

const initDesc = `This command initializes an empty repository on AWS S3.

'helm s3 init' takes one argument:
- URI - URI of the repository.
`

const initExample = `  helm s3 init s3://awesome-bucket/charts - inits chart repository in 'awesome-bucket' bucket under 'charts' path.`

func newInitCommand(opts *options) *cobra.Command { _ = "STUB: not implemented"; return nil }

// No completions for the URI argument.

// We don't use cobra's feature
//
//  cmd.MarkFlagsMutuallyExclusive("force", "ignore-if-exists")
//
// because the error message is confusing. Instead, we check the flags
// manually in run().

type initAction struct {
	printer printer

	// global flags

	acl string

	// args

	uri string

	// flags

	force          bool
	ignoreIfExists bool
}

func (act *initAction) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// fallthrough on --force

// TODO:
// do we need to automatically do `helm repo add <name> <uri>`,
// like we are doing `helm repo update` when we push a chart
// with this plugin?

func (act *initAction) checkRepoEntry() error { _ = "STUB: not implemented"; return nil }

// Repo file may not exist, this is OK for instance when the helm is
// just installed (e.g. in docker).

// Repo entry not found - all is good.

// Repo entry exists.

// fallthrough on --force

func (act *initAction) ignoreIfExistsError(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (act *initAction) ignoreIfExistsInStorageError() error { _ = "STUB: not implemented"; return nil }

func (act *initAction) alreadyExistsError(name string) error { _ = "STUB: not implemented"; return nil }

func (act *initAction) alreadyExistsInStorageError() error { _ = "STUB: not implemented"; return nil }
