package view

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Routes struct{}

func (r *Routes) Index(ctx *gin.Context) {
	pageData := PageData{
		Title:    "Helper",
		AssetURL: "/assets",
	}

	component := IndexPage(&pageData)

	ctx.Status(200)

	component.Render(context.Background(), ctx.Writer)
}
