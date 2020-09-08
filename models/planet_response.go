package models

type PlanetResponse struct {
	Links   Links    `json:"links"`
	Count   int      `json:"count"`
	Planets []Planet `json:"planets"`
}

type Links struct {
	Actual       string `json:"actual"`
	NextPage     string `json:"next_page"`
	PreviousPage string `json:"previous_page"`
}
