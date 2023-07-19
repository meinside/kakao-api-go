package kakaoapi

import (
	"log"
	"os"
	"testing"
)

func isVerbose() bool {
	verbose := os.Getenv("VERBOSE")

	return verbose == "true"
}

func TestKoGPT(t *testing.T) {
	_apiKey := os.Getenv("KAKAO_API_KEY")
	_verbose := isVerbose()

	client := NewClient(_apiKey)
	client.Verbose = _verbose

	if len(_apiKey) <= 0 {
		t.Errorf("environment variable `KAKAO_API_KEY` is needed")
	}

	params := NewParamsTextGeneration("오늘 아침 하늘은 곧 비가 올 것 같아서", 120).
		SetN(2)

	if generated, err := client.GenerateTexts(params); err != nil {
		t.Errorf("failed to generate texts: %s", err)
	} else {
		if len(generated.Generations) != 2 {
			t.Errorf("count of generated texts is different from request: %d", len(generated.Generations))
		}

		if _verbose {
			log.Printf("generated texts = %+v", generated)
		}
	}
}

func TestKarlo(t *testing.T) {
	_apiKey := os.Getenv("KAKAO_API_KEY")
	_verbose := isVerbose()

	client := NewClient(_apiKey)
	client.Verbose = _verbose

	if len(_apiKey) <= 0 {
		t.Errorf("environment variable `KAKAO_API_KEY` is needed")
	}

	// generation
	paramsGeneration := NewParamsImageGeneration("A cat with white fur").
		SetNegativePrompt("sleeping cat, dog, human, ugly face, cropped").
		SetImageFormat(ImageFormatJPEG).
		SetReturnType(ImageReturnURL).
		SetNSFWChecker(true).
		SetSamples(2)

	if generated, err := client.GenerateImages(paramsGeneration); err != nil {
		t.Errorf("failed to generate images: %s", err)
	} else {
		if len(generated.Images) != 2 {
			t.Errorf("count of generated images is different from request: %d", len(generated.Images))
		}

		if _verbose {
			log.Printf("generated images = %+v", generated)
		}
	}

	imageBytes, err := os.ReadFile("./sample/image.jpg")
	if err != nil {
		t.Fatalf("failed to read sample image file: %s", err)
	}
	image := EncodeBase64(imageBytes)

	// upscale
	paramsUpscale := NewParamsImageUpscale([]string{image}).
		SetImageFormat(ImageFormatPNG).
		SetReturnType(ImageReturnURL)
	if scaledUp, err := client.UpscaleImages(paramsUpscale); err != nil {
		t.Errorf("failed to upscale image: %s", err)
	} else {
		if _verbose {
			log.Printf("upscaled images = %+v", scaledUp)
		}
	}

	// variation
	paramsVariation := NewParamsImageVariation(image, "make it looking modern, futuristic, and avant-garde").
		SetImageFormat(ImageFormatPNG).
		SetReturnType(ImageReturnURL)
	if varied, err := client.VaryImage(paramsVariation); err != nil {
		t.Errorf("failed to vary image: %s", err)
	} else {
		if _verbose {
			log.Printf("varied images = %+v", varied)
		}
	}

	// check NSFW
	if nsfw, err := client.CheckNSFW([]string{image}); err != nil {
		t.Errorf("failed to check NSFW: %s", err)
	} else {
		if _verbose {
			log.Printf("NSFW result = %+v", nsfw)
		}
	}
}
