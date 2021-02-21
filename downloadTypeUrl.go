package vangogh_urls

import (
	"fmt"
	"github.com/arelate/vangogh_types"
)

func DstDownloadTypeUrl(dt vangogh_types.DownloadType) (string, error) {
	imagesDst := "images"
	switch dt {
	case vangogh_types.ProductImage:
		return imagesDst + dt.String(), nil
	default:
		return "", fmt.Errorf("vangogh_urls: no local destination for %s", dt)
	}
}
