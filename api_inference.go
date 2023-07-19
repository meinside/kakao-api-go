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
// https://developers.kakao.com/docs/latest/ko/karlo/rest-api#text-to-image
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

// UpscaleImages upscales given images using Karlo.
//
// https://developers.kakao.com/docs/latest/ko/karlo/rest-api#upscale
func (c *Client) UpscaleImages(params ParamsImageUpscale) (res ResponseUpscaledImages, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURLKarlo+"/upscale", authTypeKakaoAK, nil, params)

	if err == nil {
		err = json.Unmarshal(bytes, &res)
		if err == nil {
			return res, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while upscaling images: %s", string(bytes))
		}
	}

	return ResponseUpscaledImages{}, err
}

// VaryImage generates varations for given image using Karlo.
//
// https://developers.kakao.com/docs/latest/ko/karlo/rest-api#variations
func (c *Client) VaryImage(params ParamsImageVariation) (res ResponseVariedImages, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURLKarlo+"/variations", authTypeKakaoAK, nil, params)

	if err == nil {
		err = json.Unmarshal(bytes, &res)
		if err == nil {
			return res, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while varying images: %s", string(bytes))
		}
	}

	return ResponseVariedImages{}, err
}

// CheckNSFW checks whether given image is NSFW using Karlo.
//
// https://developers.kakao.com/docs/latest/ko/karlo/rest-api#nsfw
func (c *Client) CheckNSFW(base64EncodedImages []string) (res ResponseNSFWResult, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURLKarlo+"/nsfw_checker", authTypeKakaoAK, nil, map[string]any{
		"images": base64EncodedImages,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &res)
		if err == nil {
			return res, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while checking NSFW: %s", string(bytes))
		}
	}

	return ResponseNSFWResult{}, err
}
