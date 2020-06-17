package main

import (
	"encoding/json"
	"fmt"
	"log"

	kakaoapi "github.com/meinside/kakao-api-go"
)

const (
	//apiKey = "0123456789abcdefghijklmnopqrstuvwxyz" // XXX - change this to your REST API key
	apiKey = "1783501ee04a097a01c26e1d49564f4a" // XXX - change this to your REST API key

	poseFilepath = "./pose.jpg"

	// https://github.com/intel-iot-devkit/sample-videos
	poseVideoURL = "https://raw.githubusercontent.com/intel-iot-devkit/sample-videos/master/face-demographics-walking.mp4"
)

func main() {
	client := kakaoapi.NewClient(apiKey)
	//client.Verbose = true

	// analyze pose from an image
	if pose, err := client.AnalyzePoseFromImageFilepath(poseFilepath); err == nil {
		log.Printf("Pose: %s", prettify(pose))
	} else {
		log.Printf("Failed to analyze pose: %s", err)
	}

	// analyze pose from a video url
	// TODO: not tested (only for affiliated developers)
	/*
		if requested, err := client.AnalyzePoseFromVideoURL(poseVideoURL, true, ""); err == nil {
			log.Printf("Requested: %s", prettify(requested))

			log.Printf("Wait for a while...")
			time.Sleep(5 * time.Second)

			if retrieved, err := client.RetrieveAnalyzedPoseFromVideoURL(requested.JobID); err == nil {
				if retrieved.Status == kakaoapi.PoseAnalysisStatusSuccess {
					log.Printf("Retrieved: %s", prettify(retrieved))
				}
			} else {
				log.Printf("Failed to retrieve analysis result: %s", prettify(err))
			}
		} else {
			log.Printf("Failed to request analysis: %s", err)
		}
	*/
}

func prettify(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		return string(bytes)
	}

	return fmt.Sprintf("%v", obj)
}
