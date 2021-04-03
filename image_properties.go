package vangogh_urls

import (
	"fmt"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_images"
	"net/url"
)

func PropImageUrls(imageIds []string, it vangogh_images.ImageType) ([]*url.URL, error) {
	urls := make([]*url.URL, 0, len(imageIds))

	var getUrl func(string) (*url.URL, error)

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
		getUrl = gog_urls.ImageJpg
	}

	if getUrl == nil {
		return urls, fmt.Errorf("vangogh_urls: no download urls for %s", it)
	}

	for _, imageId := range imageIds {
		if imageId == "" {
			continue
		}
		imageUrl, err := getUrl(imageId)
		if err != nil {
			return urls, err
		}
		urls = append(urls, imageUrl)
	}

	return urls, nil
}
