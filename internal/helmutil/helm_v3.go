package helmutil

// This file contains helpers for helm v3.

import (
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/repo"
)

// setupHelm3 sets up environment and function bindings for helm v3.
func setupHelm3() { _ = "STUB: not implemented"; return }

var (
	helm3Env *cli.EnvSettings

	// func that loads helm repo file.
	// Defined for testing purposes.
	helm3LoadRepoFile func(path string) (*repo.File, error)
)

func repoFilePathV3() string { _ = "STUB: not implemented"; return "" }

func cacheDirPathV3() string { _ = "STUB: not implemented"; return "" }
