package domain

import "mime/multipart"

type Row struct {
	Site string
	ID   string
}

type UploadFile struct {
	File *multipart.FileHeader
}
