package helpers

import (
	"fmt"
	"io"

	"golang.org/x/crypto/blake2b"
)

func Checksum(r io.Reader) ([]byte, error) {
	hash, err := blake2b.New512(nil)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(hash, r); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func ByteCountIEC(b int) string {
	var (
		unit = 1024
		div  = unit
		exp  = 0
	)

	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
