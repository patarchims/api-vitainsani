package dto

import "mime/multipart"

type (
	DTOTransferFile struct {
		File *multipart.FileHeader `form:"file" binding:"required" bson:"file"`
	}

	DTOTransferFileResponse struct {
		FileName string `json:"fileName"  bson:"fileName"`
	}
)
