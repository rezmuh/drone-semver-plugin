package storage

import (
	"github.com/rezmuh/drone-semver-plugin/config"
	"github.com/rezmuh/drone-semver-plugin/storage/s3"
)

// DownloadVersionFile is the main function to be called from main
// which will download version file from its storage
func DownloadVersionFile(c *config.Configuration) {

	switch c.Storage {
	case "s3":
		bucket := c.StorageConfig.Source
		key := c.StorageConfig.Path
		versionFile := c.VersionFile
		initialVersion := c.InitialVersion

		s3.DownloadVersionFile(bucket, key, initialVersion, versionFile)
	}
}
