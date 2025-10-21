package fake

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed data/**
var embeddedData embed.FS

// FS returns a filesystem rooted at `data`. If useExternal is true it returns
// the OS dir filesystem so external files under `data/` are used.
func FS(useExternal bool) fs.FS {
	if useExternal {
		return os.DirFS("data")
	}
	sub, err := fs.Sub(embeddedData, "data")
	if err != nil {
		// fallback to the raw embeddedFS (shouldn't normally happen)
		return embeddedData
	}
	return sub
}
