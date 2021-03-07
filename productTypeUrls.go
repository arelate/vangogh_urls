package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_types"
	"net/url"
	"path"
)

type ProductTypeUrl func(string, gog_types.Media) *url.URL

const (
	metadataDir  = "metadata"
	denylistsDir = "_denylists"
	txtExt       = ".txt"
)

func SrcProductTypeUrl(pt vangogh_types.ProductType) (ProductTypeUrl, error) {
	if !vangogh_types.ValidProductType(pt) {
		return nil, fmt.Errorf("vangogh_urls: no remote source for %s\n", pt)
	}

	switch pt {
	case vangogh_types.StorePage:
		return gog_urls.DefaultProductsPage, nil
	case vangogh_types.AccountPage:
		return gog_urls.DefaultAccountProductsPage, nil
	case vangogh_types.WishlistPage:
		return gog_urls.DefaultWishlistPage, nil
	case vangogh_types.Details:
		return gog_urls.Details, nil
	case vangogh_types.ApiProductsV1:
		return gog_urls.ApiProductV1, nil
	case vangogh_types.ApiProductsV2:
		return gog_urls.ApiProductV2, nil
	default:
		return nil, fmt.Errorf("vangogh_urls: no remote source for %s\n", pt)
	}
}

func DstProductTypeUrl(pt vangogh_types.ProductType, mt gog_types.Media) (string, error) {
	if !vangogh_types.ValidProductType(pt) {
		return "", fmt.Errorf("vangogh_urls: no local destination for product type %s", pt)
	}
	if !gog_types.ValidMedia(mt) {
		return "", fmt.Errorf("vangogh_urls: no local destination for media %s", pt)
	}

	return path.Join(metadataDir, pt.String(), mt.String()), nil
}

func DenylistUrl(pt vangogh_types.ProductType) string {
	if !vangogh_types.ValidProductType(pt) {
		return ""
	}

	return path.Join(metadataDir, denylistsDir, pt.String()+txtExt)
}
