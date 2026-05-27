package awss3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"

	"github.com/hypnoglow/helm-s3/internal/helmutil"
)

const (
	// selects serverside encryption for bucket.
	awsS3encryption = "AWS_S3_SSE"

	// s3MetadataSoftLimitBytes is application-specific soft limit
	// for the number of bytes in S3 object metadata.
	s3MetadataSoftLimitBytes = 1900
)

var (
	// ErrBucketNotFound signals that a bucket was not found.
	ErrBucketNotFound = errors.New("bucket not found")

	// ErrObjectNotFound signals that an object was not found.
	ErrObjectNotFound = errors.New("object not found")
)

// New returns a new Storage.
func New(session *session.Session) *Storage { _ = "STUB: not implemented"; return nil }

// Returns desired encryption.
func getSSE() *string { _ = "STUB: not implemented"; return nil }

// Storage provides an interface to work with AWS S3 objects by s3 protocol.
type Storage struct {
	session *session.Session
}

// Traverse traverses all charts in the repository.
func (s *Storage) Traverse(ctx context.Context, repoURI string) (<-chan ChartInfo, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// traverse traverses all charts in the repository.
// It writes an info item about every chart to items, and errors to errs.
// It always closes both channels when returns.
func (s *Storage) traverse(ctx context.Context, repoURI string, items chan<- ChartInfo, errs chan<- error) {
	_ = "STUB: not implemented" //nolint:funcorder // TODO: needs fixing
	return
}

// We need to make object key relative to repo root.

// Additionally trim prefix slash if exists, because repos can be:
// s3://bucket/repo/subdir OR s3://bucket/repo/subdir/

// This is a subfolder. Ignore it, because chart repository
// is flat and cannot contain nested directories.

// Ignore any file that isn't a chart
// This could include index.yaml
// or any other kind of file that might be in the repo

//nolint:staticcheck // Safe use of strings.Title
//nolint:staticcheck // Safe use of strings.Title

// Some charts in the repository can have no metadata.
//
// This might happen in few cases:
// - Chart was uploaded manually, not using 'helm s3 push';
// - Chart was pushed before we started adding metadata to objects;
// - Chart metadata was too big to add to the S3 object metadata (see issues
//   https://github.com/hypnoglow/helm-s3/issues/120 and
//   https://github.com/hypnoglow/helm-s3/issues/112 )
//
// In this case we have to download the ch file itself.

// process meta and hash

// Decide if need to load more objects.

// ChartInfo contains info about particular chart.
type ChartInfo struct {
	Meta     helmutil.ChartMetadata
	Filename string
	Hash     string
}

// FetchRaw downloads the object from URI and returns it in the form of byte slice.
// uri must be in the form of s3 protocol: s3://bucket-name/key[...].
func (s *Storage) FetchRaw(ctx context.Context, uri string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exists returns true if an object exists in the storage.
func (s *Storage) Exists(ctx context.Context, uri string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// That's weird that there is no NotFound constant in aws sdk.

// PutChart puts the chart file to the storage.
// uri must be in the form of s3 protocol: s3://bucket-name/key[...].
func (s *Storage) PutChart(
	ctx context.Context,
	uri string,
	r io.Reader,
	chartMeta, acl string,
	chartDigest string,
	contentType string,
	prov bool,
	provReader io.Reader,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// PutIndex puts the index file to the storage.
// uri must be in the form of s3 protocol: s3://bucket-name/key[...].
func (s *Storage) PutIndex(ctx context.Context, uri string, acl string, r io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// IndexExists returns true if index file exists in the storage for repository
// with the provided uri.
// uri must be in the form of s3 protocol: s3://bucket-name/key[...].
func (s *Storage) IndexExists(ctx context.Context, uri string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Delete deletes the object by uri.
// uri must be in the form of s3 protocol: s3://bucket-name/key[...].
func (s *Storage) Delete(ctx context.Context, uri string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteChart deletes the chart object by uri. Also deletes .prov file if exists.
// uri must be in the form of s3 protocol: s3://bucket-name/key[...].
func (s *Storage) DeleteChart(ctx context.Context, uri string) error {
	_ = "STUB: not implemented"
	return nil
}

// parseURI returns bucket and key from URIs like:
//   - s3://bucket-name/dir
//   - s3://bucket-name/dir/file.ext
func parseURI(uri string) (bucket, key string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// assembleObjectMetadata assembles and returns S3 object metadata.
// May return empty metadata if chart metadata is too big.
//
// The user-defined metadata for the object is limited to 2 KB in size.
// To mitigate the issue with large charts which metadata is more than 2 KB,
// we simply drop it. This affects 'reindex' operation, so that it has to download
// the chart file (GET Request) instead of only fetching its metadata (HEAD request).
func assembleObjectMetadata(chartMeta, chartDigest string) map[string]*string {
	_ = "STUB: not implemented"
	return nil
}

// objectMetadataSize calculates object metadata size as described in https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html
// "The size of user-defined metadata is measured by taking the sum of the number of bytes in the UTF-8 encoding of each key and value.".
func objectMetadataSize(m map[string]*string) int { _ = "STUB: not implemented"; return 0 }

const (
	// metaChartMetadata is a s3 object metadata key that represents chart metadata.
	metaChartMetadata = "chart-metadata"

	// metaChartDigest is a s3 object metadata key that represents chart digest.
	metaChartDigest = "chart-digest"
)
