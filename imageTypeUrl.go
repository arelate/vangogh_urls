package vangogh_urls

import (
	"fmt"
	"github.com/arelate/vangogh_types"
	"path"
)

const imagesDir = "images"

var imageTypeRootDirs = map[vangogh_types.ImageType]string{
	vangogh_types.Image:            imagesDir,
	vangogh_types.BoxArt:           imagesDir,
	vangogh_types.Background:       imagesDir,
	vangogh_types.GalaxyBackground: imagesDir,
	vangogh_types.Logo:             imagesDir,
	vangogh_types.Icon:             imagesDir,
	vangogh_types.Screenshots:      imagesDir,
}

func DstImageTypeUrl(it vangogh_types.ImageType) (string, error) {
	if !vangogh_types.ValidImageType(it) {
		return "", fmt.Errorf("vangogh_urls: no local destination for %s", it)
	}

	return path.Join(imageTypeRootDirs[it], it.String()), nil
}
