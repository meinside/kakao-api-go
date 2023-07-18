package kakaoapi

import "os"

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
	if bytes, err = os.ReadFile(path); err == nil {
		return fileParam{
			bytes: bytes,
		}, nil
	}

	return fileParam{}, err
}

///////////////////////////////
// common structs
//

// ResponseError is the struct for API error responses
type ResponseError struct {
	Code int `json:"code,omitempty"`

	ErrorType string `json:"errorType,omitempty"`

	Msg     string `json:"msg,omitempty"`
	Message string `json:"message,omitempty"`
}

///////////////////////////////
// API request & response structs
//

type ParamsTextGeneration map[string]any

// NewParamsTextGeneration creates a new ParamsTextGeneration.
func NewParamsTextGeneration(prompt string, maxTokens int) ParamsTextGeneration {
	return ParamsTextGeneration{
		"prompt":     prompt,
		"max_tokens": maxTokens,
	}
}

// SetTemp sets the temperature of ParamsTextGeneration.
func (p ParamsTextGeneration) SetTemp(temp float64) ParamsTextGeneration {
	p["temperature"] = temp
	return p
}

// SetTopP sets the top_p of ParamsTextGeneration.
func (p ParamsTextGeneration) SetTopP(topP float64) ParamsTextGeneration {
	p["top_p"] = topP
	return p
}

// SetN sets the n of ParamsTextGeneration.
func (p ParamsTextGeneration) SetN(n int) ParamsTextGeneration {
	p["n"] = n
	return p
}

// ResponseGeneratedTexts is the struct for generated texts
type ResponseGeneratedTexts struct {
	ID          string `json:"id"`
	Generations []struct {
		Text   string `json:"text"`
		Tokens int    `json:"tokens"`
	} `json:"generations"`
	Usage struct {
		PromptTokens    int `json:"prompt_tokens"`
		GeneratedTokens int `json:"generated_tokens"`
		TotalTokens     int `json:"total_tokens"`
	} `json:"usage"`
}

type ImageFormat string

const (
	ImageFormatWEBP ImageFormat = "webp"
	ImageFormatJPEG ImageFormat = "jpeg"
	ImageFormatPNG  ImageFormat = "png"
)

type ImageReturnType string

const (
	ImageReturnURL    ImageReturnType = "url"
	ImageReturnBase64 ImageReturnType = "base64_string"
)

type ImageDecodeScheduler string

const (
	ImageDecodeSchedulerDDIM ImageDecodeScheduler = "decoder_ddim_v_prediction"
	ImageDecodeSchedulerDDPM ImageDecodeScheduler = "decoder_ddpm_v_prediction"
)

type ParamsImageGeneration map[string]any

// NewParamsImageGeneration creates a new ParamsImageGeneration.
func NewParamsImageGeneration(prompt string) ParamsImageGeneration {
	return ParamsImageGeneration{
		"prompt": prompt,
	}
}

// SetNegativePrompt sets the negative prompt of ParamsImageGeneration.
func (p ParamsImageGeneration) SetNegativePrompt(negativePrompt string) ParamsImageGeneration {
	p["negative_prompt"] = negativePrompt
	return p
}

// SetWidth sets the width of ParamsImageGeneration.
func (p ParamsImageGeneration) SetWidth(width int) ParamsImageGeneration {
	p["width"] = width
	return p
}

// SetHeight sets the height of ParamsImageGeneration.
func (p ParamsImageGeneration) SetHeight(height int) ParamsImageGeneration {
	p["height"] = height
	return p
}

// SetUpscale sets the upscale of ParamsImageGeneration.
func (p ParamsImageGeneration) SetUpscale(upscale bool) ParamsImageGeneration {
	p["upscale"] = upscale
	return p
}

// SetScale sets the scale of ParamsImageGeneration.
func (p ParamsImageGeneration) SetScale(scale int) ParamsImageGeneration {
	p["scale"] = scale
	return p
}

// SetImageFormat sets the image format of ParamsImageGeneration.
func (p ParamsImageGeneration) SetImageFormat(format ImageFormat) ParamsImageGeneration {
	p["image_format"] = format
	return p
}

// SetImageQuality sets the image quality of ParamsImageGeneration.
func (p ParamsImageGeneration) SetImageQuality(quality int) ParamsImageGeneration {
	p["image_quality"] = quality
	return p
}

// SetSamples sets the samples of ParamsImageGeneration.
func (p ParamsImageGeneration) SetSamples(samples int) ParamsImageGeneration {
	p["samples"] = samples
	return p
}

// SetReturnType sets the return type of ParamsImageGeneration.
func (p ParamsImageGeneration) SetReturnType(returnType ImageReturnType) ParamsImageGeneration {
	p["return_type"] = returnType
	return p
}

// SetPriorNumInferenceSteps sets the prior num inference steps of ParamsImageGeneration.
func (p ParamsImageGeneration) SetPriorNumInferenceSteps(steps int) ParamsImageGeneration {
	p["prior_num_inference_steps"] = steps
	return p
}

// SetPriorGuidanceScale sets the prior guidance scale of ParamsImageGeneration.
func (p ParamsImageGeneration) SetPriorGuidanceScale(scale float64) ParamsImageGeneration {
	p["prior_guidance_scale"] = scale
	return p
}

// SetNumInferenceSteps sets the num inference steps of ParamsImageGeneration.
func (p ParamsImageGeneration) SetNumInferenceSteps(steps int) ParamsImageGeneration {
	p["num_inference_steps"] = steps
	return p
}

// SetGuidanceScale sets the guidance scale of ParamsImageGeneration.
func (p ParamsImageGeneration) SetGuidanceScale(scale float64) ParamsImageGeneration {
	p["guidance_scale"] = scale
	return p
}

// SetScheduler sets the scheduler of ParamsImageGeneration.
func (p ParamsImageGeneration) SetScheduler(scheduler ImageDecodeScheduler) ParamsImageGeneration {
	p["scheduler"] = scheduler
	return p
}

// SetSeed sets the seed of ParamsImageGeneration.
func (p ParamsImageGeneration) SetSeed(seed []int) ParamsImageGeneration {
	p["seed"] = seed
	return p
}

// SetNSFWChecker sets the NSFW checker of ParamsImageGeneration.
func (p ParamsImageGeneration) SetNSFWChecker(nsfwChecker bool) ParamsImageGeneration {
	p["nsfw_checker"] = nsfwChecker
	return p
}

// ResponseGeneratedImages is the struct for generated images
type ResponseGeneratedImages struct {
	ID           string `json:"id"`
	ModelVersion string `json:"model_version"`
	Images       []struct {
		ID                  string   `json:"id"`
		Seed                int64    `json:"seed"`
		Image               string   `json:"image"`
		NSFWContentDetected bool     `json:"nsfw_content_detected,omitempty"`
		NSFWScore           *float64 `json:"nsfw_score,omitempty"`
	} `json:"images"`
}
