package model

type Schema struct {
	ID         int    `json:"id"`
	Version    string `json:"version"`
	Definition string `json:"definition"`
}