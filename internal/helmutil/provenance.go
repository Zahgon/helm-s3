package helmutil

import (
	"io"
)

// Digest hashes a reader and returns a SHA256 digest.
func Digest(in io.Reader) (string, error) { _ = "STUB: not implemented"; return "", nil }

// DigestFile calculates a SHA256 hash for a given file.
func DigestFile(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }
