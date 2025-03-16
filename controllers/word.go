package controllers

import (
	wordModels "lobmto-echo-example/model/words"
	"lobmto-echo-example/responses/words"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WordController struct {
}

func NewWordController() *WordController {
	return &WordController{}
}

func (c *WordController) GetWord(ctx echo.Context) error {
	// TODO: 実装
	// TODO: エラーハンドリング
	meaning, _ := wordModels.NewMeaning("単語")
	word, _ := wordModels.NewWord("word", []wordModels.Meaning{meaning}, []wordModels.Tag{})

	return ctx.JSON(http.StatusOK, words.NewGetWordResponse(word))
}
