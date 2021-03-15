package vangogh_urls

import (
	"fmt"
	"path"
)

const imagesDir = "images"

func DstImageUrl(imageId string) (string, error) {
	if imageId == "" {
		return "", fmt.Errorf("vangogh_urls: empty image-id")
	}
	if len(imageId) < 2 {
		return "", fmt.Errorf("vangogh_urls: image-id is too short")
	}

	return path.Join(imagesDir, imageId[0:2]), nil
}
