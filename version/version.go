package version

import (
	"io/ioutil"
	"log"

	"github.com/coreos/go-semver/semver"
	"github.com/rezmuh/drone-semver-plugin/config"
	"github.com/rezmuh/drone-semver-plugin/util"
)

// GetVersion returns current version number from file
func GetVersion(versionFile string) *semver.Version {
	currentVersion, err := ioutil.ReadFile(versionFile)
	if err != nil {
		log.Fatalln(err)
	}

	return semver.New(string(currentVersion))
}

// BumpVersion bumps version number and write the
// updated version to c.VersionFile
func BumpVersion(c *config.Configuration) {
	v := GetVersion(c.VersionFile)

	switch c.Bump.Increment {
	case "major":
		v.BumpMajor()
	case "minor":
		v.BumpMinor()
	case "patch":
		v.BumpPatch()
	}

	if c.Bump.Metadata != "" {
		v.Metadata = c.Bump.Metadata
	}

	if c.Bump.PreRelease != "" {
		v.PreRelease = semver.PreRelease(c.Bump.PreRelease)
	}
	util.WriteToFile(c.VersionFile, v.String())
}
