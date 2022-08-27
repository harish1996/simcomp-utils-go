package httputil

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var auth_req *HelperClient

// var noauth *HelperClient

func AuthFetchFromJSONData(url string, out interface{}) error {

	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		return fmt.Errorf("out should be a pointer type.")
	}

	/*
		Did I suddenly turn into a new leaf and decide to declare this err explicitly ?

		Nope. Apparently since auth_req is a global variable, and when we use a global variable with :=
		it assumes that it needs to create a new local variable with the same name, eventhough we have used
		it for the other variable ( err in this case. )

		Since now there is a local variable called auth_req, It will just go ahead and override the global
		variable ( Since that is the default behaviour in Go. ). Also there seems to be no way to override this
		behaviour like in C++.

		https://developmentality.wordpress.com/2014/03/03/go-gotcha-1-variable-shadowing-within-inner-scope-due-to-use-of-operator/
		https://stackoverflow.com/q/47624326
	*/
	var err error

	if auth_req == nil {
		auth_req, err = GetAuthenticatedSession()
		if err != nil {
			auth_req = nil
			return fmt.Errorf("Authenticated session creation failed: %w", err)
		}
	}

	resp, err := auth_req.Httpget(url, nil)
	if err != nil {
		return fmt.Errorf("HTTP get failed: %w", err)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(out)
	if err != nil {
		return fmt.Errorf("Readall failed \n %w", err)
	}

	return nil
}
