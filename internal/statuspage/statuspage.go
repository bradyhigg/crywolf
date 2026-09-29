package statuspage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
func GetStatus(baseUrl string, specificName string) (*Status, error) {
	summary, err := getSummary(baseUrl)
	if err != nil {
		return nil, err
	}

	returnStatus := Status{
		IsIncident: false,
		OtherIncient: false,
		isInvestigating: false,
		IncidentDocumentedTime: nil,
	}

	// no current incidents
	if len(summary.Incidents) == 0 {
		return &returnStatus, nil
	}

	var component Component
	for _,c := range summary.Components {
		if c.Name == specificName {
			component = c
		}
	}

	for _,i := range summary.Incidents {
		if component.Id == i.Id {
			returnStatus.IsIncident = true
			if strings.Contains(i.Status,"investigat"){
				returnStatus.isInvestigating = true
			}
			returnStatus.IncidentDocumentedTime = i.CreatedAt
			return &returnStatus, nil
		}
	}
	// incident not related
	returnStatus.OtherIncient = true
	return &returnStatus, nil
}