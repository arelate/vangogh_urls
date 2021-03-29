package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_images"
	"net/url"
	"strings"
)

func PropImageUrls(propValue string, it vangogh_images.ImageType) ([]*url.URL, error) {
	urls := make([]*url.URL, 0)

	var getUrl func(string) (*url.URL, error)
	var getUrls func([]string) ([]*url.URL, error)

	switch it {
	case vangogh_images.Image:
		fallthrough
	case vangogh_images.BoxArt:
		fallthrough
	case vangogh_images.Background:
		fallthrough
	case vangogh_images.GalaxyBackground:
		fallthrough
	case vangogh_images.Logo:
		fallthrough
	case vangogh_images.Icon:
		getUrl = gog_urls.Image
	case vangogh_images.Screenshots:
		getUrls = gog_urls.Screenshots
	}

	if getUrl == nil && getUrls == nil {
		return urls, fmt.Errorf("vangogh_urls: no download urls for %s", it)
	}

	if getUrl != nil {
		singleUrl, err := getUrl(propValue)
		if err != nil {
			return urls, err
		}
		urls = append(urls, singleUrl)
	}

	if getUrls != nil {
		manyUrls, err := getUrls(strings.Split(propValue, ","))
		if err != nil {
			return urls, err
		}
		urls = append(urls, manyUrls...)
	}

	return urls, nil
}
