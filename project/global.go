package project

import (
	"embed"
)

const API = "https://groupietrackers.herokuapp.com/api"

//go:embed frontEnd/template/*.html frontEnd/static/*.css frontEnd/static/icons/*.svg
var StaticFS embed.FS

// PageData holds the list of all artists for the index template.
type PageData struct {
	Artists []Artist
}

// PageDataArtist holds one artist and their location-date pairs for the artist template.
type PageDataArtist struct {
	Artist      Artist
	LocDateList []LocDate
	PrevID      int
    NextID      int
}

// Artist represents the structure of each artist returned by the /api/artists endpoint.
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// LocationEntry maps each artist's locations from the /api/locations endpoint.
type LocationEntry struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// DateEntry maps each artist's dates from the /api/dates endpoint.
type DateEntry struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// RelationEntry maps the /api/relation endpoint.
type RelationEntry struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// LocDate pairs a location with its associated dates for rendering.
type LocDate struct {
	Location string   `json:"location"`
	Dates    []string `json:"dates"`
}
