package main

import (
	"fmt"
	"sync"
)

func main() {
	documentationPages := getPages()

	var services []string
	services = make([]string, 0, len(documentationPages))

	for service := range documentationPages {
		services = append(services, service)
	}

	var waitGroup sync.WaitGroup
	var changeHistories map[string][]string
	changeHistories = make(map[string][]string, 0)

	for _, service := range services {
		waitGroup.Add(1)
		go func(service string) {
			defer waitGroup.Done()
			for _, value := range documentationPages[service] {
				_ChangeHistoryURL := "https://docs.otc.t-systems.com" + value["prefix"] + "/" + service + "/" + value["pageID"]
				_isChangeHistory := identifyChangeHistory(_ChangeHistoryURL)
				if _isChangeHistory {
					changeHistories[service] = append(changeHistories[service], _ChangeHistoryURL)
				}
			}
		}(service)
	}
	waitGroup.Wait()
	// fmt.Println(changeHistories)
	for key, value := range changeHistories {
		// fmt.Println(key + ":")
		// fmt.Println("  -", strings.Join(value, "\n"))
		for _, changeHistory := range value {
			fmt.Println(key)
			parseChangeHistory(changeHistory)
		}
	}
}
