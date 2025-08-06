package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"groupie-tracker-visualizations/project"
)

// fetchData performs an HTTP GET and returns the response body.
func fetchData(Type string) ([]byte, error) {
	url := project.API + "/" + Type
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned status %s", url, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}
	return data, nil
}

// fetchArtists retrieves and unmarshals the list of artists.
func fetchArtists() ([]project.Artist, error) {
	data, err := fetchData("artists")
	if err != nil {
		return nil, err
	}
	var artists []project.Artist
	if err := json.Unmarshal(data, &artists); err != nil {
		return nil, fmt.Errorf("unmarshal artists: %w", err)
	}
	return artists, nil
}

// fetchLocations retrieves location entries and replaces underscores with spaces.
// func fetchLocations() ([]project.LocationEntry, error) {
// 	data, err := fetchData("locations")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var resp struct {
// 		Index []project.LocationEntry `json:"index"`
// 	}
// 	if err := json.Unmarshal(data, &resp); err != nil {
// 		return nil, fmt.Errorf("unmarshal locations: %w", err)
// 	}
// 	for i := range resp.Index {
// 		for j, loc := range resp.Index[i].Locations {
// 			resp.Index[i].Locations[j] = strings.ReplaceAll(loc, "_", " ")
// 		}
// 	}
// 	return resp.Index, nil
// }

// fetchDates retrieves date entries for each artist.
// func fetchDates() ([]project.DateEntry, error) {
// 	data, err := fetchData("dates")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var resp struct {
// 		Index []project.DateEntry `json:"index"`
// 	}
// 	if err := json.Unmarshal(data, &resp); err != nil {
// 		return nil, fmt.Errorf("unmarshal dates: %w", err)
// 	}
// 	return resp.Index, nil
// }

// fetchRelations retrieves relationship data (if ever needed).
func fetchRelations() ([]project.RelationEntry, error) {
	data, err := fetchData("relation")
	if err != nil {
		return nil, err
	}
	var resp struct {
		Index []project.RelationEntry `json:"index"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal relations: %w", err)
	}
	return resp.Index, nil
}
