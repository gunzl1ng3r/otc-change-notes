package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func identifyChangeHistory(url string) bool {
	// fmt.Println("The provided url is:", url)
	resp, err := http.Get(url)

	if err != nil {
		log.Println("ERROR - Message is", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR - Message is", err)
	}

	bodyAsSlice := strings.Split(string(body), "\n")
	for _, element := range bodyAsSlice {
		// The title "Change History" indicates this page contains what we want.
		if strings.HasPrefix(element, "    <title>Change History</title>") {
			// fmt.Println("Found string \"Change History\"")
			return true
		}
	}
	return false
}
