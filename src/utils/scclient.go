package httputil

// package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type UnreadableCookie struct {
	filename string
}
type NoCookieFile struct {
	filename string
}
type CorruptedCookie struct {
	filename string
	err      error
}

func (e *UnreadableCookie) Error() string {
	return fmt.Sprintf("Cookie file  %s is unreadable", e.filename)
}
func (e *NoCookieFile) Error() string {
	return fmt.Sprintf("No file named %s to read.", e.filename)
}
func (e *NoCookieFile) Is(target error) bool {
	_, ok := target.(*NoCookieFile)
	return ok
}
func (e *CorruptedCookie) Error() string {
	return fmt.Sprintf("Unable to decode cookie file %s due \n %w", e.filename, e.err)
}
func (e *CorruptedCookie) Unwrap() error {
	return e.err
}

var Defaultheaders = map[string]string{
	`Referer`:    `https://www.simcompanies.com/`,
	`Connection`: `keep-alive`,
	`User-agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
}

func ReadExistingCookies() ([]*http.Cookie, error) {

	ret := make([]*http.Cookie, 2)
	fname := "go_simcomp_cookie.json"
	fp, err := os.Open(fname)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return nil, &UnreadableCookie{filename: fname}
		} else if errors.Is(err, os.ErrNotExist) {
			return nil, &NoCookieFile{filename: fname}
		} else {
			err = fmt.Errorf("Reading of file %s failed due to %w", fname, err)
			return nil, err
		}
	}

	decoder := json.NewDecoder(fp)
	err = decoder.Decode(&ret)
	if err != nil {
		return nil, &CorruptedCookie{filename: fp.Name(), err: err}
	}

	return ret, nil
}

func checkExpired(cookie http.Cookie) bool {

	if cookie.Expires.After(time.Now()) {
		return false
	}
	return true
}

func getPwd(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cant open file with name %s \n %w", f.Name(), err)
		os.Exit(1)
	}

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Cant read file with name %s \n %w", f.Name(), err)
		os.Exit(1)
	}

	pwd := strings.Trim(string(contents), " \n\t")
	return pwd
}

func GetNewCookies() ([]*http.Cookie, error) {

	headers := Defaultheaders

	data := map[string]string{
		/* Shift the email to a config file */
		"email":           "harishganesan96@gmail.com",
		"timezone_offset": "-330",
		/* TODO: Get the path of the password file from the input */
		"password": getPwd("simcomp.password"),
	}

	client := NewHelperClient()

	response, err := client.Httpget("https://www.simcompanies.com/", nil, nil, nil)
	if err != nil {
		err = fmt.Errorf("HTTP GET during New cookie fetch failed due to \n %w", err)
		return nil, err
	}
	if response.StatusCode != 200 {
		err = fmt.Errorf("CSRF cookie fetch failed. Status code is %d", response.StatusCode)
		return nil, err
	}

	cookies := response.Cookies()
	for _, v := range cookies {
		if v.Name == "csrftoken" {
			headers["X-CSRFToken"] = v.Value
		}
	}

	response, err = client.Httppost("https://www.simcompanies.com/api/v2/auth/email/auth/", headers, data, nil)
	if err != nil {
		err = fmt.Errorf("HTTP POST during New cookie fetch failed due to \n %w", err)
		return nil, err
	}
	if response.StatusCode != 200 {
		err = fmt.Errorf("Session ID fetch failure. Response code= %d", response.StatusCode)
		return nil, err
	}

	outfile, err := os.OpenFile("go_simcomp_cookie.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		err = fmt.Errorf("Error while writing to file go_simcomp_cookie.json during New Cookie fetch \n %w", err)
		return nil, err
	}

	cookiejson, err := json.MarshalIndent(response.Cookies(), "", "\t")
	if err != nil {
		err = fmt.Errorf("Error while marshalling cookies. \n %w", err)
		return nil, err
	}

	_, err = outfile.Write(cookiejson)
	if err != nil {
		err = fmt.Errorf("Error while writing to file. %s \n %w", outfile.Name(), err)
		return nil, err
	}

	return response.Cookies(), nil

}

func GetAuthenticatedSession() (*HelperClient, error) {

	var get_new_cookie bool = false
	cookies, err := ReadExistingCookies()
	if err != nil {
		if errors.Is(err, &NoCookieFile{}) {
			get_new_cookie = true
		} else {
			err = fmt.Errorf("Getting Authenticated Session failed \n %w", err)
			return nil, err
		}
	}

	if get_new_cookie == false {
		for _, v := range cookies {
			if checkExpired(*v) {
				get_new_cookie = true
				break
			}
		}
	}

	if get_new_cookie {
		fmt.Println("Getting new cookies either because there are no existing cookies or because existing ones have expired...")
		cookies, err = GetNewCookies()
	}

	ur, err := url.Parse("https://www.simcompanies.com/")
	if err != nil {
		err = fmt.Errorf("URL parse failed \n %w", err)
		return nil, err
	}

	client := NewHelperClient()
	client.Jar.SetCookies(ur, cookies)

	return client, nil
}
