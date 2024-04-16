package domain

import "mime/multipart"

// Row is an entity that represents a Line/Row on reading file
type Row struct {
	Site string
	ID   string
}

type UploadFile struct {
	File *multipart.FileHeader
}
