package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rezmuh/drone-semver-plugin/util"

	// "bytes"
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

// func createVersionFile(versionFile string) {
// 	bucket := os.Getenv("PLUGIN_AWS_BUCKET")
// 	key := os.Getenv("PLUGIN_AWS_KEY")

// 	if bucket == "" {
// 		log.Fatalln("aws_bucket field is required")
// 	}

// 	if key == "" {
// 		log.Fatalln("aws_key field is required")
// 	}

// 	sess := session.Must(session.NewSession())
// 	uploader := s3manager.NewUploaderWithClient(sess)

//     file, err := os.Open(versionFile)
//     if err != nil {
//         log.Fatalln("could not read version file")
//     }
//     defer file.Close()

//     // Get file size and read the file content into a buffer
//     fileInfo, _ := file.Stat()
//     var size = fileInfo.Size()
//     buffer := make([]byte, size)
//     file.Read(buffer)

// 	params := &s3manager.UploadInput{
//     	Bucket: &bucket,
//     	Key:    &key,
//     	Body:   bytes.NewReader(buffer),
// 	}

// 	// Perform an upload.
// 	result, err := uploader.Upload(params)
// }
