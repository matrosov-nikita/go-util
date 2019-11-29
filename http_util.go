package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 20,
	// disable auto follow redirects
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

// GetHTTPClient returns http client with predefined settings such as timeout
func GetHTTPClient() *http.Client {
	return client
}

// DoHTTP makes http request, returns response and error. Error not nil if resp status not 2**
func DoHTTP(req *http.Request) (*http.Response, error) {
	return DoHTTPWithClient(GetHTTPClient(), req)
}

func DoHTTPString(req *http.Request) (string, error) {
	body, err := DoHTTPBytes(req)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func DoHTTPBytes(req *http.Request) ([]byte, error) {
	resp, err := DoHTTPWithClient(client, req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// DoHTTPWithClient makes http request, returns response and error. Error not nil if resp status not 2**
func DoHTTPWithClient(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DoHTTPJson makes http request and unmarshall its body to v. Returns error if something goes wrong
func DoHTTPJson(req *http.Request, v interface{}) error {
	return DoHTTPJsonWithClient(GetHTTPClient(), req, v)
}

// DoHTTPJsonWithClient same as DoHttpJson, but with passed http client
func DoHTTPJsonWithClient(client *http.Client, req *http.Request, v interface{}) error {
	resp, err := DoHTTPWithClient(client, req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		if resp == nil {
			return fmt.Errorf("error: '%s'", err.Error()) //nolint
		}
		return fmt.Errorf("error: '%s', status: %d", err.Error(), resp.StatusCode) //nolint
	}

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		var raw []byte
		raw, err = GetBodyRaw(resp, nil)
		return fmt.Errorf("error: '%s', status: %d. errorBody: '%s'", ErrMsgOrEmptyString(err), resp.StatusCode, string(raw))
	}
	if v != nil {
		if err = UnmarshalBody(resp, v); err != nil {
			return fmt.Errorf("error: '%s', status: %d", err.Error(), resp.StatusCode)
		}
	}
	return nil
}

// GetBodyRaw return body from response
func GetBodyRaw(resp *http.Response, err error) ([]byte, error) {
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// UnmarshalBody reads body of response and unmarshall its to v
func UnmarshalBody(resp *http.Response, v interface{}) error {
	body, err := GetBodyRaw(resp, nil)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

// WriteReq writes http request without body
func WriteReq(req *http.Request) []byte {
	b := &bytes.Buffer{}
	b.WriteString(fmt.Sprintf("%s %s HTTP/1.1\r\n", req.Method, req.URL.RequestURI()))
	for key, value := range req.Header {
		b.WriteString(fmt.Sprintf("%s: %s\r\n", key, value[0]))
	}

	b.WriteString("\r\n")

	return b.Bytes()
}

// WriteResp writes http response with body
func WriteResp(status int, text string) []byte {
	b := &bytes.Buffer{}

	b.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", status, http.StatusText(status)))
	b.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	b.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(text)))
	b.WriteString("\r\n")

	b.WriteString(text)
	return b.Bytes()
}

func NewResponseWriterDecorator(writer http.ResponseWriter) *ResponseWriterDecorator {
	return &ResponseWriterDecorator{writer: writer}
}

type ResponseWriterDecorator struct {
	writer http.ResponseWriter

	StatusCode int
}

func (d *ResponseWriterDecorator) Header() http.Header {
	return d.writer.Header()
}

func (d *ResponseWriterDecorator) Write(body []byte) (int, error) {
	return d.writer.Write(body)
}

func (d *ResponseWriterDecorator) WriteHeader(statusCode int) {
	d.StatusCode = statusCode
	d.writer.WriteHeader(statusCode)
}

// GetClientIP returns client ip from request
func GetClientIPByRequest(req http.Request) (ip string) {
	defer func() {
		if r := recover(); r != nil {
			ip = ""
		}
	}()
	clientIP := strings.TrimSpace(requestHeader(req, "X-Real-Ip"))
	if len(clientIP) > 0 {
		return clientIP
	}
	clientIP = requestHeader(req, "X-Forwarded-For")
	if index := strings.IndexByte(clientIP, ','); index >= 0 {
		clientIP = clientIP[0:index]
	}
	clientIP = strings.TrimSpace(clientIP)
	if len(clientIP) > 0 {
		return clientIP
	}
	return ip
}

func requestHeader(r http.Request, key string) string {
	if values := r.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}
