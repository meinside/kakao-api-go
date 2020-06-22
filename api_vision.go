package kakaoapi

import (
	"encoding/json"
	"log"
)

// Created on 2018.05.09.
//
// Functions for [Vision APIs](https://developers.kakao.com/docs/latest/ko/vision/dev-guide)

func (c *Client) detectFace(fp fileParam, threshold float32) (detected ResponseDetectedFace, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/face/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"file":      fp,
		"threshold": threshold,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &detected)
		if err == nil {
			return detected, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while detecting face from file: %s", string(bytes))
		}
	}

	return ResponseDetectedFace{}, err
}

// DetectFaceFromBytes detects faces from the given bytes array
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-face
func (c *Client) DetectFaceFromBytes(imgBytes []byte, threshold float32) (ResponseDetectedFace, error) {
	img := newFileParamFromBytes(imgBytes)

	return c.detectFace(img, threshold)
}

// DetectFaceFromFilepath detects faces from the given filepath of an image
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-face
func (c *Client) DetectFaceFromFilepath(path string, threshold float32) (detected ResponseDetectedFace, err error) {
	var img fileParam
	if img, err = newFileParamFromFilepath(path); err == nil {
		return c.detectFace(img, threshold)
	}

	return ResponseDetectedFace{}, err
}

// DetectFaceFromURL detects faces from the given url of an image
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-face
func (c *Client) DetectFaceFromURL(url string, threshold float32) (detected ResponseDetectedFace, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/face/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
		"threshold": threshold,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &detected)
		if err == nil {
			return detected, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while detecting face from url: %s", string(bytes))
		}
	}

	return ResponseDetectedFace{}, err
}

func (c *Client) detectNSFW(fp fileParam) (detected ResponseDetectedNSFW, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/adult/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"file": fp,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &detected)
		if err == nil {
			return detected, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while detecting NSFW from file: %s", string(bytes))
		}
	}

	return ResponseDetectedNSFW{}, err
}

// DetectNSFWFromBytes detects NSFW from given bytes array
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-adult-content
func (c *Client) DetectNSFWFromBytes(bytes []byte) (detected ResponseDetectedNSFW, err error) {
	img := newFileParamFromBytes(bytes)

	return c.detectNSFW(img)
}

// DetectNSFWFromFilepath detects NSFW from given image of filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-adult-content
func (c *Client) DetectNSFWFromFilepath(path string) (detected ResponseDetectedNSFW, err error) {
	if img, err := newFileParamFromFilepath(path); err == nil {
		return c.detectNSFW(img)
	}

	return ResponseDetectedNSFW{}, err
}

// DetectNSFWFromURL detects NSFW from given image of filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-adult-content
func (c *Client) DetectNSFWFromURL(url string) (detected ResponseDetectedNSFW, err error) {
	bytes, err := c.post(APIBaseURL+"/v1/vision/adult/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &detected)
		if err == nil {
			return detected, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while detecting NSFW from url: %s", string(bytes))
		}
	}

	return ResponseDetectedNSFW{}, err
}

func (c *Client) detectProduct(fp fileParam, threshold float32) (detected ResponseDetectedProduct, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/product/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"file":      fp,
		"threshold": threshold,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &detected)
		if err == nil {
			return detected, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while detecting product from file: %s", string(bytes))
		}
	}

	return ResponseDetectedProduct{}, err
}

// DetectProductFromBytes detects products from the given bytes array
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-product
func (c *Client) DetectProductFromBytes(bytes []byte, threshold float32) (detected ResponseDetectedProduct, err error) {
	img := newFileParamFromBytes(bytes)

	return c.detectProduct(img, threshold)
}

// DetectProductFromFilepath detects products from the given filepath of an image
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-product
func (c *Client) DetectProductFromFilepath(path string, threshold float32) (detected ResponseDetectedProduct, err error) {
	var img fileParam
	if img, err = newFileParamFromFilepath(path); err == nil {
		return c.detectProduct(img, threshold)
	}

	return ResponseDetectedProduct{}, err
}

// DetectProductFromURL detects products from the given url of an image
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-product
func (c *Client) DetectProductFromURL(url string, threshold float32) (detected ResponseDetectedProduct, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/product/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
		"threshold": threshold,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &detected)
		if err == nil {
			return detected, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while detecting product from url: %s", string(bytes))
		}
	}

	return ResponseDetectedProduct{}, err
}

func (c *Client) cropThumbnail(fp fileParam, width, height int) (cropped ResponseCroppedThumbnail, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/thumbnail/crop", authTypeKakaoAK, nil, map[string]interface{}{
		"file":   fp,
		"width":  width,
		"height": height,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &cropped)
		if err == nil {
			return cropped, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while cropping thumbnail from file: %s", string(bytes))
		}
	}

	return ResponseCroppedThumbnail{}, err
}

// CropThumbnailFromBytes crops given bytes array
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-thumbnail
func (c *Client) CropThumbnailFromBytes(bytes []byte, width, height int) (cropped ResponseCroppedThumbnail, err error) {
	img := newFileParamFromBytes(bytes)

	return c.cropThumbnail(img, width, height)
}

// CropThumbnailFromFilepath crops given image from filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-thumbnail
func (c *Client) CropThumbnailFromFilepath(path string, width, height int) (cropped ResponseCroppedThumbnail, err error) {
	var img fileParam
	if img, err = newFileParamFromFilepath(path); err == nil {
		return c.cropThumbnail(img, width, height)
	}

	return ResponseCroppedThumbnail{}, err
}

