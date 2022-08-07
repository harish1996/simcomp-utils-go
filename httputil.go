package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

type stringmap map[string]string

type HelperClient http.Client

func ResponseHandler(response *http.Response, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewHelperClient() *HelperClient {
	jar, _ := cookiejar.New(nil)
	return &http.Client{Jar: jar}
}

func getHeader(dict stringmap) http.Header {
	header := make(http.Header)
	for k, v := range dict {
		header.Add(k, v)
	}
	return header
}

func fillJar(jar *http.CookieJar, cookies stringmap, link string) {

	if jar == nil {
		jar, _ = cookiejar.New(nil)
	}

	linkobj, _ := url.Parse(link)

	var cs []*http.Cookie

	for k, v := range cookies {
		// fmt.Println(k, v)
		cs = append(cs, &http.Cookie{Name: k, Value: v})
	}

	/* Check if Setting cookies, removes the cookies that are already present in the jar */
	jar.SetCookies(linkobj, cs)

}

func process_headers(header stringmap) http.Header {

}

func process_data(date []byte) (bytes.Buffer, error) {
	var ret bytes.Buffer
	if data != nil {
		databytes, err := json.Marshal(data)
		check_error(err)
		data_buf = bytes.NewBuffer(databytes)
	}
}
func (h *HelperClient) Httpget(link string, header stringmap, data []byte, cookies stringmap) (http.Request, error) {
	var headerObj http.Header
	var data_buf bytes.Buffer

	if header != nil {
		headerObj = getHeader(header)
	}
	if data != nil {
		databytes, err := json.Marshal(data)
		check_error(err)
		data_buf = bytes.NewBuffer(databytes)
	}
	if cookies != nil {
		fillJar(&h.Jar, cookies, link)
	}

	req, err := http.NewRequest("GET", link, data_buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Request creation failed")
		return nil, err
	}
	req.Header = headerObj

	response, err := h.Do(req)
	return response, err
}

func (h *HelperClient) Httppost(url string, header stringmap, data []byte, cookies stringmap) (http.Request, error) {

}
