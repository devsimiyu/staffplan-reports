package model

type RouteStatsReponse struct {
	Total     int `json:"total"`
	Ongoing   int `json:"ongoing"`
	Completed int `json:"completed"`
}