// CropThumbnailFromURL crops given image from url
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-thumbnail
func (c *Client) CropThumbnailFromURL(url string, width, height int) (cropped ResponseCroppedThumbnail, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/thumbnail/crop", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
		"width":     width,
		"height":    height,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &cropped)
		if err == nil {
			return cropped, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while cropping thumbnail from url: %s", string(bytes))
		}
	}

	return ResponseCroppedThumbnail{}, err
}

func (c *Client) suggestThumbnail(fp fileParam, width, height int) (suggested ResponseSuggestedThumbnail, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/thumbnail/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"file":   fp,
		"width":  width,
		"height": height,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &suggested)
		if err == nil {
			return suggested, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while suggesting thumbnail from file: %s", string(bytes))
		}
	}

	return ResponseSuggestedThumbnail{}, err
}

// SuggestThumbnailFromBytes generates a thumbnail from an image of bytes array
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#extract-thumbnail
func (c *Client) SuggestThumbnailFromBytes(bytes []byte, width, height int) (suggested ResponseSuggestedThumbnail, err error) {
	img := newFileParamFromBytes(bytes)

	return c.suggestThumbnail(img, width, height)
}

// SuggestThumbnailFromFilepath generates a thumbnail from an image of given filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#extract-thumbnail
func (c *Client) SuggestThumbnailFromFilepath(path string, width, height int) (suggested ResponseSuggestedThumbnail, err error) {
	var img fileParam
	if img, err = newFileParamFromFilepath(path); err == nil {
		return c.suggestThumbnail(img, width, height)
	}

	return ResponseSuggestedThumbnail{}, err
}

// SuggestThumbnailFromURL generates a thumbnail from an image of given url
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#extract-thumbnail
func (c *Client) SuggestThumbnailFromURL(url string, width, height int) (suggested ResponseSuggestedThumbnail, err error) {
	bytes, err := c.post(APIBaseURL+"/v1/vision/thumbnail/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
		"width":     width,
		"height":    height,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &suggested)
		if err == nil {
			return suggested, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while suggesting thumbnail from url: %s", string(bytes))
		}
	}

	return ResponseSuggestedThumbnail{}, err
}

func (c *Client) generateTags(fp fileParam) (generated ResponseGeneratedTags, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/multitag/generate", authTypeKakaoAK, nil, map[string]interface{}{
		"file": fp,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &generated)
		if err == nil {
			return generated, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while generating tags from file: %s", string(bytes))
		}
	}

	return ResponseGeneratedTags{}, err
}

// GenerateTagsFromBytes generates tags from an image of given bytes array
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-multi-tag
func (c *Client) GenerateTagsFromBytes(bytes []byte) (generated ResponseGeneratedTags, err error) {
	img := newFileParamFromBytes(bytes)

	return c.generateTags(img)
}

// GenerateTagsFromFilepath generates tags from an image of given filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-multi-tag
func (c *Client) GenerateTagsFromFilepath(path string) (generated ResponseGeneratedTags, err error) {
	var img fileParam
	if img, err = newFileParamFromFilepath(path); err == nil {
		return c.generateTags(img)
	}

	return ResponseGeneratedTags{}, err
}

// GenerateTagsFromURL generates tags from an image of given url
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-multi-tag
func (c *Client) GenerateTagsFromURL(url string) (generated ResponseGeneratedTags, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/multitag/generate", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &generated)
		if err == nil {
			return generated, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while generating tags from url: %s", string(bytes))
		}
	}

	return ResponseGeneratedTags{}, err
}

func (c *Client) detectText(fp fileParam) (detected ResponseDetectedText, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/text/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"file": fp,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &detected)
		if err == nil {
			return detected, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while detecting text from file: %s", string(bytes))
		}
	}

	return ResponseDetectedText{}, err
}

// DetectTextFromBytes detects text area from an image of given bytes array
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#detect-char
func (c *Client) DetectTextFromBytes(bytes []byte) (detected ResponseDetectedText, err error) {
	img := newFileParamFromBytes(bytes)

	return c.detectText(img)
}

// DetectTextFromFilepath detects text area from an image of given filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#detect-char
func (c *Client) DetectTextFromFilepath(path string) (detected ResponseDetectedText, err error) {
	var img fileParam
	if img, err = newFileParamFromFilepath(path); err == nil {
		return c.detectText(img)
	}

	return ResponseDetectedText{}, err
}

func (c *Client) recognizeText(fp fileParam, boxes []DetectedTextBounds) (recognized ResponseRecognizedText, err error) {
	var bytes []byte
	bytes, err = c.post(APIBaseURL+"/v1/vision/text/recognize", authTypeKakaoAK, nil, map[string]interface{}{
		"file":  fp,
		"boxes": boxes,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &recognized)
		if err == nil {
			return recognized, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while recognizing text from file: %s", string(bytes))
		}
	}

	return ResponseRecognizedText{}, err
}

// RecognizeTextFromBytes recognizes text from an image of given bytes array and areas
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-char
func (c *Client) RecognizeTextFromBytes(bytes []byte, boxes []DetectedTextBounds) (recognized ResponseRecognizedText, err error) {
	img := newFileParamFromBytes(bytes)

	return c.recognizeText(img, boxes)
}

// RecognizeTextFromFilepath recognizes text from an image of given filpath and areas
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-char
func (c *Client) RecognizeTextFromFilepath(path string, boxes []DetectedTextBounds) (recognized ResponseRecognizedText, err error) {
	var img fileParam
	if img, err = newFileParamFromFilepath(path); err == nil {
		return c.recognizeText(img, boxes)
	}

	return ResponseRecognizedText{}, err
}
