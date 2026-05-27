package helmutil

import (
	"io"
	"os"

	"k8s.io/helm/pkg/repo"
)

type IndexV2 struct {
	index *repo.IndexFile
}

func (idx *IndexV2) Add(metadata interface{}, filename, baseURL, digest string) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *IndexV2) AddOrReplace(metadata interface{}, filename, baseURL, digest string) error {
	_ = "STUB: not implemented"
	// TODO: this looks like a workaround.
	// Think how we can rework this in the future.
	// Ref: https://github.com/kubernetes/helm/issues/3230
	return nil
}

// TODO: this code is the same as for Helm v3, only chart.Medata struct is from Helm v2 SDK.
// We probably should reduce duplicate code .

// If no chart with such name exists in the index, just create a new
// list of versions.

// If such version exists, replace it.

// Otherwise just add to the list of versions

func (idx *IndexV2) Delete(name, version string) (url string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (idx *IndexV2) Has(name, version string) bool { _ = "STUB: not implemented"; return false }

func (idx *IndexV2) SortEntries() { _ = "STUB: not implemented"; return }

func (idx *IndexV2) UpdateGeneratedTime() { _ = "STUB: not implemented"; return }

func (idx *IndexV2) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *IndexV2) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (idx *IndexV2) Reader() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (idx *IndexV2) WriteFile(dest string, mode os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func newIndexV2() *IndexV2 { _ = "STUB: not implemented"; return nil }

func loadIndexV2(fpath string) (*IndexV2, error) { _ = "STUB: not implemented"; return nil, nil }
