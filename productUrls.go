package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_types"
	"net/url"
)

type ProductTypeUrl func(string, gog_types.Media) *url.URL

func SourceUrl(pt vangogh_types.ProductType) (ProductTypeUrl, error) {
	switch pt {
	case vangogh_types.Store:
		return gog_urls.DefaultProductsPage, nil
	case vangogh_types.Account:
		return gog_urls.DefaultAccountProductsPage, nil
	case vangogh_types.Wishlist:
		return gog_urls.DefaultWishlistPage, nil
	case vangogh_types.Details:
		return gog_urls.Details, nil
	case vangogh_types.ApiProducts:
		return gog_urls.ApiProduct, nil
	default:
		return nil, fmt.Errorf("cannot provide a source url for a type %s\n", pt)
	}
}

func DestinationUrl(pt vangogh_types.ProductType, mt gog_types.Media) (string, error) {
	return fmt.Sprintf("metadata/%s/%s", pt.String(), mt.String()), nil
}
