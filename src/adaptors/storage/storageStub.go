package storage

import (
	"context"
	"mime/multipart"
	"rent/src/entities"
	"rent/src/utils/errors"
)

type StorageStub struct {
}

func (s StorageStub) SaveImage(ctx context.Context, bucketName string, fileName string, file *multipart.FileHeader) *errors.RestError {
	return nil
}

func NewStorageStub() entities.Storage {
	return StorageStub{}
}
