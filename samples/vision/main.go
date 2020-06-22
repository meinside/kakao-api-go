package main

import (
	"encoding/json"
	"fmt"
	"log"

	kakaoapi "github.com/meinside/kakao-api-go"
)

const (
	apiKey = "0123456789abcdefghijklmnopqrstuvwxyz" // XXX - change this to your REST API key

	facesFilepath = "./faces.jpg"
	nsfwFilepath  = "./nsfw.jpg"
	textFilepath  = "./text.jpg"
)

func main() {
	client := kakaoapi.NewClient(apiKey)
	//client.Verbose = true

	// detect faces
	if faces, err := client.DetectFaceFromFilepath(facesFilepath, 0.6); err == nil {
		log.Printf("Faces: %s", prettify(faces))
	} else {
		log.Printf("Failed to detect faces: %s", err)
	}

	// detect nsfw
	if nsfw, err := client.DetectNSFWFromFilepath(nsfwFilepath); err == nil {
		log.Printf("NSFW: %+v", prettify(nsfw))
	} else {
		log.Printf("Failed to detect NSFW: %s", err)
	}

	// detect products
	if products, err := client.DetectProductFromFilepath(facesFilepath, 0.6); err == nil {
		log.Printf("Products: %s", prettify(products))
	} else {
		log.Printf("Failed to detect products: %s", err)
	}

	// crop thumbnail
	if cropped, err := client.CropThumbnailFromFilepath(facesFilepath, 200, 200); err == nil {
		log.Printf("Cropped: %s", prettify(cropped))
	} else {
		log.Printf("Failed to crop thumbnail: %s", err)
	}

	// suggest thumbnail
	if suggested, err := client.SuggestThumbnailFromFilepath(facesFilepath, 200, 200); err == nil {
		log.Printf("Suggested thumbnail: %s", prettify(suggested))
	} else {
		log.Printf("Failed to suggest thumbnail: %s", err)
	}

	// generate tags
	if tags, err := client.GenerateTagsFromFilepath(facesFilepath); err == nil {
		log.Printf("Generated tags: %s", prettify(tags))
	} else {
		log.Printf("Failed to generate tags: %s", err)
	}

	// detect text
	if text, err := client.DetectTextFromFilepath(textFilepath); err == nil {
		log.Printf("Detected text: %s", prettify(text))

		// and then recognize it
		if recognized, err := client.RecognizeTextFromFilepath(textFilepath, text.Result.Boxes); err == nil {
			log.Printf("Recognized text: %s", prettify(recognized))
		} else {
			log.Printf("Failed to recognize text: %s", err)
		}
	} else {
		log.Printf("Failed to detect text: %s", err)
	}
}

func prettify(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		return string(bytes)
	}

	return fmt.Sprintf("%v", obj)
}
