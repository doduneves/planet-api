package models

import "time"

type Response struct {
	LastUpdate string    `json:"last_update"`
	Count      string    `json:"count"`
	Next       string    `json:"next"`
	Previous   string    `json:"previous"`
	Results    []Results `json:"results"`
}

type Results struct {
	Name           string   `json:"name"`
	RotationPeriod int      `json:"rotation_period"`
	OrbitalPeriod  int      `json:"orbital_period"`
	Diameter       int      `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   int      `json:"surface_water"`
	Population     int      `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Appearances    int      `json:"appearances"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	Url            string   `json:"url"`
}

func (resp Response) UpdateLastUpdate() {
	resp.LastUpdate = time.Now().Format("2006-01-02")
}

func (res Results) SetAppearances() {
	res.Appearances = len(res.Films)

}
