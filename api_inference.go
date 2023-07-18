package kakaoapi

import (
	"encoding/json"
	"log"
)

// GenerateTexts generates texts with given params using KoGPT.
//
// https://developers.kakao.com/docs/latest/ko/kogpt/rest-api
func (c *Client) GenerateTexts(params ParamsTextGeneration) (res ResponseGeneratedTexts, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURLKoGPT+"/generation", authTypeKakaoAK, nil, params)

	if err == nil {
		err = json.Unmarshal(bytes, &res)
		if err == nil {
			return res, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while generating texts: %s", string(bytes))
		}
	}

	return ResponseGeneratedTexts{}, err
}

// GenerateImages generates images with given params using Karlo.
//
// https://developers.kakao.com/docs/latest/ko/karlo/rest-api
func (c *Client) GenerateImages(params ParamsImageGeneration) (res ResponseGeneratedImages, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURLKarlo+"/t2i", authTypeKakaoAK, nil, params)

	if err == nil {
		err = json.Unmarshal(bytes, &res)
		if err == nil {
			return res, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while generating images: %s", string(bytes))
		}
	}

	return ResponseGeneratedImages{}, err
}
