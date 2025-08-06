package helpers

import "groupie-tracker-visualizations/project"

// UnmarshalData loads artists, locations, dates, and relations from the API.
// It returns an error if any fetch or unmarshal fails.
func LoadDataSet() ([]project.Artist, []project.RelationEntry, error) {
	artists, err := fetchArtists()
	if err != nil {
		return nil, nil, err
	}
	// locations, err := fetchLocations()
	// if err != nil {
	// 	return nil, nil, nil, nil, err
	// }
	// dates, err := fetchDates()
	// if err != nil {
	// 	return nil, nil, nil, nil, err
	// }
	relations, err := fetchRelations()
	if err != nil {
		return nil, nil, err
	}
	return artists, relations, nil
}
