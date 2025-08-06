package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"groupie-tracker-visualizations/project"
)

// BuildLocationDateList fetches the relation data and returns a sorted slice of LocDate.
// Dates are sorted chronologically by reformatting "DD-MM-YYYY" to "YYYYMMDD".
func BuildLocationDateList(idx int) []project.LocDate {
	type RelationResponse struct {
		Index []project.RelationEntry `json:"index"`
	}

	resp, err := http.Get(project.API + "/relation")
	if err != nil {
		fmt.Println("Error fetching relation data:", err)
		return nil
	}
	defer resp.Body.Close()

	var data RelationResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	if idx < 0 || idx >= len(data.Index) {
		return nil
	}
	relation := data.Index[idx]

	// 1. Extract and sort locations alphabetically
	locs := make([]string, 0, len(relation.DatesLocations))
	for loc := range relation.DatesLocations {
		locs = append(locs, loc)
	}
	sort.Strings(locs)

	// 2. Build and sort dates for each location
	var locDates []project.LocDate
	for _, loc := range locs {
		dates := relation.DatesLocations[loc]

		// Chronological sort without parsing: reformat to YYYYMMDD

		// 01 - 01 - 2001
		// 012  345  678910
		// 20010101 (YYYYMMDD)

		sort.Slice(dates, func(i, j int) bool {
			key := func(s string) string {
				return s[6:10] + s[3:5] + s[0:2]
			}
			return key(dates[i]) < key(dates[j])
		})

		locDates = append(locDates, project.LocDate{
			Location: loc,
			Dates:    dates,
		})
	}

	return locDates
}
