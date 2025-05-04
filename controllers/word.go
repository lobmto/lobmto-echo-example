package controllers

import (
	"lobmto-echo-example/domain/repositories"
	"lobmto-echo-example/http/words"
	wordModels "lobmto-echo-example/model/words"
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
	// TODO: ユースケース層が欲しい
	// TODO: エラーハンドラーを整備して、JSON ではなくエラーを返すようにする
	id, err := wordModels.NewIDFromString(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id format"})
	}

	word, err := c.wordRepository.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "word not found"})
	}

	return ctx.JSON(http.StatusOK, words.NewGetWordResponse(word))
}

func (c *WordController) PostWord(ctx echo.Context) error {
	// TODO: ユースケース層が欲しい
	// TODO: エラーハンドラーを整備して、JSON ではなくエラーを返すようにする
	request := new(words.RegisterWordRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	word, err := words.NewPostWordModel(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to create word"})
	}

	word, err = c.wordRepository.Create(word)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create word"})
	}

	return ctx.JSON(http.StatusCreated, words.NewPostWordResponse(word))
}

func (c *WordController) DeleteWord(ctx echo.Context) error {
	id, err := wordModels.NewIDFromString(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id format"})
	}

	word, err := c.wordRepository.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "word not found"})
	}

	err = c.wordRepository.Delete(word)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete word"})
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
