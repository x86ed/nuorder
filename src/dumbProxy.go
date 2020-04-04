package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func githubProxy() (string, error) {

	url := "https://api.github.com/repos/facebook/react/issues"

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}

func hub(w http.ResponseWriter, req *http.Request) {
	body, err := githubProxy()
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		fmt.Fprint(w, body)
	}
	w.WriteHeader(500)
	log.Printf("{\"error\":, %v}", err)
}

func main() {

	// We register our handlers on server routes using the
	// `http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/http` package and
	// takes a function as an argument.
	http.HandleFunc("/hub", hub)

	http.ListenAndServe(":1337", nil)
}
