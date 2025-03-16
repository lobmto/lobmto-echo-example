package words

import (
	"lobmto-echo-example/model/words"
)

func NewGetWordResponse(word words.Word) *GetWordResponse {
	type ResponseMeaning struct {
		Meaning Meaning "json:\"meaning\""
	}
	responseMeaningList := make(MeaningList, len(word.MeaningList()))
	for i, meaning := range word.MeaningList() {
		responseMeaningList[i] = ResponseMeaning{
			Meaning: meaning.String(),
		}
	}

	type ResponseTag struct {
		Id Id `json:"id"`
	}
	responseTagList := make(TagList, len(word.TagList()))
	for i, tag := range word.TagList() {
		responseTagList[i] = ResponseTag{
			Id: tag.ID().Value(),
		}
	}

	return &GetWordResponse{
		Id:          word.ID().Value(),
		Word:        word.Word(),
		MeaningList: responseMeaningList,
		TagList:     responseTagList,
	}
}
