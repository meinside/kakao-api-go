package kakaoapi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

// Created on 2019.03.25.
//
// Functions for [Speech APIs](https://developers.kakao.com/docs/latest/ko/voice/rest-api)
//
// SSML guide: https://developers.kakao.com/assets/guide/kakao_ssml_guide.pdf

// TypeVoice is type for voice types
type TypeVoice string

// Voice types
const (
	VoiceWomanReadCalm     TypeVoice = "WOMAN_READ_CALM"
	VoiceManReadCalm       TypeVoice = "MAN_READ_CALM"
	VoiceWomanDialogBright TypeVoice = "WOMAN_DIALOG_BRIGHT"
	VoiceManDialogBright   TypeVoice = "MAN_DIALOG_BRIGHT"
	VoiceDefault                     = VoiceWomanReadCalm
)

// TypeSpeechStyle is type for speech styles
type TypeSpeechStyle string

// Speech styles
const (
	StyleReadSpeech TypeSpeechStyle = "SS_READ_SPEECH" // default
	StyleAltFast    TypeSpeechStyle = "SS_ALT_FAST_1"
	StyleAltSlow    TypeSpeechStyle = "SS_ALT_SLOW_1"
	StyleDefault                    = StyleReadSpeech
)

// Speak struct for speeches
type Speak struct {
	XMLName xml.Name `xml:"speak"`

	Voices []Voice `xml:"voice"`
}

// Voice struct for speech
type Voice struct {
	XMLName xml.Name `xml:"voice"`

	Text        string          `xml:",chardata"`
	VoiceType   TypeVoice       `xml:"name,attr,omitempty"`
	SpeechStyle TypeSpeechStyle `xml:"speechStyle,attr,omitempty"`
}

// SpeechToText recognizes given speech (should be in mono, 16kHz, 16bit raw pcm format)
//
// https://developers.kakao.com/docs/latest/ko/voice/rest-api#speech-to-text
func (c *Client) SpeechToText(bs []byte) (result ResponseSpeechToText, err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", APINewtoneURL+"/v1/recognize", bytes.NewBuffer(bs)); err == nil {
		// set HTTP headers
		req.Header.Set("Content-Type", "application/octet-stream")
		req.Header.Set("Transfer-Encoding", "chunked")
		req.Header.Set("Authorization", c.authHeader(authTypeKakaoAK)) // set auth header

		var resp *http.Response
		resp, err = c.httpClient.Do(req)

		if resp != nil {
			defer resp.Body.Close()
		}
		if err == nil {
			if resp.StatusCode == 200 {
				if resp.Header.Get("Content-Type") == "multipart/form-data" {
					var reader *multipart.Reader
					if reader, err = resp.Request.MultipartReader(); err == nil {
						for {
							part, pErr := reader.NextPart()
							if pErr == io.EOF { // no more parts
								break
							}

							var bs []byte
							if _, err = part.Read(bs); err == nil {
								err = json.Unmarshal(bs, &result)
							}
						}

						if err == nil {
							if result.Type == "finalResult" { // final result
								return result, nil
							} else if result.Type == "errorCalled" { // on error
								return result, fmt.Errorf(result.Value)
							}

							err = fmt.Errorf("did not receive final result")
						}
					}
				} else {
					/*
						* 서버에서 잘못된 응답이 내려옴 (multipart/form-data가 아님):

						------newtoneFeldpXiEUcgOlZIe
						Content-Type: application/json; charset=UTF-8

						{"type":"beginPointDetection","value":"BPD"}
						------newtoneFeldpXiEUcgOlZIe
						Content-Type: application/json; charset=UTF-8

						{"type":"partialResult","value":"헤이"}
						------newtoneFeldpXiEUcgOlZIe
						Content-Type: application/json; charset=UTF-8

						{"type":"partialResult","value":"헤이 카카오"}
						------newtoneFeldpXiEUcgOlZIe
						Content-Type: application/json; charset=UTF-8

						{"type":"endPointDetection","value":"EPD"}
						------newtoneFeldpXiEUcgOlZIe
						Content-Type: application/json; charset=UTF-8
						Speech-Length: 6

						{"type":"finalResult","value":"헤이 카카오","nBest":[{"value":"헤이 카카오","score":24}]}
						------newtoneFeldpXiEUcgOlZIe--
					*/
					var body []byte
					if body, err = ioutil.ReadAll(resp.Body); err == nil {
						// XXX - parse it manually...
						lines := strings.Split(string(body), "\n")
						boundary := strings.TrimSpace(lines[0])

						for _, line := range lines {
							line = strings.TrimSpace(line)

							// skip unneeded lines
							if strings.HasPrefix(line, boundary) ||
								strings.HasPrefix(line, "Content-Type") ||
								strings.HasPrefix(line, "Speech-Length") ||
								len(line) <= 0 {
								continue
							}

							err = json.Unmarshal([]byte(line), &result)
						}

						if err == nil {
							if result.Type == "finalResult" { // final result
								return result, nil
							} else if result.Type == "errorCalled" { // on error
								return result, fmt.Errorf(result.Value)
							}

							err = fmt.Errorf("did not receive final result")
						}

						log.Printf("response is not multipart/forma-data:\n%s", string(body))
					}
				}
			} else {
				err = fmt.Errorf("HTTP status %d", resp.StatusCode)
			}
		}
	}

	return ResponseSpeechToText{}, err
}

// TextToSpeech synthesizes a .mp3 file with given SSML(xml)
//
// https://developers.kakao.com/docs/latest/ko/voice/rest-api#text-to-speech
func (c *Client) TextToSpeech(ssml []byte) (file []byte, contentType string, err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", APINewtoneURL+"/v1/synthesize", bytes.NewBuffer(ssml)); err == nil {
		// set HTTP headers
		req.Header.Set("Content-Type", "application/xml")
		req.Header.Set("Authorization", c.authHeader(authTypeKakaoAK)) // set auth header

		var res []byte
		if res, err = c.fetchHTTPResponse(req); err == nil {
			return res, http.DetectContentType(res), nil
		}
	}

	return nil, "application/octet-stream", err
}
