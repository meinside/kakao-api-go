package kakaoapi

import (
	"encoding/json"
	"log"
)

// Created on 2018.05.09.
//
// Functions for [Vision APIs](https://developers.kakao.com/docs/latest/ko/vision/dev-guide)

// DetectFaceFromFilepath detects faces from the given filepath of an image
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-face
func (c *Client) DetectFaceFromFilepath(path string, threshold float32) (ResponseDetectedFace, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		bytes, err = c.post(APIBaseURL+"/v1/vision/face/detect", authTypeKakaoAK, nil, map[string]interface{}{
			"file":      file,
			"threshold": threshold,
		})

		if err == nil {
			var response ResponseDetectedFace
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseDetectedFace{}, err
}

// DetectFaceFromURL detects faces from the given url of an image
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-face
func (c *Client) DetectFaceFromURL(url string, threshold float32) (ResponseDetectedFace, error) {
	bytes, err := c.post(APIBaseURL+"/v1/vision/face/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
		"threshold": threshold,
	})

	if err == nil {
		var response ResponseDetectedFace
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseDetectedFace{}, err
}

// DetectNSFWFromFilepath detects NSFW from given image of filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-adult-content
func (c *Client) DetectNSFWFromFilepath(path string) (ResponseDetectedNSFW, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		bytes, err = c.post(APIBaseURL+"/v1/vision/adult/detect", authTypeKakaoAK, nil, map[string]interface{}{
			"file": file,
		})

		if err == nil {
			var response ResponseDetectedNSFW
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseDetectedNSFW{}, err
}

// DetectNSFWFromURL detects NSFW from given image of filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-adult-content
func (c *Client) DetectNSFWFromURL(url string) (ResponseDetectedNSFW, error) {
	bytes, err := c.post(APIBaseURL+"/v1/vision/adult/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
	})

	if err == nil {
		var response ResponseDetectedNSFW
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseDetectedNSFW{}, err
}

// DetectProductFromFilepath detects products from the given filepath of an image
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-product
func (c *Client) DetectProductFromFilepath(path string, threshold float32) (ResponseDetectedProduct, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		bytes, err = c.post(APIBaseURL+"/v1/vision/product/detect", authTypeKakaoAK, nil, map[string]interface{}{
			"file":      file,
			"threshold": threshold,
		})

		if err == nil {
			var response ResponseDetectedProduct
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseDetectedProduct{}, err
}

// DetectProductFromURL detects products from the given url of an image
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-product
func (c *Client) DetectProductFromURL(url string, threshold float32) (ResponseDetectedProduct, error) {
	bytes, err := c.post(APIBaseURL+"/v1/vision/product/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
		"threshold": threshold,
	})

	if err == nil {
		var response ResponseDetectedProduct
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseDetectedProduct{}, err
}

// CropThumbnailFromFilepath crops given image from filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-thumbnail
func (c *Client) CropThumbnailFromFilepath(path string, width, height int) (ResponseCroppedThumbnail, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		bytes, err = c.post(APIBaseURL+"/v1/vision/thumbnail/crop", authTypeKakaoAK, nil, map[string]interface{}{
			"file":   file,
			"width":  width,
			"height": height,
		})

		if err == nil {
			var response ResponseCroppedThumbnail
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseCroppedThumbnail{}, err
}

// CropThumbnailFromURL crops given image from url
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-thumbnail
func (c *Client) CropThumbnailFromURL(url string, width, height int) (ResponseCroppedThumbnail, error) {
	var bytes []byte
	bytes, err := c.post(APIBaseURL+"/v1/vision/thumbnail/crop", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
		"width":     width,
		"height":    height,
	})

	if err == nil {
		var response ResponseCroppedThumbnail
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseCroppedThumbnail{}, err
}

// SuggestThumbnailFromFilepath generates a thumbnail from an image of given filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#extract-thumbnail
func (c *Client) SuggestThumbnailFromFilepath(path string, width, height int) (ResponseSuggestedThumbnail, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		bytes, err = c.post(APIBaseURL+"/v1/vision/thumbnail/detect", authTypeKakaoAK, nil, map[string]interface{}{
			"file":   file,
			"width":  width,
			"height": height,
		})

		if err == nil {
			var response ResponseSuggestedThumbnail
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseSuggestedThumbnail{}, err
}

// SuggestThumbnailFromURL generates a thumbnail from an image of given url
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#extract-thumbnail
func (c *Client) SuggestThumbnailFromURL(url string, width, height int) (ResponseSuggestedThumbnail, error) {
	bytes, err := c.post(APIBaseURL+"/v1/vision/thumbnail/detect", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
		"width":     width,
		"height":    height,
	})

	if err == nil {
		var response ResponseSuggestedThumbnail
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseSuggestedThumbnail{}, err
}

// GenerateTagsFromFilepath generates tags from an image of given filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-multi-tag
func (c *Client) GenerateTagsFromFilepath(path string) (ResponseGeneratedTags, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		bytes, err = c.post(APIBaseURL+"/v1/vision/multitag/generate", authTypeKakaoAK, nil, map[string]interface{}{
			"file": file,
		})

		if err == nil {
			var response ResponseGeneratedTags
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseGeneratedTags{}, err
}

// GenerateTagsFromURL generates tags from an image of given url
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-multi-tag
func (c *Client) GenerateTagsFromURL(url string) (ResponseGeneratedTags, error) {
	bytes, err := c.post(APIBaseURL+"/v1/vision/multitag/generate", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
	})

	if err == nil {
		var response ResponseGeneratedTags
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseGeneratedTags{}, err
}

// DetectTextFromFilepath detects text area from an image of given filepath
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#detect-char
func (c *Client) DetectTextFromFilepath(path string) (ResponseDetectedText, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		bytes, err = c.post(APIBaseURL+"/v1/vision/text/detect", authTypeKakaoAK, nil, map[string]interface{}{
			"file": file,
		})

		if err == nil {
			var response ResponseDetectedText
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseDetectedText{}, err
}

// RecognizeTextFromFilepath recognizes text from an image of given filpath and areas
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-char
func (c *Client) RecognizeTextFromFilepath(path string, boxes []DetectedTextBounds) (ResponseRecognizedText, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		bytes, err = c.post(APIBaseURL+"/v1/vision/text/recognize", authTypeKakaoAK, nil, map[string]interface{}{
			"file":  file,
			"boxes": boxes,
		})

		if err == nil {
			var response ResponseRecognizedText
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseRecognizedText{}, err
}
