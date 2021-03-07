package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"path"
)

const (
	memoriesDst = "_memories"
)

func MemoriesUrl(pt vangogh_types.ProductType, mt gog_types.Media) (string, error) {
	if !vangogh_types.ValidProductType(pt) {
		return "", fmt.Errorf("vangogh_urls: memories not supported for product type %s", pt)
	}
	if !gog_types.ValidMedia(mt) {
		return "", fmt.Errorf("vangogh_urls: memories not supported for media %s", pt)
	}

	return path.Join(metadataDst, memoriesDst, pt.String(), mt.String()), nil
}
