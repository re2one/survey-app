package repository

import (
	"image"
)

type AssetsRepository interface {
	Post(string, string) error
	Upload(string, string) error
	SaveFile(string, string, image.Image, string) error
	GetFilenames(string, string) ([]string, error)
}
