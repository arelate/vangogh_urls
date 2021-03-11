package vangogh_urls

import (
	"fmt"
	"github.com/arelate/vangogh_types"
	"path"
)

const imagesDir = "images"

var downloadTypeRootDirs = map[vangogh_types.DownloadType]string{
	vangogh_types.Image:                 imagesDir,
	vangogh_types.BoxArt:                imagesDir,
	vangogh_types.BackgroundImage:       imagesDir,
	vangogh_types.GalaxyBackgroundImage: imagesDir,
	vangogh_types.Logo:                  imagesDir,
	vangogh_types.Icon:                  imagesDir,
	vangogh_types.Screenshots:           imagesDir,
}

func DstDownloadTypeUrl(dt vangogh_types.DownloadType) (string, error) {
	if !vangogh_types.ValidDownloadType(dt) {
		return "", fmt.Errorf("vangogh_urls: no local destination for %s", dt)
	}

	return path.Join(downloadTypeRootDirs[dt], dt.String()), nil
}
