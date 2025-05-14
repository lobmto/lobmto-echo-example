package controllers

import (
	domain "lobmto-echo-example/domain/words"
	"lobmto-echo-example/http/words"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WordsController struct {
	wordRepository domain.WordsRepository
}

func NewWordController(wordRepository domain.WordsRepository) *WordsController {
	return &WordsController{
		wordRepository: wordRepository,
	}
}

func (c *WordsController) GetWord(ctx echo.Context) error {
	// TODO: ユースケース層が欲しい
	// TODO: エラーハンドラーを整備して、JSON ではなくエラーを返すようにする
	id, err := domain.NewIDFromString(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id format"})
	}

	word, err := c.wordRepository.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "word not found"})
	}

	return ctx.JSON(http.StatusOK, words.NewGetWordResponse(word))
}

func (c *WordsController) PostWord(ctx echo.Context) error {
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

func (c *WordsController) DeleteWord(ctx echo.Context) error {
	id, err := domain.NewIDFromString(ctx.Param("id"))
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
