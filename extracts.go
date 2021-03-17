package vangogh_urls

import (
	"path"
)

const (
	extractsDir = "_extracts"
)

func ExtractsDir() string {
	return path.Join(metadataDir, extractsDir)
}
