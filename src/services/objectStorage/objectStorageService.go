package objectStorage

import (
	"context"
	"os"
	"rent/src/storage"
	"rent/src/utils/errors"
)

type ObjectService interface {
	GetObject(fileName string) ([]byte, *string, *errors.RestError)
}

type DefaultObjectStorageService struct {
	Storage storage.Storage
}

func (s DefaultObjectStorageService) GetObject(fileName string) ([]byte, *string, *errors.RestError) {
	buf, contentType, er := s.Storage.GetObject(context.Background(), fileName, os.Getenv("MINIO_BUCKET_NAME"))
	if er != nil {
		return nil, nil, errors.NewNotFoundError(er.Error(), "")
	}
	return buf, contentType, nil
}

func NewDefaultObjectStorageService(storage storage.Storage) ObjectService {
	return DefaultObjectStorageService{
		Storage: storage,
	}
}
