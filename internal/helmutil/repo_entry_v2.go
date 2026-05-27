package helmutil

import (
	"k8s.io/helm/pkg/repo"
)

// RepoEntryV2 implements RepoEntry in Helm v2.
type RepoEntryV2 struct {
	entry *repo.Entry
}

func (r RepoEntryV2) Name() string { _ = "STUB: not implemented"; return "" }

func (r RepoEntryV2) URL() string { _ = "STUB: not implemented"; return "" }

func (r RepoEntryV2) IndexURL() string { _ = "STUB: not implemented"; return "" }

func (r RepoEntryV2) CacheFile() string { _ = "STUB: not implemented"; return "" }

func lookupV2(name string) (RepoEntryV2, error) {
	_ = "STUB: not implemented"
	return *new(RepoEntryV2), nil
}

func lookupByURLV2(url string) (RepoEntryV2, bool, error) {
	_ = "STUB: not implemented"
	return *new(RepoEntryV2), false, nil
}
