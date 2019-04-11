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
		region := c.StorageConfig.Region
		bucket := c.StorageConfig.Source
		key := c.StorageConfig.Path
		versionFile := c.VersionFile
		initialVersion := c.InitialVersion

		s3.DownloadVersionFile(region, bucket, key, initialVersion, versionFile)
	}
}

// UpdateVersionFile updates bumped version number
// and then stores it back to the storage
func UpdateVersionFile(c *config.Configuration) {
	switch c.Storage {
	case "s3":
		region := c.StorageConfig.Region
		bucket := c.StorageConfig.Source
		key := c.StorageConfig.Path
		versionFile := c.VersionFile

		s3.UpdateVersionFile(region, bucket, key, versionFile)
	}

}
