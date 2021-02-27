package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_types"
	"net/url"
)

type ProductTypeUrl func(string, gog_types.Media) *url.URL

func SrcProductTypeUrl(pt vangogh_types.ProductType) (ProductTypeUrl, error) {
	switch pt {
	case vangogh_types.Store:
		return gog_urls.DefaultProductsPage, nil
	case vangogh_types.Account:
		return gog_urls.DefaultAccountProductsPage, nil
	case vangogh_types.Wishlist:
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
	switch pt {
	case vangogh_types.Store:
		fallthrough
	case vangogh_types.StoreProducts:
		fallthrough
	case vangogh_types.Account:
		fallthrough
	case vangogh_types.AccountProducts:
		fallthrough
	case vangogh_types.Wishlist:
		fallthrough
	case vangogh_types.WishlistProducts:
		fallthrough
	case vangogh_types.Details:
		fallthrough
	case vangogh_types.ApiProductsV1:
		fallthrough
	case vangogh_types.ApiProductsV2:
		return fmt.Sprintf("metadata/%s/%s", pt.String(), mt.String()), nil
	default:
		return "", fmt.Errorf("vangogh_urls: no local destination for %s", pt)
	}
}

func Denylist(pt vangogh_types.ProductType) string {
	switch pt {
	case vangogh_types.Store:
		fallthrough
	case vangogh_types.StoreProducts:
		fallthrough
	case vangogh_types.Account:
		fallthrough
	case vangogh_types.AccountProducts:
		fallthrough
	case vangogh_types.Wishlist:
		fallthrough
	case vangogh_types.WishlistProducts:
		fallthrough
	case vangogh_types.Details:
		fallthrough
	case vangogh_types.ApiProductsV1:
		fallthrough
	case vangogh_types.ApiProductsV2:
		return fmt.Sprintf("denylists/%s.txt", pt.String())
	default:
		return ""
	}
}
