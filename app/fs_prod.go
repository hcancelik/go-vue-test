//go:build prod

package app

import (
	"io/fs"
)

func getFrontendAssets() fs.FS {
	f, err := fs.Sub(EmbedFrontend, "frontend/dist")
	if err != nil {
		panic(err)
	}

	return f
}
