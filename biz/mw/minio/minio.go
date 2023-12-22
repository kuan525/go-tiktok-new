package minio

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go-tiktok-new/pkg/constants"
	"log"
	"mime/multipart"
	"net/url"
	"time"
)

var (
	Client *minio.Client
	err    error
)

func Init() {
	ctx := context.Background()
	Client, err = minio.New(constants.MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessID, constants.MinioSecretAccessKey, ""),
		Secure: constants.MiniUseSSL,
	})
	if err != nil {
		log.Fatal("minio连接错误: ", err)
	}

	log.Printf("%#v\n", Client)

	MakeBucket(ctx, constants.MinioImgBucketName)
	MakeBucket(ctx, constants.MinioImgBucketName)
}

func MakeBucket(ctx context.Context, bucketName string) {
	exist, err := Client.BucketExists(ctx, bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !exist {
		err = Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Successfully created mybucket %v\n", bucketName)
	}
}

func PutToBucket(ctx context.Context, bucketName string, file *multipart.FileHeader) (info minio.UploadInfo, err error) {
	fileObj, _ := file.Open()
	defer fileObj.Close()

	info, err = Client.PutObject(ctx, bucketName, file.Filename, fileObj, file.Size, minio.PutObjectOptions{})
	return info, err
}

func GetObjURL(ctx context.Context, bucketName, filename string) (u *url.URL, err error) {
	exp := time.Hour * 24
	reqParams := make(url.Values)
	u, err = Client.PresignedGetObject(ctx, bucketName, filename, exp, reqParams)
	return u, err
}

func PutToBucketByBuf(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = Client.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return info, err
}

func PutToBucketByFilePath(ctx context.Context, bucketName, filename, filepath string) (info minio.UploadInfo, err error) {
	info, err = Client.FPutObject(ctx, bucketName, filename, filepath, minio.PutObjectOptions{})
	return info, err
}
