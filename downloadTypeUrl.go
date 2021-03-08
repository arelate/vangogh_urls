package vangogh_urls

import (
	"fmt"
	"github.com/arelate/vangogh_types"
	"path"
)

func DstDownloadTypeUrl(dt vangogh_types.DownloadType) (string, error) {
	if !vangogh_types.ValidDownloadType(dt) {
		return "", fmt.Errorf("vangogh_urls: no local destination for %s", dt)
	}

	dstRootDir := ""

	switch dt {
	case vangogh_types.Image:
		fallthrough
	case vangogh_types.BoxArt:
		fallthrough
	case vangogh_types.BackgroundImage:
		fallthrough
	case vangogh_types.GalaxyBackgroundImage:
		fallthrough
	case vangogh_types.Logo:
		fallthrough
	case vangogh_types.Icon:
		dstRootDir = "images"
	}

	return path.Join(dstRootDir, dt.String()), nil
}
