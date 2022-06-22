package storage

import (
	"context"
	"mime/multipart"
	"rent/src/entities"
	"rent/src/logger"
	"rent/src/storage"
	"rent/src/utils/errors"
)

type MinioStorage struct {
	Storage storage.Storage
}

func (s MinioStorage) SaveImage(ctx context.Context, bucketName string, fileName string, file *multipart.FileHeader) *errors.RestError {
	if err := s.Storage.Save(ctx, bucketName, fileName, file); err != nil {
		logger.Error("Error happened on saving Image " + err.Error())
		return errors.NewInternalServerError("Internal server error happened", "")
	}
	return nil
}

func NewMinioStorage(s storage.Storage) entities.Storage {
	return MinioStorage{
		Storage: s,
	}
}
