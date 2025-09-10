package foodrecipe

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// สร้าง struct สำหรับจัดการ food recipes
}

func NewHandler() Handler {
	return Handler{}
}

func (handler Handler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List of food recipes"})
}
