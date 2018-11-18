package models

import (
	"fmt"
	"io"
	"os"
)

type ImageService interface {
	Create(galleryID uint, r io.ReadCloser, filename string) error
	// ByGalleryID(galleryID uint) []string
}

func NewImageService() ImageService {
	return &imageService{}
}

type imageService struct{}

func (is *imageService) Create(galleryID uint, r io.ReadCloser, filename string) error {

	path, err := is.mkImagePath(galleryID)
	if err != nil {
		return err
	}

	// create the destination file
	dst, err := os.Create(path + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// copy the reader data to the destination file
	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}
	return nil
}

func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := fmt.Sprintf("images/galleries/%v/", galleryID)
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}