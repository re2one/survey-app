package repository

import (
	"image"
	"os"
)

type AssetsRepository interface {
	Post(string, string) error
	Upload(string, string) error
	SaveFile(string, string, image.Image, string) error
	GetFilenames(string, string) ([]string, error)
	Get(string, string, string) (*os.File, error)
}
