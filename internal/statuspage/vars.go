package statuspage

import "time"


const statusPage = "%s/api/v2/summary.json"


type Summary struct {
    Components []Component `json:"components"`
    Incidents  []Incident  `json:"incidents"`
}

type Component struct {
	Id string `json:"id"`
	Name string `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Incident struct {
	Id string `json:"id"`
    Name string `json:"name"`
    Status string `json:"status"`
    Impact string `json:"impact"`
    CreatedAt *time.Time `json:"created_at"`
    StartedAt *time.Time `json:"started_at"`
    ResolvedAt *time.Time `json:"resolved_at"`
    Shortlink string `json:"shortlink"`
}

type Status struct {

}