package kakaoapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Created on 2018.05.09.

// Constants
const (
	APIBaseURL = "https://kapi.kakao.com"
)

// Client struct
type Client struct {
	apiKey string

	Verbose bool // log verbose message or not
}

// NewClient returns a new API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		Verbose: false,
	}
}

// HTTP functions

// HTTP GET
func (c *Client) get(apiURL string, authType authType, headers map[string]string, params map[string]interface{}) ([]byte, error) {
	httpClient := &http.Client{}

	var err error
	var req *http.Request
	if req, err = http.NewRequest("GET", apiURL, nil); err == nil {
		// set HTTP headers
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		req.Header.Set("Authorization", c.authHeader(authType)) // set auth header

		// set parameters
		queries := req.URL.Query()
		for key, value := range params {
			queries.Add(key, fmt.Sprintf("%v", value))
		}
		req.URL.RawQuery = queries.Encode()

		return c.fetchHTTPResponse(httpClient, req)
	}

	return []byte{}, err
}

// HTTP POST
func (c *Client) post(apiURL string, authType authType, headers map[string]string, params map[string]interface{}) ([]byte, error) {
	httpClient := &http.Client{}

	var err error

	if hasFileInParams(params) {
		// multipart/form-data

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for key, value := range params {
			switch value.(type) {
			case fileParam:
				file, _ := value.(fileParam)
				filename := fmt.Sprintf("%s.%s", key, getExtension(file.bytes))

				if part, err := writer.CreateFormFile(key, filename); err == nil {
					if _, err := io.Copy(part, bytes.NewReader(file.bytes)); err != nil {
						log.Printf("* Could not write bytes to multipart for param '%s': %s", key, err)
					}
				} else {
					log.Printf("* Could not create part for param '%s': %s", key, err)
				}
			default:
				writer.WriteField(key, fmt.Sprintf("%v", value))
			}
		}

		if err := writer.Close(); err != nil {
			log.Printf("* Error while closing multipart form data writer: %s", err)
		}

		var req *http.Request
		if req, err = http.NewRequest("POST", apiURL, body); err == nil {
			// set HTTP headers
			for k, v := range headers {
				req.Header.Set(k, v)
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", c.authHeader(authType)) // set auth header

			return c.fetchHTTPResponse(httpClient, req)
		}
	} else {
		// application/x-www-form-urlencoded

		// parameters
		data := url.Values{}
		for k, v := range params {
			data.Set(k, fmt.Sprintf("%v", v))
		}

		var req *http.Request
		if req, err = http.NewRequest("POST", apiURL, strings.NewReader(data.Encode())); err == nil {
			// set HTTP headers
			for k, v := range headers {
				req.Header.Set(k, v)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Authorization", c.authHeader(authType)) // set auth header

			return c.fetchHTTPResponse(httpClient, req)
		}
	}

	return []byte{}, err
}

func (c *Client) fetchHTTPResponse(httpClient *http.Client, req *http.Request) (response []byte, err error) {
	// verbose message for debugging
	if c.Verbose {
		if dumped, err := httputil.DumpRequest(req, true); err == nil {
			log.Printf(`>>>>>> Request dump of %s %s:
%s
----------------`,
				req.Method,
				req.URL.Path,
				string(dumped),
			)
		}
	}

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err == nil {
		// verbose message for debugging
		if c.Verbose {
			if dumped, err := httputil.DumpResponse(resp, true); err == nil {
				log.Printf(`>>>>>> Response dump of %s %s:
%s
----------------`,
					req.Method,
					req.URL.Path,
					string(dumped),
				)
			}
		}

		var bytes []byte
		if bytes, err = ioutil.ReadAll(resp.Body); err == nil {
			if resp.StatusCode == 200 {
				return bytes, nil
			}

			var errResponse ResponseError
			if err := json.Unmarshal(bytes, &errResponse); err == nil {
				return bytes, fmt.Errorf("API error with response code: %d, message: %s", errResponse.Code, errResponse.Message)
			}

			return bytes, fmt.Errorf("HTTP status %d %s", resp.StatusCode, string(bytes))
		} else if c.Verbose {
			// verbose message for debugging
			log.Printf(`****** Error on %s %s request:
%s
----------------`,
				req.Method,
				req.URL.Path,
				string(response),
			)
		}

	}

	return []byte{}, err
}

func (c *Client) authHeader(method authType) string {
	return fmt.Sprintf("%s %s", method, c.apiKey)
}

// checks if given `params` has any fileParam in it
func hasFileInParams(params map[string]interface{}) bool {
	for _, v := range params {
		if _, ok := v.(fileParam); ok {
			return true
		}
	}

	return false
}

// get file extension from bytes array
//
// https://www.w3.org/Protocols/rfc1341/4_Content-Type.html
func getExtension(bytes []byte) string {
	types := strings.Split(http.DetectContentType(bytes), "/") // ex: "image/jpeg"
	if len(types) >= 2 {
		splitted := strings.Split(types[1], ";") // for removing subtype parameter
		if len(splitted) >= 1 {
			return splitted[0] // return subtype only
		}
	}
	return "" // default
}
