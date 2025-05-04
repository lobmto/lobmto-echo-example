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

func NewPostWordModel(request *RegisterWordRequest) (words.Word, error) {
	meaningList := make([]words.Meaning, len(request.MeaningList))
	for i, meaning := range request.MeaningList {
		meaning, err := words.NewMeaning(meaning.Meaning)
		if err != nil {
			return words.Word{}, err
		}
		meaningList[i] = meaning
	}
	tagList := make([]words.Tag, len(request.TagIdList))
	for i, tagId := range request.TagIdList {
		tagIdModel, err := words.NewIDFromString(tagId.String())
		if err != nil {
			return words.Word{}, err
		}
		tag := words.NewTag(tagIdModel)
		tagList[i] = tag
	}
	word, err := words.NewWord(request.Word, meaningList, tagList)
	if err != nil {
		return words.Word{}, err
	}
	return word, nil
}

func NewPostWordResponse(word words.Word) *RegisterWordResponse {
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

	return &RegisterWordResponse{
		Id:          word.ID().Value(),
		Word:        word.Word(),
		MeaningList: responseMeaningList,
		TagList:     responseTagList,
	}
}
