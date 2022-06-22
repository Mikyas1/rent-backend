package entities

import (
	"context"
	"mime/multipart"
	"rent/src/utils/errors"
)

type Storage interface {
	SaveImage(ctx context.Context, bucketName string, fileName string, file *multipart.FileHeader) *errors.RestError
}
