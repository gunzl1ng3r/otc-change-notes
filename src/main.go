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
	var ChangeNotes map[string]map[string]string

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
	var ReadingChangeNotes bool
	var InsideTableBody bool
	var currentKey string

	// Loop over returned HTML, do not register index
	for _, element := range bodyAsSlice {

		// The heading "Change History" marks the beginning of the interesting part
		if strings.Contains(element, ">Change History</h1>") {
			ReadingChangeNotes = true
			continue
		}
		// fmt.Println("ReadingChangeNotes is set to:", ReadingChangeNotes)
		if ReadingChangeNotes {
			// First occurrence of "<tbody>" after entering the Change Notes marks the beginning
			if strings.Contains(element, "<tbody>") {
				InsideTableBody = true
			}
			// First occurrence of "</tbody>" after entering the Change Notes marks the end
			if strings.Contains(element, "</tbody>") {
				ReadingChangeNotes = false
				break
			}
		}

		// fmt.Println("InsideTableBody is set to:", InsideTableBody)
		if InsideTableBody {
			re := regexp.MustCompile(`>[0-9]{4}-[0-9]{2}-[0-9]{2}<`)
			if len(re.Find([]byte(element))) > 0 {
				currentKey = (stripHTML(element))
				ChangeNotes[currentKey] = make(map[string]string)
				// fmt.Println(index, element)
				// output := stripHTML(element)
				// fmt.Println(index, output)
			} else {
				ChangeNotes[currentKey][stripHTML(element)] = ""
			}
		}
	}
	// fmt.Println(string(body))
	for key, value := range ChangeNotes {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func stripHTML(html string) string {
	re := regexp.MustCompile(`<.+?>`)
	return strings.TrimSpace(re.ReplaceAllString(html, ""))
}
