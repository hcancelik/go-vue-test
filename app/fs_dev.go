//go:build !prod

package app

import (
	"io/fs"
	"os"
)

func getFrontendAssets() fs.FS {
	return os.DirFS("frontend/dist")
}
