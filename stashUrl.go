package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
	"path"
)

const (
	stashDir          = "_stash"
	distilledStashDir = "_distilled"
)

func DistilledStashUrl() string {
	return path.Join(metadataDir, stashDir, distilledStashDir)
}

func ProductTypeStashUrl(pt vangogh_types.ProductType, mt gog_types.Media) (string, error) {
	if !vangogh_types.ValidProductType(pt) {
		return "", fmt.Errorf("vangogh_urls: can't stash invalid product type %s", pt)
	}
	if !gog_types.ValidMedia(mt) {
		return "", fmt.Errorf("vangogh_urls: can't stash invalid media %s", pt)
	}

	return path.Join(metadataDir, stashDir, pt.String(), mt.String()), nil
}
