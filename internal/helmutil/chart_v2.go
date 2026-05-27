package helmutil

import (
	"io"

	"k8s.io/helm/pkg/proto/hapi/chart"
)

// ChartV2 implements Chart in Helm v2.
type ChartV2 struct {
	chart *chart.Chart
}

func (c ChartV2) Name() string { _ = "STUB: not implemented"; return "" }

func (c ChartV2) Version() string { _ = "STUB: not implemented"; return "" }

func (c ChartV2) Metadata() ChartMetadata { _ = "STUB: not implemented"; return *new(ChartMetadata) }

func loadChartV2(fpath string) (ChartV2, error) {
	_ = "STUB: not implemented"
	return *new(ChartV2), nil
}

func loadArchiveV2(r io.Reader) (ChartV2, error) {
	_ = "STUB: not implemented"
	return *new(ChartV2), nil
}

type chartMetadataV2 struct {
	meta *chart.Metadata
}

func (c *chartMetadataV2) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *chartMetadataV2) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (c *chartMetadataV2) Value() interface{} { _ = "STUB: not implemented"; return nil }

func newChartMetadataV2() *chartMetadataV2 { _ = "STUB: not implemented"; return nil }
