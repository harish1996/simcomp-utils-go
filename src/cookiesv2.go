package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	hu "hg33.com/simcomp/utils"
)

func check_error(e error) {
	if e != nil {
		e = fmt.Errorf("%#v", e)
		panic(e)
	}
}

func getPwd(filename string) string {
	f, err := os.Open(filename)
	check_error(err)

	contents, err := ioutil.ReadAll(f)
	check_error(err)

	pwd := strings.Trim(string(contents), " \n\t")
	return pwd
}

func ResponseHandler(response *http.Response, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func getAuthenticatedSession() http.Client {

// }

func main() {

	headers := map[string]string{
		`Referer`:    `https://www.simcompanies.com/`,
		`Connection`: `keep-alive`,
		`User-agent`: `Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0`,
	}

	data := map[string]string{
		"email":           "harishganesan96@gmail.com",
		"timezone_offset": "-330",
		/* TODO: Get the path of the password file from the input */
		"password": getPwd("simcomp.password"),
	}

	client := hu.NewHelperClient()

	response, err := client.Httpget("https://www.simcompanies.com/", nil, nil, nil)
	ResponseHandler(response, err)
	if response.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "CSRF cookie fetch failed. Status code is ", response.StatusCode)
		os.Exit(1)
	}

	cookies := response.Cookies()
	for _, v := range cookies {
		if v.Name == "csrftoken" {
			headers["X-CSRFToken"] = v.Value
		}
	}

	response, err = client.Httppost("https://www.simcompanies.com/api/v2/auth/email/auth/", headers, data, nil)

	ResponseHandler(response, err)
	if response.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "Session ID fetch failure. Response code= ", response.StatusCode)
		os.Exit(1)
	}

	outfile, err := os.OpenFile("go_simcomp_cookie.json", os.O_CREATE|os.O_WRONLY, 0644)
	check_error(err)

	cookiejson, err := json.MarshalIndent(response.Cookies(), "", "\t")
	written_n, err := outfile.Write(cookiejson)
	check_error(err)

	fmt.Println("Written ", written_n, " bytes into go_simcomp_cookie.json")
}
