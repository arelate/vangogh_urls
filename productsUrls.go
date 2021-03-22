package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_media"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_products"
	"net/url"
	"path"
)

type ProductTypeUrl func(string, gog_media.Media) *url.URL

const (
	metadataDir  = "metadata"
	denylistsDir = "_denylists"
	txtExt       = ".txt"
)

var productTypeUrls = map[vangogh_products.ProductType]ProductTypeUrl{
	vangogh_products.StorePage:     gog_urls.DefaultProductsPage,
	vangogh_products.AccountPage:   gog_urls.DefaultAccountProductsPage,
	vangogh_products.WishlistPage:  gog_urls.DefaultWishlistPage,
	vangogh_products.Details:       gog_urls.Details,
	vangogh_products.ApiProductsV1: gog_urls.ApiProductV1,
	vangogh_products.ApiProductsV2: gog_urls.ApiProductV2,
}

func RemoteProductsUrl(pt vangogh_products.ProductType) (ptUrl ProductTypeUrl, err error) {
	if !vangogh_products.Valid(pt) {
		return nil, fmt.Errorf("vangogh_urls: no remote source for %s\n", pt)
	}

	ptUrl, ok := productTypeUrls[pt]
	if !ok {
		err = fmt.Errorf("vangogh_urls: no remote source for %s\n", pt)
	}

	return ptUrl, err
}

func LocalProductsDir(pt vangogh_products.ProductType, mt gog_media.Media) (string, error) {
	if !vangogh_products.Valid(pt) {
		return "", fmt.Errorf("vangogh_urls: no local destination for product type %s", pt)
	}
	if !gog_media.Valid(mt) {
		return "", fmt.Errorf("vangogh_urls: no local destination for media %s", pt)
	}

	return path.Join(metadataDir, pt.String(), mt.String()), nil
}

func Denylist(pt vangogh_products.ProductType) string {
	if !vangogh_products.Valid(pt) {
		return ""
	}

	return path.Join(metadataDir, denylistsDir, pt.String()+txtExt)
}
