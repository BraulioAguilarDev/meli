package domain

import "mime/multipart"

// Entity on parsing file
type Row struct {
	Site string
	ID   string
}

type UploadFile struct {
	File *multipart.FileHeader
}
