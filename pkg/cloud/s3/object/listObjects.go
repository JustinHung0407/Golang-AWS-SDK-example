package object

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func ListObjectsV2(svc *s3.S3, bucket string, prefix string) {

	var MaxKeys int64 = 1000

	if len(bucket) < 1 {
		exitErrorf("Bucket name required")
	}

	// Get the list of items
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket), Prefix: aws.String(prefix), MaxKeys: &MaxKeys})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Owner:        ", *item.Owner.DisplayName)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	fmt.Println("Found", len(resp.Contents), "items in bucket", bucket)
	fmt.Println("ContinuationToken: ", resp.ContinuationToken)
	fmt.Println("IsTruncated: ", *resp.IsTruncated)
	fmt.Println("NextContinuationToken: ", *resp.NextContinuationToken)
	fmt.Println("")
}

func ListObjects(svc *s3.S3, bucket string, prefix string) {

	var MaxKeys int64 = 1000

	if len(bucket) < 1 {
		exitErrorf("Bucket name required")
	}

	// Get the list of items
	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket), Prefix: aws.String(prefix), MaxKeys: &MaxKeys})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Owner:        ", *item.Owner.DisplayName)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	fmt.Println("Found", len(resp.Contents), "items in bucket", bucket)
	fmt.Println("IsTruncated: ", *resp.IsTruncated)
	fmt.Println("NextMarker: ", *resp.NextMarker)
	fmt.Println("")
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
