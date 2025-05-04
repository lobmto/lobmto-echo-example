package controllers

import (
	"lobmto-echo-example/domain/repositories"
	wordModels "lobmto-echo-example/model/words"
	"lobmto-echo-example/responses/words"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WordController struct {
	wordRepository repositories.WordRepository
}

func NewWordController(wordRepository repositories.WordRepository) *WordController {
	return &WordController{
		wordRepository: wordRepository,
	}
}

func (c *WordController) GetWord(ctx echo.Context) error {
	// TODO: 実装
	// TODO: エラーハンドリング
	meaning, _ := wordModels.NewMeaning("単語")
	word, _ := wordModels.NewWord("word", []wordModels.Meaning{meaning}, []wordModels.Tag{})

	return ctx.JSON(http.StatusOK, words.NewGetWordResponse(word))
}
