package vangogh_urls

import (
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

func LocalImageIds() (imageIds map[string]bool, err error) {
	imageIds = make(map[string]bool, 0)

	err = filepath.WalkDir(RootImagesDir, func(p string, _ fs.DirEntry, err error) error {
		_, imageIdFile := path.Split(p)
		if !strings.HasSuffix(imageIdFile, ".download") {
			imageIds[strings.TrimSuffix(imageIdFile, path.Ext(imageIdFile))] = true
		}
		return err
	})

	return imageIds, err
}
