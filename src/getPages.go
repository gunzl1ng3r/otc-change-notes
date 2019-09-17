package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getPages() (SplitPageLinks map[string][]map[string]string) {

	// var PageLinks map[string][]string
	// PageLinks = make(map[string][]string)
	var PageLinks []string

	resp, err := http.Get("https://docs.otc.t-systems.com/")

	if err != nil {
		log.Println("ERROR - Message is", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("ERROR - Message is", err)
	}

	bodyAsSlice := strings.Split(string(body), "\n")

	// Loop over returned HTML, do not register index
	for _, element := range bodyAsSlice {
		// all the links are appended at the end of the page, so we look for the lines starting with a link
		if strings.HasPrefix(element, "<a href=") {
			// example: /en-us/usermanual/eip/faq_connect.html
			// remove the HTML part of the line
			_cleanedElement := strings.TrimSuffix(strings.TrimPrefix(element, "<a href=\""), "\" ></a>")
			// put all the elements in a nice little slice
			PageLinks = append(PageLinks, _cleanedElement)
		}
	}

	SplitPageLinks = make(map[string][]map[string]string)
	for _, element := range PageLinks {
		tempSplitPageLinks := strings.Split(element, "/")
		var _PageID string
		var _ServiceName string
		var _Prefix string

		_PageID, tempSplitPageLinks = tempSplitPageLinks[len(tempSplitPageLinks)-1], tempSplitPageLinks[:len(tempSplitPageLinks)-1]
		// fmt.Println("_PageID is:", _PageID)
		_ServiceName, tempSplitPageLinks = tempSplitPageLinks[len(tempSplitPageLinks)-1], tempSplitPageLinks[:len(tempSplitPageLinks)-1]
		// fmt.Println("_ServiceName is:", _ServiceName)
		_Prefix = strings.Join(tempSplitPageLinks, "/")
		// fmt.Println("_Prefix is:", _Prefix)
		// fmt.Println(element)

		if _, ok := SplitPageLinks[_ServiceName]; ok == false {
			SplitPageLinks[_ServiceName] = make([]map[string]string, 0)
		}
		SplitPageLinks[_ServiceName] = append(SplitPageLinks[_ServiceName], map[string]string{
			"pageID": _PageID,
			"prefix": _Prefix,
		})
	}
	// fmt.Println(SplitPageLinks)

	return SplitPageLinks
}
