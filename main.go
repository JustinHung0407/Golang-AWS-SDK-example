package main

import (
	s3object "awsS3/pkg/cloud/s3/object"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	var (
		bucket string
		prefix string
	)

	//bucket = "test-backup"
	//prefix = "gitops-backup/restic/gitops/snapshots"

	bucket = "prod-tc-system-backup"
	prefix = "restic/gitops/snapshots"

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := s3.New(sess, &aws.Config{
		Region:      aws.String("Default"),
		Endpoint:    aws.String("https://cos.twcc.ai"),
		Credentials: credentials.NewSharedCredentials(".aws/credentials", "default"),
		//LogLevel:    aws.LogLevel(aws.LogDebug),
	})

	//s3bucket.ListBuckets(svc, bucket)
	//
	//fmt.Println("ListObjectsV2")
	//fmt.Println("-------------------------------------------------------------------------------------------")
	//s3object.ListObjectsV2(svc, bucket, prefix)

	fmt.Println("ListObjects")
	fmt.Println("-------------------------------------------------------------------------------------------")
	s3object.ListObjects(svc, bucket, prefix)

}
