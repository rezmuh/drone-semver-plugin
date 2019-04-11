package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rezmuh/drone-semver-plugin/util"

	"bytes"
	"log"
	"os"
)

// DownloadVersionFile downloads version file from S3
func DownloadVersionFile(bucket, key, initialVersion, versionFile string) {
	file, err := os.OpenFile(versionFile, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)

	// we don't care if it fails as it already created a versionFile
	// with an empty content
	_, err = downloader.Download(file,
		&awss3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key: aws.String(key),
		})

	if err != nil {
		v := initialVersion

		if v == "" {
			v = "0.0.0"
		}
		util.WriteToFile(versionFile, v)
	}
}

// UpdateVersionFile for S3 updates the content within versionFile to S3
func UpdateVersionFile(bucket, key, versionFile string) {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	file, err := os.OpenFile(versionFile, os.O_RDONLY, 0644)

	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	params := &s3manager.UploadInput{
		Bucket: &bucket,
		Key   : &key,
		Body  : bytes.NewReader(buffer),
	}
	_, err = uploader.Upload(params)

	if err != nil {
		log.Fatalln(err)
	}
}
