package main

import (
	"log"

	"github.com/rezmuh/drone-semver-plugin/config"
	"github.com/rezmuh/drone-semver-plugin/storage"
	"github.com/rezmuh/drone-semver-plugin/version"
)

func main() {
	c, err := config.FromEnv()

	if err != nil {
		log.Fatalln(err)
	}

	switch c.Operation {
	case "bump":
		storage.DownloadVersionFile(&c)
		version.BumpVersion(&c)
	case "put":
		storage.UpdateVersionFile(&c)
	}
}
