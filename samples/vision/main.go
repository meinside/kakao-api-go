package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/meinside/kakao-api-go"
)

const (
	apiKey = "0123456789abcdefghijklmnopqrstuvwxyz" // XXX - change this to your REST API key

	facesFilepath = "./faces.jpg"
	nsfwFilepath  = "./nsfw.jpg"
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

	// detect nsfw
	if nsfw, err := client.DetectNSFWFromFilepath(nsfwFilepath); err == nil {
		log.Printf("NSFW: %+v", prettify(nsfw))
	} else {
		log.Printf("Failed to detect NSFW: %s", err)
	}
}

func prettify(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		return string(bytes)
	}

	return fmt.Sprintf("%v", obj)
}
