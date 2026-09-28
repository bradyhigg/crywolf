package statuspage

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getSummary(baseUrl string) (*Summary, error) {
	statusPageURL := fmt.Sprintf(statusPage,baseUrl)
	resp, err := http.Get(statusPageURL)
	if err != nil{
		return nil, err
	}

	var summary *Summary
	err = json.NewDecoder(resp.Body).Decode(&summary)
	if err != nil{
		return nil, err
	}
	return summary, nil
}

// GetStatus: 
// return status string, is incident, error
func GetStatus(baseUrl string, specificName string) (string, bool, error) {
	summary, err := getSummary(baseUrl)
	if err != nil {
		return "", false, err
	}

	// no current incidents
	if len(summary.Incidents) == 0 {
		return "no incident", false, nil
	}

	var component Component
	for _,c := range summary.Components {
		if c.Name == specificName {
			component = c
		}
	}

	for _,i := range summary.Incidents {
		if component.Id == i.Id {
			return i.Status, true, err
		}
	}
	// incident not related
	return "Incident unrelated to specific service", false, nil
}