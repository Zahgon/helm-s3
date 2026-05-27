package main

import (
	"github.com/spf13/cobra"
)

const versionDesc = `This command prints plugin version.

You can also check the Helm mode in which the plugin operates, either v2 or v3.
If the plugin does not detect Helm version properly, you can forcefully change 
the mode: set HELM_S3_MODE environment variable to either 2 or 3.
`

func newVersionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// No completions.

type versionAction struct {
	printer printer

	// flags

	mode bool
}

func (act *versionAction) run() error { _ = "STUB: not implemented"; return nil }
