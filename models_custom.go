package plex

import (
	"encoding/json"
	"fmt"
)

// PlexRating represents a rating object in the Plex API (found in GetMetadata responses)
type PlexRating struct {
	Image string  `json:"image"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// UnmarshalJSON implements custom unmarshaling for Metadata to handle polymorphic 'rating' field
func (m *Metadata) UnmarshalJSON(data []byte) error {
	// Create a type alias to avoid recursion
	type Alias Metadata

	// Create an auxiliary struct that overrides the Rating field
	aux := &struct {
		Rating interface{} `json:"rating"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	// Unmarshal into the auxiliary struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Handle the polymorphic Rating field
	if aux.Rating != nil {
		switch v := aux.Rating.(type) {
		case float64:
			m.Rating = v
		case []interface{}:
			// It's an array of ratings (GetMetadata response)
			// We need to parse this manually into PlexRating structs
			// And then decide what to put in m.Rating (to preserve backward compatibility)

			// Re-marshal the interface{} to bytes to unmarshal into []PlexRating
			// This is inefficient but safe. Alternatively we can type assert the map[string]interface{}

			// Optimization: Iterate the slice and manualy map if possible, but Unmarshal is cleaner code
			ratingBytes, err := json.Marshal(v)
			if err != nil {
				return fmt.Errorf("failed to marshal rating array: %v", err)
			}

			var ratings []PlexRating
			if err := json.Unmarshal(ratingBytes, &ratings); err != nil {
				return fmt.Errorf("failed to unmarshal rating array: %v", err)
			}

			m.Ratings = ratings

			// Populate m.Rating (float64) from the best available rating
			// Logic: Prioritize Rotten Tomatoes Critic, then Audience, then first available
			if len(ratings) > 0 {
				found := false
				for _, r := range ratings {
					if r.Type == "critic" {
						m.Rating = r.Value
						found = true
						break
					}
				}
				if !found {
					m.Rating = ratings[0].Value
				}
			}
		}
	}

	return nil
}
