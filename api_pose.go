package kakaoapi

import (
	"encoding/json"
	"fmt"
	"log"
)

// Created on 2020.06.17.
//
// Functions for [Pose APIs](https://developers.kakao.com/docs/latest/ko/pose/dev-guide)

func (c *Client) analyzePose(fp fileParam) (analyzed ResponseAnalyzedPose, err error) {
	var bytes []byte
	bytes, err = c.post(APICVURL+"/pose", authTypeKakaoAK, nil, map[string]interface{}{
		"file": fp,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &analyzed)
		if err == nil {
			return analyzed, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while analyzing pose from file: %s", string(bytes))
		}
	}

	return ResponseAnalyzedPose{}, err
}

// AnalyzePoseFromBytes analyzes pose from an image with given bytes array
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-submit
func (c *Client) AnalyzePoseFromBytes(bytes []byte) (analyzed ResponseAnalyzedPose, err error) {
	img := newFileParamFromBytes(bytes)

	return c.analyzePose(img)
}

// AnalyzePoseFromImageFilepath analyzes pose from an image with given filepath
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-submit
func (c *Client) AnalyzePoseFromImageFilepath(path string) (analyzed ResponseAnalyzedPose, err error) {
	var img fileParam
	if img, err = newFileParamFromFilepath(path); err == nil {
		return c.analyzePose(img)
	}

	return ResponseAnalyzedPose{}, err
}

// AnalyzePoseFromImageURL analyzes pose from an image with given url
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-submit
func (c *Client) AnalyzePoseFromImageURL(url string) (analyzed ResponseAnalyzedPose, err error) {
	var bytes []byte
	bytes, err = c.post(APICVURL+"/pose", authTypeKakaoAK, nil, map[string]interface{}{
		"image_url": url,
	})

	if err == nil {
		err = json.Unmarshal(bytes, &analyzed)
		if err == nil {
			return analyzed, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while analyzing pose from url: %s", string(bytes))
		}
	}

	return ResponseAnalyzedPose{}, err
}

// AnalyzePoseFromVideoURL analyzes pose from a video with given url
//
// `callbackURL` can be "" (= don't callback)
//
// TODO: not tested yet (only for affiliated developers)
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-submit
func (c *Client) AnalyzePoseFromVideoURL(videoURL string, smoothing bool, callbackURL string) (requested ResponseAnalyzedPoseFromVideoURLRequested, err error) {
	params := map[string]interface{}{
		"video_url": videoURL,
		"smoothing": smoothing,
	}
	if callbackURL != "" {
		params["callback_url"] = callbackURL
	}

	var bytes []byte
	bytes, err = c.post(APICVURL+"/pose/job", authTypeKakaoAK, nil, params)

	if err == nil {
		err = json.Unmarshal(bytes, &requested)
		if err == nil {
			return requested, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes: %s", string(bytes))
		}
	}

	return ResponseAnalyzedPoseFromVideoURLRequested{}, err
}

// RetrieveAnalyzedPoseFromVideoURL retrives analyzed pose from video url
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-retrieval
func (c *Client) RetrieveAnalyzedPoseFromVideoURL(jobID string) (retrieved ResponseAnalyzedPoseFromVideoURL, err error) {
	var bytes []byte
	bytes, err = c.get(APICVURL+fmt.Sprintf("/pose/job/%s", jobID), authTypeKakaoAK, nil, nil)

	if err == nil {
		err = json.Unmarshal(bytes, &retrieved)
		if err == nil {
			return retrieved, nil
		} else if c.Verbose {
			log.Printf("* Failed to decode bytes while retrieving analyzed pose from job id: %s", string(bytes))
		}
	}

	return ResponseAnalyzedPoseFromVideoURL{}, err
}
