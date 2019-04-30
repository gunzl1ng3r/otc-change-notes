package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("https://docs.otc.t-systems.com/en-us/usermanual/obs/en-us_topic_0071293550.html")

	if err != nil {
		log.Println("ERROR - Message is", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR - Message is", err)
	}

	bodyAsSlice := strings.Split(string(body), "\n")
	for index, element := range bodyAsSlice {
		output := stripHTML(element)
		fmt.Println(index, output)
	}
	// fmt.Println(string(body))
}

func stripHTML(html string) string {
	re := regexp.MustCompile(`<.+?>`)
	return re.ReplaceAllString(html, "")
}
