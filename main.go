package main

import (
	"os"
	"log"

	"github.com/coreos/go-semver/semver"
)

func main() {
	currentVersion := os.Getenv("PLUGIN_VERSION")
	bump := os.Getenv("PLUGIN_BUMP")
	metadata := os.Getenv("PLUGIN_METADATA")
	preRelease := os.Getenv("PLUGIN_PRERELEASE")
	filename := os.Getenv("PLUGIN_FILENAME")

	ver, err := semver.NewVersion(currentVersion)
	
	if err != nil {
		log.Fatalln(err)
	}
	switch bump {
	case "major":
		ver.BumpMajor()
	case "minor":
		ver.BumpMinor()
	case "patch":
		ver.BumpPatch()
	}

	if preRelease != "" {
		ver.PreRelease = semver.PreRelease(preRelease)
	}

	if metadata != "" {
		ver.Metadata = metadata
	}
	newVersion := ver.String()

	if filename == "" {
		filename = "version"
	}

	err = writeToFile(filename, newVersion)

	if err != nil {
		log.Fatalln(err)
	}
}

func writeToFile(filename, content string) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	_, err = f.WriteString(content)

	if err != nil {
		f.Close()
		return err
	}

	err = f.Close()

	if err != nil {
		return err
	}
	return nil
}
