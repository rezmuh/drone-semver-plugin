package version

import (
	"fmt"
	"log"
	"strings"
	"io/ioutil"

	"github.com/coreos/go-semver/semver"
	"github.com/rezmuh/drone-semver-plugin/config"
	"github.com/rezmuh/drone-semver-plugin/util"
)

// GetVersion returns current version number from file
func GetVersion(c *config.Configuration) *semver.Version {
	currentVersion, err := ioutil.ReadFile(c.VersionFile)
	if err != nil {
		log.Fatalln(err)
	}
	versionText := fmt.Sprintf("%s", currentVersion)
	v := strings.TrimSpace(versionText)
	return semver.New(v)
}

// BumpVersion bumps version number and write the
// updated version to c.VersionFile
func BumpVersion(c *config.Configuration) {
	v := GetVersion(c)
	log.Println("current version is: ", v)

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

	log.Println("bumped version to: ", v)
	util.WriteToFile(c.VersionFile, v.String())
}
