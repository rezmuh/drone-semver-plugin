package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepopulateEnvironment() {
	os.Setenv("PLUGIN_STORAGE", "s3")
	os.Setenv("PLUGIN_AWS_BUCKET", "my-bucket")
	os.Setenv("PLUGIN_AWS_KEY", "file")
}


// This plugin requires a `PLUGIN_OPERATION` or `settings.operation`.
func TestConfiguration_RequireOperation(t *testing.T) {
	_, err := FromEnv()

	assert.EqualError(t, err, "Operation is a required field")
}

// Minimum configuration for a `bump` operation. By default, bump will
// do a `patch` increment if `PLUGIN_INCREMENT ` or `settings.increment`
// is not set.
func TestConfiguration_Bump(t *testing.T) {
	prepopulateEnvironment()

	os.Setenv("PLUGIN_OPERATION", "bump")
	c, err := FromEnv()

	assert.Nil(t, err)

	assert.Equal(t, c.Operation, "bump")
	assert.Equal(t, c.Storage, "s3")
	assert.Equal(t, c.Bump.Increment, "patch")
}

// Test to validate when PLUGIN_INCREMENT is set to `major`
func TestConfiguration_BumpMajor(t *testing.T) {
	prepopulateEnvironment()

	increment := "major"

	os.Setenv("PLUGIN_OPERATION", "bump")
	os.Setenv("PLUGIN_INCREMENT", increment)
	c, err := FromEnv()

	assert.Nil(t, err)
	assert.Equal(t, c.Bump.Increment, increment)
}

// Test to validate when PLUGIN_INCREMENT is set to `minor`
func TestConfiguration_BumpMinor(t *testing.T) {
	prepopulateEnvironment()

	increment := "minor"

	os.Setenv("PLUGIN_OPERATION", "bump")
	os.Setenv("PLUGIN_INCREMENT", increment)
	c, err := FromEnv()

	assert.Nil(t, err)
	assert.Equal(t, c.Bump.Increment, increment)
}

// When `PLUGIN_AWS_REGION` is not set, the default points
// to `ap-southeast-1`
func TestConfiguration_AWSRegion(t *testing.T) {
	prepopulateEnvironment()
	os.Setenv("PLUGIN_OPERATION", "bump")
	c, err := FromEnv()

	assert.Nil(t, err)
	assert.Equal(t, c.StorageConfig.Region, "ap-southeast-1")
}

func TestConfiguration_NeedBucket(t *testing.T) {
	prepopulateEnvironment()
	os.Setenv("PLUGIN_OPERATION", "bump")
	os.Unsetenv("PLUGIN_AWS_BUCKET")

	_, err := FromEnv()

	assert.EqualError(t, err, "Bucket is required when choosing aws storage")
}

func TestConfiguration_NeedKey(t *testing.T) {
	prepopulateEnvironment()
	os.Setenv("PLUGIN_OPERATION", "bump")
	os.Unsetenv("PLUGIN_AWS_KEY")

	_, err := FromEnv()

	assert.EqualError(t, err, "Key is required when choosing aws storage")
}

func TestConfiguration_UnsupportedStorage(t *testing.T) {
	prepopulateEnvironment()
	os.Setenv("PLUGIN_STORAGE", "nots3")

	_, err := FromEnv()

	assert.EqualError(t, err, "Storage type is not supported")
}
