package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"path"
)

const (
	memoriesDst = "_memories"
	gobExt      = ".gob"
)

func MemoriesUrl(pt vangogh_types.ProductType, mt gog_types.Media, property string) (string, error) {
	if !vangogh_types.ValidProductType(pt) {
		return "", fmt.Errorf("vangogh_urls: no memories destination for product type %s", pt)
	}
	if !gog_types.ValidMedia(mt) {
		return "", fmt.Errorf("vangogh_urls: no memories destination for media %s", pt)
	}

	return path.Join(metadataDst, memoriesDst, fmt.Sprintf("%s-%s-%s%s", pt, mt, property, gobExt)), nil
}
