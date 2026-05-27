package main

import (
	"context"

	"github.com/spf13/cobra"
)

const downloadDesc = `This command downloads a chart from AWS S3.

Note that this command basically implements downloader plugin for Helm
and not intended to be run explicitly. For more information, see:
https://helm.sh/docs/topics/plugins/#downloader-plugins

'helm s3 download' takes four arguments:
- CERT - certificate file,
- KEY - key file,
- CA - certificate authority file,
- URL - full url.
`

func newDownloadCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// No completions for the arguments.

type downloadAction struct {
	printer printer

	// args

	certFile string
	keyFile  string
	caFile   string
	url      string
}

func (act *downloadAction) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Do not use printer, use os.Stdout directly, as required by Helm.
