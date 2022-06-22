package img

import (
	"errors"
	"github.com/google/uuid"
	"mime/multipart"
)

func ValidateImage(img *multipart.FileHeader) (*multipart.FileHeader, string, error) {

	imgContentType := img.Header.Get("Content-Type")
	if imgContentType == "image/jpeg" {
		return img, uuid.NewString() + ".jpg", nil
	}
	if imgContentType == "image/png" {
		return img, uuid.NewString() + ".png", nil
	}

	return nil, "", errors.New("image is not jpeg or png")
}
