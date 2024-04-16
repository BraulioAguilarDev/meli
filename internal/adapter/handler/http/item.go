package http

import (
	"meli/internal/core/domain"
	"meli/internal/core/port"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type ItemHdlr struct {
	service port.ItemService
}

func ProvideItemHandler(
	service port.ItemService,

) *ItemHdlr {
	return &ItemHdlr{
		service,
	}
}

type uploadFileRequest struct {
	File *multipart.FileHeader `form:"file"`
}

func (ih *ItemHdlr) UploadFile(ctx *gin.Context) {
	var request uploadFileRequest

	if err := ctx.ShouldBind(&request); err != nil {
		// tmp
		ctx.AbortWithError(400, err)
		return
	}

	if err := ih.service.UploadFile(ctx, &domain.UploadFile{
		File: request.File,
	}); err != nil {
		// tmp
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, "success")
}
