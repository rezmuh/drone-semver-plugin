package config

import (
	"errors"
	"os"
)

// Configuration base
type Configuration struct {
	Operation string
	Storage string
	InitialVersion string 
	VersionFile string
	Bump *Bump
	StorageConfig *StorageConfiguration
}

// Bump operation struct
type Bump struct {
	Increment string
	Metadata string
	PreRelease string
}

// StorageConfiguration normalizes different field from different
// storage type
type StorageConfiguration struct {
	Source string // this can be url, bucket, etc
	Path string
}

// FromEnv creates new Configuration based on Environment variables.
// Returns an error if minimum conditions are not met
func FromEnv() (Configuration, error) {
	var c Configuration
	var b Bump
	var sc StorageConfiguration

	operation := os.Getenv("PLUGIN_OPERATION")
	storage := os.Getenv("PLUGIN_STORAGE")
	versionFile := os.Getenv("PLUGIN_VERSION_FILE")

	increment := os.Getenv("PLUGIN_INCREMENT")
	metadata := os.Getenv("PLUGIN_METADATA")
	preRelease := os.Getenv("PLUGIN_PRERELEASE")
	initialVersion := os.Getenv("PLUGIN_INITIAL_VERSION")

	switch storage {
	case "s3":
		bucket := os.Getenv("PLUGIN_AWS_BUCKET")
		key := os.Getenv("PLUGIN_AWS_KEY")
		
		if bucket == "" {
			err := errors.New("Bucket is required when choosing aws storage")
			return c, err
		}

		if key == "" {
			err := errors.New("key is required when choosing aws storage")
			return c, err
		}
		sc = StorageConfiguration{
			Source: bucket,
			Path: key,
		}
	default:
		err := errors.New("Storage type is not supported")
		return c, err
	}

	switch operation {
	case "bump":
		b = Bump{
			Increment: increment,
			Metadata: metadata,
			PreRelease: preRelease,
		}
	default:
		err := errors.New("Operation is a required field")
		return c, err
	}

	if versionFile == "" {
		versionFile = "version.txt"
	}

	c = Configuration{
		Operation: operation,
		Storage: storage,
		InitialVersion: initialVersion,
		VersionFile: versionFile,
	}

	if b.Increment != "" {
		c.Bump = &b
	}

	if sc.Source != "" {
		c.StorageConfig = &sc
	}
	return c, nil
}
