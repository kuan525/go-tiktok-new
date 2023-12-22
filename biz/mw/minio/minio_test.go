package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"go-tiktok-new/pkg/constants"
	"testing"
)

func TestBucketExist(t *testing.T) {
	ctx := context.Background()
	exist, err := Client.BucketExists(ctx, constants.MinioVideoBucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exist {
		fmt.Printf("%v found!\n", constants.MinioVideoBucketName)
	} else {
		fmt.Println("not found!")
	}
}

func TestBucketMake(t *testing.T) {
	ctx := context.Background()
	exist, err := Client.BucketExists(ctx, constants.MinioVideoBucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exist {
		fmt.Printf("%v found!\n", constants.MinioVideoBucketName)
	} else {
		err = Client.MakeBucket(ctx, constants.MinioVideoBucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Successfully create mybucket %v\n", constants.MinioVideoBucketName)
	}
}

func TestGetObjURL(t *testing.T) {
	Init()
	ctx := context.Background()
	url, _ := GetObjURL(ctx, constants.MinioVideoBucketName, "1000.1212121212.mp4")
	fmt.Println(url.String())
}
