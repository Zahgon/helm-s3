package helmutil

import (
	"io"

	"helm.sh/helm/v3/pkg/chart"
)

// ChartV3 implements Chart in Helm v3.
type ChartV3 struct {
	chart *chart.Chart
}

func (c ChartV3) Name() string { _ = "STUB: not implemented"; return "" }

func (c ChartV3) Version() string { _ = "STUB: not implemented"; return "" }

func (c ChartV3) Metadata() ChartMetadata { _ = "STUB: not implemented"; return *new(ChartMetadata) }

func loadChartV3(fpath string) (ChartV3, error) {
	_ = "STUB: not implemented"
	return *new(ChartV3), nil
}

func loadArchiveV3(r io.Reader) (ChartV3, error) {
	_ = "STUB: not implemented"
	return *new(ChartV3), nil
}

type chartMetadataV3 struct {
	meta *chart.Metadata
}

func (c *chartMetadataV3) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *chartMetadataV3) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (c *chartMetadataV3) Value() interface{} { _ = "STUB: not implemented"; return nil }

func newChartMetadataV3() *chartMetadataV3 { _ = "STUB: not implemented"; return nil }
