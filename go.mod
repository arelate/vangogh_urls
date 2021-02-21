module github.com/arelate/vangogh_urls

go 1.16

require (
	github.com/arelate/gog_types v0.1.6-alpha
	github.com/arelate/gog_urls v0.1.7-alpha
	github.com/arelate/vangogh_types v0.1.0-alpha
)

replace (
	github.com/arelate/gog_types => ../gog_types
	github.com/arelate/gog_urls => ../gog_urls
	github.com/arelate/vangogh_types => ../vangogh_types
)
