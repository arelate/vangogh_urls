package vangogh_urls

import (
	"fmt"
	"path"
	"strings"
)

const RootVideosDir = "videos"

func VideoDir(videoId string) (string, error) {
	if videoId == "" {
		return "", fmt.Errorf("vangogh_urls: empty video-id")
	}
	if len(videoId) < 2 {
		return "", fmt.Errorf("vangogh_urls: video-id is too short")
	}

	return path.Join(RootVideosDir, strings.ToLower(videoId[0:1])), nil
}
