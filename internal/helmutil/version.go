package helmutil

// IsHelm3 returns true if helm is version 3+.
func IsHelm3() bool {
	_ = "STUB: not implemented"
	// Support explicit mode configuration via environment variable.
	return false
}

// continue to other detection methods.

// helm3Detected returns true if helm is v3.
var helm3Detected func() bool

func helmVersionCommand() bool { _ = "STUB: not implemented"; return false }

// Should not happen in normal cases (when helm is properly installed).
// Anyway, fallback to v3 since helm v2 was deprecated a long time ago
// and the majority of helm-s3 users use v3.

// setupHelmVersionDetection sets up the command used to detect helm version.
func setupHelmVersionDetection() { _ = "STUB: not implemented"; return }
