package ingest

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type Manifest struct {
	ID       string
	Checksum string
	Files    []string
}

func NewManifest(id string, files []string) Manifest {
	dup := append([]string(nil), files...)
	h := sha256.Sum256([]byte(strings.Join(dup, "|")))
	return Manifest{ID: id, Checksum: hex.EncodeToString(h[:]), Files: dup}
}
