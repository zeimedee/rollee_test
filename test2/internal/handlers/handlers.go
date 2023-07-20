package handlers

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/zeimedee/test2/internal/services"
)

type WordHandler struct {
	wordService *services.WordService
}

func NewWordHandler(wordService *services.WordService) *WordHandler {
	return &WordHandler{
		wordService: wordService,
	}
}

func (w *WordHandler) StoreWord(ctx *gin.Context) {
	raw := ctx.Query("word")
	word, err := url.QueryUnescape(raw)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid Word"})
		return
	}
	if !services.IsValid(word) {
		ctx.JSON(400, gin.H{"message": "Invalid Word"})
		return
	}

	w.wordService.StoreWord(word)

	ctx.JSON(200, gin.H{"message": word + " Stored Successfully"})
}

func (w *WordHandler) RetrieveWord(ctx *gin.Context) {
	raw := ctx.Query("prefix")
	prefix, err := url.QueryUnescape(raw)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid Word"})
		return
	}
	if !services.IsValid(prefix) {
		ctx.JSON(400, gin.H{"message": "Invalid Word"})
		return
	}

	frequentWord := w.wordService.GetFrequentWord(prefix)
	if frequentWord == "" {
		ctx.JSON(400, gin.H{"word": nil})
		return
	}

	ctx.JSON(200, gin.H{"word": frequentWord})
}
