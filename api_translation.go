package kakaoapi

import (
	"encoding/json"
	"log"
)

// Created on 2018.07.13.
//
// Functions for [Translation APIs](https://developers.kakao.com/docs/restapi/translation)

// TypeLanguage is a type for language type for translation
type TypeLanguage string

// Language types
const (
	LanguageKorean   TypeLanguage = "kr"
	LanguageEnglish  TypeLanguage = "en"
	LanguageJapanese TypeLanguage = "jp"
	LanguageChinese  TypeLanguage = "cn"
	LanguageInvalid  TypeLanguage = "_"
)

// TranslateText translates given text
//
// https://developers.kakao.com/docs/restapi/vision#얼굴-검출
func (c *Client) TranslateText(text string, fromLanguage, toLanguage TypeLanguage) (ResponseTranslatedText, error) {
	var bytes []byte
	var err error
	bytes, err = c.post(APIBaseURL+"/v1/translation/translate", authTypeKakaoAK, nil, map[string]interface{}{
		"query":       text,
		"src_lang":    fromLanguage,
		"target_lang": toLanguage,
	})

	if err == nil {
		var response ResponseTranslatedText
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseTranslatedText{}, err
}
