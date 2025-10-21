package fake

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed data/**
var embeddedData embed.FS

// FS returns a filesystem rooted at `data`.
// the OS dir filesystem so external files under `path` are used.
func FS(path string) fs.FS {
	if path != "" {
		return os.DirFS(path)
	}
	sub, err := fs.Sub(embeddedData, "data")
	if err != nil {
		// fallback to the raw embeddedFS (shouldn't normally happen)
		return embeddedData
	}
	return sub
}
