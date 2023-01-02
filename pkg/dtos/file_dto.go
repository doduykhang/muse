package dtos

type FileDTO struct {
	Bytes []byte
	Name  string
	Size  int64
	Field string
	Path  string
}
