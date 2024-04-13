package http

import (
	"meli/internal/core/domain"
	"meli/internal/core/port"

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

func (ih *ItemHdlr) CreateItem(ctx *gin.Context) {
	data := &domain.Item{}
	res, err := ih.service.CreateItem(ctx, data)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err.Error())
	}

	ctx.JSON(200, res)
}
