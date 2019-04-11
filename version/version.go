package version

import (
	"io/ioutil"
	"log"

	"github.com/coreos/go-semver/semver"
	"github.com/rezmuh/drone-semver-plugin/config"
	"github.com/rezmuh/drone-semver-plugin/util"
)

// BumpVersion bumps version number and write the
// updated version to c.VersionFile
func BumpVersion(c *config.Configuration) {
	currentVersion, err := ioutil.ReadFile(c.VersionFile)

	if err != nil {
		log.Fatalln(err)
	}
	v := semver.New(string(currentVersion))

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
