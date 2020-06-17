package kakaoapi

import (
	"encoding/json"
	"log"
)

// Created on 2020.06.17.
//
// Functions for [Pose APIs](https://developers.kakao.com/docs/latest/ko/pose/dev-guide)

// AnalyzePoseFromImageFilepath analyzes pose from an image with given filepath
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-submit
func (c *Client) AnalyzePoseFromImageFilepath(path string) (ResponseAnalyzedPose, error) {
	file, err := newFileParamFromFilepath(path)

	if err == nil {
		var bytes []byte
		var err error
		bytes, err = c.post(APICVURL+"/pose", authTypeKakaoAK, nil, map[string]interface{}{
			"file": file,
		})

		if err == nil {
			var response ResponseAnalyzedPose
			err = json.Unmarshal(bytes, &response)
			if err == nil {
				return response, nil
			} else if c.Verbose {
				log.Printf("* Failed to decode bytes: %s", string(bytes))
			}
		}
	}

	return ResponseAnalyzedPose{}, err
}

// AnalyzePoseFromImageURL analyzes pose from an image with given url
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-submit
func (c *Client) AnalyzePoseFromImageURL(url string) (ResponseAnalyzedPose, error) {
	bytes, err := c.post(APICVURL+"/pose", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
	})

	if err == nil {
		var response ResponseAnalyzedPose
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseAnalyzedPose{}, err
}

// AnalyzePoseFromVideoURL analyzes pose from a video with given url
//
// `callbackURL` can be "" (= don't callback)
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-submit
func (c *Client) AnalyzePoseFromVideoURL(videoURL string, smoothing bool, callbackURL string) (ResponseAnalyzedPoseFromVideoURLRequested, error) {
	params := map[string]interface{}{
		"video_url": videoURL,
		"smoothing": smoothing,
	}
	if callbackURL != "" {
		params["callback_url"] = callbackURL
	}

	bytes, err := c.post(APICVURL+"/pose/job", authTypeKakaoAK, nil, params)

	if err == nil {
		var response ResponseAnalyzedPoseFromVideoURLRequested
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseAnalyzedPoseFromVideoURLRequested{}, err
}

// RetrieveAnalyzedPoseFromVideoURL retrives analyzed pose from video url
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-retrieval
func (c *Client) RetrieveAnalyzedPoseFromVideoURL(jobID string) (ResponseAnalyzedPoseFromVideoURL, error) {
	bytes, err := c.post(APICVURL+"/pose/job", authTypeKakaoAK, nil, map[string]interface{}{
		"job_id": jobID,
	})

	if err == nil {
		var response ResponseAnalyzedPoseFromVideoURL
		err = json.Unmarshal(bytes, &response)
		if err == nil {
			return response, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseAnalyzedPoseFromVideoURL{}, err
}
