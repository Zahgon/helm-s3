package helmutil

// This file contains helpers for helm v2.

import (
	"k8s.io/helm/pkg/helm/helmpath"
	"k8s.io/helm/pkg/repo"
)

// setupHelm2 sets up environment and function bindings for helm v2.
func setupHelm2() { _ = "STUB: not implemented"; return }

// This is needed because LoadRepositoriesFile returns custom (not
// wrapped) error if the file does not exist.

var (
	helm2Home helmpath.Home

	// func that loads helm repo file.
	// Defined for testing purposes.
	helm2LoadRepoFile func(path string) (*repo.RepoFile, error)
)

const (
	envHelmHome = "HELM_HOME"
)

func resolveHome() helmpath.Home { _ = "STUB: not implemented"; return *new(helmpath.Home) }

func repoFilePathV2() string { _ = "STUB: not implemented"; return "" }

func cacheDirPathV2() string { _ = "STUB: not implemented"; return "" }
