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

	modelWord, err := c.wordRepository.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "word not found"})
	}

	meaning, err := wordModels.NewMeaning(modelWord.Meaning)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create meaning"})
	}

	word, err := wordModels.ReconstructWord(id, modelWord.Word, []wordModels.Meaning{meaning}, []wordModels.Tag{})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to reconstruct word"})
	}

	return ctx.JSON(http.StatusOK, words.NewGetWordResponse(word))
}
