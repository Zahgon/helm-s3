package main

import (
	"github.com/spf13/cobra"
)

const rootDesc = `Manage chart repositories on Amazon S3.

This plugin provides AWS S3 integration for Helm.

Basic usage:

  $ helm s3 init s3://bucket-name/charts

  $ helm repo add mynewrepo s3://bucket-name/charts

  $ helm s3 push ./epicservice-0.7.2.tgz mynewrepo

  $ helm search repo mynewrepo

  $ helm fetch mynewrepo/epicservice --version 0.7.2

  $ helm s3 delete epicservice --version 0.7.2 mynewrepo

For detailed documentation, see README at https://github.com/hypnoglow/helm-s3

[ACL]

You can use ACL with the charts and index files in your repository.
See: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
Note that if you do use ACL, you need to add '--acl' flag for all commands, even
for 'delete', because the index file is still updated when you remove a chart.

[Timeouts]

The default timeout for all commands is 5 minutes. If you don't use MFA, it may
be reasonable to lower the timeout for the most commands, e.g. to 10 seconds.
In contrast, in cases where you want to reindex big repository with thousands of
charts, you definitely want to increase the timeout.

[Verbose output]

You can enable verbose output with '--verbose' flag.
`

func newRootCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Completion is disabled for now.
// Also, see: https://helm.sh/docs/topics/plugins/#static-auto-completion

// The command may produce system error, even if the usage is correct.

// We handle errors by ourselves.
