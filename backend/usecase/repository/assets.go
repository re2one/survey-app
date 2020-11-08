package repository

type AssetsRepository interface {
	Post(string, string) error
	Upload(string, string) error
}
