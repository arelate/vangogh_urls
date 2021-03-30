module github.com/arelate/vangogh_urls

go 1.16

require (
	github.com/arelate/gog_media v0.1.2-alpha
	github.com/arelate/gog_urls v0.1.8-alpha
	github.com/arelate/vangogh_images v0.1.0-alpha
	github.com/arelate/vangogh_products v0.1.0-alpha
)

replace (
	github.com/arelate/gog_media => ../gog_media
	github.com/arelate/gog_urls => ../gog_urls
	github.com/arelate/gog_types => ../gog_types
	github.com/arelate/vangogh_images => ../vangogh_images
	github.com/arelate/vangogh_products => ../vangogh_products
)
