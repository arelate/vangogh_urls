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
		_, imageId := path.Split(p)
		imageIds[strings.TrimSuffix(imageId, path.Ext(imageId))] = true
		return err
	})

	return imageIds, err
}
