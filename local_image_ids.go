package vangogh_urls

import (
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

func id(p string) string {
	_, idFile := path.Split(p)
	if !strings.HasSuffix(idFile, ".download") {
		return strings.TrimSuffix(idFile, path.Ext(idFile))
	}
	return ""
}

func LocalImageIds() (imageIds map[string]bool, err error) {
	imageIds = make(map[string]bool, 0)

	err = filepath.WalkDir(RootImagesDir, func(p string, _ fs.DirEntry, err error) error {
		imageId := id(p)
		if imageId != "" {
			imageIds[imageId] = true
		}
		return err
	})

	return imageIds, err
}
