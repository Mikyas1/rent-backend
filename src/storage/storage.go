package storage

import (
	"context"
	"mime/multipart"
	"net/url"
)

type Storage interface {
	Save(ctx context.Context, bucketName string, fileName string, file *multipart.FileHeader) error
	Delete(ctx context.Context, bucketName string, fileName string) error
	GetUrl(ctx context.Context, bucketName string, fileName string) (*url.URL, error)
	CreateBucket(ctx context.Context, bucketName string) error
	LoadUrlToString(str *string, bucketName string)
	GetObject(ctx context.Context, fileName string, bucketName string) (buf []byte, contentType *string, err error)
}
