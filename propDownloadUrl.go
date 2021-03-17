package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_types"
	"net/url"
	"strings"
)

func PropImageUrls(property string, it vangogh_types.ImageType) ([]*url.URL, error) {
	urls := make([]*url.URL, 0)

	var getUrl func(string) (*url.URL, error)
	var getUrls func([]string) ([]*url.URL, error)

	switch it {
	case vangogh_types.Image:
		fallthrough
	case vangogh_types.BoxArt:
		fallthrough
	case vangogh_types.Background:
		fallthrough
	case vangogh_types.GalaxyBackground:
		fallthrough
	case vangogh_types.Logo:
		fallthrough
	case vangogh_types.Icon:
		getUrl = gog_urls.Image
	case vangogh_types.Screenshots:
		getUrls = gog_urls.Screenshots
	}

	if getUrl == nil && getUrls == nil {
		return urls, fmt.Errorf("vangogh_urls: no download urls for %s", it)
	}

	if getUrl != nil {
		singleUrl, err := getUrl(property)
		if err != nil {
			return urls, err
		}
		urls = append(urls, singleUrl)
	}

	if getUrls != nil {
		manyUrls, err := getUrls(strings.Split(property, ","))
		if err != nil {
			return urls, err
		}
		urls = append(urls, manyUrls...)
	}

	return urls, nil
}
