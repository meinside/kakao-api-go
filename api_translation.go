package kakaoapi

import (
	"encoding/json"
	"log"
)

// Created on 2018.07.13.
//
// Functions for [Translation APIs](https://developers.kakao.com/docs/latest/ko/translate/dev-guide)

// TypeLanguage is a type for language type for translation
type TypeLanguage string

// Language types
const (
	LanguageKorean     TypeLanguage = "kr"
	LanguageEnglish    TypeLanguage = "en"
	LanguageJapanese   TypeLanguage = "jp"
	LanguageChinese    TypeLanguage = "cn"
	LanguageVietnamese TypeLanguage = "vi"
	LanguageIndonesian TypeLanguage = "id"
	LanguageArabian    TypeLanguage = "ar"
	LanguageBengali    TypeLanguage = "bn"
	LanguageGerman     TypeLanguage = "de"
	LanguageSpanish    TypeLanguage = "es"
	LanguageFrench     TypeLanguage = "fr"
	LanguageHindi      TypeLanguage = "hi"
	LanguageItalian    TypeLanguage = "it"
	LanguageMalaysian  TypeLanguage = "ms"
	LanguageDutch      TypeLanguage = "nl"
	LanguagePortuguese TypeLanguage = "pt"
	LanguageRussian    TypeLanguage = "ru"
	LanguageThai       TypeLanguage = "th"
	LanguageTurkish    TypeLanguage = "tr"
	LanguageInvalid    TypeLanguage = "_"
)

// TranslateText translates given text
//
// https://developers.kakao.com/docs/latest/ko/translate/dev-guide#trans-sentence
func (c *Client) TranslateText(text string, fromLanguage, toLanguage TypeLanguage) (ResponseTranslatedText, error) {
	var bytes []byte
	var err error
	bytes, err = c.post(APIBaseURL+"/v2/translation/translate", authTypeKakaoAK, nil, map[string]interface{}{
		"query":       text,
		"src_lang":    string(fromLanguage),
		"target_lang": string(toLanguage),
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

// DetectLanguage detects language from given text
//
// https://developers.kakao.com/docs/latest/ko/translate/dev-guide#language-detect
func (c *Client) DetectLanguage(text string) (ResponseDetectedLanguage, error) {
	var bytes []byte
	var err error
	bytes, err = c.post(APIBaseURL+"/v3/translation/language/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"query": text,
	})

	if err == nil {
		var response ResponseDetectedLanguage
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseDetectedLanguage{}, err
}
