# Drone Semantic Versioning (SemVer) Plugin

Drone plugin that lets you increment a version as well as updating version file to external storage. Currently, only S3 external storage is supported.

Planned Support for External Storages:
- [x] Amazon S3
- [ ] Google's GCS
- [ ] HashiCorp's Vault

# Use Case

For my Drone pipeline, I'd like to standardize version number between GitHub tags, Docker tags as well as Helm chart version. With this plugin,
it will retrieve the version number from a version file in external storage. By default, it will save the version number in `.tags` (to make it compatible
with [Drone Docker Plugin](http://plugins.drone.io/drone-plugins/drone-docker/))

Then this file (.tags) can be used as a source for versioning / tagging on other build steps
