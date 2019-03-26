package main

import (
	"io/ioutil"
	"log"

	kakaoapi "github.com/meinside/kakao-api-go"
)

/*
	$ ffmpeg -i original.mp3 -acodec pcm_s16le -ac 1 -ar 16000 converted.wav
*/

const (
	apiKey = "0123456789abcdefghijklmnopqrstuvwxyz" // XXX - change this to your REST API key

	sampleFilename      = "heykakao.wav"
	synthesizedFilename = "synthesized.mp3"
)

func main() {
	client := kakaoapi.NewClient(apiKey)
	//client.Verbose = true

	// recognize voices
	if bs, err := ioutil.ReadFile(sampleFilename); err == nil {
		if res, err := client.SpeechToText(bs, kakaoapi.ServiceModeDictation); err == nil {
			log.Printf("Recognized speech: %+v", res)
		} else {
			log.Printf("Failed to recognize speech: %s", err)
		}
	} else {
		log.Printf("Failed to read sample speech: %s", err)
	}

	// synthesize voices
	if bs, contentType, err := client.TextToSpeech(kakaoapi.Speak{
		Voices: []kakaoapi.Voice{
			kakaoapi.Voice{
				Text:        "헤이 카카오!",
				VoiceType:   kakaoapi.VoiceDefault,
				SpeechStyle: kakaoapi.StyleDefault,
			},
			kakaoapi.Voice{
				Text:        "니가 가라 하와이.",
				VoiceType:   kakaoapi.VoiceManDialogBright,
				SpeechStyle: kakaoapi.StyleAltSlow,
			},
			kakaoapi.Voice{
				Text:        "고니는 제가 아는 타짜 중에 최고였어요.",
				VoiceType:   kakaoapi.VoiceWomanReadCalm,
				SpeechStyle: kakaoapi.StyleAltFast,
			},
		}}); err == nil {
		if err := ioutil.WriteFile(synthesizedFilename, bs, 0644); err == nil {
			log.Printf("Synthesized voice saved to: %s (%d bytes, %s)", synthesizedFilename, len(bs), contentType)
		} else {
			log.Printf("Failed to save synthesized voice: %s", err)
		}
	} else {
		log.Printf("Failed to synthesize voice: %s", err)
	}
}
