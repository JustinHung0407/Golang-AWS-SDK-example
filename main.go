package main

import (
	s3sdk "awsS3/pkg/cloud/s3/object"
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

	bucket = "velero-test"
	prefix = "restic/morphy-test/snapshots/"

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := s3.New(sess, &aws.Config{
		Region:      aws.String("Default"),
		Endpoint:    aws.String("https://cos.twcc.ai"),
		Credentials: credentials.NewSharedCredentials(".aws/credentials", "default"),
		//LogLevel:    aws.LogLevel(aws.LogDebug),
	})

	//s3sdk.ListBuckets(svc, bucket)

	fmt.Println("ListObjectsV2")
	fmt.Println("-------------------------------------------------------------------------------------------")
	s3sdk.ListObjectsV2(svc, bucket, prefix)

	//fmt.Println("ListObjects")
	//fmt.Println("-------------------------------------------------------------------------------------------")
	//s3sdk.ListObjects(svc, bucket, prefix)

}
