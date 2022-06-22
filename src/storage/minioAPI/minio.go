package minioAPI

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"rent/src/storage"
	"time"
)

type MinIOStorage struct {
	Client        *minio.Client
	UrlExpiryTime time.Duration
}

func (s MinIOStorage) CreateBucket(ctx context.Context, bucketName string) error {
	err := s.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := s.Client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
			return nil
		} else {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (s MinIOStorage) Save(ctx context.Context, bucketName, fileName string, file *multipart.FileHeader) error {
	size := file.Size
	f, _ := file.Open()
	_, err := s.Client.PutObject(ctx, bucketName, fileName, f, size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	return nil
}

func (s MinIOStorage) Delete(ctx context.Context, bucketName, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := s.Client.RemoveObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (s MinIOStorage) GetUrl(ctx context.Context, bucketName, objectName string) (*url.URL, error) {
	reqParams := make(url.Values)
	preSignedURL, err := s.Client.PresignedGetObject(ctx, bucketName, objectName, s.UrlExpiryTime, reqParams)
	if err != nil {
		return nil, err
	}
	return preSignedURL, nil
}

func (s MinIOStorage) LoadUrlToString(str *string, bucketName string) {
	if *str != "" {
		url, err := s.GetUrl(context.Background(), bucketName, *str)
		if err != nil {
			*str = ""
		}
		*str = url.String()
	}
}

func (s MinIOStorage) GetObject(ctx context.Context, fileName string, bucketName string) (buf []byte, contentType *string, err error) {
	object, err := s.Client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Printf("Error getting object from object storage: %v", err)
		return
	}
	defer object.Close()

	fileInfo, err := object.Stat()
	if err != nil {
		return
	}

	//fmt.Println(fileInfo.)

	contentType = &fileInfo.ContentType
	buf = make([]byte, fileInfo.Size)
	object.Read(buf)
	if err != nil {
		fmt.Printf("Error reading object: %v", err)
		return
	}
	return
}

func NewMinIOStorage() (storage.Storage, error) {
	endpoint := os.Getenv("MINIO_END_POINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	useSSL := os.Getenv("MINIO_USE_SSH") == "true"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, err
	}

	fmt.Println("created minio client")
	return MinIOStorage{
		Client:        minioClient,
		UrlExpiryTime: time.Second * 24 * 60 * 60 * 7,
	}, nil
}
