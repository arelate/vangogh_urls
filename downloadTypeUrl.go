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

	return path.Join("images", dt.String()), nil
}
