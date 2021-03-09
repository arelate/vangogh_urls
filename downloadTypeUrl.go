package vangogh_urls

import (
	"fmt"
	"github.com/arelate/vangogh_types"
	"path"
)

var downloadTypeRootDirs = map[vangogh_types.DownloadType]string{
	vangogh_types.Image:                 "images",
	vangogh_types.BoxArt:                "images",
	vangogh_types.BackgroundImage:       "images",
	vangogh_types.GalaxyBackgroundImage: "images",
	vangogh_types.Logo:                  "images",
	vangogh_types.Icon:                  "images",
}

func DstDownloadTypeUrl(dt vangogh_types.DownloadType) (string, error) {
	if !vangogh_types.ValidDownloadType(dt) {
		return "", fmt.Errorf("vangogh_urls: no local destination for %s", dt)
	}

	return path.Join(downloadTypeRootDirs[dt], dt.String()), nil
}
