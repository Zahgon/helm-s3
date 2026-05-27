package helmutil

import (
	"io"
	"os"

	"helm.sh/helm/v3/pkg/repo"
	// Note that this is from Helm v2 SDK because in Helm v3 this package is internal.
)

type IndexV3 struct {
	index *repo.IndexFile
}

func (idx *IndexV3) Add(metadata interface{}, filename, baseURL, digest string) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *IndexV3) AddOrReplace(metadata interface{}, filename, baseURL, digest string) error {
	_ = "STUB: not implemented"
	// TODO: this looks like a workaround.
	// Think how we can rework this in the future.
	// Ref: https://github.com/kubernetes/helm/issues/3230
	return nil
}

// TODO: this code is the same as for Helm v2, only chart.Medata struct is from Helm v3 SDK.
// We probably should reduce duplicate code .

// If no chart with such name exists in the index, just create a new
// list of versions.

// If such version exists, replace it.

// Otherwise just add to the list of versions

func (idx *IndexV3) Delete(name, version string) (url string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (idx *IndexV3) Has(name, version string) bool { _ = "STUB: not implemented"; return false }

func (idx *IndexV3) SortEntries() { _ = "STUB: not implemented"; return }

func (idx *IndexV3) UpdateGeneratedTime() { _ = "STUB: not implemented"; return }

func (idx *IndexV3) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *IndexV3) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (idx *IndexV3) Reader() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (idx *IndexV3) WriteFile(dest string, mode os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func newIndexV3() *IndexV3 { _ = "STUB: not implemented"; return nil }

func loadIndexV3(fpath string) (*IndexV3, error) { _ = "STUB: not implemented"; return nil, nil }
