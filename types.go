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
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-face
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

// ResponseDetectedNSFW struct
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-adult-content
type ResponseDetectedNSFW struct {
	Result struct {
		Normal float64 `json:"normal"`
		Soft   float64 `json:"soft"`
		Adult  float64 `json:"adult"`
	} `json:"result"`
}

// ResponseDetectedProduct struct
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-product
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
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-thumbnail
type ResponseCroppedThumbnail struct {
	ThumbnailImageURL string `json:"thumbnail_image_url"`
}

// ResponseSuggestedThumbnail struct
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#extract-thumbnail
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
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-multi-tag
type ResponseGeneratedTags struct {
	Result struct {
		Labels       []string `json:"label"`
		LabelsKorean []string `json:"label_kr"`
	} `json:"result"`
}

// DetectedTextBoundPoint type
type DetectedTextBoundPoint []int // [x, y]

// DetectedTextBounds type
type DetectedTextBounds []DetectedTextBoundPoint // [left-upper-point, right-upper-point, right-lower-point, left-lower-point]

// ResponseDetectedText struct
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#detect-char
type ResponseDetectedText struct {
	Result struct {
		Boxes []DetectedTextBounds `json:"boxes"`
	} `json:"result"`
}

// ResponseRecognizedText struct
//
// https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-char
type ResponseRecognizedText struct {
	Result struct {
		RecognizedWords []string `json:"recognition_words"`
	} `json:"result"`
}

// ResponseTranslatedText struct
//
// https://developers.kakao.com/docs/latest/ko/translate/dev-guide#trans-sentence
type ResponseTranslatedText struct {
	Phrases [][]string `json:"translated_text"`
}

// ResponseSpeechToText struct
//
// https://developers.kakao.com/docs/latest/ko/voice/rest-api#speech-to-text
type ResponseSpeechToText struct {
	Type  string                  `json:"type"`
	Value string                  `json:"value"`
	NBest []SpeechToTextCandidate `json:"nBest"`
}

// SpeechToTextCandidate struct
type SpeechToTextCandidate struct {
	Value string `json:"value"`
	Score int    `json:"score"`
}

// ResponseAnalyzedPose struct
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#image-pose-estimation
type ResponseAnalyzedPose []AnalyzedPose

// AnalyzedPose struct
type AnalyzedPose struct {
	Area          float64   `json:"area"`
	BoundingBoxes []float64 `json:"bbox"`
	CategoryID    int       `json:"category_id"` // 1 for Person
	KeyPoints     []float64 `json:"keypoints"`
	Score         float64   `json:"score"`
}

// KeyPointIndex for indexing keypoints
type KeyPointIndex int

// KeyPointIndex values
const (
	KeyPointIndexNose          KeyPointIndex = 0
	KeyPointIndexLeftEye       KeyPointIndex = 1
	KeyPointIndexRightEye      KeyPointIndex = 2
	KeyPointIndexLeftEar       KeyPointIndex = 3
	KeyPointIndexRightEar      KeyPointIndex = 4
	KeyPointIndexLeftShoulder  KeyPointIndex = 5
	KeyPointIndexRightShoulder KeyPointIndex = 6
	KeyPointIndexLeftElbow     KeyPointIndex = 7
	KeyPointIndexRightElbow    KeyPointIndex = 8
	KeyPointIndexLeftWrist     KeyPointIndex = 9
	KeyPointIndexRightWrist    KeyPointIndex = 10
	KeyPointIndexLeftHip       KeyPointIndex = 11
	KeyPointIndexRightHip      KeyPointIndex = 12
	KeyPointIndexLeftKnee      KeyPointIndex = 13
	KeyPointIndexRightKnee     KeyPointIndex = 14
	KeyPointIndexLeftAnkle     KeyPointIndex = 15
	KeyPointIndexRightAnkle    KeyPointIndex = 16
)

// KeyPointFor returns a keypoint value for given keypoint index
func (p AnalyzedPose) KeyPointFor(index KeyPointIndex) (x, y, score float64) {
	i := int(index)
	idx1, idx2, idx3 := i*3, i*3+1, i*3+2

	count := len(p.KeyPoints)
	if idx1 < count && idx2 < count && idx3 < count {
		x = p.KeyPoints[idx1]
		y = p.KeyPoints[idx2]
		score = p.KeyPoints[idx3]
	}

	return x, y, score
}

// ResponseAnalyzedPoseFromVideoURLRequested struct
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-submit
type ResponseAnalyzedPoseFromVideoURLRequested struct {
	JobID string `json:"job_id"`
}

// PoseAnalysisStatus type
type PoseAnalysisStatus string

// PoseAnalysisStatus constants
const (
	PoseAnalysisStatusWaiting    PoseAnalysisStatus = "waiting"
	PoseAnalysisStatusProcessing PoseAnalysisStatus = "processing"
	PoseAnalysisStatusSuccess    PoseAnalysisStatus = "success"
	PoseAnalysisStatusFailed     PoseAnalysisStatus = "failed"
	PoseAnalysisStatusNotFound   PoseAnalysisStatus = "not found"
)

// PoseAnnotation struct
type PoseAnnotation struct {
	FrameNum int            `json:"frame_num"` // 0 ~ n-1
	Objects  []AnalyzedPose `json:"objects"`
}

// KeyPointCategory struct
type KeyPointCategory struct {
	ID            int      `json:"id"` // = 1 (for person)
	KeyPoints     []string `json:"keypoints"`
	Name          string   `json:"name"` // = "person"
	Skeleton      [][]int  `json:"skeleton"`
	SuperCategory string   `json:"supercategory"` // = "person"
}

// PoseInfo struct
type PoseInfo struct {
	Contributor string `json:"contributor"`
	DateCreated string `json:"date_created"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Version     string `json:"version"`
	Year        int    `json:"year"`
}

// PoseVideo struct
type PoseVideo struct {
	FPS    float32 `json:"fps"`
	Frames int     `json:"frames"`
	Height int     `json:"height"`
	Width  int     `json:"width"`
}

// ResponseAnalyzedPoseFromVideoURL struct
//
// https://developers.kakao.com/docs/latest/ko/pose/dev-guide#job-retrieval
type ResponseAnalyzedPoseFromVideoURL struct {
	Annotations []PoseAnnotation   `json:"annotations,omitempty"` // returns only when `status` == "success"
	Categories  []KeyPointCategory `json:"categories,omitempty"`  // returns only when `status` == "success"
	Info        PoseInfo           `json:"info,omitempty"`        // returns only when `status` == "success"
	JobID       string             `json:"job_id"`
	Status      PoseAnalysisStatus `json:"status"`
	Video       PoseVideo          `json:"video,omitempty"`       // returns only when `status` == "success"
	Description string             `json:"description,omitempty"` // returns only when `status` == "failed"
}
