package vangogh_urls

import (
	"io/fs"
	"path/filepath"
)

func LocalVideoIds() (videoIds map[string]bool, err error) {
	videoIds = make(map[string]bool, 0)

	err = filepath.WalkDir(RootVideosDir, func(p string, _ fs.DirEntry, err error) error {
		videoId := id(p)
		if videoId != "" {
			videoIds[videoId] = true
		}
		return err
	})

	return videoIds, err
}
