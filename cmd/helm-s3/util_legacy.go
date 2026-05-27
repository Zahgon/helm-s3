package main

import (
	"github.com/spf13/cobra"
)

// detectDownload detects if the plugin runs in downloader mode.
//
// This function is used for compatibility with older implementation
// before the plugin CLI moved to cobra module.
// Remove before v1 release.
func detectDownload(root *cobra.Command, args []string) bool {
	_ = "STUB: not implemented"
	// Downloader plugins has exactly 5 arguments.
	// See https://helm.sh/docs/topics/plugins/#downloader-plugins
	return false
}

func isKnownCommand(commands []*cobra.Command, name string) bool {
	_ = "STUB: not implemented"
	return false
}
