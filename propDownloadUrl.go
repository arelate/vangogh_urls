package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_types"
	"net/url"
)

func PropDownloadUrl(property string, dt vangogh_types.DownloadType) ([]*url.URL, error) {
	urls := make([]*url.URL, 0)

	var getUrl func(string) (*url.URL, error)
	var getUrls func(string) ([]*url.URL, error)

	switch dt {
	case vangogh_types.Image:
		fallthrough
	case vangogh_types.BoxArt:
		fallthrough
	case vangogh_types.BackgroundImage:
		fallthrough
	case vangogh_types.GalaxyBackgroundImage:
		fallthrough
	case vangogh_types.Logo:
		fallthrough
	case vangogh_types.Icon:
		getUrl = gog_urls.Image
	}

	if getUrl == nil && getUrls == nil {
		return urls, fmt.Errorf("vangogh_urls: no download urls for %s", dt)
	}

	if getUrl != nil {
		singleUrl, err := getUrl(property)
		if err != nil {
			return urls, err
		}
		urls = append(urls, singleUrl)
	}

	if getUrls != nil {
		manyUrls, err := getUrls(property)
		if err != nil {
			return urls, err
		}
		urls = append(urls, manyUrls...)
	}

	return urls, nil
}
