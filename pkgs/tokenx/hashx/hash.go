package hashx

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256HexString returns the lowercase hex-encoded SHA-256 digest of s.
func SHA256HexString(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
