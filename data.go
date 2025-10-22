package fake

import (
	"embed"
	"io/fs"
)

//go:embed data/**
var embeddedData embed.FS

// FS returns a filesystem rooted at `data`.
// the OS dir filesystem so external files under `path` are used.
func (f *Fake) getFS() fs.FS {
	if f.fs != nil {
		return f.fs
	}
	sub, err := fs.Sub(embeddedData, "data")
	if err != nil {
		// fallback to the raw embeddedFS (shouldn't normally happen)
		return embeddedData
	}
	return sub
}
