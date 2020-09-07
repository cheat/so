package api

// Response models a response from the Search API
type Response struct {
	Items []Item `json:"items"`
}
