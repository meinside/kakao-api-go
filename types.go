package kakaoapi

import (
	"io/ioutil"
)

///////////////////////////////
// types, structs, and functions for HTTP
//

type authType string

const (
	authTypeBearer  authType = "Bearer"
	authTypeKakaoAK authType = "KakaoAK"
)

// struct for HTTp response
type httpResponse struct {
	status int    // http status
	body   []byte // response body in bytes array
}

// file parameter struct for HTTP POST/PUT
type fileParam struct {
	bytes []byte
}

// newFileParamFromBytes creates a new fileParam from given bytes
func newFileParamFromBytes(bytes []byte) fileParam {
	return fileParam{
		bytes: bytes,
	}
}

// newFileParamFromFilepath creates a new fileParam from given file location
func newFileParamFromFilepath(path string) (fileParam, error) {
	var bytes []byte
	var err error
	if bytes, err = ioutil.ReadFile(path); err == nil {
		return fileParam{
			bytes: bytes,
		}, nil
	}

	return fileParam{}, err
}

///////////////////////////////
// common structs
//

// Point type
type Point []float64 // float64 array with length=2

// X returns `x` from this Point
func (p Point) X() float64 {
	if len(p) > 0 {
		return p[0]
	}
	return 0
}

// Y returns `y` from this Point
func (p Point) Y() float64 {
	if len(p) > 1 {
		return p[1]
	}
	return 0
}

///////////////////////////////
// structs for vision API
//

// ResponseError is the struct for vision API error responses
type ResponseError struct {
	Message string `json:"msg"`
	Code    int    `json:"code"`
}

// ResponseDetectedFace struct
//
// https://developers.kakao.com/docs/restapi/vision#얼굴-검출
type ResponseDetectedFace struct {
	Result struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Faces  []struct {
			FacialAttributes struct {
				Gender struct {
					Male   float64 `json:"male"`
					Female float64 `json:"female"`
				} `json:"gender"`
			} `json:"facial_attributes"`
			FacialPoints struct {
				Jaw          []Point `json:"jaw"`
				RightEyebrow []Point `json:"right_eyebrow"`
				LeftEyebrow  []Point `json:"left_eyebrow"`
				Nose         []Point `json:"nose"`
				RightEye     []Point `json:"right_eye"`
				LeftEye      []Point `json:"left_eye"`
				Lip          []Point `json:"lip"`
			} `json:"facial_points"`
			Score      float64 `json:"score"`
			ClassIndex int     `json:"class_idx"`
			X          float64 `json:"x"`
			Y          float64 `json:"y"`
			W          float64 `json:"w"`
			H          float64 `json:"h"`
			Pitch      float64 `json:"pitch"`
			Yaw        float64 `json:"yaw"`
			Roll       float64 `json:"roll"`
		} `json:"faces"`
	} `json:"result"`
}

// ResponseDetectedProduct struct
//
// https://developers.kakao.com/docs/restapi/vision#상품-검출
type ResponseDetectedProduct struct {
	Result struct {
		Width   int `json:"width"`
		Height  int `json:"height"`
		Objects []struct {
			X1    float64 `json:"x1"`
			Y1    float64 `json:"y1"`
			X2    float64 `json:"x2"`
			Y2    float64 `json:"y2"`
			Class string  `json:"class"`
		} `json:"objects"`
	} `json:"result"`
}

// ResponseCroppedThumbnail struct
//
// https://developers.kakao.com/docs/restapi/vision#썸네일-생성
type ResponseCroppedThumbnail struct {
	ThumbnailImageURL string `json:"thumbnail_image_url"`
}

// ResponseSuggestedThumbnail struct
//
// https://developers.kakao.com/docs/restapi/vision#썸네일-검출
type ResponseSuggestedThumbnail struct {
	Result struct {
		Width     int `json:"width"`
		Height    int `json:"height"`
		Thumbnail struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"thumbnail"`
	} `json:"result"`
}

// ResponseGeneratedTags struct
//
// https://developers.kakao.com/docs/restapi/vision#멀티태그-생성
type ResponseGeneratedTags struct {
	Result struct {
		Labels       []string `json:"label"`
		LabelsKorean []string `json:"label_kr"`
	} `json:"result"`
}

// ResponseDetectedNSFW struct
//
// https://developers.kakao.com/docs/restapi/vision#성인-이미지-판별
type ResponseDetectedNSFW struct {
	Result struct {
		Normal float64 `json:"normal"`
		Soft   float64 `json:"soft"`
		Adult  float64 `json:"adult"`
	} `json:"result"`
}

// ResponseTranslatedText struct
//
// https://developers.kakao.com/docs/restapi/translation#문장번역
type ResponseTranslatedText struct {
	Phrases [][]string `json:"translated_text"`
}
