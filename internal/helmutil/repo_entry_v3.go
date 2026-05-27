package helmutil

import (
	"helm.sh/helm/v3/pkg/repo"
)

// RepoEntryV3 implements RepoEntry in Helm v3.
type RepoEntryV3 struct {
	entry *repo.Entry
}

func (r RepoEntryV3) Name() string { _ = "STUB: not implemented"; return "" }

func (r RepoEntryV3) URL() string { _ = "STUB: not implemented"; return "" }

func (r RepoEntryV3) IndexURL() string { _ = "STUB: not implemented"; return "" }

func (r RepoEntryV3) CacheFile() string { _ = "STUB: not implemented"; return "" }

func lookupV3(name string) (RepoEntryV3, error) {
	_ = "STUB: not implemented"
	return *new(RepoEntryV3), nil
}

func lookupByURLV3(url string) (RepoEntryV3, bool, error) {
	_ = "STUB: not implemented"
	return *new(RepoEntryV3), false, nil
}
