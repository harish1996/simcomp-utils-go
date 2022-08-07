package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type stringmap map[string]string

type HelperClient struct {
	http.Client
}

func NewHelperClient() *HelperClient {
	jar, _ := cookiejar.New(nil)
	return &HelperClient{Client: http.Client{Jar: jar}}
}

func getHeader(dict stringmap) http.Header {
	header := make(http.Header)
	for k, v := range dict {
		header.Add(k, v)
	}
	return header
}

func fillJar(jar http.CookieJar, cookies stringmap, link string) {

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

func (h *HelperClient) httpRequest(request_type string, link string, header stringmap, data interface{}, cookies stringmap) (*http.Response, error) {
	var data_buf bytes.Buffer
	var headerobj http.Header

	if header != nil {
		headerobj = getHeader(header)
	}

	if data != nil {
		databytes, err := json.Marshal(data)
		if err != nil {
			err = fmt.Errorf("Json marshalling failed during request %s: %w", request_type, err)
			return nil, err
		}
		/* Have to use the dereferenced version, instead of the pointer directly, because if data_buf is nil,
		http.NewRequest is still trying to read from the object, which will cause a segmentation fault */
		data_buf = *bytes.NewBuffer(databytes)
	}

	if cookies != nil {
		fillJar(h.Jar, cookies, link)
	}

	req, err := http.NewRequest(request_type, link, &data_buf)
	if err != nil {
		err = fmt.Errorf("Request creation fails during %s: %w", request_type, err)
		return nil, err
	}
	req.Header = headerobj

	return h.Do(req)
}

func (h *HelperClient) Httpget(link string, header stringmap, data interface{}, cookies stringmap) (*http.Response, error) {
	return h.httpRequest("GET", link, header, data, cookies)
}

func (h *HelperClient) Httppost(link string, header stringmap, data interface{}, cookies stringmap) (*http.Response, error) {
	return h.httpRequest("POST", link, header, data, cookies)
}
